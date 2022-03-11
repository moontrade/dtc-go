package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size                       = CancelReplaceOrderSize  (80)
//     Type                       = CANCEL_REPLACE_ORDER  (204)
//     BaseSize                   = CancelReplaceOrderSize  (80)
//     ServerOrderID              = ""
//     ClientOrderID              = ""
//     Price1                     = 0
//     Price2                     = 0
//     Quantity                   = 0
//     Price1IsSet                = 1
//     Price2IsSet                = 1
//     Unused                     = 0
//     TimeInForce                = 0
//     GoodTillDateTime           = 0
//     UpdatePrice1OffsetToParent = 0
//     TradeAccount               = ""
//     Price1AsString             = ""
//     Price2AsString             = ""
// }
var _CancelReplaceOrderDefault = []byte{80, 0, 204, 0, 80, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const CancelReplaceOrderSize = 80

// CancelReplaceOrder This message is sent by the Client to the Server to cancel and replace
// an existing order. This is also known as an order modification.
//
// When the cancel and replace operation is completed, an OrderUpdate message
// is sent by the Server with the OrderUpdateReasonfield set to ORDER_CANCEL_REPLACE_COMPLETE.
// If the cancel and replace operation cannot be completed, an OrderUpdate
// message is sent by the Server with the OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED.
// message is sent by the Server with the OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED.
type CancelReplaceOrder struct {
	message.VLS
}

// CancelReplaceOrderBuilder This message is sent by the Client to the Server to cancel and replace
// an existing order. This is also known as an order modification.
//
// When the cancel and replace operation is completed, an OrderUpdate message
// is sent by the Server with the OrderUpdateReasonfield set to ORDER_CANCEL_REPLACE_COMPLETE.
// If the cancel and replace operation cannot be completed, an OrderUpdate
// message is sent by the Server with the OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED.
// message is sent by the Server with the OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED.
type CancelReplaceOrderBuilder struct {
	message.VLSBuilder
}

func NewCancelReplaceOrderFrom(b []byte) CancelReplaceOrder {
	return CancelReplaceOrder{VLS: message.NewVLSFrom(b)}
}

func WrapCancelReplaceOrder(b []byte) CancelReplaceOrder {
	return CancelReplaceOrder{VLS: message.WrapVLS(b)}
}

func NewCancelReplaceOrder() CancelReplaceOrderBuilder {
	a := CancelReplaceOrderBuilder{message.NewVLSBuilder(80)}
	a.Unsafe().SetBytes(0, _CancelReplaceOrderDefault)
	return a
}

// Clear
// {
//     Size                       = CancelReplaceOrderSize  (80)
//     Type                       = CANCEL_REPLACE_ORDER  (204)
//     BaseSize                   = CancelReplaceOrderSize  (80)
//     ServerOrderID              = ""
//     ClientOrderID              = ""
//     Price1                     = 0
//     Price2                     = 0
//     Quantity                   = 0
//     Price1IsSet                = 1
//     Price2IsSet                = 1
//     Unused                     = 0
//     TimeInForce                = 0
//     GoodTillDateTime           = 0
//     UpdatePrice1OffsetToParent = 0
//     TradeAccount               = ""
//     Price1AsString             = ""
//     Price2AsString             = ""
// }
func (m CancelReplaceOrderBuilder) Clear() {
	m.Unsafe().SetBytes(0, _CancelReplaceOrderDefault)
}

// ToBuilder
func (m CancelReplaceOrder) ToBuilder() CancelReplaceOrderBuilder {
	return CancelReplaceOrderBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m CancelReplaceOrderBuilder) Finish() CancelReplaceOrder {
	return CancelReplaceOrder{m.VLSBuilder.Finish()}
}

// ServerOrderID This is the order identifier for the order to modify. The Client needs
// to set this to the ServerOrderID field received back in the most recent
// OrderUpdate message for the order.
//
// The Server will rely upon this ServerOrderID and only this order identifier
// to identify the order to be canceled and replaced. Although the given
// ClientOrderID by the Client must not change.
func (m CancelReplaceOrder) ServerOrderID() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// ServerOrderID This is the order identifier for the order to modify. The Client needs
// to set this to the ServerOrderID field received back in the most recent
// OrderUpdate message for the order.
//
// The Server will rely upon this ServerOrderID and only this order identifier
// to identify the order to be canceled and replaced. Although the given
// ClientOrderID by the Client must not change.
func (m CancelReplaceOrderBuilder) ServerOrderID() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// ClientOrderID This is the Client's own order identifier for the order.
//
// This must be the same throughout the life of the order. If the Server
// sees that this order identifier has changed in relation to the ServerOrderID,
// then it should reject this message with a OrderUpdate message with the
// OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED
//
// In the case where the order modification cannot be performed because the
// ServerOrderID does not exist, the Server will send a OrderUpdate message
// with the OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED and ClientOrderID
// set to the given ClientOrderID in this message. ServerOrderID will be
// unset because an invalid server order identifier was given.
func (m CancelReplaceOrder) ClientOrderID() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// ClientOrderID This is the Client's own order identifier for the order.
//
// This must be the same throughout the life of the order. If the Server
// sees that this order identifier has changed in relation to the ServerOrderID,
// then it should reject this message with a OrderUpdate message with the
// OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED
//
// In the case where the order modification cannot be performed because the
// ServerOrderID does not exist, the Server will send a OrderUpdate message
// with the OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED and ClientOrderID
// set to the given ClientOrderID in this message. ServerOrderID will be
// unset because an invalid server order identifier was given.
func (m CancelReplaceOrderBuilder) ClientOrderID() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// Price1 For orders that require a price, this is the new order price.
//
// This value can be left unset indicating to the Server that Price1 must
// not change and only the Quantity. In this case it is necessary to set
// Price1IsSet to a 0 value.
func (m CancelReplaceOrder) Price1() float64 {
	return message.Float64VLS(m.Unsafe(), 24, 16)
}

// Price1 For orders that require a price, this is the new order price.
//
// This value can be left unset indicating to the Server that Price1 must
// not change and only the Quantity. In this case it is necessary to set
// Price1IsSet to a 0 value.
func (m CancelReplaceOrderBuilder) Price1() float64 {
	return message.Float64VLS(m.Unsafe(), 24, 16)
}

// Price2 For Stop-Limit orders this is the new Limit price. For other order types
// it is not used.
//
// This value can be left unset indicating to the Server that Price2 must
// not change and only the Quantity. In this case it is necessary to set
// Price2IsSet to a 0 value.
func (m CancelReplaceOrder) Price2() float64 {
	return message.Float64VLS(m.Unsafe(), 32, 24)
}

// Price2 For Stop-Limit orders this is the new Limit price. For other order types
// it is not used.
//
// This value can be left unset indicating to the Server that Price2 must
// not change and only the Quantity. In this case it is necessary to set
// Price2IsSet to a 0 value.
func (m CancelReplaceOrderBuilder) Price2() float64 {
	return message.Float64VLS(m.Unsafe(), 32, 24)
}

// Quantity This is the new order quantity. It this is 0, then this means the order
// quantity must not be changed by the Server.
//
// If the order has partially filled, then this is going to be the order
// quantity which also includes the amount which has partially filled.
//
// For example, if the original quantity was 10 and there has been a partial
// fill of 3, the Client wants a fill of 2 more making a total of 5, then
// the Client will set this to 5.
func (m CancelReplaceOrder) Quantity() float64 {
	return message.Float64VLS(m.Unsafe(), 40, 32)
}

// Quantity This is the new order quantity. It this is 0, then this means the order
// quantity must not be changed by the Server.
//
// If the order has partially filled, then this is going to be the order
// quantity which also includes the amount which has partially filled.
//
// For example, if the original quantity was 10 and there has been a partial
// fill of 3, the Client wants a fill of 2 more making a total of 5, then
// the Client will set this to 5.
func (m CancelReplaceOrderBuilder) Quantity() float64 {
	return message.Float64VLS(m.Unsafe(), 40, 32)
}

// Price1IsSet When this field is set to a nonzero value, it indicates that Price1 is
// set and the server should use the value, if it applies to the order type.
// set and the server should use the value, if it applies to the order type.
//
// The default value is 1.
func (m CancelReplaceOrder) Price1IsSet() uint8 {
	return message.Uint8VLS(m.Unsafe(), 41, 40)
}

// Price1IsSet When this field is set to a nonzero value, it indicates that Price1 is
// set and the server should use the value, if it applies to the order type.
// set and the server should use the value, if it applies to the order type.
//
// The default value is 1.
func (m CancelReplaceOrderBuilder) Price1IsSet() uint8 {
	return message.Uint8VLS(m.Unsafe(), 41, 40)
}

// Price2IsSet When this field is set to a nonzero value, it indicates that Price2 is
// set and the server should use the value, if it applies to the order type.
// set and the server should use the value, if it applies to the order type.
//
// The default value is 1.
func (m CancelReplaceOrder) Price2IsSet() uint8 {
	return message.Uint8VLS(m.Unsafe(), 42, 41)
}

// Price2IsSet When this field is set to a nonzero value, it indicates that Price2 is
// set and the server should use the value, if it applies to the order type.
// set and the server should use the value, if it applies to the order type.
//
// The default value is 1.
func (m CancelReplaceOrderBuilder) Price2IsSet() uint8 {
	return message.Uint8VLS(m.Unsafe(), 42, 41)
}

// Unused
func (m CancelReplaceOrder) Unused() int32 {
	return message.Int32VLS(m.Unsafe(), 48, 44)
}

// Unused
func (m CancelReplaceOrderBuilder) Unused() int32 {
	return message.Int32VLS(m.Unsafe(), 48, 44)
}

// TimeInForce The Time in Force for the order. For a list of Time in Force values, refer
// to TimeInForceEnum.
//
// The default value is TIF_UNSET.
//
// When this field is set to a value other than TIF_UNSET, it indicates that
// the TimeInForce is being changed.
//
// If the server does not support changing the Time in Force of the order,
// it needs to reject this CancelReplaceOrder message and send an OrderUpdate
// message with the OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED.
// message with the OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED.
//
// The server is under no obligation to support changing the Time in Force.
// The server is under no obligation to support changing the Time in Force.
func (m CancelReplaceOrder) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(message.Int32VLS(m.Unsafe(), 52, 48))
}

