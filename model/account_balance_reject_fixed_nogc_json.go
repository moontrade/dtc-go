//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m AccountBalanceRejectFixedPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":602,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.RejectText())
	return w.Finish(), nil
}

func (m AccountBalanceRejectFixedPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":602,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.RejectText())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m AccountBalanceRejectFixedPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 602)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("RejectText", m.RejectText())
	return w.Finish(), nil
}

func (m AccountBalanceRejectFixedPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 602)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("RejectText", m.RejectText())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *AccountBalanceRejectFixedPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 602 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetRejectText(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *AccountBalanceRejectFixedPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 602 {
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
		case "RequestID":
			m.SetRequestID(r.Int32())
		case "RejectText":
			m.SetRejectText(r.String())
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
