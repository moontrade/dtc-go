package model

import (
	"github.com/moontrade/dtc-go/message"
)

// OrderUpdateFixedPointer The OrderUpdate is a unified message from the Server to the Client which
// communicates the complete details of an order, the Order Status, and the
// reason for sending the message (OrderUpdateReason).
//
// DTC uses this single unified message to provide an update for an order.
// The OrderUpdateReason field provides a clear indication for each reason
// this message is being sent.
type OrderUpdateFixedPointer struct {
	message.FixedPointer
}

// OrderUpdateFixedPointerBuilder The OrderUpdate is a unified message from the Server to the Client which
// communicates the complete details of an order, the Order Status, and the
// reason for sending the message (OrderUpdateReason).
//
// DTC uses this single unified message to provide an update for an order.
// The OrderUpdateReason field provides a clear indication for each reason
// this message is being sent.
type OrderUpdateFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocOrderUpdateFixed() OrderUpdateFixedPointerBuilder {
	a := OrderUpdateFixedPointerBuilder{message.AllocFixed(688)}
	a.Ptr.SetBytes(0, _OrderUpdateFixedDefault)
	return a
}

func AllocOrderUpdateFixedFrom(b []byte) OrderUpdateFixedPointer {
	return OrderUpdateFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size                        = OrderUpdateFixedSize  (688)
//     Type                        = ORDER_UPDATE  (301)
//     RequestID                   = 0
//     TotalNumMessages            = 0
//     MessageNumber               = 0
//     Symbol                      = ""
//     Exchange                    = ""
//     PreviousServerOrderID       = ""
//     ServerOrderID               = ""
//     ClientOrderID               = ""
//     ExchangeOrderID             = ""
//     OrderStatus                 = 0
//     OrderUpdateReason           = 0
//     OrderType                   = 0
//     BuySell                     = 0
//     Price1                      = math.MaxFloat64
//     Price2                      = math.MaxFloat64
//     TimeInForce                 = 0
//     GoodTillDateTime            = 0
//     OrderQuantity               = math.MaxFloat64
//     FilledQuantity              = math.MaxFloat64
//     RemainingQuantity           = math.MaxFloat64
//     AverageFillPrice            = math.MaxFloat64
//     LastFillPrice               = math.MaxFloat64
//     LastFillDateTime            = 0
//     LastFillQuantity            = math.MaxFloat64
//     LastFillExecutionID         = ""
//     TradeAccount                = ""
//     InfoText                    = ""
//     NoOrders                    = 0
//     ParentServerOrderID         = ""
//     OCOLinkedOrderServerOrderID = ""
//     OpenOrClose                 = 0
//     PreviousClientOrderID       = ""
//     FreeFormText                = ""
//     OrderReceivedDateTime       = 0
//     LatestTransactionDateTime   = 0
// }
func (m OrderUpdateFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _OrderUpdateFixedDefault)
}

// ToBuilder
func (m OrderUpdateFixedPointer) ToBuilder() OrderUpdateFixedPointerBuilder {
	return OrderUpdateFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *OrderUpdateFixedPointerBuilder) Finish() OrderUpdateFixedPointer {
	return OrderUpdateFixedPointer{m.FixedPointer}
}

