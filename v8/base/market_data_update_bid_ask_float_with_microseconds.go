// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-31 22:08:18.145964 +0800 WITA m=+0.011497918

package v8

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"io"
)

const MarketDataUpdateBidAskFloatWithMicrosecondsSize = 32

//     Size         uint16                       = MarketDataUpdateBidAskFloatWithMicrosecondsSize  (32)
//     Type         uint16                       = MARKET_DATA_UPDATE_BID_ASK_FLOAT_WITH_MICROSECONDS  (144)
//     SymbolID     uint32                       = 0
//     BidPrice     float32                      = math.MaxFloat32
//     BidQuantity  float32                      = 0.000000
//     AskPrice     float32                      = math.MaxFloat32
//     AskQuantity  float32                      = 0.000000
//     DateTime     DateTimeWithMicrosecondsInt  = 0
var _MarketDataUpdateBidAskFloatWithMicrosecondsDefault = []byte{32, 0, 144, 0, 0, 0, 0, 0, 255, 255, 127, 127, 0, 0, 0, 0, 255, 255, 127, 127, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// MarketDataUpdateBidAskFloatWithMicroseconds This message is optional.
//
// Sent by the Server to the Client when there is an update to the Bid Ask
// prices and/or quantities. This message is identical to the MarketDataUpdateBidAsk
// message except it does not have a timestamp. It needs to be sent when
// there is no change with the timestamp for the Bid Ask update as compared
// to the prior update.
//
// When the Server sends this message to the Client, the Client needs to
// use the prior received Bid Ask update timestamp to know what the timestamp
// is for this message.
type MarketDataUpdateBidAskFloatWithMicroseconds struct {
	p message.Fixed
}

func NewMarketDataUpdateBidAskFloatWithMicrosecondsFrom(b []byte) MarketDataUpdateBidAskFloatWithMicroseconds {
	return MarketDataUpdateBidAskFloatWithMicroseconds{p: message.NewFixed(b)}
}

func WrapMarketDataUpdateBidAskFloatWithMicroseconds(b []byte) MarketDataUpdateBidAskFloatWithMicroseconds {
	return MarketDataUpdateBidAskFloatWithMicroseconds{p: message.WrapFixed(b)}
}

func NewMarketDataUpdateBidAskFloatWithMicroseconds() *MarketDataUpdateBidAskFloatWithMicroseconds {
	return &MarketDataUpdateBidAskFloatWithMicroseconds{p: message.NewFixed(_MarketDataUpdateBidAskFloatWithMicrosecondsDefault)}
}

func ParseMarketDataUpdateBidAskFloatWithMicroseconds(b []byte) (MarketDataUpdateBidAskFloatWithMicroseconds, error) {
	if len(b) < 4 {
		return MarketDataUpdateBidAskFloatWithMicroseconds{}, message.ErrShortBuffer
	}
	m := WrapMarketDataUpdateBidAskFloatWithMicroseconds(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return MarketDataUpdateBidAskFloatWithMicroseconds{}, message.ErrOverflow
	}
	size := int(m.p.AsUint16LE())
	if size > len(b) {
		return MarketDataUpdateBidAskFloatWithMicroseconds{}, message.ErrBaseSizeOverflow
	}
	if size < 32 {
		clone := *NewMarketDataUpdateBidAskFloatWithMicroseconds()
		clone.p.SetBytes(0, b[0:size])
		clone.p.SetBytes(size, _MarketDataUpdateBidAskFloatWithMicrosecondsDefault[size:])
		return clone, nil
	}
	return m, nil
}

// Size The standard message size field. Automatically set by constructor.
func (m MarketDataUpdateBidAskFloatWithMicroseconds) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type The standard message type field. Automatically set by constructor.
func (m MarketDataUpdateBidAskFloatWithMicroseconds) Type() uint16 {
	return m.p.Uint16LE(2)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateBidAskFloatWithMicroseconds) SymbolID() uint32 {
	return m.p.Uint32LE(4)
}

// BidPrice The current Bid price. Leave unset if there is no price available.
func (m MarketDataUpdateBidAskFloatWithMicroseconds) BidPrice() float32 {
	return m.p.Float32LE(8)
}

// BidQuantity The current number of contracts/shares at the bid price.
func (m MarketDataUpdateBidAskFloatWithMicroseconds) BidQuantity() float32 {
	return m.p.Float32LE(12)
}

// AskPrice The current ask or offer price. Leave unset if there is no price available.
// The current ask or offer price. Leave unset if there is no price available.
func (m MarketDataUpdateBidAskFloatWithMicroseconds) AskPrice() float32 {
	return m.p.Float32LE(16)
}

// AskQuantity The current number of contracts/shares at the ask price.
func (m MarketDataUpdateBidAskFloatWithMicroseconds) AskQuantity() float32 {
	return m.p.Float32LE(20)
}

// DateTime The timestamp of the trade in UNIX microseconds time format.
func (m MarketDataUpdateBidAskFloatWithMicroseconds) DateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(m.p.Int64LE(24))
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m *MarketDataUpdateBidAskFloatWithMicroseconds) SetSymbolID(value uint32) *MarketDataUpdateBidAskFloatWithMicroseconds {
	m.p.SetUint32LE(4, value)
	return m
}

