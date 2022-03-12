package sc

import (
	"github.com/moontrade/dtc-go/message"
)

const HeartbeatExtendedSize = 62

// {
//     Size                                   = HeartbeatExtendedSize  (62)
//     Type                                   = HEARTBEAT  (3)
//     NumDroppedMessages                     = 0
//     CurrentDateTime                        = 0
//     SecondsSinceLastReceivedHeartbeat      = 0
//     NumberOfOutstandingSentBuffers         = 0
//     PendingTransmissionDelayInMilliseconds = 0
//     CurrentSendBufferSizeInBytes           = 0
//     SendingDateTimeMicroseconds            = 0
//     DataCompressionRatio                   = 0
//     TotalUncompressedBytes                 = 0
//     TotalCompressionTime                   = 0
//     NumberOfCompressions                   = 0
//     SourceIPAddress                        = 0
// }
var _HeartbeatExtendedDefault = []byte{62, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type HeartbeatExtended struct {
	message.Fixed
}

type HeartbeatExtendedBuilder struct {
	message.Fixed
}

type HeartbeatExtendedPointer struct {
	message.FixedPointer
}

type HeartbeatExtendedPointerBuilder struct {
	message.FixedPointer
}

func NewHeartbeatExtendedFrom(b []byte) HeartbeatExtended {
	return HeartbeatExtended{Fixed: message.NewFixedFrom(b)}
}

func WrapHeartbeatExtended(b []byte) HeartbeatExtended {
	return HeartbeatExtended{Fixed: message.WrapFixed(b)}
}

func NewHeartbeatExtended() HeartbeatExtendedBuilder {
	a := HeartbeatExtendedBuilder{message.NewFixed(62)}
	a.Unsafe().SetBytes(0, _HeartbeatExtendedDefault)
	return a
}

func AllocHeartbeatExtended() HeartbeatExtendedPointerBuilder {
	a := HeartbeatExtendedPointerBuilder{message.AllocFixed(62)}
	a.Ptr.SetBytes(0, _HeartbeatExtendedDefault)
	return a
}

func AllocHeartbeatExtendedFrom(b []byte) HeartbeatExtendedPointer {
	return HeartbeatExtendedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size                                   = HeartbeatExtendedSize  (62)
//     Type                                   = HEARTBEAT  (3)
//     NumDroppedMessages                     = 0
//     CurrentDateTime                        = 0
//     SecondsSinceLastReceivedHeartbeat      = 0
//     NumberOfOutstandingSentBuffers         = 0
//     PendingTransmissionDelayInMilliseconds = 0
//     CurrentSendBufferSizeInBytes           = 0
//     SendingDateTimeMicroseconds            = 0
//     DataCompressionRatio                   = 0
//     TotalUncompressedBytes                 = 0
//     TotalCompressionTime                   = 0
//     NumberOfCompressions                   = 0
//     SourceIPAddress                        = 0
// }
func (m HeartbeatExtendedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HeartbeatExtendedDefault)
}

// Clear
// {
//     Size                                   = HeartbeatExtendedSize  (62)
//     Type                                   = HEARTBEAT  (3)
//     NumDroppedMessages                     = 0
//     CurrentDateTime                        = 0
//     SecondsSinceLastReceivedHeartbeat      = 0
//     NumberOfOutstandingSentBuffers         = 0
//     PendingTransmissionDelayInMilliseconds = 0
//     CurrentSendBufferSizeInBytes           = 0
//     SendingDateTimeMicroseconds            = 0
//     DataCompressionRatio                   = 0
//     TotalUncompressedBytes                 = 0
//     TotalCompressionTime                   = 0
//     NumberOfCompressions                   = 0
//     SourceIPAddress                        = 0
// }
func (m HeartbeatExtendedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HeartbeatExtendedDefault)
}

// ToBuilder
func (m HeartbeatExtended) ToBuilder() HeartbeatExtendedBuilder {
	return HeartbeatExtendedBuilder{m.Fixed}
}

