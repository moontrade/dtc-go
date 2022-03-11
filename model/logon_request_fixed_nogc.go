package model

import (
	"github.com/moontrade/dtc-go/message"
)

// LogonRequestFixedPointer The LogonRequest message is sent from the Client to the Server requesting
// to logon to the Server.
//
// This is the very first message the Client sends to the Server before being
// allowed to send any other message other than the EncodingRequest.
type LogonRequestFixedPointer struct {
	message.FixedPointer
}

// LogonRequestFixedPointerBuilder The LogonRequest message is sent from the Client to the Server requesting
// to logon to the Server.
//
// This is the very first message the Client sends to the Server before being
// allowed to send any other message other than the EncodingRequest.
type LogonRequestFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocLogonRequestFixed() LogonRequestFixedPointerBuilder {
	a := LogonRequestFixedPointerBuilder{message.AllocFixed(284)}
	a.Ptr.SetBytes(0, _LogonRequestFixedDefault)
	return a
}

func AllocLogonRequestFixedFrom(b []byte) LogonRequestFixedPointer {
	return LogonRequestFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size                           = LogonRequestFixedSize  (284)
//     Type                           = LOGON_REQUEST  (1)
//     ProtocolVersion                = CURRENT_VERSION  (8)
//     Username                       = ""
//     Password                       = ""
//     GeneralTextData                = ""
//     Integer_1                      = 0
//     Integer_2                      = 0
//     HeartbeatIntervalInSeconds     = 0
//     TradeMode                      = 0
//     TradeAccount                   = ""
//     HardwareIdentifier             = ""
//     ClientName                     = ""
//     MarketDataTransmissionInterval = 0
// }
func (m LogonRequestFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _LogonRequestFixedDefault)
}

// ToBuilder
func (m LogonRequestFixedPointer) ToBuilder() LogonRequestFixedPointerBuilder {
	return LogonRequestFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *LogonRequestFixedPointerBuilder) Finish() LogonRequestFixedPointer {
	return LogonRequestFixedPointer{m.FixedPointer}
}

// ProtocolVersion The protocol version supported by the Client. Automatically set by constructor.
// The protocol version supported by the Client. Automatically set by constructor.
func (m LogonRequestFixedPointer) ProtocolVersion() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// ProtocolVersion The protocol version supported by the Client. Automatically set by constructor.
// The protocol version supported by the Client. Automatically set by constructor.
func (m LogonRequestFixedPointerBuilder) ProtocolVersion() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// Username Optional username for the server to authenticate the Client.
func (m LogonRequestFixedPointer) Username() string {
	return message.StringFixed(m.Ptr, 40, 8)
}

// Username Optional username for the server to authenticate the Client.
func (m LogonRequestFixedPointerBuilder) Username() string {
	return message.StringFixed(m.Ptr, 40, 8)
}

// Password Optional password for the server to authenticate the Client.
func (m LogonRequestFixedPointer) Password() string {
	return message.StringFixed(m.Ptr, 72, 40)
}

// Password Optional password for the server to authenticate the Client.
func (m LogonRequestFixedPointerBuilder) Password() string {
	return message.StringFixed(m.Ptr, 72, 40)
}

// GeneralTextData Optional general-purpose text string. For example, this could be used
// to pass a license key that the Server may require.
func (m LogonRequestFixedPointer) GeneralTextData() string {
	return message.StringFixed(m.Ptr, 136, 72)
}

// GeneralTextData Optional general-purpose text string. For example, this could be used
// to pass a license key that the Server may require.
func (m LogonRequestFixedPointerBuilder) GeneralTextData() string {
	return message.StringFixed(m.Ptr, 136, 72)
}

// Integer_1 Optional. General-purpose integer.
func (m LogonRequestFixedPointer) Integer_1() int32 {
	return message.Int32Fixed(m.Ptr, 140, 136)
}

// Integer_1 Optional. General-purpose integer.
func (m LogonRequestFixedPointerBuilder) Integer_1() int32 {
	return message.Int32Fixed(m.Ptr, 140, 136)
}

// Integer_2 Optional. General-purpose integer.
func (m LogonRequestFixedPointer) Integer_2() int32 {
	return message.Int32Fixed(m.Ptr, 144, 140)
}

// Integer_2 Optional. General-purpose integer.
func (m LogonRequestFixedPointerBuilder) Integer_2() int32 {
	return message.Int32Fixed(m.Ptr, 144, 140)
}

