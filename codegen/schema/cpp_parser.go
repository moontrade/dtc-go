package schema

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/yoheimuta/go-protoparser/v4/parser"
)

func parseError(lineIndex int, line, message string) error {
	return fmt.Errorf("line: %d, '%s' -> %s", lineIndex+1, line, message)
}

type CHeaderFile struct {
	Path             string
	Schema           *Schema
	currentNamespace *Namespace
	Alias            []*Alias
	AliasByName      map[string]*Alias
	Constants        []*Const
	ConstantsByName  map[string]*Const
	Enums            []*Enum
	EnumsByName      map[string]*Enum
	EnumOptions      map[string]*EnumOption
	Structs          []*Struct
	StructsByName    map[string]*Struct
}

func LoadSchemaFromCHeaders(docsPath, protoPath string, headerPaths ...string) (*Schema, error) {
	schema := NewSchema()
	if len(docsPath) > 0 {
		docsJson, err := os.ReadFile(docsPath)
		if err != nil {
			return nil, err
		}
		schema.Docs = &Docs{
			MessagesByName: make(map[string]*MessageDoc),
			TypesByName:    make(map[string]*TypeDoc),
		}
		if err = json.Unmarshal(docsJson, schema.Docs); err != nil {
			return nil, err
		}
	}
	err := schema.AddCHeaders(headerPaths...)
	if err != nil {
		return nil, err
	}
	if len(protoPath) > 0 {
		if err = schema.AddProto(protoPath); err != nil {
			return nil, err
		}
	}
	if err = schema.Validate(); err != nil {
		return nil, err
	}
	return schema, nil
}

func (schema *Schema) AddProto(path string) error {
	proto, err := ParseProto(path)
	if err != nil {
		return err
	}
LOOP:
	for _, v := range proto.ProtoBody {
		switch v := v.(type) {
		case *parser.Enum:
			enum := schema.EnumsByName[v.EnumName]
			if enum == nil {
				switch v.EnumName {
				case "DTCVersion":
				case "DTCMessageType":
				case "MessageSupportedEnum":
					continue LOOP
				default:
					return errors.New("could not locate enum named: " + v.EnumName)
				}
				continue LOOP
			}
			if err = enum.Bind(v); err != nil {
				return err
			}

		case *parser.Message:
			msg := schema.MessagesByName[v.MessageName]
			if msg == nil {
				msg = schema.MessagesByName["s_"+v.MessageName]
			}
			if msg == nil {
				switch v.MessageName {
				case "MarketDataUpdateBidAsk2":
					// ignore
					continue LOOP
				}
				fmt.Printf("WARN: protobuf message named '%s' does not have a match in DTCProtocol.h\n", v.MessageName)
				//return errors.New("could not locate struct named: " + v.MessageName)
				continue
			}
			if msg.Fixed != nil {
				if err = msg.Fixed.Bind(v); err != nil {
					return err
				}
			}
			if msg.VLS != nil {
				if err = msg.VLS.Bind(v); err != nil {
					return err
				}
			}

		case *parser.Package:
		}
	}
	return nil
}

func (schema *Schema) AddCHeaders(paths ...string) error {
	for _, path := range paths {
		if _, err := schema.AddCHeader(path); err != nil {
			return err
		}
	}
	return nil
}

