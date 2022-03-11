package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size                                          = LogonResponseSize  (52)
//     Type                                          = LOGON_RESPONSE  (2)
//     BaseSize                                      = LogonResponseSize  (52)
//     ProtocolVersion                               = CURRENT_VERSION  (8)
//     Result                                        = 0
//     ResultText                                    = ""
//     ReconnectAddress                              = ""
//     Integer_1                                     = 0
//     ServerName                                    = ""
//     MarketDepthUpdatesBestBidAndAsk               = 0
//     TradingIsSupported                            = 0
//     OCOOrdersSupported                            = 0
//     OrderCancelReplaceSupported                   = 1
//     SymbolExchangeDelimiter                       = ""
//     SecurityDefinitionsSupported                  = 0
//     HistoricalPriceDataSupported                  = 0
//     ResubscribeWhenMarketDataFeedAvailable        = 0
//     MarketDepthIsSupported                        = 1
//     OneHistoricalPriceDataRequestPerConnection    = 0
//     BracketOrdersSupported                        = 0
//     UseIntegerPriceOrderMessages                  = 0
//     UsesMultiplePositionsPerSymbolAndTradeAccount = 0
//     MarketDataSupported                           = 0
// }
var _LogonResponseDefault = []byte{52, 0, 2, 0, 52, 0, 0, 0, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}

const LogonResponseSize = 52

// LogonResponse This is a response message indicating either success or an error logging
// on to the Server.
type LogonResponse struct {
	message.VLS
}

// LogonResponseBuilder This is a response message indicating either success or an error logging
// on to the Server.
type LogonResponseBuilder struct {
	message.VLSBuilder
}

func NewLogonResponseFrom(b []byte) LogonResponse {
	return LogonResponse{VLS: message.NewVLSFrom(b)}
}

func WrapLogonResponse(b []byte) LogonResponse {
	return LogonResponse{VLS: message.WrapVLS(b)}
}

func NewLogonResponse() LogonResponseBuilder {
	a := LogonResponseBuilder{message.NewVLSBuilder(52)}
	a.Unsafe().SetBytes(0, _LogonResponseDefault)
	return a
}

// Clear
// {
//     Size                                          = LogonResponseSize  (52)
//     Type                                          = LOGON_RESPONSE  (2)
//     BaseSize                                      = LogonResponseSize  (52)
//     ProtocolVersion                               = CURRENT_VERSION  (8)
//     Result                                        = 0
//     ResultText                                    = ""
//     ReconnectAddress                              = ""
//     Integer_1                                     = 0
//     ServerName                                    = ""
//     MarketDepthUpdatesBestBidAndAsk               = 0
//     TradingIsSupported                            = 0
//     OCOOrdersSupported                            = 0
//     OrderCancelReplaceSupported                   = 1
//     SymbolExchangeDelimiter                       = ""
//     SecurityDefinitionsSupported                  = 0
//     HistoricalPriceDataSupported                  = 0
//     ResubscribeWhenMarketDataFeedAvailable        = 0
//     MarketDepthIsSupported                        = 1
//     OneHistoricalPriceDataRequestPerConnection    = 0
//     BracketOrdersSupported                        = 0
//     UseIntegerPriceOrderMessages                  = 0
//     UsesMultiplePositionsPerSymbolAndTradeAccount = 0
//     MarketDataSupported                           = 0
// }
func (m LogonResponseBuilder) Clear() {
	m.Unsafe().SetBytes(0, _LogonResponseDefault)
}

// ToBuilder
func (m LogonResponse) ToBuilder() LogonResponseBuilder {
	return LogonResponseBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m LogonResponseBuilder) Finish() LogonResponse {
	return LogonResponse{m.VLSBuilder.Finish()}
}

// ProtocolVersion This is automatically set by the constructor.
func (m LogonResponse) ProtocolVersion() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// ProtocolVersion This is automatically set by the constructor.
func (m LogonResponseBuilder) ProtocolVersion() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// Result This can be set to one of the following constants:
//
// LOGON_SUCCESS
// LOGON_ERROR
// LOGON_ERROR_NO_RECONNECT
// LOGON_RECONNECT_NEW_ADDRESS
// LOGON_ERROR_NO_RECONNECT means that there has been an error logging on
// and the Client should not try to reconnect.
//
// The Server can set this field to LOGON_RECONNECT_NEW_ADDRESS to instruct
// the Client to reconnect to the Server at a different address. The new
// address is specified through the ReconnectAddress field. This supports
// dynamic connections to a server farm.
func (m LogonResponse) Result() LogonStatusEnum {
	return LogonStatusEnum(message.Int32VLS(m.Unsafe(), 16, 12))
}

// Result This can be set to one of the following constants:
//
// LOGON_SUCCESS
// LOGON_ERROR
// LOGON_ERROR_NO_RECONNECT
// LOGON_RECONNECT_NEW_ADDRESS
// LOGON_ERROR_NO_RECONNECT means that there has been an error logging on
// and the Client should not try to reconnect.
//
// The Server can set this field to LOGON_RECONNECT_NEW_ADDRESS to instruct
// the Client to reconnect to the Server at a different address. The new
// address is specified through the ReconnectAddress field. This supports
// dynamic connections to a server farm.
func (m LogonResponseBuilder) Result() LogonStatusEnum {
	return LogonStatusEnum(message.Int32VLS(m.Unsafe(), 16, 12))
}

