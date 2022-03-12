// generated by github.com/moontrade/dtc-go/codegen/go at 2022-03-12 12:55:11.079124 +0800 WITA m=+0.026505376

package sc

import (
	"github.com/moontrade/dtc-go/message"
)

const TradeAccountDataUsernameToShareWithAddSize = 19

// {
//     Size         = TradeAccountDataUsernameToShareWithAddSize  (19)
//     Type         = TRADE_ACCOUNT_DATA_USERNAME_TO_SHARE_WITH_ADD  (10128)
//     BaseSize     = TradeAccountDataUsernameToShareWithAddSize  (19)
//     RequestID    = 0
//     TradeAccount = ""
//     Username     = ""
//     IsReadOnly   = false
// }
var _TradeAccountDataUsernameToShareWithAddDefault = []byte{19, 0, 144, 39, 19, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type TradeAccountDataUsernameToShareWithAdd struct {
	message.VLS
}

type TradeAccountDataUsernameToShareWithAddBuilder struct {
	message.VLSBuilder
}

type TradeAccountDataUsernameToShareWithAddPointer struct {
	message.VLSPointer
}

type TradeAccountDataUsernameToShareWithAddPointerBuilder struct {
	message.VLSPointerBuilder
}

func NewTradeAccountDataUsernameToShareWithAddFrom(b []byte) TradeAccountDataUsernameToShareWithAdd {
	return TradeAccountDataUsernameToShareWithAdd{VLS: message.NewVLSFrom(b)}
}

func WrapTradeAccountDataUsernameToShareWithAdd(b []byte) TradeAccountDataUsernameToShareWithAdd {
	return TradeAccountDataUsernameToShareWithAdd{VLS: message.WrapVLS(b)}
}

func NewTradeAccountDataUsernameToShareWithAdd() TradeAccountDataUsernameToShareWithAddBuilder {
	a := TradeAccountDataUsernameToShareWithAddBuilder{message.NewVLSBuilder(19)}
	a.Unsafe().SetBytes(0, _TradeAccountDataUsernameToShareWithAddDefault)
	return a
}

func AllocTradeAccountDataUsernameToShareWithAdd() TradeAccountDataUsernameToShareWithAddPointerBuilder {
	a := TradeAccountDataUsernameToShareWithAddPointerBuilder{message.AllocVLSBuilder(19)}
	a.Ptr.SetBytes(0, _TradeAccountDataUsernameToShareWithAddDefault)
	return a
}

func AllocTradeAccountDataUsernameToShareWithAddFrom(b []byte) TradeAccountDataUsernameToShareWithAddPointer {
	return TradeAccountDataUsernameToShareWithAddPointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size         = TradeAccountDataUsernameToShareWithAddSize  (19)
//     Type         = TRADE_ACCOUNT_DATA_USERNAME_TO_SHARE_WITH_ADD  (10128)
//     BaseSize     = TradeAccountDataUsernameToShareWithAddSize  (19)
//     RequestID    = 0
//     TradeAccount = ""
//     Username     = ""
//     IsReadOnly   = false
// }
func (m TradeAccountDataUsernameToShareWithAddBuilder) Clear() {
	m.Unsafe().SetBytes(0, _TradeAccountDataUsernameToShareWithAddDefault)
}

// Clear
// {
//     Size         = TradeAccountDataUsernameToShareWithAddSize  (19)
//     Type         = TRADE_ACCOUNT_DATA_USERNAME_TO_SHARE_WITH_ADD  (10128)
//     BaseSize     = TradeAccountDataUsernameToShareWithAddSize  (19)
//     RequestID    = 0
//     TradeAccount = ""
//     Username     = ""
//     IsReadOnly   = false
// }
func (m TradeAccountDataUsernameToShareWithAddPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _TradeAccountDataUsernameToShareWithAddDefault)
}

// ToBuilder
func (m TradeAccountDataUsernameToShareWithAdd) ToBuilder() TradeAccountDataUsernameToShareWithAddBuilder {
	return TradeAccountDataUsernameToShareWithAddBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m TradeAccountDataUsernameToShareWithAddPointer) ToBuilder() TradeAccountDataUsernameToShareWithAddPointerBuilder {
	return TradeAccountDataUsernameToShareWithAddPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m TradeAccountDataUsernameToShareWithAddBuilder) Finish() TradeAccountDataUsernameToShareWithAdd {
	return TradeAccountDataUsernameToShareWithAdd{m.VLSBuilder.Finish()}
}

// Finish
func (m *TradeAccountDataUsernameToShareWithAddPointerBuilder) Finish() TradeAccountDataUsernameToShareWithAddPointer {
	return TradeAccountDataUsernameToShareWithAddPointer{m.VLSPointerBuilder.Finish()}
}

// RequestID
func (m TradeAccountDataUsernameToShareWithAdd) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m TradeAccountDataUsernameToShareWithAddBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m TradeAccountDataUsernameToShareWithAddPointer) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// RequestID
func (m TradeAccountDataUsernameToShareWithAddPointerBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// TradeAccount
func (m TradeAccountDataUsernameToShareWithAdd) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// TradeAccount
func (m TradeAccountDataUsernameToShareWithAddBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// TradeAccount
func (m TradeAccountDataUsernameToShareWithAddPointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// TradeAccount
func (m TradeAccountDataUsernameToShareWithAddPointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// Username
func (m TradeAccountDataUsernameToShareWithAdd) Username() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// Username
func (m TradeAccountDataUsernameToShareWithAddBuilder) Username() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// Username
func (m TradeAccountDataUsernameToShareWithAddPointer) Username() string {
	return message.StringVLS(m.Ptr, 18, 14)
}

// Username
func (m TradeAccountDataUsernameToShareWithAddPointerBuilder) Username() string {
	return message.StringVLS(m.Ptr, 18, 14)
}

// IsReadOnly
func (m TradeAccountDataUsernameToShareWithAdd) IsReadOnly() bool {
	return message.BoolVLS(m.Unsafe(), 19, 18)
}

// IsReadOnly
func (m TradeAccountDataUsernameToShareWithAddBuilder) IsReadOnly() bool {
	return message.BoolVLS(m.Unsafe(), 19, 18)
}

// IsReadOnly
func (m TradeAccountDataUsernameToShareWithAddPointer) IsReadOnly() bool {
	return message.BoolVLS(m.Ptr, 19, 18)
}

// IsReadOnly
func (m TradeAccountDataUsernameToShareWithAddPointerBuilder) IsReadOnly() bool {
	return message.BoolVLS(m.Ptr, 19, 18)
}

// SetRequestID
func (m TradeAccountDataUsernameToShareWithAddBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 10, 6, value)
}

// SetRequestID
func (m TradeAccountDataUsernameToShareWithAddPointerBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Ptr, 10, 6, value)
}

// SetTradeAccount
func (m *TradeAccountDataUsernameToShareWithAddBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 14, 10, value)
}

// SetTradeAccount
func (m *TradeAccountDataUsernameToShareWithAddPointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 14, 10, value)
}