func (schema *Schema) AddCHeader(path string) (*CHeaderFile, error) {
	contents_, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	contents := string(contents_)
	var (
		file = &CHeaderFile{
			Path:             path,
			Schema:           schema,
			currentNamespace: schema.GetNamespace(""), // Global namespace
			AliasByName:      make(map[string]*Alias),
			ConstantsByName:  make(map[string]*Const),
			EnumsByName:      make(map[string]*Enum),
			StructsByName:    make(map[string]*Struct),
			EnumOptions:      make(map[string]*EnumOption),
		}
		packStack    []int
		blockName    string
		blockExtends Type
		blockType    Kind
		block        []string
		state        = 0 // 0 = root, 1 == expect '{', 2 == expect '};'
		curlyCount   = 0
		nonStandard  = !strings.Contains(path, "DTCProtocol.h") && !strings.Contains(path, "DTCProtocolVLS.h")
	)

	lines := strings.Split(contents, "\n")
	for lineIndex, line := range lines {
		line = strings.TrimSpace(line)

		switch state {
		case 1:
			index := strings.Index(line, "{")
			if index == -1 {
				continue
			}
			line = strings.TrimSpace(line[index+1:])
			if len(line) > 0 {
				block = append(block, line)
			}
			state = 2
			curlyCount = 0
			continue
		case 2:
			for _, c := range line {
				switch c {
				case '{':
					curlyCount++
				case '}':
					curlyCount--
				}
			}

			if curlyCount == -1 {
				//if line == "};" && curlyCount == 0 {
				switch blockType {
				case KindStruct:
					pack := 8
					if len(packStack) > 0 {
						pack = packStack[len(packStack)-1]
					}
					if blockName == "s_MessageBase" {
						state = 0
						continue
					}
					var s *Struct
					s, err = file.parseStruct(pack, blockName, block)
					if s == nil {
						state = 0
						continue
					}
					if err != nil {
						return nil, parseError(lineIndex, line, err.Error())
					}

					message := schema.MessagesByName[s.Name]
					if message == nil {
						message = &Message{NonStandard: nonStandard}
						schema.MessagesByName[s.Name] = message
						schema.Messages = append(schema.Messages, message)
					}
					switch s.Namespace {
					case NamespaceKindFixed:
						if message.Fixed != nil {
							return nil, parseError(lineIndex, line, "duplicated fixed struct named: "+s.Name)
						}
						message.Fixed = s
					case NamespaceKindVLS:
						if message.VLS != nil {
							return nil, parseError(lineIndex, line, "duplicated vls struct named: "+s.Name)
						}

						message.VLS = s
					default:
						// Non-Standard
						return nil, parseError(lineIndex, line, "unknown namespace for struct: "+s.Namespace.String())
					}

				case KindEnum:
					_, err = file.parseEnum(blockName, blockExtends.Kind, block)
					if err != nil {
						return nil, parseError(lineIndex, line, err.Error())
					}
				}
				state = 0
				continue
			} else {
				block = append(block, line)
				continue
			}
		}

		switch {
		case strings.HasPrefix(line, "namespace"):
			line = strings.TrimSpace(line[len("namespace"):])
			index := strings.LastIndex(line, "{")
			namespaceName := ""
			if index > -1 {
				namespaceName = strings.TrimSpace(line[0:index])
			} else {
				index = strings.Index(line, "/")
				if index > -1 {
					namespaceName = strings.TrimSpace(line[0:index])
				} else {
					namespaceName = strings.TrimSpace(line)
				}
			}
			file.currentNamespace = schema.GetNamespace(namespaceName)

		case strings.HasPrefix(line, "const"):
			line = strings.TrimSpace(line[len("const"):])
			_, err = file.parseConst(line)
			if err != nil {
				return nil, parseError(lineIndex, line, err.Error())
			}

		case strings.HasPrefix(line, "enum"):
			line = strings.TrimSpace(line[len("enum"):])

			blockType = KindEnum
			block = nil

			index := strings.Index(line, ":")
			if index == -1 {
				return nil, parseError(lineIndex, line, "expected ':' and data type")
			}
			blockName = strings.TrimSpace(line[0:index])
			line = strings.TrimSpace(line[index+1:])

			curlyCount = 0
			index = strings.Index(line, "{")
			if index > -1 {
				blockExtends = file.typeOf(strings.TrimSpace(line[0:index]))
				state = 2
			} else {
				blockExtends = file.typeOf(strings.TrimSpace(line))
				state = 1
			}

		case strings.HasPrefix(line, "#pragma"):
			line = strings.TrimSpace(line[len("#pragma"):])
			if !strings.HasPrefix(line, "pack(") {
				continue
			}
			line = line[len("pack("):]
			line = strings.ReplaceAll(line, ")", "")
			line = strings.TrimSpace(line)
			parts := strings.Split(line, ",")
			switch strings.TrimSpace(parts[0]) {
			case "push":
				align, err := strconv.ParseInt(strings.TrimSpace(parts[1]), 10, 64)
				if err != nil {
					return nil, parseError(lineIndex, line, "invalid #pragma pack(push, #ERROR_VALUE#) "+strings.TrimSpace(parts[1]))
				}
				packStack = append(packStack, int(align))
			case "pop":
				if len(packStack) > 0 {
					packStack = packStack[0 : len(packStack)-1]
				}
			}

		case strings.HasPrefix(line, "struct"):
			line = strings.TrimSpace(line[len("struct"):])
			blockType = KindStruct
			block = nil
			blockName = ""
			blockExtends = Type{}
			curlyCount = 0

			index := strings.Index(line, "{")
			if index > -1 {
				blockName = strings.TrimSpace(line[0:index])
				state = 2
			} else {
				blockName = line
				state = 1
			}
			index = strings.Index(blockName, "//")
			if index > -1 {
				blockName = strings.TrimSpace(blockName[0:index])
			}
			index = strings.Index(blockName, "/*")
			if index > -1 {
				blockName = strings.TrimSpace(blockName[0:index])
			}

		case strings.HasPrefix(line, "typedef"), strings.HasPrefix(line, "using"):
			if strings.HasPrefix(line, "using") {
				line = strings.TrimSpace(line[len("using"):])
			} else {
				line = strings.TrimSpace(line[len("typedef"):])
			}
			index := strings.Index(line, "//")
			if index > -1 {
				line = strings.TrimSpace(line[:index])
			}
			if strings.HasSuffix(line, ";") {
				line = strings.TrimSpace(line[0 : len(line)-1])
			}
			name := line
			base := line
			index = strings.LastIndex(line, " ")
			if index > -1 {
				base = strings.TrimSpace(line[0:index])
				name = strings.TrimSpace(line[index+1:])
			}

			if strings.HasPrefix(name, "std::") {
				continue
			}
			switch name {
			case "uint8_t", "uint16_t", "uint32_t", "uint64_t", "int8_t", "int16_t", "int32_t", "int64_t":
				continue
			}

			if strings.Contains(name, "vls_t") || strings.Contains(name, "uint8_t") {
				continue
			}
			alias := &Alias{
				Namespace: file.currentNamespace.Kind,
				Name:      name,
				Type:      file.typeOf(base),
			}
			if alias.Type.Kind == KindUnknown {
				alias.Type = file.typeOf(name)
			}
			if alias.Type.Kind == KindUnknown {
				continue
			}
			if file.Schema.Docs != nil {
				alias.Doc = file.Schema.Docs.TypeNamed(alias.Name)
			}
			file.Schema.Aliases = append(file.Schema.Aliases, alias)
			if file.Schema.AliasesByMame[alias.Name] != nil {
				return nil, parseError(lineIndex, line, "duplicate alias name: "+alias.Name)
			}
			file.Alias = append(file.Alias, alias)
			file.AliasByName[alias.Name] = alias
			file.currentNamespace.Alias = append(file.currentNamespace.Alias, alias)
			file.currentNamespace.AliasByName[alias.Name] = alias
		}
	}

	for _, s := range file.Structs {
		s.Layout()
	}

	return file, nil
}

