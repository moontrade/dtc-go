package model

import (
	"github.com/moontrade/dtc-go/message"
)

const MarketDataUpdateBidAskCompactSize = 28

// {
//     Size        = MarketDataUpdateBidAskCompactSize  (28)
//     Type        = MARKET_DATA_UPDATE_BID_ASK_COMPACT  (117)
//     BidPrice    = math.MaxFloat32
//     BidQuantity = 0
//     AskPrice    = math.MaxFloat32
//     AskQuantity = 0
//     DateTime    = 0
//     SymbolID    = 0
// }
var _MarketDataUpdateBidAskCompactDefault = []byte{28, 0, 117, 0, 255, 255, 127, 127, 0, 0, 0, 0, 255, 255, 127, 127, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type MarketDataUpdateBidAskCompact struct {
	message.Fixed
}

type MarketDataUpdateBidAskCompactBuilder struct {
	message.Fixed
}

type MarketDataUpdateBidAskCompactPointer struct {
	message.FixedPointer
}

type MarketDataUpdateBidAskCompactPointerBuilder struct {
	message.FixedPointer
}

func NewMarketDataUpdateBidAskCompactFrom(b []byte) MarketDataUpdateBidAskCompact {
	return MarketDataUpdateBidAskCompact{Fixed: message.NewFixedFrom(b)}
}

func WrapMarketDataUpdateBidAskCompact(b []byte) MarketDataUpdateBidAskCompact {
	return MarketDataUpdateBidAskCompact{Fixed: message.WrapFixed(b)}
}

func NewMarketDataUpdateBidAskCompact() MarketDataUpdateBidAskCompactBuilder {
	a := MarketDataUpdateBidAskCompactBuilder{message.NewFixed(28)}
	a.Unsafe().SetBytes(0, _MarketDataUpdateBidAskCompactDefault)
	return a
}

func AllocMarketDataUpdateBidAskCompact() MarketDataUpdateBidAskCompactPointerBuilder {
	a := MarketDataUpdateBidAskCompactPointerBuilder{message.AllocFixed(28)}
	a.Ptr.SetBytes(0, _MarketDataUpdateBidAskCompactDefault)
	return a
}

func AllocMarketDataUpdateBidAskCompactFrom(b []byte) MarketDataUpdateBidAskCompactPointer {
	return MarketDataUpdateBidAskCompactPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size        = MarketDataUpdateBidAskCompactSize  (28)
//     Type        = MARKET_DATA_UPDATE_BID_ASK_COMPACT  (117)
//     BidPrice    = math.MaxFloat32
//     BidQuantity = 0
//     AskPrice    = math.MaxFloat32
//     AskQuantity = 0
//     DateTime    = 0
//     SymbolID    = 0
// }
func (m MarketDataUpdateBidAskCompactBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketDataUpdateBidAskCompactDefault)
}

// Clear
// {
//     Size        = MarketDataUpdateBidAskCompactSize  (28)
//     Type        = MARKET_DATA_UPDATE_BID_ASK_COMPACT  (117)
//     BidPrice    = math.MaxFloat32
//     BidQuantity = 0
//     AskPrice    = math.MaxFloat32
//     AskQuantity = 0
//     DateTime    = 0
//     SymbolID    = 0
// }
func (m MarketDataUpdateBidAskCompactPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketDataUpdateBidAskCompactDefault)
}

// ToBuilder
func (m MarketDataUpdateBidAskCompact) ToBuilder() MarketDataUpdateBidAskCompactBuilder {
	return MarketDataUpdateBidAskCompactBuilder{m.Fixed}
}

// ToBuilder
func (m MarketDataUpdateBidAskCompactPointer) ToBuilder() MarketDataUpdateBidAskCompactPointerBuilder {
	return MarketDataUpdateBidAskCompactPointerBuilder{m.FixedPointer}
}

// Finish
func (m MarketDataUpdateBidAskCompactBuilder) Finish() MarketDataUpdateBidAskCompact {
	return MarketDataUpdateBidAskCompact{m.Fixed}
}

// Finish
func (m *MarketDataUpdateBidAskCompactPointerBuilder) Finish() MarketDataUpdateBidAskCompactPointer {
	return MarketDataUpdateBidAskCompactPointer{m.FixedPointer}
}

// BidPrice
func (m MarketDataUpdateBidAskCompact) BidPrice() float32 {
	return message.Float32Fixed(m.Unsafe(), 8, 4)
}

// BidPrice
func (m MarketDataUpdateBidAskCompactBuilder) BidPrice() float32 {
	return message.Float32Fixed(m.Unsafe(), 8, 4)
}

