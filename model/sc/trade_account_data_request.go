package sc

import (
	"github.com/moontrade/dtc-go/message"
)

const TradeAccountDataRequestSize = 14

// {
//     Size         = TradeAccountDataRequestSize  (14)
//     Type         = TRADE_ACCOUNT_DATA_REQUEST  (10115)
//     BaseSize     = TradeAccountDataRequestSize  (14)
//     RequestID    = 0
//     TradeAccount = ""
// }
var _TradeAccountDataRequestDefault = []byte{14, 0, 131, 39, 14, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type TradeAccountDataRequest struct {
	message.VLS
}

type TradeAccountDataRequestBuilder struct {
	message.VLSBuilder
}

type TradeAccountDataRequestPointer struct {
	message.VLSPointer
}

type TradeAccountDataRequestPointerBuilder struct {
	message.VLSPointerBuilder
}

func NewTradeAccountDataRequestFrom(b []byte) TradeAccountDataRequest {
	return TradeAccountDataRequest{VLS: message.NewVLSFrom(b)}
}

func WrapTradeAccountDataRequest(b []byte) TradeAccountDataRequest {
	return TradeAccountDataRequest{VLS: message.WrapVLS(b)}
}

func NewTradeAccountDataRequest() TradeAccountDataRequestBuilder {
	a := TradeAccountDataRequestBuilder{message.NewVLSBuilder(14)}
	a.Unsafe().SetBytes(0, _TradeAccountDataRequestDefault)
	return a
}

func AllocTradeAccountDataRequest() TradeAccountDataRequestPointerBuilder {
	a := TradeAccountDataRequestPointerBuilder{message.AllocVLSBuilder(14)}
	a.Ptr.SetBytes(0, _TradeAccountDataRequestDefault)
	return a
}

func AllocTradeAccountDataRequestFrom(b []byte) TradeAccountDataRequestPointer {
	return TradeAccountDataRequestPointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size         = TradeAccountDataRequestSize  (14)
//     Type         = TRADE_ACCOUNT_DATA_REQUEST  (10115)
//     BaseSize     = TradeAccountDataRequestSize  (14)
//     RequestID    = 0
//     TradeAccount = ""
// }
func (m TradeAccountDataRequestBuilder) Clear() {
	m.Unsafe().SetBytes(0, _TradeAccountDataRequestDefault)
}

// Clear
// {
//     Size         = TradeAccountDataRequestSize  (14)
//     Type         = TRADE_ACCOUNT_DATA_REQUEST  (10115)
//     BaseSize     = TradeAccountDataRequestSize  (14)
//     RequestID    = 0
//     TradeAccount = ""
// }
func (m TradeAccountDataRequestPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _TradeAccountDataRequestDefault)
}

// ToBuilder
func (m TradeAccountDataRequest) ToBuilder() TradeAccountDataRequestBuilder {
	return TradeAccountDataRequestBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m TradeAccountDataRequestPointer) ToBuilder() TradeAccountDataRequestPointerBuilder {
	return TradeAccountDataRequestPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m TradeAccountDataRequestBuilder) Finish() TradeAccountDataRequest {
	return TradeAccountDataRequest{m.VLSBuilder.Finish()}
}

// Finish
func (m *TradeAccountDataRequestPointerBuilder) Finish() TradeAccountDataRequestPointer {
	return TradeAccountDataRequestPointer{m.VLSPointerBuilder.Finish()}
}

// RequestID
func (m TradeAccountDataRequest) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m TradeAccountDataRequestBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m TradeAccountDataRequestPointer) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// RequestID
func (m TradeAccountDataRequestPointerBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// TradeAccount
func (m TradeAccountDataRequest) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// TradeAccount
func (m TradeAccountDataRequestBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// TradeAccount
func (m TradeAccountDataRequestPointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// TradeAccount
func (m TradeAccountDataRequestPointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// SetRequestID
func (m TradeAccountDataRequestBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 10, 6, value)
}

// SetRequestID
func (m TradeAccountDataRequestPointerBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Ptr, 10, 6, value)
}

// SetTradeAccount
func (m *TradeAccountDataRequestBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 14, 10, value)
}

// SetTradeAccount
func (m *TradeAccountDataRequestPointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 14, 10, value)
}

// Copy
func (m TradeAccountDataRequest) Copy(to TradeAccountDataRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// Copy
func (m TradeAccountDataRequestBuilder) Copy(to TradeAccountDataRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m TradeAccountDataRequest) CopyPointer(to TradeAccountDataRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m TradeAccountDataRequestBuilder) CopyPointer(to TradeAccountDataRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// Copy
func (m TradeAccountDataRequestPointer) Copy(to TradeAccountDataRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// Copy
func (m TradeAccountDataRequestPointerBuilder) Copy(to TradeAccountDataRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m TradeAccountDataRequestPointer) CopyPointer(to TradeAccountDataRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m TradeAccountDataRequestPointerBuilder) CopyPointer(to TradeAccountDataRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}