func (f *CHeaderFile) parseConst(line string) (*Const, error) {
	line = strings.TrimSpace(line)
	var (
		constant = &Const{
			Namespace: f.currentNamespace,
		}
		index = strings.Index(line, "//")
		err   error
	)
	if index > -1 {
		constant.Comment = strings.TrimSpace(line[index:])
		line = strings.TrimSpace(line[0:index])
	}
	index = strings.Index(line, "=")
	if index == -1 {
		return nil, errors.New("constant missing '=' expression")
	}
	constant.Value, err = f.parseValue(strings.TrimSpace(line[index+1:]))
	if err != nil {
		return nil, err
	}
	line = strings.TrimSpace(line[0:index])

	line = strings.ReplaceAll(line, "\t", " ")
	index = strings.LastIndex(line, " ")
	constant.Name = strings.TrimSpace(line[index+1:])
	constant.Type = f.typeOf(strings.TrimSpace(line[0:index]))

	if f.ConstantsByName[constant.Name] != nil {
		return nil, errors.New("duplicate constant declared: " + line)
	}
	if f.currentNamespace.Schema.ConstantsByName[constant.Name] != nil {
		// Use existing constant
		f.currentNamespace.Schema.DuplicateConstants = append(f.currentNamespace.Schema.DuplicateConstants, constant)
		constant = f.currentNamespace.Schema.ConstantsByName[constant.Name]
	} else {
		f.currentNamespace.Schema.Constants = append(f.currentNamespace.Schema.Constants, constant)
		f.currentNamespace.Schema.ConstantsByName[constant.Name] = constant
	}

	f.Constants = append(f.Constants, constant)
	f.ConstantsByName[constant.Name] = constant
	f.currentNamespace.Constants = append(f.currentNamespace.Constants, constant)
	f.currentNamespace.ConstantsByName[constant.Name] = constant

	return constant, nil
}

func setLengthAndAlign(t *Type) {
	kind := t.Kind
	if kind == KindUnion {
		for _, field := range t.Union.Fields {
			setLengthAndAlign(&field.Type)
			if field.Type.Size > t.Size {
				t.Size = field.Type.Size
			}
			if field.Type.Align > t.Align {
				t.Align = field.Type.Align
			}
		}
		return
	}
	if kind == KindEnum {
		kind = t.Enum.Type
	} else if kind == KindAlias {
		kind = t.Alias.Type.Kind
	}
	switch kind {
	case KindUnknown:
	case KindInt8, KindUint8:
		t.Align = 1
		t.Size = 1
	case KindInt16, KindUint16:
		t.Align = 2
		t.Size = 2
	case KindInt32, KindUint32:
		t.Align = 4
		t.Size = 4
	case KindInt64, KindUint64:
		t.Align = 8
		t.Size = 8
	case KindFloat32:
		t.Align = 4
		t.Size = 4
	case KindFloat64:
		t.Align = 8
		t.Size = 8
	case KindStringFixed:
		t.Align = 1
	case KindStringVLS:
		t.Align = 2
		t.Size = 4
	case KindAlias:
		// impossible
	case KindEnum:
		// impossible
	case KindStruct:
		if t.Struct != nil {
			t.Align = t.Struct.Size
		}
	}
}

