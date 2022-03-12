//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m SymbolsForExchangeRequest) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":502,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.SecurityType()))
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.RequestAction()))
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	return w.Finish(), nil
}

func (m SymbolsForExchangeRequestBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":502,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.SecurityType()))
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.RequestAction()))
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m SymbolsForExchangeRequestPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":502,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.SecurityType()))
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.RequestAction()))
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	return w.Finish(), nil
}

func (m SymbolsForExchangeRequestPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":502,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.SecurityType()))
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.RequestAction()))
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m SymbolsForExchangeRequest) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 502)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Exchange", m.Exchange())
	w.Int32Field("SecurityType", int32(m.SecurityType()))
	w.Int32Field("RequestAction", int32(m.RequestAction()))
	w.StringField("Symbol", m.Symbol())
	return w.Finish(), nil
}

func (m SymbolsForExchangeRequestBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 502)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Exchange", m.Exchange())
	w.Int32Field("SecurityType", int32(m.SecurityType()))
	w.Int32Field("RequestAction", int32(m.RequestAction()))
	w.StringField("Symbol", m.Symbol())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m SymbolsForExchangeRequestPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 502)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Exchange", m.Exchange())
	w.Int32Field("SecurityType", int32(m.SecurityType()))
	w.Int32Field("RequestAction", int32(m.RequestAction()))
	w.StringField("Symbol", m.Symbol())
	return w.Finish(), nil
}

func (m SymbolsForExchangeRequestPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 502)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Exchange", m.Exchange())
	w.Int32Field("SecurityType", int32(m.SecurityType()))
	w.Int32Field("RequestAction", int32(m.RequestAction()))
	w.StringField("Symbol", m.Symbol())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SymbolsForExchangeRequestBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 502 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetExchange(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetSecurityType(SecurityTypeEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetRequestAction(RequestActionEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetSymbol(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SymbolsForExchangeRequestPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 502 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetExchange(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetSecurityType(SecurityTypeEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetRequestAction(RequestActionEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetSymbol(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SymbolsForExchangeRequestBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 502 {
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
		case "Exchange":
			m.SetExchange(r.String())
		case "SecurityType":
			m.SetSecurityType(SecurityTypeEnum(r.Int32()))
		case "RequestAction":
			m.SetRequestAction(RequestActionEnum(r.Int32()))
		case "Symbol":
			m.SetSymbol(r.String())
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

func (m *SymbolsForExchangeRequestPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 502 {
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
		case "Exchange":
			m.SetExchange(r.String())
		case "SecurityType":
			m.SetSecurityType(SecurityTypeEnum(r.Int32()))
		case "RequestAction":
			m.SetRequestAction(RequestActionEnum(r.Int32()))
		case "Symbol":
			m.SetSymbol(r.String())
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

func (m SymbolsForExchangeRequestFixed) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":502,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.SecurityType()))
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.RequestAction()))
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	return w.Finish(), nil
}

func (m SymbolsForExchangeRequestFixedBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":502,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.SecurityType()))
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.RequestAction()))
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m SymbolsForExchangeRequestFixedPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":502,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.SecurityType()))
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.RequestAction()))
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	return w.Finish(), nil
}

func (m SymbolsForExchangeRequestFixedPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":502,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.SecurityType()))
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.RequestAction()))
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m SymbolsForExchangeRequestFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 502)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Exchange", m.Exchange())
	w.Int32Field("SecurityType", int32(m.SecurityType()))
	w.Int32Field("RequestAction", int32(m.RequestAction()))
	w.StringField("Symbol", m.Symbol())
	return w.Finish(), nil
}

func (m SymbolsForExchangeRequestFixedBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 502)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Exchange", m.Exchange())
	w.Int32Field("SecurityType", int32(m.SecurityType()))
	w.Int32Field("RequestAction", int32(m.RequestAction()))
	w.StringField("Symbol", m.Symbol())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m SymbolsForExchangeRequestFixedPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 502)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Exchange", m.Exchange())
	w.Int32Field("SecurityType", int32(m.SecurityType()))
	w.Int32Field("RequestAction", int32(m.RequestAction()))
	w.StringField("Symbol", m.Symbol())
	return w.Finish(), nil
}

func (m SymbolsForExchangeRequestFixedPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 502)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Exchange", m.Exchange())
	w.Int32Field("SecurityType", int32(m.SecurityType()))
	w.Int32Field("RequestAction", int32(m.RequestAction()))
	w.StringField("Symbol", m.Symbol())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SymbolsForExchangeRequestFixedBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 502 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetExchange(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetSecurityType(SecurityTypeEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetRequestAction(RequestActionEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetSymbol(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SymbolsForExchangeRequestFixedPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 502 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetExchange(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetSecurityType(SecurityTypeEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetRequestAction(RequestActionEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetSymbol(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SymbolsForExchangeRequestFixedBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 502 {
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
		case "Exchange":
			m.SetExchange(r.String())
		case "SecurityType":
			m.SetSecurityType(SecurityTypeEnum(r.Int32()))
		case "RequestAction":
			m.SetRequestAction(RequestActionEnum(r.Int32()))
		case "Symbol":
			m.SetSymbol(r.String())
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

func (m *SymbolsForExchangeRequestFixedPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 502 {
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
		case "Exchange":
			m.SetExchange(r.String())
		case "SecurityType":
			m.SetSecurityType(SecurityTypeEnum(r.Int32()))
		case "RequestAction":
			m.SetRequestAction(RequestActionEnum(r.Int32()))
		case "Symbol":
			m.SetSymbol(r.String())
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

func (m *SymbolsForExchangeRequestBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetExchange(r.ReadString())
		case 3:
			m.SetSecurityType(SecurityTypeEnum(r.ReadInt32()))
		case 4:
			m.SetRequestAction(RequestActionEnum(r.ReadInt32()))
		case 5:
			m.SetSymbol(r.ReadString())
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

func (m *SymbolsForExchangeRequestPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetExchange(r.ReadString())
		case 3:
			m.SetSecurityType(SecurityTypeEnum(r.ReadInt32()))
		case 4:
			m.SetRequestAction(RequestActionEnum(r.ReadInt32()))
		case 5:
			m.SetSymbol(r.ReadString())
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

func (m SymbolsForExchangeRequestBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 502)
	w.WriteVarint32(1, m.RequestID())
	w.WriteString(2, m.Exchange())
	w.WriteVarint32(3, int32(m.SecurityType()))
	w.WriteVarint32(4, int32(m.RequestAction()))
	w.WriteString(5, m.Symbol())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m SymbolsForExchangeRequestPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 502)
	w.WriteVarint32(1, m.RequestID())
	w.WriteString(2, m.Exchange())
	w.WriteVarint32(3, int32(m.SecurityType()))
	w.WriteVarint32(4, int32(m.RequestAction()))
	w.WriteString(5, m.Symbol())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SymbolsForExchangeRequestFixedBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetExchange(r.ReadString())
		case 3:
			m.SetSecurityType(SecurityTypeEnum(r.ReadInt32()))
		case 4:
			m.SetRequestAction(RequestActionEnum(r.ReadInt32()))
		case 5:
			m.SetSymbol(r.ReadString())
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

func (m *SymbolsForExchangeRequestFixedPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetExchange(r.ReadString())
		case 3:
			m.SetSecurityType(SecurityTypeEnum(r.ReadInt32()))
		case 4:
			m.SetRequestAction(RequestActionEnum(r.ReadInt32()))
		case 5:
			m.SetSymbol(r.ReadString())
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

func (m SymbolsForExchangeRequestFixedBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 502)
	w.WriteVarint32(1, m.RequestID())
	w.WriteString(2, m.Exchange())
	w.WriteVarint32(3, int32(m.SecurityType()))
	w.WriteVarint32(4, int32(m.RequestAction()))
	w.WriteString(5, m.Symbol())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m SymbolsForExchangeRequestFixedPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 502)
	w.WriteVarint32(1, m.RequestID())
	w.WriteString(2, m.Exchange())
	w.WriteVarint32(3, int32(m.SecurityType()))
	w.WriteVarint32(4, int32(m.RequestAction()))
	w.WriteString(5, m.Symbol())
	return w.Finish(), nil
}