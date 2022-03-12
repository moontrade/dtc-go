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

func (m MarketDataSnapshot_Int) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":125,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Int32(m.SessionSettlementPrice())
	w.Data = append(w.Data, ',')
	w.Int32(m.SessionOpenPrice())
	w.Data = append(w.Data, ',')
	w.Int32(m.SessionHighPrice())
	w.Data = append(w.Data, ',')
	w.Int32(m.SessionLowPrice())
	w.Data = append(w.Data, ',')
	w.Int32(m.SessionVolume())
	w.Data = append(w.Data, ',')
	w.Uint32(m.SessionNumTrades())
	w.Data = append(w.Data, ',')
	w.Uint32(m.OpenInterest())
	w.Data = append(w.Data, ',')
	w.Int32(m.BidPrice())
	w.Data = append(w.Data, ',')
	w.Int32(m.AskPrice())
	w.Data = append(w.Data, ',')
	w.Int32(m.AskQuantity())
	w.Data = append(w.Data, ',')
	w.Int32(m.BidQuantity())
	w.Data = append(w.Data, ',')
	w.Int32(m.LastTradePrice())
	w.Data = append(w.Data, ',')
	w.Int32(m.LastTradeVolume())
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.LastTradeDateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.BidAskDateTime()))
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.SessionSettlementDateTime()))
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.TradingSessionDate()))
	w.Data = append(w.Data, ',')
	w.Int8(int8(m.TradingStatus()))
	return w.Finish(), nil
}

func (m MarketDataSnapshot_IntBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":125,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Int32(m.SessionSettlementPrice())
	w.Data = append(w.Data, ',')
	w.Int32(m.SessionOpenPrice())
	w.Data = append(w.Data, ',')
	w.Int32(m.SessionHighPrice())
	w.Data = append(w.Data, ',')
	w.Int32(m.SessionLowPrice())
	w.Data = append(w.Data, ',')
	w.Int32(m.SessionVolume())
	w.Data = append(w.Data, ',')
	w.Uint32(m.SessionNumTrades())
	w.Data = append(w.Data, ',')
	w.Uint32(m.OpenInterest())
	w.Data = append(w.Data, ',')
	w.Int32(m.BidPrice())
	w.Data = append(w.Data, ',')
	w.Int32(m.AskPrice())
	w.Data = append(w.Data, ',')
	w.Int32(m.AskQuantity())
	w.Data = append(w.Data, ',')
	w.Int32(m.BidQuantity())
	w.Data = append(w.Data, ',')
	w.Int32(m.LastTradePrice())
	w.Data = append(w.Data, ',')
	w.Int32(m.LastTradeVolume())
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.LastTradeDateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.BidAskDateTime()))
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.SessionSettlementDateTime()))
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.TradingSessionDate()))
	w.Data = append(w.Data, ',')
	w.Int8(int8(m.TradingStatus()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataSnapshot_IntPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":125,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Int32(m.SessionSettlementPrice())
	w.Data = append(w.Data, ',')
	w.Int32(m.SessionOpenPrice())
	w.Data = append(w.Data, ',')
	w.Int32(m.SessionHighPrice())
	w.Data = append(w.Data, ',')
	w.Int32(m.SessionLowPrice())
	w.Data = append(w.Data, ',')
	w.Int32(m.SessionVolume())
	w.Data = append(w.Data, ',')
	w.Uint32(m.SessionNumTrades())
	w.Data = append(w.Data, ',')
	w.Uint32(m.OpenInterest())
	w.Data = append(w.Data, ',')
	w.Int32(m.BidPrice())
	w.Data = append(w.Data, ',')
	w.Int32(m.AskPrice())
	w.Data = append(w.Data, ',')
	w.Int32(m.AskQuantity())
	w.Data = append(w.Data, ',')
	w.Int32(m.BidQuantity())
	w.Data = append(w.Data, ',')
	w.Int32(m.LastTradePrice())
	w.Data = append(w.Data, ',')
	w.Int32(m.LastTradeVolume())
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.LastTradeDateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.BidAskDateTime()))
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.SessionSettlementDateTime()))
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.TradingSessionDate()))
	w.Data = append(w.Data, ',')
	w.Int8(int8(m.TradingStatus()))
	return w.Finish(), nil
}

