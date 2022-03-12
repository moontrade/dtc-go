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

func (m MarketDepthUpdateLevel) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":106,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Uint16(uint16(m.Side()))
	w.Data = append(w.Data, ',')
	w.Float64(m.Price())
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity())
	w.Data = append(w.Data, ',')
	w.Uint8(uint8(m.UpdateType()))
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.NumOrders())
	return w.Finish(), nil
}

func (m MarketDepthUpdateLevelBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":106,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Uint16(uint16(m.Side()))
	w.Data = append(w.Data, ',')
	w.Float64(m.Price())
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity())
	w.Data = append(w.Data, ',')
	w.Uint8(uint8(m.UpdateType()))
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.NumOrders())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDepthUpdateLevelPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":106,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Uint16(uint16(m.Side()))
	w.Data = append(w.Data, ',')
	w.Float64(m.Price())
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity())
	w.Data = append(w.Data, ',')
	w.Uint8(uint8(m.UpdateType()))
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.NumOrders())
	return w.Finish(), nil
}

func (m MarketDepthUpdateLevelPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":106,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Uint16(uint16(m.Side()))
	w.Data = append(w.Data, ',')
	w.Float64(m.Price())
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity())
	w.Data = append(w.Data, ',')
	w.Uint8(uint8(m.UpdateType()))
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.NumOrders())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDepthUpdateLevel) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 106)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Uint16Field("Side", uint16(m.Side()))
	w.Float64Field("Price", m.Price())
	w.Float64Field("Quantity", m.Quantity())
	w.Uint8Field("UpdateType", uint8(m.UpdateType()))
	w.Float64Field("DateTime", float64(m.DateTime()))
	w.Uint32Field("NumOrders", m.NumOrders())
	return w.Finish(), nil
}

func (m MarketDepthUpdateLevelBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 106)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Uint16Field("Side", uint16(m.Side()))
	w.Float64Field("Price", m.Price())
	w.Float64Field("Quantity", m.Quantity())
	w.Uint8Field("UpdateType", uint8(m.UpdateType()))
	w.Float64Field("DateTime", float64(m.DateTime()))
	w.Uint32Field("NumOrders", m.NumOrders())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDepthUpdateLevelPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 106)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Uint16Field("Side", uint16(m.Side()))
	w.Float64Field("Price", m.Price())
	w.Float64Field("Quantity", m.Quantity())
	w.Uint8Field("UpdateType", uint8(m.UpdateType()))
	w.Float64Field("DateTime", float64(m.DateTime()))
	w.Uint32Field("NumOrders", m.NumOrders())
	return w.Finish(), nil
}

func (m MarketDepthUpdateLevelPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 106)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Uint16Field("Side", uint16(m.Side()))
	w.Float64Field("Price", m.Price())
	w.Float64Field("Quantity", m.Quantity())
	w.Uint8Field("UpdateType", uint8(m.UpdateType()))
	w.Float64Field("DateTime", float64(m.DateTime()))
	w.Uint32Field("NumOrders", m.NumOrders())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDepthUpdateLevelBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 106 {
		return message.ErrWrongType
	}
	m.SetSymbolID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSide(AtBidOrAskEnum(r.Uint16()))
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetUpdateType(MarketDepthUpdateTypeEnum(r.Uint8()))
	if r.IsError() {
		return r.Error()
	}
	m.SetDateTime(DateTimeWithMilliseconds(r.Float64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetNumOrders(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDepthUpdateLevelPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 106 {
		return message.ErrWrongType
	}
	m.SetSymbolID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSide(AtBidOrAskEnum(r.Uint16()))
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetUpdateType(MarketDepthUpdateTypeEnum(r.Uint8()))
	if r.IsError() {
		return r.Error()
	}
	m.SetDateTime(DateTimeWithMilliseconds(r.Float64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetNumOrders(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDepthUpdateLevelBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 106 {
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
		case "Side":
			m.SetSide(AtBidOrAskEnum(r.Uint16()))
		case "Price":
			m.SetPrice(r.Float64())
		case "Quantity":
			m.SetQuantity(r.Float64())
		case "UpdateType":
			m.SetUpdateType(MarketDepthUpdateTypeEnum(r.Uint8()))
		case "DateTime":
			m.SetDateTime(DateTimeWithMilliseconds(r.Float64()))
		case "NumOrders":
			m.SetNumOrders(r.Uint32())
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

func (m *MarketDepthUpdateLevelPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 106 {
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
		case "Side":
			m.SetSide(AtBidOrAskEnum(r.Uint16()))
		case "Price":
			m.SetPrice(r.Float64())
		case "Quantity":
			m.SetQuantity(r.Float64())
		case "UpdateType":
			m.SetUpdateType(MarketDepthUpdateTypeEnum(r.Uint8()))
		case "DateTime":
			m.SetDateTime(DateTimeWithMilliseconds(r.Float64()))
		case "NumOrders":
			m.SetNumOrders(r.Uint32())
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

func (m *MarketDepthUpdateLevelBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetSide(AtBidOrAskEnum(r.ReadUint16()))
		case 3:
			m.SetPrice(r.ReadFloat64())
		case 4:
			m.SetQuantity(r.ReadFloat64())
		case 5:
			m.SetUpdateType(MarketDepthUpdateTypeEnum(r.ReadUint8()))
		case 6:
			m.SetDateTime(DateTimeWithMilliseconds(r.ReadFloat64()))
		case 7:
			m.SetNumOrders(r.ReadUint32())
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

func (m *MarketDepthUpdateLevelPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetSide(AtBidOrAskEnum(r.ReadUint16()))
		case 3:
			m.SetPrice(r.ReadFloat64())
		case 4:
			m.SetQuantity(r.ReadFloat64())
		case 5:
			m.SetUpdateType(MarketDepthUpdateTypeEnum(r.ReadUint8()))
		case 6:
			m.SetDateTime(DateTimeWithMilliseconds(r.ReadFloat64()))
		case 7:
			m.SetNumOrders(r.ReadUint32())
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

func (m MarketDepthUpdateLevelBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 106)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteUvarint16(2, uint16(m.Side()))
	w.WriteFixed64Float64(3, m.Price())
	w.WriteFixed64Float64(4, m.Quantity())
	w.WriteUvarint8(5, uint8(m.UpdateType()))
	w.WriteFixed64Float64(6, float64(m.DateTime()))
	w.WriteUvarint32(7, m.NumOrders())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDepthUpdateLevelPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 106)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteUvarint16(2, uint16(m.Side()))
	w.WriteFixed64Float64(3, m.Price())
	w.WriteFixed64Float64(4, m.Quantity())
	w.WriteUvarint8(5, uint8(m.UpdateType()))
	w.WriteFixed64Float64(6, float64(m.DateTime()))
	w.WriteUvarint32(7, m.NumOrders())
	return w.Finish(), nil
}