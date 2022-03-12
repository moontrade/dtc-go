//go:build !tinygo

package sc

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalTradesRequest) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10100,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.StartDateTime()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.SubAccountIdentifier())
	w.Data = append(w.Data, ',')
	w.Uint8(m.CreateFlatToFlatTrades())
	return w.Finish(), nil
}

func (m HistoricalTradesRequestBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10100,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.StartDateTime()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.SubAccountIdentifier())
	w.Data = append(w.Data, ',')
	w.Uint8(m.CreateFlatToFlatTrades())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalTradesRequestPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10100,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.StartDateTime()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.SubAccountIdentifier())
	w.Data = append(w.Data, ',')
	w.Uint8(m.CreateFlatToFlatTrades())
	return w.Finish(), nil
}

func (m HistoricalTradesRequestPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10100,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.StartDateTime()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.SubAccountIdentifier())
	w.Data = append(w.Data, ',')
	w.Uint8(m.CreateFlatToFlatTrades())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalTradesRequest) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10100)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Int64Field("StartDateTime", int64(m.StartDateTime()))
	w.Uint32Field("SubAccountIdentifier", m.SubAccountIdentifier())
	w.Uint8Field("CreateFlatToFlatTrades", m.CreateFlatToFlatTrades())
	return w.Finish(), nil
}

func (m HistoricalTradesRequestBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10100)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Int64Field("StartDateTime", int64(m.StartDateTime()))
	w.Uint32Field("SubAccountIdentifier", m.SubAccountIdentifier())
	w.Uint8Field("CreateFlatToFlatTrades", m.CreateFlatToFlatTrades())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalTradesRequestPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10100)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Int64Field("StartDateTime", int64(m.StartDateTime()))
	w.Uint32Field("SubAccountIdentifier", m.SubAccountIdentifier())
	w.Uint8Field("CreateFlatToFlatTrades", m.CreateFlatToFlatTrades())
	return w.Finish(), nil
}

func (m HistoricalTradesRequestPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10100)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Int64Field("StartDateTime", int64(m.StartDateTime()))
	w.Uint32Field("SubAccountIdentifier", m.SubAccountIdentifier())
	w.Uint8Field("CreateFlatToFlatTrades", m.CreateFlatToFlatTrades())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalTradesRequestBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10100 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSymbol(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetStartDateTime(DateTime(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetSubAccountIdentifier(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetCreateFlatToFlatTrades(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalTradesRequestPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10100 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSymbol(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetStartDateTime(DateTime(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetSubAccountIdentifier(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetCreateFlatToFlatTrades(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalTradesRequestBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 10100 {
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
		case "Symbol":
			m.SetSymbol(r.String())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "StartDateTime":
			m.SetStartDateTime(DateTime(r.Int64()))
		case "SubAccountIdentifier":
			m.SetSubAccountIdentifier(r.Uint32())
		case "CreateFlatToFlatTrades":
			m.SetCreateFlatToFlatTrades(r.Uint8())
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

func (m *HistoricalTradesRequestPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 10100 {
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
		case "Symbol":
			m.SetSymbol(r.String())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "StartDateTime":
			m.SetStartDateTime(DateTime(r.Int64()))
		case "SubAccountIdentifier":
			m.SetSubAccountIdentifier(r.Uint32())
		case "CreateFlatToFlatTrades":
			m.SetCreateFlatToFlatTrades(r.Uint8())
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

func (m HistoricalTradesRequestFixed) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10100,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.StartDateTime()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.SubAccountIdentifier())
	w.Data = append(w.Data, ',')
	w.Uint8(m.CreateFlatToFlatTrades())
	return w.Finish(), nil
}

func (m HistoricalTradesRequestFixedBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10100,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.StartDateTime()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.SubAccountIdentifier())
	w.Data = append(w.Data, ',')
	w.Uint8(m.CreateFlatToFlatTrades())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalTradesRequestFixedPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10100,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.StartDateTime()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.SubAccountIdentifier())
	w.Data = append(w.Data, ',')
	w.Uint8(m.CreateFlatToFlatTrades())
	return w.Finish(), nil
}

func (m HistoricalTradesRequestFixedPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10100,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.StartDateTime()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.SubAccountIdentifier())
	w.Data = append(w.Data, ',')
	w.Uint8(m.CreateFlatToFlatTrades())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalTradesRequestFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10100)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Int64Field("StartDateTime", int64(m.StartDateTime()))
	w.Uint32Field("SubAccountIdentifier", m.SubAccountIdentifier())
	w.Uint8Field("CreateFlatToFlatTrades", m.CreateFlatToFlatTrades())
	return w.Finish(), nil
}

func (m HistoricalTradesRequestFixedBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10100)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Int64Field("StartDateTime", int64(m.StartDateTime()))
	w.Uint32Field("SubAccountIdentifier", m.SubAccountIdentifier())
	w.Uint8Field("CreateFlatToFlatTrades", m.CreateFlatToFlatTrades())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalTradesRequestFixedPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10100)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Int64Field("StartDateTime", int64(m.StartDateTime()))
	w.Uint32Field("SubAccountIdentifier", m.SubAccountIdentifier())
	w.Uint8Field("CreateFlatToFlatTrades", m.CreateFlatToFlatTrades())
	return w.Finish(), nil
}

func (m HistoricalTradesRequestFixedPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10100)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Int64Field("StartDateTime", int64(m.StartDateTime()))
	w.Uint32Field("SubAccountIdentifier", m.SubAccountIdentifier())
	w.Uint8Field("CreateFlatToFlatTrades", m.CreateFlatToFlatTrades())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalTradesRequestFixedBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10100 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSymbol(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetStartDateTime(DateTime(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetSubAccountIdentifier(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetCreateFlatToFlatTrades(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalTradesRequestFixedPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10100 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSymbol(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetStartDateTime(DateTime(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetSubAccountIdentifier(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetCreateFlatToFlatTrades(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalTradesRequestFixedBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 10100 {
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
		case "Symbol":
			m.SetSymbol(r.String())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "StartDateTime":
			m.SetStartDateTime(DateTime(r.Int64()))
		case "SubAccountIdentifier":
			m.SetSubAccountIdentifier(r.Uint32())
		case "CreateFlatToFlatTrades":
			m.SetCreateFlatToFlatTrades(r.Uint8())
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

func (m *HistoricalTradesRequestFixedPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 10100 {
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
		case "Symbol":
			m.SetSymbol(r.String())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "StartDateTime":
			m.SetStartDateTime(DateTime(r.Int64()))
		case "SubAccountIdentifier":
			m.SetSubAccountIdentifier(r.Uint32())
		case "CreateFlatToFlatTrades":
			m.SetCreateFlatToFlatTrades(r.Uint8())
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