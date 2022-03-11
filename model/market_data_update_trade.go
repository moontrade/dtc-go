package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size       = MarketDataUpdateTradeSize  (40)
//     Type       = MARKET_DATA_UPDATE_TRADE  (107)
//     SymbolID   = 0
//     AtBidOrAsk = 0
//     Price      = 0
//     Volume     = 0
//     DateTime   = 0
// }
var _MarketDataUpdateTradeDefault = []byte{40, 0, 107, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const MarketDataUpdateTradeSize = 40

// MarketDataUpdateTrade The Server sends this market data feed message to the Client when a trade
// occurs.
type MarketDataUpdateTrade struct {
	message.Fixed
}

// MarketDataUpdateTradeBuilder The Server sends this market data feed message to the Client when a trade
// occurs.
type MarketDataUpdateTradeBuilder struct {
	message.Fixed
}

func NewMarketDataUpdateTradeFrom(b []byte) MarketDataUpdateTrade {
	return MarketDataUpdateTrade{Fixed: message.NewFixedFrom(b)}
}

func WrapMarketDataUpdateTrade(b []byte) MarketDataUpdateTrade {
	return MarketDataUpdateTrade{Fixed: message.WrapFixed(b)}
}

func NewMarketDataUpdateTrade() MarketDataUpdateTradeBuilder {
	a := MarketDataUpdateTradeBuilder{message.NewFixed(40)}
	a.Unsafe().SetBytes(0, _MarketDataUpdateTradeDefault)
	return a
}

// Clear
// {
//     Size       = MarketDataUpdateTradeSize  (40)
//     Type       = MARKET_DATA_UPDATE_TRADE  (107)
//     SymbolID   = 0
//     AtBidOrAsk = 0
//     Price      = 0
//     Volume     = 0
//     DateTime   = 0
// }
func (m MarketDataUpdateTradeBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketDataUpdateTradeDefault)
}

// ToBuilder
func (m MarketDataUpdateTrade) ToBuilder() MarketDataUpdateTradeBuilder {
	return MarketDataUpdateTradeBuilder{m.Fixed}
}

// Finish
func (m MarketDataUpdateTradeBuilder) Finish() MarketDataUpdateTrade {
	return MarketDataUpdateTrade{m.Fixed}
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateTrade) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateTradeBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// AtBidOrAsk Indicator whether the trade occurred at the bid or ask.
func (m MarketDataUpdateTrade) AtBidOrAsk() AtBidOrAskEnum {
	return AtBidOrAskEnum(message.Uint16Fixed(m.Unsafe(), 10, 8))
}

// AtBidOrAsk Indicator whether the trade occurred at the bid or ask.
func (m MarketDataUpdateTradeBuilder) AtBidOrAsk() AtBidOrAskEnum {
	return AtBidOrAskEnum(message.Uint16Fixed(m.Unsafe(), 10, 8))
}

// Price The price of the trade.
func (m MarketDataUpdateTrade) Price() float64 {
	return message.Float64Fixed(m.Unsafe(), 24, 16)
}

// Price The price of the trade.
func (m MarketDataUpdateTradeBuilder) Price() float64 {
	return message.Float64Fixed(m.Unsafe(), 24, 16)
}

// Volume The volume of the trade.
func (m MarketDataUpdateTrade) Volume() float64 {
	return message.Float64Fixed(m.Unsafe(), 32, 24)
}

// Volume The volume of the trade.
func (m MarketDataUpdateTradeBuilder) Volume() float64 {
	return message.Float64Fixed(m.Unsafe(), 32, 24)
}

// DateTime The Date-Time of the trade.
func (m MarketDataUpdateTrade) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Unsafe(), 40, 32))
}

// DateTime The Date-Time of the trade.
func (m MarketDataUpdateTradeBuilder) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Unsafe(), 40, 32))
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateTradeBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetAtBidOrAsk Indicator whether the trade occurred at the bid or ask.
func (m MarketDataUpdateTradeBuilder) SetAtBidOrAsk(value AtBidOrAskEnum) {
	message.SetUint16Fixed(m.Unsafe(), 10, 8, uint16(value))
}

// SetPrice The price of the trade.
func (m MarketDataUpdateTradeBuilder) SetPrice(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 24, 16, value)
}

// SetVolume The volume of the trade.
func (m MarketDataUpdateTradeBuilder) SetVolume(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 32, 24, value)
}

// SetDateTime The Date-Time of the trade.
func (m MarketDataUpdateTradeBuilder) SetDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64Fixed(m.Unsafe(), 40, 32, float64(value))
}

// Copy
func (m MarketDataUpdateTrade) Copy(to MarketDataUpdateTradeBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetDateTime(m.DateTime())
}

// Copy
func (m MarketDataUpdateTradeBuilder) Copy(to MarketDataUpdateTradeBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetDateTime(m.DateTime())
}

// CopyPointer
func (m MarketDataUpdateTrade) CopyPointer(to MarketDataUpdateTradePointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetDateTime(m.DateTime())
}

// CopyPointer
func (m MarketDataUpdateTradeBuilder) CopyPointer(to MarketDataUpdateTradePointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetDateTime(m.DateTime())
}
