package golang

import (
	"fmt"
	"github.com/moontrade/dtc-go/codegen/schema"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

type Generator struct {
	config          *Config
	schema          *schema.Schema
	packageName     string
	all             map[string]interface{}
	aliases         []*Alias
	aliasesByName   map[string]*Alias
	constants       []*Constant
	constantsByName map[string]*Constant
	enums           []*Enum
	enumsByName     map[string]*Enum
	messages        []*Message
	messagesByName  map[string]*Message
	messagesByType  map[uint16]*Message
}

func NewGenerator(config *Config, s *schema.Schema) (*Generator, error) {
	generator := &Generator{
		config:          config,
		schema:          s,
		packageName:     filepath.Base(config.RootPackage),
		all:             make(map[string]interface{}),
		aliasesByName:   make(map[string]*Alias),
		enumsByName:     make(map[string]*Enum),
		constantsByName: make(map[string]*Constant),
		messagesByName:  make(map[string]*Message),
		messagesByType:  make(map[uint16]*Message),
	}

	for _, alias := range s.Aliases {
		a := &Alias{
			Alias: alias,
			Name:  cleanName(alias.Name),
		}
		if !alias.Type.Kind.IsPrimitive() {
			return nil, fmt.Errorf("alias kind is not primitive: %d", alias.Type.Kind)
		}
		if generator.all[a.Name] != nil {
			return nil, fmt.Errorf("alias name conflict: %s", a.Name)
		}
		generator.all[a.Name] = a
		if generator.aliasesByName[a.Name] != nil {
			return nil, fmt.Errorf("duplicate alias name: %s", a.Name)
		}
		generator.aliases = append(generator.aliases, a)
		generator.aliasesByName[a.Name] = a
	}

	for _, constant := range s.Constants {
		c := &Constant{
			Const: constant,
			Name:  constant.Name,
		}
		if !constant.Type.Kind.IsPrimitive() {
			return nil, fmt.Errorf("constant kind is not primitive: %d", constant.Type.Kind)
		}

		if generator.all[c.Name] != nil {
			return nil, fmt.Errorf("constant name conflict: %s", c.Name)
		}
		generator.all[c.Name] = c
		if generator.constantsByName[c.Name] != nil {
			return nil, fmt.Errorf("duplicate constant name: %s", c.Name)
		}
		generator.constants = append(generator.constants, c)
		generator.constantsByName[c.Name] = c
	}
	for _, enum := range s.Enums {
		e := &Enum{
			Enum:          enum,
			Name:          cleanName(enum.Name),
			OptionsByName: make(map[string]*EnumOption),
		}
		if !e.Type.IsPrimitive() {
			return nil, fmt.Errorf("enum kind is not primitive: %d", enum.Type)
		}

		for _, option := range enum.Options {
			o := &EnumOption{
				EnumOption: option,
				Enum:       e,
				Name:       option.Name,
			}
			if generator.all[o.Name] != nil {
				return nil, fmt.Errorf("EnumOption name conflict: %s", o.Name)
			}
			generator.all[o.Name] = o
			e.Options = append(e.Options, o)
			e.OptionsByName[o.Name] = o
		}

		if generator.all[e.Name] != nil {
			return nil, fmt.Errorf("enum name conflict: %s", e.Name)
		}
		generator.all[e.Name] = e
		if generator.enumsByName[e.Name] != nil {
			return nil, fmt.Errorf("duplicate enum name: %s", e.Name)
		}

		generator.enums = append(generator.enums, e)
		generator.enumsByName[e.Name] = e
	}

	createStruct := func(st *schema.Struct) (*Struct, error) {
		if st == nil {
			return nil, nil
		}

		msg := &Struct{
			Struct:       st,
			Name:         cleanName(st.Name),
			FieldsByName: make(map[string]*Field),
		}
		if !st.VLS && s.MessagesByName[st.Name].VLS != nil {
			msg.Name = fmt.Sprintf("%sFixed", msg.Name)
		}
		for _, field := range st.Fields {
			var (
				f = &Field{
					Struct: msg,
					Field:  field,
					Name:   cleanName(field.Name),
				}
			)
			if field.Type.Union != nil {
				for _, field := range field.Type.Union.Fields {
					ff := &Field{
						Struct: msg,
						Field:  field,
						Name:   cleanName(field.Name),
					}
					if msg.FieldsByName[ff.Name] != nil {
						return nil, fmt.Errorf("%s.%s duplicate field name after cleanName", st.Name, field.Name)
					}
					f.Fields = append(f.Fields, ff)
					msg.FieldsByName[ff.Name] = ff
				}
			} else {
				if msg.FieldsByName[f.Name] != nil {
					return nil, fmt.Errorf("%s.%s duplicate field name after cleanName", st.Name, field.Name)
				}
				msg.FieldsByName[f.Name] = f
			}
			msg.Fields = append(msg.Fields, f)
		}
		return msg, nil
	}
	for _, message := range s.Messages {
		var (
			m = &Message{
				NonStandard: message.NonStandard,
			}
			err error
		)
		if m.Fixed, err = createStruct(message.Fixed); err != nil {
			return nil, err
		}
		if m.VLS, err = createStruct(message.VLS); err != nil {
			return nil, err
		}
		name := ""
		if m.Fixed != nil {
			name = m.Fixed.Name
			m.Fixed.Message = m
		}
		if m.VLS != nil {
			name = m.VLS.Name
			m.VLS.Message = m
		}

		existing := generator.messagesByType[message.Type()]
		if existing != nil {
			existing.Extension = m
			m.Extends = existing
		} else {
			generator.messagesByType[message.Type()] = m
		}

		if len(name) == 0 {
			return nil, fmt.Errorf("nil message")
		}
		if generator.messagesByName[name] != nil {
			return nil, fmt.Errorf("duplicate message name after cleanName: %s", name)
		}
		generator.messages = append(generator.messages, m)
		generator.messagesByName[name] = m
	}

	return generator, nil
}

func (g *Generator) Run() error {
	_ = os.MkdirAll(g.config.Dir, 0755)
	files, err := ioutil.ReadDir(g.config.Dir)
	if err != nil {
		return err
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		n := file.Name()
		if err = os.Remove(filepath.Join(g.config.Dir, n)); err != nil {
			return err
		}
	}
	if err := g.generateAliases(); err != nil {
		return err
	}
	if err := g.generateConstants(); err != nil {
		return err
	}
	if err := g.generateEnums(); err != nil {
		return err
	}
	if err := g.generateStructs(); err != nil {
		return err
	}
	if err := g.generateHandler(); err != nil {
		return err
	}
	return nil
}

func (g *Generator) writeFile(name string, data []byte) error {
	path := filepath.Join(g.config.Dir, name)
	if err := os.WriteFile(path, data, 0755); err != nil {
		return err
	}
	if g.config.GoFmt {
		if err := exec.Command("gofmt", "-w", path).Run(); err != nil {
			return err
		}
	}
	return nil
}

func toSnakeCase(name string) string {
	out := make([]byte, 0, len(name)*2)

	// NewOCOOrder
	// TradeOrder
	for i := 0; i < len(name); i++ {
		var (
			isUpper     = string(name[i]) != strings.ToLower(string(name[i]))
			nextIsLower = false
			prevIsLower = false
		)
		if i > 0 {
			prevIsLower = string(name[i-1]) == strings.ToLower(string(name[i-1]))
		}
		if i < len(name)-1 {
			nextIsLower = strings.ToLower(string(name[i+1])) == string(name[i+1])
		}
		if name[i] == '_' {
			continue
		}
		if isUpper {
			if prevIsLower || nextIsLower {
				if len(out) > 0 {
					out = append(out, '_')
				}
			}
		}

		out = append(out, strings.ToLower(string(name[i]))...)
	}

	return string(out)
}

func valueToGo(v *schema.Value) string {
	switch v.Type {
	case schema.ValueTypeBool:
		if v.Int == 0 {
			return "false"
		} else {
			return "true"
		}
	case schema.ValueTypeInt:
		return strconv.FormatInt(v.Int, 10)
	case schema.ValueTypeUint:
		return strconv.FormatUint(v.Uint, 10)
	case schema.ValueTypeFloat:
		return strconv.FormatFloat(v.Float, 'g', 10, 64)
	case schema.ValueTypeFloat32Max:
		return "math.MaxFloat32"
	case schema.ValueTypeFloat64Max:
		return "math.MaxFloat64"
	case schema.ValueTypeString:
		return fmt.Sprintf("\"%s\"", v.Str)
	case schema.ValueTypeConst:
		return cleanName(v.Const.Name)
	case schema.ValueTypeEnumOption:
		return cleanName(v.EnumOption.Name)
	case schema.ValueTypeSizeof:
		return fmt.Sprintf("%d", v.Struct.Size)
	}
	return ""
}

func primitiveTypeName(kind schema.Kind) string {
	switch kind {
	case schema.KindBool:
		return "bool"
	case schema.KindInt8:
		return "int8"
	case schema.KindInt16:
		return "int16"
	case schema.KindInt32:
		return "int32"
	case schema.KindInt64:
		return "int64"
	case schema.KindUint8:
		return "uint8"
	case schema.KindUint16:
		return "uint16"
	case schema.KindUint32:
		return "uint32"
	case schema.KindUint64:
		return "uint64"
	case schema.KindFloat32:
		return "float32"
	case schema.KindFloat64:
		return "float64"
	case schema.KindStringFixed, schema.KindStringVLS:
		return "string"
	default:
		return "invalid"
	}
}

func (g *Generator) GoTypeName(t *schema.Type) string {
	switch t.Kind {
	case schema.KindBool:
		return "bool"
	case schema.KindInt8:
		return "int8"
	case schema.KindInt16:
		return "int16"
	case schema.KindInt32:
		return "int32"
	case schema.KindInt64:
		return "int64"
	case schema.KindUint8:
		return "uint8"
	case schema.KindUint16:
		return "uint16"
	case schema.KindUint32:
		return "uint32"
	case schema.KindUint64:
		return "uint64"
	case schema.KindFloat32:
		return "float32"
	case schema.KindFloat64:
		return "float64"
	case schema.KindStringFixed, schema.KindStringVLS:
		return "string"
	case schema.KindEnum:
		return cleanName(t.Enum.Name)
	case schema.KindAlias:
		return cleanName(t.Alias.Name)
	default:
		return "invalid"
	}
}
