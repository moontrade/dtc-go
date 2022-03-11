//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarginDataResponsePointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10142,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.String(m.ErrorText())
	w.Data = append(w.Data, ',')
	w.Float64(m.ExchangeMargin())
	w.Data = append(w.Data, ',')
	w.Float64(m.AccountMargin())
	return w.Finish(), nil
}

func (m MarginDataResponsePointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10142,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.String(m.ErrorText())
	w.Data = append(w.Data, ',')
	w.Float64(m.ExchangeMargin())
	w.Data = append(w.Data, ',')
	w.Float64(m.AccountMargin())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarginDataResponsePointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10142)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("ErrorText", m.ErrorText())
	w.Float64Field("ExchangeMargin", m.ExchangeMargin())
	w.Float64Field("AccountMargin", m.AccountMargin())
	return w.Finish(), nil
}

func (m MarginDataResponsePointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10142)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("ErrorText", m.ErrorText())
	w.Float64Field("ExchangeMargin", m.ExchangeMargin())
	w.Float64Field("AccountMargin", m.AccountMargin())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarginDataResponsePointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10142 {
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
	m.SetErrorText(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetExchangeMargin(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetAccountMargin(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarginDataResponsePointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 10142 {
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
		case "ErrorText":
			m.SetErrorText(r.String())
		case "ExchangeMargin":
			m.SetExchangeMargin(r.Float64())
		case "AccountMargin":
			m.SetAccountMargin(r.Float64())
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