// ResultText Optional freeform text to provide information related to a successful
// or unsuccessful logon. The Client will display this text to the user.
func (m LogonResponse) ResultText() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
}

// ResultText Optional freeform text to provide information related to a successful
// or unsuccessful logon. The Client will display this text to the user.
func (m LogonResponseBuilder) ResultText() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
}

// ReconnectAddress Server address/IP number and optional port number to reconnect to. Format:
// [Server Address:Port Number]. Only used if Result is set to LOGON_RECONNECT_NEW_ADDRESS.
// [Server Address:Port Number]. Only used if Result is set to LOGON_RECONNECT_NEW_ADDRESS.
func (m LogonResponse) ReconnectAddress() string {
	return message.StringVLS(m.Unsafe(), 24, 20)
}

// ReconnectAddress Server address/IP number and optional port number to reconnect to. Format:
// [Server Address:Port Number]. Only used if Result is set to LOGON_RECONNECT_NEW_ADDRESS.
// [Server Address:Port Number]. Only used if Result is set to LOGON_RECONNECT_NEW_ADDRESS.
func (m LogonResponseBuilder) ReconnectAddress() string {
	return message.StringVLS(m.Unsafe(), 24, 20)
}

// Integer_1 Optional. General-purpose integer for the Server to communicate to the
// Client an integer value on logon.
func (m LogonResponse) Integer_1() int32 {
	return message.Int32VLS(m.Unsafe(), 28, 24)
}

// Integer_1 Optional. General-purpose integer for the Server to communicate to the
// Client an integer value on logon.
func (m LogonResponseBuilder) Integer_1() int32 {
	return message.Int32VLS(m.Unsafe(), 28, 24)
}

// ServerName Optional free-form text for the Server to fill out.
//
// It is recommended that the Server fill this in with descriptive text identifying
// itself to the Client.
//
// The length of this text string is 60 characters when fixed length strings
// are used.
func (m LogonResponse) ServerName() string {
	return message.StringVLS(m.Unsafe(), 32, 28)
}

// ServerName Optional free-form text for the Server to fill out.
//
// It is recommended that the Server fill this in with descriptive text identifying
// itself to the Client.
//
// The length of this text string is 60 characters when fixed length strings
// are used.
func (m LogonResponseBuilder) ServerName() string {
	return message.StringVLS(m.Unsafe(), 32, 28)
}

// MarketDepthUpdatesBestBidAndAsk Set this to 1 to indicate that the Server will only be sending market
// depth updates and not best bid and ask updates. The Client will use depth
// at level 1 to update the best bid and ask prices.
//
// Some Clients will maintain separate best bid and ask prices from market
// depth data.
func (m LogonResponse) MarketDepthUpdatesBestBidAndAsk() uint8 {
	return message.Uint8VLS(m.Unsafe(), 33, 32)
}

// MarketDepthUpdatesBestBidAndAsk Set this to 1 to indicate that the Server will only be sending market
// depth updates and not best bid and ask updates. The Client will use depth
// at level 1 to update the best bid and ask prices.
//
// Some Clients will maintain separate best bid and ask prices from market
// depth data.
func (m LogonResponseBuilder) MarketDepthUpdatesBestBidAndAsk() uint8 {
	return message.Uint8VLS(m.Unsafe(), 33, 32)
}

// TradingIsSupported Set this to 1 to indicate the Server supports trading. Otherwise, the
// Client will not send through any trading messages.
func (m LogonResponse) TradingIsSupported() uint8 {
	return message.Uint8VLS(m.Unsafe(), 34, 33)
}

// TradingIsSupported Set this to 1 to indicate the Server supports trading. Otherwise, the
// Client will not send through any trading messages.
func (m LogonResponseBuilder) TradingIsSupported() uint8 {
	return message.Uint8VLS(m.Unsafe(), 34, 33)
}

// OCOOrdersSupported Set this to 1 to indicate the Server supports OCO orders.
func (m LogonResponse) OCOOrdersSupported() uint8 {
	return message.Uint8VLS(m.Unsafe(), 35, 34)
}

// OCOOrdersSupported Set this to 1 to indicate the Server supports OCO orders.
func (m LogonResponseBuilder) OCOOrdersSupported() uint8 {
	return message.Uint8VLS(m.Unsafe(), 35, 34)
}

// OrderCancelReplaceSupported Set this to 0 if Server does not support the CancelReplaceOrder message.
// Set this to 0 if Server does not support the CancelReplaceOrder message.
func (m LogonResponse) OrderCancelReplaceSupported() uint8 {
	return message.Uint8VLS(m.Unsafe(), 36, 35)
}

