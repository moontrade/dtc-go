// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-31 21:46:24.141427 +0800 WITA m=+0.018610501

package v8

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"io"
)

const TradeAccountDataSymbolCommissionUpdateSize = 27

//     Size                 uint16   = TradeAccountDataSymbolCommissionUpdateSize  (27)
//     Type                 uint16   = TRADE_ACCOUNT_DATA_SYMBOL_COMMISSION_UPDATE  (10120)
//     BaseSize             uint16   = TradeAccountDataSymbolCommissionUpdateSize  (27)
//     RequestID            uint32   = 0
//     IsDeleted            bool     = false
//     TradeAccount         string   = ""
//     Symbol               string   = ""
//     TradeFeePerContract  float64  = 0.000000
var _TradeAccountDataSymbolCommissionUpdateDefault = []byte{27, 0, 136, 39, 27, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type TradeAccountDataSymbolCommissionUpdate struct {
	p message.VLS
}

func NewTradeAccountDataSymbolCommissionUpdateFrom(b []byte) TradeAccountDataSymbolCommissionUpdate {
	return TradeAccountDataSymbolCommissionUpdate{p: message.NewVLS(b)}
}

func WrapTradeAccountDataSymbolCommissionUpdate(b []byte) TradeAccountDataSymbolCommissionUpdate {
	return TradeAccountDataSymbolCommissionUpdate{p: message.WrapVLS(b)}
}

func NewTradeAccountDataSymbolCommissionUpdate() *TradeAccountDataSymbolCommissionUpdate {
	return &TradeAccountDataSymbolCommissionUpdate{p: message.NewVLS(_TradeAccountDataSymbolCommissionUpdateDefault)}
}

func ParseTradeAccountDataSymbolCommissionUpdate(b []byte) (TradeAccountDataSymbolCommissionUpdate, error) {
	if len(b) < 6 {
		return TradeAccountDataSymbolCommissionUpdate{}, message.ErrShortBuffer
	}
	m := WrapTradeAccountDataSymbolCommissionUpdate(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return TradeAccountDataSymbolCommissionUpdate{}, message.ErrOverflow
	}
	baseSize := int(m.p.Uint16LE(4))
	if baseSize > len(b) {
		return TradeAccountDataSymbolCommissionUpdate{}, message.ErrBaseSizeOverflow
	}
	if baseSize < 27 {
		newSize := len(b) + (27 - baseSize)
		if newSize > message.MaxSize {
			return TradeAccountDataSymbolCommissionUpdate{}, message.ErrOverflow
		}
		clone := TradeAccountDataSymbolCommissionUpdate{message.WrapVLSUnsafe(message.Alloc(uintptr(newSize)), len(b))}
		clone.p.SetBytes(0, b[0:baseSize])
		clone.p.SetBytes(baseSize, _TradeAccountDataSymbolCommissionUpdateDefault[baseSize:])
		if len(b) > baseSize {
			shift := uint16(27 - baseSize)
			var offset uint16
			offset = clone.p.Uint16LE(11)
			if offset > 0 {
				clone.p.SetUint16LE(11, offset+shift)
			}
			offset = clone.p.Uint16LE(15)
			if offset > 0 {
				clone.p.SetUint16LE(15, offset+shift)
			}
		}
		return clone, nil
	}
	return m, nil
}

// Size
func (m TradeAccountDataSymbolCommissionUpdate) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type
func (m TradeAccountDataSymbolCommissionUpdate) Type() uint16 {
	return m.p.Uint16LE(2)
}

// BaseSize
func (m TradeAccountDataSymbolCommissionUpdate) BaseSize() uint16 {
	return m.p.Uint16LE(4)
}

// RequestID
func (m TradeAccountDataSymbolCommissionUpdate) RequestID() uint32 {
	return m.p.Uint32LE(6)
}

// IsDeleted
func (m TradeAccountDataSymbolCommissionUpdate) IsDeleted() bool {
	return m.p.Bool(10)
}

// TradeAccount
func (m TradeAccountDataSymbolCommissionUpdate) TradeAccount() string {
	return m.p.StringVLS(11)
}

// Symbol
func (m TradeAccountDataSymbolCommissionUpdate) Symbol() string {
	return m.p.StringVLS(15)
}

// TradeFeePerContract
func (m TradeAccountDataSymbolCommissionUpdate) TradeFeePerContract() float64 {
	return m.p.Float64LE(19)
}

// SetRequestID
func (m *TradeAccountDataSymbolCommissionUpdate) SetRequestID(value uint32) *TradeAccountDataSymbolCommissionUpdate {
	m.p.SetUint32LE(6, value)
	return m
}

// SetIsDeleted
func (m *TradeAccountDataSymbolCommissionUpdate) SetIsDeleted(value bool) *TradeAccountDataSymbolCommissionUpdate {
	m.p.SetBool(10, value)
	return m
}

// SetTradeAccount
func (m *TradeAccountDataSymbolCommissionUpdate) SetTradeAccount(value string) *TradeAccountDataSymbolCommissionUpdate {
	m.p.SetStringVLS(11, value)
	return m
}

// SetSymbol
func (m *TradeAccountDataSymbolCommissionUpdate) SetSymbol(value string) *TradeAccountDataSymbolCommissionUpdate {
	m.p.SetStringVLS(15, value)
	return m
}

// SetTradeFeePerContract
func (m *TradeAccountDataSymbolCommissionUpdate) SetTradeFeePerContract(value float64) *TradeAccountDataSymbolCommissionUpdate {
	m.p.SetFloat64LE(19, value)
	return m
}

func (m TradeAccountDataSymbolCommissionUpdate) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m TradeAccountDataSymbolCommissionUpdate) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

// Copy
func (m TradeAccountDataSymbolCommissionUpdate) Copy(to TradeAccountDataSymbolCommissionUpdate) {
	to.SetRequestID(m.RequestID())
	to.SetIsDeleted(m.IsDeleted())
	to.SetTradeAccount(m.TradeAccount())
	to.SetSymbol(m.Symbol())
	to.SetTradeFeePerContract(m.TradeFeePerContract())
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountDataSymbolCommissionUpdate) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10120)
	w.Uint32Field("RequestID", m.RequestID())
	w.BoolField("IsDeleted", m.IsDeleted())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("Symbol", m.Symbol())
	w.Float64Field("TradeFeePerContract", m.TradeFeePerContract())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradeAccountDataSymbolCommissionUpdate) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 10120 {
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
		case "IsDeleted":
			m.SetIsDeleted(r.Bool())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "Symbol":
			m.SetSymbol(r.String())
		case "TradeFeePerContract":
			m.SetTradeFeePerContract(r.Float64())
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

func (m *TradeAccountDataSymbolCommissionUpdate) UnmarshalJSONReader(r *json.Reader) error {
	return m.UnmarshalJSONDoc(r)
}