// RequestID Set to 0 unless this is in response to an OpenOrdersRequest, in which
// case this must be set to the RequestID given in the OpenOrdersRequest.
//
// If this OrderUpdate is unsolicited, for example a real-time fill or other
// unsolicited order event, the Server must leave this at 0.
func (m OrderUpdateFixedPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID Set to 0 unless this is in response to an OpenOrdersRequest, in which
// case this must be set to the RequestID given in the OpenOrdersRequest.
//
// If this OrderUpdate is unsolicited, for example a real-time fill or other
// unsolicited order event, the Server must leave this at 0.
func (m OrderUpdateFixedPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// TotalNumMessages This indicates the total number of OrderUpdate messages when a batch of
// reports is being sent in response to an OpenOrdersRequest. If there is
// only one order being sent, this will be 1. The Server must use a value
// of 1 for an unsolicited report. A Client should not rely on this field
// for an unsolicited report.
func (m OrderUpdateFixedPointer) TotalNumMessages() int32 {
	return message.Int32Fixed(m.Ptr, 12, 8)
}

// TotalNumMessages This indicates the total number of OrderUpdate messages when a batch of
// reports is being sent in response to an OpenOrdersRequest. If there is
// only one order being sent, this will be 1. The Server must use a value
// of 1 for an unsolicited report. A Client should not rely on this field
// for an unsolicited report.
func (m OrderUpdateFixedPointerBuilder) TotalNumMessages() int32 {
	return message.Int32Fixed(m.Ptr, 12, 8)
}

// MessageNumber This indicates the 1-based index of the OrderUpdate message when a batch
// of reports is being sent in response to an OpenOrdersRequest. If there
// is only one order being sent, this will be 1. Use a value of 1 for an
// unsolicited report. A Client should not rely on this field for an unsolicited
// report.
func (m OrderUpdateFixedPointer) MessageNumber() int32 {
	return message.Int32Fixed(m.Ptr, 16, 12)
}

// MessageNumber This indicates the 1-based index of the OrderUpdate message when a batch
// of reports is being sent in response to an OpenOrdersRequest. If there
// is only one order being sent, this will be 1. Use a value of 1 for an
// unsolicited report. A Client should not rely on this field for an unsolicited
// report.
func (m OrderUpdateFixedPointerBuilder) MessageNumber() int32 {
	return message.Int32Fixed(m.Ptr, 16, 12)
}

// Symbol The symbol for the order.
func (m OrderUpdateFixedPointer) Symbol() string {
	return message.StringFixed(m.Ptr, 80, 16)
}

// Symbol The symbol for the order.
func (m OrderUpdateFixedPointerBuilder) Symbol() string {
	return message.StringFixed(m.Ptr, 80, 16)
}

// Exchange The optional exchange for the symbol.
func (m OrderUpdateFixedPointer) Exchange() string {
	return message.StringFixed(m.Ptr, 96, 80)
}

// Exchange The optional exchange for the symbol.
func (m OrderUpdateFixedPointerBuilder) Exchange() string {
	return message.StringFixed(m.Ptr, 96, 80)
}

// PreviousServerOrderID Used upon a Cancel and Replace operation (ORDER_CANCEL_REPLACE_COMPLETE)
// where a new Server Order identifier is given.
//
// In this case this field needs to be set to the previous Server Order identifier.
// In this case this field needs to be set to the previous Server Order identifier.
//
// This should be left at the default setting of empty in the case where
// the Server does not change the Server Order identifier upon a Cancel and
// Replace operation.
func (m OrderUpdateFixedPointer) PreviousServerOrderID() string {
	return message.StringFixed(m.Ptr, 128, 96)
}

// PreviousServerOrderID Used upon a Cancel and Replace operation (ORDER_CANCEL_REPLACE_COMPLETE)
// where a new Server Order identifier is given.
//
// In this case this field needs to be set to the previous Server Order identifier.
// In this case this field needs to be set to the previous Server Order identifier.
//
// This should be left at the default setting of empty in the case where
// the Server does not change the Server Order identifier upon a Cancel and
// Replace operation.
func (m OrderUpdateFixedPointerBuilder) PreviousServerOrderID() string {
	return message.StringFixed(m.Ptr, 128, 96)
}

// ServerOrderID The ServerOrderID is set by the server and uniquely identifies the order.
// When a new order is submitted by the Client and the Server responds with
// an OrderUpdate, this field needs to be set to the order identifier which
// is good for the life of the order.
//
// This ServerOrderID can optionally change on a Cancel and Replace operation.
// In this case, the PreviousServerOrderID field will contain the previous
// ServerOrderID upon a ORDER_CANCEL_REPLACE_COMPLETE OrderUpdateReason.
//
// ServerOrderID must always be set except it is not required in the cases
// when the OrderUpdateReason is one of the following: NEW_ORDER_REJECTED,
// ORDER_CANCEL_REJECTED, ORDER_CANCEL_REPLACE_REJECTED. If it is not set,
// then the ClientOrderID must be set.
func (m OrderUpdateFixedPointer) ServerOrderID() string {
	return message.StringFixed(m.Ptr, 160, 128)
}

// ServerOrderID The ServerOrderID is set by the server and uniquely identifies the order.
// When a new order is submitted by the Client and the Server responds with
// an OrderUpdate, this field needs to be set to the order identifier which
// is good for the life of the order.
//
// This ServerOrderID can optionally change on a Cancel and Replace operation.
// In this case, the PreviousServerOrderID field will contain the previous
// ServerOrderID upon a ORDER_CANCEL_REPLACE_COMPLETE OrderUpdateReason.
//
// ServerOrderID must always be set except it is not required in the cases
// when the OrderUpdateReason is one of the following: NEW_ORDER_REJECTED,
// ORDER_CANCEL_REJECTED, ORDER_CANCEL_REPLACE_REJECTED. If it is not set,
// then the ClientOrderID must be set.
func (m OrderUpdateFixedPointerBuilder) ServerOrderID() string {
	return message.StringFixed(m.Ptr, 160, 128)
}

// ClientOrderID The ClientOrderID is the order identifier provided by the Client. When
// the Client submits a new order, cancels and replaces an existing order,
// or cancels an order, then the Client needs to specify this identifier.
//
// The Client must maintain the same order identifier throughout the life
// of the order.
//
// The Server should persist the ClientOrderID across sessions. A session
// is defined as the period of time from the start of the network connection
// between the Client and Server to the end of that connection.
//
// The Client should only rely upon the ClientOrderID being set when the
// OrderUpdateReason is one of the following: NEW_ORDER_ACCEPTED, NEW_ORDER_REJECTED,
// ORDER_CANCELED, ORDER_CANCEL_REPLACE_COMPLETE, ORDER_CANCEL_REJECTED,
// ORDER_CANCEL_REPLACE_REJECTED.
//
// After a new order has been accepted, the Client will rely upon the given
// ServerOrderID from the server to identify the order and should no longer
// rely upon the given ClientOrderID. However, the Client needs to maintain
// a copy of the ClientOrderID for any subsequent order modifications and
// cancellations because this is a required field for those.
func (m OrderUpdateFixedPointer) ClientOrderID() string {
	return message.StringFixed(m.Ptr, 192, 160)
}

// ClientOrderID The ClientOrderID is the order identifier provided by the Client. When
// the Client submits a new order, cancels and replaces an existing order,
// or cancels an order, then the Client needs to specify this identifier.
//
// The Client must maintain the same order identifier throughout the life
// of the order.
//
// The Server should persist the ClientOrderID across sessions. A session
// is defined as the period of time from the start of the network connection
// between the Client and Server to the end of that connection.
//
// The Client should only rely upon the ClientOrderID being set when the
// OrderUpdateReason is one of the following: NEW_ORDER_ACCEPTED, NEW_ORDER_REJECTED,
// ORDER_CANCELED, ORDER_CANCEL_REPLACE_COMPLETE, ORDER_CANCEL_REJECTED,
// ORDER_CANCEL_REPLACE_REJECTED.
//
// After a new order has been accepted, the Client will rely upon the given
// ServerOrderID from the server to identify the order and should no longer
// rely upon the given ClientOrderID. However, the Client needs to maintain
// a copy of the ClientOrderID for any subsequent order modifications and
// cancellations because this is a required field for those.
func (m OrderUpdateFixedPointerBuilder) ClientOrderID() string {
	return message.StringFixed(m.Ptr, 192, 160)
}

// ExchangeOrderID The order identifier from the exchange that handles the order. This is
// optional.
func (m OrderUpdateFixedPointer) ExchangeOrderID() string {
	return message.StringFixed(m.Ptr, 224, 192)
}

// ExchangeOrderID The order identifier from the exchange that handles the order. This is
// optional.
func (m OrderUpdateFixedPointerBuilder) ExchangeOrderID() string {
	return message.StringFixed(m.Ptr, 224, 192)
}

// OrderStatus This is required. Needs to be set to one of the following values by the
// Server to indicate the current status of the order, unless NoneOrder =
// 1:
//
// ORDER_STATUS_UNSPECIFIED (0): The status of the order is unset. Use this
// when NoneOrder = 1, or when the Server does not know the true status of
// an order when the Order Update Report message is sent.
// ORDER_STATUS_ORDER_SENT (1): When a Client sends an order to the Server,
// then the Client internally will set the status to ORDER_STATUS_ORDER_SENT.
// The Server will not set this Status.
// ORDER_STATUS_PENDING_OPEN (2): This means the Server has accepted the
// order but it is not yet considered in a fully working state for any reason.
// order but it is not yet considered in a fully working state for any reason.
// ORDER_STATUS_PENDING_CHILD (3): This status applies to a Limit or Stop
// order attached to a parent order. It will have this status if the parent
// order has not yet filled.
// ORDER_STATUS_OPEN (4): Order is open and working.
// ORDER_STATUS_PENDING_CANCEL_REPLACE (5): Order is pending a Cancel and
// Replace operation. The Server should send a OrderUpdate message with the
// OrderUpdateReason set to GENERAL_OrderUpdate and the OrderStatus set to
// this status to indicate the pending Cancel and Replace operation.
// ORDER_STATUS_PENDING_CANCEL (6): Order is pending a Cancel operation.
// The Server should send a OrderUpdate message with the OrderUpdateReason
// set to GENERAL_OrderUpdate and the OrderStatus set to this status to indicate
// the pending Cancel operation.
// ORDER_STATUS_FILLED (7): Order is filled and no longer working.
// ORDER_STATUS_CANCELED (8): Order is canceled. If the user tries to cancel
// an order that has already been canceled, then continue to return this
// Order Status for it.
// ORDER_STATUS_REJECTED (9): Order has been rejected after the initial order
// submission. It is not working.
// ORDER_STATUS_PARTIALLY_FILLED (10): Order is partially filled and still
// working.
func (m OrderUpdateFixedPointer) OrderStatus() OrderStatusEnum {
	return OrderStatusEnum(message.Int32Fixed(m.Ptr, 228, 224))
}

// OrderStatus This is required. Needs to be set to one of the following values by the
// Server to indicate the current status of the order, unless NoneOrder =
// 1:
//
// ORDER_STATUS_UNSPECIFIED (0): The status of the order is unset. Use this
// when NoneOrder = 1, or when the Server does not know the true status of
// an order when the Order Update Report message is sent.
// ORDER_STATUS_ORDER_SENT (1): When a Client sends an order to the Server,
// then the Client internally will set the status to ORDER_STATUS_ORDER_SENT.
// The Server will not set this Status.
// ORDER_STATUS_PENDING_OPEN (2): This means the Server has accepted the
// order but it is not yet considered in a fully working state for any reason.
// order but it is not yet considered in a fully working state for any reason.
// ORDER_STATUS_PENDING_CHILD (3): This status applies to a Limit or Stop
// order attached to a parent order. It will have this status if the parent
// order has not yet filled.
// ORDER_STATUS_OPEN (4): Order is open and working.
// ORDER_STATUS_PENDING_CANCEL_REPLACE (5): Order is pending a Cancel and
// Replace operation. The Server should send a OrderUpdate message with the
// OrderUpdateReason set to GENERAL_OrderUpdate and the OrderStatus set to
// this status to indicate the pending Cancel and Replace operation.
// ORDER_STATUS_PENDING_CANCEL (6): Order is pending a Cancel operation.
// The Server should send a OrderUpdate message with the OrderUpdateReason
// set to GENERAL_OrderUpdate and the OrderStatus set to this status to indicate
// the pending Cancel operation.
// ORDER_STATUS_FILLED (7): Order is filled and no longer working.
// ORDER_STATUS_CANCELED (8): Order is canceled. If the user tries to cancel
// an order that has already been canceled, then continue to return this
// Order Status for it.
// ORDER_STATUS_REJECTED (9): Order has been rejected after the initial order
// submission. It is not working.
// ORDER_STATUS_PARTIALLY_FILLED (10): Order is partially filled and still
// working.
func (m OrderUpdateFixedPointerBuilder) OrderStatus() OrderStatusEnum {
	return OrderStatusEnum(message.Int32Fixed(m.Ptr, 228, 224))
}

// OrderUpdateReason This is required. This field needs to be set to one of the following values
// by the Server to indicate the reason the OrderUpdate is being sent.
//
// OpenOrdersRequest_RESPONSE (1): When the OrderUpdate is specifically sent
// in response to an OpenOrdersRequest request, this is the OrderUpdateReason.
// in response to an OpenOrdersRequest request, this is the OrderUpdateReason.
// NEW_ORDER_ACCEPTED (2): This OrderUpdate indicates a new order has been
// accepted.
// GENERAL_OrderUpdate (3): A general order update. For example, when an
// order is in the process of being canceled, the Server may send an OrderUpdate
// message with the OrderStatus set to ORDER_STATUS_PENDING_CANCEL. In this
// particular case the OrderUpdateReason needs to be set to GENERAL_OrderUpdate
// (3).
// ORDER_FILLED (4): Upon a complete fill of the order, this is the OrderUpdateReason.
// This OrderUpdateReason must only be used when an OrderUpdate is sent at
// the moment in time of a fill. A previously filled order that is being
// restated in response to an OpenOrdersRequest must not use this OrderUpdateReason.
// restated in response to an OpenOrdersRequest must not use this OrderUpdateReason.
// ORDER_FILLED_PARTIALLY (5): Upon a partial fill of the order, this is
// the OrderUpdateReason. This OrderUpdateReason must only be used when an
// OrderUpdate is sent at the moment in time of a fill. A previously filled
// order that is being restated in response to an OpenOrdersRequest must
// not use this OrderUpdateReason.
// ORDER_CANCELED (6): This OrderUpdateReason indicates the order is now
// successfully canceled.
// ORDER_CANCEL_REPLACE_COMPLETE (7): This OrderUpdateReason indicates the
// order is now successfully canceled and replaced (modified).
// NEW_ORDER_REJECTED (8): After an order has been submitted by the Client,
// it has been rejected for any reason and was never working, the Server
// will send through an OrderUpdate with this OrderUpdateReason. In this
// case the Server needs to set the OrderStatus in the OrderUpdate to ORDER_STATUS_REJECTED.
// case the Server needs to set the OrderStatus in the OrderUpdate to ORDER_STATUS_REJECTED.
//
// The following fields need to be set for a NEW_ORDER_REJECTED OrderUpdateReason:
// OrderUpdateReason, OrderStatus, MessageNumber, TotalNumMessages, ClientOrderID,
// Symbol, Exchange, TradeAccount.
// ORDER_CANCEL_REJECTED (9): A request by the Client to cancel the order
// with the CancelOrder message has been rejected.
//
// The current status of the order must be set in the OrderStatus member
// of the OrderUpdate message.
//
// In the case where the given ServerOrderID in a CancelOrder message from
// the Client is not known, then respond with an OrderUpdate message and
// set the OrderUpdateReason to ORDER_CANCEL_REJECTED and set the OrderStatus
// to ORDER_STATUS_REJECTED.
//
// In the case where the ServerOrderID is not known, then in the OrderUpdate
// message, it should not be set. However, in this case it is necessary to
// set the ClientOrderID in the OrderUpdate message to the given ClientOrderID
// in the CancelOrder message. The ClientOrderID must always be set in this
// case in an OrderUpdate.
//
// In the case where the order has already been canceled for the given ServerOrderID
// in a CancelOrder message, then respond with an OrderUpdate message and
// set the OrderUpdateReason to ORDER_CANCEL_REJECTED and set the OrderStatus
// to ORDER_STATUS_CANCELED.
//
// If for some reason the Server is uncertain as to the status of the order
// that was attempted to be canceled, then set the OrderStatus to ORDER_STATUS_UNSPECIFIED.
// that was attempted to be canceled, then set the OrderStatus to ORDER_STATUS_UNSPECIFIED.
// ORDER_CANCEL_REPLACE_REJECTED (10): A request by the Client to cancel
// and replace the order with the CancelReplaceOrder message has been rejected.
// and replace the order with the CancelReplaceOrder message has been rejected.
//
// The current status of the order must be set in the OrderStatus member
// of the OrderUpdate message.
//
// In the case where the given ServerOrderID in a CancelReplaceOrder message
// from the Client is not known, then respond with an OrderUpdate message
// and set the OrderUpdateReason to ORDER_CANCEL_REPLACE_REJECTED and set
// the OrderStatus to ORDER_STATUS_REJECTED.
//
// In the case where the ServerOrderID is not known, then in the OrderUpdate
// message, it should not be set. However, in this case it is necessary to
// set the ClientOrderID in the OrderUpdate message to the given ClientOrderID
// in the CancelReplaceOrder message. The ClientOrderID must always be set
// in this case in an OrderUpdate.
//
// If for some reason the Server is uncertain as to the status of the order
// that was attempted to be canceled and replaced, then set the OrderStatus
// to ORDER_STATUS_UNSPECIFIED.
func (m OrderUpdateFixedPointer) OrderUpdateReason() OrderUpdateReasonEnum {
	return OrderUpdateReasonEnum(message.Int32Fixed(m.Ptr, 232, 228))
}

// OrderUpdateReason This is required. This field needs to be set to one of the following values
// by the Server to indicate the reason the OrderUpdate is being sent.
//
// OpenOrdersRequest_RESPONSE (1): When the OrderUpdate is specifically sent
// in response to an OpenOrdersRequest request, this is the OrderUpdateReason.
// in response to an OpenOrdersRequest request, this is the OrderUpdateReason.
// NEW_ORDER_ACCEPTED (2): This OrderUpdate indicates a new order has been
// accepted.
// GENERAL_OrderUpdate (3): A general order update. For example, when an
// order is in the process of being canceled, the Server may send an OrderUpdate
// message with the OrderStatus set to ORDER_STATUS_PENDING_CANCEL. In this
// particular case the OrderUpdateReason needs to be set to GENERAL_OrderUpdate
// (3).
// ORDER_FILLED (4): Upon a complete fill of the order, this is the OrderUpdateReason.
// This OrderUpdateReason must only be used when an OrderUpdate is sent at
// the moment in time of a fill. A previously filled order that is being
// restated in response to an OpenOrdersRequest must not use this OrderUpdateReason.
// restated in response to an OpenOrdersRequest must not use this OrderUpdateReason.
// ORDER_FILLED_PARTIALLY (5): Upon a partial fill of the order, this is
// the OrderUpdateReason. This OrderUpdateReason must only be used when an
// OrderUpdate is sent at the moment in time of a fill. A previously filled
// order that is being restated in response to an OpenOrdersRequest must
// not use this OrderUpdateReason.
// ORDER_CANCELED (6): This OrderUpdateReason indicates the order is now
// successfully canceled.
// ORDER_CANCEL_REPLACE_COMPLETE (7): This OrderUpdateReason indicates the
// order is now successfully canceled and replaced (modified).
// NEW_ORDER_REJECTED (8): After an order has been submitted by the Client,
// it has been rejected for any reason and was never working, the Server
// will send through an OrderUpdate with this OrderUpdateReason. In this
// case the Server needs to set the OrderStatus in the OrderUpdate to ORDER_STATUS_REJECTED.
// case the Server needs to set the OrderStatus in the OrderUpdate to ORDER_STATUS_REJECTED.
//
// The following fields need to be set for a NEW_ORDER_REJECTED OrderUpdateReason:
// OrderUpdateReason, OrderStatus, MessageNumber, TotalNumMessages, ClientOrderID,
// Symbol, Exchange, TradeAccount.
// ORDER_CANCEL_REJECTED (9): A request by the Client to cancel the order
// with the CancelOrder message has been rejected.
//
// The current status of the order must be set in the OrderStatus member
// of the OrderUpdate message.
//
// In the case where the given ServerOrderID in a CancelOrder message from
// the Client is not known, then respond with an OrderUpdate message and
// set the OrderUpdateReason to ORDER_CANCEL_REJECTED and set the OrderStatus
// to ORDER_STATUS_REJECTED.
//
// In the case where the ServerOrderID is not known, then in the OrderUpdate
// message, it should not be set. However, in this case it is necessary to
// set the ClientOrderID in the OrderUpdate message to the given ClientOrderID
// in the CancelOrder message. The ClientOrderID must always be set in this
// case in an OrderUpdate.
//
// In the case where the order has already been canceled for the given ServerOrderID
// in a CancelOrder message, then respond with an OrderUpdate message and
// set the OrderUpdateReason to ORDER_CANCEL_REJECTED and set the OrderStatus
// to ORDER_STATUS_CANCELED.
//
// If for some reason the Server is uncertain as to the status of the order
// that was attempted to be canceled, then set the OrderStatus to ORDER_STATUS_UNSPECIFIED.
// that was attempted to be canceled, then set the OrderStatus to ORDER_STATUS_UNSPECIFIED.
// ORDER_CANCEL_REPLACE_REJECTED (10): A request by the Client to cancel
// and replace the order with the CancelReplaceOrder message has been rejected.
// and replace the order with the CancelReplaceOrder message has been rejected.
//
// The current status of the order must be set in the OrderStatus member
// of the OrderUpdate message.
//
// In the case where the given ServerOrderID in a CancelReplaceOrder message
// from the Client is not known, then respond with an OrderUpdate message
// and set the OrderUpdateReason to ORDER_CANCEL_REPLACE_REJECTED and set
// the OrderStatus to ORDER_STATUS_REJECTED.
//
// In the case where the ServerOrderID is not known, then in the OrderUpdate
// message, it should not be set. However, in this case it is necessary to
// set the ClientOrderID in the OrderUpdate message to the given ClientOrderID
// in the CancelReplaceOrder message. The ClientOrderID must always be set
// in this case in an OrderUpdate.
//
// If for some reason the Server is uncertain as to the status of the order
// that was attempted to be canceled and replaced, then set the OrderStatus
// to ORDER_STATUS_UNSPECIFIED.
func (m OrderUpdateFixedPointerBuilder) OrderUpdateReason() OrderUpdateReasonEnum {
	return OrderUpdateReasonEnum(message.Int32Fixed(m.Ptr, 232, 228))
}

// OrderType The order type. Can be set to one of the following.
//
// ORDER_TYPE_MARKET
// ORDER_TYPE_LIMIT
// ORDER_TYPE_STOP
// ORDER_TYPE_STOP_LIMIT
// ORDER_TYPE_MARKET_IF_TOUCHED
func (m OrderUpdateFixedPointer) OrderType() OrderTypeEnum {
	return OrderTypeEnum(message.Int32Fixed(m.Ptr, 236, 232))
}

// OrderType The order type. Can be set to one of the following.
//
// ORDER_TYPE_MARKET
// ORDER_TYPE_LIMIT
// ORDER_TYPE_STOP
// ORDER_TYPE_STOP_LIMIT
// ORDER_TYPE_MARKET_IF_TOUCHED
func (m OrderUpdateFixedPointerBuilder) OrderType() OrderTypeEnum {
	return OrderTypeEnum(message.Int32Fixed(m.Ptr, 236, 232))
}

// BuySell Indicates whether the order is a Buy or Sell order. Can be set to one
// of the following constants: DTC::BUY (1) or DTC::SELL (2).
func (m OrderUpdateFixedPointer) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Ptr, 240, 236))
}

