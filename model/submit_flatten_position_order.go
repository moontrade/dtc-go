// generated by github.com/moontrade/dtc-go/codegen/go at 2022-03-12 12:55:01.923209 +0800 WITA m=+0.028978918

package model

import (
	"github.com/moontrade/dtc-go/message"
)

const SubmitFlattenPositionOrderSize = 28

const SubmitFlattenPositionOrderFixedSize = 198

// {
//     Size             = SubmitFlattenPositionOrderSize  (28)
//     Type             = SUBMIT_FLATTEN_POSITION_ORDER  (209)
//     BaseSize         = SubmitFlattenPositionOrderSize  (28)
//     Symbol           = ""
//     Exchange         = ""
//     TradeAccount     = ""
//     ClientOrderID    = ""
//     FreeFormText     = ""
//     IsAutomatedOrder = false
// }
var _SubmitFlattenPositionOrderDefault = []byte{28, 0, 209, 0, 28, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// {
//     Size             = SubmitFlattenPositionOrderFixedSize  (198)
//     Type             = SUBMIT_FLATTEN_POSITION_ORDER  (209)
//     Symbol           = ""
//     Exchange         = ""
//     TradeAccount     = ""
//     ClientOrderID    = ""
//     FreeFormText     = ""
//     IsAutomatedOrder = false
// }
var _SubmitFlattenPositionOrderFixedDefault = []byte{198, 0, 209, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type SubmitFlattenPositionOrder struct {
	message.VLS
}

type SubmitFlattenPositionOrderBuilder struct {
	message.VLSBuilder
}

type SubmitFlattenPositionOrderFixed struct {
	message.Fixed
}

type SubmitFlattenPositionOrderFixedBuilder struct {
	message.Fixed
}

type SubmitFlattenPositionOrderPointer struct {
	message.VLSPointer
}

type SubmitFlattenPositionOrderPointerBuilder struct {
	message.VLSPointerBuilder
}

type SubmitFlattenPositionOrderFixedPointer struct {
	message.FixedPointer
}

type SubmitFlattenPositionOrderFixedPointerBuilder struct {
	message.FixedPointer
}

func NewSubmitFlattenPositionOrderFrom(b []byte) SubmitFlattenPositionOrder {
	return SubmitFlattenPositionOrder{VLS: message.NewVLSFrom(b)}
}

func WrapSubmitFlattenPositionOrder(b []byte) SubmitFlattenPositionOrder {
	return SubmitFlattenPositionOrder{VLS: message.WrapVLS(b)}
}

func NewSubmitFlattenPositionOrder() SubmitFlattenPositionOrderBuilder {
	a := SubmitFlattenPositionOrderBuilder{message.NewVLSBuilder(28)}
	a.Unsafe().SetBytes(0, _SubmitFlattenPositionOrderDefault)
	return a
}

func NewSubmitFlattenPositionOrderFixedFrom(b []byte) SubmitFlattenPositionOrderFixed {
	return SubmitFlattenPositionOrderFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapSubmitFlattenPositionOrderFixed(b []byte) SubmitFlattenPositionOrderFixed {
	return SubmitFlattenPositionOrderFixed{Fixed: message.WrapFixed(b)}
}

func NewSubmitFlattenPositionOrderFixed() SubmitFlattenPositionOrderFixedBuilder {
	a := SubmitFlattenPositionOrderFixedBuilder{message.NewFixed(198)}
	a.Unsafe().SetBytes(0, _SubmitFlattenPositionOrderFixedDefault)
	return a
}

func AllocSubmitFlattenPositionOrder() SubmitFlattenPositionOrderPointerBuilder {
	a := SubmitFlattenPositionOrderPointerBuilder{message.AllocVLSBuilder(28)}
	a.Ptr.SetBytes(0, _SubmitFlattenPositionOrderDefault)
	return a
}

func AllocSubmitFlattenPositionOrderFrom(b []byte) SubmitFlattenPositionOrderPointer {
	return SubmitFlattenPositionOrderPointer{VLSPointer: message.AllocVLSFrom(b)}
}

func AllocSubmitFlattenPositionOrderFixed() SubmitFlattenPositionOrderFixedPointerBuilder {
	a := SubmitFlattenPositionOrderFixedPointerBuilder{message.AllocFixed(198)}
	a.Ptr.SetBytes(0, _SubmitFlattenPositionOrderFixedDefault)
	return a
}

func AllocSubmitFlattenPositionOrderFixedFrom(b []byte) SubmitFlattenPositionOrderFixedPointer {
	return SubmitFlattenPositionOrderFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size             = SubmitFlattenPositionOrderSize  (28)
//     Type             = SUBMIT_FLATTEN_POSITION_ORDER  (209)
//     BaseSize         = SubmitFlattenPositionOrderSize  (28)
//     Symbol           = ""
//     Exchange         = ""
//     TradeAccount     = ""
//     ClientOrderID    = ""
//     FreeFormText     = ""
//     IsAutomatedOrder = false
// }
func (m SubmitFlattenPositionOrderBuilder) Clear() {
	m.Unsafe().SetBytes(0, _SubmitFlattenPositionOrderDefault)
}

// Clear
// {
//     Size             = SubmitFlattenPositionOrderFixedSize  (198)
//     Type             = SUBMIT_FLATTEN_POSITION_ORDER  (209)
//     Symbol           = ""
//     Exchange         = ""
//     TradeAccount     = ""
//     ClientOrderID    = ""
//     FreeFormText     = ""
//     IsAutomatedOrder = false
// }
func (m SubmitFlattenPositionOrderFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _SubmitFlattenPositionOrderFixedDefault)
}

// Clear
// {
//     Size             = SubmitFlattenPositionOrderSize  (28)
//     Type             = SUBMIT_FLATTEN_POSITION_ORDER  (209)
//     BaseSize         = SubmitFlattenPositionOrderSize  (28)
//     Symbol           = ""
//     Exchange         = ""
//     TradeAccount     = ""
//     ClientOrderID    = ""
//     FreeFormText     = ""
//     IsAutomatedOrder = false
// }
func (m SubmitFlattenPositionOrderPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _SubmitFlattenPositionOrderDefault)
}

// Clear
// {
//     Size             = SubmitFlattenPositionOrderFixedSize  (198)
//     Type             = SUBMIT_FLATTEN_POSITION_ORDER  (209)
//     Symbol           = ""
//     Exchange         = ""
//     TradeAccount     = ""
//     ClientOrderID    = ""
//     FreeFormText     = ""
//     IsAutomatedOrder = false
// }
func (m SubmitFlattenPositionOrderFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _SubmitFlattenPositionOrderFixedDefault)
}

// ToBuilder
func (m SubmitFlattenPositionOrder) ToBuilder() SubmitFlattenPositionOrderBuilder {
	return SubmitFlattenPositionOrderBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m SubmitFlattenPositionOrderFixed) ToBuilder() SubmitFlattenPositionOrderFixedBuilder {
	return SubmitFlattenPositionOrderFixedBuilder{m.Fixed}
}

// ToBuilder
func (m SubmitFlattenPositionOrderPointer) ToBuilder() SubmitFlattenPositionOrderPointerBuilder {
	return SubmitFlattenPositionOrderPointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m SubmitFlattenPositionOrderFixedPointer) ToBuilder() SubmitFlattenPositionOrderFixedPointerBuilder {
	return SubmitFlattenPositionOrderFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m SubmitFlattenPositionOrderBuilder) Finish() SubmitFlattenPositionOrder {
	return SubmitFlattenPositionOrder{m.VLSBuilder.Finish()}
}

// Finish
func (m SubmitFlattenPositionOrderFixedBuilder) Finish() SubmitFlattenPositionOrderFixed {
	return SubmitFlattenPositionOrderFixed{m.Fixed}
}

// Finish
func (m *SubmitFlattenPositionOrderPointerBuilder) Finish() SubmitFlattenPositionOrderPointer {
	return SubmitFlattenPositionOrderPointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *SubmitFlattenPositionOrderFixedPointerBuilder) Finish() SubmitFlattenPositionOrderFixedPointer {
	return SubmitFlattenPositionOrderFixedPointer{m.FixedPointer}
}

// Symbol The symbol of the Trade Position to flatten.
func (m SubmitFlattenPositionOrder) Symbol() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// Symbol The symbol of the Trade Position to flatten.
func (m SubmitFlattenPositionOrderBuilder) Symbol() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// Symbol The symbol of the Trade Position to flatten.
func (m SubmitFlattenPositionOrderPointer) Symbol() string {
	return message.StringVLS(m.Ptr, 10, 6)
}

// Symbol The symbol of the Trade Position to flatten.
func (m SubmitFlattenPositionOrderPointerBuilder) Symbol() string {
	return message.StringVLS(m.Ptr, 10, 6)
}

// Exchange The optional exchange for the Symbol.
func (m SubmitFlattenPositionOrder) Exchange() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// Exchange The optional exchange for the Symbol.
func (m SubmitFlattenPositionOrderBuilder) Exchange() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// Exchange The optional exchange for the Symbol.
func (m SubmitFlattenPositionOrderPointer) Exchange() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// Exchange The optional exchange for the Symbol.
func (m SubmitFlattenPositionOrderPointerBuilder) Exchange() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// TradeAccount The trade account as a text string of the Trade Position to flatten.
func (m SubmitFlattenPositionOrder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// TradeAccount The trade account as a text string of the Trade Position to flatten.
func (m SubmitFlattenPositionOrderBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// TradeAccount The trade account as a text string of the Trade Position to flatten.
func (m SubmitFlattenPositionOrderPointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 18, 14)
}

// TradeAccount The trade account as a text string of the Trade Position to flatten.
func (m SubmitFlattenPositionOrderPointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 18, 14)
}

// ClientOrderID The Client supplied order identifier for the order which will be created
// to flatten the Trade Position.
//
// The Server will remember this for the life of the order.
func (m SubmitFlattenPositionOrder) ClientOrderID() string {
	return message.StringVLS(m.Unsafe(), 22, 18)
}

// ClientOrderID The Client supplied order identifier for the order which will be created
// to flatten the Trade Position.
//
// The Server will remember this for the life of the order.
func (m SubmitFlattenPositionOrderBuilder) ClientOrderID() string {
	return message.StringVLS(m.Unsafe(), 22, 18)
}

// ClientOrderID The Client supplied order identifier for the order which will be created
// to flatten the Trade Position.
//
// The Server will remember this for the life of the order.
func (m SubmitFlattenPositionOrderPointer) ClientOrderID() string {
	return message.StringVLS(m.Ptr, 22, 18)
}

// ClientOrderID The Client supplied order identifier for the order which will be created
// to flatten the Trade Position.
//
// The Server will remember this for the life of the order.
func (m SubmitFlattenPositionOrderPointerBuilder) ClientOrderID() string {
	return message.StringVLS(m.Ptr, 22, 18)
}

// FreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order which will be created to flatten the
// Trade Position.
//
// The Server is not under any obligation to use this text and it may place
// a limitation on the length of this text.
func (m SubmitFlattenPositionOrder) FreeFormText() string {
	return message.StringVLS(m.Unsafe(), 26, 22)
}

// FreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order which will be created to flatten the
// Trade Position.
//
// The Server is not under any obligation to use this text and it may place
// a limitation on the length of this text.
func (m SubmitFlattenPositionOrderBuilder) FreeFormText() string {
	return message.StringVLS(m.Unsafe(), 26, 22)
}

// FreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order which will be created to flatten the
// Trade Position.
//
// The Server is not under any obligation to use this text and it may place
// a limitation on the length of this text.
func (m SubmitFlattenPositionOrderPointer) FreeFormText() string {
	return message.StringVLS(m.Ptr, 26, 22)
}

// FreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order which will be created to flatten the
// Trade Position.
//
// The Server is not under any obligation to use this text and it may place
// a limitation on the length of this text.
func (m SubmitFlattenPositionOrderPointerBuilder) FreeFormText() string {
	return message.StringVLS(m.Ptr, 26, 22)
}

// IsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitFlattenPositionOrder) IsAutomatedOrder() bool {
	return message.BoolVLS(m.Unsafe(), 27, 26)
}

// IsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitFlattenPositionOrderBuilder) IsAutomatedOrder() bool {
	return message.BoolVLS(m.Unsafe(), 27, 26)
}

// IsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitFlattenPositionOrderPointer) IsAutomatedOrder() bool {
	return message.BoolVLS(m.Ptr, 27, 26)
}

// IsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitFlattenPositionOrderPointerBuilder) IsAutomatedOrder() bool {
	return message.BoolVLS(m.Ptr, 27, 26)
}

// Symbol The symbol of the Trade Position to flatten.
func (m SubmitFlattenPositionOrderFixed) Symbol() string {
	return message.StringFixed(m.Unsafe(), 68, 4)
}

// Symbol The symbol of the Trade Position to flatten.
func (m SubmitFlattenPositionOrderFixedBuilder) Symbol() string {
	return message.StringFixed(m.Unsafe(), 68, 4)
}

// Symbol The symbol of the Trade Position to flatten.
func (m SubmitFlattenPositionOrderFixedPointer) Symbol() string {
	return message.StringFixed(m.Ptr, 68, 4)
}

// Symbol The symbol of the Trade Position to flatten.
func (m SubmitFlattenPositionOrderFixedPointerBuilder) Symbol() string {
	return message.StringFixed(m.Ptr, 68, 4)
}