// TimeInForce The Time in Force for the order. For a list of Time in Force values, refer
// to TimeInForceEnum.
//
// The default value is TIF_UNSET.
//
// When this field is set to a value other than TIF_UNSET, it indicates that
// the TimeInForce is being changed.
//
// If the server does not support changing the Time in Force of the order,
// it needs to reject this CancelReplaceOrder message and send an OrderUpdate
// message with the OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED.
// message with the OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED.
//
// The server is under no obligation to support changing the Time in Force.
// The server is under no obligation to support changing the Time in Force.
func (m CancelReplaceOrderBuilder) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(message.Int32VLS(m.Unsafe(), 52, 48))
}

// GoodTillDateTime In the case of when the TimeInForce field is TIF_GOOD_TILL_DATE_TIME,
// this specifies the expiration Date-Time of the order.
func (m CancelReplaceOrder) GoodTillDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Unsafe(), 64, 56))
}

// GoodTillDateTime In the case of when the TimeInForce field is TIF_GOOD_TILL_DATE_TIME,
// this specifies the expiration Date-Time of the order.
func (m CancelReplaceOrderBuilder) GoodTillDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Unsafe(), 64, 56))
}

// UpdatePrice1OffsetToParent This is an optional field. If modifying a child order which is part of
// a Server managed bracket order, then when this variable is set to 1 it
// provides an indication to the Server to update the internal server managed
// price offset to the parent order that this child order has to the parent.
// price offset to the parent order that this child order has to the parent.
//
// This will ensure the Server will maintain the proper offset of the child
// order to the fill price of the parent order when the parent order fills.
// order to the fill price of the parent order when the parent order fills.
func (m CancelReplaceOrder) UpdatePrice1OffsetToParent() uint8 {
	return message.Uint8VLS(m.Unsafe(), 65, 64)
}

