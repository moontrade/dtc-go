package model

import (
	"github.com/moontrade/dtc-go/message"
)

const HistoricalPriceDataTickRecordResponse_IntSize = 32

// {
//     Size          = HistoricalPriceDataTickRecordResponse_IntSize  (32)
//     Type          = HISTORICAL_PRICE_DATA_TICK_RECORD_RESPONSE_INT  (806)
//     RequestID     = 0
//     DateTime      = 0
//     Price         = 0
//     Volume        = 0
//     AtBidOrAsk    = BID_ASK_UNSET  (0)
//     IsFinalRecord = false
// }
var _HistoricalPriceDataTickRecordResponse_IntDefault = []byte{32, 0, 38, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type HistoricalPriceDataTickRecordResponse_Int struct {
	message.Fixed
}

type HistoricalPriceDataTickRecordResponse_IntBuilder struct {
	message.Fixed
}

type HistoricalPriceDataTickRecordResponse_IntPointer struct {
	message.FixedPointer
}

type HistoricalPriceDataTickRecordResponse_IntPointerBuilder struct {
	message.FixedPointer
}

func NewHistoricalPriceDataTickRecordResponse_IntFrom(b []byte) HistoricalPriceDataTickRecordResponse_Int {
	return HistoricalPriceDataTickRecordResponse_Int{Fixed: message.NewFixedFrom(b)}
}

func WrapHistoricalPriceDataTickRecordResponse_Int(b []byte) HistoricalPriceDataTickRecordResponse_Int {
	return HistoricalPriceDataTickRecordResponse_Int{Fixed: message.WrapFixed(b)}
}

func NewHistoricalPriceDataTickRecordResponse_Int() HistoricalPriceDataTickRecordResponse_IntBuilder {
	a := HistoricalPriceDataTickRecordResponse_IntBuilder{message.NewFixed(32)}
	a.Unsafe().SetBytes(0, _HistoricalPriceDataTickRecordResponse_IntDefault)
	return a
}

func AllocHistoricalPriceDataTickRecordResponse_Int() HistoricalPriceDataTickRecordResponse_IntPointerBuilder {
	a := HistoricalPriceDataTickRecordResponse_IntPointerBuilder{message.AllocFixed(32)}
	a.Ptr.SetBytes(0, _HistoricalPriceDataTickRecordResponse_IntDefault)
	return a
}

func AllocHistoricalPriceDataTickRecordResponse_IntFrom(b []byte) HistoricalPriceDataTickRecordResponse_IntPointer {
	return HistoricalPriceDataTickRecordResponse_IntPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size          = HistoricalPriceDataTickRecordResponse_IntSize  (32)
//     Type          = HISTORICAL_PRICE_DATA_TICK_RECORD_RESPONSE_INT  (806)
//     RequestID     = 0
//     DateTime      = 0
//     Price         = 0
//     Volume        = 0
//     AtBidOrAsk    = BID_ASK_UNSET  (0)
//     IsFinalRecord = false
// }
func (m HistoricalPriceDataTickRecordResponse_IntBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalPriceDataTickRecordResponse_IntDefault)
}

// Clear
// {
//     Size          = HistoricalPriceDataTickRecordResponse_IntSize  (32)
//     Type          = HISTORICAL_PRICE_DATA_TICK_RECORD_RESPONSE_INT  (806)
//     RequestID     = 0
//     DateTime      = 0
//     Price         = 0
//     Volume        = 0
//     AtBidOrAsk    = BID_ASK_UNSET  (0)
//     IsFinalRecord = false
// }
func (m HistoricalPriceDataTickRecordResponse_IntPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalPriceDataTickRecordResponse_IntDefault)
}

// ToBuilder
func (m HistoricalPriceDataTickRecordResponse_Int) ToBuilder() HistoricalPriceDataTickRecordResponse_IntBuilder {
	return HistoricalPriceDataTickRecordResponse_IntBuilder{m.Fixed}
}

// ToBuilder
func (m HistoricalPriceDataTickRecordResponse_IntPointer) ToBuilder() HistoricalPriceDataTickRecordResponse_IntPointerBuilder {
	return HistoricalPriceDataTickRecordResponse_IntPointerBuilder{m.FixedPointer}
}

// Finish
func (m HistoricalPriceDataTickRecordResponse_IntBuilder) Finish() HistoricalPriceDataTickRecordResponse_Int {
	return HistoricalPriceDataTickRecordResponse_Int{m.Fixed}
}