// Exchange The optional exchange for the Symbol.
func (m SubmitFlattenPositionOrderFixed) Exchange() string {
	return message.StringFixed(m.Unsafe(), 84, 68)
}

// Exchange The optional exchange for the Symbol.
func (m SubmitFlattenPositionOrderFixedBuilder) Exchange() string {
	return message.StringFixed(m.Unsafe(), 84, 68)
}

// Exchange The optional exchange for the Symbol.
func (m SubmitFlattenPositionOrderFixedPointer) Exchange() string {
	return message.StringFixed(m.Ptr, 84, 68)
}

// Exchange The optional exchange for the Symbol.
func (m SubmitFlattenPositionOrderFixedPointerBuilder) Exchange() string {
	return message.StringFixed(m.Ptr, 84, 68)
}

// TradeAccount The trade account as a text string of the Trade Position to flatten.
func (m SubmitFlattenPositionOrderFixed) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 116, 84)
}

// TradeAccount The trade account as a text string of the Trade Position to flatten.
func (m SubmitFlattenPositionOrderFixedBuilder) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 116, 84)
}

// TradeAccount The trade account as a text string of the Trade Position to flatten.
func (m SubmitFlattenPositionOrderFixedPointer) TradeAccount() string {
	return message.StringFixed(m.Ptr, 116, 84)
}

