//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m GeneralLogMessageFixedPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":701,\"F\":["...)
	w.String(m.MessageText())
	return w.Finish(), nil
}

func (m GeneralLogMessageFixedPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":701,\"F\":["...)
	w.String(m.MessageText())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m GeneralLogMessageFixedPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 701)
	w.StringField("MessageText", m.MessageText())
	return w.Finish(), nil
}

func (m GeneralLogMessageFixedPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 701)
	w.StringField("MessageText", m.MessageText())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *GeneralLogMessageFixedPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 701 {
		return message.ErrWrongType
	}
	m.SetMessageText(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *GeneralLogMessageFixedPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 701 {
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
		case "MessageText":
			m.SetMessageText(r.String())
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