// BuySell Indicates whether the order is a Buy or Sell order. Can be set to one
// of the following constants: DTC::BUY (1) or DTC::SELL (2).
func (m OrderUpdateFixedPointerBuilder) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Ptr, 240, 236))
}

// Price1 For orders that require a price, this is the order price.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedPointer) Price1() float64 {
	return message.Float64Fixed(m.Ptr, 248, 240)
}

// Price1 For orders that require a price, this is the order price.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedPointerBuilder) Price1() float64 {
	return message.Float64Fixed(m.Ptr, 248, 240)
}

// Price2 For Stop-Limit orders this is the Limit price. Otherwise, this is unset.
// For Stop-Limit orders this is the Limit price. Otherwise, this is unset.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedPointer) Price2() float64 {
	return message.Float64Fixed(m.Ptr, 256, 248)
}

// Price2 For Stop-Limit orders this is the Limit price. Otherwise, this is unset.
// For Stop-Limit orders this is the Limit price. Otherwise, this is unset.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedPointerBuilder) Price2() float64 {
	return message.Float64Fixed(m.Ptr, 256, 248)
}

// TimeInForce The Time in Force of the order. Can be any of the following:
//
// TIF_DAY
// TIF_GOOD_TILL_CANCELED
// TIF_GOOD_TILL_DATE_TIME
// TIF_IMMEDIATE_OR_CANCEL
// TIF_ALL_OR_NONE
// TIF_FILL_OR_KILL
func (m OrderUpdateFixedPointer) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(message.Int32Fixed(m.Ptr, 260, 256))
}

