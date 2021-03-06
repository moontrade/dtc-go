// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-31 22:08:18.145964 +0800 WITA m=+0.011497918

package v8

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"io"
)

const LogoffSize = 12

const LogoffFixedSize = 102

//     Size            uint16  = LogoffSize  (12)
//     Type            uint16  = LOGOFF  (5)
//     BaseSize        uint16  = LogoffSize  (12)
//     Reason          string  = ""
//     DoNotReconnect  uint8   = 0
var _LogoffDefault = []byte{12, 0, 5, 0, 12, 0, 0, 0, 0, 0, 0, 0}

//     Size            uint16      = LogoffFixedSize  (102)
//     Type            uint16      = LOGOFF  (5)
//     Reason          string[96]  = ""
//     DoNotReconnect  uint8       = 0
var _LogoffFixedDefault = []byte{102, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// Logoff A Logoff is a message which can be sent either by the Client or the Server
// to the other side. It indicates that the Client or the Server is logging
// off and going to be closing the connection.
//
// When one side receives this message, it should expect the connection will
// be closed. It should not be expected that any messages will follow the
// Logoff message, and it should close the network connection and consider
// it finished. The side receiving this message can send a Logoff message
// to the other side before closing the connection.
type Logoff struct {
	p message.VLS
}

// LogoffFixed A Logoff is a message which can be sent either by the Client or the Server
// to the other side. It indicates that the Client or the Server is logging
// off and going to be closing the connection.
//
// When one side receives this message, it should expect the connection will
// be closed. It should not be expected that any messages will follow the
// Logoff message, and it should close the network connection and consider
// it finished. The side receiving this message can send a Logoff message
// to the other side before closing the connection.
type LogoffFixed struct {
	p message.Fixed
}

func NewLogoffFrom(b []byte) Logoff {
	return Logoff{p: message.NewVLS(b)}
}

func WrapLogoff(b []byte) Logoff {
	return Logoff{p: message.WrapVLS(b)}
}

func NewLogoff() *Logoff {
	return &Logoff{p: message.NewVLS(_LogoffDefault)}
}

func ParseLogoff(b []byte) (Logoff, error) {
	if len(b) < 6 {
		return Logoff{}, message.ErrShortBuffer
	}
	m := WrapLogoff(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return Logoff{}, message.ErrOverflow
	}
	baseSize := int(m.p.Uint16LE(4))
	if baseSize > len(b) {
		return Logoff{}, message.ErrBaseSizeOverflow
	}
	if baseSize < 12 {
		newSize := len(b) + (12 - baseSize)
		if newSize > message.MaxSize {
			return Logoff{}, message.ErrOverflow
		}
		clone := Logoff{message.WrapVLSUnsafe(message.Alloc(uintptr(newSize)), len(b))}
		clone.p.SetBytes(0, b[0:baseSize])
		clone.p.SetBytes(baseSize, _LogoffDefault[baseSize:])
		if len(b) > baseSize {
			shift := uint16(12 - baseSize)
			var offset uint16
			offset = clone.p.Uint16LE(6)
			if offset > 0 {
				clone.p.SetUint16LE(6, offset+shift)
			}
		}
		return clone, nil
	}
	return m, nil
}

func NewLogoffFixedFrom(b []byte) LogoffFixed {
	return LogoffFixed{p: message.NewFixed(b)}
}

func WrapLogoffFixed(b []byte) LogoffFixed {
	return LogoffFixed{p: message.WrapFixed(b)}
}

func NewLogoffFixed() *LogoffFixed {
	return &LogoffFixed{p: message.NewFixed(_LogoffFixedDefault)}
}

func ParseLogoffFixed(b []byte) (LogoffFixed, error) {
	if len(b) < 4 {
		return LogoffFixed{}, message.ErrShortBuffer
	}
	m := WrapLogoffFixed(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return LogoffFixed{}, message.ErrOverflow
	}
	size := int(m.p.AsUint16LE())
	if size > len(b) {
		return LogoffFixed{}, message.ErrBaseSizeOverflow
	}
	if size < 102 {
		clone := *NewLogoffFixed()
		clone.p.SetBytes(0, b[0:size])
		clone.p.SetBytes(size, _LogoffFixedDefault[size:])
		return clone, nil
	}
	return m, nil
}

// Size The standard message size field. Automatically set by constructor.
func (m Logoff) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type The standard message type field. Automatically set by constructor.
//
// To determine the field number for JSON, refer to this message type constant
// in the DTCProtocol.h file.
func (m Logoff) Type() uint16 {
	return m.p.Uint16LE(2)
}

// BaseSize
func (m Logoff) BaseSize() uint16 {
	return m.p.Uint16LE(4)
}

// Reason is a character string indicating the reason for the log off from
// either the Client or the Server.
func (m Logoff) Reason() string {
	return m.p.StringVLS(6)
}

// DoNotReconnect When DoNotReconnect is set to a 1, this indicates to the other side that
// a reconnect to the opposite side should not occur automatically.
func (m Logoff) DoNotReconnect() uint8 {
	return m.p.Uint8(10)
}

// Size The standard message size field. Automatically set by constructor.
func (m LogoffFixed) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type The standard message type field. Automatically set by constructor.
//
// To determine the field number for JSON, refer to this message type constant
// in the DTCProtocol.h file.
func (m LogoffFixed) Type() uint16 {
	return m.p.Uint16LE(2)
}

// Reason is a character string indicating the reason for the log off from
// either the Client or the Server.
func (m LogoffFixed) Reason() string {
	return m.p.StringFixed(4, 96)
}

// DoNotReconnect When DoNotReconnect is set to a 1, this indicates to the other side that
// a reconnect to the opposite side should not occur automatically.
func (m LogoffFixed) DoNotReconnect() uint8 {
	return m.p.Uint8(100)
}

// SetReason Reason is a character string indicating the reason for the log off from
// either the Client or the Server.
func (m *Logoff) SetReason(value string) *Logoff {
	m.p.SetStringVLS(6, value)
	return m
}

// SetDoNotReconnect When DoNotReconnect is set to a 1, this indicates to the other side that
// a reconnect to the opposite side should not occur automatically.
func (m *Logoff) SetDoNotReconnect(value uint8) *Logoff {
	m.p.SetUint8(10, value)
	return m
}

// SetReason Reason is a character string indicating the reason for the log off from
// either the Client or the Server.
func (m *LogoffFixed) SetReason(value string) *LogoffFixed {
	m.p.SetStringFixed(4, 96, value)
	return m
}

// SetDoNotReconnect When DoNotReconnect is set to a 1, this indicates to the other side that
// a reconnect to the opposite side should not occur automatically.
func (m *LogoffFixed) SetDoNotReconnect(value uint8) *LogoffFixed {
	m.p.SetUint8(100, value)
	return m
}

func (m *Logoff) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m *Logoff) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