// OrderCancelReplaceSupported Set this to 0 if Server does not support the CancelReplaceOrder message.
// Set this to 0 if Server does not support the CancelReplaceOrder message.
func (m LogonResponseBuilder) OrderCancelReplaceSupported() uint8 {
	return message.Uint8VLS(m.Unsafe(), 36, 35)
}

// SymbolExchangeDelimiter Some Clients will usually consider the Symbol and Exchange fields as a
// single text string. If the Server will be using the Exchange field in
// DTC messages that have a Symbol and Exchange fields, it must specify the
// SymbolExchangeDelimiter field to provide a standard delimiter for the
// Client to use to combine the Symbol and the Exchange into a single text
// string.
//
// It is recommended to use a "-" or ".". Examples of how the Client will
// then combine the Symbol and exchange.
//
// Symbol-Exchange
// Symbol.Exchange
// If this field is unset, then this is an indication to the Client that
// the Exchange field in DTC Protocol messages are not used.
//
// Even if the symbols supported by a Server have an Exchange text string,
// does not mean the Server has to use the Exchange field in DTC messages.
// The Server can combine the Symbol and the Exchange in Security Definition
// responses into the Symbol field only.
//
// When a Client sees that the SymbolExchangeDelimiter field is set, then
// it can use this delimiter to combine the Symbol and Exchange into a single
// text string. When the Client is setting the Symbol and Exchange in DTC
// messages, it needs to separate out the Symbol and Exchange from the larger
// text string and set those fields separately.
func (m LogonResponse) SymbolExchangeDelimiter() string {
	return message.StringVLS(m.Unsafe(), 40, 36)
}

// SymbolExchangeDelimiter Some Clients will usually consider the Symbol and Exchange fields as a
// single text string. If the Server will be using the Exchange field in
// DTC messages that have a Symbol and Exchange fields, it must specify the
// SymbolExchangeDelimiter field to provide a standard delimiter for the
// Client to use to combine the Symbol and the Exchange into a single text
// string.
//
// It is recommended to use a "-" or ".". Examples of how the Client will
// then combine the Symbol and exchange.
//
// Symbol-Exchange
// Symbol.Exchange
// If this field is unset, then this is an indication to the Client that
// the Exchange field in DTC Protocol messages are not used.
//
// Even if the symbols supported by a Server have an Exchange text string,
// does not mean the Server has to use the Exchange field in DTC messages.
// The Server can combine the Symbol and the Exchange in Security Definition
// responses into the Symbol field only.
//
// When a Client sees that the SymbolExchangeDelimiter field is set, then
// it can use this delimiter to combine the Symbol and Exchange into a single
// text string. When the Client is setting the Symbol and Exchange in DTC
// messages, it needs to separate out the Symbol and Exchange from the larger
// text string and set those fields separately.
func (m LogonResponseBuilder) SymbolExchangeDelimiter() string {
	return message.StringVLS(m.Unsafe(), 40, 36)
}

// SecurityDefinitionsSupported Set to 1 if the Server supports Security Definition messages.
func (m LogonResponse) SecurityDefinitionsSupported() uint8 {
	return message.Uint8VLS(m.Unsafe(), 41, 40)
}

// SecurityDefinitionsSupported Set to 1 if the Server supports Security Definition messages.
func (m LogonResponseBuilder) SecurityDefinitionsSupported() uint8 {
	return message.Uint8VLS(m.Unsafe(), 41, 40)
}

// HistoricalPriceDataSupported Set this to 1 if the Server supports the HistoricalPriceDataRequest message.
// Set this to 1 if the Server supports the HistoricalPriceDataRequest message.
func (m LogonResponse) HistoricalPriceDataSupported() uint8 {
	return message.Uint8VLS(m.Unsafe(), 42, 41)
}

// HistoricalPriceDataSupported Set this to 1 if the Server supports the HistoricalPriceDataRequest message.
// Set this to 1 if the Server supports the HistoricalPriceDataRequest message.
func (m LogonResponseBuilder) HistoricalPriceDataSupported() uint8 {
	return message.Uint8VLS(m.Unsafe(), 42, 41)
}

// ResubscribeWhenMarketDataFeedAvailable Set this to 1, so that when the Client receives a MarketDataFeedStatus
// indicating the market data feed is restored, it will resubscribe to market
// data and market depth for all of the symbols it was previously tracking.
// data and market depth for all of the symbols it was previously tracking.
func (m LogonResponse) ResubscribeWhenMarketDataFeedAvailable() uint8 {
	return message.Uint8VLS(m.Unsafe(), 43, 42)
}

// ResubscribeWhenMarketDataFeedAvailable Set this to 1, so that when the Client receives a MarketDataFeedStatus
// indicating the market data feed is restored, it will resubscribe to market
// data and market depth for all of the symbols it was previously tracking.
// data and market depth for all of the symbols it was previously tracking.
func (m LogonResponseBuilder) ResubscribeWhenMarketDataFeedAvailable() uint8 {
	return message.Uint8VLS(m.Unsafe(), 43, 42)
}

// MarketDepthIsSupported Set this to 1, if the Server supports the MarketDepthRequest message.
//
// The default is 0.
func (m LogonResponse) MarketDepthIsSupported() uint8 {
	return message.Uint8VLS(m.Unsafe(), 44, 43)
}