func (m MarketDataSnapshot_IntPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":125,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Int32(m.SessionSettlementPrice())
	w.Data = append(w.Data, ',')
	w.Int32(m.SessionOpenPrice())
	w.Data = append(w.Data, ',')
	w.Int32(m.SessionHighPrice())
	w.Data = append(w.Data, ',')
	w.Int32(m.SessionLowPrice())
	w.Data = append(w.Data, ',')
	w.Int32(m.SessionVolume())
	w.Data = append(w.Data, ',')
	w.Uint32(m.SessionNumTrades())
	w.Data = append(w.Data, ',')
	w.Uint32(m.OpenInterest())
	w.Data = append(w.Data, ',')
	w.Int32(m.BidPrice())
	w.Data = append(w.Data, ',')
	w.Int32(m.AskPrice())
	w.Data = append(w.Data, ',')
	w.Int32(m.AskQuantity())
	w.Data = append(w.Data, ',')
	w.Int32(m.BidQuantity())
	w.Data = append(w.Data, ',')
	w.Int32(m.LastTradePrice())
	w.Data = append(w.Data, ',')
	w.Int32(m.LastTradeVolume())
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.LastTradeDateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.BidAskDateTime()))
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.SessionSettlementDateTime()))
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.TradingSessionDate()))
	w.Data = append(w.Data, ',')
	w.Int8(int8(m.TradingStatus()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataSnapshot_Int) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 125)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Int32Field("SessionSettlementPrice", m.SessionSettlementPrice())
	w.Int32Field("SessionOpenPrice", m.SessionOpenPrice())
	w.Int32Field("SessionHighPrice", m.SessionHighPrice())
	w.Int32Field("SessionLowPrice", m.SessionLowPrice())
	w.Int32Field("SessionVolume", m.SessionVolume())
	w.Uint32Field("SessionNumTrades", m.SessionNumTrades())
	w.Uint32Field("OpenInterest", m.OpenInterest())
	w.Int32Field("BidPrice", m.BidPrice())
	w.Int32Field("AskPrice", m.AskPrice())
	w.Int32Field("AskQuantity", m.AskQuantity())
	w.Int32Field("BidQuantity", m.BidQuantity())
	w.Int32Field("LastTradePrice", m.LastTradePrice())
	w.Int32Field("LastTradeVolume", m.LastTradeVolume())
	w.Float64Field("LastTradeDateTime", float64(m.LastTradeDateTime()))
	w.Float64Field("BidAskDateTime", float64(m.BidAskDateTime()))
	w.Uint32Field("SessionSettlementDateTime", uint32(m.SessionSettlementDateTime()))
	w.Uint32Field("TradingSessionDate", uint32(m.TradingSessionDate()))
	w.Int8Field("TradingStatus", int8(m.TradingStatus()))
	return w.Finish(), nil
}

func (m MarketDataSnapshot_IntBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 125)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Int32Field("SessionSettlementPrice", m.SessionSettlementPrice())
	w.Int32Field("SessionOpenPrice", m.SessionOpenPrice())
	w.Int32Field("SessionHighPrice", m.SessionHighPrice())
	w.Int32Field("SessionLowPrice", m.SessionLowPrice())
	w.Int32Field("SessionVolume", m.SessionVolume())
	w.Uint32Field("SessionNumTrades", m.SessionNumTrades())
	w.Uint32Field("OpenInterest", m.OpenInterest())
	w.Int32Field("BidPrice", m.BidPrice())
	w.Int32Field("AskPrice", m.AskPrice())
	w.Int32Field("AskQuantity", m.AskQuantity())
	w.Int32Field("BidQuantity", m.BidQuantity())
	w.Int32Field("LastTradePrice", m.LastTradePrice())
	w.Int32Field("LastTradeVolume", m.LastTradeVolume())
	w.Float64Field("LastTradeDateTime", float64(m.LastTradeDateTime()))
	w.Float64Field("BidAskDateTime", float64(m.BidAskDateTime()))
	w.Uint32Field("SessionSettlementDateTime", uint32(m.SessionSettlementDateTime()))
	w.Uint32Field("TradingSessionDate", uint32(m.TradingSessionDate()))
	w.Int8Field("TradingStatus", int8(m.TradingStatus()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataSnapshot_IntPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 125)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Int32Field("SessionSettlementPrice", m.SessionSettlementPrice())
	w.Int32Field("SessionOpenPrice", m.SessionOpenPrice())
	w.Int32Field("SessionHighPrice", m.SessionHighPrice())
	w.Int32Field("SessionLowPrice", m.SessionLowPrice())
	w.Int32Field("SessionVolume", m.SessionVolume())
	w.Uint32Field("SessionNumTrades", m.SessionNumTrades())
	w.Uint32Field("OpenInterest", m.OpenInterest())
	w.Int32Field("BidPrice", m.BidPrice())
	w.Int32Field("AskPrice", m.AskPrice())
	w.Int32Field("AskQuantity", m.AskQuantity())
	w.Int32Field("BidQuantity", m.BidQuantity())
	w.Int32Field("LastTradePrice", m.LastTradePrice())
	w.Int32Field("LastTradeVolume", m.LastTradeVolume())
	w.Float64Field("LastTradeDateTime", float64(m.LastTradeDateTime()))
	w.Float64Field("BidAskDateTime", float64(m.BidAskDateTime()))
	w.Uint32Field("SessionSettlementDateTime", uint32(m.SessionSettlementDateTime()))
	w.Uint32Field("TradingSessionDate", uint32(m.TradingSessionDate()))
	w.Int8Field("TradingStatus", int8(m.TradingStatus()))
	return w.Finish(), nil
}

