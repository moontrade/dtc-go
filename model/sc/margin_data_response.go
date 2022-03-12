// generated by github.com/moontrade/dtc-go/codegen/go at 2022-03-12 12:55:11.079124 +0800 WITA m=+0.026505376

package sc

import (
	"github.com/moontrade/dtc-go/message"
)

const MarginDataResponseSize = 34

// {
//     Size           = MarginDataResponseSize  (34)
//     Type           = MARGIN_DATA_RESPONSE  (10142)
//     BaseSize       = MarginDataResponseSize  (34)
//     RequestID      = 0
//     TradeAccount   = ""
//     ErrorText      = ""
//     ExchangeMargin = 0
//     AccountMargin  = 0
// }
var _MarginDataResponseDefault = []byte{34, 0, 158, 39, 34, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type MarginDataResponse struct {
	message.VLS
}

type MarginDataResponseBuilder struct {
	message.VLSBuilder
}

type MarginDataResponsePointer struct {
	message.VLSPointer
}

type MarginDataResponsePointerBuilder struct {
	message.VLSPointerBuilder
}

func NewMarginDataResponseFrom(b []byte) MarginDataResponse {
	return MarginDataResponse{VLS: message.NewVLSFrom(b)}
}

func WrapMarginDataResponse(b []byte) MarginDataResponse {
	return MarginDataResponse{VLS: message.WrapVLS(b)}
}

func NewMarginDataResponse() MarginDataResponseBuilder {
	a := MarginDataResponseBuilder{message.NewVLSBuilder(34)}
	a.Unsafe().SetBytes(0, _MarginDataResponseDefault)
	return a
}

func AllocMarginDataResponse() MarginDataResponsePointerBuilder {
	a := MarginDataResponsePointerBuilder{message.AllocVLSBuilder(34)}
	a.Ptr.SetBytes(0, _MarginDataResponseDefault)
	return a
}

func AllocMarginDataResponseFrom(b []byte) MarginDataResponsePointer {
	return MarginDataResponsePointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size           = MarginDataResponseSize  (34)
//     Type           = MARGIN_DATA_RESPONSE  (10142)
//     BaseSize       = MarginDataResponseSize  (34)
//     RequestID      = 0
//     TradeAccount   = ""
//     ErrorText      = ""
//     ExchangeMargin = 0
//     AccountMargin  = 0
// }
func (m MarginDataResponseBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarginDataResponseDefault)
}

// Clear
// {
//     Size           = MarginDataResponseSize  (34)
//     Type           = MARGIN_DATA_RESPONSE  (10142)
//     BaseSize       = MarginDataResponseSize  (34)
//     RequestID      = 0
//     TradeAccount   = ""
//     ErrorText      = ""
//     ExchangeMargin = 0
//     AccountMargin  = 0
// }
func (m MarginDataResponsePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarginDataResponseDefault)
}

// ToBuilder
func (m MarginDataResponse) ToBuilder() MarginDataResponseBuilder {
	return MarginDataResponseBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m MarginDataResponsePointer) ToBuilder() MarginDataResponsePointerBuilder {
	return MarginDataResponsePointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m MarginDataResponseBuilder) Finish() MarginDataResponse {
	return MarginDataResponse{m.VLSBuilder.Finish()}
}

// Finish
func (m *MarginDataResponsePointerBuilder) Finish() MarginDataResponsePointer {
	return MarginDataResponsePointer{m.VLSPointerBuilder.Finish()}
}

// RequestID
func (m MarginDataResponse) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m MarginDataResponseBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m MarginDataResponsePointer) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 10, 6)
}

// RequestID
func (m MarginDataResponsePointerBuilder) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 10, 6)
}

