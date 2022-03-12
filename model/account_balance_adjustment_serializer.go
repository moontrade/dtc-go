//go:build !tinygo

// generated by github.com/moontrade/dtc-go/codegen/go at 2022-03-12 12:55:01.923209 +0800 WITA m=+0.028978918
package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m AccountBalanceAdjustment) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":607,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Float64(m.CreditAmount())
	w.Data = append(w.Data, ',')
	w.Float64(m.DebitAmount())
	w.Data = append(w.Data, ',')
	w.String(m.Currency())
	w.Data = append(w.Data, ',')
	w.String(m.Reason())
	return w.Finish(), nil
}

func (m AccountBalanceAdjustmentBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":607,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Float64(m.CreditAmount())
	w.Data = append(w.Data, ',')
	w.Float64(m.DebitAmount())
	w.Data = append(w.Data, ',')
	w.String(m.Currency())
	w.Data = append(w.Data, ',')
	w.String(m.Reason())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m AccountBalanceAdjustmentPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":607,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Float64(m.CreditAmount())
	w.Data = append(w.Data, ',')
	w.Float64(m.DebitAmount())
	w.Data = append(w.Data, ',')
	w.String(m.Currency())
	w.Data = append(w.Data, ',')
	w.String(m.Reason())
	return w.Finish(), nil
}

func (m AccountBalanceAdjustmentPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":607,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Float64(m.CreditAmount())
	w.Data = append(w.Data, ',')
	w.Float64(m.DebitAmount())
	w.Data = append(w.Data, ',')
	w.String(m.Currency())
	w.Data = append(w.Data, ',')
	w.String(m.Reason())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m AccountBalanceAdjustment) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 607)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Float64Field("CreditAmount", m.CreditAmount())
	w.Float64Field("DebitAmount", m.DebitAmount())
	w.StringField("Currency", m.Currency())
	w.StringField("Reason", m.Reason())
	return w.Finish(), nil
}

func (m AccountBalanceAdjustmentBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 607)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Float64Field("CreditAmount", m.CreditAmount())
	w.Float64Field("DebitAmount", m.DebitAmount())
	w.StringField("Currency", m.Currency())
	w.StringField("Reason", m.Reason())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m AccountBalanceAdjustmentPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 607)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Float64Field("CreditAmount", m.CreditAmount())
	w.Float64Field("DebitAmount", m.DebitAmount())
	w.StringField("Currency", m.Currency())
	w.StringField("Reason", m.Reason())
	return w.Finish(), nil
}

func (m AccountBalanceAdjustmentPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 607)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Float64Field("CreditAmount", m.CreditAmount())
	w.Float64Field("DebitAmount", m.DebitAmount())
	w.StringField("Currency", m.Currency())
	w.StringField("Reason", m.Reason())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *AccountBalanceAdjustmentBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 607 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetCreditAmount(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetDebitAmount(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetCurrency(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetReason(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *AccountBalanceAdjustmentPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 607 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetCreditAmount(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetDebitAmount(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetCurrency(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetReason(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *AccountBalanceAdjustmentBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 607 {
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
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "CreditAmount":
			m.SetCreditAmount(r.Float64())
		case "DebitAmount":
			m.SetDebitAmount(r.Float64())
		case "Currency":
			m.SetCurrency(r.String())
		case "Reason":
			m.SetReason(r.String())
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

func (m *AccountBalanceAdjustmentPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 607 {
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
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "CreditAmount":
			m.SetCreditAmount(r.Float64())
		case "DebitAmount":
			m.SetDebitAmount(r.Float64())
		case "Currency":
			m.SetCurrency(r.String())
		case "Reason":
			m.SetReason(r.String())
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
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m AccountBalanceAdjustmentFixed) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":607,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Float64(m.CreditAmount())
	w.Data = append(w.Data, ',')
	w.Float64(m.DebitAmount())
	w.Data = append(w.Data, ',')
	w.String(m.Currency())
	w.Data = append(w.Data, ',')
	w.String(m.Reason())
	return w.Finish(), nil
}

func (m AccountBalanceAdjustmentFixedBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":607,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Float64(m.CreditAmount())
	w.Data = append(w.Data, ',')
	w.Float64(m.DebitAmount())
	w.Data = append(w.Data, ',')
	w.String(m.Currency())
	w.Data = append(w.Data, ',')
	w.String(m.Reason())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m AccountBalanceAdjustmentFixedPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":607,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Float64(m.CreditAmount())
	w.Data = append(w.Data, ',')
	w.Float64(m.DebitAmount())
	w.Data = append(w.Data, ',')
	w.String(m.Currency())
	w.Data = append(w.Data, ',')
	w.String(m.Reason())
	return w.Finish(), nil
}

func (m AccountBalanceAdjustmentFixedPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":607,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Float64(m.CreditAmount())
	w.Data = append(w.Data, ',')
	w.Float64(m.DebitAmount())
	w.Data = append(w.Data, ',')
	w.String(m.Currency())
	w.Data = append(w.Data, ',')
	w.String(m.Reason())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m AccountBalanceAdjustmentFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 607)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Float64Field("CreditAmount", m.CreditAmount())
	w.Float64Field("DebitAmount", m.DebitAmount())
	w.StringField("Currency", m.Currency())
	w.StringField("Reason", m.Reason())
	return w.Finish(), nil
}

func (m AccountBalanceAdjustmentFixedBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 607)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Float64Field("CreditAmount", m.CreditAmount())
	w.Float64Field("DebitAmount", m.DebitAmount())
	w.StringField("Currency", m.Currency())
	w.StringField("Reason", m.Reason())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m AccountBalanceAdjustmentFixedPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 607)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Float64Field("CreditAmount", m.CreditAmount())
	w.Float64Field("DebitAmount", m.DebitAmount())
	w.StringField("Currency", m.Currency())
	w.StringField("Reason", m.Reason())
	return w.Finish(), nil
}

func (m AccountBalanceAdjustmentFixedPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 607)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Float64Field("CreditAmount", m.CreditAmount())
	w.Float64Field("DebitAmount", m.DebitAmount())
	w.StringField("Currency", m.Currency())
	w.StringField("Reason", m.Reason())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *AccountBalanceAdjustmentFixedBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 607 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetCreditAmount(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetDebitAmount(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetCurrency(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetReason(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *AccountBalanceAdjustmentFixedPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 607 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetCreditAmount(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetDebitAmount(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetCurrency(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetReason(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *AccountBalanceAdjustmentFixedBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 607 {
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
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "CreditAmount":
			m.SetCreditAmount(r.Float64())
		case "DebitAmount":
			m.SetDebitAmount(r.Float64())
		case "Currency":
			m.SetCurrency(r.String())
		case "Reason":
			m.SetReason(r.String())
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

func (m *AccountBalanceAdjustmentFixedPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 607 {
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
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "CreditAmount":
			m.SetCreditAmount(r.Float64())
		case "DebitAmount":
			m.SetDebitAmount(r.Float64())
		case "Currency":
			m.SetCurrency(r.String())
		case "Reason":
			m.SetReason(r.String())
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
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *AccountBalanceAdjustmentBuilder) UnmarshalProtobuf(b []byte) error {
	var (
		r   = pb.NewBuffer(b)
		f   pb.Number
		err error
	)
	for !r.EOF() {
		f, _, err = r.DecodeTag()
		if err != nil {
			break
		}
		switch f {
		case 1:
			m.SetRequestID(r.ReadInt32())
		case 2:
			m.SetCreditAmount(r.ReadFloat64())
		case 3:
			m.SetDebitAmount(r.ReadFloat64())
		case 4:
			m.SetCurrency(r.ReadString())
		case 5:
			m.SetReason(r.ReadString())
		case 6:
			m.SetTradeAccount(r.ReadString())
		default:
			r.Consume()
		}
		if err = r.Error(); err != nil {
			return err
		}
	}
	return err
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *AccountBalanceAdjustmentPointerBuilder) UnmarshalProtobuf(b []byte) error {
	var (
		r   = pb.NewBuffer(b)
		f   pb.Number
		err error
	)
	for !r.EOF() {
		f, _, err = r.DecodeTag()
		if err != nil {
			break
		}
		switch f {
		case 1:
			m.SetRequestID(r.ReadInt32())
		case 2:
			m.SetCreditAmount(r.ReadFloat64())
		case 3:
			m.SetDebitAmount(r.ReadFloat64())
		case 4:
			m.SetCurrency(r.ReadString())
		case 5:
			m.SetReason(r.ReadString())
		case 6:
			m.SetTradeAccount(r.ReadString())
		default:
			r.Consume()
		}
		if err = r.Error(); err != nil {
			return err
		}
	}
	return err
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m AccountBalanceAdjustmentBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 607)
	w.WriteVarint32(1, m.RequestID())
	w.WriteFixed64Float64(2, m.CreditAmount())
	w.WriteFixed64Float64(3, m.DebitAmount())
	w.WriteString(4, m.Currency())
	w.WriteString(5, m.Reason())
	w.WriteString(6, m.TradeAccount())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m AccountBalanceAdjustmentPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 607)
	w.WriteVarint32(1, m.RequestID())
	w.WriteFixed64Float64(2, m.CreditAmount())
	w.WriteFixed64Float64(3, m.DebitAmount())
	w.WriteString(4, m.Currency())
	w.WriteString(5, m.Reason())
	w.WriteString(6, m.TradeAccount())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *AccountBalanceAdjustmentFixedBuilder) UnmarshalProtobuf(b []byte) error {
	var (
		r   = pb.NewBuffer(b)
		f   pb.Number
		err error
	)
	for !r.EOF() {
		f, _, err = r.DecodeTag()
		if err != nil {
			break
		}
		switch f {
		case 1:
			m.SetRequestID(r.ReadInt32())
		case 2:
			m.SetCreditAmount(r.ReadFloat64())
		case 3:
			m.SetDebitAmount(r.ReadFloat64())
		case 4:
			m.SetCurrency(r.ReadString())
		case 5:
			m.SetReason(r.ReadString())
		case 6:
			m.SetTradeAccount(r.ReadString())
		default:
			r.Consume()
		}
		if err = r.Error(); err != nil {
			return err
		}
	}
	return err
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *AccountBalanceAdjustmentFixedPointerBuilder) UnmarshalProtobuf(b []byte) error {
	var (
		r   = pb.NewBuffer(b)
		f   pb.Number
		err error
	)
	for !r.EOF() {
		f, _, err = r.DecodeTag()
		if err != nil {
			break
		}
		switch f {
		case 1:
			m.SetRequestID(r.ReadInt32())
		case 2:
			m.SetCreditAmount(r.ReadFloat64())
		case 3:
			m.SetDebitAmount(r.ReadFloat64())
		case 4:
			m.SetCurrency(r.ReadString())
		case 5:
			m.SetReason(r.ReadString())
		case 6:
			m.SetTradeAccount(r.ReadString())
		default:
			r.Consume()
		}
		if err = r.Error(); err != nil {
			return err
		}
	}
	return err
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m AccountBalanceAdjustmentFixedBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 607)
	w.WriteVarint32(1, m.RequestID())
	w.WriteFixed64Float64(2, m.CreditAmount())
	w.WriteFixed64Float64(3, m.DebitAmount())
	w.WriteString(4, m.Currency())
	w.WriteString(5, m.Reason())
	w.WriteString(6, m.TradeAccount())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m AccountBalanceAdjustmentFixedPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 607)
	w.WriteVarint32(1, m.RequestID())
	w.WriteFixed64Float64(2, m.CreditAmount())
	w.WriteFixed64Float64(3, m.DebitAmount())
	w.WriteString(4, m.Currency())
	w.WriteString(5, m.Reason())
	w.WriteString(6, m.TradeAccount())
	return w.Finish(), nil
}