// ToBuilder
func (m HeartbeatExtendedPointer) ToBuilder() HeartbeatExtendedPointerBuilder {
	return HeartbeatExtendedPointerBuilder{m.FixedPointer}
}

// Finish
func (m HeartbeatExtendedBuilder) Finish() HeartbeatExtended {
	return HeartbeatExtended{m.Fixed}
}

// Finish
func (m *HeartbeatExtendedPointerBuilder) Finish() HeartbeatExtendedPointer {
	return HeartbeatExtendedPointer{m.FixedPointer}
}

// NumDroppedMessages
func (m HeartbeatExtended) NumDroppedMessages() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// NumDroppedMessages
func (m HeartbeatExtendedBuilder) NumDroppedMessages() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// NumDroppedMessages
func (m HeartbeatExtendedPointer) NumDroppedMessages() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// NumDroppedMessages
func (m HeartbeatExtendedPointerBuilder) NumDroppedMessages() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// CurrentDateTime
func (m HeartbeatExtended) CurrentDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Unsafe(), 16, 8))
}

// CurrentDateTime
func (m HeartbeatExtendedBuilder) CurrentDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Unsafe(), 16, 8))
}

// CurrentDateTime
func (m HeartbeatExtendedPointer) CurrentDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 16, 8))
}

// CurrentDateTime
func (m HeartbeatExtendedPointerBuilder) CurrentDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 16, 8))
}

// SecondsSinceLastReceivedHeartbeat
func (m HeartbeatExtended) SecondsSinceLastReceivedHeartbeat() uint16 {
	return message.Uint16Fixed(m.Unsafe(), 18, 16)
}

// SecondsSinceLastReceivedHeartbeat
func (m HeartbeatExtendedBuilder) SecondsSinceLastReceivedHeartbeat() uint16 {
	return message.Uint16Fixed(m.Unsafe(), 18, 16)
}

// SecondsSinceLastReceivedHeartbeat
func (m HeartbeatExtendedPointer) SecondsSinceLastReceivedHeartbeat() uint16 {
	return message.Uint16Fixed(m.Ptr, 18, 16)
}

// SecondsSinceLastReceivedHeartbeat
func (m HeartbeatExtendedPointerBuilder) SecondsSinceLastReceivedHeartbeat() uint16 {
	return message.Uint16Fixed(m.Ptr, 18, 16)
}

// NumberOfOutstandingSentBuffers
func (m HeartbeatExtended) NumberOfOutstandingSentBuffers() uint16 {
	return message.Uint16Fixed(m.Unsafe(), 20, 18)
}

// NumberOfOutstandingSentBuffers
func (m HeartbeatExtendedBuilder) NumberOfOutstandingSentBuffers() uint16 {
	return message.Uint16Fixed(m.Unsafe(), 20, 18)
}

// NumberOfOutstandingSentBuffers
func (m HeartbeatExtendedPointer) NumberOfOutstandingSentBuffers() uint16 {
	return message.Uint16Fixed(m.Ptr, 20, 18)
}

// NumberOfOutstandingSentBuffers
func (m HeartbeatExtendedPointerBuilder) NumberOfOutstandingSentBuffers() uint16 {
	return message.Uint16Fixed(m.Ptr, 20, 18)
}

// PendingTransmissionDelayInMilliseconds
func (m HeartbeatExtended) PendingTransmissionDelayInMilliseconds() uint16 {
	return message.Uint16Fixed(m.Unsafe(), 22, 20)
}

// PendingTransmissionDelayInMilliseconds
func (m HeartbeatExtendedBuilder) PendingTransmissionDelayInMilliseconds() uint16 {
	return message.Uint16Fixed(m.Unsafe(), 22, 20)
}

// PendingTransmissionDelayInMilliseconds
func (m HeartbeatExtendedPointer) PendingTransmissionDelayInMilliseconds() uint16 {
	return message.Uint16Fixed(m.Ptr, 22, 20)
}