// TimeInForce The Time in Force of the order. Can be any of the following:
//
// TIF_DAY
// TIF_GOOD_TILL_CANCELED
// TIF_GOOD_TILL_DATE_TIME
// TIF_IMMEDIATE_OR_CANCEL
// TIF_ALL_OR_NONE
// TIF_FILL_OR_KILL
func (m OrderUpdateFixedPointerBuilder) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(message.Int32Fixed(m.Ptr, 260, 256))
}

// GoodTillDateTime The expiration Date and Time of the order in the case when TimeInForce
// is TIF_GOOD_TILL_DATE_TIME.
func (m OrderUpdateFixedPointer) GoodTillDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 272, 264))
}

// GoodTillDateTime The expiration Date and Time of the order in the case when TimeInForce
// is TIF_GOOD_TILL_DATE_TIME.
func (m OrderUpdateFixedPointerBuilder) GoodTillDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 272, 264))
}

// OrderQuantity The quantity of the order.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedPointer) OrderQuantity() float64 {
	return message.Float64Fixed(m.Ptr, 280, 272)
}

// OrderQuantity The quantity of the order.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedPointerBuilder) OrderQuantity() float64 {
	return message.Float64Fixed(m.Ptr, 280, 272)
}

// FilledQuantity The number of shares or contracts that have filled in the order.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedPointer) FilledQuantity() float64 {
	return message.Float64Fixed(m.Ptr, 288, 280)
}

// FilledQuantity The number of shares or contracts that have filled in the order.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedPointerBuilder) FilledQuantity() float64 {
	return message.Float64Fixed(m.Ptr, 288, 280)
}

// RemainingQuantity The number of shares or contracts that still remain to be filled.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedPointer) RemainingQuantity() float64 {
	return message.Float64Fixed(m.Ptr, 296, 288)
}

// RemainingQuantity The number of shares or contracts that still remain to be filled.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedPointerBuilder) RemainingQuantity() float64 {
	return message.Float64Fixed(m.Ptr, 296, 288)
}

// AverageFillPrice The average price of all of the fills for the order.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedPointer) AverageFillPrice() float64 {
	return message.Float64Fixed(m.Ptr, 304, 296)
}

// AverageFillPrice The average price of all of the fills for the order.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedPointerBuilder) AverageFillPrice() float64 {
	return message.Float64Fixed(m.Ptr, 304, 296)
}

// LastFillPrice The price of the most recent fill.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedPointer) LastFillPrice() float64 {
	return message.Float64Fixed(m.Ptr, 312, 304)
}

// LastFillPrice The price of the most recent fill.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedPointerBuilder) LastFillPrice() float64 {
	return message.Float64Fixed(m.Ptr, 312, 304)
}

// LastFillDateTime The date and Time of the most recent fill.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
func (m OrderUpdateFixedPointer) LastFillDateTime() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64Fixed(m.Ptr, 320, 312))
}

// LastFillDateTime The date and Time of the most recent fill.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
func (m OrderUpdateFixedPointerBuilder) LastFillDateTime() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64Fixed(m.Ptr, 320, 312))
}

// LastFillQuantity The number of contracts/shares that has filled for the specific order
// fill that is currently reported through the OrderUpdate message.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedPointer) LastFillQuantity() float64 {
	return message.Float64Fixed(m.Ptr, 328, 320)
}

// LastFillQuantity The number of contracts/shares that has filled for the specific order
// fill that is currently reported through the OrderUpdate message.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedPointerBuilder) LastFillQuantity() float64 {
	return message.Float64Fixed(m.Ptr, 328, 320)
}

// LastFillExecutionID The unique identifier for the most recent fill.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
func (m OrderUpdateFixedPointer) LastFillExecutionID() string {
	return message.StringFixed(m.Ptr, 392, 328)
}

// LastFillExecutionID The unique identifier for the most recent fill.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
func (m OrderUpdateFixedPointerBuilder) LastFillExecutionID() string {
	return message.StringFixed(m.Ptr, 392, 328)
}

// TradeAccount The trade account the order belongs to.
func (m OrderUpdateFixedPointer) TradeAccount() string {
	return message.StringFixed(m.Ptr, 424, 392)
}

// TradeAccount The trade account the order belongs to.
func (m OrderUpdateFixedPointerBuilder) TradeAccount() string {
	return message.StringFixed(m.Ptr, 424, 392)
}

// InfoText Free-form text with information to communicate about the order. When an
// order is rejected, this should be set by the Server to indicate the reason
// for the rejection.
func (m OrderUpdateFixedPointer) InfoText() string {
	return message.StringFixed(m.Ptr, 520, 424)
}

