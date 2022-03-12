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

func (m LogonResponse) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":2,\"F\":["...)
	w.Int32(m.ProtocolVersion())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.Result()))
	w.Data = append(w.Data, ',')
	w.String(m.ResultText())
	w.Data = append(w.Data, ',')
	w.String(m.ReconnectAddress())
	w.Data = append(w.Data, ',')
	w.Int32(m.Integer_1())
	w.Data = append(w.Data, ',')
	w.String(m.ServerName())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MarketDepthUpdatesBestBidAndAsk())
	w.Data = append(w.Data, ',')
	w.Bool(m.TradingIsSupported())
	w.Data = append(w.Data, ',')
	w.Bool(m.OCOOrdersSupported())
	w.Data = append(w.Data, ',')
	w.Bool(m.OrderCancelReplaceSupported())
	w.Data = append(w.Data, ',')
	w.String(m.SymbolExchangeDelimiter())
	w.Data = append(w.Data, ',')
	w.Bool(m.SecurityDefinitionsSupported())
	w.Data = append(w.Data, ',')
	w.Bool(m.HistoricalPriceDataSupported())
	w.Data = append(w.Data, ',')
	w.Uint8(m.ResubscribeWhenMarketDataFeedAvailable())
	w.Data = append(w.Data, ',')
	w.Bool(m.MarketDepthIsSupported())
	w.Data = append(w.Data, ',')
	w.Uint8(m.OneHistoricalPriceDataRequestPerConnection())
	w.Data = append(w.Data, ',')
	w.Bool(m.BracketOrdersSupported())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseIntegerPriceOrderMessages())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	w.Data = append(w.Data, ',')
	w.Bool(m.MarketDataSupported())
	return w.Finish(), nil
}

func (m LogonResponseBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":2,\"F\":["...)
	w.Int32(m.ProtocolVersion())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.Result()))
	w.Data = append(w.Data, ',')
	w.String(m.ResultText())
	w.Data = append(w.Data, ',')
	w.String(m.ReconnectAddress())
	w.Data = append(w.Data, ',')
	w.Int32(m.Integer_1())
	w.Data = append(w.Data, ',')
	w.String(m.ServerName())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MarketDepthUpdatesBestBidAndAsk())
	w.Data = append(w.Data, ',')
	w.Bool(m.TradingIsSupported())
	w.Data = append(w.Data, ',')
	w.Bool(m.OCOOrdersSupported())
	w.Data = append(w.Data, ',')
	w.Bool(m.OrderCancelReplaceSupported())
	w.Data = append(w.Data, ',')
	w.String(m.SymbolExchangeDelimiter())
	w.Data = append(w.Data, ',')
	w.Bool(m.SecurityDefinitionsSupported())
	w.Data = append(w.Data, ',')
	w.Bool(m.HistoricalPriceDataSupported())
	w.Data = append(w.Data, ',')
	w.Uint8(m.ResubscribeWhenMarketDataFeedAvailable())
	w.Data = append(w.Data, ',')
	w.Bool(m.MarketDepthIsSupported())
	w.Data = append(w.Data, ',')
	w.Uint8(m.OneHistoricalPriceDataRequestPerConnection())
	w.Data = append(w.Data, ',')
	w.Bool(m.BracketOrdersSupported())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseIntegerPriceOrderMessages())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	w.Data = append(w.Data, ',')
	w.Bool(m.MarketDataSupported())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m LogonResponsePointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":2,\"F\":["...)
	w.Int32(m.ProtocolVersion())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.Result()))
	w.Data = append(w.Data, ',')
	w.String(m.ResultText())
	w.Data = append(w.Data, ',')
	w.String(m.ReconnectAddress())
	w.Data = append(w.Data, ',')
	w.Int32(m.Integer_1())
	w.Data = append(w.Data, ',')
	w.String(m.ServerName())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MarketDepthUpdatesBestBidAndAsk())
	w.Data = append(w.Data, ',')
	w.Bool(m.TradingIsSupported())
	w.Data = append(w.Data, ',')
	w.Bool(m.OCOOrdersSupported())
	w.Data = append(w.Data, ',')
	w.Bool(m.OrderCancelReplaceSupported())
	w.Data = append(w.Data, ',')
	w.String(m.SymbolExchangeDelimiter())
	w.Data = append(w.Data, ',')
	w.Bool(m.SecurityDefinitionsSupported())
	w.Data = append(w.Data, ',')
	w.Bool(m.HistoricalPriceDataSupported())
	w.Data = append(w.Data, ',')
	w.Uint8(m.ResubscribeWhenMarketDataFeedAvailable())
	w.Data = append(w.Data, ',')
	w.Bool(m.MarketDepthIsSupported())
	w.Data = append(w.Data, ',')
	w.Uint8(m.OneHistoricalPriceDataRequestPerConnection())
	w.Data = append(w.Data, ',')
	w.Bool(m.BracketOrdersSupported())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseIntegerPriceOrderMessages())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	w.Data = append(w.Data, ',')
	w.Bool(m.MarketDataSupported())
	return w.Finish(), nil
}

