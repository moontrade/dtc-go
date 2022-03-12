package model

import (
	"github.com/moontrade/dtc-go/message"
)

const AccountBalanceRequestSize = 16

const AccountBalanceRequestFixedSize = 40

// {
//     Size         = AccountBalanceRequestSize  (16)
//     Type         = ACCOUNT_BALANCE_REQUEST  (601)
//     BaseSize     = AccountBalanceRequestSize  (16)
//     RequestID    = 0
//     TradeAccount = ""
// }
var _AccountBalanceRequestDefault = []byte{16, 0, 89, 2, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// {
//     Size         = AccountBalanceRequestFixedSize  (40)
//     Type         = ACCOUNT_BALANCE_REQUEST  (601)
//     RequestID    = 0
//     TradeAccount = ""
// }
var _AccountBalanceRequestFixedDefault = []byte{40, 0, 89, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// AccountBalanceRequest This is a message from the Client to the Server to request Trade Account
// Balance data.
//
// The Server will respond with an AccountBalanceUpdate or reject the request.
// The Server will respond with an AccountBalanceUpdate or reject the request.
//
// The Server will set the RequestID in the AccountBalanceUpdate message
// to match the RequestID in the AccountBalanceRequest.
//
// The Server will periodically send AccountBalanceUpdate messages as the
// Account Balance data changes. The frequency of the updates is determined
// by the Server. Account Balance updates are considered automatically subscribed
// to. When unsolicited AccountBalanceUpdate messages are sent by the Server,
// the RequestID will be 0.
type AccountBalanceRequest struct {
	message.VLS
}

// AccountBalanceRequestBuilder This is a message from the Client to the Server to request Trade Account
// Balance data.
//
// The Server will respond with an AccountBalanceUpdate or reject the request.
// The Server will respond with an AccountBalanceUpdate or reject the request.
//
// The Server will set the RequestID in the AccountBalanceUpdate message
// to match the RequestID in the AccountBalanceRequest.
//
// The Server will periodically send AccountBalanceUpdate messages as the
// Account Balance data changes. The frequency of the updates is determined
// by the Server. Account Balance updates are considered automatically subscribed
// to. When unsolicited AccountBalanceUpdate messages are sent by the Server,
// the RequestID will be 0.
type AccountBalanceRequestBuilder struct {
	message.VLSBuilder
}

// AccountBalanceRequestFixed This is a message from the Client to the Server to request Trade Account
// Balance data.
//
// The Server will respond with an AccountBalanceUpdate or reject the request.
// The Server will respond with an AccountBalanceUpdate or reject the request.
//
// The Server will set the RequestID in the AccountBalanceUpdate message
// to match the RequestID in the AccountBalanceRequest.
//
// The Server will periodically send AccountBalanceUpdate messages as the
// Account Balance data changes. The frequency of the updates is determined
// by the Server. Account Balance updates are considered automatically subscribed
// to. When unsolicited AccountBalanceUpdate messages are sent by the Server,
// the RequestID will be 0.
type AccountBalanceRequestFixed struct {
	message.Fixed
}

// AccountBalanceRequestFixedBuilder This is a message from the Client to the Server to request Trade Account
// Balance data.
//
// The Server will respond with an AccountBalanceUpdate or reject the request.
// The Server will respond with an AccountBalanceUpdate or reject the request.
//
// The Server will set the RequestID in the AccountBalanceUpdate message
// to match the RequestID in the AccountBalanceRequest.
//
// The Server will periodically send AccountBalanceUpdate messages as the
// Account Balance data changes. The frequency of the updates is determined
// by the Server. Account Balance updates are considered automatically subscribed
// to. When unsolicited AccountBalanceUpdate messages are sent by the Server,
// the RequestID will be 0.
type AccountBalanceRequestFixedBuilder struct {
	message.Fixed
}

// AccountBalanceRequestPointer This is a message from the Client to the Server to request Trade Account
// Balance data.
//
// The Server will respond with an AccountBalanceUpdate or reject the request.
// The Server will respond with an AccountBalanceUpdate or reject the request.
//
// The Server will set the RequestID in the AccountBalanceUpdate message
// to match the RequestID in the AccountBalanceRequest.
//
// The Server will periodically send AccountBalanceUpdate messages as the
// Account Balance data changes. The frequency of the updates is determined
// by the Server. Account Balance updates are considered automatically subscribed
// to. When unsolicited AccountBalanceUpdate messages are sent by the Server,
// the RequestID will be 0.
type AccountBalanceRequestPointer struct {
	message.VLSPointer
}

// AccountBalanceRequestPointerBuilder This is a message from the Client to the Server to request Trade Account
// Balance data.
//
// The Server will respond with an AccountBalanceUpdate or reject the request.
// The Server will respond with an AccountBalanceUpdate or reject the request.
//
// The Server will set the RequestID in the AccountBalanceUpdate message
// to match the RequestID in the AccountBalanceRequest.
//
// The Server will periodically send AccountBalanceUpdate messages as the
// Account Balance data changes. The frequency of the updates is determined
// by the Server. Account Balance updates are considered automatically subscribed
// to. When unsolicited AccountBalanceUpdate messages are sent by the Server,
// the RequestID will be 0.
type AccountBalanceRequestPointerBuilder struct {
	message.VLSPointerBuilder
}

// AccountBalanceRequestFixedPointer This is a message from the Client to the Server to request Trade Account
// Balance data.
//
// The Server will respond with an AccountBalanceUpdate or reject the request.
// The Server will respond with an AccountBalanceUpdate or reject the request.
//
// The Server will set the RequestID in the AccountBalanceUpdate message
// to match the RequestID in the AccountBalanceRequest.
//
// The Server will periodically send AccountBalanceUpdate messages as the
// Account Balance data changes. The frequency of the updates is determined
// by the Server. Account Balance updates are considered automatically subscribed
// to. When unsolicited AccountBalanceUpdate messages are sent by the Server,
// the RequestID will be 0.
type AccountBalanceRequestFixedPointer struct {
	message.FixedPointer
}

// AccountBalanceRequestFixedPointerBuilder This is a message from the Client to the Server to request Trade Account
// Balance data.
//
// The Server will respond with an AccountBalanceUpdate or reject the request.
// The Server will respond with an AccountBalanceUpdate or reject the request.
//
// The Server will set the RequestID in the AccountBalanceUpdate message
// to match the RequestID in the AccountBalanceRequest.
//
// The Server will periodically send AccountBalanceUpdate messages as the
// Account Balance data changes. The frequency of the updates is determined
// by the Server. Account Balance updates are considered automatically subscribed
// to. When unsolicited AccountBalanceUpdate messages are sent by the Server,
// the RequestID will be 0.
type AccountBalanceRequestFixedPointerBuilder struct {
	message.FixedPointer
}

func NewAccountBalanceRequestFrom(b []byte) AccountBalanceRequest {
	return AccountBalanceRequest{VLS: message.NewVLSFrom(b)}
}

func WrapAccountBalanceRequest(b []byte) AccountBalanceRequest {
	return AccountBalanceRequest{VLS: message.WrapVLS(b)}
}

func NewAccountBalanceRequest() AccountBalanceRequestBuilder {
	a := AccountBalanceRequestBuilder{message.NewVLSBuilder(16)}
	a.Unsafe().SetBytes(0, _AccountBalanceRequestDefault)
	return a
}

func NewAccountBalanceRequestFixedFrom(b []byte) AccountBalanceRequestFixed {
	return AccountBalanceRequestFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapAccountBalanceRequestFixed(b []byte) AccountBalanceRequestFixed {
	return AccountBalanceRequestFixed{Fixed: message.WrapFixed(b)}
}

func NewAccountBalanceRequestFixed() AccountBalanceRequestFixedBuilder {
	a := AccountBalanceRequestFixedBuilder{message.NewFixed(40)}
	a.Unsafe().SetBytes(0, _AccountBalanceRequestFixedDefault)
	return a
}

func AllocAccountBalanceRequest() AccountBalanceRequestPointerBuilder {
	a := AccountBalanceRequestPointerBuilder{message.AllocVLSBuilder(16)}
	a.Ptr.SetBytes(0, _AccountBalanceRequestDefault)
	return a
}

func AllocAccountBalanceRequestFrom(b []byte) AccountBalanceRequestPointer {
	return AccountBalanceRequestPointer{VLSPointer: message.AllocVLSFrom(b)}
}

func AllocAccountBalanceRequestFixed() AccountBalanceRequestFixedPointerBuilder {
	a := AccountBalanceRequestFixedPointerBuilder{message.AllocFixed(40)}
	a.Ptr.SetBytes(0, _AccountBalanceRequestFixedDefault)
	return a
}

func AllocAccountBalanceRequestFixedFrom(b []byte) AccountBalanceRequestFixedPointer {
	return AccountBalanceRequestFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size         = AccountBalanceRequestSize  (16)
//     Type         = ACCOUNT_BALANCE_REQUEST  (601)
//     BaseSize     = AccountBalanceRequestSize  (16)
//     RequestID    = 0
//     TradeAccount = ""
// }
func (m AccountBalanceRequestBuilder) Clear() {
	m.Unsafe().SetBytes(0, _AccountBalanceRequestDefault)
}

// Clear
// {
//     Size         = AccountBalanceRequestFixedSize  (40)
//     Type         = ACCOUNT_BALANCE_REQUEST  (601)
//     RequestID    = 0
//     TradeAccount = ""
// }
func (m AccountBalanceRequestFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _AccountBalanceRequestFixedDefault)
}

// Clear
// {
//     Size         = AccountBalanceRequestSize  (16)
//     Type         = ACCOUNT_BALANCE_REQUEST  (601)
//     BaseSize     = AccountBalanceRequestSize  (16)
//     RequestID    = 0
//     TradeAccount = ""
// }
func (m AccountBalanceRequestPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _AccountBalanceRequestDefault)
}

// Clear
// {
//     Size         = AccountBalanceRequestFixedSize  (40)
//     Type         = ACCOUNT_BALANCE_REQUEST  (601)
//     RequestID    = 0
//     TradeAccount = ""
// }
func (m AccountBalanceRequestFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _AccountBalanceRequestFixedDefault)
}

// ToBuilder
func (m AccountBalanceRequest) ToBuilder() AccountBalanceRequestBuilder {
	return AccountBalanceRequestBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m AccountBalanceRequestFixed) ToBuilder() AccountBalanceRequestFixedBuilder {
	return AccountBalanceRequestFixedBuilder{m.Fixed}
}

// ToBuilder
func (m AccountBalanceRequestPointer) ToBuilder() AccountBalanceRequestPointerBuilder {
	return AccountBalanceRequestPointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m AccountBalanceRequestFixedPointer) ToBuilder() AccountBalanceRequestFixedPointerBuilder {
	return AccountBalanceRequestFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m AccountBalanceRequestBuilder) Finish() AccountBalanceRequest {
	return AccountBalanceRequest{m.VLSBuilder.Finish()}
}

// Finish
func (m AccountBalanceRequestFixedBuilder) Finish() AccountBalanceRequestFixed {
	return AccountBalanceRequestFixed{m.Fixed}
}

// Finish
func (m *AccountBalanceRequestPointerBuilder) Finish() AccountBalanceRequestPointer {
	return AccountBalanceRequestPointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *AccountBalanceRequestFixedPointerBuilder) Finish() AccountBalanceRequestFixedPointer {
	return AccountBalanceRequestFixedPointer{m.FixedPointer}
}

// RequestID A unique request identifier for this request.
func (m AccountBalanceRequest) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID A unique request identifier for this request.
func (m AccountBalanceRequestBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID A unique request identifier for this request.
func (m AccountBalanceRequestPointer) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RequestID A unique request identifier for this request.
func (m AccountBalanceRequestPointerBuilder) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// TradeAccount This is an optional field. Leave this empty to request the Server to return
// Account Balance data for all Trade Accounts on the logged in Username.
// Otherwise, specify a particular Trade Account to request Account Balance
// data for.
func (m AccountBalanceRequest) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// TradeAccount This is an optional field. Leave this empty to request the Server to return
// Account Balance data for all Trade Accounts on the logged in Username.
// Otherwise, specify a particular Trade Account to request Account Balance
// data for.
func (m AccountBalanceRequestBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// TradeAccount This is an optional field. Leave this empty to request the Server to return
// Account Balance data for all Trade Accounts on the logged in Username.
// Otherwise, specify a particular Trade Account to request Account Balance
// data for.
func (m AccountBalanceRequestPointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// TradeAccount This is an optional field. Leave this empty to request the Server to return
// Account Balance data for all Trade Accounts on the logged in Username.
// Otherwise, specify a particular Trade Account to request Account Balance
// data for.
func (m AccountBalanceRequestPointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// RequestID A unique request identifier for this request.
func (m AccountBalanceRequestFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID A unique request identifier for this request.
func (m AccountBalanceRequestFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID A unique request identifier for this request.
func (m AccountBalanceRequestFixedPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID A unique request identifier for this request.
func (m AccountBalanceRequestFixedPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// TradeAccount This is an optional field. Leave this empty to request the Server to return
// Account Balance data for all Trade Accounts on the logged in Username.
// Otherwise, specify a particular Trade Account to request Account Balance
// data for.
func (m AccountBalanceRequestFixed) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 40, 8)
}

// TradeAccount This is an optional field. Leave this empty to request the Server to return
// Account Balance data for all Trade Accounts on the logged in Username.
// Otherwise, specify a particular Trade Account to request Account Balance
// data for.
func (m AccountBalanceRequestFixedBuilder) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 40, 8)
}

// TradeAccount This is an optional field. Leave this empty to request the Server to return
// Account Balance data for all Trade Accounts on the logged in Username.
// Otherwise, specify a particular Trade Account to request Account Balance
// data for.
func (m AccountBalanceRequestFixedPointer) TradeAccount() string {
	return message.StringFixed(m.Ptr, 40, 8)
}

// TradeAccount This is an optional field. Leave this empty to request the Server to return
// Account Balance data for all Trade Accounts on the logged in Username.
// Otherwise, specify a particular Trade Account to request Account Balance
// data for.
func (m AccountBalanceRequestFixedPointerBuilder) TradeAccount() string {
	return message.StringFixed(m.Ptr, 40, 8)
}

// SetRequestID A unique request identifier for this request.
func (m AccountBalanceRequestBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetRequestID A unique request identifier for this request.
func (m AccountBalanceRequestPointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetTradeAccount This is an optional field. Leave this empty to request the Server to return
// Account Balance data for all Trade Accounts on the logged in Username.
// Otherwise, specify a particular Trade Account to request Account Balance
// data for.
func (m *AccountBalanceRequestBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// SetTradeAccount This is an optional field. Leave this empty to request the Server to return
// Account Balance data for all Trade Accounts on the logged in Username.
// Otherwise, specify a particular Trade Account to request Account Balance
// data for.
func (m *AccountBalanceRequestPointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// SetRequestID A unique request identifier for this request.
func (m AccountBalanceRequestFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRequestID A unique request identifier for this request.
func (m AccountBalanceRequestFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetTradeAccount This is an optional field. Leave this empty to request the Server to return
// Account Balance data for all Trade Accounts on the logged in Username.
// Otherwise, specify a particular Trade Account to request Account Balance
// data for.
func (m AccountBalanceRequestFixedBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Unsafe(), 40, 8, value)
}

// SetTradeAccount This is an optional field. Leave this empty to request the Server to return
// Account Balance data for all Trade Accounts on the logged in Username.
// Otherwise, specify a particular Trade Account to request Account Balance
// data for.
func (m AccountBalanceRequestFixedPointerBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Ptr, 40, 8, value)
}

// Copy
func (m AccountBalanceRequest) Copy(to AccountBalanceRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// Copy
func (m AccountBalanceRequestBuilder) Copy(to AccountBalanceRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m AccountBalanceRequest) CopyPointer(to AccountBalanceRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m AccountBalanceRequestBuilder) CopyPointer(to AccountBalanceRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyToPointer
func (m AccountBalanceRequest) CopyToPointer(to AccountBalanceRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyToPointer
func (m AccountBalanceRequestBuilder) CopyToPointer(to AccountBalanceRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyTo
func (m AccountBalanceRequest) CopyTo(to AccountBalanceRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyTo
func (m AccountBalanceRequestBuilder) CopyTo(to AccountBalanceRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// Copy
func (m AccountBalanceRequestFixed) Copy(to AccountBalanceRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// Copy
func (m AccountBalanceRequestFixedBuilder) Copy(to AccountBalanceRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m AccountBalanceRequestFixed) CopyPointer(to AccountBalanceRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m AccountBalanceRequestFixedBuilder) CopyPointer(to AccountBalanceRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyToPointer
func (m AccountBalanceRequestFixed) CopyToPointer(to AccountBalanceRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyToPointer
func (m AccountBalanceRequestFixedBuilder) CopyToPointer(to AccountBalanceRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyTo
func (m AccountBalanceRequestFixed) CopyTo(to AccountBalanceRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyTo
func (m AccountBalanceRequestFixedBuilder) CopyTo(to AccountBalanceRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// Copy
func (m AccountBalanceRequestPointer) Copy(to AccountBalanceRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// Copy
func (m AccountBalanceRequestPointerBuilder) Copy(to AccountBalanceRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m AccountBalanceRequestPointer) CopyPointer(to AccountBalanceRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m AccountBalanceRequestPointerBuilder) CopyPointer(to AccountBalanceRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyTo
func (m AccountBalanceRequestPointer) CopyTo(to AccountBalanceRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyTo
func (m AccountBalanceRequestPointerBuilder) CopyTo(to AccountBalanceRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyToPointer
func (m AccountBalanceRequestPointer) CopyToPointer(to AccountBalanceRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyToPointer
func (m AccountBalanceRequestPointerBuilder) CopyToPointer(to AccountBalanceRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// Copy
func (m AccountBalanceRequestFixedPointer) Copy(to AccountBalanceRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// Copy
func (m AccountBalanceRequestFixedPointerBuilder) Copy(to AccountBalanceRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m AccountBalanceRequestFixedPointer) CopyPointer(to AccountBalanceRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m AccountBalanceRequestFixedPointerBuilder) CopyPointer(to AccountBalanceRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyTo
func (m AccountBalanceRequestFixedPointer) CopyTo(to AccountBalanceRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyTo
func (m AccountBalanceRequestFixedPointerBuilder) CopyTo(to AccountBalanceRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyToPointer
func (m AccountBalanceRequestFixedPointer) CopyToPointer(to AccountBalanceRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyToPointer
func (m AccountBalanceRequestFixedPointerBuilder) CopyToPointer(to AccountBalanceRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}