// MarketDepthIsSupported Set this to 1, if the Server supports the MarketDepthRequest message.
//
// The default is 0.
func (m LogonResponseBuilder) MarketDepthIsSupported() uint8 {
	return message.Uint8VLS(m.Unsafe(), 44, 43)
}

// OneHistoricalPriceDataRequestPerConnection The server can optionally set the OneHistoricalPriceDataRequestPerConnection
// field to 1 in the LogonResponse message to indicate that it only will
// accept one historical price data request per network connection.
//
// After the first request is served or rejected, the network connection
// will be gracefully closed at the appropriate time by the Server. This
// method simplifies the serving of historical price data on the Server side
// and the implementation on the Client side when data compression is used.
// and the implementation on the Client side when data compression is used.
func (m LogonResponse) OneHistoricalPriceDataRequestPerConnection() uint8 {
	return message.Uint8VLS(m.Unsafe(), 45, 44)
}

// OneHistoricalPriceDataRequestPerConnection The server can optionally set the OneHistoricalPriceDataRequestPerConnection
// field to 1 in the LogonResponse message to indicate that it only will
// accept one historical price data request per network connection.
//
// After the first request is served or rejected, the network connection
// will be gracefully closed at the appropriate time by the Server. This
// method simplifies the serving of historical price data on the Server side
// and the implementation on the Client side when data compression is used.
// and the implementation on the Client side when data compression is used.
func (m LogonResponseBuilder) OneHistoricalPriceDataRequestPerConnection() uint8 {
	return message.Uint8VLS(m.Unsafe(), 45, 44)
}

// BracketOrdersSupported Set this to 1 to indicate the Server supports bracket orders.
func (m LogonResponse) BracketOrdersSupported() uint8 {
	return message.Uint8VLS(m.Unsafe(), 46, 45)
}

// BracketOrdersSupported Set this to 1 to indicate the Server supports bracket orders.
func (m LogonResponseBuilder) BracketOrdersSupported() uint8 {
	return message.Uint8VLS(m.Unsafe(), 46, 45)
}

// UseIntegerPriceOrderMessages With the integer trading messages discontinued as of August 2020, this
// field is no longer relevant.
func (m LogonResponse) UseIntegerPriceOrderMessages() uint8 {
	return message.Uint8VLS(m.Unsafe(), 47, 46)
}

// UseIntegerPriceOrderMessages With the integer trading messages discontinued as of August 2020, this
// field is no longer relevant.
func (m LogonResponseBuilder) UseIntegerPriceOrderMessages() uint8 {
	return message.Uint8VLS(m.Unsafe(), 47, 46)
}

// UsesMultiplePositionsPerSymbolAndTradeAccount If the Server can report more than one Trade Position for a specific Symbol
// and Trade Account, then it needs to set UsesMultiplePositionsPerSymbolAndTradeAccount
// to 1.
//
// When the server has set to 1, it must always set PositionIdentifier in
// the PositionUpdate message to the identifier of the Trade Position.
//
// When the Client checks that this is set to 1, then it knows that it can
// expect there potentially can be more than one Trade Position for a specific
// Symbol and Trade Account being reported by the PositionUpdate messages.
// The Client can then handle this appropriately.
func (m LogonResponse) UsesMultiplePositionsPerSymbolAndTradeAccount() uint8 {
	return message.Uint8VLS(m.Unsafe(), 48, 47)
}

// UsesMultiplePositionsPerSymbolAndTradeAccount If the Server can report more than one Trade Position for a specific Symbol
// and Trade Account, then it needs to set UsesMultiplePositionsPerSymbolAndTradeAccount
// to 1.
//
// When the server has set to 1, it must always set PositionIdentifier in
// the PositionUpdate message to the identifier of the Trade Position.
//
// When the Client checks that this is set to 1, then it knows that it can
// expect there potentially can be more than one Trade Position for a specific
// Symbol and Trade Account being reported by the PositionUpdate messages.
// The Client can then handle this appropriately.
func (m LogonResponseBuilder) UsesMultiplePositionsPerSymbolAndTradeAccount() uint8 {
	return message.Uint8VLS(m.Unsafe(), 48, 47)
}

// MarketDataSupported Set this to 1, if the Server supports the MarketDataRequest message.
//
// The default is 1.
func (m LogonResponse) MarketDataSupported() uint8 {
	return message.Uint8VLS(m.Unsafe(), 49, 48)
}

// MarketDataSupported Set this to 1, if the Server supports the MarketDataRequest message.
//
// The default is 1.
func (m LogonResponseBuilder) MarketDataSupported() uint8 {
	return message.Uint8VLS(m.Unsafe(), 49, 48)
}