func (m LogonResponsePointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":2,\"F\":["...)
	w.Int32(m.ProtocolVersion())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.Result()))
	w.Data = append(w.Data, ',')
	w.String(m.ResultText())
	w.Data = append(w.Data, ',')
	w.String(m.ReconnectAddress())
	w.Data = append(w.Data, ',')
	w.Int32(m.Integer_1())
	w.Data = append(w.Data, ',')
	w.String(m.ServerName())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MarketDepthUpdatesBestBidAndAsk())
	w.Data = append(w.Data, ',')
	w.Bool(m.TradingIsSupported())
	w.Data = append(w.Data, ',')
	w.Bool(m.OCOOrdersSupported())
	w.Data = append(w.Data, ',')
	w.Bool(m.OrderCancelReplaceSupported())
	w.Data = append(w.Data, ',')
	w.String(m.SymbolExchangeDelimiter())
	w.Data = append(w.Data, ',')
	w.Bool(m.SecurityDefinitionsSupported())
	w.Data = append(w.Data, ',')
	w.Bool(m.HistoricalPriceDataSupported())
	w.Data = append(w.Data, ',')
	w.Uint8(m.ResubscribeWhenMarketDataFeedAvailable())
	w.Data = append(w.Data, ',')
	w.Bool(m.MarketDepthIsSupported())
	w.Data = append(w.Data, ',')
	w.Uint8(m.OneHistoricalPriceDataRequestPerConnection())
	w.Data = append(w.Data, ',')
	w.Bool(m.BracketOrdersSupported())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseIntegerPriceOrderMessages())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	w.Data = append(w.Data, ',')
	w.Bool(m.MarketDataSupported())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m LogonResponse) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 2)
	w.Int32Field("ProtocolVersion", m.ProtocolVersion())
	w.Int32Field("Result", int32(m.Result()))
	w.StringField("ResultText", m.ResultText())
	w.StringField("ReconnectAddress", m.ReconnectAddress())
	w.Int32Field("Integer_1", m.Integer_1())
	w.StringField("ServerName", m.ServerName())
	w.Uint8Field("MarketDepthUpdatesBestBidAndAsk", m.MarketDepthUpdatesBestBidAndAsk())
	w.BoolField("TradingIsSupported", m.TradingIsSupported())
	w.BoolField("OCOOrdersSupported", m.OCOOrdersSupported())
	w.BoolField("OrderCancelReplaceSupported", m.OrderCancelReplaceSupported())
	w.StringField("SymbolExchangeDelimiter", m.SymbolExchangeDelimiter())
	w.BoolField("SecurityDefinitionsSupported", m.SecurityDefinitionsSupported())
	w.BoolField("HistoricalPriceDataSupported", m.HistoricalPriceDataSupported())
	w.Uint8Field("ResubscribeWhenMarketDataFeedAvailable", m.ResubscribeWhenMarketDataFeedAvailable())
	w.BoolField("MarketDepthIsSupported", m.MarketDepthIsSupported())
	w.Uint8Field("OneHistoricalPriceDataRequestPerConnection", m.OneHistoricalPriceDataRequestPerConnection())
	w.BoolField("BracketOrdersSupported", m.BracketOrdersSupported())
	w.BoolField("UseIntegerPriceOrderMessages", m.UseIntegerPriceOrderMessages())
	w.Uint8Field("UsesMultiplePositionsPerSymbolAndTradeAccount", m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	w.BoolField("MarketDataSupported", m.MarketDataSupported())
	return w.Finish(), nil
}

func (m LogonResponseBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 2)
	w.Int32Field("ProtocolVersion", m.ProtocolVersion())
	w.Int32Field("Result", int32(m.Result()))
	w.StringField("ResultText", m.ResultText())
	w.StringField("ReconnectAddress", m.ReconnectAddress())
	w.Int32Field("Integer_1", m.Integer_1())
	w.StringField("ServerName", m.ServerName())
	w.Uint8Field("MarketDepthUpdatesBestBidAndAsk", m.MarketDepthUpdatesBestBidAndAsk())
	w.BoolField("TradingIsSupported", m.TradingIsSupported())
	w.BoolField("OCOOrdersSupported", m.OCOOrdersSupported())
	w.BoolField("OrderCancelReplaceSupported", m.OrderCancelReplaceSupported())
	w.StringField("SymbolExchangeDelimiter", m.SymbolExchangeDelimiter())
	w.BoolField("SecurityDefinitionsSupported", m.SecurityDefinitionsSupported())
	w.BoolField("HistoricalPriceDataSupported", m.HistoricalPriceDataSupported())
	w.Uint8Field("ResubscribeWhenMarketDataFeedAvailable", m.ResubscribeWhenMarketDataFeedAvailable())
	w.BoolField("MarketDepthIsSupported", m.MarketDepthIsSupported())
	w.Uint8Field("OneHistoricalPriceDataRequestPerConnection", m.OneHistoricalPriceDataRequestPerConnection())
	w.BoolField("BracketOrdersSupported", m.BracketOrdersSupported())
	w.BoolField("UseIntegerPriceOrderMessages", m.UseIntegerPriceOrderMessages())
	w.Uint8Field("UsesMultiplePositionsPerSymbolAndTradeAccount", m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	w.BoolField("MarketDataSupported", m.MarketDataSupported())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m LogonResponsePointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 2)
	w.Int32Field("ProtocolVersion", m.ProtocolVersion())
	w.Int32Field("Result", int32(m.Result()))
	w.StringField("ResultText", m.ResultText())
	w.StringField("ReconnectAddress", m.ReconnectAddress())
	w.Int32Field("Integer_1", m.Integer_1())
	w.StringField("ServerName", m.ServerName())
	w.Uint8Field("MarketDepthUpdatesBestBidAndAsk", m.MarketDepthUpdatesBestBidAndAsk())
	w.BoolField("TradingIsSupported", m.TradingIsSupported())
	w.BoolField("OCOOrdersSupported", m.OCOOrdersSupported())
	w.BoolField("OrderCancelReplaceSupported", m.OrderCancelReplaceSupported())
	w.StringField("SymbolExchangeDelimiter", m.SymbolExchangeDelimiter())
	w.BoolField("SecurityDefinitionsSupported", m.SecurityDefinitionsSupported())
	w.BoolField("HistoricalPriceDataSupported", m.HistoricalPriceDataSupported())
	w.Uint8Field("ResubscribeWhenMarketDataFeedAvailable", m.ResubscribeWhenMarketDataFeedAvailable())
	w.BoolField("MarketDepthIsSupported", m.MarketDepthIsSupported())
	w.Uint8Field("OneHistoricalPriceDataRequestPerConnection", m.OneHistoricalPriceDataRequestPerConnection())
	w.BoolField("BracketOrdersSupported", m.BracketOrdersSupported())
	w.BoolField("UseIntegerPriceOrderMessages", m.UseIntegerPriceOrderMessages())
	w.Uint8Field("UsesMultiplePositionsPerSymbolAndTradeAccount", m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	w.BoolField("MarketDataSupported", m.MarketDataSupported())
	return w.Finish(), nil
}