// PendingTransmissionDelayInMilliseconds
func (m HeartbeatExtendedPointerBuilder) PendingTransmissionDelayInMilliseconds() uint16 {
	return message.Uint16Fixed(m.Ptr, 22, 20)
}

// CurrentSendBufferSizeInBytes
func (m HeartbeatExtended) CurrentSendBufferSizeInBytes() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 26, 22)
}

// CurrentSendBufferSizeInBytes
func (m HeartbeatExtendedBuilder) CurrentSendBufferSizeInBytes() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 26, 22)
}

// CurrentSendBufferSizeInBytes
func (m HeartbeatExtendedPointer) CurrentSendBufferSizeInBytes() uint32 {
	return message.Uint32Fixed(m.Ptr, 26, 22)
}

// CurrentSendBufferSizeInBytes
func (m HeartbeatExtendedPointerBuilder) CurrentSendBufferSizeInBytes() uint32 {
	return message.Uint32Fixed(m.Ptr, 26, 22)
}

// SendingDateTimeMicroseconds
func (m HeartbeatExtended) SendingDateTimeMicroseconds() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Unsafe(), 34, 26))
}

// SendingDateTimeMicroseconds
func (m HeartbeatExtendedBuilder) SendingDateTimeMicroseconds() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Unsafe(), 34, 26))
}

// SendingDateTimeMicroseconds
func (m HeartbeatExtendedPointer) SendingDateTimeMicroseconds() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Ptr, 34, 26))
}

// SendingDateTimeMicroseconds
func (m HeartbeatExtendedPointerBuilder) SendingDateTimeMicroseconds() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Ptr, 34, 26))
}

// DataCompressionRatio
func (m HeartbeatExtended) DataCompressionRatio() float32 {
	return message.Float32Fixed(m.Unsafe(), 38, 34)
}

// DataCompressionRatio
func (m HeartbeatExtendedBuilder) DataCompressionRatio() float32 {
	return message.Float32Fixed(m.Unsafe(), 38, 34)
}

// DataCompressionRatio
func (m HeartbeatExtendedPointer) DataCompressionRatio() float32 {
	return message.Float32Fixed(m.Ptr, 38, 34)
}

// DataCompressionRatio
func (m HeartbeatExtendedPointerBuilder) DataCompressionRatio() float32 {
	return message.Float32Fixed(m.Ptr, 38, 34)
}

// TotalUncompressedBytes
func (m HeartbeatExtended) TotalUncompressedBytes() uint64 {
	return message.Uint64Fixed(m.Unsafe(), 46, 38)
}

// TotalUncompressedBytes
func (m HeartbeatExtendedBuilder) TotalUncompressedBytes() uint64 {
	return message.Uint64Fixed(m.Unsafe(), 46, 38)
}

// TotalUncompressedBytes
func (m HeartbeatExtendedPointer) TotalUncompressedBytes() uint64 {
	return message.Uint64Fixed(m.Ptr, 46, 38)
}

// TotalUncompressedBytes
func (m HeartbeatExtendedPointerBuilder) TotalUncompressedBytes() uint64 {
	return message.Uint64Fixed(m.Ptr, 46, 38)
}

// TotalCompressionTime
func (m HeartbeatExtended) TotalCompressionTime() float64 {
	return message.Float64Fixed(m.Unsafe(), 54, 46)
}

// TotalCompressionTime
func (m HeartbeatExtendedBuilder) TotalCompressionTime() float64 {
	return message.Float64Fixed(m.Unsafe(), 54, 46)
}

// TotalCompressionTime
func (m HeartbeatExtendedPointer) TotalCompressionTime() float64 {
	return message.Float64Fixed(m.Ptr, 54, 46)
}

// TotalCompressionTime
func (m HeartbeatExtendedPointerBuilder) TotalCompressionTime() float64 {
	return message.Float64Fixed(m.Ptr, 54, 46)
}