// SetProtocolVersion This is automatically set by the constructor.
func (m LogonResponseBuilder) SetProtocolVersion(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetResult This can be set to one of the following constants:
//
// LOGON_SUCCESS
// LOGON_ERROR
// LOGON_ERROR_NO_RECONNECT
// LOGON_RECONNECT_NEW_ADDRESS
// LOGON_ERROR_NO_RECONNECT means that there has been an error logging on
// and the Client should not try to reconnect.
//
// The Server can set this field to LOGON_RECONNECT_NEW_ADDRESS to instruct
// the Client to reconnect to the Server at a different address. The new
// address is specified through the ReconnectAddress field. This supports
// dynamic connections to a server farm.
func (m LogonResponseBuilder) SetResult(value LogonStatusEnum) {
	message.SetInt32VLS(m.Unsafe(), 16, 12, int32(value))
}

// SetResultText Optional freeform text to provide information related to a successful
// or unsuccessful logon. The Client will display this text to the user.
func (m *LogonResponseBuilder) SetResultText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 20, 16, value)
}

// SetReconnectAddress Server address/IP number and optional port number to reconnect to. Format:
// [Server Address:Port Number]. Only used if Result is set to LOGON_RECONNECT_NEW_ADDRESS.
// [Server Address:Port Number]. Only used if Result is set to LOGON_RECONNECT_NEW_ADDRESS.
func (m *LogonResponseBuilder) SetReconnectAddress(value string) {
	message.SetStringVLS(&m.VLSBuilder, 24, 20, value)
}

// SetInteger_1 Optional. General-purpose integer for the Server to communicate to the
// Client an integer value on logon.
func (m LogonResponseBuilder) SetInteger_1(value int32) {
	message.SetInt32VLS(m.Unsafe(), 28, 24, value)
}

// SetServerName Optional free-form text for the Server to fill out.
//
// It is recommended that the Server fill this in with descriptive text identifying
// itself to the Client.
//
// The length of this text string is 60 characters when fixed length strings
// are used.
func (m *LogonResponseBuilder) SetServerName(value string) {
	message.SetStringVLS(&m.VLSBuilder, 32, 28, value)
}

// SetMarketDepthUpdatesBestBidAndAsk Set this to 1 to indicate that the Server will only be sending market
// depth updates and not best bid and ask updates. The Client will use depth
// at level 1 to update the best bid and ask prices.
//
// Some Clients will maintain separate best bid and ask prices from market
// depth data.
func (m LogonResponseBuilder) SetMarketDepthUpdatesBestBidAndAsk(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 33, 32, value)
}

// SetTradingIsSupported Set this to 1 to indicate the Server supports trading. Otherwise, the
// Client will not send through any trading messages.
func (m LogonResponseBuilder) SetTradingIsSupported(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 34, 33, value)
}

// SetOCOOrdersSupported Set this to 1 to indicate the Server supports OCO orders.
func (m LogonResponseBuilder) SetOCOOrdersSupported(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 35, 34, value)
}

// SetOrderCancelReplaceSupported Set this to 0 if Server does not support the CancelReplaceOrder message.
// Set this to 0 if Server does not support the CancelReplaceOrder message.
func (m LogonResponseBuilder) SetOrderCancelReplaceSupported(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 36, 35, value)
}

// SetSymbolExchangeDelimiter Some Clients will usually consider the Symbol and Exchange fields as a
// single text string. If the Server will be using the Exchange field in
// DTC messages that have a Symbol and Exchange fields, it must specify the
// SymbolExchangeDelimiter field to provide a standard delimiter for the
// Client to use to combine the Symbol and the Exchange into a single text
// string.
//
// It is recommended to use a "-" or ".". Examples of how the Client will
// then combine the Symbol and exchange.
//
// Symbol-Exchange
// Symbol.Exchange
// If this field is unset, then this is an indication to the Client that
// the Exchange field in DTC Protocol messages are not used.
//
// Even if the symbols supported by a Server have an Exchange text string,
// does not mean the Server has to use the Exchange field in DTC messages.
// The Server can combine the Symbol and the Exchange in Security Definition
// responses into the Symbol field only.
//
// When a Client sees that the SymbolExchangeDelimiter field is set, then
// it can use this delimiter to combine the Symbol and Exchange into a single
// text string. When the Client is setting the Symbol and Exchange in DTC
// messages, it needs to separate out the Symbol and Exchange from the larger
// text string and set those fields separately.
func (m *LogonResponseBuilder) SetSymbolExchangeDelimiter(value string) {
	message.SetStringVLS(&m.VLSBuilder, 40, 36, value)
}

// SetSecurityDefinitionsSupported Set to 1 if the Server supports Security Definition messages.
func (m LogonResponseBuilder) SetSecurityDefinitionsSupported(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 41, 40, value)
}

// SetHistoricalPriceDataSupported Set this to 1 if the Server supports the HistoricalPriceDataRequest message.
// Set this to 1 if the Server supports the HistoricalPriceDataRequest message.
func (m LogonResponseBuilder) SetHistoricalPriceDataSupported(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 42, 41, value)
}

// SetResubscribeWhenMarketDataFeedAvailable Set this to 1, so that when the Client receives a MarketDataFeedStatus
// indicating the market data feed is restored, it will resubscribe to market
// data and market depth for all of the symbols it was previously tracking.
// data and market depth for all of the symbols it was previously tracking.
func (m LogonResponseBuilder) SetResubscribeWhenMarketDataFeedAvailable(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 43, 42, value)
}