func (f *CHeaderFile) typeOf(value string) Type {
	if strings.HasPrefix(value, "std::") {
		value = value[len("std::"):]
	}
	value = strings.TrimSpace(value)
	switch value {
	case "char":
		return Type{Kind: KindStringFixed}
	case "bool":
		return Type{Kind: KindBool, Align: 1, Size: 1}
	case "int8_t", "schar":
		return Type{Kind: KindInt8, Align: 1, Size: 1}
	case "uint8_t", "unsigned char":
		return Type{Kind: KindUint8, Align: 1, Size: 1}
	case "int16_t", "short":
		return Type{Kind: KindInt16, Align: 2, Size: 2}
	case "uint16_t", "unsigned short":
		return Type{Kind: KindUint16, Align: 2, Size: 2}
	case "int32_t", "int":
		return Type{Kind: KindInt32, Align: 4, Size: 4}
	case "uint32_t", "unsigned int", "unsigned long":
		return Type{Kind: KindUint32, Align: 4, Size: 4}
	case "int64_t", "long long":
		return Type{Kind: KindInt64, Align: 8, Size: 8}
	case "uint64_t", "unsigned long long":
		return Type{Kind: KindUint64, Align: 8, Size: 8}
	case "float":
		return Type{Kind: KindFloat32, Align: 4, Size: 4}
	case "double":
		return Type{Kind: KindFloat64, Align: 8, Size: 8}
	case "DTC_VLS::vls_t", "vls_t":
		return Type{Kind: KindStringVLS, Align: 2, Size: 4}
	}

	namespace := f.currentNamespace
	index := strings.LastIndex(value, "::")
	if index > -1 {
		namespaceName := strings.TrimSpace(value[0:index])
		namespace = f.Schema.GetNamespace(namespaceName)
		value = strings.TrimSpace(value[index+2:])
	}

	enum := namespace.EnumsByName[value]
	if enum != nil {
		t := Type{
			Namespace: enum.Namespace,
			Kind:      KindEnum,
			Enum:      enum,
		}
		setLengthAndAlign(&t)
		return t
	}

	message := namespace.StructsByName[value]
	if message != nil {
		t := Type{
			Namespace: message.Namespace,
			Kind:      KindStruct,
			Struct:    message,
		}
		setLengthAndAlign(&t)
		return t
	}

	alias := namespace.AliasByName[value]
	if alias != nil {
		t := Type{
			Namespace: alias.Namespace,
			Kind:      KindAlias,
			Alias:     alias,
		}
		setLengthAndAlign(&t)
		return t
	}

	return Type{Kind: KindUnknown}
}

func isFunc(name, value string) bool {
	index := strings.Index(value, "(")
	if index > -1 {
		value = strings.TrimSpace(value[0:index])
	}
	index = strings.LastIndex(value, " ")
	if index == -1 {
		index = strings.LastIndex(value, "\t")
		if index == -1 {
			index = strings.LastIndex(value, "\n")
		}
	}
	if index > -1 {
		value = strings.TrimSpace(value[index+1:])
	}
	return name == value
}

