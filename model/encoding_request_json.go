//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m EncodingRequest) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":6,\"F\":["...)
	w.Int32(m.ProtocolVersion())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.Encoding()))
	w.Data = append(w.Data, ',')
	w.String(m.ProtocolType())
	return w.Finish(), nil
}

func (m EncodingRequestBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":6,\"F\":["...)
	w.Int32(m.ProtocolVersion())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.Encoding()))
	w.Data = append(w.Data, ',')
	w.String(m.ProtocolType())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m EncodingRequest) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 6)
	w.Int32Field("ProtocolVersion", m.ProtocolVersion())
	w.Int32Field("Encoding", int32(m.Encoding()))
	w.StringField("ProtocolType", m.ProtocolType())
	return w.Finish(), nil
}

func (m EncodingRequestBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 6)
	w.Int32Field("ProtocolVersion", m.ProtocolVersion())
	w.Int32Field("Encoding", int32(m.Encoding()))
	w.StringField("ProtocolType", m.ProtocolType())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *EncodingRequestBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 6 {
		return message.ErrWrongType
	}
	m.SetProtocolVersion(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetEncoding(EncodingEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetProtocolType(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *EncodingRequestBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 6 {
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
		case "ProtocolVersion":
			m.SetProtocolVersion(r.Int32())
		case "Encoding":
			m.SetEncoding(EncodingEnum(r.Int32()))
		case "ProtocolType":
			m.SetProtocolType(r.String())
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
