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

func (m HistoricalOrderFillResponse) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":304,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Int32(m.TotalNumberMessages())
	w.Data = append(w.Data, ',')
	w.Int32(m.MessageNumber())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.String(m.ServerOrderID())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.BuySell()))
	w.Data = append(w.Data, ',')
	w.Float64(m.Price())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity())
	w.Data = append(w.Data, ',')
	w.String(m.UniqueExecutionID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.OpenClose()))
	w.Data = append(w.Data, ',')
	w.Uint8(m.NoOrderFills())
	w.Data = append(w.Data, ',')
	w.String(m.InfoText())
	w.Data = append(w.Data, ',')
	w.Float64(m.HighPriceDuringPosition())
	w.Data = append(w.Data, ',')
	w.Float64(m.LowPriceDuringPosition())
	w.Data = append(w.Data, ',')
	w.Float64(m.PositionQuantity())
	return w.Finish(), nil
}

func (m HistoricalOrderFillResponseBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":304,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Int32(m.TotalNumberMessages())
	w.Data = append(w.Data, ',')
	w.Int32(m.MessageNumber())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.String(m.ServerOrderID())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.BuySell()))
	w.Data = append(w.Data, ',')
	w.Float64(m.Price())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity())
	w.Data = append(w.Data, ',')
	w.String(m.UniqueExecutionID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.OpenClose()))
	w.Data = append(w.Data, ',')
	w.Uint8(m.NoOrderFills())
	w.Data = append(w.Data, ',')
	w.String(m.InfoText())
	w.Data = append(w.Data, ',')
	w.Float64(m.HighPriceDuringPosition())
	w.Data = append(w.Data, ',')
	w.Float64(m.LowPriceDuringPosition())
	w.Data = append(w.Data, ',')
	w.Float64(m.PositionQuantity())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalOrderFillResponsePointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":304,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Int32(m.TotalNumberMessages())
	w.Data = append(w.Data, ',')
	w.Int32(m.MessageNumber())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.String(m.ServerOrderID())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.BuySell()))
	w.Data = append(w.Data, ',')
	w.Float64(m.Price())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity())
	w.Data = append(w.Data, ',')
	w.String(m.UniqueExecutionID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.OpenClose()))
	w.Data = append(w.Data, ',')
	w.Uint8(m.NoOrderFills())
	w.Data = append(w.Data, ',')
	w.String(m.InfoText())
	w.Data = append(w.Data, ',')
	w.Float64(m.HighPriceDuringPosition())
	w.Data = append(w.Data, ',')
	w.Float64(m.LowPriceDuringPosition())
	w.Data = append(w.Data, ',')
	w.Float64(m.PositionQuantity())
	return w.Finish(), nil
}

func (m HistoricalOrderFillResponsePointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":304,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Int32(m.TotalNumberMessages())
	w.Data = append(w.Data, ',')
	w.Int32(m.MessageNumber())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.String(m.ServerOrderID())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.BuySell()))
	w.Data = append(w.Data, ',')
	w.Float64(m.Price())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity())
	w.Data = append(w.Data, ',')
	w.String(m.UniqueExecutionID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.OpenClose()))
	w.Data = append(w.Data, ',')
	w.Uint8(m.NoOrderFills())
	w.Data = append(w.Data, ',')
	w.String(m.InfoText())
	w.Data = append(w.Data, ',')
	w.Float64(m.HighPriceDuringPosition())
	w.Data = append(w.Data, ',')
	w.Float64(m.LowPriceDuringPosition())
	w.Data = append(w.Data, ',')
	w.Float64(m.PositionQuantity())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalOrderFillResponse) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 304)
	w.Int32Field("RequestID", m.RequestID())
	w.Int32Field("TotalNumberMessages", m.TotalNumberMessages())
	w.Int32Field("MessageNumber", m.MessageNumber())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.Int32Field("BuySell", int32(m.BuySell()))
	w.Float64Field("Price", m.Price())
	w.Int64Field("DateTime", int64(m.DateTime()))
	w.Float64Field("Quantity", m.Quantity())
	w.StringField("UniqueExecutionID", m.UniqueExecutionID())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Int32Field("OpenClose", int32(m.OpenClose()))
	w.Uint8Field("NoOrderFills", m.NoOrderFills())
	w.StringField("InfoText", m.InfoText())
	w.Float64Field("HighPriceDuringPosition", m.HighPriceDuringPosition())
	w.Float64Field("LowPriceDuringPosition", m.LowPriceDuringPosition())
	w.Float64Field("PositionQuantity", m.PositionQuantity())
	return w.Finish(), nil
}

func (m HistoricalOrderFillResponseBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 304)
	w.Int32Field("RequestID", m.RequestID())
	w.Int32Field("TotalNumberMessages", m.TotalNumberMessages())
	w.Int32Field("MessageNumber", m.MessageNumber())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.Int32Field("BuySell", int32(m.BuySell()))
	w.Float64Field("Price", m.Price())
	w.Int64Field("DateTime", int64(m.DateTime()))
	w.Float64Field("Quantity", m.Quantity())
	w.StringField("UniqueExecutionID", m.UniqueExecutionID())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Int32Field("OpenClose", int32(m.OpenClose()))
	w.Uint8Field("NoOrderFills", m.NoOrderFills())
	w.StringField("InfoText", m.InfoText())
	w.Float64Field("HighPriceDuringPosition", m.HighPriceDuringPosition())
	w.Float64Field("LowPriceDuringPosition", m.LowPriceDuringPosition())
	w.Float64Field("PositionQuantity", m.PositionQuantity())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalOrderFillResponsePointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 304)
	w.Int32Field("RequestID", m.RequestID())
	w.Int32Field("TotalNumberMessages", m.TotalNumberMessages())
	w.Int32Field("MessageNumber", m.MessageNumber())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.Int32Field("BuySell", int32(m.BuySell()))
	w.Float64Field("Price", m.Price())
	w.Int64Field("DateTime", int64(m.DateTime()))
	w.Float64Field("Quantity", m.Quantity())
	w.StringField("UniqueExecutionID", m.UniqueExecutionID())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Int32Field("OpenClose", int32(m.OpenClose()))
	w.Uint8Field("NoOrderFills", m.NoOrderFills())
	w.StringField("InfoText", m.InfoText())
	w.Float64Field("HighPriceDuringPosition", m.HighPriceDuringPosition())
	w.Float64Field("LowPriceDuringPosition", m.LowPriceDuringPosition())
	w.Float64Field("PositionQuantity", m.PositionQuantity())
	return w.Finish(), nil
}

func (m HistoricalOrderFillResponsePointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 304)
	w.Int32Field("RequestID", m.RequestID())
	w.Int32Field("TotalNumberMessages", m.TotalNumberMessages())
	w.Int32Field("MessageNumber", m.MessageNumber())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.Int32Field("BuySell", int32(m.BuySell()))
	w.Float64Field("Price", m.Price())
	w.Int64Field("DateTime", int64(m.DateTime()))
	w.Float64Field("Quantity", m.Quantity())
	w.StringField("UniqueExecutionID", m.UniqueExecutionID())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Int32Field("OpenClose", int32(m.OpenClose()))
	w.Uint8Field("NoOrderFills", m.NoOrderFills())
	w.StringField("InfoText", m.InfoText())
	w.Float64Field("HighPriceDuringPosition", m.HighPriceDuringPosition())
	w.Float64Field("LowPriceDuringPosition", m.LowPriceDuringPosition())
	w.Float64Field("PositionQuantity", m.PositionQuantity())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalOrderFillResponseBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 304 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetTotalNumberMessages(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetMessageNumber(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSymbol(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetExchange(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetServerOrderID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetBuySell(BuySellEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetDateTime(DateTime(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetUniqueExecutionID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetOpenClose(OpenCloseTradeEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetNoOrderFills(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetInfoText(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetHighPriceDuringPosition(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetLowPriceDuringPosition(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetPositionQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalOrderFillResponsePointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 304 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetTotalNumberMessages(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetMessageNumber(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSymbol(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetExchange(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetServerOrderID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetBuySell(BuySellEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetDateTime(DateTime(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetUniqueExecutionID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetOpenClose(OpenCloseTradeEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetNoOrderFills(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetInfoText(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetHighPriceDuringPosition(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetLowPriceDuringPosition(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetPositionQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalOrderFillResponseBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 304 {
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
		case "TotalNumberMessages":
			m.SetTotalNumberMessages(r.Int32())
		case "MessageNumber":
			m.SetMessageNumber(r.Int32())
		case "Symbol":
			m.SetSymbol(r.String())
		case "Exchange":
			m.SetExchange(r.String())
		case "ServerOrderID":
			m.SetServerOrderID(r.String())
		case "BuySell":
			m.SetBuySell(BuySellEnum(r.Int32()))
		case "Price":
			m.SetPrice(r.Float64())
		case "DateTime":
			m.SetDateTime(DateTime(r.Int64()))
		case "Quantity":
			m.SetQuantity(r.Float64())
		case "UniqueExecutionID":
			m.SetUniqueExecutionID(r.String())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "OpenClose":
			m.SetOpenClose(OpenCloseTradeEnum(r.Int32()))
		case "NoOrderFills":
			m.SetNoOrderFills(r.Uint8())
		case "InfoText":
			m.SetInfoText(r.String())
		case "HighPriceDuringPosition":
			m.SetHighPriceDuringPosition(r.Float64())
		case "LowPriceDuringPosition":
			m.SetLowPriceDuringPosition(r.Float64())
		case "PositionQuantity":
			m.SetPositionQuantity(r.Float64())
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

func (m *HistoricalOrderFillResponsePointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 304 {
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
		case "TotalNumberMessages":
			m.SetTotalNumberMessages(r.Int32())
		case "MessageNumber":
			m.SetMessageNumber(r.Int32())
		case "Symbol":
			m.SetSymbol(r.String())
		case "Exchange":
			m.SetExchange(r.String())
		case "ServerOrderID":
			m.SetServerOrderID(r.String())
		case "BuySell":
			m.SetBuySell(BuySellEnum(r.Int32()))
		case "Price":
			m.SetPrice(r.Float64())
		case "DateTime":
			m.SetDateTime(DateTime(r.Int64()))
		case "Quantity":
			m.SetQuantity(r.Float64())
		case "UniqueExecutionID":
			m.SetUniqueExecutionID(r.String())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "OpenClose":
			m.SetOpenClose(OpenCloseTradeEnum(r.Int32()))
		case "NoOrderFills":
			m.SetNoOrderFills(r.Uint8())
		case "InfoText":
			m.SetInfoText(r.String())
		case "HighPriceDuringPosition":
			m.SetHighPriceDuringPosition(r.Float64())
		case "LowPriceDuringPosition":
			m.SetLowPriceDuringPosition(r.Float64())
		case "PositionQuantity":
			m.SetPositionQuantity(r.Float64())
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

func (m HistoricalOrderFillResponseFixed) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":304,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Int32(m.TotalNumberMessages())
	w.Data = append(w.Data, ',')
	w.Int32(m.MessageNumber())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.String(m.ServerOrderID())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.BuySell()))
	w.Data = append(w.Data, ',')
	w.Float64(m.Price())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity())
	w.Data = append(w.Data, ',')
	w.String(m.UniqueExecutionID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.OpenClose()))
	w.Data = append(w.Data, ',')
	w.Uint8(m.NoOrderFills())
	w.Data = append(w.Data, ',')
	w.String(m.InfoText())
	w.Data = append(w.Data, ',')
	w.Float64(m.HighPriceDuringPosition())
	w.Data = append(w.Data, ',')
	w.Float64(m.LowPriceDuringPosition())
	w.Data = append(w.Data, ',')
	w.Float64(m.PositionQuantity())
	return w.Finish(), nil
}

func (m HistoricalOrderFillResponseFixedBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":304,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Int32(m.TotalNumberMessages())
	w.Data = append(w.Data, ',')
	w.Int32(m.MessageNumber())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.String(m.ServerOrderID())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.BuySell()))
	w.Data = append(w.Data, ',')
	w.Float64(m.Price())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity())
	w.Data = append(w.Data, ',')
	w.String(m.UniqueExecutionID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.OpenClose()))
	w.Data = append(w.Data, ',')
	w.Uint8(m.NoOrderFills())
	w.Data = append(w.Data, ',')
	w.String(m.InfoText())
	w.Data = append(w.Data, ',')
	w.Float64(m.HighPriceDuringPosition())
	w.Data = append(w.Data, ',')
	w.Float64(m.LowPriceDuringPosition())
	w.Data = append(w.Data, ',')
	w.Float64(m.PositionQuantity())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalOrderFillResponseFixedPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":304,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Int32(m.TotalNumberMessages())
	w.Data = append(w.Data, ',')
	w.Int32(m.MessageNumber())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.String(m.ServerOrderID())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.BuySell()))
	w.Data = append(w.Data, ',')
	w.Float64(m.Price())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity())
	w.Data = append(w.Data, ',')
	w.String(m.UniqueExecutionID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.OpenClose()))
	w.Data = append(w.Data, ',')
	w.Uint8(m.NoOrderFills())
	w.Data = append(w.Data, ',')
	w.String(m.InfoText())
	w.Data = append(w.Data, ',')
	w.Float64(m.HighPriceDuringPosition())
	w.Data = append(w.Data, ',')
	w.Float64(m.LowPriceDuringPosition())
	w.Data = append(w.Data, ',')
	w.Float64(m.PositionQuantity())
	return w.Finish(), nil
}

func (m HistoricalOrderFillResponseFixedPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":304,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Int32(m.TotalNumberMessages())
	w.Data = append(w.Data, ',')
	w.Int32(m.MessageNumber())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.String(m.ServerOrderID())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.BuySell()))
	w.Data = append(w.Data, ',')
	w.Float64(m.Price())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity())
	w.Data = append(w.Data, ',')
	w.String(m.UniqueExecutionID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.OpenClose()))
	w.Data = append(w.Data, ',')
	w.Uint8(m.NoOrderFills())
	w.Data = append(w.Data, ',')
	w.String(m.InfoText())
	w.Data = append(w.Data, ',')
	w.Float64(m.HighPriceDuringPosition())
	w.Data = append(w.Data, ',')
	w.Float64(m.LowPriceDuringPosition())
	w.Data = append(w.Data, ',')
	w.Float64(m.PositionQuantity())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalOrderFillResponseFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 304)
	w.Int32Field("RequestID", m.RequestID())
	w.Int32Field("TotalNumberMessages", m.TotalNumberMessages())
	w.Int32Field("MessageNumber", m.MessageNumber())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.Int32Field("BuySell", int32(m.BuySell()))
	w.Float64Field("Price", m.Price())
	w.Int64Field("DateTime", int64(m.DateTime()))
	w.Float64Field("Quantity", m.Quantity())
	w.StringField("UniqueExecutionID", m.UniqueExecutionID())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Int32Field("OpenClose", int32(m.OpenClose()))
	w.Uint8Field("NoOrderFills", m.NoOrderFills())
	w.StringField("InfoText", m.InfoText())
	w.Float64Field("HighPriceDuringPosition", m.HighPriceDuringPosition())
	w.Float64Field("LowPriceDuringPosition", m.LowPriceDuringPosition())
	w.Float64Field("PositionQuantity", m.PositionQuantity())
	return w.Finish(), nil
}

func (m HistoricalOrderFillResponseFixedBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 304)
	w.Int32Field("RequestID", m.RequestID())
	w.Int32Field("TotalNumberMessages", m.TotalNumberMessages())
	w.Int32Field("MessageNumber", m.MessageNumber())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.Int32Field("BuySell", int32(m.BuySell()))
	w.Float64Field("Price", m.Price())
	w.Int64Field("DateTime", int64(m.DateTime()))
	w.Float64Field("Quantity", m.Quantity())
	w.StringField("UniqueExecutionID", m.UniqueExecutionID())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Int32Field("OpenClose", int32(m.OpenClose()))
	w.Uint8Field("NoOrderFills", m.NoOrderFills())
	w.StringField("InfoText", m.InfoText())
	w.Float64Field("HighPriceDuringPosition", m.HighPriceDuringPosition())
	w.Float64Field("LowPriceDuringPosition", m.LowPriceDuringPosition())
	w.Float64Field("PositionQuantity", m.PositionQuantity())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalOrderFillResponseFixedPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 304)
	w.Int32Field("RequestID", m.RequestID())
	w.Int32Field("TotalNumberMessages", m.TotalNumberMessages())
	w.Int32Field("MessageNumber", m.MessageNumber())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.Int32Field("BuySell", int32(m.BuySell()))
	w.Float64Field("Price", m.Price())
	w.Int64Field("DateTime", int64(m.DateTime()))
	w.Float64Field("Quantity", m.Quantity())
	w.StringField("UniqueExecutionID", m.UniqueExecutionID())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Int32Field("OpenClose", int32(m.OpenClose()))
	w.Uint8Field("NoOrderFills", m.NoOrderFills())
	w.StringField("InfoText", m.InfoText())
	w.Float64Field("HighPriceDuringPosition", m.HighPriceDuringPosition())
	w.Float64Field("LowPriceDuringPosition", m.LowPriceDuringPosition())
	w.Float64Field("PositionQuantity", m.PositionQuantity())
	return w.Finish(), nil
}

func (m HistoricalOrderFillResponseFixedPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 304)
	w.Int32Field("RequestID", m.RequestID())
	w.Int32Field("TotalNumberMessages", m.TotalNumberMessages())
	w.Int32Field("MessageNumber", m.MessageNumber())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.Int32Field("BuySell", int32(m.BuySell()))
	w.Float64Field("Price", m.Price())
	w.Int64Field("DateTime", int64(m.DateTime()))
	w.Float64Field("Quantity", m.Quantity())
	w.StringField("UniqueExecutionID", m.UniqueExecutionID())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Int32Field("OpenClose", int32(m.OpenClose()))
	w.Uint8Field("NoOrderFills", m.NoOrderFills())
	w.StringField("InfoText", m.InfoText())
	w.Float64Field("HighPriceDuringPosition", m.HighPriceDuringPosition())
	w.Float64Field("LowPriceDuringPosition", m.LowPriceDuringPosition())
	w.Float64Field("PositionQuantity", m.PositionQuantity())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalOrderFillResponseFixedBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 304 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetTotalNumberMessages(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetMessageNumber(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSymbol(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetExchange(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetServerOrderID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetBuySell(BuySellEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetDateTime(DateTime(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetUniqueExecutionID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetOpenClose(OpenCloseTradeEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetNoOrderFills(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetInfoText(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetHighPriceDuringPosition(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetLowPriceDuringPosition(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetPositionQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalOrderFillResponseFixedPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 304 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetTotalNumberMessages(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetMessageNumber(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSymbol(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetExchange(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetServerOrderID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetBuySell(BuySellEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetDateTime(DateTime(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetUniqueExecutionID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetOpenClose(OpenCloseTradeEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetNoOrderFills(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetInfoText(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetHighPriceDuringPosition(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetLowPriceDuringPosition(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetPositionQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalOrderFillResponseFixedBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 304 {
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
		case "TotalNumberMessages":
			m.SetTotalNumberMessages(r.Int32())
		case "MessageNumber":
			m.SetMessageNumber(r.Int32())
		case "Symbol":
			m.SetSymbol(r.String())
		case "Exchange":
			m.SetExchange(r.String())
		case "ServerOrderID":
			m.SetServerOrderID(r.String())
		case "BuySell":
			m.SetBuySell(BuySellEnum(r.Int32()))
		case "Price":
			m.SetPrice(r.Float64())
		case "DateTime":
			m.SetDateTime(DateTime(r.Int64()))
		case "Quantity":
			m.SetQuantity(r.Float64())
		case "UniqueExecutionID":
			m.SetUniqueExecutionID(r.String())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "OpenClose":
			m.SetOpenClose(OpenCloseTradeEnum(r.Int32()))
		case "NoOrderFills":
			m.SetNoOrderFills(r.Uint8())
		case "InfoText":
			m.SetInfoText(r.String())
		case "HighPriceDuringPosition":
			m.SetHighPriceDuringPosition(r.Float64())
		case "LowPriceDuringPosition":
			m.SetLowPriceDuringPosition(r.Float64())
		case "PositionQuantity":
			m.SetPositionQuantity(r.Float64())
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

func (m *HistoricalOrderFillResponseFixedPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 304 {
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
		case "TotalNumberMessages":
			m.SetTotalNumberMessages(r.Int32())
		case "MessageNumber":
			m.SetMessageNumber(r.Int32())
		case "Symbol":
			m.SetSymbol(r.String())
		case "Exchange":
			m.SetExchange(r.String())
		case "ServerOrderID":
			m.SetServerOrderID(r.String())
		case "BuySell":
			m.SetBuySell(BuySellEnum(r.Int32()))
		case "Price":
			m.SetPrice(r.Float64())
		case "DateTime":
			m.SetDateTime(DateTime(r.Int64()))
		case "Quantity":
			m.SetQuantity(r.Float64())
		case "UniqueExecutionID":
			m.SetUniqueExecutionID(r.String())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "OpenClose":
			m.SetOpenClose(OpenCloseTradeEnum(r.Int32()))
		case "NoOrderFills":
			m.SetNoOrderFills(r.Uint8())
		case "InfoText":
			m.SetInfoText(r.String())
		case "HighPriceDuringPosition":
			m.SetHighPriceDuringPosition(r.Float64())
		case "LowPriceDuringPosition":
			m.SetLowPriceDuringPosition(r.Float64())
		case "PositionQuantity":
			m.SetPositionQuantity(r.Float64())
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

func (m *HistoricalOrderFillResponseBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetRequestID(r.ReadInt32())
		case 2:
			m.SetTotalNumberMessages(r.ReadInt32())
		case 3:
			m.SetMessageNumber(r.ReadInt32())
		case 4:
			m.SetSymbol(r.ReadString())
		case 5:
			m.SetExchange(r.ReadString())
		case 6:
			m.SetServerOrderID(r.ReadString())
		case 7:
			m.SetBuySell(BuySellEnum(r.ReadInt32()))
		case 8:
			m.SetPrice(r.ReadFloat64())
		case 9:
			m.SetDateTime(DateTime(r.ReadInt64()))
		case 10:
			m.SetQuantity(r.ReadFloat64())
		case 11:
			m.SetUniqueExecutionID(r.ReadString())
		case 12:
			m.SetTradeAccount(r.ReadString())
		case 13:
			m.SetOpenClose(OpenCloseTradeEnum(r.ReadInt32()))
		case 14:
			m.SetNoOrderFills(r.ReadUint8())
		case 15:
			m.SetInfoText(r.ReadString())
		case 16:
			m.SetHighPriceDuringPosition(r.ReadFloat64())
		case 17:
			m.SetLowPriceDuringPosition(r.ReadFloat64())
		case 18:
			m.SetPositionQuantity(r.ReadFloat64())
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

func (m *HistoricalOrderFillResponsePointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetRequestID(r.ReadInt32())
		case 2:
			m.SetTotalNumberMessages(r.ReadInt32())
		case 3:
			m.SetMessageNumber(r.ReadInt32())
		case 4:
			m.SetSymbol(r.ReadString())
		case 5:
			m.SetExchange(r.ReadString())
		case 6:
			m.SetServerOrderID(r.ReadString())
		case 7:
			m.SetBuySell(BuySellEnum(r.ReadInt32()))
		case 8:
			m.SetPrice(r.ReadFloat64())
		case 9:
			m.SetDateTime(DateTime(r.ReadInt64()))
		case 10:
			m.SetQuantity(r.ReadFloat64())
		case 11:
			m.SetUniqueExecutionID(r.ReadString())
		case 12:
			m.SetTradeAccount(r.ReadString())
		case 13:
			m.SetOpenClose(OpenCloseTradeEnum(r.ReadInt32()))
		case 14:
			m.SetNoOrderFills(r.ReadUint8())
		case 15:
			m.SetInfoText(r.ReadString())
		case 16:
			m.SetHighPriceDuringPosition(r.ReadFloat64())
		case 17:
			m.SetLowPriceDuringPosition(r.ReadFloat64())
		case 18:
			m.SetPositionQuantity(r.ReadFloat64())
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

func (m HistoricalOrderFillResponseBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 304)
	w.WriteVarint32(1, m.RequestID())
	w.WriteVarint32(2, m.TotalNumberMessages())
	w.WriteVarint32(3, m.MessageNumber())
	w.WriteString(4, m.Symbol())
	w.WriteString(5, m.Exchange())
	w.WriteString(6, m.ServerOrderID())
	w.WriteVarint32(7, int32(m.BuySell()))
	w.WriteFixed64Float64(8, m.Price())
	w.WriteFixed64Int64(9, int64(m.DateTime()))
	w.WriteFixed64Float64(10, m.Quantity())
	w.WriteString(11, m.UniqueExecutionID())
	w.WriteString(12, m.TradeAccount())
	w.WriteVarint32(13, int32(m.OpenClose()))
	w.WriteUvarint8(14, m.NoOrderFills())
	w.WriteString(15, m.InfoText())
	w.WriteFixed64Float64(16, m.HighPriceDuringPosition())
	w.WriteFixed64Float64(17, m.LowPriceDuringPosition())
	w.WriteFixed64Float64(18, m.PositionQuantity())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalOrderFillResponsePointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 304)
	w.WriteVarint32(1, m.RequestID())
	w.WriteVarint32(2, m.TotalNumberMessages())
	w.WriteVarint32(3, m.MessageNumber())
	w.WriteString(4, m.Symbol())
	w.WriteString(5, m.Exchange())
	w.WriteString(6, m.ServerOrderID())
	w.WriteVarint32(7, int32(m.BuySell()))
	w.WriteFixed64Float64(8, m.Price())
	w.WriteFixed64Int64(9, int64(m.DateTime()))
	w.WriteFixed64Float64(10, m.Quantity())
	w.WriteString(11, m.UniqueExecutionID())
	w.WriteString(12, m.TradeAccount())
	w.WriteVarint32(13, int32(m.OpenClose()))
	w.WriteUvarint8(14, m.NoOrderFills())
	w.WriteString(15, m.InfoText())
	w.WriteFixed64Float64(16, m.HighPriceDuringPosition())
	w.WriteFixed64Float64(17, m.LowPriceDuringPosition())
	w.WriteFixed64Float64(18, m.PositionQuantity())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalOrderFillResponseFixedBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetRequestID(r.ReadInt32())
		case 2:
			m.SetTotalNumberMessages(r.ReadInt32())
		case 3:
			m.SetMessageNumber(r.ReadInt32())
		case 4:
			m.SetSymbol(r.ReadString())
		case 5:
			m.SetExchange(r.ReadString())
		case 6:
			m.SetServerOrderID(r.ReadString())
		case 7:
			m.SetBuySell(BuySellEnum(r.ReadInt32()))
		case 8:
			m.SetPrice(r.ReadFloat64())
		case 9:
			m.SetDateTime(DateTime(r.ReadInt64()))
		case 10:
			m.SetQuantity(r.ReadFloat64())
		case 11:
			m.SetUniqueExecutionID(r.ReadString())
		case 12:
			m.SetTradeAccount(r.ReadString())
		case 13:
			m.SetOpenClose(OpenCloseTradeEnum(r.ReadInt32()))
		case 14:
			m.SetNoOrderFills(r.ReadUint8())
		case 15:
			m.SetInfoText(r.ReadString())
		case 16:
			m.SetHighPriceDuringPosition(r.ReadFloat64())
		case 17:
			m.SetLowPriceDuringPosition(r.ReadFloat64())
		case 18:
			m.SetPositionQuantity(r.ReadFloat64())
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

func (m *HistoricalOrderFillResponseFixedPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetRequestID(r.ReadInt32())
		case 2:
			m.SetTotalNumberMessages(r.ReadInt32())
		case 3:
			m.SetMessageNumber(r.ReadInt32())
		case 4:
			m.SetSymbol(r.ReadString())
		case 5:
			m.SetExchange(r.ReadString())
		case 6:
			m.SetServerOrderID(r.ReadString())
		case 7:
			m.SetBuySell(BuySellEnum(r.ReadInt32()))
		case 8:
			m.SetPrice(r.ReadFloat64())
		case 9:
			m.SetDateTime(DateTime(r.ReadInt64()))
		case 10:
			m.SetQuantity(r.ReadFloat64())
		case 11:
			m.SetUniqueExecutionID(r.ReadString())
		case 12:
			m.SetTradeAccount(r.ReadString())
		case 13:
			m.SetOpenClose(OpenCloseTradeEnum(r.ReadInt32()))
		case 14:
			m.SetNoOrderFills(r.ReadUint8())
		case 15:
			m.SetInfoText(r.ReadString())
		case 16:
			m.SetHighPriceDuringPosition(r.ReadFloat64())
		case 17:
			m.SetLowPriceDuringPosition(r.ReadFloat64())
		case 18:
			m.SetPositionQuantity(r.ReadFloat64())
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

func (m HistoricalOrderFillResponseFixedBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 304)
	w.WriteVarint32(1, m.RequestID())
	w.WriteVarint32(2, m.TotalNumberMessages())
	w.WriteVarint32(3, m.MessageNumber())
	w.WriteString(4, m.Symbol())
	w.WriteString(5, m.Exchange())
	w.WriteString(6, m.ServerOrderID())
	w.WriteVarint32(7, int32(m.BuySell()))
	w.WriteFixed64Float64(8, m.Price())
	w.WriteFixed64Int64(9, int64(m.DateTime()))
	w.WriteFixed64Float64(10, m.Quantity())
	w.WriteString(11, m.UniqueExecutionID())
	w.WriteString(12, m.TradeAccount())
	w.WriteVarint32(13, int32(m.OpenClose()))
	w.WriteUvarint8(14, m.NoOrderFills())
	w.WriteString(15, m.InfoText())
	w.WriteFixed64Float64(16, m.HighPriceDuringPosition())
	w.WriteFixed64Float64(17, m.LowPriceDuringPosition())
	w.WriteFixed64Float64(18, m.PositionQuantity())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalOrderFillResponseFixedPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 304)
	w.WriteVarint32(1, m.RequestID())
	w.WriteVarint32(2, m.TotalNumberMessages())
	w.WriteVarint32(3, m.MessageNumber())
	w.WriteString(4, m.Symbol())
	w.WriteString(5, m.Exchange())
	w.WriteString(6, m.ServerOrderID())
	w.WriteVarint32(7, int32(m.BuySell()))
	w.WriteFixed64Float64(8, m.Price())
	w.WriteFixed64Int64(9, int64(m.DateTime()))
	w.WriteFixed64Float64(10, m.Quantity())
	w.WriteString(11, m.UniqueExecutionID())
	w.WriteString(12, m.TradeAccount())
	w.WriteVarint32(13, int32(m.OpenClose()))
	w.WriteUvarint8(14, m.NoOrderFills())
	w.WriteString(15, m.InfoText())
	w.WriteFixed64Float64(16, m.HighPriceDuringPosition())
	w.WriteFixed64Float64(17, m.LowPriceDuringPosition())
	w.WriteFixed64Float64(18, m.PositionQuantity())
	return w.Finish(), nil
}