// Finish
func (m *HistoricalPriceDataTickRecordResponse_IntPointerBuilder) Finish() HistoricalPriceDataTickRecordResponse_IntPointer {
	return HistoricalPriceDataTickRecordResponse_IntPointer{m.FixedPointer}
}

// RequestID
func (m HistoricalPriceDataTickRecordResponse_Int) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m HistoricalPriceDataTickRecordResponse_IntBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m HistoricalPriceDataTickRecordResponse_IntPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID
func (m HistoricalPriceDataTickRecordResponse_IntPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// DateTime
func (m HistoricalPriceDataTickRecordResponse_Int) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Unsafe(), 16, 8))
}

// DateTime
func (m HistoricalPriceDataTickRecordResponse_IntBuilder) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Unsafe(), 16, 8))
}

// DateTime
func (m HistoricalPriceDataTickRecordResponse_IntPointer) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Ptr, 16, 8))
}

// DateTime
func (m HistoricalPriceDataTickRecordResponse_IntPointerBuilder) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Ptr, 16, 8))
}

// Price
func (m HistoricalPriceDataTickRecordResponse_Int) Price() int32 {
	return message.Int32Fixed(m.Unsafe(), 20, 16)
}

// Price
func (m HistoricalPriceDataTickRecordResponse_IntBuilder) Price() int32 {
	return message.Int32Fixed(m.Unsafe(), 20, 16)
}

// Price
func (m HistoricalPriceDataTickRecordResponse_IntPointer) Price() int32 {
	return message.Int32Fixed(m.Ptr, 20, 16)
}

// Price
func (m HistoricalPriceDataTickRecordResponse_IntPointerBuilder) Price() int32 {
	return message.Int32Fixed(m.Ptr, 20, 16)
}

// Volume
func (m HistoricalPriceDataTickRecordResponse_Int) Volume() int32 {
	return message.Int32Fixed(m.Unsafe(), 24, 20)
}

// Volume
func (m HistoricalPriceDataTickRecordResponse_IntBuilder) Volume() int32 {
	return message.Int32Fixed(m.Unsafe(), 24, 20)
}

// Volume
func (m HistoricalPriceDataTickRecordResponse_IntPointer) Volume() int32 {
	return message.Int32Fixed(m.Ptr, 24, 20)
}

// Volume
func (m HistoricalPriceDataTickRecordResponse_IntPointerBuilder) Volume() int32 {
	return message.Int32Fixed(m.Ptr, 24, 20)
}

// AtBidOrAsk
func (m HistoricalPriceDataTickRecordResponse_Int) AtBidOrAsk() AtBidOrAskEnum {
	return AtBidOrAskEnum(message.Uint16Fixed(m.Unsafe(), 26, 24))
}

// AtBidOrAsk
func (m HistoricalPriceDataTickRecordResponse_IntBuilder) AtBidOrAsk() AtBidOrAskEnum {
	return AtBidOrAskEnum(message.Uint16Fixed(m.Unsafe(), 26, 24))
}

// AtBidOrAsk
func (m HistoricalPriceDataTickRecordResponse_IntPointer) AtBidOrAsk() AtBidOrAskEnum {
	return AtBidOrAskEnum(message.Uint16Fixed(m.Ptr, 26, 24))
}

// AtBidOrAsk
func (m HistoricalPriceDataTickRecordResponse_IntPointerBuilder) AtBidOrAsk() AtBidOrAskEnum {
	return AtBidOrAskEnum(message.Uint16Fixed(m.Ptr, 26, 24))
}

// IsFinalRecord
func (m HistoricalPriceDataTickRecordResponse_Int) IsFinalRecord() bool {
	return message.BoolFixed(m.Unsafe(), 27, 26)
}

// IsFinalRecord
func (m HistoricalPriceDataTickRecordResponse_IntBuilder) IsFinalRecord() bool {
	return message.BoolFixed(m.Unsafe(), 27, 26)
}

// IsFinalRecord
func (m HistoricalPriceDataTickRecordResponse_IntPointer) IsFinalRecord() bool {
	return message.BoolFixed(m.Ptr, 27, 26)
}

// IsFinalRecord
func (m HistoricalPriceDataTickRecordResponse_IntPointerBuilder) IsFinalRecord() bool {
	return message.BoolFixed(m.Ptr, 27, 26)
}

