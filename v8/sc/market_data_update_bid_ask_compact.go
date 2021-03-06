// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-31 21:46:24.141427 +0800 WITA m=+0.018610501

package v8

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"io"
)

const MarketDataUpdateBidAskCompactSize = 28

//     Size         uint16         = MarketDataUpdateBidAskCompactSize  (28)
//     Type         uint16         = MARKET_DATA_UPDATE_BID_ASK_COMPACT  (117)
//     BidPrice     float32        = math.MaxFloat32
//     BidQuantity  float32        = 0.000000
//     AskPrice     float32        = math.MaxFloat32
//     AskQuantity  float32        = 0.000000
//     DateTime     DateTime4Byte  = 0
//     SymbolID     uint32         = 0
var _MarketDataUpdateBidAskCompactDefault = []byte{28, 0, 117, 0, 255, 255, 127, 127, 0, 0, 0, 0, 255, 255, 127, 127, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type MarketDataUpdateBidAskCompact struct {
	p message.Fixed
}

func NewMarketDataUpdateBidAskCompactFrom(b []byte) MarketDataUpdateBidAskCompact {
	return MarketDataUpdateBidAskCompact{p: message.NewFixed(b)}
}

func WrapMarketDataUpdateBidAskCompact(b []byte) MarketDataUpdateBidAskCompact {
	return MarketDataUpdateBidAskCompact{p: message.WrapFixed(b)}
}

func NewMarketDataUpdateBidAskCompact() *MarketDataUpdateBidAskCompact {
	return &MarketDataUpdateBidAskCompact{p: message.NewFixed(_MarketDataUpdateBidAskCompactDefault)}
}

func ParseMarketDataUpdateBidAskCompact(b []byte) (MarketDataUpdateBidAskCompact, error) {
	if len(b) < 4 {
		return MarketDataUpdateBidAskCompact{}, message.ErrShortBuffer
	}
	m := WrapMarketDataUpdateBidAskCompact(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return MarketDataUpdateBidAskCompact{}, message.ErrOverflow
	}
	size := int(m.p.AsUint16LE())
	if size > len(b) {
		return MarketDataUpdateBidAskCompact{}, message.ErrBaseSizeOverflow
	}
	if size < 28 {
		clone := *NewMarketDataUpdateBidAskCompact()
		clone.p.SetBytes(0, b[0:size])
		clone.p.SetBytes(size, _MarketDataUpdateBidAskCompactDefault[size:])
		return clone, nil
	}
	return m, nil
}

// Size
func (m MarketDataUpdateBidAskCompact) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type
func (m MarketDataUpdateBidAskCompact) Type() uint16 {
	return m.p.Uint16LE(2)
}

// BidPrice
func (m MarketDataUpdateBidAskCompact) BidPrice() float32 {
	return m.p.Float32LE(4)
}

// BidQuantity
func (m MarketDataUpdateBidAskCompact) BidQuantity() float32 {
	return m.p.Float32LE(8)
}

// AskPrice
func (m MarketDataUpdateBidAskCompact) AskPrice() float32 {
	return m.p.Float32LE(12)
}

// AskQuantity
func (m MarketDataUpdateBidAskCompact) AskQuantity() float32 {
	return m.p.Float32LE(16)
}

// DateTime
func (m MarketDataUpdateBidAskCompact) DateTime() DateTime4Byte {
	return DateTime4Byte(m.p.Uint32LE(20))
}

// SymbolID
func (m MarketDataUpdateBidAskCompact) SymbolID() uint32 {
	return m.p.Uint32LE(24)
}

// SetBidPrice
func (m *MarketDataUpdateBidAskCompact) SetBidPrice(value float32) *MarketDataUpdateBidAskCompact {
	m.p.SetFloat32LE(4, value)
	return m
}

// SetBidQuantity
func (m *MarketDataUpdateBidAskCompact) SetBidQuantity(value float32) *MarketDataUpdateBidAskCompact {
	m.p.SetFloat32LE(8, value)
	return m
}

// SetAskPrice
func (m *MarketDataUpdateBidAskCompact) SetAskPrice(value float32) *MarketDataUpdateBidAskCompact {
	m.p.SetFloat32LE(12, value)
	return m
}

// SetAskQuantity
func (m *MarketDataUpdateBidAskCompact) SetAskQuantity(value float32) *MarketDataUpdateBidAskCompact {
	m.p.SetFloat32LE(16, value)
	return m
}

// SetDateTime
func (m *MarketDataUpdateBidAskCompact) SetDateTime(value DateTime4Byte) *MarketDataUpdateBidAskCompact {
	m.p.SetUint32LE(20, uint32(value))
	return m
}

// SetSymbolID
func (m *MarketDataUpdateBidAskCompact) SetSymbolID(value uint32) *MarketDataUpdateBidAskCompact {
	m.p.SetUint32LE(24, value)
	return m
}

func (m MarketDataUpdateBidAskCompact) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m MarketDataUpdateBidAskCompact) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

// Copy
func (m MarketDataUpdateBidAskCompact) Copy(to MarketDataUpdateBidAskCompact) {
	to.SetBidPrice(m.BidPrice())
	to.SetBidQuantity(m.BidQuantity())
	to.SetAskPrice(m.AskPrice())
	to.SetAskQuantity(m.AskQuantity())
	to.SetDateTime(m.DateTime())
	to.SetSymbolID(m.SymbolID())
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateBidAskCompact) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 117)
	w.Float32Field("BidPrice", m.BidPrice())
	w.Float32Field("BidQuantity", m.BidQuantity())
	w.Float32Field("AskPrice", m.AskPrice())
	w.Float32Field("AskQuantity", m.AskQuantity())
	w.Uint32Field("DateTime", uint32(m.DateTime()))
	w.Uint32Field("SymbolID", m.SymbolID())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateBidAskCompact) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 117 {
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
		case "BidPrice":
			m.SetBidPrice(r.Float32())
		case "BidQuantity":
			m.SetBidQuantity(r.Float32())
		case "AskPrice":
			m.SetAskPrice(r.Float32())
		case "AskQuantity":
			m.SetAskQuantity(r.Float32())
		case "DateTime":
			m.SetDateTime(DateTime4Byte(r.Uint32()))
		case "SymbolID":
			m.SetSymbolID(r.Uint32())
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

func (m *MarketDataUpdateBidAskCompact) UnmarshalJSONReader(r *json.Reader) error {
	return m.UnmarshalJSONDoc(r)
}
