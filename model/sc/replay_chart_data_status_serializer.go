//go:build !tinygo

package sc

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m ReplayChartDataStatus) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10106,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.ErrorMessage())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.Status()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.SubAccountIdentifier())
	return w.Finish(), nil
}

func (m ReplayChartDataStatusBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10106,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.ErrorMessage())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.Status()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.SubAccountIdentifier())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m ReplayChartDataStatusPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10106,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.ErrorMessage())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.Status()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.SubAccountIdentifier())
	return w.Finish(), nil
}

func (m ReplayChartDataStatusPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10106,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.ErrorMessage())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.Status()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.SubAccountIdentifier())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m ReplayChartDataStatus) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10106)
	w.Uint32Field("RequestID", m.RequestID())
	w.StringField("ErrorMessage", m.ErrorMessage())
	w.Int32Field("Status", int32(m.Status()))
	w.Uint32Field("SubAccountIdentifier", m.SubAccountIdentifier())
	return w.Finish(), nil
}

func (m ReplayChartDataStatusBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10106)
	w.Uint32Field("RequestID", m.RequestID())
	w.StringField("ErrorMessage", m.ErrorMessage())
	w.Int32Field("Status", int32(m.Status()))
	w.Uint32Field("SubAccountIdentifier", m.SubAccountIdentifier())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m ReplayChartDataStatusPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10106)
	w.Uint32Field("RequestID", m.RequestID())
	w.StringField("ErrorMessage", m.ErrorMessage())
	w.Int32Field("Status", int32(m.Status()))
	w.Uint32Field("SubAccountIdentifier", m.SubAccountIdentifier())
	return w.Finish(), nil
}

func (m ReplayChartDataStatusPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10106)
	w.Uint32Field("RequestID", m.RequestID())
	w.StringField("ErrorMessage", m.ErrorMessage())
	w.Int32Field("Status", int32(m.Status()))
	w.Uint32Field("SubAccountIdentifier", m.SubAccountIdentifier())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *ReplayChartDataStatusBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10106 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetErrorMessage(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetStatus(ReplayChartDataStatusEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetSubAccountIdentifier(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *ReplayChartDataStatusPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10106 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetErrorMessage(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetStatus(ReplayChartDataStatusEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetSubAccountIdentifier(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *ReplayChartDataStatusBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 10106 {
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
		case "ErrorMessage":
			m.SetErrorMessage(r.String())
		case "Status":
			m.SetStatus(ReplayChartDataStatusEnum(r.Int32()))
		case "SubAccountIdentifier":
			m.SetSubAccountIdentifier(r.Uint32())
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

func (m *ReplayChartDataStatusPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 10106 {
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
		case "ErrorMessage":
			m.SetErrorMessage(r.String())
		case "Status":
			m.SetStatus(ReplayChartDataStatusEnum(r.Int32()))
		case "SubAccountIdentifier":
			m.SetSubAccountIdentifier(r.Uint32())
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

func (m ReplayChartDataStatusFixed) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10106,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.ErrorMessage())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.Status()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.SubAccountIdentifier())
	return w.Finish(), nil
}

func (m ReplayChartDataStatusFixedBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10106,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.ErrorMessage())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.Status()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.SubAccountIdentifier())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m ReplayChartDataStatusFixedPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10106,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.ErrorMessage())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.Status()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.SubAccountIdentifier())
	return w.Finish(), nil
}

func (m ReplayChartDataStatusFixedPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10106,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.ErrorMessage())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.Status()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.SubAccountIdentifier())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m ReplayChartDataStatusFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10106)
	w.Uint32Field("RequestID", m.RequestID())
	w.StringField("ErrorMessage", m.ErrorMessage())
	w.Int32Field("Status", int32(m.Status()))
	w.Uint32Field("SubAccountIdentifier", m.SubAccountIdentifier())
	return w.Finish(), nil
}

func (m ReplayChartDataStatusFixedBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10106)
	w.Uint32Field("RequestID", m.RequestID())
	w.StringField("ErrorMessage", m.ErrorMessage())
	w.Int32Field("Status", int32(m.Status()))
	w.Uint32Field("SubAccountIdentifier", m.SubAccountIdentifier())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m ReplayChartDataStatusFixedPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10106)
	w.Uint32Field("RequestID", m.RequestID())
	w.StringField("ErrorMessage", m.ErrorMessage())
	w.Int32Field("Status", int32(m.Status()))
	w.Uint32Field("SubAccountIdentifier", m.SubAccountIdentifier())
	return w.Finish(), nil
}

func (m ReplayChartDataStatusFixedPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10106)
	w.Uint32Field("RequestID", m.RequestID())
	w.StringField("ErrorMessage", m.ErrorMessage())
	w.Int32Field("Status", int32(m.Status()))
	w.Uint32Field("SubAccountIdentifier", m.SubAccountIdentifier())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *ReplayChartDataStatusFixedBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10106 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetErrorMessage(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetStatus(ReplayChartDataStatusEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetSubAccountIdentifier(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *ReplayChartDataStatusFixedPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10106 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetErrorMessage(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetStatus(ReplayChartDataStatusEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetSubAccountIdentifier(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *ReplayChartDataStatusFixedBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 10106 {
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
		case "ErrorMessage":
			m.SetErrorMessage(r.String())
		case "Status":
			m.SetStatus(ReplayChartDataStatusEnum(r.Int32()))
		case "SubAccountIdentifier":
			m.SetSubAccountIdentifier(r.Uint32())
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

func (m *ReplayChartDataStatusFixedPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 10106 {
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
		case "ErrorMessage":
			m.SetErrorMessage(r.String())
		case "Status":
			m.SetStatus(ReplayChartDataStatusEnum(r.Int32()))
		case "SubAccountIdentifier":
			m.SetSubAccountIdentifier(r.Uint32())
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