func (f *CHeaderFile) parseStruct(pack int, name string, lines []string) (*Struct, error) {
	var (
		message = &Struct{
			Namespace:    f.currentNamespace.Kind,
			Name:         name,
			MaxAlign:     pack,
			VLS:          strings.Contains(f.currentNamespace.Name, "VLS"),
			Fields:       nil,
			FieldsByName: make(map[string]*Field),
		}
		funcLines  []string
		funcDecl   string
		state      = 0 // 0 = begin, 1 = waiting for end '}'
		err        error
		curlyCount = 0
		comment    = ""
	)
	_ = comment

	for i, line := range lines {
		line = strings.TrimSpace(line)
		index := strings.Index(line, "//")
		if index > -1 {
			comment = strings.TrimSpace(line[index:])
			line = strings.TrimSpace(line[0:index])
		}
		_ = i

		if len(line) == 0 {
			continue
		}

		if state == 1 {
			for _, c := range line {
				switch c {
				case '{':
					curlyCount++
				case '}':
					curlyCount--
				}
			}
			if curlyCount == -1 {
				state = 0

				var (
					isClear = isFunc("clear", strings.ToLower(funcDecl))
					isCtor  = isFunc(message.Name, funcDecl)
				)

				if isClear || isCtor {
					// Process default field values
					for _, line = range funcLines {
						parts := strings.Split(line, "=")
						left := strings.TrimSpace(parts[0])
						right := ""
						field := message.FieldsByName[left]
						if field == nil {
							if strings.HasPrefix(left, "Set") {
								left = strings.TrimSpace(left[3:])
								index = strings.Index(left, "(")
								if index == -1 {
									continue
								}
								field = message.FieldsByName[left[0:index]]
								if field == nil {
									continue
								}
								right = left[index+1:]
								index = strings.LastIndex(right, ")")
								if index == -1 {
									continue
								}
								right = strings.TrimSpace(right[0:index])
							} else {
								continue
							}
						} else {
							right = strings.TrimSpace(parts[1])
						}
						field.ClearExpression = right
						if strings.HasSuffix(field.ClearExpression, ";") {
							field.ClearExpression = strings.TrimSpace(field.ClearExpression[0 : len(field.ClearExpression)-1])
						}
						field.Initial, err = f.parseInitValue(message, field.ClearExpression)
						if err != nil {
							return nil, err
						}
						if field.Initial == nil {
							return nil, errors.New("could not parse init clear expression: " + field.ClearExpression)
						}
						if field.Initial.Type == ValueTypeBool {
							switch field.Type.Kind {
							case KindInt8, KindUint8:
								field.Type.Kind = KindBool
							default:
								field.Initial.Type = ValueTypeInt
							}
						}
					}
				} else if funcDecl == "union" {
					union := Type{
						Kind:  KindUnion,
						Union: &Union{},
					}

					for _, line := range funcLines {
						// Should be field declaration
						field, err := f.parseField(message, line)
						if err != nil {
							return nil, err
						}

						union.Union.Fields = append(union.Union.Fields, field)

						if field.Type.Size > union.Size {
							union.Size = field.Type.Size
						}
						if field.Type.Align > union.Align {
							union.Align = field.Type.Align
						}

						message.FieldsByName[field.Name] = field
					}

					message.Fields = append(message.Fields, &Field{
						Name: "",
						Type: union,
					})
				}

				funcDecl = ""
			} else {
				funcLines = append(funcLines, line)
			}
			continue
		}

		if !strings.HasSuffix(line, ";") {
			index = strings.Index(line, "{")
			if index > -1 {
				funcLines = nil
				begin := strings.TrimSpace(line[0:index])
				if len(begin) > 0 {
					funcDecl = begin
				}
				line = strings.TrimSpace(line[index+1:])
				if len(line) > 0 {
					funcLines = append(funcLines, line)
				}
				curlyCount = 0
				state = 1
			} else if len(line) > 0 {
				funcDecl = line
			}

			continue
		}

		line = strings.TrimSpace(line[0 : len(line)-1])
		if strings.HasSuffix(line, "const") || strings.HasSuffix(line, "noexcept") {
			continue
		}

		if strings.HasSuffix(line, ")") {
			if !strings.Contains(line, "sizeof(") {
				continue
			}
		}

		// Should be field declaration
		field, err := f.parseField(message, line)
		if err != nil {
			return nil, err
		}

		if field != nil {
			message.Fields = append(message.Fields, field)
			message.FieldsByName[field.Name] = field
		}
	}

	if f.Schema.Docs != nil {
		message.Doc = f.Schema.Docs.MessageNamed(message.Name)
	}

	if len(message.Fields) == 0 || strings.ToLower(message.Fields[0].Name) != "size" {
		return nil, nil
	}

	for _, field := range message.Fields {
		if message.Doc != nil {
			field.Doc = message.Doc.FieldNamed(field.Name)
		}
		if field.Type.Kind == KindStringVLS {
			message.VLS = true
		}
		if field.Initial == nil && len(field.InitExpression) > 0 {
			field.Initial, err = f.parseInitValue(message, field.InitExpression)
			if err != nil {
				return nil, err
			}
		}

		if field.Type.Union != nil {
			for _, field := range field.Type.Union.Fields {
				if message.Doc != nil {
					field.Doc = message.Doc.FieldNamed(field.Name)
				}
				if field.Type.Kind == KindStringVLS {
					message.VLS = true
				}
				if field.Initial == nil && len(field.InitExpression) > 0 {
					field.Initial, err = f.parseInitValue(message, field.InitExpression)
					if err != nil {
						return nil, err
					}
				}
			}
		}

		if strings.ToLower(field.Name) == "type" && field.Initial != nil {
			if field.Initial.Int > 0 {
				message.Type = uint16(field.Initial.Int)
			} else if field.Initial.Uint > 0 {
				message.Type = uint16(field.Initial.Uint)
			} else if field.Initial.Const != nil {
				if field.Initial.Const.Value.Int > 0 {
					message.Type = uint16(field.Initial.Const.Value.Int)
				} else if field.Initial.Const.Value.Uint > 0 {
					message.Type = uint16(field.Initial.Const.Value.Uint)
				}
			}
		}
	}

	f.Structs = append(f.Structs, message)
	f.StructsByName[message.Name] = message
	f.currentNamespace.Structs = append(f.currentNamespace.Structs, message)
	f.currentNamespace.StructsByName[message.Name] = message

	return message, nil
}

