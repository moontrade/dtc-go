// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-31 21:46:24.141427 +0800 WITA m=+0.018610501

package v8

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"io"
)

const MarketDepthSnapshotLevelFloatSize = 24

//     Size                uint16                  = MarketDepthSnapshotLevelFloatSize  (24)
//     Type                uint16                  = MARKET_DEPTH_SNAPSHOT_LEVEL_FLOAT  (145)
//     SymbolID            uint32                  = 0
//     Price               float32                 = 0.000000
//     Quantity            float32                 = 0.000000
//     NumOrders           uint32                  = 0
//     Level               uint16                  = 0
//     Side                AtBidOrAskEnum8         = BID_ASK_UNSET_8  (0)
//     FinalUpdateInBatch  FinalUpdateInBatchEnum  = FINAL_UPDATE_UNSET  (0)
var _MarketDepthSnapshotLevelFloatDefault = []byte{24, 0, 145, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// MarketDepthSnapshotLevelFloat This is a message sent by Server to provide the initial market depth data
// entries to the Client after the Client subscribes to market data or separately
// subscribes to market depth data. The Client will need to separately subscribe
// to market depth data if the Server requires it.
//
// Each message provides a single entry of depth data. Therefore, the Server
// will send multiple MarketDepthSnapshotLevelFloat messages in a series
// in order for the Client to build up its initial market depth book.
//
// The first message will be identified by the IsFirstMessageInBatch field
// being set to 1. The last message will be identified by the IsLastMessageInBatch
// field being set to 1.
//
// In the case where the market depth book is empty, the Server still needs
// to send through one single message with the SymbolID set, IsFirstMessageInBatch
// equal to 1 and IsLastMessageInBatch equal to 1. All other members will
// be at the default values. The Client will understand this as an empty
// book.
type MarketDepthSnapshotLevelFloat struct {
	p message.Fixed
}

func NewMarketDepthSnapshotLevelFloatFrom(b []byte) MarketDepthSnapshotLevelFloat {
	return MarketDepthSnapshotLevelFloat{p: message.NewFixed(b)}
}

func WrapMarketDepthSnapshotLevelFloat(b []byte) MarketDepthSnapshotLevelFloat {
	return MarketDepthSnapshotLevelFloat{p: message.WrapFixed(b)}
}

func NewMarketDepthSnapshotLevelFloat() *MarketDepthSnapshotLevelFloat {
	return &MarketDepthSnapshotLevelFloat{p: message.NewFixed(_MarketDepthSnapshotLevelFloatDefault)}
}

func ParseMarketDepthSnapshotLevelFloat(b []byte) (MarketDepthSnapshotLevelFloat, error) {
	if len(b) < 4 {
		return MarketDepthSnapshotLevelFloat{}, message.ErrShortBuffer
	}
	m := WrapMarketDepthSnapshotLevelFloat(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return MarketDepthSnapshotLevelFloat{}, message.ErrOverflow
	}
	size := int(m.p.AsUint16LE())
	if size > len(b) {
		return MarketDepthSnapshotLevelFloat{}, message.ErrBaseSizeOverflow
	}
	if size < 24 {
		clone := *NewMarketDepthSnapshotLevelFloat()
		clone.p.SetBytes(0, b[0:size])
		clone.p.SetBytes(size, _MarketDepthSnapshotLevelFloatDefault[size:])
		return clone, nil
	}
	return m, nil
}

// Size The standard message size field. Automatically set by constructor.
func (m MarketDepthSnapshotLevelFloat) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type The standard message type field. Automatically set by constructor.
func (m MarketDepthSnapshotLevelFloat) Type() uint16 {
	return m.p.Uint16LE(2)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest/MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthSnapshotLevelFloat) SymbolID() uint32 {
	return m.p.Uint32LE(4)
}

// Price This is the price of the market depth entry.
func (m MarketDepthSnapshotLevelFloat) Price() float32 {
	return m.p.Float32LE(8)
}

// Quantity This is the quantity of orders at the Price.
func (m MarketDepthSnapshotLevelFloat) Quantity() float32 {
	return m.p.Float32LE(12)
}

// NumOrders The number of orders at the Price.
func (m MarketDepthSnapshotLevelFloat) NumOrders() uint32 {
	return m.p.Uint32LE(16)
}

// Level This indicates the level of the price within the market depth book. The
// minimum value is 1. There is no maximum value. A value of 1 is considered
// the best bid or ask data.
func (m MarketDepthSnapshotLevelFloat) Level() uint16 {
	return m.p.Uint16LE(20)
}

// Side Set to AT_BID = 1 if this is a bid side market depth entry. Set to AT_ASK
// = 2, if this is an ask side market depth entry.
func (m MarketDepthSnapshotLevelFloat) Side() AtBidOrAskEnum8 {
	return AtBidOrAskEnum8(m.p.Uint8(22))
}

// FinalUpdateInBatch An indicator whether this is the final message or not in a batch of updates.
// An indicator whether this is the final message or not in a batch of updates.
func (m MarketDepthSnapshotLevelFloat) FinalUpdateInBatch() FinalUpdateInBatchEnum {
	return FinalUpdateInBatchEnum(m.p.Uint8(23))
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest/MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m *MarketDepthSnapshotLevelFloat) SetSymbolID(value uint32) *MarketDepthSnapshotLevelFloat {
	m.p.SetUint32LE(4, value)
	return m
}

// SetPrice This is the price of the market depth entry.
func (m *MarketDepthSnapshotLevelFloat) SetPrice(value float32) *MarketDepthSnapshotLevelFloat {
	m.p.SetFloat32LE(8, value)
	return m
}

// SetQuantity This is the quantity of orders at the Price.
func (m *MarketDepthSnapshotLevelFloat) SetQuantity(value float32) *MarketDepthSnapshotLevelFloat {
	m.p.SetFloat32LE(12, value)
	return m
}

// SetNumOrders The number of orders at the Price.
func (m *MarketDepthSnapshotLevelFloat) SetNumOrders(value uint32) *MarketDepthSnapshotLevelFloat {
	m.p.SetUint32LE(16, value)
	return m
}

// SetLevel This indicates the level of the price within the market depth book. The
// minimum value is 1. There is no maximum value. A value of 1 is considered
// the best bid or ask data.
func (m *MarketDepthSnapshotLevelFloat) SetLevel(value uint16) *MarketDepthSnapshotLevelFloat {
	m.p.SetUint16LE(20, value)
	return m
}

// SetSide Set to AT_BID = 1 if this is a bid side market depth entry. Set to AT_ASK
// = 2, if this is an ask side market depth entry.
func (m *MarketDepthSnapshotLevelFloat) SetSide(value AtBidOrAskEnum8) *MarketDepthSnapshotLevelFloat {
	m.p.SetUint8(22, uint8(value))
	return m
}

// SetFinalUpdateInBatch An indicator whether this is the final message or not in a batch of updates.
// An indicator whether this is the final message or not in a batch of updates.
func (m *MarketDepthSnapshotLevelFloat) SetFinalUpdateInBatch(value FinalUpdateInBatchEnum) *MarketDepthSnapshotLevelFloat {
	m.p.SetUint8(23, uint8(value))
	return m
}

func (m MarketDepthSnapshotLevelFloat) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m MarketDepthSnapshotLevelFloat) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