// BidPrice
func (m MarketDataUpdateBidAskCompactPointer) BidPrice() float32 {
	return message.Float32Fixed(m.Ptr, 8, 4)
}

// BidPrice
func (m MarketDataUpdateBidAskCompactPointerBuilder) BidPrice() float32 {
	return message.Float32Fixed(m.Ptr, 8, 4)
}

// BidQuantity
func (m MarketDataUpdateBidAskCompact) BidQuantity() float32 {
	return message.Float32Fixed(m.Unsafe(), 12, 8)
}

// BidQuantity
func (m MarketDataUpdateBidAskCompactBuilder) BidQuantity() float32 {
	return message.Float32Fixed(m.Unsafe(), 12, 8)
}

// BidQuantity
func (m MarketDataUpdateBidAskCompactPointer) BidQuantity() float32 {
	return message.Float32Fixed(m.Ptr, 12, 8)
}

// BidQuantity
func (m MarketDataUpdateBidAskCompactPointerBuilder) BidQuantity() float32 {
	return message.Float32Fixed(m.Ptr, 12, 8)
}

// AskPrice
func (m MarketDataUpdateBidAskCompact) AskPrice() float32 {
	return message.Float32Fixed(m.Unsafe(), 16, 12)
}

// AskPrice
func (m MarketDataUpdateBidAskCompactBuilder) AskPrice() float32 {
	return message.Float32Fixed(m.Unsafe(), 16, 12)
}

// AskPrice
func (m MarketDataUpdateBidAskCompactPointer) AskPrice() float32 {
	return message.Float32Fixed(m.Ptr, 16, 12)
}

// AskPrice
func (m MarketDataUpdateBidAskCompactPointerBuilder) AskPrice() float32 {
	return message.Float32Fixed(m.Ptr, 16, 12)
}

// AskQuantity
func (m MarketDataUpdateBidAskCompact) AskQuantity() float32 {
	return message.Float32Fixed(m.Unsafe(), 20, 16)
}

// AskQuantity
func (m MarketDataUpdateBidAskCompactBuilder) AskQuantity() float32 {
	return message.Float32Fixed(m.Unsafe(), 20, 16)
}

// AskQuantity
func (m MarketDataUpdateBidAskCompactPointer) AskQuantity() float32 {
	return message.Float32Fixed(m.Ptr, 20, 16)
}

// AskQuantity
func (m MarketDataUpdateBidAskCompactPointerBuilder) AskQuantity() float32 {
	return message.Float32Fixed(m.Ptr, 20, 16)
}

// DateTime
func (m MarketDataUpdateBidAskCompact) DateTime() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Unsafe(), 24, 20))
}

// DateTime
func (m MarketDataUpdateBidAskCompactBuilder) DateTime() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Unsafe(), 24, 20))
}

// DateTime
func (m MarketDataUpdateBidAskCompactPointer) DateTime() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Ptr, 24, 20))
}

// DateTime
func (m MarketDataUpdateBidAskCompactPointerBuilder) DateTime() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Ptr, 24, 20))
}

// SymbolID
func (m MarketDataUpdateBidAskCompact) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 28, 24)
}

// SymbolID
func (m MarketDataUpdateBidAskCompactBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 28, 24)
}

// SymbolID
func (m MarketDataUpdateBidAskCompactPointer) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 28, 24)
}

// SymbolID
func (m MarketDataUpdateBidAskCompactPointerBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 28, 24)
}

// SetBidPrice
func (m MarketDataUpdateBidAskCompactBuilder) SetBidPrice(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 8, 4, value)
}

// SetBidPrice
func (m MarketDataUpdateBidAskCompactPointerBuilder) SetBidPrice(value float32) {
	message.SetFloat32Fixed(m.Ptr, 8, 4, value)
}

// SetBidQuantity
func (m MarketDataUpdateBidAskCompactBuilder) SetBidQuantity(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 12, 8, value)
}

// SetBidQuantity
func (m MarketDataUpdateBidAskCompactPointerBuilder) SetBidQuantity(value float32) {
	message.SetFloat32Fixed(m.Ptr, 12, 8, value)
}

// SetAskPrice
func (m MarketDataUpdateBidAskCompactBuilder) SetAskPrice(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 16, 12, value)
}

// SetAskPrice
func (m MarketDataUpdateBidAskCompactPointerBuilder) SetAskPrice(value float32) {
	message.SetFloat32Fixed(m.Ptr, 16, 12, value)
}

// SetAskQuantity
func (m MarketDataUpdateBidAskCompactBuilder) SetAskQuantity(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 20, 16, value)
}