func (m MarketDataSnapshot_IntPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 125)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Int32Field("SessionSettlementPrice", m.SessionSettlementPrice())
	w.Int32Field("SessionOpenPrice", m.SessionOpenPrice())
	w.Int32Field("SessionHighPrice", m.SessionHighPrice())
	w.Int32Field("SessionLowPrice", m.SessionLowPrice())
	w.Int32Field("SessionVolume", m.SessionVolume())
	w.Uint32Field("SessionNumTrades", m.SessionNumTrades())
	w.Uint32Field("OpenInterest", m.OpenInterest())
	w.Int32Field("BidPrice", m.BidPrice())
	w.Int32Field("AskPrice", m.AskPrice())
	w.Int32Field("AskQuantity", m.AskQuantity())
	w.Int32Field("BidQuantity", m.BidQuantity())
	w.Int32Field("LastTradePrice", m.LastTradePrice())
	w.Int32Field("LastTradeVolume", m.LastTradeVolume())
	w.Float64Field("LastTradeDateTime", float64(m.LastTradeDateTime()))
	w.Float64Field("BidAskDateTime", float64(m.BidAskDateTime()))
	w.Uint32Field("SessionSettlementDateTime", uint32(m.SessionSettlementDateTime()))
	w.Uint32Field("TradingSessionDate", uint32(m.TradingSessionDate()))
	w.Int8Field("TradingStatus", int8(m.TradingStatus()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataSnapshot_IntBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 125 {
		return message.ErrWrongType
	}
	m.SetSymbolID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSessionSettlementPrice(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSessionOpenPrice(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSessionHighPrice(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSessionLowPrice(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSessionVolume(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSessionNumTrades(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetOpenInterest(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetBidPrice(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetAskPrice(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetAskQuantity(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetBidQuantity(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetLastTradePrice(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetLastTradeVolume(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetLastTradeDateTime(DateTimeWithMilliseconds(r.Float64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetBidAskDateTime(DateTimeWithMilliseconds(r.Float64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetSessionSettlementDateTime(DateTime4Byte(r.Uint32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetTradingSessionDate(DateTime4Byte(r.Uint32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetTradingStatus(TradingStatusEnum(r.Int8()))
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataSnapshot_IntPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 125 {
		return message.ErrWrongType
	}
	m.SetSymbolID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSessionSettlementPrice(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSessionOpenPrice(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSessionHighPrice(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSessionLowPrice(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSessionVolume(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSessionNumTrades(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetOpenInterest(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetBidPrice(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetAskPrice(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetAskQuantity(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetBidQuantity(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetLastTradePrice(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetLastTradeVolume(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetLastTradeDateTime(DateTimeWithMilliseconds(r.Float64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetBidAskDateTime(DateTimeWithMilliseconds(r.Float64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetSessionSettlementDateTime(DateTime4Byte(r.Uint32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetTradingSessionDate(DateTime4Byte(r.Uint32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetTradingStatus(TradingStatusEnum(r.Int8()))
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataSnapshot_IntBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 125 {
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
		case "SessionSettlementPrice":
			m.SetSessionSettlementPrice(r.Int32())
		case "SessionOpenPrice":
			m.SetSessionOpenPrice(r.Int32())
		case "SessionHighPrice":
			m.SetSessionHighPrice(r.Int32())
		case "SessionLowPrice":
			m.SetSessionLowPrice(r.Int32())
		case "SessionVolume":
			m.SetSessionVolume(r.Int32())
		case "SessionNumTrades":
			m.SetSessionNumTrades(r.Uint32())
		case "OpenInterest":
			m.SetOpenInterest(r.Uint32())
		case "BidPrice":
			m.SetBidPrice(r.Int32())
		case "AskPrice":
			m.SetAskPrice(r.Int32())
		case "AskQuantity":
			m.SetAskQuantity(r.Int32())
		case "BidQuantity":
			m.SetBidQuantity(r.Int32())
		case "LastTradePrice":
			m.SetLastTradePrice(r.Int32())
		case "LastTradeVolume":
			m.SetLastTradeVolume(r.Int32())
		case "LastTradeDateTime":
			m.SetLastTradeDateTime(DateTimeWithMilliseconds(r.Float64()))
		case "BidAskDateTime":
			m.SetBidAskDateTime(DateTimeWithMilliseconds(r.Float64()))
		case "SessionSettlementDateTime":
			m.SetSessionSettlementDateTime(DateTime4Byte(r.Uint32()))
		case "TradingSessionDate":
			m.SetTradingSessionDate(DateTime4Byte(r.Uint32()))
		case "TradingStatus":
			m.SetTradingStatus(TradingStatusEnum(r.Int8()))
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

func (m *MarketDataSnapshot_IntPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 125 {
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
		case "SessionSettlementPrice":
			m.SetSessionSettlementPrice(r.Int32())
		case "SessionOpenPrice":
			m.SetSessionOpenPrice(r.Int32())
		case "SessionHighPrice":
			m.SetSessionHighPrice(r.Int32())
		case "SessionLowPrice":
			m.SetSessionLowPrice(r.Int32())
		case "SessionVolume":
			m.SetSessionVolume(r.Int32())
		case "SessionNumTrades":
			m.SetSessionNumTrades(r.Uint32())
		case "OpenInterest":
			m.SetOpenInterest(r.Uint32())
		case "BidPrice":
			m.SetBidPrice(r.Int32())
		case "AskPrice":
			m.SetAskPrice(r.Int32())
		case "AskQuantity":
			m.SetAskQuantity(r.Int32())
		case "BidQuantity":
			m.SetBidQuantity(r.Int32())
		case "LastTradePrice":
			m.SetLastTradePrice(r.Int32())
		case "LastTradeVolume":
			m.SetLastTradeVolume(r.Int32())
		case "LastTradeDateTime":
			m.SetLastTradeDateTime(DateTimeWithMilliseconds(r.Float64()))
		case "BidAskDateTime":
			m.SetBidAskDateTime(DateTimeWithMilliseconds(r.Float64()))
		case "SessionSettlementDateTime":
			m.SetSessionSettlementDateTime(DateTime4Byte(r.Uint32()))
		case "TradingSessionDate":
			m.SetTradingSessionDate(DateTime4Byte(r.Uint32()))
		case "TradingStatus":
			m.SetTradingStatus(TradingStatusEnum(r.Int8()))
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

func (m *MarketDataSnapshot_IntBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetSymbolID(r.ReadUint32())
		case 2:
			m.SetSessionSettlementPrice(r.ReadInt32())
		case 3:
			m.SetSessionOpenPrice(r.ReadInt32())
		case 4:
			m.SetSessionHighPrice(r.ReadInt32())
		case 5:
			m.SetSessionLowPrice(r.ReadInt32())
		case 6:
			m.SetSessionVolume(r.ReadInt32())
		case 7:
			m.SetSessionNumTrades(r.ReadUint32())
		case 8:
			m.SetOpenInterest(r.ReadUint32())
		case 9:
			m.SetBidPrice(r.ReadInt32())
		case 10:
			m.SetAskPrice(r.ReadInt32())
		case 11:
			m.SetAskQuantity(r.ReadInt32())
		case 12:
			m.SetBidQuantity(r.ReadInt32())
		case 13:
			m.SetLastTradePrice(r.ReadInt32())
		case 14:
			m.SetLastTradeVolume(r.ReadInt32())
		case 15:
			m.SetLastTradeDateTime(DateTimeWithMilliseconds(r.ReadFloat64()))
		case 16:
			m.SetBidAskDateTime(DateTimeWithMilliseconds(r.ReadFloat64()))
		case 17:
			m.SetSessionSettlementDateTime(DateTime4Byte(r.ReadUint32()))
		case 18:
			m.SetTradingSessionDate(DateTime4Byte(r.ReadUint32()))
		case 19:
			m.SetTradingStatus(TradingStatusEnum(r.ReadInt8()))
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

func (m *MarketDataSnapshot_IntPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetSymbolID(r.ReadUint32())
		case 2:
			m.SetSessionSettlementPrice(r.ReadInt32())
		case 3:
			m.SetSessionOpenPrice(r.ReadInt32())
		case 4:
			m.SetSessionHighPrice(r.ReadInt32())
		case 5:
			m.SetSessionLowPrice(r.ReadInt32())
		case 6:
			m.SetSessionVolume(r.ReadInt32())
		case 7:
			m.SetSessionNumTrades(r.ReadUint32())
		case 8:
			m.SetOpenInterest(r.ReadUint32())
		case 9:
			m.SetBidPrice(r.ReadInt32())
		case 10:
			m.SetAskPrice(r.ReadInt32())
		case 11:
			m.SetAskQuantity(r.ReadInt32())
		case 12:
			m.SetBidQuantity(r.ReadInt32())
		case 13:
			m.SetLastTradePrice(r.ReadInt32())
		case 14:
			m.SetLastTradeVolume(r.ReadInt32())
		case 15:
			m.SetLastTradeDateTime(DateTimeWithMilliseconds(r.ReadFloat64()))
		case 16:
			m.SetBidAskDateTime(DateTimeWithMilliseconds(r.ReadFloat64()))
		case 17:
			m.SetSessionSettlementDateTime(DateTime4Byte(r.ReadUint32()))
		case 18:
			m.SetTradingSessionDate(DateTime4Byte(r.ReadUint32()))
		case 19:
			m.SetTradingStatus(TradingStatusEnum(r.ReadInt8()))
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

func (m MarketDataSnapshot_IntBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 125)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteVarint32(2, m.SessionSettlementPrice())
	w.WriteVarint32(3, m.SessionOpenPrice())
	w.WriteVarint32(4, m.SessionHighPrice())
	w.WriteVarint32(5, m.SessionLowPrice())
	w.WriteVarint32(6, m.SessionVolume())
	w.WriteUvarint32(7, m.SessionNumTrades())
	w.WriteUvarint32(8, m.OpenInterest())
	w.WriteVarint32(9, m.BidPrice())
	w.WriteVarint32(10, m.AskPrice())
	w.WriteVarint32(11, m.AskQuantity())
	w.WriteVarint32(12, m.BidQuantity())
	w.WriteVarint32(13, m.LastTradePrice())
	w.WriteVarint32(14, m.LastTradeVolume())
	w.WriteFixed64Float64(15, float64(m.LastTradeDateTime()))
	w.WriteFixed64Float64(16, float64(m.BidAskDateTime()))
	w.WriteUvarint32(17, uint32(m.SessionSettlementDateTime()))
	w.WriteUvarint32(18, uint32(m.TradingSessionDate()))
	w.WriteVarint8(19, int8(m.TradingStatus()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataSnapshot_IntPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 125)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteVarint32(2, m.SessionSettlementPrice())
	w.WriteVarint32(3, m.SessionOpenPrice())
	w.WriteVarint32(4, m.SessionHighPrice())
	w.WriteVarint32(5, m.SessionLowPrice())
	w.WriteVarint32(6, m.SessionVolume())
	w.WriteUvarint32(7, m.SessionNumTrades())
	w.WriteUvarint32(8, m.OpenInterest())
	w.WriteVarint32(9, m.BidPrice())
	w.WriteVarint32(10, m.AskPrice())
	w.WriteVarint32(11, m.AskQuantity())
	w.WriteVarint32(12, m.BidQuantity())
	w.WriteVarint32(13, m.LastTradePrice())
	w.WriteVarint32(14, m.LastTradeVolume())
	w.WriteFixed64Float64(15, float64(m.LastTradeDateTime()))
	w.WriteFixed64Float64(16, float64(m.BidAskDateTime()))
	w.WriteUvarint32(17, uint32(m.SessionSettlementDateTime()))
	w.WriteUvarint32(18, uint32(m.TradingSessionDate()))
	w.WriteVarint8(19, int8(m.TradingStatus()))
	return w.Finish(), nil
}