// SetMarketDepthIsSupported Set this to 1, if the Server supports the MarketDepthRequest message.
//
// The default is 0.
func (m LogonResponseBuilder) SetMarketDepthIsSupported(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 44, 43, value)
}

// SetOneHistoricalPriceDataRequestPerConnection The server can optionally set the OneHistoricalPriceDataRequestPerConnection
// field to 1 in the LogonResponse message to indicate that it only will
// accept one historical price data request per network connection.
//
// After the first request is served or rejected, the network connection
// will be gracefully closed at the appropriate time by the Server. This
// method simplifies the serving of historical price data on the Server side
// and the implementation on the Client side when data compression is used.
// and the implementation on the Client side when data compression is used.
func (m LogonResponseBuilder) SetOneHistoricalPriceDataRequestPerConnection(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 45, 44, value)
}

// SetBracketOrdersSupported Set this to 1 to indicate the Server supports bracket orders.
func (m LogonResponseBuilder) SetBracketOrdersSupported(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 46, 45, value)
}

// SetUseIntegerPriceOrderMessages With the integer trading messages discontinued as of August 2020, this
// field is no longer relevant.
func (m LogonResponseBuilder) SetUseIntegerPriceOrderMessages(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 47, 46, value)
}

// SetUsesMultiplePositionsPerSymbolAndTradeAccount If the Server can report more than one Trade Position for a specific Symbol
// and Trade Account, then it needs to set UsesMultiplePositionsPerSymbolAndTradeAccount
// to 1.
//
// When the server has set to 1, it must always set PositionIdentifier in
// the PositionUpdate message to the identifier of the Trade Position.
//
// When the Client checks that this is set to 1, then it knows that it can
// expect there potentially can be more than one Trade Position for a specific
// Symbol and Trade Account being reported by the PositionUpdate messages.
// The Client can then handle this appropriately.
func (m LogonResponseBuilder) SetUsesMultiplePositionsPerSymbolAndTradeAccount(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 48, 47, value)
}

// SetMarketDataSupported Set this to 1, if the Server supports the MarketDataRequest message.
//
// The default is 1.
func (m LogonResponseBuilder) SetMarketDataSupported(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 49, 48, value)
}

