// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-31 22:08:18.145964 +0800 WITA m=+0.011497918

package v8

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"io"
)

const JournalEntriesRejectSize = 16

const JournalEntriesRejectFixedSize = 104

//     Size        uint16  = JournalEntriesRejectSize  (16)
//     Type        uint16  = JOURNAL_ENTRIES_REJECT  (705)
//     BaseSize    uint16  = JournalEntriesRejectSize  (16)
//     RequestID   int32   = 0
//     RejectText  string  = ""
var _JournalEntriesRejectDefault = []byte{16, 0, 193, 2, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

//     Size        uint16      = JournalEntriesRejectFixedSize  (104)
//     Type        uint16      = JOURNAL_ENTRIES_REJECT  (705)
//     RequestID   int32       = 0
//     RejectText  string[96]  = ""
var _JournalEntriesRejectFixedDefault = []byte{104, 0, 193, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type JournalEntriesReject struct {
	p message.VLS
}

type JournalEntriesRejectFixed struct {
	p message.Fixed
}

func NewJournalEntriesRejectFrom(b []byte) JournalEntriesReject {
	return JournalEntriesReject{p: message.NewVLS(b)}
}

func WrapJournalEntriesReject(b []byte) JournalEntriesReject {
	return JournalEntriesReject{p: message.WrapVLS(b)}
}

func NewJournalEntriesReject() *JournalEntriesReject {
	return &JournalEntriesReject{p: message.NewVLS(_JournalEntriesRejectDefault)}
}

func ParseJournalEntriesReject(b []byte) (JournalEntriesReject, error) {
	if len(b) < 6 {
		return JournalEntriesReject{}, message.ErrShortBuffer
	}
	m := WrapJournalEntriesReject(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return JournalEntriesReject{}, message.ErrOverflow
	}
	baseSize := int(m.p.Uint16LE(4))
	if baseSize > len(b) {
		return JournalEntriesReject{}, message.ErrBaseSizeOverflow
	}
	if baseSize < 16 {
		newSize := len(b) + (16 - baseSize)
		if newSize > message.MaxSize {
			return JournalEntriesReject{}, message.ErrOverflow
		}
		clone := JournalEntriesReject{message.WrapVLSUnsafe(message.Alloc(uintptr(newSize)), len(b))}
		clone.p.SetBytes(0, b[0:baseSize])
		clone.p.SetBytes(baseSize, _JournalEntriesRejectDefault[baseSize:])
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

func NewJournalEntriesRejectFixedFrom(b []byte) JournalEntriesRejectFixed {
	return JournalEntriesRejectFixed{p: message.NewFixed(b)}
}

func WrapJournalEntriesRejectFixed(b []byte) JournalEntriesRejectFixed {
	return JournalEntriesRejectFixed{p: message.WrapFixed(b)}
}

func NewJournalEntriesRejectFixed() *JournalEntriesRejectFixed {
	return &JournalEntriesRejectFixed{p: message.NewFixed(_JournalEntriesRejectFixedDefault)}
}

func ParseJournalEntriesRejectFixed(b []byte) (JournalEntriesRejectFixed, error) {
	if len(b) < 4 {
		return JournalEntriesRejectFixed{}, message.ErrShortBuffer
	}
	m := WrapJournalEntriesRejectFixed(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return JournalEntriesRejectFixed{}, message.ErrOverflow
	}
	size := int(m.p.AsUint16LE())
	if size > len(b) {
		return JournalEntriesRejectFixed{}, message.ErrBaseSizeOverflow
	}
	if size < 104 {
		clone := *NewJournalEntriesRejectFixed()
		clone.p.SetBytes(0, b[0:size])
		clone.p.SetBytes(size, _JournalEntriesRejectFixedDefault[size:])
		return clone, nil
	}
	return m, nil
}

// Size
func (m JournalEntriesReject) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type
func (m JournalEntriesReject) Type() uint16 {
	return m.p.Uint16LE(2)
}

// BaseSize
func (m JournalEntriesReject) BaseSize() uint16 {
	return m.p.Uint16LE(4)
}

// RequestID
func (m JournalEntriesReject) RequestID() int32 {
	return m.p.Int32LE(8)
}

// RejectText
func (m JournalEntriesReject) RejectText() string {
	return m.p.StringVLS(12)
}

// Size
func (m JournalEntriesRejectFixed) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type
func (m JournalEntriesRejectFixed) Type() uint16 {
	return m.p.Uint16LE(2)
}

// RequestID
func (m JournalEntriesRejectFixed) RequestID() int32 {
	return m.p.Int32LE(4)
}

// RejectText
func (m JournalEntriesRejectFixed) RejectText() string {
	return m.p.StringFixed(8, 96)
}

// SetRequestID
func (m *JournalEntriesReject) SetRequestID(value int32) *JournalEntriesReject {
	m.p.SetInt32LE(8, value)
	return m
}

// SetRejectText
func (m *JournalEntriesReject) SetRejectText(value string) *JournalEntriesReject {
	m.p.SetStringVLS(12, value)
	return m
}

// SetRequestID
func (m *JournalEntriesRejectFixed) SetRequestID(value int32) *JournalEntriesRejectFixed {
	m.p.SetInt32LE(4, value)
	return m
}

// SetRejectText
func (m *JournalEntriesRejectFixed) SetRejectText(value string) *JournalEntriesRejectFixed {
	m.p.SetStringFixed(8, 96, value)
	return m
}

func (m *JournalEntriesReject) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m *JournalEntriesReject) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

func (m *JournalEntriesReject) Clone() *JournalEntriesReject {
	return &JournalEntriesReject{message.WrapVLSPointer(m.p.Clone(uintptr(m.Size())), int(m.Size()))}
}

func (m *JournalEntriesRejectFixed) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m *JournalEntriesRejectFixed) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

func (m *JournalEntriesRejectFixed) Clone() *JournalEntriesRejectFixed {
	return &JournalEntriesRejectFixed{message.WrapFixedPointer(m.p.Clone(uintptr(m.Size())))}
}

// Copy
func (m JournalEntriesReject) Copy(to JournalEntriesReject) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m JournalEntriesReject) CopyTo(to JournalEntriesRejectFixed) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m JournalEntriesRejectFixed) Copy(to JournalEntriesRejectFixed) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m JournalEntriesRejectFixed) CopyTo(to JournalEntriesReject) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *JournalEntriesReject) MarshalJSON() ([]byte, error) {
	return m.MarshalJSONTo(nil)
}

func (m *JournalEntriesReject) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 705)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("RejectText", m.RejectText())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *JournalEntriesReject) UnmarshalJSON(b []byte) error {
	r, err := json.OpenReader(b)
	if err != nil {
		return err
	}
	return m.UnmarshalJSONFromReader(&r)
}

func (m *JournalEntriesReject) UnmarshalJSONFromReader(r *json.Reader) error {
	if r.Type != 705 {
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
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *JournalEntriesRejectFixed) MarshalJSON() ([]byte, error) {
	return m.MarshalJSONTo(nil)
}

func (m *JournalEntriesRejectFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 705)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("RejectText", m.RejectText())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *JournalEntriesRejectFixed) UnmarshalJSON(b []byte) error {
	r, err := json.OpenReader(b)
	if err != nil {
		return err
	}
	return m.UnmarshalJSONFromReader(&r)
}

func (m *JournalEntriesRejectFixed) UnmarshalJSONFromReader(r *json.Reader) error {
	if r.Type != 705 {
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