// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-31 22:08:18.145964 +0800 WITA m=+0.011497918

package v8

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"io"
)

const MarketDataRejectSize = 16

const MarketDataRejectFixedSize = 104

//     Size        uint16  = MarketDataRejectSize  (16)
//     Type        uint16  = MARKET_DATA_REJECT  (103)
//     BaseSize    uint16  = MarketDataRejectSize  (16)
//     SymbolID    uint32  = 0
//     RejectText  string  = ""
var _MarketDataRejectDefault = []byte{16, 0, 103, 0, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

//     Size        uint16      = MarketDataRejectFixedSize  (104)
//     Type        uint16      = MARKET_DATA_REJECT  (103)
//     SymbolID    uint32      = 0
//     RejectText  string[96]  = ""
var _MarketDataRejectFixedDefault = []byte{104, 0, 103, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// MarketDataReject The MarketDataReject message is sent by the Server to the Client to reject
// a MarketDataRequest message for any reason.
type MarketDataReject struct {
	p message.VLS
}

// MarketDataRejectFixed The MarketDataReject message is sent by the Server to the Client to reject
// a MarketDataRequest message for any reason.
type MarketDataRejectFixed struct {
	p message.Fixed
}

func NewMarketDataRejectFrom(b []byte) MarketDataReject {
	return MarketDataReject{p: message.NewVLS(b)}
}

func WrapMarketDataReject(b []byte) MarketDataReject {
	return MarketDataReject{p: message.WrapVLS(b)}
}

func NewMarketDataReject() *MarketDataReject {
	return &MarketDataReject{p: message.NewVLS(_MarketDataRejectDefault)}
}

func ParseMarketDataReject(b []byte) (MarketDataReject, error) {
	if len(b) < 6 {
		return MarketDataReject{}, message.ErrShortBuffer
	}
	m := WrapMarketDataReject(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return MarketDataReject{}, message.ErrOverflow
	}
	baseSize := int(m.p.Uint16LE(4))
	if baseSize > len(b) {
		return MarketDataReject{}, message.ErrBaseSizeOverflow
	}
	if baseSize < 16 {
		newSize := len(b) + (16 - baseSize)
		if newSize > message.MaxSize {
			return MarketDataReject{}, message.ErrOverflow
		}
		clone := MarketDataReject{message.WrapVLSUnsafe(message.Alloc(uintptr(newSize)), len(b))}
		clone.p.SetBytes(0, b[0:baseSize])
		clone.p.SetBytes(baseSize, _MarketDataRejectDefault[baseSize:])
		if len(b) > baseSize {
			shift := uint16(16 - baseSize)
			var offset uint16
			offset = clone.p.Uint16LE(12)
			if offset > 0 {
				clone.p.SetUint16LE(12, offset+shift)
			}
		}
		return clone, nil
	}
	return m, nil
}

func NewMarketDataRejectFixedFrom(b []byte) MarketDataRejectFixed {
	return MarketDataRejectFixed{p: message.NewFixed(b)}
}

func WrapMarketDataRejectFixed(b []byte) MarketDataRejectFixed {
	return MarketDataRejectFixed{p: message.WrapFixed(b)}
}

func NewMarketDataRejectFixed() *MarketDataRejectFixed {
	return &MarketDataRejectFixed{p: message.NewFixed(_MarketDataRejectFixedDefault)}
}

func ParseMarketDataRejectFixed(b []byte) (MarketDataRejectFixed, error) {
	if len(b) < 4 {
		return MarketDataRejectFixed{}, message.ErrShortBuffer
	}
	m := WrapMarketDataRejectFixed(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return MarketDataRejectFixed{}, message.ErrOverflow
	}
	size := int(m.p.AsUint16LE())
	if size > len(b) {
		return MarketDataRejectFixed{}, message.ErrBaseSizeOverflow
	}
	if size < 104 {
		clone := *NewMarketDataRejectFixed()
		clone.p.SetBytes(0, b[0:size])
		clone.p.SetBytes(size, _MarketDataRejectFixedDefault[size:])
		return clone, nil
	}
	return m, nil
}

// Size The standard message size field. Automatically set by constructor.
func (m MarketDataReject) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type The standard message type field. Automatically set by constructor.
func (m MarketDataReject) Type() uint16 {
	return m.p.Uint16LE(2)
}

// BaseSize
func (m MarketDataReject) BaseSize() uint16 {
	return m.p.Uint16LE(4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataReject) SymbolID() uint32 {
	return m.p.Uint32LE(8)
}

// RejectText Free-form text explaining the reason for the reject.
func (m MarketDataReject) RejectText() string {
	return m.p.StringVLS(12)
}

// Size The standard message size field. Automatically set by constructor.
func (m MarketDataRejectFixed) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type The standard message type field. Automatically set by constructor.
func (m MarketDataRejectFixed) Type() uint16 {
	return m.p.Uint16LE(2)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataRejectFixed) SymbolID() uint32 {
	return m.p.Uint32LE(4)
}

// RejectText Free-form text explaining the reason for the reject.
func (m MarketDataRejectFixed) RejectText() string {
	return m.p.StringFixed(8, 96)
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m *MarketDataReject) SetSymbolID(value uint32) *MarketDataReject {
	m.p.SetUint32LE(8, value)
	return m
}

// SetRejectText Free-form text explaining the reason for the reject.
func (m *MarketDataReject) SetRejectText(value string) *MarketDataReject {
	m.p.SetStringVLS(12, value)
	return m
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m *MarketDataRejectFixed) SetSymbolID(value uint32) *MarketDataRejectFixed {
	m.p.SetUint32LE(4, value)
	return m
}

// SetRejectText Free-form text explaining the reason for the reject.
func (m *MarketDataRejectFixed) SetRejectText(value string) *MarketDataRejectFixed {
	m.p.SetStringFixed(8, 96, value)
	return m
}

func (m *MarketDataReject) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m *MarketDataReject) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

func (m *MarketDataReject) Clone() *MarketDataReject {
	return &MarketDataReject{message.WrapVLSPointer(m.p.Clone(uintptr(m.Size())), int(m.Size()))}
}

func (m *MarketDataRejectFixed) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m *MarketDataRejectFixed) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

func (m *MarketDataRejectFixed) Clone() *MarketDataRejectFixed {
	return &MarketDataRejectFixed{message.WrapFixedPointer(m.p.Clone(uintptr(m.Size())))}
}

// Copy
func (m MarketDataReject) Copy(to MarketDataReject) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m MarketDataReject) CopyTo(to MarketDataRejectFixed) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m MarketDataRejectFixed) Copy(to MarketDataRejectFixed) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m MarketDataRejectFixed) CopyTo(to MarketDataReject) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataReject) MarshalJSON() ([]byte, error) {
	return m.MarshalJSONTo(nil)
}

func (m *MarketDataReject) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 103)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.StringField("RejectText", m.RejectText())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataReject) UnmarshalJSON(b []byte) error {
	r, err := json.OpenReader(b)
	if err != nil {
		return err
	}
	return m.UnmarshalJSONFromReader(&r)
}

func (m *MarketDataReject) UnmarshalJSONFromReader(r *json.Reader) error {
	if r.Type != 103 {
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
		case "RejectText":
			m.SetRejectText(r.String())
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
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataRejectFixed) MarshalJSON() ([]byte, error) {
	return m.MarshalJSONTo(nil)
}

func (m *MarketDataRejectFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 103)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.StringField("RejectText", m.RejectText())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataRejectFixed) UnmarshalJSON(b []byte) error {
	r, err := json.OpenReader(b)
	if err != nil {
		return err
	}
	return m.UnmarshalJSONFromReader(&r)
}

func (m *MarketDataRejectFixed) UnmarshalJSONFromReader(r *json.Reader) error {
	if r.Type != 103 {
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
		case "RejectText":
			m.SetRejectText(r.String())
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