// UpdatePrice1OffsetToParent This is an optional field. If modifying a child order which is part of
// a Server managed bracket order, then when this variable is set to 1 it
// provides an indication to the Server to update the internal server managed
// price offset to the parent order that this child order has to the parent.
// price offset to the parent order that this child order has to the parent.
//
// This will ensure the Server will maintain the proper offset of the child
// order to the fill price of the parent order when the parent order fills.
// order to the fill price of the parent order when the parent order fills.
func (m CancelReplaceOrderBuilder) UpdatePrice1OffsetToParent() uint8 {
	return message.Uint8VLS(m.Unsafe(), 65, 64)
}

// TradeAccount
func (m CancelReplaceOrder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 70, 66)
}

// TradeAccount
func (m CancelReplaceOrderBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 70, 66)
}

// Price1AsString
func (m CancelReplaceOrder) Price1AsString() string {
	return message.StringVLS(m.Unsafe(), 74, 70)
}

// Price1AsString
func (m CancelReplaceOrderBuilder) Price1AsString() string {
	return message.StringVLS(m.Unsafe(), 74, 70)
}

// Price2AsString
func (m CancelReplaceOrder) Price2AsString() string {
	return message.StringVLS(m.Unsafe(), 78, 74)
}

// Price2AsString
func (m CancelReplaceOrderBuilder) Price2AsString() string {
	return message.StringVLS(m.Unsafe(), 78, 74)
}

