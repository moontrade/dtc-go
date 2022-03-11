//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountDataSymbolLimitsResponsePointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10121,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.Float64(m.TradePositionLimit())
	w.Data = append(w.Data, ',')
	w.Float64(m.OrderQuantityLimit())
	w.Data = append(w.Data, ',')
	w.Int32(m.UsePercentOfMargin())
	w.Data = append(w.Data, ',')
	w.Uint8(m.OverrideMarginOtherAccounts())
	w.Data = append(w.Data, ',')
	w.Int32(m.UsePercentOfMarginForDayTrading())
	w.Data = append(w.Data, ',')
	w.Int32(m.NumberOfDaysBeforeLastTradingDateToDisallowOrders())
	return w.Finish(), nil
}

func (m TradeAccountDataSymbolLimitsResponsePointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10121,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.Float64(m.TradePositionLimit())
	w.Data = append(w.Data, ',')
	w.Float64(m.OrderQuantityLimit())
	w.Data = append(w.Data, ',')
	w.Int32(m.UsePercentOfMargin())
	w.Data = append(w.Data, ',')
	w.Uint8(m.OverrideMarginOtherAccounts())
	w.Data = append(w.Data, ',')
	w.Int32(m.UsePercentOfMarginForDayTrading())
	w.Data = append(w.Data, ',')
	w.Int32(m.NumberOfDaysBeforeLastTradingDateToDisallowOrders())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountDataSymbolLimitsResponsePointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10121)
	w.Uint32Field("RequestID", m.RequestID())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("Symbol", m.Symbol())
	w.Float64Field("TradePositionLimit", m.TradePositionLimit())
	w.Float64Field("OrderQuantityLimit", m.OrderQuantityLimit())
	w.Int32Field("UsePercentOfMargin", m.UsePercentOfMargin())
	w.Uint8Field("OverrideMarginOtherAccounts", m.OverrideMarginOtherAccounts())
	w.Int32Field("UsePercentOfMarginForDayTrading", m.UsePercentOfMarginForDayTrading())
	w.Int32Field("NumberOfDaysBeforeLastTradingDateToDisallowOrders", m.NumberOfDaysBeforeLastTradingDateToDisallowOrders())
	return w.Finish(), nil
}

func (m TradeAccountDataSymbolLimitsResponsePointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10121)
	w.Uint32Field("RequestID", m.RequestID())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("Symbol", m.Symbol())
	w.Float64Field("TradePositionLimit", m.TradePositionLimit())
	w.Float64Field("OrderQuantityLimit", m.OrderQuantityLimit())
	w.Int32Field("UsePercentOfMargin", m.UsePercentOfMargin())
	w.Uint8Field("OverrideMarginOtherAccounts", m.OverrideMarginOtherAccounts())
	w.Int32Field("UsePercentOfMarginForDayTrading", m.UsePercentOfMarginForDayTrading())
	w.Int32Field("NumberOfDaysBeforeLastTradingDateToDisallowOrders", m.NumberOfDaysBeforeLastTradingDateToDisallowOrders())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradeAccountDataSymbolLimitsResponsePointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10121 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetSymbol(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradePositionLimit(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetOrderQuantityLimit(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetUsePercentOfMargin(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetOverrideMarginOtherAccounts(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUsePercentOfMarginForDayTrading(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetNumberOfDaysBeforeLastTradingDateToDisallowOrders(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradeAccountDataSymbolLimitsResponsePointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 10121 {
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
			m.SetRequestID(r.Uint32())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "Symbol":
			m.SetSymbol(r.String())
		case "TradePositionLimit":
			m.SetTradePositionLimit(r.Float64())
		case "OrderQuantityLimit":
			m.SetOrderQuantityLimit(r.Float64())
		case "UsePercentOfMargin":
			m.SetUsePercentOfMargin(r.Int32())
		case "OverrideMarginOtherAccounts":
			m.SetOverrideMarginOtherAccounts(r.Uint8())
		case "UsePercentOfMarginForDayTrading":
			m.SetUsePercentOfMarginForDayTrading(r.Int32())
		case "NumberOfDaysBeforeLastTradingDateToDisallowOrders":
			m.SetNumberOfDaysBeforeLastTradingDateToDisallowOrders(r.Int32())
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