// HeartbeatIntervalInSeconds The interval in seconds that each side, the Client and the Server, needs
// to use to send Heartbeat messages to the other side.
//
// This should be a value from anywhere from 5 to 60 seconds.
//
// This field is required.
func (m LogonRequestFixedPointer) HeartbeatIntervalInSeconds() int32 {
	return message.Int32Fixed(m.Ptr, 148, 144)
}

// HeartbeatIntervalInSeconds The interval in seconds that each side, the Client and the Server, needs
// to use to send Heartbeat messages to the other side.
//
// This should be a value from anywhere from 5 to 60 seconds.
//
// This field is required.
func (m LogonRequestFixedPointerBuilder) HeartbeatIntervalInSeconds() int32 {
	return message.Int32Fixed(m.Ptr, 148, 144)
}

// TradeMode
func (m LogonRequestFixedPointer) TradeMode() TradeModeEnum {
	return TradeModeEnum(message.Int32Fixed(m.Ptr, 152, 148))
}

// TradeMode
func (m LogonRequestFixedPointerBuilder) TradeMode() TradeModeEnum {
	return TradeModeEnum(message.Int32Fixed(m.Ptr, 152, 148))
}

// TradeAccount This is an optional field and this should only be set to a Trade Account
// identifier if that is required to logon by the Server. this would only
// be implemented in rare cases. Usually this would be the case if the logon
// is bound to a particular Trade Account and not changeable after the log
// in.
//
// The server is still required to implement the TradeAccountsRequest and
// TradeAccountResponsemessages.
func (m LogonRequestFixedPointer) TradeAccount() string {
	return message.StringFixed(m.Ptr, 184, 152)
}

// TradeAccount This is an optional field and this should only be set to a Trade Account
// identifier if that is required to logon by the Server. this would only
// be implemented in rare cases. Usually this would be the case if the logon
// is bound to a particular Trade Account and not changeable after the log
// in.
//
// The server is still required to implement the TradeAccountsRequest and
// TradeAccountResponsemessages.
func (m LogonRequestFixedPointerBuilder) TradeAccount() string {
	return message.StringFixed(m.Ptr, 184, 152)
}

// HardwareIdentifier Optional: This is the computer hardware identifier. The intention of this
// is that this will be implemented by the Client program developer on a
// case-by-case basis for specific Data/Trading service providers. It will
// be a reasonable implementation to uniquely identify a system and will
// not be publicly disclosed. It will never contain personally identifiable
// information.
func (m LogonRequestFixedPointer) HardwareIdentifier() string {
	return message.StringFixed(m.Ptr, 248, 184)
}

// HardwareIdentifier Optional: This is the computer hardware identifier. The intention of this
// is that this will be implemented by the Client program developer on a
// case-by-case basis for specific Data/Trading service providers. It will
// be a reasonable implementation to uniquely identify a system and will
// not be publicly disclosed. It will never contain personally identifiable
// information.
func (m LogonRequestFixedPointerBuilder) HardwareIdentifier() string {
	return message.StringFixed(m.Ptr, 248, 184)
}

// ClientName The Client name. This is a free-form text string.
func (m LogonRequestFixedPointer) ClientName() string {
	return message.StringFixed(m.Ptr, 280, 248)
}

// ClientName The Client name. This is a free-form text string.
func (m LogonRequestFixedPointerBuilder) ClientName() string {
	return message.StringFixed(m.Ptr, 280, 248)
}

// MarketDataTransmissionInterval This is an optional field to be used by the Server which specifies in
// milliseconds, the delay with transmitting market data to the Client.
//
// For reasons of efficiency, the server may buffer data over this timeframe,
// and send data after this time frame expires.
func (m LogonRequestFixedPointer) MarketDataTransmissionInterval() int32 {
	return message.Int32Fixed(m.Ptr, 284, 280)
}

// MarketDataTransmissionInterval This is an optional field to be used by the Server which specifies in
// milliseconds, the delay with transmitting market data to the Client.
//
// For reasons of efficiency, the server may buffer data over this timeframe,
// and send data after this time frame expires.
func (m LogonRequestFixedPointerBuilder) MarketDataTransmissionInterval() int32 {
	return message.Int32Fixed(m.Ptr, 284, 280)
}

