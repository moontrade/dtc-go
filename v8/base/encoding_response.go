// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-31 22:08:18.145964 +0800 WITA m=+0.011497918

package v8

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"io"
)

const EncodingResponseSize = 16

//     Size             uint16        = EncodingResponseSize  (16)
//     Type             uint16        = ENCODING_RESPONSE  (7)
//     ProtocolVersion  int32         = CURRENT_VERSION  (8)
//     Encoding         EncodingEnum  = BINARY_ENCODING  (0)
//     ProtocolType     string[4]     = "DTC"
var _EncodingResponseDefault = []byte{16, 0, 7, 0, 8, 0, 0, 0, 0, 0, 0, 0, 68, 84, 67, 0}

// EncodingResponse Requirements: Required for Servers. Required for Clients if the Client
// needs to discover the encoding the Server uses.
//
// The EncodingResponse is a message from the Server to the Client, telling
// the Client what message encoding it must use to communicate with the Server.
// the Client what message encoding it must use to communicate with the Server.
//
// For the procedure to work with this message, refer to Encoding Request
// Sequence.
type EncodingResponse struct {
	p message.Fixed
}

func NewEncodingResponseFrom(b []byte) EncodingResponse {
	return EncodingResponse{p: message.NewFixed(b)}
}

func WrapEncodingResponse(b []byte) EncodingResponse {
	return EncodingResponse{p: message.WrapFixed(b)}
}

func NewEncodingResponse() *EncodingResponse {
	return &EncodingResponse{p: message.NewFixed(_EncodingResponseDefault)}
}

func ParseEncodingResponse(b []byte) (EncodingResponse, error) {
	if len(b) < 4 {
		return EncodingResponse{}, message.ErrShortBuffer
	}
	m := WrapEncodingResponse(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return EncodingResponse{}, message.ErrOverflow
	}
	size := int(m.p.AsUint16LE())
	if size > len(b) {
		return EncodingResponse{}, message.ErrBaseSizeOverflow
	}
	if size < 16 {
		clone := *NewEncodingResponse()
		clone.p.SetBytes(0, b[0:size])
		clone.p.SetBytes(size, _EncodingResponseDefault[size:])
		return clone, nil
	}
	return m, nil
}

// Size The standard message size field. Automatically set by constructor.
func (m EncodingResponse) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type The standard message type field. Automatically set by constructor.
//
// To determine the field number for JSON, refer to this message type constant
// in the DTCProtocol.h file.
func (m EncodingResponse) Type() uint16 {
	return m.p.Uint16LE(2)
}

// ProtocolVersion This field is automatically set by the constructor.
func (m EncodingResponse) ProtocolVersion() int32 {
	return m.p.Int32LE(4)
}

// Encoding The DTC message encoding to be used.
//
// This value may be different from the requested DTC encoding if the Server
// does not support the requested encoding from the Client.
func (m EncodingResponse) Encoding() EncodingEnum {
	return EncodingEnum(m.p.Int32LE(8))
}

// ProtocolType The ProtocolType field needs to be set to the text string "DTC".
//
// This field is automatically set with the binary encoding data structures.
// This field is automatically set with the binary encoding data structures.
//
// This field is used for the Client to know that it is communicating with
// a DTC compliant Server.
func (m EncodingResponse) ProtocolType() string {
	return m.p.StringFixed(12, 4)
}

// SetProtocolVersion This field is automatically set by the constructor.
func (m *EncodingResponse) SetProtocolVersion(value int32) *EncodingResponse {
	m.p.SetInt32LE(4, value)
	return m
}

// SetEncoding The DTC message encoding to be used.
//
// This value may be different from the requested DTC encoding if the Server
// does not support the requested encoding from the Client.
func (m *EncodingResponse) SetEncoding(value EncodingEnum) *EncodingResponse {
	m.p.SetInt32LE(8, int32(value))
	return m
}

// SetProtocolType The ProtocolType field needs to be set to the text string "DTC".
//
// This field is automatically set with the binary encoding data structures.
// This field is automatically set with the binary encoding data structures.
//
// This field is used for the Client to know that it is communicating with
// a DTC compliant Server.
func (m *EncodingResponse) SetProtocolType(value string) *EncodingResponse {
	m.p.SetStringFixed(12, 4, value)
	return m
}

func (m *EncodingResponse) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m *EncodingResponse) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

func (m *EncodingResponse) Clone() *EncodingResponse {
	return &EncodingResponse{message.WrapFixedPointer(m.p.Clone(uintptr(m.Size())))}
}

// Copy
func (m EncodingResponse) Copy(to EncodingResponse) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetEncoding(m.Encoding())
	to.SetProtocolType(m.ProtocolType())
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *EncodingResponse) MarshalJSON() ([]byte, error) {
	return m.MarshalJSONTo(nil)
}

func (m *EncodingResponse) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 7)
	w.Int32Field("ProtocolVersion", m.ProtocolVersion())
	w.Int32Field("Encoding", int32(m.Encoding()))
	w.StringField("ProtocolType", m.ProtocolType())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *EncodingResponse) UnmarshalJSON(b []byte) error {
	r, err := json.OpenReader(b)
	if err != nil {
		return err
	}
	return m.UnmarshalJSONFromReader(&r)
}

func (m *EncodingResponse) UnmarshalJSONFromReader(r *json.Reader) error {
	if r.Type != 7 {
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
		case "ProtocolVersion":
			m.SetProtocolVersion(r.Int32())
		case "Encoding":
			m.SetEncoding(EncodingEnum(r.Int32()))
		case "ProtocolType":
			m.SetProtocolType(r.String())
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