func (f *CHeaderFile) parseField(s *Struct, line string) (*Field, error) {
	line = strings.TrimSpace(line)
	index := strings.Index(line, "//")
	comment := ""
	if index > -1 {
		comment = strings.TrimSpace(line[index:])
		line = strings.TrimSpace(line[0:index])
	}
	_ = comment

	if strings.HasSuffix(line, ";") {
		line = strings.TrimSpace(line[0 : len(line)-1])
	}

	var (
		expression = ""
		field      = &Field{Struct: s}
		err        error
	)
	index = strings.Index(line, "=")
	if index > -1 {
		expression = strings.TrimSpace(line[index+1:])
		line = strings.TrimSpace(line[0:index])
	}

	line = strings.ReplaceAll(line, "\t", " ")
	index = strings.LastIndex(line, " ")
	if index == -1 {
		return nil, errors.New("expected a space between data type and name")
	}

	field.Type = f.typeOf(strings.TrimSpace(line[0:index]))
	field.Name = strings.TrimSpace(line[index+1:])

	index = strings.Index(field.Name, "[")
	if index > -1 {
		endIndex := strings.Index(field.Name, "]")
		if index >= endIndex {
			return nil, errors.New("invalid fixed string declaration")
		}
		lengthExpression := strings.TrimSpace(field.Name[index+1 : endIndex])
		field.Name = strings.TrimSpace(field.Name[0:index])

		length, err := strconv.ParseInt(lengthExpression, 10, 64)
		if err != nil {
			value := f.findConstValue(lengthExpression)
			if value == nil || value.Const == nil {
				return nil, errors.New("could not determine fixed string length")
			}
			field.Type.SizeConst = value.Const
			field.Type.Size = int(value.Const.Value.Int)
		} else {
			field.Type.Size = int(length)
		}
	}

	if field.Name == "" {
		return nil, nil
	}

	field.InitExpression = expression

	if len(expression) > 0 {
		field.Initial, err = f.parseValue(expression)
		if err != nil {
			return nil, err
		}
	}

	switch field.Type.Kind {
	case KindInt8, KindUint8:
		if nameIsBool(field.Name) {
			if field.Initial != nil {
				switch field.Initial.Type {
				case ValueTypeBool:
					field.Type.Kind = KindBool
				case ValueTypeInt:
					switch field.Initial.Int {
					case 0:
						field.Type.Kind = KindBool
						field.Initial.Type = ValueTypeBool
					case 1:
						field.Type.Kind = KindBool
						field.Initial.Type = ValueTypeBool
					}
				case ValueTypeUint:
					switch field.Initial.Uint {
					case 0:
						field.Type.Kind = KindBool
						field.Initial.Type = ValueTypeBool
						field.Initial.Int = 0
					case 1:
						field.Type.Kind = KindBool
						field.Initial.Type = ValueTypeBool
						field.Initial.Int = 1
					}
				}
			} else {
				field.Type.Kind = KindBool
			}
		}
	}

	return field, nil
}

func nameIsBool(name string) bool {
	strings.TrimSpace(name)
	if strings.HasPrefix(name, "m_") {
		name = strings.TrimSpace(name[2:])
	}
	return isCamelCaseWord("Is", name) ||
		isCamelCaseWord("Supported", name) ||
		isCamelCaseWord("Use", name)
}

func isCamelCaseWord(word, in string) bool {
	index := strings.Index(in, word)
	if index == -1 {
		return false
	}
	in = in[index+len(word):]
	if len(in) == 0 {
		return true
	}
	return in[0:1] == "_" || strings.ToLower(in[0:1]) != in[0:1]
}

func (f *CHeaderFile) findConstValue(str string) *Value {
	index := strings.Index(str, "::")
	namespace := f.currentNamespace
	if index > -1 {
		namespace = f.Schema.GetNamespace(strings.TrimSpace(str[0:index]))
		str = strings.TrimSpace(str[index+2:])
	}
	constant := f.currentNamespace.Schema.ConstantsByName[str]
	if constant != nil {
		return &Value{
			Namespace: namespace.Kind,
			Type:      ValueTypeConst,
			Int:       constant.Value.Int,
			Float:     constant.Value.Float,
			Const:     constant,
		}
	}

	option := f.currentNamespace.Schema.EnumOptionsByName[str]
	if option != nil {
		return &Value{
			Int:        option.Value,
			Type:       ValueTypeEnumOption,
			EnumOption: option,
		}
	}

	return nil
}