// SetProtocolVersion The protocol version supported by the Client. Automatically set by constructor.
// The protocol version supported by the Client. Automatically set by constructor.
func (m LogonRequestFixedPointerBuilder) SetProtocolVersion(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetUsername Optional username for the server to authenticate the Client.
func (m LogonRequestFixedPointerBuilder) SetUsername(value string) {
	message.SetStringFixed(m.Ptr, 40, 8, value)
}

// SetPassword Optional password for the server to authenticate the Client.
func (m LogonRequestFixedPointerBuilder) SetPassword(value string) {
	message.SetStringFixed(m.Ptr, 72, 40, value)
}

// SetGeneralTextData Optional general-purpose text string. For example, this could be used
// to pass a license key that the Server may require.
func (m LogonRequestFixedPointerBuilder) SetGeneralTextData(value string) {
	message.SetStringFixed(m.Ptr, 136, 72, value)
}

// SetInteger_1 Optional. General-purpose integer.
func (m LogonRequestFixedPointerBuilder) SetInteger_1(value int32) {
	message.SetInt32Fixed(m.Ptr, 140, 136, value)
}

// SetInteger_2 Optional. General-purpose integer.
func (m LogonRequestFixedPointerBuilder) SetInteger_2(value int32) {
	message.SetInt32Fixed(m.Ptr, 144, 140, value)
}

// SetHeartbeatIntervalInSeconds The interval in seconds that each side, the Client and the Server, needs
// to use to send Heartbeat messages to the other side.
//
// This should be a value from anywhere from 5 to 60 seconds.
//
// This field is required.
func (m LogonRequestFixedPointerBuilder) SetHeartbeatIntervalInSeconds(value int32) {
	message.SetInt32Fixed(m.Ptr, 148, 144, value)
}

// SetTradeMode
func (m LogonRequestFixedPointerBuilder) SetTradeMode(value TradeModeEnum) {
	message.SetInt32Fixed(m.Ptr, 152, 148, int32(value))
}

// SetTradeAccount This is an optional field and this should only be set to a Trade Account
// identifier if that is required to logon by the Server. this would only
// be implemented in rare cases. Usually this would be the case if the logon
// is bound to a particular Trade Account and not changeable after the log
// in.
//
// The server is still required to implement the TradeAccountsRequest and
// TradeAccountResponsemessages.
func (m LogonRequestFixedPointerBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Ptr, 184, 152, value)
}

// SetHardwareIdentifier Optional: This is the computer hardware identifier. The intention of this
// is that this will be implemented by the Client program developer on a
// case-by-case basis for specific Data/Trading service providers. It will
// be a reasonable implementation to uniquely identify a system and will
// not be publicly disclosed. It will never contain personally identifiable
// information.
func (m LogonRequestFixedPointerBuilder) SetHardwareIdentifier(value string) {
	message.SetStringFixed(m.Ptr, 248, 184, value)
}

// SetClientName The Client name. This is a free-form text string.
func (m LogonRequestFixedPointerBuilder) SetClientName(value string) {
	message.SetStringFixed(m.Ptr, 280, 248, value)
}

// SetMarketDataTransmissionInterval This is an optional field to be used by the Server which specifies in
// milliseconds, the delay with transmitting market data to the Client.
//
// For reasons of efficiency, the server may buffer data over this timeframe,
// and send data after this time frame expires.
func (m LogonRequestFixedPointerBuilder) SetMarketDataTransmissionInterval(value int32) {
	message.SetInt32Fixed(m.Ptr, 284, 280, value)
}

// Copy
func (m LogonRequestFixedPointer) Copy(to LogonRequestFixedBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetUsername(m.Username())
	to.SetPassword(m.Password())
	to.SetGeneralTextData(m.GeneralTextData())
	to.SetInteger_1(m.Integer_1())
	to.SetInteger_2(m.Integer_2())
	to.SetHeartbeatIntervalInSeconds(m.HeartbeatIntervalInSeconds())
	to.SetTradeMode(m.TradeMode())
	to.SetTradeAccount(m.TradeAccount())
	to.SetHardwareIdentifier(m.HardwareIdentifier())
	to.SetClientName(m.ClientName())
	to.SetMarketDataTransmissionInterval(m.MarketDataTransmissionInterval())
}

// Copy
func (m LogonRequestFixedPointerBuilder) Copy(to LogonRequestFixedBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetUsername(m.Username())
	to.SetPassword(m.Password())
	to.SetGeneralTextData(m.GeneralTextData())
	to.SetInteger_1(m.Integer_1())
	to.SetInteger_2(m.Integer_2())
	to.SetHeartbeatIntervalInSeconds(m.HeartbeatIntervalInSeconds())
	to.SetTradeMode(m.TradeMode())
	to.SetTradeAccount(m.TradeAccount())
	to.SetHardwareIdentifier(m.HardwareIdentifier())
	to.SetClientName(m.ClientName())
	to.SetMarketDataTransmissionInterval(m.MarketDataTransmissionInterval())
}

