// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-31 21:46:24.141427 +0800 WITA m=+0.018610501

package v8

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"io"
)

const MarketDataUpdateBidAskSize = 40

//     Size         uint16         = MarketDataUpdateBidAskSize  (40)
//     Type         uint16         = MARKET_DATA_UPDATE_BID_ASK  (108)
//     SymbolID     uint32         = 0
//     BidPrice     float64        = math.MaxFloat64
//     BidQuantity  float32        = 0.000000
//     AskPrice     float64        = math.MaxFloat64
//     AskQuantity  float32        = 0.000000
//     DateTime     DateTime4Byte  = 0
var _MarketDataUpdateBidAskDefault = []byte{40, 0, 108, 0, 0, 0, 0, 0, 255, 255, 255, 255, 255, 255, 239, 127, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 255, 255, 255, 255, 239, 127, 0, 0, 0, 0, 0, 0, 0, 0}

// MarketDataUpdateBidAsk The Server sends this market data feed message to the Client when the
// best bid or ask price or size changes.
type MarketDataUpdateBidAsk struct {
	p message.Fixed
}

func NewMarketDataUpdateBidAskFrom(b []byte) MarketDataUpdateBidAsk {
	return MarketDataUpdateBidAsk{p: message.NewFixed(b)}
}

func WrapMarketDataUpdateBidAsk(b []byte) MarketDataUpdateBidAsk {
	return MarketDataUpdateBidAsk{p: message.WrapFixed(b)}
}

func NewMarketDataUpdateBidAsk() *MarketDataUpdateBidAsk {
	return &MarketDataUpdateBidAsk{p: message.NewFixed(_MarketDataUpdateBidAskDefault)}
}

func ParseMarketDataUpdateBidAsk(b []byte) (MarketDataUpdateBidAsk, error) {
	if len(b) < 4 {
		return MarketDataUpdateBidAsk{}, message.ErrShortBuffer
	}
	m := WrapMarketDataUpdateBidAsk(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return MarketDataUpdateBidAsk{}, message.ErrOverflow
	}
	size := int(m.p.AsUint16LE())
	if size > len(b) {
		return MarketDataUpdateBidAsk{}, message.ErrBaseSizeOverflow
	}
	if size < 40 {
		clone := *NewMarketDataUpdateBidAsk()
		clone.p.SetBytes(0, b[0:size])
		clone.p.SetBytes(size, _MarketDataUpdateBidAskDefault[size:])
		return clone, nil
	}
	return m, nil
}

// Size The standard message size field. Automatically set by constructor.
func (m MarketDataUpdateBidAsk) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type The standard message type field. Automatically set by constructor.
func (m MarketDataUpdateBidAsk) Type() uint16 {
	return m.p.Uint16LE(2)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateBidAsk) SymbolID() uint32 {
	return m.p.Uint32LE(4)
}

// BidPrice The current Bid price. Leave unset if there is no price available.
func (m MarketDataUpdateBidAsk) BidPrice() float64 {
	return m.p.Float64LE(8)
}

// BidQuantity The current number of contracts/shares at the bid price.
func (m MarketDataUpdateBidAsk) BidQuantity() float32 {
	return m.p.Float32LE(16)
}

// AskPrice The current ask or offer price. Leave unset if there is no price available.
// The current ask or offer price. Leave unset if there is no price available.
func (m MarketDataUpdateBidAsk) AskPrice() float64 {
	return m.p.Float64LE(24)
}

// AskQuantity The current number of contracts/shares at the ask price.
func (m MarketDataUpdateBidAsk) AskQuantity() float32 {
	return m.p.Float32LE(32)
}

// DateTime The Date-Time of the Bid and Ask update.
func (m MarketDataUpdateBidAsk) DateTime() DateTime4Byte {
	return DateTime4Byte(m.p.Uint32LE(36))
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m *MarketDataUpdateBidAsk) SetSymbolID(value uint32) *MarketDataUpdateBidAsk {
	m.p.SetUint32LE(4, value)
	return m
}

// SetBidPrice The current Bid price. Leave unset if there is no price available.
func (m *MarketDataUpdateBidAsk) SetBidPrice(value float64) *MarketDataUpdateBidAsk {
	m.p.SetFloat64LE(8, value)
	return m
}

// SetBidQuantity The current number of contracts/shares at the bid price.
func (m *MarketDataUpdateBidAsk) SetBidQuantity(value float32) *MarketDataUpdateBidAsk {
	m.p.SetFloat32LE(16, value)
	return m
}

// SetAskPrice The current ask or offer price. Leave unset if there is no price available.
// The current ask or offer price. Leave unset if there is no price available.
func (m *MarketDataUpdateBidAsk) SetAskPrice(value float64) *MarketDataUpdateBidAsk {
	m.p.SetFloat64LE(24, value)
	return m
}

// SetAskQuantity The current number of contracts/shares at the ask price.
func (m *MarketDataUpdateBidAsk) SetAskQuantity(value float32) *MarketDataUpdateBidAsk {
	m.p.SetFloat32LE(32, value)
	return m
}

// SetDateTime The Date-Time of the Bid and Ask update.
func (m *MarketDataUpdateBidAsk) SetDateTime(value DateTime4Byte) *MarketDataUpdateBidAsk {
	m.p.SetUint32LE(36, uint32(value))
	return m
}

func (m MarketDataUpdateBidAsk) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m MarketDataUpdateBidAsk) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

// Copy
func (m MarketDataUpdateBidAsk) Copy(to MarketDataUpdateBidAsk) {
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

func (m MarketDataUpdateBidAsk) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 108)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Float64Field("BidPrice", m.BidPrice())
	w.Float32Field("BidQuantity", m.BidQuantity())
	w.Float64Field("AskPrice", m.AskPrice())
	w.Float32Field("AskQuantity", m.AskQuantity())
	w.Uint32Field("DateTime", uint32(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateBidAsk) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 108 {
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
			m.SetBidPrice(r.Float64())
		case "BidQuantity":
			m.SetBidQuantity(r.Float32())
		case "AskPrice":
			m.SetAskPrice(r.Float64())
		case "AskQuantity":
			m.SetAskQuantity(r.Float32())
		case "DateTime":
			m.SetDateTime(DateTime4Byte(r.Uint32()))
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

func (m *MarketDataUpdateBidAsk) UnmarshalJSONReader(r *json.Reader) error {
	return m.UnmarshalJSONDoc(r)
}