func (m LogonResponsePointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 2)
	w.Int32Field("ProtocolVersion", m.ProtocolVersion())
	w.Int32Field("Result", int32(m.Result()))
	w.StringField("ResultText", m.ResultText())
	w.StringField("ReconnectAddress", m.ReconnectAddress())
	w.Int32Field("Integer_1", m.Integer_1())
	w.StringField("ServerName", m.ServerName())
	w.Uint8Field("MarketDepthUpdatesBestBidAndAsk", m.MarketDepthUpdatesBestBidAndAsk())
	w.BoolField("TradingIsSupported", m.TradingIsSupported())
	w.BoolField("OCOOrdersSupported", m.OCOOrdersSupported())
	w.BoolField("OrderCancelReplaceSupported", m.OrderCancelReplaceSupported())
	w.StringField("SymbolExchangeDelimiter", m.SymbolExchangeDelimiter())
	w.BoolField("SecurityDefinitionsSupported", m.SecurityDefinitionsSupported())
	w.BoolField("HistoricalPriceDataSupported", m.HistoricalPriceDataSupported())
	w.Uint8Field("ResubscribeWhenMarketDataFeedAvailable", m.ResubscribeWhenMarketDataFeedAvailable())
	w.BoolField("MarketDepthIsSupported", m.MarketDepthIsSupported())
	w.Uint8Field("OneHistoricalPriceDataRequestPerConnection", m.OneHistoricalPriceDataRequestPerConnection())
	w.BoolField("BracketOrdersSupported", m.BracketOrdersSupported())
	w.BoolField("UseIntegerPriceOrderMessages", m.UseIntegerPriceOrderMessages())
	w.Uint8Field("UsesMultiplePositionsPerSymbolAndTradeAccount", m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	w.BoolField("MarketDataSupported", m.MarketDataSupported())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *LogonResponseBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 2 {
		return message.ErrWrongType
	}
	m.SetProtocolVersion(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetResult(LogonStatusEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetResultText(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetReconnectAddress(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetInteger_1(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetServerName(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetMarketDepthUpdatesBestBidAndAsk(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradingIsSupported(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetOCOOrdersSupported(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetOrderCancelReplaceSupported(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetSymbolExchangeDelimiter(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetSecurityDefinitionsSupported(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetHistoricalPriceDataSupported(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetResubscribeWhenMarketDataFeedAvailable(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetMarketDepthIsSupported(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetOneHistoricalPriceDataRequestPerConnection(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetBracketOrdersSupported(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseIntegerPriceOrderMessages(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUsesMultiplePositionsPerSymbolAndTradeAccount(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetMarketDataSupported(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *LogonResponsePointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 2 {
		return message.ErrWrongType
	}
	m.SetProtocolVersion(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetResult(LogonStatusEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetResultText(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetReconnectAddress(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetInteger_1(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetServerName(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetMarketDepthUpdatesBestBidAndAsk(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradingIsSupported(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetOCOOrdersSupported(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetOrderCancelReplaceSupported(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetSymbolExchangeDelimiter(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetSecurityDefinitionsSupported(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetHistoricalPriceDataSupported(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetResubscribeWhenMarketDataFeedAvailable(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetMarketDepthIsSupported(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetOneHistoricalPriceDataRequestPerConnection(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetBracketOrdersSupported(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseIntegerPriceOrderMessages(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUsesMultiplePositionsPerSymbolAndTradeAccount(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetMarketDataSupported(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *LogonResponseBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 2 {
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
		case "Result":
			m.SetResult(LogonStatusEnum(r.Int32()))
		case "ResultText":
			m.SetResultText(r.String())
		case "ReconnectAddress":
			m.SetReconnectAddress(r.String())
		case "Integer_1":
			m.SetInteger_1(r.Int32())
		case "ServerName":
			m.SetServerName(r.String())
		case "MarketDepthUpdatesBestBidAndAsk":
			m.SetMarketDepthUpdatesBestBidAndAsk(r.Uint8())
		case "TradingIsSupported":
			m.SetTradingIsSupported(r.Bool())
		case "OCOOrdersSupported":
			m.SetOCOOrdersSupported(r.Bool())
		case "OrderCancelReplaceSupported":
			m.SetOrderCancelReplaceSupported(r.Bool())
		case "SymbolExchangeDelimiter":
			m.SetSymbolExchangeDelimiter(r.String())
		case "SecurityDefinitionsSupported":
			m.SetSecurityDefinitionsSupported(r.Bool())
		case "HistoricalPriceDataSupported":
			m.SetHistoricalPriceDataSupported(r.Bool())
		case "ResubscribeWhenMarketDataFeedAvailable":
			m.SetResubscribeWhenMarketDataFeedAvailable(r.Uint8())
		case "MarketDepthIsSupported":
			m.SetMarketDepthIsSupported(r.Bool())
		case "OneHistoricalPriceDataRequestPerConnection":
			m.SetOneHistoricalPriceDataRequestPerConnection(r.Uint8())
		case "BracketOrdersSupported":
			m.SetBracketOrdersSupported(r.Bool())
		case "UseIntegerPriceOrderMessages":
			m.SetUseIntegerPriceOrderMessages(r.Bool())
		case "UsesMultiplePositionsPerSymbolAndTradeAccount":
			m.SetUsesMultiplePositionsPerSymbolAndTradeAccount(r.Uint8())
		case "MarketDataSupported":
			m.SetMarketDataSupported(r.Bool())
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

func (m *LogonResponsePointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 2 {
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
		case "Result":
			m.SetResult(LogonStatusEnum(r.Int32()))
		case "ResultText":
			m.SetResultText(r.String())
		case "ReconnectAddress":
			m.SetReconnectAddress(r.String())
		case "Integer_1":
			m.SetInteger_1(r.Int32())
		case "ServerName":
			m.SetServerName(r.String())
		case "MarketDepthUpdatesBestBidAndAsk":
			m.SetMarketDepthUpdatesBestBidAndAsk(r.Uint8())
		case "TradingIsSupported":
			m.SetTradingIsSupported(r.Bool())
		case "OCOOrdersSupported":
			m.SetOCOOrdersSupported(r.Bool())
		case "OrderCancelReplaceSupported":
			m.SetOrderCancelReplaceSupported(r.Bool())
		case "SymbolExchangeDelimiter":
			m.SetSymbolExchangeDelimiter(r.String())
		case "SecurityDefinitionsSupported":
			m.SetSecurityDefinitionsSupported(r.Bool())
		case "HistoricalPriceDataSupported":
			m.SetHistoricalPriceDataSupported(r.Bool())
		case "ResubscribeWhenMarketDataFeedAvailable":
			m.SetResubscribeWhenMarketDataFeedAvailable(r.Uint8())
		case "MarketDepthIsSupported":
			m.SetMarketDepthIsSupported(r.Bool())
		case "OneHistoricalPriceDataRequestPerConnection":
			m.SetOneHistoricalPriceDataRequestPerConnection(r.Uint8())
		case "BracketOrdersSupported":
			m.SetBracketOrdersSupported(r.Bool())
		case "UseIntegerPriceOrderMessages":
			m.SetUseIntegerPriceOrderMessages(r.Bool())
		case "UsesMultiplePositionsPerSymbolAndTradeAccount":
			m.SetUsesMultiplePositionsPerSymbolAndTradeAccount(r.Uint8())
		case "MarketDataSupported":
			m.SetMarketDataSupported(r.Bool())
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

func (m LogonResponseFixed) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":2,\"F\":["...)
	w.Int32(m.ProtocolVersion())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.Result()))
	w.Data = append(w.Data, ',')
	w.String(m.ResultText())
	w.Data = append(w.Data, ',')
	w.String(m.ReconnectAddress())
	w.Data = append(w.Data, ',')
	w.Int32(m.Integer_1())
	w.Data = append(w.Data, ',')
	w.String(m.ServerName())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MarketDepthUpdatesBestBidAndAsk())
	w.Data = append(w.Data, ',')
	w.Bool(m.TradingIsSupported())
	w.Data = append(w.Data, ',')
	w.Bool(m.OCOOrdersSupported())
	w.Data = append(w.Data, ',')
	w.Bool(m.OrderCancelReplaceSupported())
	w.Data = append(w.Data, ',')
	w.String(m.SymbolExchangeDelimiter())
	w.Data = append(w.Data, ',')
	w.Bool(m.SecurityDefinitionsSupported())
	w.Data = append(w.Data, ',')
	w.Bool(m.HistoricalPriceDataSupported())
	w.Data = append(w.Data, ',')
	w.Uint8(m.ResubscribeWhenMarketDataFeedAvailable())
	w.Data = append(w.Data, ',')
	w.Bool(m.MarketDepthIsSupported())
	w.Data = append(w.Data, ',')
	w.Uint8(m.OneHistoricalPriceDataRequestPerConnection())
	w.Data = append(w.Data, ',')
	w.Bool(m.BracketOrdersSupported())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseIntegerPriceOrderMessages())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	w.Data = append(w.Data, ',')
	w.Bool(m.MarketDataSupported())
	return w.Finish(), nil
}

func (m LogonResponseFixedBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":2,\"F\":["...)
	w.Int32(m.ProtocolVersion())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.Result()))
	w.Data = append(w.Data, ',')
	w.String(m.ResultText())
	w.Data = append(w.Data, ',')
	w.String(m.ReconnectAddress())
	w.Data = append(w.Data, ',')
	w.Int32(m.Integer_1())
	w.Data = append(w.Data, ',')
	w.String(m.ServerName())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MarketDepthUpdatesBestBidAndAsk())
	w.Data = append(w.Data, ',')
	w.Bool(m.TradingIsSupported())
	w.Data = append(w.Data, ',')
	w.Bool(m.OCOOrdersSupported())
	w.Data = append(w.Data, ',')
	w.Bool(m.OrderCancelReplaceSupported())
	w.Data = append(w.Data, ',')
	w.String(m.SymbolExchangeDelimiter())
	w.Data = append(w.Data, ',')
	w.Bool(m.SecurityDefinitionsSupported())
	w.Data = append(w.Data, ',')
	w.Bool(m.HistoricalPriceDataSupported())
	w.Data = append(w.Data, ',')
	w.Uint8(m.ResubscribeWhenMarketDataFeedAvailable())
	w.Data = append(w.Data, ',')
	w.Bool(m.MarketDepthIsSupported())
	w.Data = append(w.Data, ',')
	w.Uint8(m.OneHistoricalPriceDataRequestPerConnection())
	w.Data = append(w.Data, ',')
	w.Bool(m.BracketOrdersSupported())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseIntegerPriceOrderMessages())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	w.Data = append(w.Data, ',')
	w.Bool(m.MarketDataSupported())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m LogonResponseFixedPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":2,\"F\":["...)
	w.Int32(m.ProtocolVersion())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.Result()))
	w.Data = append(w.Data, ',')
	w.String(m.ResultText())
	w.Data = append(w.Data, ',')
	w.String(m.ReconnectAddress())
	w.Data = append(w.Data, ',')
	w.Int32(m.Integer_1())
	w.Data = append(w.Data, ',')
	w.String(m.ServerName())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MarketDepthUpdatesBestBidAndAsk())
	w.Data = append(w.Data, ',')
	w.Bool(m.TradingIsSupported())
	w.Data = append(w.Data, ',')
	w.Bool(m.OCOOrdersSupported())
	w.Data = append(w.Data, ',')
	w.Bool(m.OrderCancelReplaceSupported())
	w.Data = append(w.Data, ',')
	w.String(m.SymbolExchangeDelimiter())
	w.Data = append(w.Data, ',')
	w.Bool(m.SecurityDefinitionsSupported())
	w.Data = append(w.Data, ',')
	w.Bool(m.HistoricalPriceDataSupported())
	w.Data = append(w.Data, ',')
	w.Uint8(m.ResubscribeWhenMarketDataFeedAvailable())
	w.Data = append(w.Data, ',')
	w.Bool(m.MarketDepthIsSupported())
	w.Data = append(w.Data, ',')
	w.Uint8(m.OneHistoricalPriceDataRequestPerConnection())
	w.Data = append(w.Data, ',')
	w.Bool(m.BracketOrdersSupported())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseIntegerPriceOrderMessages())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	w.Data = append(w.Data, ',')
	w.Bool(m.MarketDataSupported())
	return w.Finish(), nil
}

func (m LogonResponseFixedPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":2,\"F\":["...)
	w.Int32(m.ProtocolVersion())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.Result()))
	w.Data = append(w.Data, ',')
	w.String(m.ResultText())
	w.Data = append(w.Data, ',')
	w.String(m.ReconnectAddress())
	w.Data = append(w.Data, ',')
	w.Int32(m.Integer_1())
	w.Data = append(w.Data, ',')
	w.String(m.ServerName())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MarketDepthUpdatesBestBidAndAsk())
	w.Data = append(w.Data, ',')
	w.Bool(m.TradingIsSupported())
	w.Data = append(w.Data, ',')
	w.Bool(m.OCOOrdersSupported())
	w.Data = append(w.Data, ',')
	w.Bool(m.OrderCancelReplaceSupported())
	w.Data = append(w.Data, ',')
	w.String(m.SymbolExchangeDelimiter())
	w.Data = append(w.Data, ',')
	w.Bool(m.SecurityDefinitionsSupported())
	w.Data = append(w.Data, ',')
	w.Bool(m.HistoricalPriceDataSupported())
	w.Data = append(w.Data, ',')
	w.Uint8(m.ResubscribeWhenMarketDataFeedAvailable())
	w.Data = append(w.Data, ',')
	w.Bool(m.MarketDepthIsSupported())
	w.Data = append(w.Data, ',')
	w.Uint8(m.OneHistoricalPriceDataRequestPerConnection())
	w.Data = append(w.Data, ',')
	w.Bool(m.BracketOrdersSupported())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseIntegerPriceOrderMessages())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	w.Data = append(w.Data, ',')
	w.Bool(m.MarketDataSupported())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m LogonResponseFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 2)
	w.Int32Field("ProtocolVersion", m.ProtocolVersion())
	w.Int32Field("Result", int32(m.Result()))
	w.StringField("ResultText", m.ResultText())
	w.StringField("ReconnectAddress", m.ReconnectAddress())
	w.Int32Field("Integer_1", m.Integer_1())
	w.StringField("ServerName", m.ServerName())
	w.Uint8Field("MarketDepthUpdatesBestBidAndAsk", m.MarketDepthUpdatesBestBidAndAsk())
	w.BoolField("TradingIsSupported", m.TradingIsSupported())
	w.BoolField("OCOOrdersSupported", m.OCOOrdersSupported())
	w.BoolField("OrderCancelReplaceSupported", m.OrderCancelReplaceSupported())
	w.StringField("SymbolExchangeDelimiter", m.SymbolExchangeDelimiter())
	w.BoolField("SecurityDefinitionsSupported", m.SecurityDefinitionsSupported())
	w.BoolField("HistoricalPriceDataSupported", m.HistoricalPriceDataSupported())
	w.Uint8Field("ResubscribeWhenMarketDataFeedAvailable", m.ResubscribeWhenMarketDataFeedAvailable())
	w.BoolField("MarketDepthIsSupported", m.MarketDepthIsSupported())
	w.Uint8Field("OneHistoricalPriceDataRequestPerConnection", m.OneHistoricalPriceDataRequestPerConnection())
	w.BoolField("BracketOrdersSupported", m.BracketOrdersSupported())
	w.BoolField("UseIntegerPriceOrderMessages", m.UseIntegerPriceOrderMessages())
	w.Uint8Field("UsesMultiplePositionsPerSymbolAndTradeAccount", m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	w.BoolField("MarketDataSupported", m.MarketDataSupported())
	return w.Finish(), nil
}

func (m LogonResponseFixedBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 2)
	w.Int32Field("ProtocolVersion", m.ProtocolVersion())
	w.Int32Field("Result", int32(m.Result()))
	w.StringField("ResultText", m.ResultText())
	w.StringField("ReconnectAddress", m.ReconnectAddress())
	w.Int32Field("Integer_1", m.Integer_1())
	w.StringField("ServerName", m.ServerName())
	w.Uint8Field("MarketDepthUpdatesBestBidAndAsk", m.MarketDepthUpdatesBestBidAndAsk())
	w.BoolField("TradingIsSupported", m.TradingIsSupported())
	w.BoolField("OCOOrdersSupported", m.OCOOrdersSupported())
	w.BoolField("OrderCancelReplaceSupported", m.OrderCancelReplaceSupported())
	w.StringField("SymbolExchangeDelimiter", m.SymbolExchangeDelimiter())
	w.BoolField("SecurityDefinitionsSupported", m.SecurityDefinitionsSupported())
	w.BoolField("HistoricalPriceDataSupported", m.HistoricalPriceDataSupported())
	w.Uint8Field("ResubscribeWhenMarketDataFeedAvailable", m.ResubscribeWhenMarketDataFeedAvailable())
	w.BoolField("MarketDepthIsSupported", m.MarketDepthIsSupported())
	w.Uint8Field("OneHistoricalPriceDataRequestPerConnection", m.OneHistoricalPriceDataRequestPerConnection())
	w.BoolField("BracketOrdersSupported", m.BracketOrdersSupported())
	w.BoolField("UseIntegerPriceOrderMessages", m.UseIntegerPriceOrderMessages())
	w.Uint8Field("UsesMultiplePositionsPerSymbolAndTradeAccount", m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	w.BoolField("MarketDataSupported", m.MarketDataSupported())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m LogonResponseFixedPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 2)
	w.Int32Field("ProtocolVersion", m.ProtocolVersion())
	w.Int32Field("Result", int32(m.Result()))
	w.StringField("ResultText", m.ResultText())
	w.StringField("ReconnectAddress", m.ReconnectAddress())
	w.Int32Field("Integer_1", m.Integer_1())
	w.StringField("ServerName", m.ServerName())
	w.Uint8Field("MarketDepthUpdatesBestBidAndAsk", m.MarketDepthUpdatesBestBidAndAsk())
	w.BoolField("TradingIsSupported", m.TradingIsSupported())
	w.BoolField("OCOOrdersSupported", m.OCOOrdersSupported())
	w.BoolField("OrderCancelReplaceSupported", m.OrderCancelReplaceSupported())
	w.StringField("SymbolExchangeDelimiter", m.SymbolExchangeDelimiter())
	w.BoolField("SecurityDefinitionsSupported", m.SecurityDefinitionsSupported())
	w.BoolField("HistoricalPriceDataSupported", m.HistoricalPriceDataSupported())
	w.Uint8Field("ResubscribeWhenMarketDataFeedAvailable", m.ResubscribeWhenMarketDataFeedAvailable())
	w.BoolField("MarketDepthIsSupported", m.MarketDepthIsSupported())
	w.Uint8Field("OneHistoricalPriceDataRequestPerConnection", m.OneHistoricalPriceDataRequestPerConnection())
	w.BoolField("BracketOrdersSupported", m.BracketOrdersSupported())
	w.BoolField("UseIntegerPriceOrderMessages", m.UseIntegerPriceOrderMessages())
	w.Uint8Field("UsesMultiplePositionsPerSymbolAndTradeAccount", m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	w.BoolField("MarketDataSupported", m.MarketDataSupported())
	return w.Finish(), nil
}

func (m LogonResponseFixedPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 2)
	w.Int32Field("ProtocolVersion", m.ProtocolVersion())
	w.Int32Field("Result", int32(m.Result()))
	w.StringField("ResultText", m.ResultText())
	w.StringField("ReconnectAddress", m.ReconnectAddress())
	w.Int32Field("Integer_1", m.Integer_1())
	w.StringField("ServerName", m.ServerName())
	w.Uint8Field("MarketDepthUpdatesBestBidAndAsk", m.MarketDepthUpdatesBestBidAndAsk())
	w.BoolField("TradingIsSupported", m.TradingIsSupported())
	w.BoolField("OCOOrdersSupported", m.OCOOrdersSupported())
	w.BoolField("OrderCancelReplaceSupported", m.OrderCancelReplaceSupported())
	w.StringField("SymbolExchangeDelimiter", m.SymbolExchangeDelimiter())
	w.BoolField("SecurityDefinitionsSupported", m.SecurityDefinitionsSupported())
	w.BoolField("HistoricalPriceDataSupported", m.HistoricalPriceDataSupported())
	w.Uint8Field("ResubscribeWhenMarketDataFeedAvailable", m.ResubscribeWhenMarketDataFeedAvailable())
	w.BoolField("MarketDepthIsSupported", m.MarketDepthIsSupported())
	w.Uint8Field("OneHistoricalPriceDataRequestPerConnection", m.OneHistoricalPriceDataRequestPerConnection())
	w.BoolField("BracketOrdersSupported", m.BracketOrdersSupported())
	w.BoolField("UseIntegerPriceOrderMessages", m.UseIntegerPriceOrderMessages())
	w.Uint8Field("UsesMultiplePositionsPerSymbolAndTradeAccount", m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	w.BoolField("MarketDataSupported", m.MarketDataSupported())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *LogonResponseFixedBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 2 {
		return message.ErrWrongType
	}
	m.SetProtocolVersion(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetResult(LogonStatusEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetResultText(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetReconnectAddress(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetInteger_1(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetServerName(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetMarketDepthUpdatesBestBidAndAsk(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradingIsSupported(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetOCOOrdersSupported(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetOrderCancelReplaceSupported(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetSymbolExchangeDelimiter(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetSecurityDefinitionsSupported(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetHistoricalPriceDataSupported(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetResubscribeWhenMarketDataFeedAvailable(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetMarketDepthIsSupported(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetOneHistoricalPriceDataRequestPerConnection(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetBracketOrdersSupported(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseIntegerPriceOrderMessages(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUsesMultiplePositionsPerSymbolAndTradeAccount(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetMarketDataSupported(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *LogonResponseFixedPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 2 {
		return message.ErrWrongType
	}
	m.SetProtocolVersion(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetResult(LogonStatusEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetResultText(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetReconnectAddress(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetInteger_1(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetServerName(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetMarketDepthUpdatesBestBidAndAsk(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradingIsSupported(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetOCOOrdersSupported(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetOrderCancelReplaceSupported(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetSymbolExchangeDelimiter(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetSecurityDefinitionsSupported(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetHistoricalPriceDataSupported(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetResubscribeWhenMarketDataFeedAvailable(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetMarketDepthIsSupported(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetOneHistoricalPriceDataRequestPerConnection(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetBracketOrdersSupported(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseIntegerPriceOrderMessages(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUsesMultiplePositionsPerSymbolAndTradeAccount(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetMarketDataSupported(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *LogonResponseFixedBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 2 {
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
		case "Result":
			m.SetResult(LogonStatusEnum(r.Int32()))
		case "ResultText":
			m.SetResultText(r.String())
		case "ReconnectAddress":
			m.SetReconnectAddress(r.String())
		case "Integer_1":
			m.SetInteger_1(r.Int32())
		case "ServerName":
			m.SetServerName(r.String())
		case "MarketDepthUpdatesBestBidAndAsk":
			m.SetMarketDepthUpdatesBestBidAndAsk(r.Uint8())
		case "TradingIsSupported":
			m.SetTradingIsSupported(r.Bool())
		case "OCOOrdersSupported":
			m.SetOCOOrdersSupported(r.Bool())
		case "OrderCancelReplaceSupported":
			m.SetOrderCancelReplaceSupported(r.Bool())
		case "SymbolExchangeDelimiter":
			m.SetSymbolExchangeDelimiter(r.String())
		case "SecurityDefinitionsSupported":
			m.SetSecurityDefinitionsSupported(r.Bool())
		case "HistoricalPriceDataSupported":
			m.SetHistoricalPriceDataSupported(r.Bool())
		case "ResubscribeWhenMarketDataFeedAvailable":
			m.SetResubscribeWhenMarketDataFeedAvailable(r.Uint8())
		case "MarketDepthIsSupported":
			m.SetMarketDepthIsSupported(r.Bool())
		case "OneHistoricalPriceDataRequestPerConnection":
			m.SetOneHistoricalPriceDataRequestPerConnection(r.Uint8())
		case "BracketOrdersSupported":
			m.SetBracketOrdersSupported(r.Bool())
		case "UseIntegerPriceOrderMessages":
			m.SetUseIntegerPriceOrderMessages(r.Bool())
		case "UsesMultiplePositionsPerSymbolAndTradeAccount":
			m.SetUsesMultiplePositionsPerSymbolAndTradeAccount(r.Uint8())
		case "MarketDataSupported":
			m.SetMarketDataSupported(r.Bool())
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

func (m *LogonResponseFixedPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 2 {
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
		case "Result":
			m.SetResult(LogonStatusEnum(r.Int32()))
		case "ResultText":
			m.SetResultText(r.String())
		case "ReconnectAddress":
			m.SetReconnectAddress(r.String())
		case "Integer_1":
			m.SetInteger_1(r.Int32())
		case "ServerName":
			m.SetServerName(r.String())
		case "MarketDepthUpdatesBestBidAndAsk":
			m.SetMarketDepthUpdatesBestBidAndAsk(r.Uint8())
		case "TradingIsSupported":
			m.SetTradingIsSupported(r.Bool())
		case "OCOOrdersSupported":
			m.SetOCOOrdersSupported(r.Bool())
		case "OrderCancelReplaceSupported":
			m.SetOrderCancelReplaceSupported(r.Bool())
		case "SymbolExchangeDelimiter":
			m.SetSymbolExchangeDelimiter(r.String())
		case "SecurityDefinitionsSupported":
			m.SetSecurityDefinitionsSupported(r.Bool())
		case "HistoricalPriceDataSupported":
			m.SetHistoricalPriceDataSupported(r.Bool())
		case "ResubscribeWhenMarketDataFeedAvailable":
			m.SetResubscribeWhenMarketDataFeedAvailable(r.Uint8())
		case "MarketDepthIsSupported":
			m.SetMarketDepthIsSupported(r.Bool())
		case "OneHistoricalPriceDataRequestPerConnection":
			m.SetOneHistoricalPriceDataRequestPerConnection(r.Uint8())
		case "BracketOrdersSupported":
			m.SetBracketOrdersSupported(r.Bool())
		case "UseIntegerPriceOrderMessages":
			m.SetUseIntegerPriceOrderMessages(r.Bool())
		case "UsesMultiplePositionsPerSymbolAndTradeAccount":
			m.SetUsesMultiplePositionsPerSymbolAndTradeAccount(r.Uint8())
		case "MarketDataSupported":
			m.SetMarketDataSupported(r.Bool())
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

func (m *LogonResponseBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetProtocolVersion(r.ReadInt32())
		case 2:
			m.SetResult(LogonStatusEnum(r.ReadInt32()))
		case 3:
			m.SetResultText(r.ReadString())
		case 4:
			m.SetReconnectAddress(r.ReadString())
		case 5:
			m.SetInteger_1(r.ReadInt32())
		case 6:
			m.SetServerName(r.ReadString())
		case 7:
			m.SetMarketDepthUpdatesBestBidAndAsk(r.ReadUint8())
		case 8:
			m.SetTradingIsSupported(r.ReadBool())
		case 9:
			m.SetOCOOrdersSupported(r.ReadBool())
		case 10:
			m.SetOrderCancelReplaceSupported(r.ReadBool())
		case 11:
			m.SetSymbolExchangeDelimiter(r.ReadString())
		case 12:
			m.SetSecurityDefinitionsSupported(r.ReadBool())
		case 13:
			m.SetHistoricalPriceDataSupported(r.ReadBool())
		case 14:
			m.SetResubscribeWhenMarketDataFeedAvailable(r.ReadUint8())
		case 15:
			m.SetMarketDepthIsSupported(r.ReadBool())
		case 16:
			m.SetOneHistoricalPriceDataRequestPerConnection(r.ReadUint8())
		case 17:
			m.SetBracketOrdersSupported(r.ReadBool())
		case 18:
			m.SetUseIntegerPriceOrderMessages(r.ReadBool())
		case 19:
			m.SetUsesMultiplePositionsPerSymbolAndTradeAccount(r.ReadUint8())
		case 20:
			m.SetMarketDataSupported(r.ReadBool())
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

func (m *LogonResponsePointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetProtocolVersion(r.ReadInt32())
		case 2:
			m.SetResult(LogonStatusEnum(r.ReadInt32()))
		case 3:
			m.SetResultText(r.ReadString())
		case 4:
			m.SetReconnectAddress(r.ReadString())
		case 5:
			m.SetInteger_1(r.ReadInt32())
		case 6:
			m.SetServerName(r.ReadString())
		case 7:
			m.SetMarketDepthUpdatesBestBidAndAsk(r.ReadUint8())
		case 8:
			m.SetTradingIsSupported(r.ReadBool())
		case 9:
			m.SetOCOOrdersSupported(r.ReadBool())
		case 10:
			m.SetOrderCancelReplaceSupported(r.ReadBool())
		case 11:
			m.SetSymbolExchangeDelimiter(r.ReadString())
		case 12:
			m.SetSecurityDefinitionsSupported(r.ReadBool())
		case 13:
			m.SetHistoricalPriceDataSupported(r.ReadBool())
		case 14:
			m.SetResubscribeWhenMarketDataFeedAvailable(r.ReadUint8())
		case 15:
			m.SetMarketDepthIsSupported(r.ReadBool())
		case 16:
			m.SetOneHistoricalPriceDataRequestPerConnection(r.ReadUint8())
		case 17:
			m.SetBracketOrdersSupported(r.ReadBool())
		case 18:
			m.SetUseIntegerPriceOrderMessages(r.ReadBool())
		case 19:
			m.SetUsesMultiplePositionsPerSymbolAndTradeAccount(r.ReadUint8())
		case 20:
			m.SetMarketDataSupported(r.ReadBool())
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

func (m LogonResponseBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 2)
	w.WriteVarint32(1, m.ProtocolVersion())
	w.WriteVarint32(2, int32(m.Result()))
	w.WriteString(3, m.ResultText())
	w.WriteString(4, m.ReconnectAddress())
	w.WriteVarint32(5, m.Integer_1())
	w.WriteString(6, m.ServerName())
	w.WriteUvarint8(7, m.MarketDepthUpdatesBestBidAndAsk())
	w.WriteBool(8, m.TradingIsSupported())
	w.WriteBool(9, m.OCOOrdersSupported())
	w.WriteBool(10, m.OrderCancelReplaceSupported())
	w.WriteString(11, m.SymbolExchangeDelimiter())
	w.WriteBool(12, m.SecurityDefinitionsSupported())
	w.WriteBool(13, m.HistoricalPriceDataSupported())
	w.WriteUvarint8(14, m.ResubscribeWhenMarketDataFeedAvailable())
	w.WriteBool(15, m.MarketDepthIsSupported())
	w.WriteUvarint8(16, m.OneHistoricalPriceDataRequestPerConnection())
	w.WriteBool(17, m.BracketOrdersSupported())
	w.WriteBool(18, m.UseIntegerPriceOrderMessages())
	w.WriteUvarint8(19, m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	w.WriteBool(20, m.MarketDataSupported())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m LogonResponsePointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 2)
	w.WriteVarint32(1, m.ProtocolVersion())
	w.WriteVarint32(2, int32(m.Result()))
	w.WriteString(3, m.ResultText())
	w.WriteString(4, m.ReconnectAddress())
	w.WriteVarint32(5, m.Integer_1())
	w.WriteString(6, m.ServerName())
	w.WriteUvarint8(7, m.MarketDepthUpdatesBestBidAndAsk())
	w.WriteBool(8, m.TradingIsSupported())
	w.WriteBool(9, m.OCOOrdersSupported())
	w.WriteBool(10, m.OrderCancelReplaceSupported())
	w.WriteString(11, m.SymbolExchangeDelimiter())
	w.WriteBool(12, m.SecurityDefinitionsSupported())
	w.WriteBool(13, m.HistoricalPriceDataSupported())
	w.WriteUvarint8(14, m.ResubscribeWhenMarketDataFeedAvailable())
	w.WriteBool(15, m.MarketDepthIsSupported())
	w.WriteUvarint8(16, m.OneHistoricalPriceDataRequestPerConnection())
	w.WriteBool(17, m.BracketOrdersSupported())
	w.WriteBool(18, m.UseIntegerPriceOrderMessages())
	w.WriteUvarint8(19, m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	w.WriteBool(20, m.MarketDataSupported())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *LogonResponseFixedBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetProtocolVersion(r.ReadInt32())
		case 2:
			m.SetResult(LogonStatusEnum(r.ReadInt32()))
		case 3:
			m.SetResultText(r.ReadString())
		case 4:
			m.SetReconnectAddress(r.ReadString())
		case 5:
			m.SetInteger_1(r.ReadInt32())
		case 6:
			m.SetServerName(r.ReadString())
		case 7:
			m.SetMarketDepthUpdatesBestBidAndAsk(r.ReadUint8())
		case 8:
			m.SetTradingIsSupported(r.ReadBool())
		case 9:
			m.SetOCOOrdersSupported(r.ReadBool())
		case 10:
			m.SetOrderCancelReplaceSupported(r.ReadBool())
		case 11:
			m.SetSymbolExchangeDelimiter(r.ReadString())
		case 12:
			m.SetSecurityDefinitionsSupported(r.ReadBool())
		case 13:
			m.SetHistoricalPriceDataSupported(r.ReadBool())
		case 14:
			m.SetResubscribeWhenMarketDataFeedAvailable(r.ReadUint8())
		case 15:
			m.SetMarketDepthIsSupported(r.ReadBool())
		case 16:
			m.SetOneHistoricalPriceDataRequestPerConnection(r.ReadUint8())
		case 17:
			m.SetBracketOrdersSupported(r.ReadBool())
		case 18:
			m.SetUseIntegerPriceOrderMessages(r.ReadBool())
		case 19:
			m.SetUsesMultiplePositionsPerSymbolAndTradeAccount(r.ReadUint8())
		case 20:
			m.SetMarketDataSupported(r.ReadBool())
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

func (m *LogonResponseFixedPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetProtocolVersion(r.ReadInt32())
		case 2:
			m.SetResult(LogonStatusEnum(r.ReadInt32()))
		case 3:
			m.SetResultText(r.ReadString())
		case 4:
			m.SetReconnectAddress(r.ReadString())
		case 5:
			m.SetInteger_1(r.ReadInt32())
		case 6:
			m.SetServerName(r.ReadString())
		case 7:
			m.SetMarketDepthUpdatesBestBidAndAsk(r.ReadUint8())
		case 8:
			m.SetTradingIsSupported(r.ReadBool())
		case 9:
			m.SetOCOOrdersSupported(r.ReadBool())
		case 10:
			m.SetOrderCancelReplaceSupported(r.ReadBool())
		case 11:
			m.SetSymbolExchangeDelimiter(r.ReadString())
		case 12:
			m.SetSecurityDefinitionsSupported(r.ReadBool())
		case 13:
			m.SetHistoricalPriceDataSupported(r.ReadBool())
		case 14:
			m.SetResubscribeWhenMarketDataFeedAvailable(r.ReadUint8())
		case 15:
			m.SetMarketDepthIsSupported(r.ReadBool())
		case 16:
			m.SetOneHistoricalPriceDataRequestPerConnection(r.ReadUint8())
		case 17:
			m.SetBracketOrdersSupported(r.ReadBool())
		case 18:
			m.SetUseIntegerPriceOrderMessages(r.ReadBool())
		case 19:
			m.SetUsesMultiplePositionsPerSymbolAndTradeAccount(r.ReadUint8())
		case 20:
			m.SetMarketDataSupported(r.ReadBool())
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

func (m LogonResponseFixedBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 2)
	w.WriteVarint32(1, m.ProtocolVersion())
	w.WriteVarint32(2, int32(m.Result()))
	w.WriteString(3, m.ResultText())
	w.WriteString(4, m.ReconnectAddress())
	w.WriteVarint32(5, m.Integer_1())
	w.WriteString(6, m.ServerName())
	w.WriteUvarint8(7, m.MarketDepthUpdatesBestBidAndAsk())
	w.WriteBool(8, m.TradingIsSupported())
	w.WriteBool(9, m.OCOOrdersSupported())
	w.WriteBool(10, m.OrderCancelReplaceSupported())
	w.WriteString(11, m.SymbolExchangeDelimiter())
	w.WriteBool(12, m.SecurityDefinitionsSupported())
	w.WriteBool(13, m.HistoricalPriceDataSupported())
	w.WriteUvarint8(14, m.ResubscribeWhenMarketDataFeedAvailable())
	w.WriteBool(15, m.MarketDepthIsSupported())
	w.WriteUvarint8(16, m.OneHistoricalPriceDataRequestPerConnection())
	w.WriteBool(17, m.BracketOrdersSupported())
	w.WriteBool(18, m.UseIntegerPriceOrderMessages())
	w.WriteUvarint8(19, m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	w.WriteBool(20, m.MarketDataSupported())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m LogonResponseFixedPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 2)
	w.WriteVarint32(1, m.ProtocolVersion())
	w.WriteVarint32(2, int32(m.Result()))
	w.WriteString(3, m.ResultText())
	w.WriteString(4, m.ReconnectAddress())
	w.WriteVarint32(5, m.Integer_1())
	w.WriteString(6, m.ServerName())
	w.WriteUvarint8(7, m.MarketDepthUpdatesBestBidAndAsk())
	w.WriteBool(8, m.TradingIsSupported())
	w.WriteBool(9, m.OCOOrdersSupported())
	w.WriteBool(10, m.OrderCancelReplaceSupported())
	w.WriteString(11, m.SymbolExchangeDelimiter())
	w.WriteBool(12, m.SecurityDefinitionsSupported())
	w.WriteBool(13, m.HistoricalPriceDataSupported())
	w.WriteUvarint8(14, m.ResubscribeWhenMarketDataFeedAvailable())
	w.WriteBool(15, m.MarketDepthIsSupported())
	w.WriteUvarint8(16, m.OneHistoricalPriceDataRequestPerConnection())
	w.WriteBool(17, m.BracketOrdersSupported())
	w.WriteBool(18, m.UseIntegerPriceOrderMessages())
	w.WriteUvarint8(19, m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	w.WriteBool(20, m.MarketDataSupported())
	return w.Finish(), nil
}