func (m *Logoff) Clone() *Logoff {
	return &Logoff{message.WrapVLSPointer(m.p.Clone(uintptr(m.Size())), int(m.Size()))}
}

func (m *LogoffFixed) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m *LogoffFixed) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

func (m *LogoffFixed) Clone() *LogoffFixed {
	return &LogoffFixed{message.WrapFixedPointer(m.p.Clone(uintptr(m.Size())))}
}

// Copy
func (m Logoff) Copy(to Logoff) {
	to.SetReason(m.Reason())
	to.SetDoNotReconnect(m.DoNotReconnect())
}

// CopyTo
func (m Logoff) CopyTo(to LogoffFixed) {
	to.SetReason(m.Reason())
	to.SetDoNotReconnect(m.DoNotReconnect())
}

// Copy
func (m LogoffFixed) Copy(to LogoffFixed) {
	to.SetReason(m.Reason())
	to.SetDoNotReconnect(m.DoNotReconnect())
}

// CopyTo
func (m LogoffFixed) CopyTo(to Logoff) {
	to.SetReason(m.Reason())
	to.SetDoNotReconnect(m.DoNotReconnect())
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *Logoff) MarshalJSON() ([]byte, error) {
	return m.MarshalJSONTo(nil)
}

func (m *Logoff) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 5)
	w.StringField("Reason", m.Reason())
	w.Uint8Field("DoNotReconnect", m.DoNotReconnect())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *Logoff) UnmarshalJSON(b []byte) error {
	r, err := json.OpenReader(b)
	if err != nil {
		return err
	}
	return m.UnmarshalJSONFromReader(&r)
}

func (m *Logoff) UnmarshalJSONFromReader(r *json.Reader) error {
	if r.Type != 5 {
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
		case "Reason":
			m.SetReason(r.String())
		case "DoNotReconnect":
			m.SetDoNotReconnect(r.Uint8())
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

func (m *LogoffFixed) MarshalJSON() ([]byte, error) {
	return m.MarshalJSONTo(nil)
}

func (m *LogoffFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 5)
	w.StringField("Reason", m.Reason())
	w.Uint8Field("DoNotReconnect", m.DoNotReconnect())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *LogoffFixed) UnmarshalJSON(b []byte) error {
	r, err := json.OpenReader(b)
	if err != nil {
		return err
	}
	return m.UnmarshalJSONFromReader(&r)
}

func (m *LogoffFixed) UnmarshalJSONFromReader(r *json.Reader) error {
	if r.Type != 5 {
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
		case "Reason":
			m.SetReason(r.String())
		case "DoNotReconnect":
			m.SetDoNotReconnect(r.Uint8())
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