// InfoText Free-form text with information to communicate about the order. When an
// order is rejected, this should be set by the Server to indicate the reason
// for the rejection.
func (m OrderUpdateFixedPointerBuilder) InfoText() string {
	return message.StringFixed(m.Ptr, 520, 424)
}

// NoOrders Set by the Server to 1 to indicate there are no orders when OpenOrdersRequest
// message has been received and is being responded to. Otherwise, leave
// at the default of 0.
func (m OrderUpdateFixedPointer) NoOrders() uint8 {
	return message.Uint8Fixed(m.Ptr, 521, 520)
}

// NoOrders Set by the Server to 1 to indicate there are no orders when OpenOrdersRequest
// message has been received and is being responded to. Otherwise, leave
// at the default of 0.
func (m OrderUpdateFixedPointerBuilder) NoOrders() uint8 {
	return message.Uint8Fixed(m.Ptr, 521, 520)
}

// ParentServerOrderID This is the ServerOrderID of the parent order when the order that this
// Order Update Report is for, is a child order in a bracket order. Otherwise,
// this is an empty text string.
func (m OrderUpdateFixedPointer) ParentServerOrderID() string {
	return message.StringFixed(m.Ptr, 553, 521)
}

// ParentServerOrderID This is the ServerOrderID of the parent order when the order that this
// Order Update Report is for, is a child order in a bracket order. Otherwise,
// this is an empty text string.
func (m OrderUpdateFixedPointerBuilder) ParentServerOrderID() string {
	return message.StringFixed(m.Ptr, 553, 521)
}

// OCOLinkedOrderServerOrderID In the case of an OCO order set submitted with SubmitNewOCOOrder, whether
// it has a Parent order or not, this is the ServerOrderID of the other order
// in the OCO pair. These two orders are considered "linked" together. Otherwise,
// this is an empty text string.
func (m OrderUpdateFixedPointer) OCOLinkedOrderServerOrderID() string {
	return message.StringFixed(m.Ptr, 585, 553)
}

// OCOLinkedOrderServerOrderID In the case of an OCO order set submitted with SubmitNewOCOOrder, whether
// it has a Parent order or not, this is the ServerOrderID of the other order
// in the OCO pair. These two orders are considered "linked" together. Otherwise,
// this is an empty text string.
func (m OrderUpdateFixedPointerBuilder) OCOLinkedOrderServerOrderID() string {
	return message.StringFixed(m.Ptr, 585, 553)
}

// OpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m OrderUpdateFixedPointer) OpenOrClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32Fixed(m.Ptr, 592, 588))
}

// OpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m OrderUpdateFixedPointerBuilder) OpenOrClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32Fixed(m.Ptr, 592, 588))
}

// PreviousClientOrderID The PreviousClientOrderID is the previous ClientOrderID provided by the
// Client for the order, if the Client changed it during order cancel and
// replace request or an order cancel request.
//
// A Server only is obligated to provide this field immediately after the
// ClientOrderID has been changed. Subsequent OrderUpdate messages do not
// need to set this field.
func (m OrderUpdateFixedPointer) PreviousClientOrderID() string {
	return message.StringFixed(m.Ptr, 624, 592)
}

// PreviousClientOrderID The PreviousClientOrderID is the previous ClientOrderID provided by the
// Client for the order, if the Client changed it during order cancel and
// replace request or an order cancel request.
//
// A Server only is obligated to provide this field immediately after the
// ClientOrderID has been changed. Subsequent OrderUpdate messages do not
// need to set this field.
func (m OrderUpdateFixedPointerBuilder) PreviousClientOrderID() string {
	return message.StringFixed(m.Ptr, 624, 592)
}

// FreeFormText This is the optional free-form text that was originally set with a new
// order using the new order messages.
func (m OrderUpdateFixedPointer) FreeFormText() string {
	return message.StringFixed(m.Ptr, 672, 624)
}

// FreeFormText This is the optional free-form text that was originally set with a new
// order using the new order messages.
func (m OrderUpdateFixedPointerBuilder) FreeFormText() string {
	return message.StringFixed(m.Ptr, 672, 624)
}

// OrderReceivedDateTime This is the Date-Time when the original order was received by the Server.
// Order modifications normally will not cause this Date-Time to be updated.
// Order modifications normally will not cause this Date-Time to be updated.
func (m OrderUpdateFixedPointer) OrderReceivedDateTime() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64Fixed(m.Ptr, 680, 672))
}

// OrderReceivedDateTime This is the Date-Time when the original order was received by the Server.
// Order modifications normally will not cause this Date-Time to be updated.
// Order modifications normally will not cause this Date-Time to be updated.
func (m OrderUpdateFixedPointerBuilder) OrderReceivedDateTime() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64Fixed(m.Ptr, 680, 672))
}

// LatestTransactionDateTime
func (m OrderUpdateFixedPointer) LatestTransactionDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Ptr, 688, 680))
}

// LatestTransactionDateTime
func (m OrderUpdateFixedPointerBuilder) LatestTransactionDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Ptr, 688, 680))
}

// SetRequestID Set to 0 unless this is in response to an OpenOrdersRequest, in which
// case this must be set to the RequestID given in the OpenOrdersRequest.
//
// If this OrderUpdate is unsolicited, for example a real-time fill or other
// unsolicited order event, the Server must leave this at 0.
func (m OrderUpdateFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetTotalNumMessages This indicates the total number of OrderUpdate messages when a batch of
// reports is being sent in response to an OpenOrdersRequest. If there is
// only one order being sent, this will be 1. The Server must use a value
// of 1 for an unsolicited report. A Client should not rely on this field
// for an unsolicited report.
func (m OrderUpdateFixedPointerBuilder) SetTotalNumMessages(value int32) {
	message.SetInt32Fixed(m.Ptr, 12, 8, value)
}

// SetMessageNumber This indicates the 1-based index of the OrderUpdate message when a batch
// of reports is being sent in response to an OpenOrdersRequest. If there
// is only one order being sent, this will be 1. Use a value of 1 for an
// unsolicited report. A Client should not rely on this field for an unsolicited
// report.
func (m OrderUpdateFixedPointerBuilder) SetMessageNumber(value int32) {
	message.SetInt32Fixed(m.Ptr, 16, 12, value)
}

// SetSymbol The symbol for the order.
func (m OrderUpdateFixedPointerBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Ptr, 80, 16, value)
}

// SetExchange The optional exchange for the symbol.
func (m OrderUpdateFixedPointerBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Ptr, 96, 80, value)
}

// SetPreviousServerOrderID Used upon a Cancel and Replace operation (ORDER_CANCEL_REPLACE_COMPLETE)
// where a new Server Order identifier is given.
//
// In this case this field needs to be set to the previous Server Order identifier.
// In this case this field needs to be set to the previous Server Order identifier.
//
// This should be left at the default setting of empty in the case where
// the Server does not change the Server Order identifier upon a Cancel and
// Replace operation.
func (m OrderUpdateFixedPointerBuilder) SetPreviousServerOrderID(value string) {
	message.SetStringFixed(m.Ptr, 128, 96, value)
}

// SetServerOrderID The ServerOrderID is set by the server and uniquely identifies the order.
// When a new order is submitted by the Client and the Server responds with
// an OrderUpdate, this field needs to be set to the order identifier which
// is good for the life of the order.
//
// This ServerOrderID can optionally change on a Cancel and Replace operation.
// In this case, the PreviousServerOrderID field will contain the previous
// ServerOrderID upon a ORDER_CANCEL_REPLACE_COMPLETE OrderUpdateReason.
//
// ServerOrderID must always be set except it is not required in the cases
// when the OrderUpdateReason is one of the following: NEW_ORDER_REJECTED,
// ORDER_CANCEL_REJECTED, ORDER_CANCEL_REPLACE_REJECTED. If it is not set,
// then the ClientOrderID must be set.
func (m OrderUpdateFixedPointerBuilder) SetServerOrderID(value string) {
	message.SetStringFixed(m.Ptr, 160, 128, value)
}

