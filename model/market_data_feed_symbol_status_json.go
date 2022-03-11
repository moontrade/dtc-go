//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataFeedSymbolStatus) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":116,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.Status()))
	return w.Finish(), nil
}

func (m MarketDataFeedSymbolStatusBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":116,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.Status()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataFeedSymbolStatus) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 116)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Int32Field("Status", int32(m.Status()))
	return w.Finish(), nil
}

func (m MarketDataFeedSymbolStatusBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 116)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Int32Field("Status", int32(m.Status()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataFeedSymbolStatusBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 116 {
		return message.ErrWrongType
	}
	m.SetSymbolID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetStatus(MarketDataFeedStatusEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataFeedSymbolStatusBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 116 {
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
		case "Status":
			m.SetStatus(MarketDataFeedStatusEnum(r.Int32()))
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
