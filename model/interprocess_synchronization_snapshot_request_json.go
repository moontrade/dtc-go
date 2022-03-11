//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m InterprocessSynchronizationSnapshotRequest) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10133,\"F\":["...)
	w.Uint32(m.RequestID())
	return w.Finish(), nil
}

func (m InterprocessSynchronizationSnapshotRequestBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10133,\"F\":["...)
	w.Uint32(m.RequestID())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m InterprocessSynchronizationSnapshotRequest) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10133)
	w.Uint32Field("RequestID", m.RequestID())
	return w.Finish(), nil
}

func (m InterprocessSynchronizationSnapshotRequestBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10133)
	w.Uint32Field("RequestID", m.RequestID())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *InterprocessSynchronizationSnapshotRequestBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10133 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *InterprocessSynchronizationSnapshotRequestBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 10133 {
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