// NumberOfCompressions
func (m HeartbeatExtended) NumberOfCompressions() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 58, 54)
}

// NumberOfCompressions
func (m HeartbeatExtendedBuilder) NumberOfCompressions() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 58, 54)
}

// NumberOfCompressions
func (m HeartbeatExtendedPointer) NumberOfCompressions() uint32 {
	return message.Uint32Fixed(m.Ptr, 58, 54)
}

// NumberOfCompressions
func (m HeartbeatExtendedPointerBuilder) NumberOfCompressions() uint32 {
	return message.Uint32Fixed(m.Ptr, 58, 54)
}

// SourceIPAddress
func (m HeartbeatExtended) SourceIPAddress() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 62, 58)
}

// SourceIPAddress
func (m HeartbeatExtendedBuilder) SourceIPAddress() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 62, 58)
}

// SourceIPAddress
func (m HeartbeatExtendedPointer) SourceIPAddress() uint32 {
	return message.Uint32Fixed(m.Ptr, 62, 58)
}

// SourceIPAddress
func (m HeartbeatExtendedPointerBuilder) SourceIPAddress() uint32 {
	return message.Uint32Fixed(m.Ptr, 62, 58)
}

// SetNumDroppedMessages
func (m HeartbeatExtendedBuilder) SetNumDroppedMessages(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetNumDroppedMessages
func (m HeartbeatExtendedPointerBuilder) SetNumDroppedMessages(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetCurrentDateTime
func (m HeartbeatExtendedBuilder) SetCurrentDateTime(value DateTime) {
	message.SetInt64Fixed(m.Unsafe(), 16, 8, int64(value))
}

// SetCurrentDateTime
func (m HeartbeatExtendedPointerBuilder) SetCurrentDateTime(value DateTime) {
	message.SetInt64Fixed(m.Ptr, 16, 8, int64(value))
}

// SetSecondsSinceLastReceivedHeartbeat
func (m HeartbeatExtendedBuilder) SetSecondsSinceLastReceivedHeartbeat(value uint16) {
	message.SetUint16Fixed(m.Unsafe(), 18, 16, value)
}

// SetSecondsSinceLastReceivedHeartbeat
func (m HeartbeatExtendedPointerBuilder) SetSecondsSinceLastReceivedHeartbeat(value uint16) {
	message.SetUint16Fixed(m.Ptr, 18, 16, value)
}

// SetNumberOfOutstandingSentBuffers
func (m HeartbeatExtendedBuilder) SetNumberOfOutstandingSentBuffers(value uint16) {
	message.SetUint16Fixed(m.Unsafe(), 20, 18, value)
}

// SetNumberOfOutstandingSentBuffers
func (m HeartbeatExtendedPointerBuilder) SetNumberOfOutstandingSentBuffers(value uint16) {
	message.SetUint16Fixed(m.Ptr, 20, 18, value)
}

// SetPendingTransmissionDelayInMilliseconds
func (m HeartbeatExtendedBuilder) SetPendingTransmissionDelayInMilliseconds(value uint16) {
	message.SetUint16Fixed(m.Unsafe(), 22, 20, value)
}

// SetPendingTransmissionDelayInMilliseconds
func (m HeartbeatExtendedPointerBuilder) SetPendingTransmissionDelayInMilliseconds(value uint16) {
	message.SetUint16Fixed(m.Ptr, 22, 20, value)
}

// SetCurrentSendBufferSizeInBytes
func (m HeartbeatExtendedBuilder) SetCurrentSendBufferSizeInBytes(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 26, 22, value)
}

// SetCurrentSendBufferSizeInBytes
func (m HeartbeatExtendedPointerBuilder) SetCurrentSendBufferSizeInBytes(value uint32) {
	message.SetUint32Fixed(m.Ptr, 26, 22, value)
}

// SetSendingDateTimeMicroseconds
func (m HeartbeatExtendedBuilder) SetSendingDateTimeMicroseconds(value DateTimeWithMicrosecondsInt) {
	message.SetInt64Fixed(m.Unsafe(), 34, 26, int64(value))
}

// SetSendingDateTimeMicroseconds
func (m HeartbeatExtendedPointerBuilder) SetSendingDateTimeMicroseconds(value DateTimeWithMicrosecondsInt) {
	message.SetInt64Fixed(m.Ptr, 34, 26, int64(value))
}

// SetDataCompressionRatio
func (m HeartbeatExtendedBuilder) SetDataCompressionRatio(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 38, 34, value)
}