// CopyPointer
func (m LogonRequestFixedPointer) CopyPointer(to LogonRequestFixedPointerBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetUsername(m.Username())
	to.SetPassword(m.Password())
	to.SetGeneralTextData(m.GeneralTextData())
	to.SetInteger_1(m.Integer_1())
	to.SetInteger_2(m.Integer_2())
	to.SetHeartbeatIntervalInSeconds(m.HeartbeatIntervalInSeconds())
	to.SetTradeMode(m.TradeMode())
	to.SetTradeAccount(m.TradeAccount())
	to.SetHardwareIdentifier(m.HardwareIdentifier())
	to.SetClientName(m.ClientName())
	to.SetMarketDataTransmissionInterval(m.MarketDataTransmissionInterval())
}

// CopyPointer
func (m LogonRequestFixedPointerBuilder) CopyPointer(to LogonRequestFixedPointerBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetUsername(m.Username())
	to.SetPassword(m.Password())
	to.SetGeneralTextData(m.GeneralTextData())
	to.SetInteger_1(m.Integer_1())
	to.SetInteger_2(m.Integer_2())
	to.SetHeartbeatIntervalInSeconds(m.HeartbeatIntervalInSeconds())
	to.SetTradeMode(m.TradeMode())
	to.SetTradeAccount(m.TradeAccount())
	to.SetHardwareIdentifier(m.HardwareIdentifier())
	to.SetClientName(m.ClientName())
	to.SetMarketDataTransmissionInterval(m.MarketDataTransmissionInterval())
}

// CopyTo
func (m LogonRequestFixedPointer) CopyTo(to LogonRequestBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetUsername(m.Username())
	to.SetPassword(m.Password())
	to.SetGeneralTextData(m.GeneralTextData())
	to.SetInteger_1(m.Integer_1())
	to.SetInteger_2(m.Integer_2())
	to.SetHeartbeatIntervalInSeconds(m.HeartbeatIntervalInSeconds())
	to.SetTradeMode(m.TradeMode())
	to.SetTradeAccount(m.TradeAccount())
	to.SetHardwareIdentifier(m.HardwareIdentifier())
	to.SetClientName(m.ClientName())
	to.SetMarketDataTransmissionInterval(m.MarketDataTransmissionInterval())
}

// CopyTo
func (m LogonRequestFixedPointerBuilder) CopyTo(to LogonRequestBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetUsername(m.Username())
	to.SetPassword(m.Password())
	to.SetGeneralTextData(m.GeneralTextData())
	to.SetInteger_1(m.Integer_1())
	to.SetInteger_2(m.Integer_2())
	to.SetHeartbeatIntervalInSeconds(m.HeartbeatIntervalInSeconds())
	to.SetTradeMode(m.TradeMode())
	to.SetTradeAccount(m.TradeAccount())
	to.SetHardwareIdentifier(m.HardwareIdentifier())
	to.SetClientName(m.ClientName())
	to.SetMarketDataTransmissionInterval(m.MarketDataTransmissionInterval())
}

// CopyToPointer
func (m LogonRequestFixedPointer) CopyToPointer(to LogonRequestPointerBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetUsername(m.Username())
	to.SetPassword(m.Password())
	to.SetGeneralTextData(m.GeneralTextData())
	to.SetInteger_1(m.Integer_1())
	to.SetInteger_2(m.Integer_2())
	to.SetHeartbeatIntervalInSeconds(m.HeartbeatIntervalInSeconds())
	to.SetTradeMode(m.TradeMode())
	to.SetTradeAccount(m.TradeAccount())
	to.SetHardwareIdentifier(m.HardwareIdentifier())
	to.SetClientName(m.ClientName())
	to.SetMarketDataTransmissionInterval(m.MarketDataTransmissionInterval())
}

// CopyToPointer
func (m LogonRequestFixedPointerBuilder) CopyToPointer(to LogonRequestPointerBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetUsername(m.Username())
	to.SetPassword(m.Password())
	to.SetGeneralTextData(m.GeneralTextData())
	to.SetInteger_1(m.Integer_1())
	to.SetInteger_2(m.Integer_2())
	to.SetHeartbeatIntervalInSeconds(m.HeartbeatIntervalInSeconds())
	to.SetTradeMode(m.TradeMode())
	to.SetTradeAccount(m.TradeAccount())
	to.SetHardwareIdentifier(m.HardwareIdentifier())
	to.SetClientName(m.ClientName())
	to.SetMarketDataTransmissionInterval(m.MarketDataTransmissionInterval())
}