// SetRequestID
func (m HistoricalPriceDataTickRecordResponse_IntBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRequestID
func (m HistoricalPriceDataTickRecordResponse_IntPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetDateTime
func (m HistoricalPriceDataTickRecordResponse_IntBuilder) SetDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64Fixed(m.Unsafe(), 16, 8, float64(value))
}

// SetDateTime
func (m HistoricalPriceDataTickRecordResponse_IntPointerBuilder) SetDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64Fixed(m.Ptr, 16, 8, float64(value))
}

// SetPrice
func (m HistoricalPriceDataTickRecordResponse_IntBuilder) SetPrice(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 20, 16, value)
}

// SetPrice
func (m HistoricalPriceDataTickRecordResponse_IntPointerBuilder) SetPrice(value int32) {
	message.SetInt32Fixed(m.Ptr, 20, 16, value)
}

// SetVolume
func (m HistoricalPriceDataTickRecordResponse_IntBuilder) SetVolume(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 24, 20, value)
}

// SetVolume
func (m HistoricalPriceDataTickRecordResponse_IntPointerBuilder) SetVolume(value int32) {
	message.SetInt32Fixed(m.Ptr, 24, 20, value)
}

// SetAtBidOrAsk
func (m HistoricalPriceDataTickRecordResponse_IntBuilder) SetAtBidOrAsk(value AtBidOrAskEnum) {
	message.SetUint16Fixed(m.Unsafe(), 26, 24, uint16(value))
}

// SetAtBidOrAsk
func (m HistoricalPriceDataTickRecordResponse_IntPointerBuilder) SetAtBidOrAsk(value AtBidOrAskEnum) {
	message.SetUint16Fixed(m.Ptr, 26, 24, uint16(value))
}

// SetIsFinalRecord
func (m HistoricalPriceDataTickRecordResponse_IntBuilder) SetIsFinalRecord(value bool) {
	message.SetBoolFixed(m.Unsafe(), 27, 26, value)
}

// SetIsFinalRecord
func (m HistoricalPriceDataTickRecordResponse_IntPointerBuilder) SetIsFinalRecord(value bool) {
	message.SetBoolFixed(m.Ptr, 27, 26, value)
}

// Copy
func (m HistoricalPriceDataTickRecordResponse_Int) Copy(to HistoricalPriceDataTickRecordResponse_IntBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetDateTime(m.DateTime())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetIsFinalRecord(m.IsFinalRecord())
}

// Copy
func (m HistoricalPriceDataTickRecordResponse_IntBuilder) Copy(to HistoricalPriceDataTickRecordResponse_IntBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetDateTime(m.DateTime())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetIsFinalRecord(m.IsFinalRecord())
}

// CopyPointer
func (m HistoricalPriceDataTickRecordResponse_Int) CopyPointer(to HistoricalPriceDataTickRecordResponse_IntPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetDateTime(m.DateTime())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetIsFinalRecord(m.IsFinalRecord())
}

// CopyPointer
func (m HistoricalPriceDataTickRecordResponse_IntBuilder) CopyPointer(to HistoricalPriceDataTickRecordResponse_IntPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetDateTime(m.DateTime())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetIsFinalRecord(m.IsFinalRecord())
}

// Copy
func (m HistoricalPriceDataTickRecordResponse_IntPointer) Copy(to HistoricalPriceDataTickRecordResponse_IntBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetDateTime(m.DateTime())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetIsFinalRecord(m.IsFinalRecord())
}

// Copy
func (m HistoricalPriceDataTickRecordResponse_IntPointerBuilder) Copy(to HistoricalPriceDataTickRecordResponse_IntBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetDateTime(m.DateTime())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetIsFinalRecord(m.IsFinalRecord())
}

// CopyPointer
func (m HistoricalPriceDataTickRecordResponse_IntPointer) CopyPointer(to HistoricalPriceDataTickRecordResponse_IntPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetDateTime(m.DateTime())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetIsFinalRecord(m.IsFinalRecord())
}

// CopyPointer
func (m HistoricalPriceDataTickRecordResponse_IntPointerBuilder) CopyPointer(to HistoricalPriceDataTickRecordResponse_IntPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetDateTime(m.DateTime())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetIsFinalRecord(m.IsFinalRecord())
}