// SetDataCompressionRatio
func (m HeartbeatExtendedPointerBuilder) SetDataCompressionRatio(value float32) {
	message.SetFloat32Fixed(m.Ptr, 38, 34, value)
}

// SetTotalUncompressedBytes
func (m HeartbeatExtendedBuilder) SetTotalUncompressedBytes(value uint64) {
	message.SetUint64Fixed(m.Unsafe(), 46, 38, value)
}

// SetTotalUncompressedBytes
func (m HeartbeatExtendedPointerBuilder) SetTotalUncompressedBytes(value uint64) {
	message.SetUint64Fixed(m.Ptr, 46, 38, value)
}

// SetTotalCompressionTime
func (m HeartbeatExtendedBuilder) SetTotalCompressionTime(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 54, 46, value)
}

// SetTotalCompressionTime
func (m HeartbeatExtendedPointerBuilder) SetTotalCompressionTime(value float64) {
	message.SetFloat64Fixed(m.Ptr, 54, 46, value)
}

// SetNumberOfCompressions
func (m HeartbeatExtendedBuilder) SetNumberOfCompressions(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 58, 54, value)
}

// SetNumberOfCompressions
func (m HeartbeatExtendedPointerBuilder) SetNumberOfCompressions(value uint32) {
	message.SetUint32Fixed(m.Ptr, 58, 54, value)
}

// SetSourceIPAddress
func (m HeartbeatExtendedBuilder) SetSourceIPAddress(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 62, 58, value)
}

// SetSourceIPAddress
func (m HeartbeatExtendedPointerBuilder) SetSourceIPAddress(value uint32) {
	message.SetUint32Fixed(m.Ptr, 62, 58, value)
}

// Copy
func (m HeartbeatExtended) Copy(to HeartbeatExtendedBuilder) {
	to.SetNumDroppedMessages(m.NumDroppedMessages())
	to.SetCurrentDateTime(m.CurrentDateTime())
	to.SetSecondsSinceLastReceivedHeartbeat(m.SecondsSinceLastReceivedHeartbeat())
	to.SetNumberOfOutstandingSentBuffers(m.NumberOfOutstandingSentBuffers())
	to.SetPendingTransmissionDelayInMilliseconds(m.PendingTransmissionDelayInMilliseconds())
	to.SetCurrentSendBufferSizeInBytes(m.CurrentSendBufferSizeInBytes())
	to.SetSendingDateTimeMicroseconds(m.SendingDateTimeMicroseconds())
	to.SetDataCompressionRatio(m.DataCompressionRatio())
	to.SetTotalUncompressedBytes(m.TotalUncompressedBytes())
	to.SetTotalCompressionTime(m.TotalCompressionTime())
	to.SetNumberOfCompressions(m.NumberOfCompressions())
	to.SetSourceIPAddress(m.SourceIPAddress())
}

