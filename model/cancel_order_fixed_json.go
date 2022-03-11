//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m CancelOrderFixed) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":203,\"F\":["...)
	w.String(m.ServerOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.ClientOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	return w.Finish(), nil
}

func (m CancelOrderFixedBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":203,\"F\":["...)
	w.String(m.ServerOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.ClientOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m CancelOrderFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 203)
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.StringField("TradeAccount", m.TradeAccount())
	return w.Finish(), nil
}

func (m CancelOrderFixedBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 203)
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.StringField("TradeAccount", m.TradeAccount())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *CancelOrderFixedBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 203 {
		return message.ErrWrongType
	}
	m.SetServerOrderID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetClientOrderID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *CancelOrderFixedBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 203 {
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
		case "ServerOrderID":
			m.SetServerOrderID(r.String())
		case "ClientOrderID":
			m.SetClientOrderID(r.String())
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
