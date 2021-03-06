// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-31 21:46:24.141427 +0800 WITA m=+0.018610501

package v8

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"io"
)

const CurrentPositionsRejectSize = 16

const CurrentPositionsRejectFixedSize = 104

//     Size        uint16  = CurrentPositionsRejectSize  (16)
//     Type        uint16  = CURRENT_POSITIONS_REJECT  (307)
//     BaseSize    uint16  = CurrentPositionsRejectSize  (16)
//     RequestID   int32   = 0
//     RejectText  string  = ""
var _CurrentPositionsRejectDefault = []byte{16, 0, 51, 1, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

//     Size        uint16      = CurrentPositionsRejectFixedSize  (104)
//     Type        uint16      = CURRENT_POSITIONS_REJECT  (307)
//     RequestID   int32       = 0
//     RejectText  string[96]  = ""
var _CurrentPositionsRejectFixedDefault = []byte{104, 0, 51, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// CurrentPositionsReject If the Server is unable to serve the request for an CurrentPositionsRequest
// message received, for a reason other than there not being any current
// Trade positions, then send this message to the Client.
//
// This must never be sent when there are actually no Trade Positions in
// the account or accounts requested.
type CurrentPositionsReject struct {
	p message.VLS
}

// CurrentPositionsRejectFixed If the Server is unable to serve the request for an CurrentPositionsRequest
// message received, for a reason other than there not being any current
// Trade positions, then send this message to the Client.
//
// This must never be sent when there are actually no Trade Positions in
// the account or accounts requested.
type CurrentPositionsRejectFixed struct {
	p message.Fixed
}

func NewCurrentPositionsRejectFrom(b []byte) CurrentPositionsReject {
	return CurrentPositionsReject{p: message.NewVLS(b)}
}

func WrapCurrentPositionsReject(b []byte) CurrentPositionsReject {
	return CurrentPositionsReject{p: message.WrapVLS(b)}
}

func NewCurrentPositionsReject() *CurrentPositionsReject {
	return &CurrentPositionsReject{p: message.NewVLS(_CurrentPositionsRejectDefault)}
}

func ParseCurrentPositionsReject(b []byte) (CurrentPositionsReject, error) {
	if len(b) < 6 {
		return CurrentPositionsReject{}, message.ErrShortBuffer
	}
	m := WrapCurrentPositionsReject(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return CurrentPositionsReject{}, message.ErrOverflow
	}
	baseSize := int(m.p.Uint16LE(4))
	if baseSize > len(b) {
		return CurrentPositionsReject{}, message.ErrBaseSizeOverflow
	}
	if baseSize < 16 {
		newSize := len(b) + (16 - baseSize)
		if newSize > message.MaxSize {
			return CurrentPositionsReject{}, message.ErrOverflow
		}
		clone := CurrentPositionsReject{message.WrapVLSUnsafe(message.Alloc(uintptr(newSize)), len(b))}
		clone.p.SetBytes(0, b[0:baseSize])
		clone.p.SetBytes(baseSize, _CurrentPositionsRejectDefault[baseSize:])
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

func NewCurrentPositionsRejectFixedFrom(b []byte) CurrentPositionsRejectFixed {
	return CurrentPositionsRejectFixed{p: message.NewFixed(b)}
}

func WrapCurrentPositionsRejectFixed(b []byte) CurrentPositionsRejectFixed {
	return CurrentPositionsRejectFixed{p: message.WrapFixed(b)}
}

func NewCurrentPositionsRejectFixed() *CurrentPositionsRejectFixed {
	return &CurrentPositionsRejectFixed{p: message.NewFixed(_CurrentPositionsRejectFixedDefault)}
}

func ParseCurrentPositionsRejectFixed(b []byte) (CurrentPositionsRejectFixed, error) {
	if len(b) < 4 {
		return CurrentPositionsRejectFixed{}, message.ErrShortBuffer
	}
	m := WrapCurrentPositionsRejectFixed(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return CurrentPositionsRejectFixed{}, message.ErrOverflow
	}
	size := int(m.p.AsUint16LE())
	if size > len(b) {
		return CurrentPositionsRejectFixed{}, message.ErrBaseSizeOverflow
	}
	if size < 104 {
		clone := *NewCurrentPositionsRejectFixed()
		clone.p.SetBytes(0, b[0:size])
		clone.p.SetBytes(size, _CurrentPositionsRejectFixedDefault[size:])
		return clone, nil
	}
	return m, nil
}

// Size The standard message size field. Automatically set by constructor.
func (m CurrentPositionsReject) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type The standard message type field. Automatically set by constructor.
func (m CurrentPositionsReject) Type() uint16 {
	return m.p.Uint16LE(2)
}

// BaseSize
func (m CurrentPositionsReject) BaseSize() uint16 {
	return m.p.Uint16LE(4)
}

// RequestID This is set to the RequestID field sent in the CurrentPositionsRequest
// message.
func (m CurrentPositionsReject) RequestID() int32 {
	return m.p.Int32LE(8)
}

// RejectText Free-form text indicating the reason for the rejection.
//
func (m CurrentPositionsReject) RejectText() string {
	return m.p.StringVLS(12)
}

// Size The standard message size field. Automatically set by constructor.
func (m CurrentPositionsRejectFixed) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type The standard message type field. Automatically set by constructor.
func (m CurrentPositionsRejectFixed) Type() uint16 {
	return m.p.Uint16LE(2)
}

// RequestID This is set to the RequestID field sent in the CurrentPositionsRequest
// message.
func (m CurrentPositionsRejectFixed) RequestID() int32 {
	return m.p.Int32LE(4)
}

// RejectText Free-form text indicating the reason for the rejection.
//
func (m CurrentPositionsRejectFixed) RejectText() string {
	return m.p.StringFixed(8, 96)
}

// SetRequestID This is set to the RequestID field sent in the CurrentPositionsRequest
// message.
func (m *CurrentPositionsReject) SetRequestID(value int32) *CurrentPositionsReject {
	m.p.SetInt32LE(8, value)
	return m
}

// SetRejectText Free-form text indicating the reason for the rejection.
//
func (m *CurrentPositionsReject) SetRejectText(value string) *CurrentPositionsReject {
	m.p.SetStringVLS(12, value)
	return m
}

// SetRequestID This is set to the RequestID field sent in the CurrentPositionsRequest
// message.
func (m *CurrentPositionsRejectFixed) SetRequestID(value int32) *CurrentPositionsRejectFixed {
	m.p.SetInt32LE(4, value)
	return m
}

// SetRejectText Free-form text indicating the reason for the rejection.
//
func (m *CurrentPositionsRejectFixed) SetRejectText(value string) *CurrentPositionsRejectFixed {
	m.p.SetStringFixed(8, 96, value)
	return m
}

func (m CurrentPositionsReject) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m CurrentPositionsReject) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

func (m CurrentPositionsRejectFixed) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m CurrentPositionsRejectFixed) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

// Copy
func (m CurrentPositionsReject) Copy(to CurrentPositionsReject) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m CurrentPositionsReject) CopyTo(to CurrentPositionsRejectFixed) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m CurrentPositionsRejectFixed) Copy(to CurrentPositionsRejectFixed) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m CurrentPositionsRejectFixed) CopyTo(to CurrentPositionsReject) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m CurrentPositionsReject) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 307)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("RejectText", m.RejectText())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *CurrentPositionsReject) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 307 {
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

func (m *CurrentPositionsReject) UnmarshalJSONReader(r *json.Reader) error {
	return m.UnmarshalJSONDoc(r)
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m CurrentPositionsRejectFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 307)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("RejectText", m.RejectText())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *CurrentPositionsRejectFixed) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 307 {
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

func (m *CurrentPositionsRejectFixed) UnmarshalJSONReader(r *json.Reader) error {
	return m.UnmarshalJSONDoc(r)
}