// SetClientOrderID The ClientOrderID is the order identifier provided by the Client. When
// the Client submits a new order, cancels and replaces an existing order,
// or cancels an order, then the Client needs to specify this identifier.
//
// The Client must maintain the same order identifier throughout the life
// of the order.
//
// The Server should persist the ClientOrderID across sessions. A session
// is defined as the period of time from the start of the network connection
// between the Client and Server to the end of that connection.
//
// The Client should only rely upon the ClientOrderID being set when the
// OrderUpdateReason is one of the following: NEW_ORDER_ACCEPTED, NEW_ORDER_REJECTED,
// ORDER_CANCELED, ORDER_CANCEL_REPLACE_COMPLETE, ORDER_CANCEL_REJECTED,
// ORDER_CANCEL_REPLACE_REJECTED.
//
// After a new order has been accepted, the Client will rely upon the given
// ServerOrderID from the server to identify the order and should no longer
// rely upon the given ClientOrderID. However, the Client needs to maintain
// a copy of the ClientOrderID for any subsequent order modifications and
// cancellations because this is a required field for those.
func (m OrderUpdateFixedPointerBuilder) SetClientOrderID(value string) {
	message.SetStringFixed(m.Ptr, 192, 160, value)
}

// SetExchangeOrderID The order identifier from the exchange that handles the order. This is
// optional.
func (m OrderUpdateFixedPointerBuilder) SetExchangeOrderID(value string) {
	message.SetStringFixed(m.Ptr, 224, 192, value)
}

// SetOrderStatus This is required. Needs to be set to one of the following values by the
// Server to indicate the current status of the order, unless NoneOrder =
// 1:
//
// ORDER_STATUS_UNSPECIFIED (0): The status of the order is unset. Use this
// when NoneOrder = 1, or when the Server does not know the true status of
// an order when the Order Update Report message is sent.
// ORDER_STATUS_ORDER_SENT (1): When a Client sends an order to the Server,
// then the Client internally will set the status to ORDER_STATUS_ORDER_SENT.
// The Server will not set this Status.
// ORDER_STATUS_PENDING_OPEN (2): This means the Server has accepted the
// order but it is not yet considered in a fully working state for any reason.
// order but it is not yet considered in a fully working state for any reason.
// ORDER_STATUS_PENDING_CHILD (3): This status applies to a Limit or Stop
// order attached to a parent order. It will have this status if the parent
// order has not yet filled.
// ORDER_STATUS_OPEN (4): Order is open and working.
// ORDER_STATUS_PENDING_CANCEL_REPLACE (5): Order is pending a Cancel and
// Replace operation. The Server should send a OrderUpdate message with the
// OrderUpdateReason set to GENERAL_OrderUpdate and the OrderStatus set to
// this status to indicate the pending Cancel and Replace operation.
// ORDER_STATUS_PENDING_CANCEL (6): Order is pending a Cancel operation.
// The Server should send a OrderUpdate message with the OrderUpdateReason
// set to GENERAL_OrderUpdate and the OrderStatus set to this status to indicate
// the pending Cancel operation.
// ORDER_STATUS_FILLED (7): Order is filled and no longer working.
// ORDER_STATUS_CANCELED (8): Order is canceled. If the user tries to cancel
// an order that has already been canceled, then continue to return this
// Order Status for it.
// ORDER_STATUS_REJECTED (9): Order has been rejected after the initial order
// submission. It is not working.
// ORDER_STATUS_PARTIALLY_FILLED (10): Order is partially filled and still
// working.
func (m OrderUpdateFixedPointerBuilder) SetOrderStatus(value OrderStatusEnum) {
	message.SetInt32Fixed(m.Ptr, 228, 224, int32(value))
}

// SetOrderUpdateReason This is required. This field needs to be set to one of the following values
// by the Server to indicate the reason the OrderUpdate is being sent.
//
// OpenOrdersRequest_RESPONSE (1): When the OrderUpdate is specifically sent
// in response to an OpenOrdersRequest request, this is the OrderUpdateReason.
// in response to an OpenOrdersRequest request, this is the OrderUpdateReason.
// NEW_ORDER_ACCEPTED (2): This OrderUpdate indicates a new order has been
// accepted.
// GENERAL_OrderUpdate (3): A general order update. For example, when an
// order is in the process of being canceled, the Server may send an OrderUpdate
// message with the OrderStatus set to ORDER_STATUS_PENDING_CANCEL. In this
// particular case the OrderUpdateReason needs to be set to GENERAL_OrderUpdate
// (3).
// ORDER_FILLED (4): Upon a complete fill of the order, this is the OrderUpdateReason.
// This OrderUpdateReason must only be used when an OrderUpdate is sent at
// the moment in time of a fill. A previously filled order that is being
// restated in response to an OpenOrdersRequest must not use this OrderUpdateReason.
// restated in response to an OpenOrdersRequest must not use this OrderUpdateReason.
// ORDER_FILLED_PARTIALLY (5): Upon a partial fill of the order, this is
// the OrderUpdateReason. This OrderUpdateReason must only be used when an
// OrderUpdate is sent at the moment in time of a fill. A previously filled
// order that is being restated in response to an OpenOrdersRequest must
// not use this OrderUpdateReason.
// ORDER_CANCELED (6): This OrderUpdateReason indicates the order is now
// successfully canceled.
// ORDER_CANCEL_REPLACE_COMPLETE (7): This OrderUpdateReason indicates the
// order is now successfully canceled and replaced (modified).
// NEW_ORDER_REJECTED (8): After an order has been submitted by the Client,
// it has been rejected for any reason and was never working, the Server
// will send through an OrderUpdate with this OrderUpdateReason. In this
// case the Server needs to set the OrderStatus in the OrderUpdate to ORDER_STATUS_REJECTED.
// case the Server needs to set the OrderStatus in the OrderUpdate to ORDER_STATUS_REJECTED.
//
// The following fields need to be set for a NEW_ORDER_REJECTED OrderUpdateReason:
// OrderUpdateReason, OrderStatus, MessageNumber, TotalNumMessages, ClientOrderID,
// Symbol, Exchange, TradeAccount.
// ORDER_CANCEL_REJECTED (9): A request by the Client to cancel the order
// with the CancelOrder message has been rejected.
//
// The current status of the order must be set in the OrderStatus member
// of the OrderUpdate message.
//
// In the case where the given ServerOrderID in a CancelOrder message from
// the Client is not known, then respond with an OrderUpdate message and
// set the OrderUpdateReason to ORDER_CANCEL_REJECTED and set the OrderStatus
// to ORDER_STATUS_REJECTED.
//
// In the case where the ServerOrderID is not known, then in the OrderUpdate
// message, it should not be set. However, in this case it is necessary to
// set the ClientOrderID in the OrderUpdate message to the given ClientOrderID
// in the CancelOrder message. The ClientOrderID must always be set in this
// case in an OrderUpdate.
//
// In the case where the order has already been canceled for the given ServerOrderID
// in a CancelOrder message, then respond with an OrderUpdate message and
// set the OrderUpdateReason to ORDER_CANCEL_REJECTED and set the OrderStatus
// to ORDER_STATUS_CANCELED.
//
// If for some reason the Server is uncertain as to the status of the order
// that was attempted to be canceled, then set the OrderStatus to ORDER_STATUS_UNSPECIFIED.
// that was attempted to be canceled, then set the OrderStatus to ORDER_STATUS_UNSPECIFIED.
// ORDER_CANCEL_REPLACE_REJECTED (10): A request by the Client to cancel
// and replace the order with the CancelReplaceOrder message has been rejected.
// and replace the order with the CancelReplaceOrder message has been rejected.
//
// The current status of the order must be set in the OrderStatus member
// of the OrderUpdate message.
//
// In the case where the given ServerOrderID in a CancelReplaceOrder message
// from the Client is not known, then respond with an OrderUpdate message
// and set the OrderUpdateReason to ORDER_CANCEL_REPLACE_REJECTED and set
// the OrderStatus to ORDER_STATUS_REJECTED.
//
// In the case where the ServerOrderID is not known, then in the OrderUpdate
// message, it should not be set. However, in this case it is necessary to
// set the ClientOrderID in the OrderUpdate message to the given ClientOrderID
// in the CancelReplaceOrder message. The ClientOrderID must always be set
// in this case in an OrderUpdate.
//
// If for some reason the Server is uncertain as to the status of the order
// that was attempted to be canceled and replaced, then set the OrderStatus
// to ORDER_STATUS_UNSPECIFIED.
func (m OrderUpdateFixedPointerBuilder) SetOrderUpdateReason(value OrderUpdateReasonEnum) {
	message.SetInt32Fixed(m.Ptr, 232, 228, int32(value))
}

// SetOrderType The order type. Can be set to one of the following.
//
// ORDER_TYPE_MARKET
// ORDER_TYPE_LIMIT
// ORDER_TYPE_STOP
// ORDER_TYPE_STOP_LIMIT
// ORDER_TYPE_MARKET_IF_TOUCHED
func (m OrderUpdateFixedPointerBuilder) SetOrderType(value OrderTypeEnum) {
	message.SetInt32Fixed(m.Ptr, 236, 232, int32(value))
}

