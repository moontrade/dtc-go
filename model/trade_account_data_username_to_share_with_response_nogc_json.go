//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountDataUsernameToShareWithResponsePointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10127,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.String(m.Username())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsReadOnly())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UpdateOperationComplete())
	w.Data = append(w.Data, ',')
	w.Uint16(m.UpdateOperationMessageType())
	w.Data = append(w.Data, ',')
	w.String(m.UpdateOperationErrorText())
	return w.Finish(), nil
}

func (m TradeAccountDataUsernameToShareWithResponsePointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10127,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.String(m.Username())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsReadOnly())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UpdateOperationComplete())
	w.Data = append(w.Data, ',')
	w.Uint16(m.UpdateOperationMessageType())
	w.Data = append(w.Data, ',')
	w.String(m.UpdateOperationErrorText())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountDataUsernameToShareWithResponsePointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10127)
	w.Uint32Field("RequestID", m.RequestID())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("Username", m.Username())
	w.Uint8Field("IsReadOnly", m.IsReadOnly())
	w.Uint8Field("UpdateOperationComplete", m.UpdateOperationComplete())
	w.Uint16Field("UpdateOperationMessageType", m.UpdateOperationMessageType())
	w.StringField("UpdateOperationErrorText", m.UpdateOperationErrorText())
	return w.Finish(), nil
}

func (m TradeAccountDataUsernameToShareWithResponsePointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10127)
	w.Uint32Field("RequestID", m.RequestID())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("Username", m.Username())
	w.Uint8Field("IsReadOnly", m.IsReadOnly())
	w.Uint8Field("UpdateOperationComplete", m.UpdateOperationComplete())
	w.Uint16Field("UpdateOperationMessageType", m.UpdateOperationMessageType())
	w.StringField("UpdateOperationErrorText", m.UpdateOperationErrorText())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradeAccountDataUsernameToShareWithResponsePointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10127 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetUsername(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsReadOnly(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUpdateOperationComplete(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUpdateOperationMessageType(r.Uint16())
	if r.IsError() {
		return r.Error()
	}
	m.SetUpdateOperationErrorText(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradeAccountDataUsernameToShareWithResponsePointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 10127 {
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
		case "Username":
			m.SetUsername(r.String())
		case "IsReadOnly":
			m.SetIsReadOnly(r.Uint8())
		case "UpdateOperationComplete":
			m.SetUpdateOperationComplete(r.Uint8())
		case "UpdateOperationMessageType":
			m.SetUpdateOperationMessageType(r.Uint16())
		case "UpdateOperationErrorText":
			m.SetUpdateOperationErrorText(r.String())
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