// Copy
func (m LogonResponse) Copy(to LogonResponseBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetResult(m.Result())
	to.SetResultText(m.ResultText())
	to.SetReconnectAddress(m.ReconnectAddress())
	to.SetInteger_1(m.Integer_1())
	to.SetServerName(m.ServerName())
	to.SetMarketDepthUpdatesBestBidAndAsk(m.MarketDepthUpdatesBestBidAndAsk())
	to.SetTradingIsSupported(m.TradingIsSupported())
	to.SetOCOOrdersSupported(m.OCOOrdersSupported())
	to.SetOrderCancelReplaceSupported(m.OrderCancelReplaceSupported())
	to.SetSymbolExchangeDelimiter(m.SymbolExchangeDelimiter())
	to.SetSecurityDefinitionsSupported(m.SecurityDefinitionsSupported())
	to.SetHistoricalPriceDataSupported(m.HistoricalPriceDataSupported())
	to.SetResubscribeWhenMarketDataFeedAvailable(m.ResubscribeWhenMarketDataFeedAvailable())
	to.SetMarketDepthIsSupported(m.MarketDepthIsSupported())
	to.SetOneHistoricalPriceDataRequestPerConnection(m.OneHistoricalPriceDataRequestPerConnection())
	to.SetBracketOrdersSupported(m.BracketOrdersSupported())
	to.SetUseIntegerPriceOrderMessages(m.UseIntegerPriceOrderMessages())
	to.SetUsesMultiplePositionsPerSymbolAndTradeAccount(m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	to.SetMarketDataSupported(m.MarketDataSupported())
}

// Copy
func (m LogonResponseBuilder) Copy(to LogonResponseBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetResult(m.Result())
	to.SetResultText(m.ResultText())
	to.SetReconnectAddress(m.ReconnectAddress())
	to.SetInteger_1(m.Integer_1())
	to.SetServerName(m.ServerName())
	to.SetMarketDepthUpdatesBestBidAndAsk(m.MarketDepthUpdatesBestBidAndAsk())
	to.SetTradingIsSupported(m.TradingIsSupported())
	to.SetOCOOrdersSupported(m.OCOOrdersSupported())
	to.SetOrderCancelReplaceSupported(m.OrderCancelReplaceSupported())
	to.SetSymbolExchangeDelimiter(m.SymbolExchangeDelimiter())
	to.SetSecurityDefinitionsSupported(m.SecurityDefinitionsSupported())
	to.SetHistoricalPriceDataSupported(m.HistoricalPriceDataSupported())
	to.SetResubscribeWhenMarketDataFeedAvailable(m.ResubscribeWhenMarketDataFeedAvailable())
	to.SetMarketDepthIsSupported(m.MarketDepthIsSupported())
	to.SetOneHistoricalPriceDataRequestPerConnection(m.OneHistoricalPriceDataRequestPerConnection())
	to.SetBracketOrdersSupported(m.BracketOrdersSupported())
	to.SetUseIntegerPriceOrderMessages(m.UseIntegerPriceOrderMessages())
	to.SetUsesMultiplePositionsPerSymbolAndTradeAccount(m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	to.SetMarketDataSupported(m.MarketDataSupported())
}

// CopyPointer
func (m LogonResponse) CopyPointer(to LogonResponsePointerBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetResult(m.Result())
	to.SetResultText(m.ResultText())
	to.SetReconnectAddress(m.ReconnectAddress())
	to.SetInteger_1(m.Integer_1())
	to.SetServerName(m.ServerName())
	to.SetMarketDepthUpdatesBestBidAndAsk(m.MarketDepthUpdatesBestBidAndAsk())
	to.SetTradingIsSupported(m.TradingIsSupported())
	to.SetOCOOrdersSupported(m.OCOOrdersSupported())
	to.SetOrderCancelReplaceSupported(m.OrderCancelReplaceSupported())
	to.SetSymbolExchangeDelimiter(m.SymbolExchangeDelimiter())
	to.SetSecurityDefinitionsSupported(m.SecurityDefinitionsSupported())
	to.SetHistoricalPriceDataSupported(m.HistoricalPriceDataSupported())
	to.SetResubscribeWhenMarketDataFeedAvailable(m.ResubscribeWhenMarketDataFeedAvailable())
	to.SetMarketDepthIsSupported(m.MarketDepthIsSupported())
	to.SetOneHistoricalPriceDataRequestPerConnection(m.OneHistoricalPriceDataRequestPerConnection())
	to.SetBracketOrdersSupported(m.BracketOrdersSupported())
	to.SetUseIntegerPriceOrderMessages(m.UseIntegerPriceOrderMessages())
	to.SetUsesMultiplePositionsPerSymbolAndTradeAccount(m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	to.SetMarketDataSupported(m.MarketDataSupported())
}

// CopyPointer
func (m LogonResponseBuilder) CopyPointer(to LogonResponsePointerBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetResult(m.Result())
	to.SetResultText(m.ResultText())
	to.SetReconnectAddress(m.ReconnectAddress())
	to.SetInteger_1(m.Integer_1())
	to.SetServerName(m.ServerName())
	to.SetMarketDepthUpdatesBestBidAndAsk(m.MarketDepthUpdatesBestBidAndAsk())
	to.SetTradingIsSupported(m.TradingIsSupported())
	to.SetOCOOrdersSupported(m.OCOOrdersSupported())
	to.SetOrderCancelReplaceSupported(m.OrderCancelReplaceSupported())
	to.SetSymbolExchangeDelimiter(m.SymbolExchangeDelimiter())
	to.SetSecurityDefinitionsSupported(m.SecurityDefinitionsSupported())
	to.SetHistoricalPriceDataSupported(m.HistoricalPriceDataSupported())
	to.SetResubscribeWhenMarketDataFeedAvailable(m.ResubscribeWhenMarketDataFeedAvailable())
	to.SetMarketDepthIsSupported(m.MarketDepthIsSupported())
	to.SetOneHistoricalPriceDataRequestPerConnection(m.OneHistoricalPriceDataRequestPerConnection())
	to.SetBracketOrdersSupported(m.BracketOrdersSupported())
	to.SetUseIntegerPriceOrderMessages(m.UseIntegerPriceOrderMessages())
	to.SetUsesMultiplePositionsPerSymbolAndTradeAccount(m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	to.SetMarketDataSupported(m.MarketDataSupported())
}

// CopyToPointer
func (m LogonResponse) CopyToPointer(to LogonResponseFixedPointerBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetResult(m.Result())
	to.SetResultText(m.ResultText())
	to.SetReconnectAddress(m.ReconnectAddress())
	to.SetInteger_1(m.Integer_1())
	to.SetServerName(m.ServerName())
	to.SetMarketDepthUpdatesBestBidAndAsk(m.MarketDepthUpdatesBestBidAndAsk())
	to.SetTradingIsSupported(m.TradingIsSupported())
	to.SetOCOOrdersSupported(m.OCOOrdersSupported())
	to.SetOrderCancelReplaceSupported(m.OrderCancelReplaceSupported())
	to.SetSymbolExchangeDelimiter(m.SymbolExchangeDelimiter())
	to.SetSecurityDefinitionsSupported(m.SecurityDefinitionsSupported())
	to.SetHistoricalPriceDataSupported(m.HistoricalPriceDataSupported())
	to.SetResubscribeWhenMarketDataFeedAvailable(m.ResubscribeWhenMarketDataFeedAvailable())
	to.SetMarketDepthIsSupported(m.MarketDepthIsSupported())
	to.SetOneHistoricalPriceDataRequestPerConnection(m.OneHistoricalPriceDataRequestPerConnection())
	to.SetBracketOrdersSupported(m.BracketOrdersSupported())
	to.SetUseIntegerPriceOrderMessages(m.UseIntegerPriceOrderMessages())
	to.SetUsesMultiplePositionsPerSymbolAndTradeAccount(m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	to.SetMarketDataSupported(m.MarketDataSupported())
}

// CopyToPointer
func (m LogonResponseBuilder) CopyToPointer(to LogonResponseFixedPointerBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetResult(m.Result())
	to.SetResultText(m.ResultText())
	to.SetReconnectAddress(m.ReconnectAddress())
	to.SetInteger_1(m.Integer_1())
	to.SetServerName(m.ServerName())
	to.SetMarketDepthUpdatesBestBidAndAsk(m.MarketDepthUpdatesBestBidAndAsk())
	to.SetTradingIsSupported(m.TradingIsSupported())
	to.SetOCOOrdersSupported(m.OCOOrdersSupported())
	to.SetOrderCancelReplaceSupported(m.OrderCancelReplaceSupported())
	to.SetSymbolExchangeDelimiter(m.SymbolExchangeDelimiter())
	to.SetSecurityDefinitionsSupported(m.SecurityDefinitionsSupported())
	to.SetHistoricalPriceDataSupported(m.HistoricalPriceDataSupported())
	to.SetResubscribeWhenMarketDataFeedAvailable(m.ResubscribeWhenMarketDataFeedAvailable())
	to.SetMarketDepthIsSupported(m.MarketDepthIsSupported())
	to.SetOneHistoricalPriceDataRequestPerConnection(m.OneHistoricalPriceDataRequestPerConnection())
	to.SetBracketOrdersSupported(m.BracketOrdersSupported())
	to.SetUseIntegerPriceOrderMessages(m.UseIntegerPriceOrderMessages())
	to.SetUsesMultiplePositionsPerSymbolAndTradeAccount(m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	to.SetMarketDataSupported(m.MarketDataSupported())
}

// CopyTo
func (m LogonResponse) CopyTo(to LogonResponseFixedBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetResult(m.Result())
	to.SetResultText(m.ResultText())
	to.SetReconnectAddress(m.ReconnectAddress())
	to.SetInteger_1(m.Integer_1())
	to.SetServerName(m.ServerName())
	to.SetMarketDepthUpdatesBestBidAndAsk(m.MarketDepthUpdatesBestBidAndAsk())
	to.SetTradingIsSupported(m.TradingIsSupported())
	to.SetOCOOrdersSupported(m.OCOOrdersSupported())
	to.SetOrderCancelReplaceSupported(m.OrderCancelReplaceSupported())
	to.SetSymbolExchangeDelimiter(m.SymbolExchangeDelimiter())
	to.SetSecurityDefinitionsSupported(m.SecurityDefinitionsSupported())
	to.SetHistoricalPriceDataSupported(m.HistoricalPriceDataSupported())
	to.SetResubscribeWhenMarketDataFeedAvailable(m.ResubscribeWhenMarketDataFeedAvailable())
	to.SetMarketDepthIsSupported(m.MarketDepthIsSupported())
	to.SetOneHistoricalPriceDataRequestPerConnection(m.OneHistoricalPriceDataRequestPerConnection())
	to.SetBracketOrdersSupported(m.BracketOrdersSupported())
	to.SetUseIntegerPriceOrderMessages(m.UseIntegerPriceOrderMessages())
	to.SetUsesMultiplePositionsPerSymbolAndTradeAccount(m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	to.SetMarketDataSupported(m.MarketDataSupported())
}

// CopyTo
func (m LogonResponseBuilder) CopyTo(to LogonResponseFixedBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetResult(m.Result())
	to.SetResultText(m.ResultText())
	to.SetReconnectAddress(m.ReconnectAddress())
	to.SetInteger_1(m.Integer_1())
	to.SetServerName(m.ServerName())
	to.SetMarketDepthUpdatesBestBidAndAsk(m.MarketDepthUpdatesBestBidAndAsk())
	to.SetTradingIsSupported(m.TradingIsSupported())
	to.SetOCOOrdersSupported(m.OCOOrdersSupported())
	to.SetOrderCancelReplaceSupported(m.OrderCancelReplaceSupported())
	to.SetSymbolExchangeDelimiter(m.SymbolExchangeDelimiter())
	to.SetSecurityDefinitionsSupported(m.SecurityDefinitionsSupported())
	to.SetHistoricalPriceDataSupported(m.HistoricalPriceDataSupported())
	to.SetResubscribeWhenMarketDataFeedAvailable(m.ResubscribeWhenMarketDataFeedAvailable())
	to.SetMarketDepthIsSupported(m.MarketDepthIsSupported())
	to.SetOneHistoricalPriceDataRequestPerConnection(m.OneHistoricalPriceDataRequestPerConnection())
	to.SetBracketOrdersSupported(m.BracketOrdersSupported())
	to.SetUseIntegerPriceOrderMessages(m.UseIntegerPriceOrderMessages())
	to.SetUsesMultiplePositionsPerSymbolAndTradeAccount(m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	to.SetMarketDataSupported(m.MarketDataSupported())
}
