// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-16 08:00:54.519198 +0800 WITA m=+0.055446543

package v8

import (
	"github.com/moontrade/dtc-go/message"
)

const HistoricalPriceDataResponseTrailerSize = 16

//     Size                     uint16                       = HistoricalPriceDataResponseTrailerSize  (16)
//     Type                     uint16                       = HISTORICAL_PRICE_DATA_RESPONSE_TRAILER  (807)
//     RequestID                int32                        = 0
//     FinalRecordLastDateTime  DateTimeWithMicrosecondsInt  = 0
var _HistoricalPriceDataResponseTrailerDefault = []byte{16, 0, 39, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type HistoricalPriceDataResponseTrailer struct {
	message.Fixed
}

type HistoricalPriceDataResponseTrailerBuilder struct {
	message.Fixed
}

type HistoricalPriceDataResponseTrailerPointer struct {
	message.FixedPointer
}

type HistoricalPriceDataResponseTrailerPointerBuilder struct {
	message.FixedPointer
}

func NewHistoricalPriceDataResponseTrailerFrom(b []byte) HistoricalPriceDataResponseTrailer {
	return HistoricalPriceDataResponseTrailer{Fixed: message.NewFixed(b)}
}

func WrapHistoricalPriceDataResponseTrailer(b []byte) HistoricalPriceDataResponseTrailer {
	return HistoricalPriceDataResponseTrailer{Fixed: message.WrapFixed(b)}
}

func NewHistoricalPriceDataResponseTrailer() HistoricalPriceDataResponseTrailerBuilder {
	return HistoricalPriceDataResponseTrailerBuilder{message.NewFixed(_HistoricalPriceDataResponseTrailerDefault)}
}

func AllocHistoricalPriceDataResponseTrailer() HistoricalPriceDataResponseTrailerPointerBuilder {
	return HistoricalPriceDataResponseTrailerPointerBuilder{message.AllocFixed(_HistoricalPriceDataResponseTrailerDefault)}
}

func AllocHistoricalPriceDataResponseTrailerFrom(b []byte) HistoricalPriceDataResponseTrailerPointer {
	return HistoricalPriceDataResponseTrailerPointer{FixedPointer: message.AllocFixed(b)}
}

// Clear
//     Size                     uint16                       = HistoricalPriceDataResponseTrailerSize  (16)
//     Type                     uint16                       = HISTORICAL_PRICE_DATA_RESPONSE_TRAILER  (807)
//     RequestID                int32                        = 0
//     FinalRecordLastDateTime  DateTimeWithMicrosecondsInt  = 0
func (m HistoricalPriceDataResponseTrailerBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalPriceDataResponseTrailerDefault)
}

// Clear
//     Size                     uint16                       = HistoricalPriceDataResponseTrailerSize  (16)
//     Type                     uint16                       = HISTORICAL_PRICE_DATA_RESPONSE_TRAILER  (807)
//     RequestID                int32                        = 0
//     FinalRecordLastDateTime  DateTimeWithMicrosecondsInt  = 0
func (m HistoricalPriceDataResponseTrailerPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalPriceDataResponseTrailerDefault)
}

// ToBuilder
func (m HistoricalPriceDataResponseTrailer) ToBuilder() HistoricalPriceDataResponseTrailerBuilder {
	return HistoricalPriceDataResponseTrailerBuilder{m.Fixed}
}

// ToBuilder
func (m HistoricalPriceDataResponseTrailerPointer) ToBuilder() HistoricalPriceDataResponseTrailerPointerBuilder {
	return HistoricalPriceDataResponseTrailerPointerBuilder{m.FixedPointer}
}

// Finish
func (m HistoricalPriceDataResponseTrailerBuilder) Finish() HistoricalPriceDataResponseTrailer {
	return HistoricalPriceDataResponseTrailer{m.Fixed}
}

// Finish
func (m *HistoricalPriceDataResponseTrailerPointerBuilder) Finish() HistoricalPriceDataResponseTrailerPointer {
	return HistoricalPriceDataResponseTrailerPointer{m.FixedPointer}
}

// RequestID
func (m HistoricalPriceDataResponseTrailer) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m HistoricalPriceDataResponseTrailerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m HistoricalPriceDataResponseTrailerPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID
func (m HistoricalPriceDataResponseTrailerPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// FinalRecordLastDateTime
func (m HistoricalPriceDataResponseTrailer) FinalRecordLastDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Unsafe(), 16, 8))
}

// FinalRecordLastDateTime
func (m HistoricalPriceDataResponseTrailerBuilder) FinalRecordLastDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Unsafe(), 16, 8))
}

// FinalRecordLastDateTime
func (m HistoricalPriceDataResponseTrailerPointer) FinalRecordLastDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Ptr, 16, 8))
}

// FinalRecordLastDateTime
func (m HistoricalPriceDataResponseTrailerPointerBuilder) FinalRecordLastDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Ptr, 16, 8))
}

// SetRequestID
func (m HistoricalPriceDataResponseTrailerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRequestID
func (m HistoricalPriceDataResponseTrailerPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetFinalRecordLastDateTime
func (m HistoricalPriceDataResponseTrailerBuilder) SetFinalRecordLastDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64Fixed(m.Unsafe(), 16, 8, int64(value))
}

// SetFinalRecordLastDateTime
func (m HistoricalPriceDataResponseTrailerPointerBuilder) SetFinalRecordLastDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64Fixed(m.Ptr, 16, 8, int64(value))
}