// TradeAccount
func (m MarginDataResponse) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// TradeAccount
func (m MarginDataResponseBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// TradeAccount
func (m MarginDataResponsePointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// TradeAccount
func (m MarginDataResponsePointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// ErrorText
func (m MarginDataResponse) ErrorText() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// ErrorText
func (m MarginDataResponseBuilder) ErrorText() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// ErrorText
func (m MarginDataResponsePointer) ErrorText() string {
	return message.StringVLS(m.Ptr, 18, 14)
}

// ErrorText
func (m MarginDataResponsePointerBuilder) ErrorText() string {
	return message.StringVLS(m.Ptr, 18, 14)
}

// ExchangeMargin
func (m MarginDataResponse) ExchangeMargin() float64 {
	return message.Float64VLS(m.Unsafe(), 26, 18)
}

// ExchangeMargin
func (m MarginDataResponseBuilder) ExchangeMargin() float64 {
	return message.Float64VLS(m.Unsafe(), 26, 18)
}

// ExchangeMargin
func (m MarginDataResponsePointer) ExchangeMargin() float64 {
	return message.Float64VLS(m.Ptr, 26, 18)
}

// ExchangeMargin
func (m MarginDataResponsePointerBuilder) ExchangeMargin() float64 {
	return message.Float64VLS(m.Ptr, 26, 18)
}

// AccountMargin
func (m MarginDataResponse) AccountMargin() float64 {
	return message.Float64VLS(m.Unsafe(), 34, 26)
}

// AccountMargin
func (m MarginDataResponseBuilder) AccountMargin() float64 {
	return message.Float64VLS(m.Unsafe(), 34, 26)
}

// AccountMargin
func (m MarginDataResponsePointer) AccountMargin() float64 {
	return message.Float64VLS(m.Ptr, 34, 26)
}

// AccountMargin
func (m MarginDataResponsePointerBuilder) AccountMargin() float64 {
	return message.Float64VLS(m.Ptr, 34, 26)
}

// SetRequestID
func (m MarginDataResponseBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 10, 6, value)
}

// SetRequestID
func (m MarginDataResponsePointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 10, 6, value)
}

// SetTradeAccount
func (m *MarginDataResponseBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 14, 10, value)
}

// SetTradeAccount
func (m *MarginDataResponsePointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 14, 10, value)
}

// SetErrorText
func (m *MarginDataResponseBuilder) SetErrorText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 18, 14, value)
}

// SetErrorText
func (m *MarginDataResponsePointerBuilder) SetErrorText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 18, 14, value)
}

// SetExchangeMargin
func (m MarginDataResponseBuilder) SetExchangeMargin(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 26, 18, value)
}

// SetExchangeMargin
func (m MarginDataResponsePointerBuilder) SetExchangeMargin(value float64) {
	message.SetFloat64VLS(m.Ptr, 26, 18, value)
}

// SetAccountMargin
func (m MarginDataResponseBuilder) SetAccountMargin(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 34, 26, value)
}

// SetAccountMargin
func (m MarginDataResponsePointerBuilder) SetAccountMargin(value float64) {
	message.SetFloat64VLS(m.Ptr, 34, 26, value)
}

// Copy
func (m MarginDataResponse) Copy(to MarginDataResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetErrorText(m.ErrorText())
	to.SetExchangeMargin(m.ExchangeMargin())
	to.SetAccountMargin(m.AccountMargin())
}

// Copy
func (m MarginDataResponseBuilder) Copy(to MarginDataResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetErrorText(m.ErrorText())
	to.SetExchangeMargin(m.ExchangeMargin())
	to.SetAccountMargin(m.AccountMargin())
}

// CopyPointer
func (m MarginDataResponse) CopyPointer(to MarginDataResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetErrorText(m.ErrorText())
	to.SetExchangeMargin(m.ExchangeMargin())
	to.SetAccountMargin(m.AccountMargin())
}

// CopyPointer
func (m MarginDataResponseBuilder) CopyPointer(to MarginDataResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetErrorText(m.ErrorText())
	to.SetExchangeMargin(m.ExchangeMargin())
	to.SetAccountMargin(m.AccountMargin())
}

// Copy
func (m MarginDataResponsePointer) Copy(to MarginDataResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetErrorText(m.ErrorText())
	to.SetExchangeMargin(m.ExchangeMargin())
	to.SetAccountMargin(m.AccountMargin())
}

// Copy
func (m MarginDataResponsePointerBuilder) Copy(to MarginDataResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetErrorText(m.ErrorText())
	to.SetExchangeMargin(m.ExchangeMargin())
	to.SetAccountMargin(m.AccountMargin())
}

// CopyPointer
func (m MarginDataResponsePointer) CopyPointer(to MarginDataResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetErrorText(m.ErrorText())
	to.SetExchangeMargin(m.ExchangeMargin())
	to.SetAccountMargin(m.AccountMargin())
}

// CopyPointer
func (m MarginDataResponsePointerBuilder) CopyPointer(to MarginDataResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetErrorText(m.ErrorText())
	to.SetExchangeMargin(m.ExchangeMargin())
	to.SetAccountMargin(m.AccountMargin())
}