// Copy
func (m MarketDepthSnapshotLevelFloat) Copy(to MarketDepthSnapshotLevelFloat) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
	to.SetNumOrders(m.NumOrders())
	to.SetLevel(m.Level())
	to.SetSide(m.Side())
	to.SetFinalUpdateInBatch(m.FinalUpdateInBatch())
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDepthSnapshotLevelFloat) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 145)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Float32Field("Price", m.Price())
	w.Float32Field("Quantity", m.Quantity())
	w.Uint32Field("NumOrders", m.NumOrders())
	w.Uint16Field("Level", m.Level())
	w.Uint8Field("Side", uint8(m.Side()))
	w.Uint8Field("FinalUpdateInBatch", uint8(m.FinalUpdateInBatch()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDepthSnapshotLevelFloat) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 145 {
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
		case "Quantity":
			m.SetQuantity(r.Float32())
		case "NumOrders":
			m.SetNumOrders(r.Uint32())
		case "Level":
			m.SetLevel(r.Uint16())
		case "Side":
			m.SetSide(AtBidOrAskEnum8(r.Uint8()))
		case "FinalUpdateInBatch":
			m.SetFinalUpdateInBatch(FinalUpdateInBatchEnum(r.Uint8()))
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

func (m *MarketDepthSnapshotLevelFloat) UnmarshalJSONReader(r *json.Reader) error {
	return m.UnmarshalJSONDoc(r)
}