// SetServerOrderID This is the order identifier for the order to modify. The Client needs
// to set this to the ServerOrderID field received back in the most recent
// OrderUpdate message for the order.
//
// The Server will rely upon this ServerOrderID and only this order identifier
// to identify the order to be canceled and replaced. Although the given
// ClientOrderID by the Client must not change.
func (m *CancelReplaceOrderBuilder) SetServerOrderID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 10, 6, value)
}

// SetClientOrderID This is the Client's own order identifier for the order.
//
// This must be the same throughout the life of the order. If the Server
// sees that this order identifier has changed in relation to the ServerOrderID,
// then it should reject this message with a OrderUpdate message with the
// OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED
//
// In the case where the order modification cannot be performed because the
// ServerOrderID does not exist, the Server will send a OrderUpdate message
// with the OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED and ClientOrderID
// set to the given ClientOrderID in this message. ServerOrderID will be
// unset because an invalid server order identifier was given.
func (m *CancelReplaceOrderBuilder) SetClientOrderID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 14, 10, value)
}

// SetPrice1 For orders that require a price, this is the new order price.
//
// This value can be left unset indicating to the Server that Price1 must
// not change and only the Quantity. In this case it is necessary to set
// Price1IsSet to a 0 value.
func (m CancelReplaceOrderBuilder) SetPrice1(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 24, 16, value)
}

// SetPrice2 For Stop-Limit orders this is the new Limit price. For other order types
// it is not used.
//
// This value can be left unset indicating to the Server that Price2 must
// not change and only the Quantity. In this case it is necessary to set
// Price2IsSet to a 0 value.
func (m CancelReplaceOrderBuilder) SetPrice2(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 32, 24, value)
}

// SetQuantity This is the new order quantity. It this is 0, then this means the order
// quantity must not be changed by the Server.
//
// If the order has partially filled, then this is going to be the order
// quantity which also includes the amount which has partially filled.
//
// For example, if the original quantity was 10 and there has been a partial
// fill of 3, the Client wants a fill of 2 more making a total of 5, then
// the Client will set this to 5.
func (m CancelReplaceOrderBuilder) SetQuantity(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 40, 32, value)
}

