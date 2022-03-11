//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m WriteIntradayDataFileSessionValue) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10140,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.ValueType()))
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.Date()))
	w.Data = append(w.Data, ',')
	w.Float64(m.Price())
	w.Data = append(w.Data, ',')
	w.Float64(m.Volume())
	return w.Finish(), nil
}

func (m WriteIntradayDataFileSessionValueBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10140,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.ValueType()))
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.Date()))
	w.Data = append(w.Data, ',')
	w.Float64(m.Price())
	w.Data = append(w.Data, ',')
	w.Float64(m.Volume())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m WriteIntradayDataFileSessionValue) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10140)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Int32Field("ValueType", int32(m.ValueType()))
	w.Int64Field("DateTime", int64(m.DateTime()))
	w.Int64Field("Date", int64(m.Date()))
	w.Float64Field("Price", m.Price())
	w.Float64Field("Volume", m.Volume())
	return w.Finish(), nil
}

func (m WriteIntradayDataFileSessionValueBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10140)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Int32Field("ValueType", int32(m.ValueType()))
	w.Int64Field("DateTime", int64(m.DateTime()))
	w.Int64Field("Date", int64(m.Date()))
	w.Float64Field("Price", m.Price())
	w.Float64Field("Volume", m.Volume())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *WriteIntradayDataFileSessionValueBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10140 {
		return message.ErrWrongType
	}
	m.SetSymbolID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetValueType(WriteIntradayDataFileSessionValueTypeEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetDateTime(DateTimeWithMicrosecondsInt(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetDate(DateTime(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetVolume(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *WriteIntradayDataFileSessionValueBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 10140 {
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
		case "SymbolID":
			m.SetSymbolID(r.Uint32())
		case "ValueType":
			m.SetValueType(WriteIntradayDataFileSessionValueTypeEnum(r.Int32()))
		case "DateTime":
			m.SetDateTime(DateTimeWithMicrosecondsInt(r.Int64()))
		case "Date":
			m.SetDate(DateTime(r.Int64()))
		case "Price":
			m.SetPrice(r.Float64())
		case "Volume":
			m.SetVolume(r.Float64())
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
