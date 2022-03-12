package model

import (
	"github.com/moontrade/dtc-go/message"
)

const ExchangeListRequestSize = 8

// {
//     Size      = ExchangeListRequestSize  (8)
//     Type      = EXCHANGE_LIST_REQUEST  (500)
//     RequestID = 0
// }
var _ExchangeListRequestDefault = []byte{8, 0, 244, 1, 0, 0, 0, 0}

// ExchangeListRequest This is a message from the Client to the Server to request a list of all
// available exchanges from the Server.
//
// The server will respond with a separate ExchangeListResponse message for
// each exchange.
//
// In the case where the Server does not specify an exchange with its symbols,
// then the Server should provide a single response with an empty Exchange.
// then the Server should provide a single response with an empty Exchange.
type ExchangeListRequest struct {
	message.Fixed
}

// ExchangeListRequestBuilder This is a message from the Client to the Server to request a list of all
// available exchanges from the Server.
//
// The server will respond with a separate ExchangeListResponse message for
// each exchange.
//
// In the case where the Server does not specify an exchange with its symbols,
// then the Server should provide a single response with an empty Exchange.
// then the Server should provide a single response with an empty Exchange.
type ExchangeListRequestBuilder struct {
	message.Fixed
}

// ExchangeListRequestPointer This is a message from the Client to the Server to request a list of all
// available exchanges from the Server.
//
// The server will respond with a separate ExchangeListResponse message for
// each exchange.
//
// In the case where the Server does not specify an exchange with its symbols,
// then the Server should provide a single response with an empty Exchange.
// then the Server should provide a single response with an empty Exchange.
type ExchangeListRequestPointer struct {
	message.FixedPointer
}

// ExchangeListRequestPointerBuilder This is a message from the Client to the Server to request a list of all
// available exchanges from the Server.
//
// The server will respond with a separate ExchangeListResponse message for
// each exchange.
//
// In the case where the Server does not specify an exchange with its symbols,
// then the Server should provide a single response with an empty Exchange.
// then the Server should provide a single response with an empty Exchange.
type ExchangeListRequestPointerBuilder struct {
	message.FixedPointer
}

func NewExchangeListRequestFrom(b []byte) ExchangeListRequest {
	return ExchangeListRequest{Fixed: message.NewFixedFrom(b)}
}

func WrapExchangeListRequest(b []byte) ExchangeListRequest {
	return ExchangeListRequest{Fixed: message.WrapFixed(b)}
}

func NewExchangeListRequest() ExchangeListRequestBuilder {
	a := ExchangeListRequestBuilder{message.NewFixed(8)}
	a.Unsafe().SetBytes(0, _ExchangeListRequestDefault)
	return a
}

func AllocExchangeListRequest() ExchangeListRequestPointerBuilder {
	a := ExchangeListRequestPointerBuilder{message.AllocFixed(8)}
	a.Ptr.SetBytes(0, _ExchangeListRequestDefault)
	return a
}

func AllocExchangeListRequestFrom(b []byte) ExchangeListRequestPointer {
	return ExchangeListRequestPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size      = ExchangeListRequestSize  (8)
//     Type      = EXCHANGE_LIST_REQUEST  (500)
//     RequestID = 0
// }
func (m ExchangeListRequestBuilder) Clear() {
	m.Unsafe().SetBytes(0, _ExchangeListRequestDefault)
}

// Clear
// {
//     Size      = ExchangeListRequestSize  (8)
//     Type      = EXCHANGE_LIST_REQUEST  (500)
//     RequestID = 0
// }
func (m ExchangeListRequestPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _ExchangeListRequestDefault)
}

// ToBuilder
func (m ExchangeListRequest) ToBuilder() ExchangeListRequestBuilder {
	return ExchangeListRequestBuilder{m.Fixed}
}

// ToBuilder
func (m ExchangeListRequestPointer) ToBuilder() ExchangeListRequestPointerBuilder {
	return ExchangeListRequestPointerBuilder{m.FixedPointer}
}

// Finish
func (m ExchangeListRequestBuilder) Finish() ExchangeListRequest {
	return ExchangeListRequest{m.Fixed}
}

// Finish
func (m *ExchangeListRequestPointerBuilder) Finish() ExchangeListRequestPointer {
	return ExchangeListRequestPointer{m.FixedPointer}
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the ExchangeListResponse message.
func (m ExchangeListRequest) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the ExchangeListResponse message.
func (m ExchangeListRequestBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the ExchangeListResponse message.
func (m ExchangeListRequestPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the ExchangeListResponse message.
func (m ExchangeListRequestPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// SetRequestID The unique identifier for this request. This same identifier will be returned
// in the ExchangeListResponse message.
func (m ExchangeListRequestBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRequestID The unique identifier for this request. This same identifier will be returned
// in the ExchangeListResponse message.
func (m ExchangeListRequestPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// Copy
func (m ExchangeListRequest) Copy(to ExchangeListRequestBuilder) {
	to.SetRequestID(m.RequestID())
}

// Copy
func (m ExchangeListRequestBuilder) Copy(to ExchangeListRequestBuilder) {
	to.SetRequestID(m.RequestID())
}

// CopyPointer
func (m ExchangeListRequest) CopyPointer(to ExchangeListRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
}

// CopyPointer
func (m ExchangeListRequestBuilder) CopyPointer(to ExchangeListRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
}

// Copy
func (m ExchangeListRequestPointer) Copy(to ExchangeListRequestBuilder) {
	to.SetRequestID(m.RequestID())
}

// Copy
func (m ExchangeListRequestPointerBuilder) Copy(to ExchangeListRequestBuilder) {
	to.SetRequestID(m.RequestID())
}

// CopyPointer
func (m ExchangeListRequestPointer) CopyPointer(to ExchangeListRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
}

// CopyPointer
func (m ExchangeListRequestPointerBuilder) CopyPointer(to ExchangeListRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
}