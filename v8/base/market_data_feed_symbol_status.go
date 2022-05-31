// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-31 22:08:18.145964 +0800 WITA m=+0.011497918

package v8

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"io"
)

const MarketDataFeedSymbolStatusSize = 12

//     Size      uint16                    = MarketDataFeedSymbolStatusSize  (12)
//     Type      uint16                    = MARKET_DATA_FEED_SYMBOL_STATUS  (116)
//     SymbolID  uint32                    = 0
//     Status    MarketDataFeedStatusEnum  = MARKET_DATA_FEED_STATUS_UNSET  (0)
var _MarketDataFeedSymbolStatusDefault = []byte{12, 0, 116, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type MarketDataFeedSymbolStatus struct {
	p message.Fixed
}

func NewMarketDataFeedSymbolStatusFrom(b []byte) MarketDataFeedSymbolStatus {
	return MarketDataFeedSymbolStatus{p: message.NewFixed(b)}
}

func WrapMarketDataFeedSymbolStatus(b []byte) MarketDataFeedSymbolStatus {
	return MarketDataFeedSymbolStatus{p: message.WrapFixed(b)}
}

func NewMarketDataFeedSymbolStatus() *MarketDataFeedSymbolStatus {
	return &MarketDataFeedSymbolStatus{p: message.NewFixed(_MarketDataFeedSymbolStatusDefault)}
}

func ParseMarketDataFeedSymbolStatus(b []byte) (MarketDataFeedSymbolStatus, error) {
	if len(b) < 4 {
		return MarketDataFeedSymbolStatus{}, message.ErrShortBuffer
	}
	m := WrapMarketDataFeedSymbolStatus(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return MarketDataFeedSymbolStatus{}, message.ErrOverflow
	}
	size := int(m.p.AsUint16LE())
	if size > len(b) {
		return MarketDataFeedSymbolStatus{}, message.ErrBaseSizeOverflow
	}
	if size < 12 {
		clone := *NewMarketDataFeedSymbolStatus()
		clone.p.SetBytes(0, b[0:size])
		clone.p.SetBytes(size, _MarketDataFeedSymbolStatusDefault[size:])
		return clone, nil
	}
	return m, nil
}

// Size
func (m MarketDataFeedSymbolStatus) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type
func (m MarketDataFeedSymbolStatus) Type() uint16 {
	return m.p.Uint16LE(2)
}

// SymbolID
func (m MarketDataFeedSymbolStatus) SymbolID() uint32 {
	return m.p.Uint32LE(4)
}

// Status
func (m MarketDataFeedSymbolStatus) Status() MarketDataFeedStatusEnum {
	return MarketDataFeedStatusEnum(m.p.Int32LE(8))
}

// SetSymbolID
func (m *MarketDataFeedSymbolStatus) SetSymbolID(value uint32) *MarketDataFeedSymbolStatus {
	m.p.SetUint32LE(4, value)
	return m
}

// SetStatus
func (m *MarketDataFeedSymbolStatus) SetStatus(value MarketDataFeedStatusEnum) *MarketDataFeedSymbolStatus {
	m.p.SetInt32LE(8, int32(value))
	return m
}

func (m *MarketDataFeedSymbolStatus) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m *MarketDataFeedSymbolStatus) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

func (m *MarketDataFeedSymbolStatus) Clone() *MarketDataFeedSymbolStatus {
	return &MarketDataFeedSymbolStatus{message.WrapFixedPointer(m.p.Clone(uintptr(m.Size())))}
}

// Copy
func (m MarketDataFeedSymbolStatus) Copy(to MarketDataFeedSymbolStatus) {
	to.SetSymbolID(m.SymbolID())
	to.SetStatus(m.Status())
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataFeedSymbolStatus) MarshalJSON() ([]byte, error) {
	return m.MarshalJSONTo(nil)
}

func (m *MarketDataFeedSymbolStatus) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 116)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Int32Field("Status", int32(m.Status()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataFeedSymbolStatus) UnmarshalJSON(b []byte) error {
	r, err := json.OpenReader(b)
	if err != nil {
		return err
	}
	return m.UnmarshalJSONFromReader(&r)
}

func (m *MarketDataFeedSymbolStatus) UnmarshalJSONFromReader(r *json.Reader) error {
	if r.Type != 116 {
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
		case "Status":
			m.SetStatus(MarketDataFeedStatusEnum(r.Int32()))
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