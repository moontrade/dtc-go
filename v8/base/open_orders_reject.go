// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-31 22:08:18.145964 +0800 WITA m=+0.011497918

package v8

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"io"
)

const OpenOrdersRejectSize = 16

const OpenOrdersRejectFixedSize = 104

//     Size        uint16  = OpenOrdersRejectSize  (16)
//     Type        uint16  = OPEN_ORDERS_REJECT  (302)
//     BaseSize    uint16  = OpenOrdersRejectSize  (16)
//     RequestID   int32   = 0
//     RejectText  string  = ""
var _OpenOrdersRejectDefault = []byte{16, 0, 46, 1, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

//     Size        uint16      = OpenOrdersRejectFixedSize  (104)
//     Type        uint16      = OPEN_ORDERS_REJECT  (302)
//     RequestID   int32       = 0
//     RejectText  string[96]  = ""
var _OpenOrdersRejectFixedDefault = []byte{104, 0, 46, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// OpenOrdersReject If the Server is unable to serve the request for an OpenOrdersRequest
// message received, for a reason other than there not being any open orders,
// then send this message to the Client.
type OpenOrdersReject struct {
	p message.VLS
}

// OpenOrdersRejectFixed If the Server is unable to serve the request for an OpenOrdersRequest
// message received, for a reason other than there not being any open orders,
// then send this message to the Client.
type OpenOrdersRejectFixed struct {
	p message.Fixed
}

func NewOpenOrdersRejectFrom(b []byte) OpenOrdersReject {
	return OpenOrdersReject{p: message.NewVLS(b)}
}

func WrapOpenOrdersReject(b []byte) OpenOrdersReject {
	return OpenOrdersReject{p: message.WrapVLS(b)}
}

func NewOpenOrdersReject() *OpenOrdersReject {
	return &OpenOrdersReject{p: message.NewVLS(_OpenOrdersRejectDefault)}
}

func ParseOpenOrdersReject(b []byte) (OpenOrdersReject, error) {
	if len(b) < 6 {
		return OpenOrdersReject{}, message.ErrShortBuffer
	}
	m := WrapOpenOrdersReject(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return OpenOrdersReject{}, message.ErrOverflow
	}
	baseSize := int(m.p.Uint16LE(4))
	if baseSize > len(b) {
		return OpenOrdersReject{}, message.ErrBaseSizeOverflow
	}
	if baseSize < 16 {
		newSize := len(b) + (16 - baseSize)
		if newSize > message.MaxSize {
			return OpenOrdersReject{}, message.ErrOverflow
		}
		clone := OpenOrdersReject{message.WrapVLSUnsafe(message.Alloc(uintptr(newSize)), len(b))}
		clone.p.SetBytes(0, b[0:baseSize])
		clone.p.SetBytes(baseSize, _OpenOrdersRejectDefault[baseSize:])
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

func NewOpenOrdersRejectFixedFrom(b []byte) OpenOrdersRejectFixed {
	return OpenOrdersRejectFixed{p: message.NewFixed(b)}
}

func WrapOpenOrdersRejectFixed(b []byte) OpenOrdersRejectFixed {
	return OpenOrdersRejectFixed{p: message.WrapFixed(b)}
}

func NewOpenOrdersRejectFixed() *OpenOrdersRejectFixed {
	return &OpenOrdersRejectFixed{p: message.NewFixed(_OpenOrdersRejectFixedDefault)}
}

func ParseOpenOrdersRejectFixed(b []byte) (OpenOrdersRejectFixed, error) {
	if len(b) < 4 {
		return OpenOrdersRejectFixed{}, message.ErrShortBuffer
	}
	m := WrapOpenOrdersRejectFixed(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return OpenOrdersRejectFixed{}, message.ErrOverflow
	}
	size := int(m.p.AsUint16LE())
	if size > len(b) {
		return OpenOrdersRejectFixed{}, message.ErrBaseSizeOverflow
	}
	if size < 104 {
		clone := *NewOpenOrdersRejectFixed()
		clone.p.SetBytes(0, b[0:size])
		clone.p.SetBytes(size, _OpenOrdersRejectFixedDefault[size:])
		return clone, nil
	}
	return m, nil
}

// Size The standard message size field. Automatically set by constructor.
func (m OpenOrdersReject) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type The standard message type field. Automatically set by constructor.
func (m OpenOrdersReject) Type() uint16 {
	return m.p.Uint16LE(2)
}

// BaseSize
func (m OpenOrdersReject) BaseSize() uint16 {
	return m.p.Uint16LE(4)
}

// RequestID This is set to the RequestID field sent in the OpenOrdersRequest message
// from the Client.
func (m OpenOrdersReject) RequestID() int32 {
	return m.p.Int32LE(8)
}

// RejectText Free-form text indicating the reason for rejection.
func (m OpenOrdersReject) RejectText() string {
	return m.p.StringVLS(12)
}

// Size The standard message size field. Automatically set by constructor.
func (m OpenOrdersRejectFixed) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type The standard message type field. Automatically set by constructor.
func (m OpenOrdersRejectFixed) Type() uint16 {
	return m.p.Uint16LE(2)
}

// RequestID This is set to the RequestID field sent in the OpenOrdersRequest message
// from the Client.
func (m OpenOrdersRejectFixed) RequestID() int32 {
	return m.p.Int32LE(4)
}

// RejectText Free-form text indicating the reason for rejection.
func (m OpenOrdersRejectFixed) RejectText() string {
	return m.p.StringFixed(8, 96)
}

// SetRequestID This is set to the RequestID field sent in the OpenOrdersRequest message
// from the Client.
func (m *OpenOrdersReject) SetRequestID(value int32) *OpenOrdersReject {
	m.p.SetInt32LE(8, value)
	return m
}

// SetRejectText Free-form text indicating the reason for rejection.
func (m *OpenOrdersReject) SetRejectText(value string) *OpenOrdersReject {
	m.p.SetStringVLS(12, value)
	return m
}

// SetRequestID This is set to the RequestID field sent in the OpenOrdersRequest message
// from the Client.
func (m *OpenOrdersRejectFixed) SetRequestID(value int32) *OpenOrdersRejectFixed {
	m.p.SetInt32LE(4, value)
	return m
}

// SetRejectText Free-form text indicating the reason for rejection.
func (m *OpenOrdersRejectFixed) SetRejectText(value string) *OpenOrdersRejectFixed {
	m.p.SetStringFixed(8, 96, value)
	return m
}

func (m *OpenOrdersReject) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m *OpenOrdersReject) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

func (m *OpenOrdersReject) Clone() *OpenOrdersReject {
	return &OpenOrdersReject{message.WrapVLSPointer(m.p.Clone(uintptr(m.Size())), int(m.Size()))}
}

func (m *OpenOrdersRejectFixed) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m *OpenOrdersRejectFixed) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

func (m *OpenOrdersRejectFixed) Clone() *OpenOrdersRejectFixed {
	return &OpenOrdersRejectFixed{message.WrapFixedPointer(m.p.Clone(uintptr(m.Size())))}
}

// Copy
func (m OpenOrdersReject) Copy(to OpenOrdersReject) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m OpenOrdersReject) CopyTo(to OpenOrdersRejectFixed) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m OpenOrdersRejectFixed) Copy(to OpenOrdersRejectFixed) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m OpenOrdersRejectFixed) CopyTo(to OpenOrdersReject) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *OpenOrdersReject) MarshalJSON() ([]byte, error) {
	return m.MarshalJSONTo(nil)
}

func (m *OpenOrdersReject) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 302)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("RejectText", m.RejectText())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *OpenOrdersReject) UnmarshalJSON(b []byte) error {
	r, err := json.OpenReader(b)
	if err != nil {
		return err
	}
	return m.UnmarshalJSONFromReader(&r)
}

func (m *OpenOrdersReject) UnmarshalJSONFromReader(r *json.Reader) error {
	if r.Type != 302 {
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

func (m *OpenOrdersRejectFixed) MarshalJSON() ([]byte, error) {
	return m.MarshalJSONTo(nil)
}

func (m *OpenOrdersRejectFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 302)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("RejectText", m.RejectText())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *OpenOrdersRejectFixed) UnmarshalJSON(b []byte) error {
	r, err := json.OpenReader(b)
	if err != nil {
		return err
	}
	return m.UnmarshalJSONFromReader(&r)
}

func (m *OpenOrdersRejectFixed) UnmarshalJSONFromReader(r *json.Reader) error {
	if r.Type != 302 {
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