// SetUsername
func (m *TradeAccountDataUsernameToShareWithAddBuilder) SetUsername(value string) {
	message.SetStringVLS(&m.VLSBuilder, 18, 14, value)
}

// SetUsername
func (m *TradeAccountDataUsernameToShareWithAddPointerBuilder) SetUsername(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 18, 14, value)
}

// SetIsReadOnly
func (m TradeAccountDataUsernameToShareWithAddBuilder) SetIsReadOnly(value bool) {
	message.SetBoolVLS(m.Unsafe(), 19, 18, value)
}

// SetIsReadOnly
func (m TradeAccountDataUsernameToShareWithAddPointerBuilder) SetIsReadOnly(value bool) {
	message.SetBoolVLS(m.Ptr, 19, 18, value)
}

// Copy
func (m TradeAccountDataUsernameToShareWithAdd) Copy(to TradeAccountDataUsernameToShareWithAddBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
	to.SetIsReadOnly(m.IsReadOnly())
}

// Copy
func (m TradeAccountDataUsernameToShareWithAddBuilder) Copy(to TradeAccountDataUsernameToShareWithAddBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
	to.SetIsReadOnly(m.IsReadOnly())
}

// CopyPointer
func (m TradeAccountDataUsernameToShareWithAdd) CopyPointer(to TradeAccountDataUsernameToShareWithAddPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
	to.SetIsReadOnly(m.IsReadOnly())
}

// CopyPointer
func (m TradeAccountDataUsernameToShareWithAddBuilder) CopyPointer(to TradeAccountDataUsernameToShareWithAddPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
	to.SetIsReadOnly(m.IsReadOnly())
}

// Copy
func (m TradeAccountDataUsernameToShareWithAddPointer) Copy(to TradeAccountDataUsernameToShareWithAddBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
	to.SetIsReadOnly(m.IsReadOnly())
}

// Copy
func (m TradeAccountDataUsernameToShareWithAddPointerBuilder) Copy(to TradeAccountDataUsernameToShareWithAddBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
	to.SetIsReadOnly(m.IsReadOnly())
}

// CopyPointer
func (m TradeAccountDataUsernameToShareWithAddPointer) CopyPointer(to TradeAccountDataUsernameToShareWithAddPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
	to.SetIsReadOnly(m.IsReadOnly())
}

// CopyPointer
func (m TradeAccountDataUsernameToShareWithAddPointerBuilder) CopyPointer(to TradeAccountDataUsernameToShareWithAddPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
	to.SetIsReadOnly(m.IsReadOnly())
}