// SetPrice1IsSet When this field is set to a nonzero value, it indicates that Price1 is
// set and the server should use the value, if it applies to the order type.
// set and the server should use the value, if it applies to the order type.
//
// The default value is 1.
func (m CancelReplaceOrderBuilder) SetPrice1IsSet(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 41, 40, value)
}

// SetPrice2IsSet When this field is set to a nonzero value, it indicates that Price2 is
// set and the server should use the value, if it applies to the order type.
// set and the server should use the value, if it applies to the order type.
//
// The default value is 1.
func (m CancelReplaceOrderBuilder) SetPrice2IsSet(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 42, 41, value)
}

// SetUnused
func (m CancelReplaceOrderBuilder) SetUnused(value int32) {
	message.SetInt32VLS(m.Unsafe(), 48, 44, value)
}

// SetTimeInForce The Time in Force for the order. For a list of Time in Force values, refer
// to TimeInForceEnum.
//
// The default value is TIF_UNSET.
//
// When this field is set to a value other than TIF_UNSET, it indicates that
// the TimeInForce is being changed.
//
// If the server does not support changing the Time in Force of the order,
// it needs to reject this CancelReplaceOrder message and send an OrderUpdate
// message with the OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED.
// message with the OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED.
//
// The server is under no obligation to support changing the Time in Force.
// The server is under no obligation to support changing the Time in Force.
func (m CancelReplaceOrderBuilder) SetTimeInForce(value TimeInForceEnum) {
	message.SetInt32VLS(m.Unsafe(), 52, 48, int32(value))
}

// SetGoodTillDateTime In the case of when the TimeInForce field is TIF_GOOD_TILL_DATE_TIME,
// this specifies the expiration Date-Time of the order.
func (m CancelReplaceOrderBuilder) SetGoodTillDateTime(value DateTime) {
	message.SetInt64VLS(m.Unsafe(), 64, 56, int64(value))
}

// SetUpdatePrice1OffsetToParent This is an optional field. If modifying a child order which is part of
// a Server managed bracket order, then when this variable is set to 1 it
// provides an indication to the Server to update the internal server managed
// price offset to the parent order that this child order has to the parent.
// price offset to the parent order that this child order has to the parent.
//
// This will ensure the Server will maintain the proper offset of the child
// order to the fill price of the parent order when the parent order fills.
// order to the fill price of the parent order when the parent order fills.
func (m CancelReplaceOrderBuilder) SetUpdatePrice1OffsetToParent(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 65, 64, value)
}

// SetTradeAccount
func (m *CancelReplaceOrderBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 70, 66, value)
}

// SetPrice1AsString
func (m *CancelReplaceOrderBuilder) SetPrice1AsString(value string) {
	message.SetStringVLS(&m.VLSBuilder, 74, 70, value)
}

// SetPrice2AsString
func (m *CancelReplaceOrderBuilder) SetPrice2AsString(value string) {
	message.SetStringVLS(&m.VLSBuilder, 78, 74, value)
}

// Copy
func (m CancelReplaceOrder) Copy(to CancelReplaceOrderBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetQuantity(m.Quantity())
	to.SetPrice1IsSet(m.Price1IsSet())
	to.SetPrice2IsSet(m.Price2IsSet())
	to.SetUnused(m.Unused())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetUpdatePrice1OffsetToParent(m.UpdatePrice1OffsetToParent())
	to.SetTradeAccount(m.TradeAccount())
	to.SetPrice1AsString(m.Price1AsString())
	to.SetPrice2AsString(m.Price2AsString())
}

