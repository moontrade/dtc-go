//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m UserMessage) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":700,\"F\":["...)
	w.String(m.UserMessage())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsPopupMessage())
	return w.Finish(), nil
}

func (m UserMessageBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":700,\"F\":["...)
	w.String(m.UserMessage())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsPopupMessage())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m UserMessagePointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":700,\"F\":["...)
	w.String(m.UserMessage())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsPopupMessage())
	return w.Finish(), nil
}

func (m UserMessagePointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":700,\"F\":["...)
	w.String(m.UserMessage())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsPopupMessage())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m UserMessage) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 700)
	w.StringField("UserMessage", m.UserMessage())
	w.BoolField("IsPopupMessage", m.IsPopupMessage())
	return w.Finish(), nil
}

func (m UserMessageBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 700)
	w.StringField("UserMessage", m.UserMessage())
	w.BoolField("IsPopupMessage", m.IsPopupMessage())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m UserMessagePointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 700)
	w.StringField("UserMessage", m.UserMessage())
	w.BoolField("IsPopupMessage", m.IsPopupMessage())
	return w.Finish(), nil
}

func (m UserMessagePointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 700)
	w.StringField("UserMessage", m.UserMessage())
	w.BoolField("IsPopupMessage", m.IsPopupMessage())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *UserMessageBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 700 {
		return message.ErrWrongType
	}
	m.SetUserMessage(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsPopupMessage(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *UserMessagePointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 700 {
		return message.ErrWrongType
	}
	m.SetUserMessage(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsPopupMessage(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *UserMessageBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 700 {
		return message.ErrWrongType
	}
	in := &r.Lexer
LOOP:
	for !in.IsDelim('}') {
		key, err := r.FieldName()
		if err != nil {
			return err
		}
		switch key {
		case "UserMessage":
			m.SetUserMessage(r.String())
		case "IsPopupMessage":
			m.SetIsPopupMessage(r.Bool())
		case "f", "F":
			return message.ErrJSONCompactDetected
		case "":
			break LOOP
		default:
			in.SkipRecursive()
		}
		if r.IsError() {
			return r.Error()
		}
		in.WantComma()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *UserMessagePointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 700 {
		return message.ErrWrongType
	}
	in := &r.Lexer
LOOP:
	for !in.IsDelim('}') {
		key, err := r.FieldName()
		if err != nil {
			return err
		}
		switch key {
		case "UserMessage":
			m.SetUserMessage(r.String())
		case "IsPopupMessage":
			m.SetIsPopupMessage(r.Bool())
		case "f", "F":
			return message.ErrJSONCompactDetected
		case "":
			break LOOP
		default:
			in.SkipRecursive()
		}
		if r.IsError() {
			return r.Error()
		}
		in.WantComma()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m UserMessageFixed) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":700,\"F\":["...)
	w.String(m.UserMessage())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsPopupMessage())
	return w.Finish(), nil
}

func (m UserMessageFixedBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":700,\"F\":["...)
	w.String(m.UserMessage())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsPopupMessage())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m UserMessageFixedPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":700,\"F\":["...)
	w.String(m.UserMessage())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsPopupMessage())
	return w.Finish(), nil
}

func (m UserMessageFixedPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":700,\"F\":["...)
	w.String(m.UserMessage())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsPopupMessage())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m UserMessageFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 700)
	w.StringField("UserMessage", m.UserMessage())
	w.BoolField("IsPopupMessage", m.IsPopupMessage())
	return w.Finish(), nil
}

func (m UserMessageFixedBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 700)
	w.StringField("UserMessage", m.UserMessage())
	w.BoolField("IsPopupMessage", m.IsPopupMessage())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m UserMessageFixedPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 700)
	w.StringField("UserMessage", m.UserMessage())
	w.BoolField("IsPopupMessage", m.IsPopupMessage())
	return w.Finish(), nil
}