func (f *CHeaderFile) parseInitValue(s *Struct, str string) (*Value, error) {
	if strings.HasSuffix(str, ";") {
		str = strings.TrimSpace(str[0 : len(str)-1])
	}
	field := s.FieldsByName[str]
	if field != nil {
		return field.Initial, nil
	}
	v, err := f.parseValue(str)
	if err != nil {
		return nil, err
	}
	if v != nil && v.Type == ValueTypeSizeof {
		v.Struct = s
	}
	return v, nil
}

func (f *CHeaderFile) parseValue(str string) (*Value, error) {
	str = strings.TrimSpace(str)
	if len(str) > 1 && strings.HasSuffix(str, ";") {
		str = strings.TrimSpace(str[0 : len(str)-1])
	}

	index := strings.Index(str, "::")
	if index > -1 {
		return f.findConstValue(str), nil
	}

	if str == "{}" {
		return &Value{
			Type: ValueTypeString,
			Str:  "",
		}, nil
	}

	index = strings.Index(str, "sizeof(")
	if index > -1 {
		str = strings.TrimSpace(str[index+len("sizeof("):])
		index = strings.Index(str, ")")
		if index > -1 {
			str = strings.TrimSpace(str[0:index])
		}
		return &Value{
			Type:   ValueTypeSizeof,
			Sizeof: str,
		}, nil
	}

	if str == "FLT_MAX" {
		return &Value{
			Type:  ValueTypeFloat32Max,
			Float: math.MaxFloat32,
		}, nil
	}
	if str == "DBL_MAX" {
		return &Value{
			Type:  ValueTypeFloat64Max,
			Float: math.MaxFloat64,
		}, nil
	}

	/*
		CHAR_BIT	Number of bits in a char object (byte)	8 or greater*
		SCHAR_MIN	Minimum value for an object of type signed char	-127 (-27+1) or less*
		SCHAR_MAX	Maximum value for an object of type signed char	127 (27-1) or greater*
		UCHAR_MAX	Maximum value for an object of type unsigned char	255 (28-1) or greater*
		CHAR_MIN	Minimum value for an object of type char	either SCHAR_MIN or 0
		CHAR_MAX	Maximum value for an object of type char	either SCHAR_MAX or UCHAR_MAX
		MB_LEN_MAX	Maximum number of bytes in a multibyte character, for any locale	1 or greater*
		SHRT_MIN	Minimum value for an object of type short int	-32767 (-215+1) or less*
		SHRT_MAX	Maximum value for an object of type short int	32767 (215-1) or greater*
		USHRT_MAX	Maximum value for an object of type unsigned short int	65535 (216-1) or greater*
		INT_MIN	Minimum value for an object of type int	-32767 (-215+1) or less*
		INT_MAX	Maximum value for an object of type int	32767 (215-1) or greater*
		UINT_MAX	Maximum value for an object of type unsigned int	65535 (216-1) or greater*
		LONG_MIN	Minimum value for an object of type long int	-2147483647 (-231+1) or less*
		LONG_MAX	Maximum value for an object of type long int	2147483647 (231-1) or greater*
		ULONG_MAX	Maximum value for an object of type unsigned long int	4294967295 (232-1) or greater*
		LLONG_MIN	Minimum value for an object of type long long int	-9223372036854775807 (-263+1) or less*
		LLONG_MAX	Maximum value for an object of type long long int	9223372036854775807 (263-1) or greater*
		ULLONG_MAX	Maximum value for an object of type unsigned long long int	18446744073709551615 (264-1) or greater*
	*/
	switch str {
	case "true":
		return &Value{
			Type: ValueTypeBool,
			Int:  1,
		}, nil
	case "false":
		return &Value{
			Type: ValueTypeBool,
			Int:  0,
		}, nil
	case "SCHAR_MIN":
		return &Value{
			Type: ValueTypeInt,
			Int:  math.MinInt8,
		}, nil
	case "SCHAR_MAX":
		return &Value{
			Type: ValueTypeInt,
			Int:  math.MaxInt8,
		}, nil
	case "UCHAR_MAX", "CHAR_MAX":
		return &Value{
			Type: ValueTypeUint,
			Uint: math.MaxUint8,
		}, nil
	case "CHAR_MIN":
		return &Value{
			Type: ValueTypeUint,
			Uint: 0,
		}, nil
	case "SHRT_MIN", "INT_MIN":
		return &Value{
			Type: ValueTypeInt,
			Int:  math.MinInt16,
		}, nil
	case "SHRT_MAX", "INT_MAX":
		return &Value{
			Type: ValueTypeInt,
			Int:  math.MaxInt16,
		}, nil
	case "USHRT_MAX", "UINT_MAX":
		return &Value{
			Type: ValueTypeUint,
			Uint: math.MaxUint16,
		}, nil
	case "LONG_MIN":
		return &Value{
			Type: ValueTypeInt,
			Int:  math.MinInt32,
		}, nil
	case "LONG_MAX":
		return &Value{
			Type: ValueTypeInt,
			Int:  math.MaxInt32,
		}, nil
	case "ULONG_MAX":
		return &Value{
			Type: ValueTypeUint,
			Uint: math.MaxUint32,
		}, nil
	case "LLONG_MIN":
		return &Value{
			Type: ValueTypeInt,
			Int:  math.MinInt64,
		}, nil
	case "LLONG_MAX":
		return &Value{
			Type: ValueTypeInt,
			Int:  math.MaxInt64,
		}, nil
	case "ULLONG_MAX":
		return &Value{
			Type: ValueTypeUint,
			Uint: math.MaxUint64,
		}, nil
	case "FLT_MAX":
		return &Value{
			Type:  ValueTypeFloat32Max,
			Float: math.MaxFloat32,
		}, nil
	case "DBL_MAX":
		return &Value{
			Type:  ValueTypeFloat64Max,
			Float: math.MaxFloat64,
		}, nil
	}

	//if strings.Contains(str, "_MIN") {
	//}
	//if strings.Contains(str, "_MAX") {
	//}

	if strings.Index(str, "\"") > -1 {
		return &Value{
			Type: ValueTypeString,
			Str:  strings.ReplaceAll(str, "\"", ""),
		}, nil
	}
	if strings.Index(str, ".") > -1 {
		if strings.HasSuffix(str, "f") {
			str = str[0 : len(str)-1]
		}
		number, err := strconv.ParseFloat(str, 64)
		if err == nil {
			return &Value{
				Type:  ValueTypeFloat,
				Float: number,
			}, nil
		}
	}
	number, err := strconv.ParseInt(str, 10, 64)
	if err == nil {
		return &Value{
			Type: ValueTypeInt,
			Int:  number,
		}, nil
	}

	return f.findConstValue(str), nil
}

