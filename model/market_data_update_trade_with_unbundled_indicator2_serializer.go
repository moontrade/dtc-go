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

func (m MarketDataUpdateTradeWithUnbundledIndicator2) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":146,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Float32(m.Price())
	w.Data = append(w.Data, ',')
	w.Uint32(m.Volume())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Uint8(uint8(m.AtBidOrAsk()))
	w.Data = append(w.Data, ',')
	w.Int8(int8(m.UnbundledTradeIndicator()))
	w.Data = append(w.Data, ',')
	w.Int8(int8(m.TradeCondition()))
	return w.Finish(), nil
}

func (m MarketDataUpdateTradeWithUnbundledIndicator2Builder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":146,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Float32(m.Price())
	w.Data = append(w.Data, ',')
	w.Uint32(m.Volume())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Uint8(uint8(m.AtBidOrAsk()))
	w.Data = append(w.Data, ',')
	w.Int8(int8(m.UnbundledTradeIndicator()))
	w.Data = append(w.Data, ',')
	w.Int8(int8(m.TradeCondition()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateTradeWithUnbundledIndicator2Pointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":146,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Float32(m.Price())
	w.Data = append(w.Data, ',')
	w.Uint32(m.Volume())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Uint8(uint8(m.AtBidOrAsk()))
	w.Data = append(w.Data, ',')
	w.Int8(int8(m.UnbundledTradeIndicator()))
	w.Data = append(w.Data, ',')
	w.Int8(int8(m.TradeCondition()))
	return w.Finish(), nil
}

func (m MarketDataUpdateTradeWithUnbundledIndicator2PointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":146,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Float32(m.Price())
	w.Data = append(w.Data, ',')
	w.Uint32(m.Volume())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Uint8(uint8(m.AtBidOrAsk()))
	w.Data = append(w.Data, ',')
	w.Int8(int8(m.UnbundledTradeIndicator()))
	w.Data = append(w.Data, ',')
	w.Int8(int8(m.TradeCondition()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateTradeWithUnbundledIndicator2) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 146)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Float32Field("Price", m.Price())
	w.Uint32Field("Volume", m.Volume())
	w.Int64Field("DateTime", int64(m.DateTime()))
	w.Uint8Field("AtBidOrAsk", uint8(m.AtBidOrAsk()))
	w.Int8Field("UnbundledTradeIndicator", int8(m.UnbundledTradeIndicator()))
	w.Int8Field("TradeCondition", int8(m.TradeCondition()))
	return w.Finish(), nil
}

func (m MarketDataUpdateTradeWithUnbundledIndicator2Builder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 146)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Float32Field("Price", m.Price())
	w.Uint32Field("Volume", m.Volume())
	w.Int64Field("DateTime", int64(m.DateTime()))
	w.Uint8Field("AtBidOrAsk", uint8(m.AtBidOrAsk()))
	w.Int8Field("UnbundledTradeIndicator", int8(m.UnbundledTradeIndicator()))
	w.Int8Field("TradeCondition", int8(m.TradeCondition()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateTradeWithUnbundledIndicator2Pointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 146)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Float32Field("Price", m.Price())
	w.Uint32Field("Volume", m.Volume())
	w.Int64Field("DateTime", int64(m.DateTime()))
	w.Uint8Field("AtBidOrAsk", uint8(m.AtBidOrAsk()))
	w.Int8Field("UnbundledTradeIndicator", int8(m.UnbundledTradeIndicator()))
	w.Int8Field("TradeCondition", int8(m.TradeCondition()))
	return w.Finish(), nil
}

func (m MarketDataUpdateTradeWithUnbundledIndicator2PointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 146)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Float32Field("Price", m.Price())
	w.Uint32Field("Volume", m.Volume())
	w.Int64Field("DateTime", int64(m.DateTime()))
	w.Uint8Field("AtBidOrAsk", uint8(m.AtBidOrAsk()))
	w.Int8Field("UnbundledTradeIndicator", int8(m.UnbundledTradeIndicator()))
	w.Int8Field("TradeCondition", int8(m.TradeCondition()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateTradeWithUnbundledIndicator2Builder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 146 {
		return message.ErrWrongType
	}
	m.SetSymbolID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetVolume(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetDateTime(DateTimeWithMicrosecondsInt(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetAtBidOrAsk(AtBidOrAskEnum8(r.Uint8()))
	if r.IsError() {
		return r.Error()
	}
	m.SetUnbundledTradeIndicator(UnbundledTradeIndicatorEnum(r.Int8()))
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeCondition(TradeConditionEnum(r.Int8()))
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateTradeWithUnbundledIndicator2PointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 146 {
		return message.ErrWrongType
	}
	m.SetSymbolID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetVolume(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetDateTime(DateTimeWithMicrosecondsInt(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetAtBidOrAsk(AtBidOrAskEnum8(r.Uint8()))
	if r.IsError() {
		return r.Error()
	}
	m.SetUnbundledTradeIndicator(UnbundledTradeIndicatorEnum(r.Int8()))
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeCondition(TradeConditionEnum(r.Int8()))
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateTradeWithUnbundledIndicator2Builder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 146 {
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
		case "Price":
			m.SetPrice(r.Float32())
		case "Volume":
			m.SetVolume(r.Uint32())
		case "DateTime":
			m.SetDateTime(DateTimeWithMicrosecondsInt(r.Int64()))
		case "AtBidOrAsk":
			m.SetAtBidOrAsk(AtBidOrAskEnum8(r.Uint8()))
		case "UnbundledTradeIndicator":
			m.SetUnbundledTradeIndicator(UnbundledTradeIndicatorEnum(r.Int8()))
		case "TradeCondition":
			m.SetTradeCondition(TradeConditionEnum(r.Int8()))
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

func (m *MarketDataUpdateTradeWithUnbundledIndicator2PointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 146 {
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
		case "Price":
			m.SetPrice(r.Float32())
		case "Volume":
			m.SetVolume(r.Uint32())
		case "DateTime":
			m.SetDateTime(DateTimeWithMicrosecondsInt(r.Int64()))
		case "AtBidOrAsk":
			m.SetAtBidOrAsk(AtBidOrAskEnum8(r.Uint8()))
		case "UnbundledTradeIndicator":
			m.SetUnbundledTradeIndicator(UnbundledTradeIndicatorEnum(r.Int8()))
		case "TradeCondition":
			m.SetTradeCondition(TradeConditionEnum(r.Int8()))
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

func (m *MarketDataUpdateTradeWithUnbundledIndicator2Builder) UnmarshalProtobuf(b []byte) error {
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
			m.SetPrice(r.ReadFloat32())
		case 3:
			m.SetVolume(r.ReadUint32())
		case 4:
			m.SetDateTime(DateTimeWithMicrosecondsInt(r.ReadInt64()))
		case 5:
			m.SetAtBidOrAsk(AtBidOrAskEnum8(r.ReadUint8()))
		case 6:
			m.SetUnbundledTradeIndicator(UnbundledTradeIndicatorEnum(r.ReadInt8()))
		case 7:
			m.SetTradeCondition(TradeConditionEnum(r.ReadInt8()))
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

func (m *MarketDataUpdateTradeWithUnbundledIndicator2PointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetPrice(r.ReadFloat32())
		case 3:
			m.SetVolume(r.ReadUint32())
		case 4:
			m.SetDateTime(DateTimeWithMicrosecondsInt(r.ReadInt64()))
		case 5:
			m.SetAtBidOrAsk(AtBidOrAskEnum8(r.ReadUint8()))
		case 6:
			m.SetUnbundledTradeIndicator(UnbundledTradeIndicatorEnum(r.ReadInt8()))
		case 7:
			m.SetTradeCondition(TradeConditionEnum(r.ReadInt8()))
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

func (m MarketDataUpdateTradeWithUnbundledIndicator2Builder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 146)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteFixed32Float32(2, m.Price())
	w.WriteUvarint32(3, m.Volume())
	w.WriteFixed64Int64(4, int64(m.DateTime()))
	w.WriteUvarint8(5, uint8(m.AtBidOrAsk()))
	w.WriteVarint8(6, int8(m.UnbundledTradeIndicator()))
	w.WriteVarint8(7, int8(m.TradeCondition()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateTradeWithUnbundledIndicator2PointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 146)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteFixed32Float32(2, m.Price())
	w.WriteUvarint32(3, m.Volume())
	w.WriteFixed64Int64(4, int64(m.DateTime()))
	w.WriteUvarint8(5, uint8(m.AtBidOrAsk()))
	w.WriteVarint8(6, int8(m.UnbundledTradeIndicator()))
	w.WriteVarint8(7, int8(m.TradeCondition()))
	return w.Finish(), nil
}