func (m UserMessageFixedPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 700)
	w.StringField("UserMessage", m.UserMessage())
	w.BoolField("IsPopupMessage", m.IsPopupMessage())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *UserMessageFixedBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 700 {
		return message.ErrWrongType
	}
	m.SetUserMessage(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsPopupMessage(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *UserMessageFixedPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 700 {
		return message.ErrWrongType
	}
	m.SetUserMessage(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsPopupMessage(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *UserMessageFixedBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 700 {
		return message.ErrWrongType
	}
	in := &r.Lexer
LOOP:
	for !in.IsDelim('}') {
		key, err := r.FieldName()
		if err != nil {
			return err
		}
		switch key {
		case "UserMessage":
			m.SetUserMessage(r.String())
		case "IsPopupMessage":
			m.SetIsPopupMessage(r.Bool())
		case "f", "F":
			return message.ErrJSONCompactDetected
		case "":
			break LOOP
		default:
			in.SkipRecursive()
		}
		if r.IsError() {
			return r.Error()
		}
		in.WantComma()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *UserMessageFixedPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 700 {
		return message.ErrWrongType
	}
	in := &r.Lexer
LOOP:
	for !in.IsDelim('}') {
		key, err := r.FieldName()
		if err != nil {
			return err
		}
		switch key {
		case "UserMessage":
			m.SetUserMessage(r.String())
		case "IsPopupMessage":
			m.SetIsPopupMessage(r.Bool())
		case "f", "F":
			return message.ErrJSONCompactDetected
		case "":
			break LOOP
		default:
			in.SkipRecursive()
		}
		if r.IsError() {
			return r.Error()
		}
		in.WantComma()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *UserMessageBuilder) UnmarshalProtobuf(b []byte) error {
	var (
		r   = pb.NewBuffer(b)
		f   pb.Number
		err error
	)
	for !r.EOF() {
		f, _, err = r.DecodeTag()
		if err != nil {
			break
		}
		switch f {
		case 1:
			m.SetUserMessage(r.ReadString())
		case 2:
			m.SetIsPopupMessage(r.ReadBool())
		default:
			r.Consume()
		}
		if err = r.Error(); err != nil {
			return err
		}
	}
	return err
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *UserMessagePointerBuilder) UnmarshalProtobuf(b []byte) error {
	var (
		r   = pb.NewBuffer(b)
		f   pb.Number
		err error
	)
	for !r.EOF() {
		f, _, err = r.DecodeTag()
		if err != nil {
			break
		}
		switch f {
		case 1:
			m.SetUserMessage(r.ReadString())
		case 2:
			m.SetIsPopupMessage(r.ReadBool())
		default:
			r.Consume()
		}
		if err = r.Error(); err != nil {
			return err
		}
	}
	return err
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m UserMessageBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 700)
	w.WriteString(1, m.UserMessage())
	w.WriteBool(2, m.IsPopupMessage())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m UserMessagePointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 700)
	w.WriteString(1, m.UserMessage())
	w.WriteBool(2, m.IsPopupMessage())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *UserMessageFixedBuilder) UnmarshalProtobuf(b []byte) error {
	var (
		r   = pb.NewBuffer(b)
		f   pb.Number
		err error
	)
	for !r.EOF() {
		f, _, err = r.DecodeTag()
		if err != nil {
			break
		}
		switch f {
		case 1:
			m.SetUserMessage(r.ReadString())
		case 2:
			m.SetIsPopupMessage(r.ReadBool())
		default:
			r.Consume()
		}
		if err = r.Error(); err != nil {
			return err
		}
	}
	return err
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *UserMessageFixedPointerBuilder) UnmarshalProtobuf(b []byte) error {
	var (
		r   = pb.NewBuffer(b)
		f   pb.Number
		err error
	)
	for !r.EOF() {
		f, _, err = r.DecodeTag()
		if err != nil {
			break
		}
		switch f {
		case 1:
			m.SetUserMessage(r.ReadString())
		case 2:
			m.SetIsPopupMessage(r.ReadBool())
		default:
			r.Consume()
		}
		if err = r.Error(); err != nil {
			return err
		}
	}
	return err
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m UserMessageFixedBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 700)
	w.WriteString(1, m.UserMessage())
	w.WriteBool(2, m.IsPopupMessage())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m UserMessageFixedPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 700)
	w.WriteString(1, m.UserMessage())
	w.WriteBool(2, m.IsPopupMessage())
	return w.Finish(), nil
}