func (f *CHeaderFile) parseEnum(name string, kind Kind, lines []string) (*Enum, error) {
	var (
		enum = &Enum{
			Namespace:     f.currentNamespace.Kind,
			Name:          name,
			Type:          kind,
			OptionsByName: make(map[string]*EnumOption),
		}
		options = strings.Split(strings.Join(lines, ""), ",")
	)

	if f.Schema.Docs != nil {
		enum.Doc = f.Schema.Docs.TypeNamed(enum.Name)
	}

	for _, option := range options {
		option = strings.TrimSpace(option)
		var (
			err   error
			opt   = &EnumOption{Enum: enum}
			index = strings.Index(option, "//")
		)
		if index == -1 {
			index = strings.Index(option, "/*")
		}
		if index > -1 {
			opt.Comment = strings.TrimSpace(option[index:])
			option = strings.TrimSpace(option[0:index])
		}

		index = strings.Index(option, "=")

		if index == -1 {
			opt.Name = strings.TrimSpace(option)
			// Find previous value
			if len(enum.Options) == 0 {
				opt.Value = 0
			} else {
				opt.Value = enum.Options[len(enum.Options)-1].Value + 1
			}
		} else {
			opt.Value, err = strconv.ParseInt(strings.TrimSpace(option[index+1:]), 10, 64)
			if err != nil {
				return nil, errors.New("invalid enum option declaration: " + option + " : " + err.Error())
			}
			opt.Name = strings.TrimSpace(option[0:index])
		}

		if enum.Doc != nil {
			opt.Doc = enum.Doc.OptionNamed(opt.Name)
		}

		if enum.OptionsByName[opt.Name] != nil {
			return nil, errors.New("duplicate enum option declaration: " + option)
		}
		if f.currentNamespace.EnumOptions[opt.Name] != nil {
			return nil, errors.New("file: duplicate enum option name: " + option)
		}

		enum.Options = append(enum.Options, opt)
		enum.OptionsByName[opt.Name] = opt
	}

	if f.currentNamespace.Schema.EnumsByName[enum.Name] != nil {
		// Use existing enum
		f.Schema.DuplicateEnums = append(f.Schema.DuplicateEnums, enum)
		enum = f.Schema.EnumsByName[enum.Name]
	} else {
		f.Schema.Enums = append(f.Schema.Enums, enum)
		f.Schema.EnumsByName[enum.Name] = enum

		for _, opt := range enum.Options {
			f.EnumOptions[opt.Name] = opt
			f.currentNamespace.EnumOptions[opt.Name] = opt

			if f.Schema.EnumOptionsByName[opt.Name] != nil {
				return nil, errors.New("duplicate enum option name used: " + opt.Name)
			}
			f.Schema.EnumOptionsByName[opt.Name] = opt
		}
	}
	f.Enums = append(f.Enums, enum)
	f.EnumsByName[enum.Name] = enum
	f.currentNamespace.Enums = append(f.currentNamespace.Enums, enum)
	f.currentNamespace.EnumsByName[enum.Name] = enum

	return enum, nil
}