// SetAskQuantity
func (m MarketDataUpdateBidAskCompactPointerBuilder) SetAskQuantity(value float32) {
	message.SetFloat32Fixed(m.Ptr, 20, 16, value)
}

// SetDateTime
func (m MarketDataUpdateBidAskCompactBuilder) SetDateTime(value DateTime4Byte) {
	message.SetUint32Fixed(m.Unsafe(), 24, 20, uint32(value))
}

// SetDateTime
func (m MarketDataUpdateBidAskCompactPointerBuilder) SetDateTime(value DateTime4Byte) {
	message.SetUint32Fixed(m.Ptr, 24, 20, uint32(value))
}

// SetSymbolID
func (m MarketDataUpdateBidAskCompactBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 28, 24, value)
}

// SetSymbolID
func (m MarketDataUpdateBidAskCompactPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 28, 24, value)
}

// Copy
func (m MarketDataUpdateBidAskCompact) Copy(to MarketDataUpdateBidAskCompactBuilder) {
	to.SetBidPrice(m.BidPrice())
	to.SetBidQuantity(m.BidQuantity())
	to.SetAskPrice(m.AskPrice())
	to.SetAskQuantity(m.AskQuantity())
	to.SetDateTime(m.DateTime())
	to.SetSymbolID(m.SymbolID())
}

// Copy
func (m MarketDataUpdateBidAskCompactBuilder) Copy(to MarketDataUpdateBidAskCompactBuilder) {
	to.SetBidPrice(m.BidPrice())
	to.SetBidQuantity(m.BidQuantity())
	to.SetAskPrice(m.AskPrice())
	to.SetAskQuantity(m.AskQuantity())
	to.SetDateTime(m.DateTime())
	to.SetSymbolID(m.SymbolID())
}

// CopyPointer
func (m MarketDataUpdateBidAskCompact) CopyPointer(to MarketDataUpdateBidAskCompactPointerBuilder) {
	to.SetBidPrice(m.BidPrice())
	to.SetBidQuantity(m.BidQuantity())
	to.SetAskPrice(m.AskPrice())
	to.SetAskQuantity(m.AskQuantity())
	to.SetDateTime(m.DateTime())
	to.SetSymbolID(m.SymbolID())
}

// CopyPointer
func (m MarketDataUpdateBidAskCompactBuilder) CopyPointer(to MarketDataUpdateBidAskCompactPointerBuilder) {
	to.SetBidPrice(m.BidPrice())
	to.SetBidQuantity(m.BidQuantity())
	to.SetAskPrice(m.AskPrice())
	to.SetAskQuantity(m.AskQuantity())
	to.SetDateTime(m.DateTime())
	to.SetSymbolID(m.SymbolID())
}

// Copy
func (m MarketDataUpdateBidAskCompactPointer) Copy(to MarketDataUpdateBidAskCompactBuilder) {
	to.SetBidPrice(m.BidPrice())
	to.SetBidQuantity(m.BidQuantity())
	to.SetAskPrice(m.AskPrice())
	to.SetAskQuantity(m.AskQuantity())
	to.SetDateTime(m.DateTime())
	to.SetSymbolID(m.SymbolID())
}

// Copy
func (m MarketDataUpdateBidAskCompactPointerBuilder) Copy(to MarketDataUpdateBidAskCompactBuilder) {
	to.SetBidPrice(m.BidPrice())
	to.SetBidQuantity(m.BidQuantity())
	to.SetAskPrice(m.AskPrice())
	to.SetAskQuantity(m.AskQuantity())
	to.SetDateTime(m.DateTime())
	to.SetSymbolID(m.SymbolID())
}

// CopyPointer
func (m MarketDataUpdateBidAskCompactPointer) CopyPointer(to MarketDataUpdateBidAskCompactPointerBuilder) {
	to.SetBidPrice(m.BidPrice())
	to.SetBidQuantity(m.BidQuantity())
	to.SetAskPrice(m.AskPrice())
	to.SetAskQuantity(m.AskQuantity())
	to.SetDateTime(m.DateTime())
	to.SetSymbolID(m.SymbolID())
}

// CopyPointer
func (m MarketDataUpdateBidAskCompactPointerBuilder) CopyPointer(to MarketDataUpdateBidAskCompactPointerBuilder) {
	to.SetBidPrice(m.BidPrice())
	to.SetBidQuantity(m.BidQuantity())
	to.SetAskPrice(m.AskPrice())
	to.SetAskQuantity(m.AskQuantity())
	to.SetDateTime(m.DateTime())
	to.SetSymbolID(m.SymbolID())
}