// SetBuySell Indicates whether the order is a Buy or Sell order. Can be set to one
// of the following constants: DTC::BUY (1) or DTC::SELL (2).
func (m OrderUpdateFixedPointerBuilder) SetBuySell(value BuySellEnum) {
	message.SetInt32Fixed(m.Ptr, 240, 236, int32(value))
}

// SetPrice1 For orders that require a price, this is the order price.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedPointerBuilder) SetPrice1(value float64) {
	message.SetFloat64Fixed(m.Ptr, 248, 240, value)
}

// SetPrice2 For Stop-Limit orders this is the Limit price. Otherwise, this is unset.
// For Stop-Limit orders this is the Limit price. Otherwise, this is unset.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedPointerBuilder) SetPrice2(value float64) {
	message.SetFloat64Fixed(m.Ptr, 256, 248, value)
}

// SetTimeInForce The Time in Force of the order. Can be any of the following:
//
// TIF_DAY
// TIF_GOOD_TILL_CANCELED
// TIF_GOOD_TILL_DATE_TIME
// TIF_IMMEDIATE_OR_CANCEL
// TIF_ALL_OR_NONE
// TIF_FILL_OR_KILL
func (m OrderUpdateFixedPointerBuilder) SetTimeInForce(value TimeInForceEnum) {
	message.SetInt32Fixed(m.Ptr, 260, 256, int32(value))
}

// SetGoodTillDateTime The expiration Date and Time of the order in the case when TimeInForce
// is TIF_GOOD_TILL_DATE_TIME.
func (m OrderUpdateFixedPointerBuilder) SetGoodTillDateTime(value DateTime) {
	message.SetInt64Fixed(m.Ptr, 272, 264, int64(value))
}

// SetOrderQuantity The quantity of the order.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedPointerBuilder) SetOrderQuantity(value float64) {
	message.SetFloat64Fixed(m.Ptr, 280, 272, value)
}

// SetFilledQuantity The number of shares or contracts that have filled in the order.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedPointerBuilder) SetFilledQuantity(value float64) {
	message.SetFloat64Fixed(m.Ptr, 288, 280, value)
}

// SetRemainingQuantity The number of shares or contracts that still remain to be filled.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedPointerBuilder) SetRemainingQuantity(value float64) {
	message.SetFloat64Fixed(m.Ptr, 296, 288, value)
}

// SetAverageFillPrice The average price of all of the fills for the order.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedPointerBuilder) SetAverageFillPrice(value float64) {
	message.SetFloat64Fixed(m.Ptr, 304, 296, value)
}

// SetLastFillPrice The price of the most recent fill.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedPointerBuilder) SetLastFillPrice(value float64) {
	message.SetFloat64Fixed(m.Ptr, 312, 304, value)
}

// SetLastFillDateTime The date and Time of the most recent fill.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
func (m OrderUpdateFixedPointerBuilder) SetLastFillDateTime(value DateTimeWithMillisecondsInt) {
	message.SetInt64Fixed(m.Ptr, 320, 312, int64(value))
}

// SetLastFillQuantity The number of contracts/shares that has filled for the specific order
// fill that is currently reported through the OrderUpdate message.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedPointerBuilder) SetLastFillQuantity(value float64) {
	message.SetFloat64Fixed(m.Ptr, 328, 320, value)
}

// SetLastFillExecutionID The unique identifier for the most recent fill.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
func (m OrderUpdateFixedPointerBuilder) SetLastFillExecutionID(value string) {
	message.SetStringFixed(m.Ptr, 392, 328, value)
}

// SetTradeAccount The trade account the order belongs to.
func (m OrderUpdateFixedPointerBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Ptr, 424, 392, value)
}

// SetInfoText Free-form text with information to communicate about the order. When an
// order is rejected, this should be set by the Server to indicate the reason
// for the rejection.
func (m OrderUpdateFixedPointerBuilder) SetInfoText(value string) {
	message.SetStringFixed(m.Ptr, 520, 424, value)
}

// SetNoOrders Set by the Server to 1 to indicate there are no orders when OpenOrdersRequest
// message has been received and is being responded to. Otherwise, leave
// at the default of 0.
func (m OrderUpdateFixedPointerBuilder) SetNoOrders(value uint8) {
	message.SetUint8Fixed(m.Ptr, 521, 520, value)
}

// SetParentServerOrderID This is the ServerOrderID of the parent order when the order that this
// Order Update Report is for, is a child order in a bracket order. Otherwise,
// this is an empty text string.
func (m OrderUpdateFixedPointerBuilder) SetParentServerOrderID(value string) {
	message.SetStringFixed(m.Ptr, 553, 521, value)
}

// SetOCOLinkedOrderServerOrderID In the case of an OCO order set submitted with SubmitNewOCOOrder, whether
// it has a Parent order or not, this is the ServerOrderID of the other order
// in the OCO pair. These two orders are considered "linked" together. Otherwise,
// this is an empty text string.
func (m OrderUpdateFixedPointerBuilder) SetOCOLinkedOrderServerOrderID(value string) {
	message.SetStringFixed(m.Ptr, 585, 553, value)
}

// SetOpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m OrderUpdateFixedPointerBuilder) SetOpenOrClose(value OpenCloseTradeEnum) {
	message.SetInt32Fixed(m.Ptr, 592, 588, int32(value))
}

// SetPreviousClientOrderID The PreviousClientOrderID is the previous ClientOrderID provided by the
// Client for the order, if the Client changed it during order cancel and
// replace request or an order cancel request.
//
// A Server only is obligated to provide this field immediately after the
// ClientOrderID has been changed. Subsequent OrderUpdate messages do not
// need to set this field.
func (m OrderUpdateFixedPointerBuilder) SetPreviousClientOrderID(value string) {
	message.SetStringFixed(m.Ptr, 624, 592, value)
}

// SetFreeFormText This is the optional free-form text that was originally set with a new
// order using the new order messages.
func (m OrderUpdateFixedPointerBuilder) SetFreeFormText(value string) {
	message.SetStringFixed(m.Ptr, 672, 624, value)
}

// SetOrderReceivedDateTime This is the Date-Time when the original order was received by the Server.
// Order modifications normally will not cause this Date-Time to be updated.
// Order modifications normally will not cause this Date-Time to be updated.
func (m OrderUpdateFixedPointerBuilder) SetOrderReceivedDateTime(value DateTimeWithMillisecondsInt) {
	message.SetInt64Fixed(m.Ptr, 680, 672, int64(value))
}

// SetLatestTransactionDateTime
func (m OrderUpdateFixedPointerBuilder) SetLatestTransactionDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64Fixed(m.Ptr, 688, 680, float64(value))
}

