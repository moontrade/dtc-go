// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-31 21:46:24.141427 +0800 WITA m=+0.018610501

package v8

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"io"
)

const TradeAccountDataDeleteSize = 14

//     Size          uint16  = TradeAccountDataDeleteSize  (14)
//     Type          uint16  = TRADE_ACCOUNT_DATA_DELETE  (10118)
//     BaseSize      uint16  = TradeAccountDataDeleteSize  (14)
//     RequestID     uint32  = 0
//     TradeAccount  string  = ""
var _TradeAccountDataDeleteDefault = []byte{14, 0, 134, 39, 14, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type TradeAccountDataDelete struct {
	p message.VLS
}

func NewTradeAccountDataDeleteFrom(b []byte) TradeAccountDataDelete {
	return TradeAccountDataDelete{p: message.NewVLS(b)}
}

func WrapTradeAccountDataDelete(b []byte) TradeAccountDataDelete {
	return TradeAccountDataDelete{p: message.WrapVLS(b)}
}

func NewTradeAccountDataDelete() *TradeAccountDataDelete {
	return &TradeAccountDataDelete{p: message.NewVLS(_TradeAccountDataDeleteDefault)}
}

func ParseTradeAccountDataDelete(b []byte) (TradeAccountDataDelete, error) {
	if len(b) < 6 {
		return TradeAccountDataDelete{}, message.ErrShortBuffer
	}
	m := WrapTradeAccountDataDelete(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return TradeAccountDataDelete{}, message.ErrOverflow
	}
	baseSize := int(m.p.Uint16LE(4))
	if baseSize > len(b) {
		return TradeAccountDataDelete{}, message.ErrBaseSizeOverflow
	}
	if baseSize < 14 {
		newSize := len(b) + (14 - baseSize)
		if newSize > message.MaxSize {
			return TradeAccountDataDelete{}, message.ErrOverflow
		}
		clone := TradeAccountDataDelete{message.WrapVLSUnsafe(message.Alloc(uintptr(newSize)), len(b))}
		clone.p.SetBytes(0, b[0:baseSize])
		clone.p.SetBytes(baseSize, _TradeAccountDataDeleteDefault[baseSize:])
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

// Size
func (m TradeAccountDataDelete) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type
func (m TradeAccountDataDelete) Type() uint16 {
	return m.p.Uint16LE(2)
}

// BaseSize
func (m TradeAccountDataDelete) BaseSize() uint16 {
	return m.p.Uint16LE(4)
}

// RequestID
func (m TradeAccountDataDelete) RequestID() uint32 {
	return m.p.Uint32LE(6)
}

// TradeAccount
func (m TradeAccountDataDelete) TradeAccount() string {
	return m.p.StringVLS(10)
}

// SetRequestID
func (m *TradeAccountDataDelete) SetRequestID(value uint32) *TradeAccountDataDelete {
	m.p.SetUint32LE(6, value)
	return m
}

// SetTradeAccount
func (m *TradeAccountDataDelete) SetTradeAccount(value string) *TradeAccountDataDelete {
	m.p.SetStringVLS(10, value)
	return m
}

func (m TradeAccountDataDelete) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m TradeAccountDataDelete) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

// Copy
func (m TradeAccountDataDelete) Copy(to TradeAccountDataDelete) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountDataDelete) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10118)
	w.Uint32Field("RequestID", m.RequestID())
	w.StringField("TradeAccount", m.TradeAccount())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradeAccountDataDelete) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 10118 {
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
			m.SetRequestID(r.Uint32())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
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

func (m *TradeAccountDataDelete) UnmarshalJSONReader(r *json.Reader) error {
	return m.UnmarshalJSONDoc(r)
}