// Copy
func (m CancelReplaceOrderBuilder) Copy(to CancelReplaceOrderBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetQuantity(m.Quantity())
	to.SetPrice1IsSet(m.Price1IsSet())
	to.SetPrice2IsSet(m.Price2IsSet())
	to.SetUnused(m.Unused())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetUpdatePrice1OffsetToParent(m.UpdatePrice1OffsetToParent())
	to.SetTradeAccount(m.TradeAccount())
	to.SetPrice1AsString(m.Price1AsString())
	to.SetPrice2AsString(m.Price2AsString())
}

// CopyPointer
func (m CancelReplaceOrder) CopyPointer(to CancelReplaceOrderPointerBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetQuantity(m.Quantity())
	to.SetPrice1IsSet(m.Price1IsSet())
	to.SetPrice2IsSet(m.Price2IsSet())
	to.SetUnused(m.Unused())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetUpdatePrice1OffsetToParent(m.UpdatePrice1OffsetToParent())
	to.SetTradeAccount(m.TradeAccount())
	to.SetPrice1AsString(m.Price1AsString())
	to.SetPrice2AsString(m.Price2AsString())
}

// CopyPointer
func (m CancelReplaceOrderBuilder) CopyPointer(to CancelReplaceOrderPointerBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetQuantity(m.Quantity())
	to.SetPrice1IsSet(m.Price1IsSet())
	to.SetPrice2IsSet(m.Price2IsSet())
	to.SetUnused(m.Unused())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetUpdatePrice1OffsetToParent(m.UpdatePrice1OffsetToParent())
	to.SetTradeAccount(m.TradeAccount())
	to.SetPrice1AsString(m.Price1AsString())
	to.SetPrice2AsString(m.Price2AsString())
}

// CopyToPointer
func (m CancelReplaceOrder) CopyToPointer(to CancelReplaceOrderFixedPointerBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetQuantity(m.Quantity())
	to.SetPrice1IsSet(m.Price1IsSet())
	to.SetPrice2IsSet(m.Price2IsSet())
	to.SetUnused(m.Unused())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetUpdatePrice1OffsetToParent(m.UpdatePrice1OffsetToParent())
	to.SetTradeAccount(m.TradeAccount())
	to.SetPrice1AsString(m.Price1AsString())
	to.SetPrice2AsString(m.Price2AsString())
}

// CopyToPointer
func (m CancelReplaceOrderBuilder) CopyToPointer(to CancelReplaceOrderFixedPointerBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetQuantity(m.Quantity())
	to.SetPrice1IsSet(m.Price1IsSet())
	to.SetPrice2IsSet(m.Price2IsSet())
	to.SetUnused(m.Unused())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetUpdatePrice1OffsetToParent(m.UpdatePrice1OffsetToParent())
	to.SetTradeAccount(m.TradeAccount())
	to.SetPrice1AsString(m.Price1AsString())
	to.SetPrice2AsString(m.Price2AsString())
}

// CopyTo
func (m CancelReplaceOrder) CopyTo(to CancelReplaceOrderFixedBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetQuantity(m.Quantity())
	to.SetPrice1IsSet(m.Price1IsSet())
	to.SetPrice2IsSet(m.Price2IsSet())
	to.SetUnused(m.Unused())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetUpdatePrice1OffsetToParent(m.UpdatePrice1OffsetToParent())
	to.SetTradeAccount(m.TradeAccount())
	to.SetPrice1AsString(m.Price1AsString())
	to.SetPrice2AsString(m.Price2AsString())
}

// CopyTo
func (m CancelReplaceOrderBuilder) CopyTo(to CancelReplaceOrderFixedBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetQuantity(m.Quantity())
	to.SetPrice1IsSet(m.Price1IsSet())
	to.SetPrice2IsSet(m.Price2IsSet())
	to.SetUnused(m.Unused())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetUpdatePrice1OffsetToParent(m.UpdatePrice1OffsetToParent())
	to.SetTradeAccount(m.TradeAccount())
	to.SetPrice1AsString(m.Price1AsString())
	to.SetPrice2AsString(m.Price2AsString())
}