// TradeAccount The trade account as a text string of the Trade Position to flatten.
func (m SubmitFlattenPositionOrderFixedPointerBuilder) TradeAccount() string {
	return message.StringFixed(m.Ptr, 116, 84)
}

// ClientOrderID The Client supplied order identifier for the order which will be created
// to flatten the Trade Position.
//
// The Server will remember this for the life of the order.
func (m SubmitFlattenPositionOrderFixed) ClientOrderID() string {
	return message.StringFixed(m.Unsafe(), 148, 116)
}

// ClientOrderID The Client supplied order identifier for the order which will be created
// to flatten the Trade Position.
//
// The Server will remember this for the life of the order.
func (m SubmitFlattenPositionOrderFixedBuilder) ClientOrderID() string {
	return message.StringFixed(m.Unsafe(), 148, 116)
}

// ClientOrderID The Client supplied order identifier for the order which will be created
// to flatten the Trade Position.
//
// The Server will remember this for the life of the order.
func (m SubmitFlattenPositionOrderFixedPointer) ClientOrderID() string {
	return message.StringFixed(m.Ptr, 148, 116)
}

// ClientOrderID The Client supplied order identifier for the order which will be created
// to flatten the Trade Position.
//
// The Server will remember this for the life of the order.
func (m SubmitFlattenPositionOrderFixedPointerBuilder) ClientOrderID() string {
	return message.StringFixed(m.Ptr, 148, 116)
}

// FreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order which will be created to flatten the
// Trade Position.
//
// The Server is not under any obligation to use this text and it may place
// a limitation on the length of this text.
func (m SubmitFlattenPositionOrderFixed) FreeFormText() string {
	return message.StringFixed(m.Unsafe(), 196, 148)
}

// FreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order which will be created to flatten the
// Trade Position.
//
// The Server is not under any obligation to use this text and it may place
// a limitation on the length of this text.
func (m SubmitFlattenPositionOrderFixedBuilder) FreeFormText() string {
	return message.StringFixed(m.Unsafe(), 196, 148)
}

// FreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order which will be created to flatten the
// Trade Position.
//
// The Server is not under any obligation to use this text and it may place
// a limitation on the length of this text.
func (m SubmitFlattenPositionOrderFixedPointer) FreeFormText() string {
	return message.StringFixed(m.Ptr, 196, 148)
}

// FreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order which will be created to flatten the
// Trade Position.
//
// The Server is not under any obligation to use this text and it may place
// a limitation on the length of this text.
func (m SubmitFlattenPositionOrderFixedPointerBuilder) FreeFormText() string {
	return message.StringFixed(m.Ptr, 196, 148)
}

// IsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitFlattenPositionOrderFixed) IsAutomatedOrder() bool {
	return message.BoolFixed(m.Unsafe(), 197, 196)
}

// IsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitFlattenPositionOrderFixedBuilder) IsAutomatedOrder() bool {
	return message.BoolFixed(m.Unsafe(), 197, 196)
}

// IsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitFlattenPositionOrderFixedPointer) IsAutomatedOrder() bool {
	return message.BoolFixed(m.Ptr, 197, 196)
}

// IsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitFlattenPositionOrderFixedPointerBuilder) IsAutomatedOrder() bool {
	return message.BoolFixed(m.Ptr, 197, 196)
}

// SetSymbol The symbol of the Trade Position to flatten.
func (m *SubmitFlattenPositionOrderBuilder) SetSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 10, 6, value)
}

// SetSymbol The symbol of the Trade Position to flatten.
func (m *SubmitFlattenPositionOrderPointerBuilder) SetSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 10, 6, value)
}

// SetExchange The optional exchange for the Symbol.
func (m *SubmitFlattenPositionOrderBuilder) SetExchange(value string) {
	message.SetStringVLS(&m.VLSBuilder, 14, 10, value)
}

// SetExchange The optional exchange for the Symbol.
func (m *SubmitFlattenPositionOrderPointerBuilder) SetExchange(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 14, 10, value)
}

// SetTradeAccount The trade account as a text string of the Trade Position to flatten.
func (m *SubmitFlattenPositionOrderBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 18, 14, value)
}

// SetTradeAccount The trade account as a text string of the Trade Position to flatten.
func (m *SubmitFlattenPositionOrderPointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 18, 14, value)
}

// SetClientOrderID The Client supplied order identifier for the order which will be created
// to flatten the Trade Position.
//
// The Server will remember this for the life of the order.
func (m *SubmitFlattenPositionOrderBuilder) SetClientOrderID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 22, 18, value)
}

// SetClientOrderID The Client supplied order identifier for the order which will be created
// to flatten the Trade Position.
//
// The Server will remember this for the life of the order.
func (m *SubmitFlattenPositionOrderPointerBuilder) SetClientOrderID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 22, 18, value)
}

// SetFreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order which will be created to flatten the
// Trade Position.
//
// The Server is not under any obligation to use this text and it may place
// a limitation on the length of this text.
func (m *SubmitFlattenPositionOrderBuilder) SetFreeFormText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 26, 22, value)
}

// SetFreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order which will be created to flatten the
// Trade Position.
//
// The Server is not under any obligation to use this text and it may place
// a limitation on the length of this text.
func (m *SubmitFlattenPositionOrderPointerBuilder) SetFreeFormText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 26, 22, value)
}

// SetIsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitFlattenPositionOrderBuilder) SetIsAutomatedOrder(value bool) {
	message.SetBoolVLS(m.Unsafe(), 27, 26, value)
}

// SetIsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitFlattenPositionOrderPointerBuilder) SetIsAutomatedOrder(value bool) {
	message.SetBoolVLS(m.Ptr, 27, 26, value)
}

// SetSymbol The symbol of the Trade Position to flatten.
func (m SubmitFlattenPositionOrderFixedBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Unsafe(), 68, 4, value)
}

// SetSymbol The symbol of the Trade Position to flatten.
func (m SubmitFlattenPositionOrderFixedPointerBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Ptr, 68, 4, value)
}

// SetExchange The optional exchange for the Symbol.
func (m SubmitFlattenPositionOrderFixedBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Unsafe(), 84, 68, value)
}

// SetExchange The optional exchange for the Symbol.
func (m SubmitFlattenPositionOrderFixedPointerBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Ptr, 84, 68, value)
}

// SetTradeAccount The trade account as a text string of the Trade Position to flatten.
func (m SubmitFlattenPositionOrderFixedBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Unsafe(), 116, 84, value)
}

// SetTradeAccount The trade account as a text string of the Trade Position to flatten.
func (m SubmitFlattenPositionOrderFixedPointerBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Ptr, 116, 84, value)
}

// SetClientOrderID The Client supplied order identifier for the order which will be created
// to flatten the Trade Position.
//
// The Server will remember this for the life of the order.
func (m SubmitFlattenPositionOrderFixedBuilder) SetClientOrderID(value string) {
	message.SetStringFixed(m.Unsafe(), 148, 116, value)
}

// SetClientOrderID The Client supplied order identifier for the order which will be created
// to flatten the Trade Position.
//
// The Server will remember this for the life of the order.
func (m SubmitFlattenPositionOrderFixedPointerBuilder) SetClientOrderID(value string) {
	message.SetStringFixed(m.Ptr, 148, 116, value)
}

// SetFreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order which will be created to flatten the
// Trade Position.
//
// The Server is not under any obligation to use this text and it may place
// a limitation on the length of this text.
func (m SubmitFlattenPositionOrderFixedBuilder) SetFreeFormText(value string) {
	message.SetStringFixed(m.Unsafe(), 196, 148, value)
}

// SetFreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order which will be created to flatten the
// Trade Position.
//
// The Server is not under any obligation to use this text and it may place
// a limitation on the length of this text.
func (m SubmitFlattenPositionOrderFixedPointerBuilder) SetFreeFormText(value string) {
	message.SetStringFixed(m.Ptr, 196, 148, value)
}

// SetIsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitFlattenPositionOrderFixedBuilder) SetIsAutomatedOrder(value bool) {
	message.SetBoolFixed(m.Unsafe(), 197, 196, value)
}

// SetIsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitFlattenPositionOrderFixedPointerBuilder) SetIsAutomatedOrder(value bool) {
	message.SetBoolFixed(m.Ptr, 197, 196, value)
}