// Copy
func (m OrderUpdateFixedPointer) Copy(to OrderUpdateFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumMessages(m.TotalNumMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetPreviousServerOrderID(m.PreviousServerOrderID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetOrderStatus(m.OrderStatus())
	to.SetOrderUpdateReason(m.OrderUpdateReason())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetOrderQuantity(m.OrderQuantity())
	to.SetFilledQuantity(m.FilledQuantity())
	to.SetRemainingQuantity(m.RemainingQuantity())
	to.SetAverageFillPrice(m.AverageFillPrice())
	to.SetLastFillPrice(m.LastFillPrice())
	to.SetLastFillDateTime(m.LastFillDateTime())
	to.SetLastFillQuantity(m.LastFillQuantity())
	to.SetLastFillExecutionID(m.LastFillExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetInfoText(m.InfoText())
	to.SetNoOrders(m.NoOrders())
	to.SetParentServerOrderID(m.ParentServerOrderID())
	to.SetOCOLinkedOrderServerOrderID(m.OCOLinkedOrderServerOrderID())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPreviousClientOrderID(m.PreviousClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOrderReceivedDateTime(m.OrderReceivedDateTime())
	to.SetLatestTransactionDateTime(m.LatestTransactionDateTime())
}

// Copy
func (m OrderUpdateFixedPointerBuilder) Copy(to OrderUpdateFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumMessages(m.TotalNumMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetPreviousServerOrderID(m.PreviousServerOrderID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetOrderStatus(m.OrderStatus())
	to.SetOrderUpdateReason(m.OrderUpdateReason())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetOrderQuantity(m.OrderQuantity())
	to.SetFilledQuantity(m.FilledQuantity())
	to.SetRemainingQuantity(m.RemainingQuantity())
	to.SetAverageFillPrice(m.AverageFillPrice())
	to.SetLastFillPrice(m.LastFillPrice())
	to.SetLastFillDateTime(m.LastFillDateTime())
	to.SetLastFillQuantity(m.LastFillQuantity())
	to.SetLastFillExecutionID(m.LastFillExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetInfoText(m.InfoText())
	to.SetNoOrders(m.NoOrders())
	to.SetParentServerOrderID(m.ParentServerOrderID())
	to.SetOCOLinkedOrderServerOrderID(m.OCOLinkedOrderServerOrderID())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPreviousClientOrderID(m.PreviousClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOrderReceivedDateTime(m.OrderReceivedDateTime())
	to.SetLatestTransactionDateTime(m.LatestTransactionDateTime())
}

// CopyPointer
func (m OrderUpdateFixedPointer) CopyPointer(to OrderUpdateFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumMessages(m.TotalNumMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetPreviousServerOrderID(m.PreviousServerOrderID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetOrderStatus(m.OrderStatus())
	to.SetOrderUpdateReason(m.OrderUpdateReason())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetOrderQuantity(m.OrderQuantity())
	to.SetFilledQuantity(m.FilledQuantity())
	to.SetRemainingQuantity(m.RemainingQuantity())
	to.SetAverageFillPrice(m.AverageFillPrice())
	to.SetLastFillPrice(m.LastFillPrice())
	to.SetLastFillDateTime(m.LastFillDateTime())
	to.SetLastFillQuantity(m.LastFillQuantity())
	to.SetLastFillExecutionID(m.LastFillExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetInfoText(m.InfoText())
	to.SetNoOrders(m.NoOrders())
	to.SetParentServerOrderID(m.ParentServerOrderID())
	to.SetOCOLinkedOrderServerOrderID(m.OCOLinkedOrderServerOrderID())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPreviousClientOrderID(m.PreviousClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOrderReceivedDateTime(m.OrderReceivedDateTime())
	to.SetLatestTransactionDateTime(m.LatestTransactionDateTime())
}

// CopyPointer
func (m OrderUpdateFixedPointerBuilder) CopyPointer(to OrderUpdateFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumMessages(m.TotalNumMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetPreviousServerOrderID(m.PreviousServerOrderID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetOrderStatus(m.OrderStatus())
	to.SetOrderUpdateReason(m.OrderUpdateReason())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetOrderQuantity(m.OrderQuantity())
	to.SetFilledQuantity(m.FilledQuantity())
	to.SetRemainingQuantity(m.RemainingQuantity())
	to.SetAverageFillPrice(m.AverageFillPrice())
	to.SetLastFillPrice(m.LastFillPrice())
	to.SetLastFillDateTime(m.LastFillDateTime())
	to.SetLastFillQuantity(m.LastFillQuantity())
	to.SetLastFillExecutionID(m.LastFillExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetInfoText(m.InfoText())
	to.SetNoOrders(m.NoOrders())
	to.SetParentServerOrderID(m.ParentServerOrderID())
	to.SetOCOLinkedOrderServerOrderID(m.OCOLinkedOrderServerOrderID())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPreviousClientOrderID(m.PreviousClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOrderReceivedDateTime(m.OrderReceivedDateTime())
	to.SetLatestTransactionDateTime(m.LatestTransactionDateTime())
}

// CopyTo
func (m OrderUpdateFixedPointer) CopyTo(to OrderUpdateBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumMessages(m.TotalNumMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetPreviousServerOrderID(m.PreviousServerOrderID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetOrderStatus(m.OrderStatus())
	to.SetOrderUpdateReason(m.OrderUpdateReason())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetOrderQuantity(m.OrderQuantity())
	to.SetFilledQuantity(m.FilledQuantity())
	to.SetRemainingQuantity(m.RemainingQuantity())
	to.SetAverageFillPrice(m.AverageFillPrice())
	to.SetLastFillPrice(m.LastFillPrice())
	to.SetLastFillDateTime(m.LastFillDateTime())
	to.SetLastFillQuantity(m.LastFillQuantity())
	to.SetLastFillExecutionID(m.LastFillExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetInfoText(m.InfoText())
	to.SetNoOrders(m.NoOrders())
	to.SetParentServerOrderID(m.ParentServerOrderID())
	to.SetOCOLinkedOrderServerOrderID(m.OCOLinkedOrderServerOrderID())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPreviousClientOrderID(m.PreviousClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOrderReceivedDateTime(m.OrderReceivedDateTime())
	to.SetLatestTransactionDateTime(m.LatestTransactionDateTime())
}

// CopyTo
func (m OrderUpdateFixedPointerBuilder) CopyTo(to OrderUpdateBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumMessages(m.TotalNumMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetPreviousServerOrderID(m.PreviousServerOrderID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetOrderStatus(m.OrderStatus())
	to.SetOrderUpdateReason(m.OrderUpdateReason())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetOrderQuantity(m.OrderQuantity())
	to.SetFilledQuantity(m.FilledQuantity())
	to.SetRemainingQuantity(m.RemainingQuantity())
	to.SetAverageFillPrice(m.AverageFillPrice())
	to.SetLastFillPrice(m.LastFillPrice())
	to.SetLastFillDateTime(m.LastFillDateTime())
	to.SetLastFillQuantity(m.LastFillQuantity())
	to.SetLastFillExecutionID(m.LastFillExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetInfoText(m.InfoText())
	to.SetNoOrders(m.NoOrders())
	to.SetParentServerOrderID(m.ParentServerOrderID())
	to.SetOCOLinkedOrderServerOrderID(m.OCOLinkedOrderServerOrderID())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPreviousClientOrderID(m.PreviousClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOrderReceivedDateTime(m.OrderReceivedDateTime())
	to.SetLatestTransactionDateTime(m.LatestTransactionDateTime())
}

// CopyToPointer
func (m OrderUpdateFixedPointer) CopyToPointer(to OrderUpdatePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumMessages(m.TotalNumMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetPreviousServerOrderID(m.PreviousServerOrderID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetOrderStatus(m.OrderStatus())
	to.SetOrderUpdateReason(m.OrderUpdateReason())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetOrderQuantity(m.OrderQuantity())
	to.SetFilledQuantity(m.FilledQuantity())
	to.SetRemainingQuantity(m.RemainingQuantity())
	to.SetAverageFillPrice(m.AverageFillPrice())
	to.SetLastFillPrice(m.LastFillPrice())
	to.SetLastFillDateTime(m.LastFillDateTime())
	to.SetLastFillQuantity(m.LastFillQuantity())
	to.SetLastFillExecutionID(m.LastFillExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetInfoText(m.InfoText())
	to.SetNoOrders(m.NoOrders())
	to.SetParentServerOrderID(m.ParentServerOrderID())
	to.SetOCOLinkedOrderServerOrderID(m.OCOLinkedOrderServerOrderID())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPreviousClientOrderID(m.PreviousClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOrderReceivedDateTime(m.OrderReceivedDateTime())
	to.SetLatestTransactionDateTime(m.LatestTransactionDateTime())
}

// CopyToPointer
func (m OrderUpdateFixedPointerBuilder) CopyToPointer(to OrderUpdatePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumMessages(m.TotalNumMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetPreviousServerOrderID(m.PreviousServerOrderID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetOrderStatus(m.OrderStatus())
	to.SetOrderUpdateReason(m.OrderUpdateReason())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetOrderQuantity(m.OrderQuantity())
	to.SetFilledQuantity(m.FilledQuantity())
	to.SetRemainingQuantity(m.RemainingQuantity())
	to.SetAverageFillPrice(m.AverageFillPrice())
	to.SetLastFillPrice(m.LastFillPrice())
	to.SetLastFillDateTime(m.LastFillDateTime())
	to.SetLastFillQuantity(m.LastFillQuantity())
	to.SetLastFillExecutionID(m.LastFillExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetInfoText(m.InfoText())
	to.SetNoOrders(m.NoOrders())
	to.SetParentServerOrderID(m.ParentServerOrderID())
	to.SetOCOLinkedOrderServerOrderID(m.OCOLinkedOrderServerOrderID())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPreviousClientOrderID(m.PreviousClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOrderReceivedDateTime(m.OrderReceivedDateTime())
	to.SetLatestTransactionDateTime(m.LatestTransactionDateTime())
}