// Copy
func (m HeartbeatExtendedBuilder) Copy(to HeartbeatExtendedBuilder) {
	to.SetNumDroppedMessages(m.NumDroppedMessages())
	to.SetCurrentDateTime(m.CurrentDateTime())
	to.SetSecondsSinceLastReceivedHeartbeat(m.SecondsSinceLastReceivedHeartbeat())
	to.SetNumberOfOutstandingSentBuffers(m.NumberOfOutstandingSentBuffers())
	to.SetPendingTransmissionDelayInMilliseconds(m.PendingTransmissionDelayInMilliseconds())
	to.SetCurrentSendBufferSizeInBytes(m.CurrentSendBufferSizeInBytes())
	to.SetSendingDateTimeMicroseconds(m.SendingDateTimeMicroseconds())
	to.SetDataCompressionRatio(m.DataCompressionRatio())
	to.SetTotalUncompressedBytes(m.TotalUncompressedBytes())
	to.SetTotalCompressionTime(m.TotalCompressionTime())
	to.SetNumberOfCompressions(m.NumberOfCompressions())
	to.SetSourceIPAddress(m.SourceIPAddress())
}

// CopyPointer
func (m HeartbeatExtended) CopyPointer(to HeartbeatExtendedPointerBuilder) {
	to.SetNumDroppedMessages(m.NumDroppedMessages())
	to.SetCurrentDateTime(m.CurrentDateTime())
	to.SetSecondsSinceLastReceivedHeartbeat(m.SecondsSinceLastReceivedHeartbeat())
	to.SetNumberOfOutstandingSentBuffers(m.NumberOfOutstandingSentBuffers())
	to.SetPendingTransmissionDelayInMilliseconds(m.PendingTransmissionDelayInMilliseconds())
	to.SetCurrentSendBufferSizeInBytes(m.CurrentSendBufferSizeInBytes())
	to.SetSendingDateTimeMicroseconds(m.SendingDateTimeMicroseconds())
	to.SetDataCompressionRatio(m.DataCompressionRatio())
	to.SetTotalUncompressedBytes(m.TotalUncompressedBytes())
	to.SetTotalCompressionTime(m.TotalCompressionTime())
	to.SetNumberOfCompressions(m.NumberOfCompressions())
	to.SetSourceIPAddress(m.SourceIPAddress())
}

// CopyPointer
func (m HeartbeatExtendedBuilder) CopyPointer(to HeartbeatExtendedPointerBuilder) {
	to.SetNumDroppedMessages(m.NumDroppedMessages())
	to.SetCurrentDateTime(m.CurrentDateTime())
	to.SetSecondsSinceLastReceivedHeartbeat(m.SecondsSinceLastReceivedHeartbeat())
	to.SetNumberOfOutstandingSentBuffers(m.NumberOfOutstandingSentBuffers())
	to.SetPendingTransmissionDelayInMilliseconds(m.PendingTransmissionDelayInMilliseconds())
	to.SetCurrentSendBufferSizeInBytes(m.CurrentSendBufferSizeInBytes())
	to.SetSendingDateTimeMicroseconds(m.SendingDateTimeMicroseconds())
	to.SetDataCompressionRatio(m.DataCompressionRatio())
	to.SetTotalUncompressedBytes(m.TotalUncompressedBytes())
	to.SetTotalCompressionTime(m.TotalCompressionTime())
	to.SetNumberOfCompressions(m.NumberOfCompressions())
	to.SetSourceIPAddress(m.SourceIPAddress())
}

// Copy
func (m HeartbeatExtendedPointer) Copy(to HeartbeatExtendedBuilder) {
	to.SetNumDroppedMessages(m.NumDroppedMessages())
	to.SetCurrentDateTime(m.CurrentDateTime())
	to.SetSecondsSinceLastReceivedHeartbeat(m.SecondsSinceLastReceivedHeartbeat())
	to.SetNumberOfOutstandingSentBuffers(m.NumberOfOutstandingSentBuffers())
	to.SetPendingTransmissionDelayInMilliseconds(m.PendingTransmissionDelayInMilliseconds())
	to.SetCurrentSendBufferSizeInBytes(m.CurrentSendBufferSizeInBytes())
	to.SetSendingDateTimeMicroseconds(m.SendingDateTimeMicroseconds())
	to.SetDataCompressionRatio(m.DataCompressionRatio())
	to.SetTotalUncompressedBytes(m.TotalUncompressedBytes())
	to.SetTotalCompressionTime(m.TotalCompressionTime())
	to.SetNumberOfCompressions(m.NumberOfCompressions())
	to.SetSourceIPAddress(m.SourceIPAddress())
}