// Copy
func (m SubmitFlattenPositionOrder) Copy(to SubmitFlattenPositionOrderBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// Copy
func (m SubmitFlattenPositionOrderBuilder) Copy(to SubmitFlattenPositionOrderBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyPointer
func (m SubmitFlattenPositionOrder) CopyPointer(to SubmitFlattenPositionOrderPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyPointer
func (m SubmitFlattenPositionOrderBuilder) CopyPointer(to SubmitFlattenPositionOrderPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyToPointer
func (m SubmitFlattenPositionOrder) CopyToPointer(to SubmitFlattenPositionOrderFixedPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyToPointer
func (m SubmitFlattenPositionOrderBuilder) CopyToPointer(to SubmitFlattenPositionOrderFixedPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyTo
func (m SubmitFlattenPositionOrder) CopyTo(to SubmitFlattenPositionOrderFixedBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyTo
func (m SubmitFlattenPositionOrderBuilder) CopyTo(to SubmitFlattenPositionOrderFixedBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// Copy
func (m SubmitFlattenPositionOrderFixed) Copy(to SubmitFlattenPositionOrderFixedBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// Copy
func (m SubmitFlattenPositionOrderFixedBuilder) Copy(to SubmitFlattenPositionOrderFixedBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyPointer
func (m SubmitFlattenPositionOrderFixed) CopyPointer(to SubmitFlattenPositionOrderFixedPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyPointer
func (m SubmitFlattenPositionOrderFixedBuilder) CopyPointer(to SubmitFlattenPositionOrderFixedPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyToPointer
func (m SubmitFlattenPositionOrderFixed) CopyToPointer(to SubmitFlattenPositionOrderPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyToPointer
func (m SubmitFlattenPositionOrderFixedBuilder) CopyToPointer(to SubmitFlattenPositionOrderPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyTo
func (m SubmitFlattenPositionOrderFixed) CopyTo(to SubmitFlattenPositionOrderBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyTo
func (m SubmitFlattenPositionOrderFixedBuilder) CopyTo(to SubmitFlattenPositionOrderBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// Copy
func (m SubmitFlattenPositionOrderPointer) Copy(to SubmitFlattenPositionOrderBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// Copy
func (m SubmitFlattenPositionOrderPointerBuilder) Copy(to SubmitFlattenPositionOrderBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyPointer
func (m SubmitFlattenPositionOrderPointer) CopyPointer(to SubmitFlattenPositionOrderPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyPointer
func (m SubmitFlattenPositionOrderPointerBuilder) CopyPointer(to SubmitFlattenPositionOrderPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyTo
func (m SubmitFlattenPositionOrderPointer) CopyTo(to SubmitFlattenPositionOrderFixedBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyTo
func (m SubmitFlattenPositionOrderPointerBuilder) CopyTo(to SubmitFlattenPositionOrderFixedBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyToPointer
func (m SubmitFlattenPositionOrderPointer) CopyToPointer(to SubmitFlattenPositionOrderFixedPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyToPointer
func (m SubmitFlattenPositionOrderPointerBuilder) CopyToPointer(to SubmitFlattenPositionOrderFixedPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// Copy
func (m SubmitFlattenPositionOrderFixedPointer) Copy(to SubmitFlattenPositionOrderFixedBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// Copy
func (m SubmitFlattenPositionOrderFixedPointerBuilder) Copy(to SubmitFlattenPositionOrderFixedBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyPointer
func (m SubmitFlattenPositionOrderFixedPointer) CopyPointer(to SubmitFlattenPositionOrderFixedPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyPointer
func (m SubmitFlattenPositionOrderFixedPointerBuilder) CopyPointer(to SubmitFlattenPositionOrderFixedPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyTo
func (m SubmitFlattenPositionOrderFixedPointer) CopyTo(to SubmitFlattenPositionOrderBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyTo
func (m SubmitFlattenPositionOrderFixedPointerBuilder) CopyTo(to SubmitFlattenPositionOrderBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyToPointer
func (m SubmitFlattenPositionOrderFixedPointer) CopyToPointer(to SubmitFlattenPositionOrderPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyToPointer
func (m SubmitFlattenPositionOrderFixedPointerBuilder) CopyToPointer(to SubmitFlattenPositionOrderPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}
