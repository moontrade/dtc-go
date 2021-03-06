// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-31 21:46:24.141427 +0800 WITA m=+0.018610501

package v8

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"io"
)

const HistoricalTradesRejectSize = 14

const HistoricalTradesRejectFixedSize = 104

//     Size        uint16  = HistoricalTradesRejectSize  (14)
//     Type        uint16  = HISTORICAL_TRADES_REJECT  (10101)
//     BaseSize    uint16  = HistoricalTradesRejectSize  (14)
//     RequestID   int32   = 0
//     RejectText  string  = ""
var _HistoricalTradesRejectDefault = []byte{14, 0, 117, 39, 14, 0, 0, 0, 0, 0, 0, 0, 0, 0}

//     Size        uint16      = HistoricalTradesRejectFixedSize  (104)
//     Type        uint16      = HISTORICAL_TRADES_REJECT  (10101)
//     RequestID   int32       = 0
//     RejectText  string[96]  = ""
var _HistoricalTradesRejectFixedDefault = []byte{104, 0, 117, 39, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type HistoricalTradesReject struct {
	p message.VLS
}

type HistoricalTradesRejectFixed struct {
	p message.Fixed
}

func NewHistoricalTradesRejectFrom(b []byte) HistoricalTradesReject {
	return HistoricalTradesReject{p: message.NewVLS(b)}
}

func WrapHistoricalTradesReject(b []byte) HistoricalTradesReject {
	return HistoricalTradesReject{p: message.WrapVLS(b)}
}

func NewHistoricalTradesReject() *HistoricalTradesReject {
	return &HistoricalTradesReject{p: message.NewVLS(_HistoricalTradesRejectDefault)}
}

func ParseHistoricalTradesReject(b []byte) (HistoricalTradesReject, error) {
	if len(b) < 6 {
		return HistoricalTradesReject{}, message.ErrShortBuffer
	}
	m := WrapHistoricalTradesReject(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return HistoricalTradesReject{}, message.ErrOverflow
	}
	baseSize := int(m.p.Uint16LE(4))
	if baseSize > len(b) {
		return HistoricalTradesReject{}, message.ErrBaseSizeOverflow
	}
	if baseSize < 14 {
		newSize := len(b) + (14 - baseSize)
		if newSize > message.MaxSize {
			return HistoricalTradesReject{}, message.ErrOverflow
		}
		clone := HistoricalTradesReject{message.WrapVLSUnsafe(message.Alloc(uintptr(newSize)), len(b))}
		clone.p.SetBytes(0, b[0:baseSize])
		clone.p.SetBytes(baseSize, _HistoricalTradesRejectDefault[baseSize:])
		if len(b) > baseSize {
			shift := uint16(14 - baseSize)
			var offset uint16
			offset = clone.p.Uint16LE(10)
			if offset > 0 {
				clone.p.SetUint16LE(10, offset+shift)
			}
		}
		return clone, nil
	}
	return m, nil
}

func NewHistoricalTradesRejectFixedFrom(b []byte) HistoricalTradesRejectFixed {
	return HistoricalTradesRejectFixed{p: message.NewFixed(b)}
}

func WrapHistoricalTradesRejectFixed(b []byte) HistoricalTradesRejectFixed {
	return HistoricalTradesRejectFixed{p: message.WrapFixed(b)}
}

func NewHistoricalTradesRejectFixed() *HistoricalTradesRejectFixed {
	return &HistoricalTradesRejectFixed{p: message.NewFixed(_HistoricalTradesRejectFixedDefault)}
}

func ParseHistoricalTradesRejectFixed(b []byte) (HistoricalTradesRejectFixed, error) {
	if len(b) < 4 {
		return HistoricalTradesRejectFixed{}, message.ErrShortBuffer
	}
	m := WrapHistoricalTradesRejectFixed(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return HistoricalTradesRejectFixed{}, message.ErrOverflow
	}
	size := int(m.p.AsUint16LE())
	if size > len(b) {
		return HistoricalTradesRejectFixed{}, message.ErrBaseSizeOverflow
	}
	if size < 104 {
		clone := *NewHistoricalTradesRejectFixed()
		clone.p.SetBytes(0, b[0:size])
		clone.p.SetBytes(size, _HistoricalTradesRejectFixedDefault[size:])
		return clone, nil
	}
	return m, nil
}

// Size
func (m HistoricalTradesReject) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type
func (m HistoricalTradesReject) Type() uint16 {
	return m.p.Uint16LE(2)
}

// BaseSize
func (m HistoricalTradesReject) BaseSize() uint16 {
	return m.p.Uint16LE(4)
}

// RequestID
func (m HistoricalTradesReject) RequestID() int32 {
	return m.p.Int32LE(6)
}

// RejectText
func (m HistoricalTradesReject) RejectText() string {
	return m.p.StringVLS(10)
}

// Size
func (m HistoricalTradesRejectFixed) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type
func (m HistoricalTradesRejectFixed) Type() uint16 {
	return m.p.Uint16LE(2)
}

// RequestID
func (m HistoricalTradesRejectFixed) RequestID() int32 {
	return m.p.Int32LE(4)
}

// RejectText
func (m HistoricalTradesRejectFixed) RejectText() string {
	return m.p.StringFixed(8, 96)
}

// SetRequestID
func (m *HistoricalTradesReject) SetRequestID(value int32) *HistoricalTradesReject {
	m.p.SetInt32LE(6, value)
	return m
}

// SetRejectText
func (m *HistoricalTradesReject) SetRejectText(value string) *HistoricalTradesReject {
	m.p.SetStringVLS(10, value)
	return m
}

// SetRequestID
func (m *HistoricalTradesRejectFixed) SetRequestID(value int32) *HistoricalTradesRejectFixed {
	m.p.SetInt32LE(4, value)
	return m
}

// SetRejectText
func (m *HistoricalTradesRejectFixed) SetRejectText(value string) *HistoricalTradesRejectFixed {
	m.p.SetStringFixed(8, 96, value)
	return m
}

func (m HistoricalTradesReject) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m HistoricalTradesReject) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

func (m HistoricalTradesRejectFixed) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m HistoricalTradesRejectFixed) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

// Copy
func (m HistoricalTradesReject) Copy(to HistoricalTradesReject) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalTradesReject) CopyTo(to HistoricalTradesRejectFixed) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m HistoricalTradesRejectFixed) Copy(to HistoricalTradesRejectFixed) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalTradesRejectFixed) CopyTo(to HistoricalTradesReject) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalTradesReject) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10101)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("RejectText", m.RejectText())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalTradesReject) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 10101 {
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
		case "RequestID":
			m.SetRequestID(r.Int32())
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
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalTradesReject) UnmarshalJSONReader(r *json.Reader) error {
	return m.UnmarshalJSONDoc(r)
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalTradesRejectFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10101)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("RejectText", m.RejectText())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalTradesRejectFixed) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 10101 {
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
		case "RequestID":
			m.SetRequestID(r.Int32())
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
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalTradesRejectFixed) UnmarshalJSONReader(r *json.Reader) error {
	return m.UnmarshalJSONDoc(r)
}