// Copy
func (m HeartbeatExtendedPointerBuilder) Copy(to HeartbeatExtendedBuilder) {
	to.SetNumDroppedMessages(m.NumDroppedMessages())
	to.SetCurrentDateTime(m.CurrentDateTime())
	to.SetSecondsSinceLastReceivedHeartbeat(m.SecondsSinceLastReceivedHeartbeat())
	to.SetNumberOfOutstandingSentBuffers(m.NumberOfOutstandingSentBuffers())
	to.SetPendingTransmissionDelayInMilliseconds(m.PendingTransmissionDelayInMilliseconds())
	to.SetCurrentSendBufferSizeInBytes(m.CurrentSendBufferSizeInBytes())
	to.SetSendingDateTimeMicroseconds(m.SendingDateTimeMicroseconds())
	to.SetDataCompressionRatio(m.DataCompressionRatio())
	to.SetTotalUncompressedBytes(m.TotalUncompressedBytes())
	to.SetTotalCompressionTime(m.TotalCompressionTime())
	to.SetNumberOfCompressions(m.NumberOfCompressions())
	to.SetSourceIPAddress(m.SourceIPAddress())
}

// CopyPointer
func (m HeartbeatExtendedPointer) CopyPointer(to HeartbeatExtendedPointerBuilder) {
	to.SetNumDroppedMessages(m.NumDroppedMessages())
	to.SetCurrentDateTime(m.CurrentDateTime())
	to.SetSecondsSinceLastReceivedHeartbeat(m.SecondsSinceLastReceivedHeartbeat())
	to.SetNumberOfOutstandingSentBuffers(m.NumberOfOutstandingSentBuffers())
	to.SetPendingTransmissionDelayInMilliseconds(m.PendingTransmissionDelayInMilliseconds())
	to.SetCurrentSendBufferSizeInBytes(m.CurrentSendBufferSizeInBytes())
	to.SetSendingDateTimeMicroseconds(m.SendingDateTimeMicroseconds())
	to.SetDataCompressionRatio(m.DataCompressionRatio())
	to.SetTotalUncompressedBytes(m.TotalUncompressedBytes())
	to.SetTotalCompressionTime(m.TotalCompressionTime())
	to.SetNumberOfCompressions(m.NumberOfCompressions())
	to.SetSourceIPAddress(m.SourceIPAddress())
}

// CopyPointer
func (m HeartbeatExtendedPointerBuilder) CopyPointer(to HeartbeatExtendedPointerBuilder) {
	to.SetNumDroppedMessages(m.NumDroppedMessages())
	to.SetCurrentDateTime(m.CurrentDateTime())
	to.SetSecondsSinceLastReceivedHeartbeat(m.SecondsSinceLastReceivedHeartbeat())
	to.SetNumberOfOutstandingSentBuffers(m.NumberOfOutstandingSentBuffers())
	to.SetPendingTransmissionDelayInMilliseconds(m.PendingTransmissionDelayInMilliseconds())
	to.SetCurrentSendBufferSizeInBytes(m.CurrentSendBufferSizeInBytes())
	to.SetSendingDateTimeMicroseconds(m.SendingDateTimeMicroseconds())
	to.SetDataCompressionRatio(m.DataCompressionRatio())
	to.SetTotalUncompressedBytes(m.TotalUncompressedBytes())
	to.SetTotalCompressionTime(m.TotalCompressionTime())
	to.SetNumberOfCompressions(m.NumberOfCompressions())
	to.SetSourceIPAddress(m.SourceIPAddress())
}