// SetBidPrice The current Bid price. Leave unset if there is no price available.
func (m *MarketDataUpdateBidAskFloatWithMicroseconds) SetBidPrice(value float32) *MarketDataUpdateBidAskFloatWithMicroseconds {
	m.p.SetFloat32LE(8, value)
	return m
}

// SetBidQuantity The current number of contracts/shares at the bid price.
func (m *MarketDataUpdateBidAskFloatWithMicroseconds) SetBidQuantity(value float32) *MarketDataUpdateBidAskFloatWithMicroseconds {
	m.p.SetFloat32LE(12, value)
	return m
}

// SetAskPrice The current ask or offer price. Leave unset if there is no price available.
// The current ask or offer price. Leave unset if there is no price available.
func (m *MarketDataUpdateBidAskFloatWithMicroseconds) SetAskPrice(value float32) *MarketDataUpdateBidAskFloatWithMicroseconds {
	m.p.SetFloat32LE(16, value)
	return m
}

// SetAskQuantity The current number of contracts/shares at the ask price.
func (m *MarketDataUpdateBidAskFloatWithMicroseconds) SetAskQuantity(value float32) *MarketDataUpdateBidAskFloatWithMicroseconds {
	m.p.SetFloat32LE(20, value)
	return m
}

// SetDateTime The timestamp of the trade in UNIX microseconds time format.
func (m *MarketDataUpdateBidAskFloatWithMicroseconds) SetDateTime(value DateTimeWithMicrosecondsInt) *MarketDataUpdateBidAskFloatWithMicroseconds {
	m.p.SetInt64LE(24, int64(value))
	return m
}

func (m *MarketDataUpdateBidAskFloatWithMicroseconds) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m *MarketDataUpdateBidAskFloatWithMicroseconds) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

func (m *MarketDataUpdateBidAskFloatWithMicroseconds) Clone() *MarketDataUpdateBidAskFloatWithMicroseconds {
	return &MarketDataUpdateBidAskFloatWithMicroseconds{message.WrapFixedPointer(m.p.Clone(uintptr(m.Size())))}
}

// Copy
func (m MarketDataUpdateBidAskFloatWithMicroseconds) Copy(to MarketDataUpdateBidAskFloatWithMicroseconds) {
	to.SetSymbolID(m.SymbolID())
	to.SetBidPrice(m.BidPrice())
	to.SetBidQuantity(m.BidQuantity())
	to.SetAskPrice(m.AskPrice())
	to.SetAskQuantity(m.AskQuantity())
	to.SetDateTime(m.DateTime())
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateBidAskFloatWithMicroseconds) MarshalJSON() ([]byte, error) {
	return m.MarshalJSONTo(nil)
}

func (m *MarketDataUpdateBidAskFloatWithMicroseconds) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 144)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Float32Field("BidPrice", m.BidPrice())
	w.Float32Field("BidQuantity", m.BidQuantity())
	w.Float32Field("AskPrice", m.AskPrice())
	w.Float32Field("AskQuantity", m.AskQuantity())
	w.Int64Field("DateTime", int64(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateBidAskFloatWithMicroseconds) UnmarshalJSON(b []byte) error {
	r, err := json.OpenReader(b)
	if err != nil {
		return err
	}
	return m.UnmarshalJSONFromReader(&r)
}

func (m *MarketDataUpdateBidAskFloatWithMicroseconds) UnmarshalJSONFromReader(r *json.Reader) error {
	if r.Type != 144 {
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
		case "BidPrice":
			m.SetBidPrice(r.Float32())
		case "BidQuantity":
			m.SetBidQuantity(r.Float32())
		case "AskPrice":
			m.SetAskPrice(r.Float32())
		case "AskQuantity":
			m.SetAskQuantity(r.Float32())
		case "DateTime":
			m.SetDateTime(DateTimeWithMicrosecondsInt(r.Int64()))
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