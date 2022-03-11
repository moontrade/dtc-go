//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m JournalEntryAddFixedPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":703,\"F\":["...)
	w.String(m.JournalEntry())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.DateTime()))
	return w.Finish(), nil
}

func (m JournalEntryAddFixedPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":703,\"F\":["...)
	w.String(m.JournalEntry())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m JournalEntryAddFixedPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 703)
	w.StringField("JournalEntry", m.JournalEntry())
	w.Int64Field("DateTime", int64(m.DateTime()))
	return w.Finish(), nil
}

func (m JournalEntryAddFixedPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 703)
	w.StringField("JournalEntry", m.JournalEntry())
	w.Int64Field("DateTime", int64(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *JournalEntryAddFixedPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 703 {
		return message.ErrWrongType
	}
	m.SetJournalEntry(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetDateTime(DateTime(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *JournalEntryAddFixedPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 703 {
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
		case "JournalEntry":
			m.SetJournalEntry(r.String())
		case "DateTime":
			m.SetDateTime(DateTime(r.Int64()))
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
