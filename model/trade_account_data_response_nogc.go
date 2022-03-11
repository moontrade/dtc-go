package model

import (
	"github.com/moontrade/dtc-go/message"
)

type TradeAccountDataResponsePointer struct {
	message.VLSPointer
}

type TradeAccountDataResponsePointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocTradeAccountDataResponse() TradeAccountDataResponsePointerBuilder {
	a := TradeAccountDataResponsePointerBuilder{message.AllocVLSBuilder(311)}
	a.Ptr.SetBytes(0, _TradeAccountDataResponseDefault)
	return a
}

func AllocTradeAccountDataResponseFrom(b []byte) TradeAccountDataResponsePointer {
	return TradeAccountDataResponsePointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size                                                               = TradeAccountDataResponseSize  (311)
//     Type                                                               = TRADE_ACCOUNT_DATA_RESPONSE  (10116)
//     BaseSize                                                           = TradeAccountDataResponseSize  (311)
//     RequestID                                                          = 0
//     TradeAccountNotExist                                               = 0
//     TradeAccount                                                       = ""
//     IsSimulated                                                        = 0
//     CurrencyCode                                                       = ""
//     CurrentCashBalance                                                 = 0
//     AvailableFundsForNewPositions                                      = 0
//     MarginRequirement                                                  = 0
//     AccountValue                                                       = 0
//     OpenPositionsProfitLoss                                            = 0
//     DailyProfitLoss                                                    = 0
//     TransactionIdentifierForCashBalanceAdjustment                      = 0
//     LastTransactionDateTime                                            = 0
//     TrailingAccountValueAtWhichToNotAllowNewPositions                  = 0
//     CalculatedDailyNetLossLimitInAccountCurrency                       = 0
//     DailyNetLossLimitHasBeenReached                                    = 0
//     LastResetDailyNetLossManagementVariablesDateTimeUTC                = 0
//     IsUnderRequiredMargin                                              = 0
//     DailyNetLossLimitInAccountCurrency                                 = 0
//     PercentOfCashBalanceForDailyNetLossLimit                           = 0
//     UseTrailingAccountValueToNotAllowIncreaseInPositions               = 0
//     DoNotAllowIncreaseInPositionsAtDailyLossLimit                      = 0
//     FlattenPositionsAtDailyLossLimit                                   = 0
//     ClosePositionsAtEndOfDay                                           = 0
//     FlattenPositionsWhenUnderMarginIntraday                            = 1
//     FlattenPositionsWhenUnderMarginAtEndOfDay                          = 0
//     SenderSubID                                                        = ""
//     SenderLocationId                                                   = ""
//     SelfMatchPreventionID                                              = ""
//     CTICode                                                            = 0
//     TradeAccountIsReadOnly                                             = 0
//     MaximumGlobalPositionQuantity                                      = 0
//     TradeFeePerContract                                                = 0
//     TradeFeePerShare                                                   = 0
//     RequireSufficientMarginForNewPositions                             = 1
//     UsePercentOfMargin                                                 = 100
//     UsePercentOfMarginForDayTrading                                    = 100
//     MaximumAllowedAccountBalanceForPositionsAsPercentage               = 100
//     FirmID                                                             = ""
//     TradingIsDisabled                                                  = 0
//     DescriptiveName                                                    = ""
//     IsMasterFirmControlAccount                                         = 0
//     MinimumRequiredAccountValue                                        = 0
//     BeginTimeForDayMargin                                              = 0
//     EndTimeForDayMargin                                                = 0
//     DayMarginTimeZone                                                  = ""
//     IsSnapshot                                                         = 0
//     IsFirstMessageInBatch                                              = 0
//     IsLastMessageInBatch                                               = 0
//     IsDeleted                                                          = 0
//     UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday              = 0
//     UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay            = 0
//     UseMasterFirm_SymbolLimitsArray                                    = 0
//     UseMasterFirm_TradeFees                                            = 0
//     UseMasterFirm_TradeFeePerShare                                     = 0
//     UseMasterFirm_RequireSufficientMarginForNewPositions               = 0
//     UseMasterFirm_UsePercentOfMargin                                   = 0
//     UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage = 0
//     UseMasterFirm_MinimumRequiredAccountValue                          = 0
//     UseMasterFirm_MarginTimeSettings                                   = 0
//     UseMasterFirm_TradingIsDisabled                                    = 0
//     IsTradeStatisticsPublicallyShared                                  = 0
//     IsReadOnlyFollowingRequestsAllowed                                 = 0
//     IsTradeAccountSharingAllowed                                       = 0
//     ReadOnlyShareToAllSCUsernames                                      = 0
//     UseMasterFirm_SymbolCommissionsArray                               = 0
//     UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit        = 0
//     UseMasterFirm_UsePercentOfMarginForDayTrading                      = 0
//     OpenPositionsProfitLossBasedOnSettlement                           = 0
//     MarginRequirementFull                                              = 0
//     MarginRequirementFullPositionsOnly                                 = 0
//     UseMasterFirm_TradeFeesFullOverride                                = 0
//     UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders         = 0
//     UseMasterFirm_UsePercentOfMarginFullOverride                       = 0
//     UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride          = 0
//     PeakMarginRequirement                                              = 0
//     LiquidationOnlyMode                                                = 0
//     FlattenPositionsWhenUnderMarginIntradayTriggered                   = 0
//     FlattenPositionsWhenUnderMinimumAccountValueTriggered              = 0
//     AccountValueAtEndOfDayCaptureTime                                  = 0
//     EndOfDayCaptureTime                                                = 0
//     CustomerOrFirm                                                     = 0
//     MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet               = 0
//     MasterFirm_FlattenCancelWhenUnderMinimumAccountValue               = 0
//     MasterFirm_FlattenCancelWhenUnderMarginIntraday                    = 0
//     MasterFirm_FlattenCancelWhenUnderMarginEndOfDay                    = 0
//     MasterFirm_MaximumOrderQuantity                                    = 0
//     LastTriggerDateTimeUTCForDailyLossLimit                            = 0
//     OpenPositionsProfitLossIsDelayed                                   = 0
//     ExchangeTraderId                                                   = ""
// }
func (m TradeAccountDataResponsePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _TradeAccountDataResponseDefault)
}

// ToBuilder
func (m TradeAccountDataResponsePointer) ToBuilder() TradeAccountDataResponsePointerBuilder {
	return TradeAccountDataResponsePointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *TradeAccountDataResponsePointerBuilder) Finish() TradeAccountDataResponsePointer {
	return TradeAccountDataResponsePointer{m.VLSPointerBuilder.Finish()}
}

// RequestID
func (m TradeAccountDataResponsePointer) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// RequestID
func (m TradeAccountDataResponsePointerBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// TradeAccountNotExist
func (m TradeAccountDataResponsePointer) TradeAccountNotExist() uint8 {
	return message.Uint8VLS(m.Ptr, 11, 10)
}

// TradeAccountNotExist
func (m TradeAccountDataResponsePointerBuilder) TradeAccountNotExist() uint8 {
	return message.Uint8VLS(m.Ptr, 11, 10)
}

// TradeAccount
func (m TradeAccountDataResponsePointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 15, 11)
}

// TradeAccount
func (m TradeAccountDataResponsePointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 15, 11)
}

// IsSimulated
func (m TradeAccountDataResponsePointer) IsSimulated() uint8 {
	return message.Uint8VLS(m.Ptr, 16, 15)
}

// IsSimulated
func (m TradeAccountDataResponsePointerBuilder) IsSimulated() uint8 {
	return message.Uint8VLS(m.Ptr, 16, 15)
}

// CurrencyCode
func (m TradeAccountDataResponsePointer) CurrencyCode() string {
	return message.StringVLS(m.Ptr, 20, 16)
}

// CurrencyCode
func (m TradeAccountDataResponsePointerBuilder) CurrencyCode() string {
	return message.StringVLS(m.Ptr, 20, 16)
}

// CurrentCashBalance
func (m TradeAccountDataResponsePointer) CurrentCashBalance() float64 {
	return message.Float64VLS(m.Ptr, 28, 20)
}

// CurrentCashBalance
func (m TradeAccountDataResponsePointerBuilder) CurrentCashBalance() float64 {
	return message.Float64VLS(m.Ptr, 28, 20)
}

// AvailableFundsForNewPositions
func (m TradeAccountDataResponsePointer) AvailableFundsForNewPositions() float64 {
	return message.Float64VLS(m.Ptr, 36, 28)
}

// AvailableFundsForNewPositions
func (m TradeAccountDataResponsePointerBuilder) AvailableFundsForNewPositions() float64 {
	return message.Float64VLS(m.Ptr, 36, 28)
}

// MarginRequirement
func (m TradeAccountDataResponsePointer) MarginRequirement() float64 {
	return message.Float64VLS(m.Ptr, 44, 36)
}

// MarginRequirement
func (m TradeAccountDataResponsePointerBuilder) MarginRequirement() float64 {
	return message.Float64VLS(m.Ptr, 44, 36)
}

// AccountValue
func (m TradeAccountDataResponsePointer) AccountValue() float64 {
	return message.Float64VLS(m.Ptr, 52, 44)
}

// AccountValue
func (m TradeAccountDataResponsePointerBuilder) AccountValue() float64 {
	return message.Float64VLS(m.Ptr, 52, 44)
}

// OpenPositionsProfitLoss
func (m TradeAccountDataResponsePointer) OpenPositionsProfitLoss() float64 {
	return message.Float64VLS(m.Ptr, 60, 52)
}

// OpenPositionsProfitLoss
func (m TradeAccountDataResponsePointerBuilder) OpenPositionsProfitLoss() float64 {
	return message.Float64VLS(m.Ptr, 60, 52)
}

// DailyProfitLoss
func (m TradeAccountDataResponsePointer) DailyProfitLoss() float64 {
	return message.Float64VLS(m.Ptr, 68, 60)
}

// DailyProfitLoss
func (m TradeAccountDataResponsePointerBuilder) DailyProfitLoss() float64 {
	return message.Float64VLS(m.Ptr, 68, 60)
}

// TransactionIdentifierForCashBalanceAdjustment
func (m TradeAccountDataResponsePointer) TransactionIdentifierForCashBalanceAdjustment() uint64 {
	return message.Uint64VLS(m.Ptr, 76, 68)
}

// TransactionIdentifierForCashBalanceAdjustment
func (m TradeAccountDataResponsePointerBuilder) TransactionIdentifierForCashBalanceAdjustment() uint64 {
	return message.Uint64VLS(m.Ptr, 76, 68)
}

// LastTransactionDateTime
func (m TradeAccountDataResponsePointer) LastTransactionDateTime() int64 {
	return message.Int64VLS(m.Ptr, 84, 76)
}

// LastTransactionDateTime
func (m TradeAccountDataResponsePointerBuilder) LastTransactionDateTime() int64 {
	return message.Int64VLS(m.Ptr, 84, 76)
}

// TrailingAccountValueAtWhichToNotAllowNewPositions
func (m TradeAccountDataResponsePointer) TrailingAccountValueAtWhichToNotAllowNewPositions() float64 {
	return message.Float64VLS(m.Ptr, 92, 84)
}

// TrailingAccountValueAtWhichToNotAllowNewPositions
func (m TradeAccountDataResponsePointerBuilder) TrailingAccountValueAtWhichToNotAllowNewPositions() float64 {
	return message.Float64VLS(m.Ptr, 92, 84)
}

// CalculatedDailyNetLossLimitInAccountCurrency
func (m TradeAccountDataResponsePointer) CalculatedDailyNetLossLimitInAccountCurrency() float64 {
	return message.Float64VLS(m.Ptr, 100, 92)
}

// CalculatedDailyNetLossLimitInAccountCurrency
func (m TradeAccountDataResponsePointerBuilder) CalculatedDailyNetLossLimitInAccountCurrency() float64 {
	return message.Float64VLS(m.Ptr, 100, 92)
}

// DailyNetLossLimitHasBeenReached
func (m TradeAccountDataResponsePointer) DailyNetLossLimitHasBeenReached() uint8 {
	return message.Uint8VLS(m.Ptr, 101, 100)
}

// DailyNetLossLimitHasBeenReached
func (m TradeAccountDataResponsePointerBuilder) DailyNetLossLimitHasBeenReached() uint8 {
	return message.Uint8VLS(m.Ptr, 101, 100)
}

// LastResetDailyNetLossManagementVariablesDateTimeUTC
func (m TradeAccountDataResponsePointer) LastResetDailyNetLossManagementVariablesDateTimeUTC() int64 {
	return message.Int64VLS(m.Ptr, 109, 101)
}

// LastResetDailyNetLossManagementVariablesDateTimeUTC
func (m TradeAccountDataResponsePointerBuilder) LastResetDailyNetLossManagementVariablesDateTimeUTC() int64 {
	return message.Int64VLS(m.Ptr, 109, 101)
}

// IsUnderRequiredMargin
func (m TradeAccountDataResponsePointer) IsUnderRequiredMargin() uint8 {
	return message.Uint8VLS(m.Ptr, 110, 109)
}

// IsUnderRequiredMargin
func (m TradeAccountDataResponsePointerBuilder) IsUnderRequiredMargin() uint8 {
	return message.Uint8VLS(m.Ptr, 110, 109)
}

// DailyNetLossLimitInAccountCurrency
func (m TradeAccountDataResponsePointer) DailyNetLossLimitInAccountCurrency() float32 {
	return message.Float32VLS(m.Ptr, 114, 110)
}

// DailyNetLossLimitInAccountCurrency
func (m TradeAccountDataResponsePointerBuilder) DailyNetLossLimitInAccountCurrency() float32 {
	return message.Float32VLS(m.Ptr, 114, 110)
}

// PercentOfCashBalanceForDailyNetLossLimit
func (m TradeAccountDataResponsePointer) PercentOfCashBalanceForDailyNetLossLimit() int32 {
	return message.Int32VLS(m.Ptr, 118, 114)
}

// PercentOfCashBalanceForDailyNetLossLimit
func (m TradeAccountDataResponsePointerBuilder) PercentOfCashBalanceForDailyNetLossLimit() int32 {
	return message.Int32VLS(m.Ptr, 118, 114)
}

// UseTrailingAccountValueToNotAllowIncreaseInPositions
func (m TradeAccountDataResponsePointer) UseTrailingAccountValueToNotAllowIncreaseInPositions() uint8 {
	return message.Uint8VLS(m.Ptr, 119, 118)
}

// UseTrailingAccountValueToNotAllowIncreaseInPositions
func (m TradeAccountDataResponsePointerBuilder) UseTrailingAccountValueToNotAllowIncreaseInPositions() uint8 {
	return message.Uint8VLS(m.Ptr, 119, 118)
}

// DoNotAllowIncreaseInPositionsAtDailyLossLimit
func (m TradeAccountDataResponsePointer) DoNotAllowIncreaseInPositionsAtDailyLossLimit() uint8 {
	return message.Uint8VLS(m.Ptr, 120, 119)
}

// DoNotAllowIncreaseInPositionsAtDailyLossLimit
func (m TradeAccountDataResponsePointerBuilder) DoNotAllowIncreaseInPositionsAtDailyLossLimit() uint8 {
	return message.Uint8VLS(m.Ptr, 120, 119)
}

// FlattenPositionsAtDailyLossLimit
func (m TradeAccountDataResponsePointer) FlattenPositionsAtDailyLossLimit() uint8 {
	return message.Uint8VLS(m.Ptr, 121, 120)
}

// FlattenPositionsAtDailyLossLimit
func (m TradeAccountDataResponsePointerBuilder) FlattenPositionsAtDailyLossLimit() uint8 {
	return message.Uint8VLS(m.Ptr, 121, 120)
}

// ClosePositionsAtEndOfDay
func (m TradeAccountDataResponsePointer) ClosePositionsAtEndOfDay() uint8 {
	return message.Uint8VLS(m.Ptr, 122, 121)
}

// ClosePositionsAtEndOfDay
func (m TradeAccountDataResponsePointerBuilder) ClosePositionsAtEndOfDay() uint8 {
	return message.Uint8VLS(m.Ptr, 122, 121)
}

// FlattenPositionsWhenUnderMarginIntraday
func (m TradeAccountDataResponsePointer) FlattenPositionsWhenUnderMarginIntraday() uint8 {
	return message.Uint8VLS(m.Ptr, 123, 122)
}

// FlattenPositionsWhenUnderMarginIntraday
func (m TradeAccountDataResponsePointerBuilder) FlattenPositionsWhenUnderMarginIntraday() uint8 {
	return message.Uint8VLS(m.Ptr, 123, 122)
}

// FlattenPositionsWhenUnderMarginAtEndOfDay
func (m TradeAccountDataResponsePointer) FlattenPositionsWhenUnderMarginAtEndOfDay() uint8 {
	return message.Uint8VLS(m.Ptr, 124, 123)
}

// FlattenPositionsWhenUnderMarginAtEndOfDay
func (m TradeAccountDataResponsePointerBuilder) FlattenPositionsWhenUnderMarginAtEndOfDay() uint8 {
	return message.Uint8VLS(m.Ptr, 124, 123)
}

// SenderSubID
func (m TradeAccountDataResponsePointer) SenderSubID() string {
	return message.StringVLS(m.Ptr, 128, 124)
}

// SenderSubID
func (m TradeAccountDataResponsePointerBuilder) SenderSubID() string {
	return message.StringVLS(m.Ptr, 128, 124)
}

// SenderLocationId
func (m TradeAccountDataResponsePointer) SenderLocationId() string {
	return message.StringVLS(m.Ptr, 132, 128)
}

// SenderLocationId
func (m TradeAccountDataResponsePointerBuilder) SenderLocationId() string {
	return message.StringVLS(m.Ptr, 132, 128)
}

// SelfMatchPreventionID
func (m TradeAccountDataResponsePointer) SelfMatchPreventionID() string {
	return message.StringVLS(m.Ptr, 136, 132)
}

// SelfMatchPreventionID
func (m TradeAccountDataResponsePointerBuilder) SelfMatchPreventionID() string {
	return message.StringVLS(m.Ptr, 136, 132)
}

// CTICode
func (m TradeAccountDataResponsePointer) CTICode() int32 {
	return message.Int32VLS(m.Ptr, 140, 136)
}

// CTICode
func (m TradeAccountDataResponsePointerBuilder) CTICode() int32 {
	return message.Int32VLS(m.Ptr, 140, 136)
}

// TradeAccountIsReadOnly
func (m TradeAccountDataResponsePointer) TradeAccountIsReadOnly() uint8 {
	return message.Uint8VLS(m.Ptr, 141, 140)
}

// TradeAccountIsReadOnly
func (m TradeAccountDataResponsePointerBuilder) TradeAccountIsReadOnly() uint8 {
	return message.Uint8VLS(m.Ptr, 141, 140)
}

// MaximumGlobalPositionQuantity
func (m TradeAccountDataResponsePointer) MaximumGlobalPositionQuantity() int32 {
	return message.Int32VLS(m.Ptr, 145, 141)
}

// MaximumGlobalPositionQuantity
func (m TradeAccountDataResponsePointerBuilder) MaximumGlobalPositionQuantity() int32 {
	return message.Int32VLS(m.Ptr, 145, 141)
}

// TradeFeePerContract
func (m TradeAccountDataResponsePointer) TradeFeePerContract() float64 {
	return message.Float64VLS(m.Ptr, 153, 145)
}

// TradeFeePerContract
func (m TradeAccountDataResponsePointerBuilder) TradeFeePerContract() float64 {
	return message.Float64VLS(m.Ptr, 153, 145)
}

// TradeFeePerShare
func (m TradeAccountDataResponsePointer) TradeFeePerShare() float64 {
	return message.Float64VLS(m.Ptr, 161, 153)
}

// TradeFeePerShare
func (m TradeAccountDataResponsePointerBuilder) TradeFeePerShare() float64 {
	return message.Float64VLS(m.Ptr, 161, 153)
}

// RequireSufficientMarginForNewPositions
func (m TradeAccountDataResponsePointer) RequireSufficientMarginForNewPositions() uint8 {
	return message.Uint8VLS(m.Ptr, 162, 161)
}

// RequireSufficientMarginForNewPositions
func (m TradeAccountDataResponsePointerBuilder) RequireSufficientMarginForNewPositions() uint8 {
	return message.Uint8VLS(m.Ptr, 162, 161)
}

// UsePercentOfMargin
func (m TradeAccountDataResponsePointer) UsePercentOfMargin() int32 {
	return message.Int32VLS(m.Ptr, 166, 162)
}

// UsePercentOfMargin
func (m TradeAccountDataResponsePointerBuilder) UsePercentOfMargin() int32 {
	return message.Int32VLS(m.Ptr, 166, 162)
}

// UsePercentOfMarginForDayTrading
func (m TradeAccountDataResponsePointer) UsePercentOfMarginForDayTrading() int32 {
	return message.Int32VLS(m.Ptr, 170, 166)
}

// UsePercentOfMarginForDayTrading
func (m TradeAccountDataResponsePointerBuilder) UsePercentOfMarginForDayTrading() int32 {
	return message.Int32VLS(m.Ptr, 170, 166)
}

// MaximumAllowedAccountBalanceForPositionsAsPercentage
func (m TradeAccountDataResponsePointer) MaximumAllowedAccountBalanceForPositionsAsPercentage() int32 {
	return message.Int32VLS(m.Ptr, 174, 170)
}

// MaximumAllowedAccountBalanceForPositionsAsPercentage
func (m TradeAccountDataResponsePointerBuilder) MaximumAllowedAccountBalanceForPositionsAsPercentage() int32 {
	return message.Int32VLS(m.Ptr, 174, 170)
}

// FirmID
func (m TradeAccountDataResponsePointer) FirmID() string {
	return message.StringVLS(m.Ptr, 178, 174)
}

// FirmID
func (m TradeAccountDataResponsePointerBuilder) FirmID() string {
	return message.StringVLS(m.Ptr, 178, 174)
}

// TradingIsDisabled
func (m TradeAccountDataResponsePointer) TradingIsDisabled() uint8 {
	return message.Uint8VLS(m.Ptr, 179, 178)
}

// TradingIsDisabled
func (m TradeAccountDataResponsePointerBuilder) TradingIsDisabled() uint8 {
	return message.Uint8VLS(m.Ptr, 179, 178)
}

// DescriptiveName
func (m TradeAccountDataResponsePointer) DescriptiveName() string {
	return message.StringVLS(m.Ptr, 183, 179)
}

// DescriptiveName
func (m TradeAccountDataResponsePointerBuilder) DescriptiveName() string {
	return message.StringVLS(m.Ptr, 183, 179)
}

// IsMasterFirmControlAccount
func (m TradeAccountDataResponsePointer) IsMasterFirmControlAccount() uint8 {
	return message.Uint8VLS(m.Ptr, 184, 183)
}

// IsMasterFirmControlAccount
func (m TradeAccountDataResponsePointerBuilder) IsMasterFirmControlAccount() uint8 {
	return message.Uint8VLS(m.Ptr, 184, 183)
}

// MinimumRequiredAccountValue
func (m TradeAccountDataResponsePointer) MinimumRequiredAccountValue() float64 {
	return message.Float64VLS(m.Ptr, 192, 184)
}

// MinimumRequiredAccountValue
func (m TradeAccountDataResponsePointerBuilder) MinimumRequiredAccountValue() float64 {
	return message.Float64VLS(m.Ptr, 192, 184)
}

// BeginTimeForDayMargin
func (m TradeAccountDataResponsePointer) BeginTimeForDayMargin() int64 {
	return message.Int64VLS(m.Ptr, 200, 192)
}

// BeginTimeForDayMargin
func (m TradeAccountDataResponsePointerBuilder) BeginTimeForDayMargin() int64 {
	return message.Int64VLS(m.Ptr, 200, 192)
}

// EndTimeForDayMargin
func (m TradeAccountDataResponsePointer) EndTimeForDayMargin() int64 {
	return message.Int64VLS(m.Ptr, 208, 200)
}

// EndTimeForDayMargin
func (m TradeAccountDataResponsePointerBuilder) EndTimeForDayMargin() int64 {
	return message.Int64VLS(m.Ptr, 208, 200)
}

// DayMarginTimeZone
func (m TradeAccountDataResponsePointer) DayMarginTimeZone() string {
	return message.StringVLS(m.Ptr, 212, 208)
}

// DayMarginTimeZone
func (m TradeAccountDataResponsePointerBuilder) DayMarginTimeZone() string {
	return message.StringVLS(m.Ptr, 212, 208)
}

// IsSnapshot
func (m TradeAccountDataResponsePointer) IsSnapshot() uint8 {
	return message.Uint8VLS(m.Ptr, 213, 212)
}

// IsSnapshot
func (m TradeAccountDataResponsePointerBuilder) IsSnapshot() uint8 {
	return message.Uint8VLS(m.Ptr, 213, 212)
}

// IsFirstMessageInBatch
func (m TradeAccountDataResponsePointer) IsFirstMessageInBatch() uint8 {
	return message.Uint8VLS(m.Ptr, 214, 213)
}

// IsFirstMessageInBatch
func (m TradeAccountDataResponsePointerBuilder) IsFirstMessageInBatch() uint8 {
	return message.Uint8VLS(m.Ptr, 214, 213)
}

// IsLastMessageInBatch
func (m TradeAccountDataResponsePointer) IsLastMessageInBatch() uint8 {
	return message.Uint8VLS(m.Ptr, 215, 214)
}

// IsLastMessageInBatch
func (m TradeAccountDataResponsePointerBuilder) IsLastMessageInBatch() uint8 {
	return message.Uint8VLS(m.Ptr, 215, 214)
}

// IsDeleted
func (m TradeAccountDataResponsePointer) IsDeleted() uint8 {
	return message.Uint8VLS(m.Ptr, 216, 215)
}

// IsDeleted
func (m TradeAccountDataResponsePointerBuilder) IsDeleted() uint8 {
	return message.Uint8VLS(m.Ptr, 216, 215)
}

// UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday
func (m TradeAccountDataResponsePointer) UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday() uint8 {
	return message.Uint8VLS(m.Ptr, 217, 216)
}

// UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday
func (m TradeAccountDataResponsePointerBuilder) UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday() uint8 {
	return message.Uint8VLS(m.Ptr, 217, 216)
}

// UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay
func (m TradeAccountDataResponsePointer) UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay() uint8 {
	return message.Uint8VLS(m.Ptr, 218, 217)
}

// UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay
func (m TradeAccountDataResponsePointerBuilder) UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay() uint8 {
	return message.Uint8VLS(m.Ptr, 218, 217)
}

// UseMasterFirm_SymbolLimitsArray
func (m TradeAccountDataResponsePointer) UseMasterFirm_SymbolLimitsArray() uint8 {
	return message.Uint8VLS(m.Ptr, 219, 218)
}

// UseMasterFirm_SymbolLimitsArray
func (m TradeAccountDataResponsePointerBuilder) UseMasterFirm_SymbolLimitsArray() uint8 {
	return message.Uint8VLS(m.Ptr, 219, 218)
}

// UseMasterFirm_TradeFees
func (m TradeAccountDataResponsePointer) UseMasterFirm_TradeFees() uint8 {
	return message.Uint8VLS(m.Ptr, 220, 219)
}

// UseMasterFirm_TradeFees
func (m TradeAccountDataResponsePointerBuilder) UseMasterFirm_TradeFees() uint8 {
	return message.Uint8VLS(m.Ptr, 220, 219)
}

// UseMasterFirm_TradeFeePerShare
func (m TradeAccountDataResponsePointer) UseMasterFirm_TradeFeePerShare() uint8 {
	return message.Uint8VLS(m.Ptr, 221, 220)
}

// UseMasterFirm_TradeFeePerShare
func (m TradeAccountDataResponsePointerBuilder) UseMasterFirm_TradeFeePerShare() uint8 {
	return message.Uint8VLS(m.Ptr, 221, 220)
}

// UseMasterFirm_RequireSufficientMarginForNewPositions
func (m TradeAccountDataResponsePointer) UseMasterFirm_RequireSufficientMarginForNewPositions() uint8 {
	return message.Uint8VLS(m.Ptr, 222, 221)
}

// UseMasterFirm_RequireSufficientMarginForNewPositions
func (m TradeAccountDataResponsePointerBuilder) UseMasterFirm_RequireSufficientMarginForNewPositions() uint8 {
	return message.Uint8VLS(m.Ptr, 222, 221)
}

// UseMasterFirm_UsePercentOfMargin
func (m TradeAccountDataResponsePointer) UseMasterFirm_UsePercentOfMargin() uint8 {
	return message.Uint8VLS(m.Ptr, 223, 222)
}

// UseMasterFirm_UsePercentOfMargin
func (m TradeAccountDataResponsePointerBuilder) UseMasterFirm_UsePercentOfMargin() uint8 {
	return message.Uint8VLS(m.Ptr, 223, 222)
}

// UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage
func (m TradeAccountDataResponsePointer) UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage() uint8 {
	return message.Uint8VLS(m.Ptr, 224, 223)
}

// UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage
func (m TradeAccountDataResponsePointerBuilder) UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage() uint8 {
	return message.Uint8VLS(m.Ptr, 224, 223)
}

// UseMasterFirm_MinimumRequiredAccountValue
func (m TradeAccountDataResponsePointer) UseMasterFirm_MinimumRequiredAccountValue() uint8 {
	return message.Uint8VLS(m.Ptr, 225, 224)
}

// UseMasterFirm_MinimumRequiredAccountValue
func (m TradeAccountDataResponsePointerBuilder) UseMasterFirm_MinimumRequiredAccountValue() uint8 {
	return message.Uint8VLS(m.Ptr, 225, 224)
}

// UseMasterFirm_MarginTimeSettings
func (m TradeAccountDataResponsePointer) UseMasterFirm_MarginTimeSettings() uint8 {
	return message.Uint8VLS(m.Ptr, 226, 225)
}

// UseMasterFirm_MarginTimeSettings
func (m TradeAccountDataResponsePointerBuilder) UseMasterFirm_MarginTimeSettings() uint8 {
	return message.Uint8VLS(m.Ptr, 226, 225)
}

// UseMasterFirm_TradingIsDisabled
func (m TradeAccountDataResponsePointer) UseMasterFirm_TradingIsDisabled() uint8 {
	return message.Uint8VLS(m.Ptr, 227, 226)
}

// UseMasterFirm_TradingIsDisabled
func (m TradeAccountDataResponsePointerBuilder) UseMasterFirm_TradingIsDisabled() uint8 {
	return message.Uint8VLS(m.Ptr, 227, 226)
}

// IsTradeStatisticsPublicallyShared
func (m TradeAccountDataResponsePointer) IsTradeStatisticsPublicallyShared() uint8 {
	return message.Uint8VLS(m.Ptr, 228, 227)
}

// IsTradeStatisticsPublicallyShared
func (m TradeAccountDataResponsePointerBuilder) IsTradeStatisticsPublicallyShared() uint8 {
	return message.Uint8VLS(m.Ptr, 228, 227)
}

// IsReadOnlyFollowingRequestsAllowed
func (m TradeAccountDataResponsePointer) IsReadOnlyFollowingRequestsAllowed() uint8 {
	return message.Uint8VLS(m.Ptr, 229, 228)
}

// IsReadOnlyFollowingRequestsAllowed
func (m TradeAccountDataResponsePointerBuilder) IsReadOnlyFollowingRequestsAllowed() uint8 {
	return message.Uint8VLS(m.Ptr, 229, 228)
}

// IsTradeAccountSharingAllowed
func (m TradeAccountDataResponsePointer) IsTradeAccountSharingAllowed() uint8 {
	return message.Uint8VLS(m.Ptr, 230, 229)
}

// IsTradeAccountSharingAllowed
func (m TradeAccountDataResponsePointerBuilder) IsTradeAccountSharingAllowed() uint8 {
	return message.Uint8VLS(m.Ptr, 230, 229)
}

// ReadOnlyShareToAllSCUsernames
func (m TradeAccountDataResponsePointer) ReadOnlyShareToAllSCUsernames() uint8 {
	return message.Uint8VLS(m.Ptr, 231, 230)
}

// ReadOnlyShareToAllSCUsernames
func (m TradeAccountDataResponsePointerBuilder) ReadOnlyShareToAllSCUsernames() uint8 {
	return message.Uint8VLS(m.Ptr, 231, 230)
}

// UseMasterFirm_SymbolCommissionsArray
func (m TradeAccountDataResponsePointer) UseMasterFirm_SymbolCommissionsArray() uint8 {
	return message.Uint8VLS(m.Ptr, 232, 231)
}

// UseMasterFirm_SymbolCommissionsArray
func (m TradeAccountDataResponsePointerBuilder) UseMasterFirm_SymbolCommissionsArray() uint8 {
	return message.Uint8VLS(m.Ptr, 232, 231)
}

// UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit
func (m TradeAccountDataResponsePointer) UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit() uint8 {
	return message.Uint8VLS(m.Ptr, 233, 232)
}

// UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit
func (m TradeAccountDataResponsePointerBuilder) UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit() uint8 {
	return message.Uint8VLS(m.Ptr, 233, 232)
}

// UseMasterFirm_UsePercentOfMarginForDayTrading
func (m TradeAccountDataResponsePointer) UseMasterFirm_UsePercentOfMarginForDayTrading() uint8 {
	return message.Uint8VLS(m.Ptr, 234, 233)
}

// UseMasterFirm_UsePercentOfMarginForDayTrading
func (m TradeAccountDataResponsePointerBuilder) UseMasterFirm_UsePercentOfMarginForDayTrading() uint8 {
	return message.Uint8VLS(m.Ptr, 234, 233)
}

// OpenPositionsProfitLossBasedOnSettlement
func (m TradeAccountDataResponsePointer) OpenPositionsProfitLossBasedOnSettlement() float64 {
	return message.Float64VLS(m.Ptr, 242, 234)
}

// OpenPositionsProfitLossBasedOnSettlement
func (m TradeAccountDataResponsePointerBuilder) OpenPositionsProfitLossBasedOnSettlement() float64 {
	return message.Float64VLS(m.Ptr, 242, 234)
}

// MarginRequirementFull
func (m TradeAccountDataResponsePointer) MarginRequirementFull() float64 {
	return message.Float64VLS(m.Ptr, 250, 242)
}

// MarginRequirementFull
func (m TradeAccountDataResponsePointerBuilder) MarginRequirementFull() float64 {
	return message.Float64VLS(m.Ptr, 250, 242)
}

// MarginRequirementFullPositionsOnly
func (m TradeAccountDataResponsePointer) MarginRequirementFullPositionsOnly() float64 {
	return message.Float64VLS(m.Ptr, 258, 250)
}

// MarginRequirementFullPositionsOnly
func (m TradeAccountDataResponsePointerBuilder) MarginRequirementFullPositionsOnly() float64 {
	return message.Float64VLS(m.Ptr, 258, 250)
}

// UseMasterFirm_TradeFeesFullOverride
func (m TradeAccountDataResponsePointer) UseMasterFirm_TradeFeesFullOverride() uint8 {
	return message.Uint8VLS(m.Ptr, 259, 258)
}

// UseMasterFirm_TradeFeesFullOverride
func (m TradeAccountDataResponsePointerBuilder) UseMasterFirm_TradeFeesFullOverride() uint8 {
	return message.Uint8VLS(m.Ptr, 259, 258)
}

// UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders
func (m TradeAccountDataResponsePointer) UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders() uint8 {
	return message.Uint8VLS(m.Ptr, 260, 259)
}

// UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders
func (m TradeAccountDataResponsePointerBuilder) UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders() uint8 {
	return message.Uint8VLS(m.Ptr, 260, 259)
}

// UseMasterFirm_UsePercentOfMarginFullOverride
func (m TradeAccountDataResponsePointer) UseMasterFirm_UsePercentOfMarginFullOverride() uint8 {
	return message.Uint8VLS(m.Ptr, 261, 260)
}

// UseMasterFirm_UsePercentOfMarginFullOverride
func (m TradeAccountDataResponsePointerBuilder) UseMasterFirm_UsePercentOfMarginFullOverride() uint8 {
	return message.Uint8VLS(m.Ptr, 261, 260)
}

// UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride
func (m TradeAccountDataResponsePointer) UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride() uint8 {
	return message.Uint8VLS(m.Ptr, 262, 261)
}

// UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride
func (m TradeAccountDataResponsePointerBuilder) UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride() uint8 {
	return message.Uint8VLS(m.Ptr, 262, 261)
}

// PeakMarginRequirement
func (m TradeAccountDataResponsePointer) PeakMarginRequirement() float64 {
	return message.Float64VLS(m.Ptr, 270, 262)
}

// PeakMarginRequirement
func (m TradeAccountDataResponsePointerBuilder) PeakMarginRequirement() float64 {
	return message.Float64VLS(m.Ptr, 270, 262)
}

// LiquidationOnlyMode
func (m TradeAccountDataResponsePointer) LiquidationOnlyMode() uint8 {
	return message.Uint8VLS(m.Ptr, 271, 270)
}

// LiquidationOnlyMode
func (m TradeAccountDataResponsePointerBuilder) LiquidationOnlyMode() uint8 {
	return message.Uint8VLS(m.Ptr, 271, 270)
}

// FlattenPositionsWhenUnderMarginIntradayTriggered
func (m TradeAccountDataResponsePointer) FlattenPositionsWhenUnderMarginIntradayTriggered() uint8 {
	return message.Uint8VLS(m.Ptr, 272, 271)
}

// FlattenPositionsWhenUnderMarginIntradayTriggered
func (m TradeAccountDataResponsePointerBuilder) FlattenPositionsWhenUnderMarginIntradayTriggered() uint8 {
	return message.Uint8VLS(m.Ptr, 272, 271)
}

// FlattenPositionsWhenUnderMinimumAccountValueTriggered
func (m TradeAccountDataResponsePointer) FlattenPositionsWhenUnderMinimumAccountValueTriggered() uint8 {
	return message.Uint8VLS(m.Ptr, 273, 272)
}

// FlattenPositionsWhenUnderMinimumAccountValueTriggered
func (m TradeAccountDataResponsePointerBuilder) FlattenPositionsWhenUnderMinimumAccountValueTriggered() uint8 {
	return message.Uint8VLS(m.Ptr, 273, 272)
}

// AccountValueAtEndOfDayCaptureTime
func (m TradeAccountDataResponsePointer) AccountValueAtEndOfDayCaptureTime() float64 {
	return message.Float64VLS(m.Ptr, 281, 273)
}

// AccountValueAtEndOfDayCaptureTime
func (m TradeAccountDataResponsePointerBuilder) AccountValueAtEndOfDayCaptureTime() float64 {
	return message.Float64VLS(m.Ptr, 281, 273)
}

// EndOfDayCaptureTime
func (m TradeAccountDataResponsePointer) EndOfDayCaptureTime() int64 {
	return message.Int64VLS(m.Ptr, 289, 281)
}

// EndOfDayCaptureTime
func (m TradeAccountDataResponsePointerBuilder) EndOfDayCaptureTime() int64 {
	return message.Int64VLS(m.Ptr, 289, 281)
}

// CustomerOrFirm
func (m TradeAccountDataResponsePointer) CustomerOrFirm() uint8 {
	return message.Uint8VLS(m.Ptr, 290, 289)
}

// CustomerOrFirm
func (m TradeAccountDataResponsePointerBuilder) CustomerOrFirm() uint8 {
	return message.Uint8VLS(m.Ptr, 290, 289)
}

// MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet
func (m TradeAccountDataResponsePointer) MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet() uint8 {
	return message.Uint8VLS(m.Ptr, 291, 290)
}

// MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet
func (m TradeAccountDataResponsePointerBuilder) MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet() uint8 {
	return message.Uint8VLS(m.Ptr, 291, 290)
}

// MasterFirm_FlattenCancelWhenUnderMinimumAccountValue
func (m TradeAccountDataResponsePointer) MasterFirm_FlattenCancelWhenUnderMinimumAccountValue() uint8 {
	return message.Uint8VLS(m.Ptr, 292, 291)
}

// MasterFirm_FlattenCancelWhenUnderMinimumAccountValue
func (m TradeAccountDataResponsePointerBuilder) MasterFirm_FlattenCancelWhenUnderMinimumAccountValue() uint8 {
	return message.Uint8VLS(m.Ptr, 292, 291)
}

// MasterFirm_FlattenCancelWhenUnderMarginIntraday
func (m TradeAccountDataResponsePointer) MasterFirm_FlattenCancelWhenUnderMarginIntraday() uint8 {
	return message.Uint8VLS(m.Ptr, 293, 292)
}

// MasterFirm_FlattenCancelWhenUnderMarginIntraday
func (m TradeAccountDataResponsePointerBuilder) MasterFirm_FlattenCancelWhenUnderMarginIntraday() uint8 {
	return message.Uint8VLS(m.Ptr, 293, 292)
}

// MasterFirm_FlattenCancelWhenUnderMarginEndOfDay
func (m TradeAccountDataResponsePointer) MasterFirm_FlattenCancelWhenUnderMarginEndOfDay() uint8 {
	return message.Uint8VLS(m.Ptr, 294, 293)
}

// MasterFirm_FlattenCancelWhenUnderMarginEndOfDay
func (m TradeAccountDataResponsePointerBuilder) MasterFirm_FlattenCancelWhenUnderMarginEndOfDay() uint8 {
	return message.Uint8VLS(m.Ptr, 294, 293)
}

// MasterFirm_MaximumOrderQuantity
func (m TradeAccountDataResponsePointer) MasterFirm_MaximumOrderQuantity() uint32 {
	return message.Uint32VLS(m.Ptr, 298, 294)
}

// MasterFirm_MaximumOrderQuantity
func (m TradeAccountDataResponsePointerBuilder) MasterFirm_MaximumOrderQuantity() uint32 {
	return message.Uint32VLS(m.Ptr, 298, 294)
}

// LastTriggerDateTimeUTCForDailyLossLimit
func (m TradeAccountDataResponsePointer) LastTriggerDateTimeUTCForDailyLossLimit() int64 {
	return message.Int64VLS(m.Ptr, 306, 298)
}

// LastTriggerDateTimeUTCForDailyLossLimit
func (m TradeAccountDataResponsePointerBuilder) LastTriggerDateTimeUTCForDailyLossLimit() int64 {
	return message.Int64VLS(m.Ptr, 306, 298)
}

// OpenPositionsProfitLossIsDelayed
func (m TradeAccountDataResponsePointer) OpenPositionsProfitLossIsDelayed() uint8 {
	return message.Uint8VLS(m.Ptr, 307, 306)
}

// OpenPositionsProfitLossIsDelayed
func (m TradeAccountDataResponsePointerBuilder) OpenPositionsProfitLossIsDelayed() uint8 {
	return message.Uint8VLS(m.Ptr, 307, 306)
}

// ExchangeTraderId
func (m TradeAccountDataResponsePointer) ExchangeTraderId() string {
	return message.StringVLS(m.Ptr, 311, 307)
}

// ExchangeTraderId
func (m TradeAccountDataResponsePointerBuilder) ExchangeTraderId() string {
	return message.StringVLS(m.Ptr, 311, 307)
}

// SetRequestID
func (m TradeAccountDataResponsePointerBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Ptr, 10, 6, value)
}

// SetTradeAccountNotExist
func (m TradeAccountDataResponsePointerBuilder) SetTradeAccountNotExist(value uint8) {
	message.SetUint8VLS(m.Ptr, 11, 10, value)
}

// SetTradeAccount
func (m *TradeAccountDataResponsePointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 15, 11, value)
}

// SetIsSimulated
func (m TradeAccountDataResponsePointerBuilder) SetIsSimulated(value uint8) {
	message.SetUint8VLS(m.Ptr, 16, 15, value)
}

// SetCurrencyCode
func (m *TradeAccountDataResponsePointerBuilder) SetCurrencyCode(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 20, 16, value)
}

// SetCurrentCashBalance
func (m TradeAccountDataResponsePointerBuilder) SetCurrentCashBalance(value float64) {
	message.SetFloat64VLS(m.Ptr, 28, 20, value)
}

// SetAvailableFundsForNewPositions
func (m TradeAccountDataResponsePointerBuilder) SetAvailableFundsForNewPositions(value float64) {
	message.SetFloat64VLS(m.Ptr, 36, 28, value)
}

// SetMarginRequirement
func (m TradeAccountDataResponsePointerBuilder) SetMarginRequirement(value float64) {
	message.SetFloat64VLS(m.Ptr, 44, 36, value)
}

// SetAccountValue
func (m TradeAccountDataResponsePointerBuilder) SetAccountValue(value float64) {
	message.SetFloat64VLS(m.Ptr, 52, 44, value)
}

// SetOpenPositionsProfitLoss
func (m TradeAccountDataResponsePointerBuilder) SetOpenPositionsProfitLoss(value float64) {
	message.SetFloat64VLS(m.Ptr, 60, 52, value)
}

// SetDailyProfitLoss
func (m TradeAccountDataResponsePointerBuilder) SetDailyProfitLoss(value float64) {
	message.SetFloat64VLS(m.Ptr, 68, 60, value)
}

// SetTransactionIdentifierForCashBalanceAdjustment
func (m TradeAccountDataResponsePointerBuilder) SetTransactionIdentifierForCashBalanceAdjustment(value uint64) {
	message.SetUint64VLS(m.Ptr, 76, 68, value)
}

// SetLastTransactionDateTime
func (m TradeAccountDataResponsePointerBuilder) SetLastTransactionDateTime(value int64) {
	message.SetInt64VLS(m.Ptr, 84, 76, value)
}

// SetTrailingAccountValueAtWhichToNotAllowNewPositions
func (m TradeAccountDataResponsePointerBuilder) SetTrailingAccountValueAtWhichToNotAllowNewPositions(value float64) {
	message.SetFloat64VLS(m.Ptr, 92, 84, value)
}

// SetCalculatedDailyNetLossLimitInAccountCurrency
func (m TradeAccountDataResponsePointerBuilder) SetCalculatedDailyNetLossLimitInAccountCurrency(value float64) {
	message.SetFloat64VLS(m.Ptr, 100, 92, value)
}

// SetDailyNetLossLimitHasBeenReached
func (m TradeAccountDataResponsePointerBuilder) SetDailyNetLossLimitHasBeenReached(value uint8) {
	message.SetUint8VLS(m.Ptr, 101, 100, value)
}

// SetLastResetDailyNetLossManagementVariablesDateTimeUTC
func (m TradeAccountDataResponsePointerBuilder) SetLastResetDailyNetLossManagementVariablesDateTimeUTC(value int64) {
	message.SetInt64VLS(m.Ptr, 109, 101, value)
}

// SetIsUnderRequiredMargin
func (m TradeAccountDataResponsePointerBuilder) SetIsUnderRequiredMargin(value uint8) {
	message.SetUint8VLS(m.Ptr, 110, 109, value)
}

// SetDailyNetLossLimitInAccountCurrency
func (m TradeAccountDataResponsePointerBuilder) SetDailyNetLossLimitInAccountCurrency(value float32) {
	message.SetFloat32VLS(m.Ptr, 114, 110, value)
}

// SetPercentOfCashBalanceForDailyNetLossLimit
func (m TradeAccountDataResponsePointerBuilder) SetPercentOfCashBalanceForDailyNetLossLimit(value int32) {
	message.SetInt32VLS(m.Ptr, 118, 114, value)
}

// SetUseTrailingAccountValueToNotAllowIncreaseInPositions
func (m TradeAccountDataResponsePointerBuilder) SetUseTrailingAccountValueToNotAllowIncreaseInPositions(value uint8) {
	message.SetUint8VLS(m.Ptr, 119, 118, value)
}

// SetDoNotAllowIncreaseInPositionsAtDailyLossLimit
func (m TradeAccountDataResponsePointerBuilder) SetDoNotAllowIncreaseInPositionsAtDailyLossLimit(value uint8) {
	message.SetUint8VLS(m.Ptr, 120, 119, value)
}

// SetFlattenPositionsAtDailyLossLimit
func (m TradeAccountDataResponsePointerBuilder) SetFlattenPositionsAtDailyLossLimit(value uint8) {
	message.SetUint8VLS(m.Ptr, 121, 120, value)
}

// SetClosePositionsAtEndOfDay
func (m TradeAccountDataResponsePointerBuilder) SetClosePositionsAtEndOfDay(value uint8) {
	message.SetUint8VLS(m.Ptr, 122, 121, value)
}

// SetFlattenPositionsWhenUnderMarginIntraday
func (m TradeAccountDataResponsePointerBuilder) SetFlattenPositionsWhenUnderMarginIntraday(value uint8) {
	message.SetUint8VLS(m.Ptr, 123, 122, value)
}

// SetFlattenPositionsWhenUnderMarginAtEndOfDay
func (m TradeAccountDataResponsePointerBuilder) SetFlattenPositionsWhenUnderMarginAtEndOfDay(value uint8) {
	message.SetUint8VLS(m.Ptr, 124, 123, value)
}

// SetSenderSubID
func (m *TradeAccountDataResponsePointerBuilder) SetSenderSubID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 128, 124, value)
}

// SetSenderLocationId
func (m *TradeAccountDataResponsePointerBuilder) SetSenderLocationId(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 132, 128, value)
}

// SetSelfMatchPreventionID
func (m *TradeAccountDataResponsePointerBuilder) SetSelfMatchPreventionID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 136, 132, value)
}

// SetCTICode
func (m TradeAccountDataResponsePointerBuilder) SetCTICode(value int32) {
	message.SetInt32VLS(m.Ptr, 140, 136, value)
}

// SetTradeAccountIsReadOnly
func (m TradeAccountDataResponsePointerBuilder) SetTradeAccountIsReadOnly(value uint8) {
	message.SetUint8VLS(m.Ptr, 141, 140, value)
}

// SetMaximumGlobalPositionQuantity
func (m TradeAccountDataResponsePointerBuilder) SetMaximumGlobalPositionQuantity(value int32) {
	message.SetInt32VLS(m.Ptr, 145, 141, value)
}

// SetTradeFeePerContract
func (m TradeAccountDataResponsePointerBuilder) SetTradeFeePerContract(value float64) {
	message.SetFloat64VLS(m.Ptr, 153, 145, value)
}

// SetTradeFeePerShare
func (m TradeAccountDataResponsePointerBuilder) SetTradeFeePerShare(value float64) {
	message.SetFloat64VLS(m.Ptr, 161, 153, value)
}

// SetRequireSufficientMarginForNewPositions
func (m TradeAccountDataResponsePointerBuilder) SetRequireSufficientMarginForNewPositions(value uint8) {
	message.SetUint8VLS(m.Ptr, 162, 161, value)
}

// SetUsePercentOfMargin
func (m TradeAccountDataResponsePointerBuilder) SetUsePercentOfMargin(value int32) {
	message.SetInt32VLS(m.Ptr, 166, 162, value)
}

// SetUsePercentOfMarginForDayTrading
func (m TradeAccountDataResponsePointerBuilder) SetUsePercentOfMarginForDayTrading(value int32) {
	message.SetInt32VLS(m.Ptr, 170, 166, value)
}

// SetMaximumAllowedAccountBalanceForPositionsAsPercentage
func (m TradeAccountDataResponsePointerBuilder) SetMaximumAllowedAccountBalanceForPositionsAsPercentage(value int32) {
	message.SetInt32VLS(m.Ptr, 174, 170, value)
}

// SetFirmID
func (m *TradeAccountDataResponsePointerBuilder) SetFirmID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 178, 174, value)
}

// SetTradingIsDisabled
func (m TradeAccountDataResponsePointerBuilder) SetTradingIsDisabled(value uint8) {
	message.SetUint8VLS(m.Ptr, 179, 178, value)
}

// SetDescriptiveName
func (m *TradeAccountDataResponsePointerBuilder) SetDescriptiveName(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 183, 179, value)
}

// SetIsMasterFirmControlAccount
func (m TradeAccountDataResponsePointerBuilder) SetIsMasterFirmControlAccount(value uint8) {
	message.SetUint8VLS(m.Ptr, 184, 183, value)
}

// SetMinimumRequiredAccountValue
func (m TradeAccountDataResponsePointerBuilder) SetMinimumRequiredAccountValue(value float64) {
	message.SetFloat64VLS(m.Ptr, 192, 184, value)
}

// SetBeginTimeForDayMargin
func (m TradeAccountDataResponsePointerBuilder) SetBeginTimeForDayMargin(value int64) {
	message.SetInt64VLS(m.Ptr, 200, 192, value)
}

// SetEndTimeForDayMargin
func (m TradeAccountDataResponsePointerBuilder) SetEndTimeForDayMargin(value int64) {
	message.SetInt64VLS(m.Ptr, 208, 200, value)
}

// SetDayMarginTimeZone
func (m *TradeAccountDataResponsePointerBuilder) SetDayMarginTimeZone(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 212, 208, value)
}

// SetIsSnapshot
func (m TradeAccountDataResponsePointerBuilder) SetIsSnapshot(value uint8) {
	message.SetUint8VLS(m.Ptr, 213, 212, value)
}

// SetIsFirstMessageInBatch
func (m TradeAccountDataResponsePointerBuilder) SetIsFirstMessageInBatch(value uint8) {
	message.SetUint8VLS(m.Ptr, 214, 213, value)
}

// SetIsLastMessageInBatch
func (m TradeAccountDataResponsePointerBuilder) SetIsLastMessageInBatch(value uint8) {
	message.SetUint8VLS(m.Ptr, 215, 214, value)
}

// SetIsDeleted
func (m TradeAccountDataResponsePointerBuilder) SetIsDeleted(value uint8) {
	message.SetUint8VLS(m.Ptr, 216, 215, value)
}

// SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntraday
func (m TradeAccountDataResponsePointerBuilder) SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntraday(value uint8) {
	message.SetUint8VLS(m.Ptr, 217, 216, value)
}

// SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay
func (m TradeAccountDataResponsePointerBuilder) SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay(value uint8) {
	message.SetUint8VLS(m.Ptr, 218, 217, value)
}

// SetUseMasterFirm_SymbolLimitsArray
func (m TradeAccountDataResponsePointerBuilder) SetUseMasterFirm_SymbolLimitsArray(value uint8) {
	message.SetUint8VLS(m.Ptr, 219, 218, value)
}

// SetUseMasterFirm_TradeFees
func (m TradeAccountDataResponsePointerBuilder) SetUseMasterFirm_TradeFees(value uint8) {
	message.SetUint8VLS(m.Ptr, 220, 219, value)
}

// SetUseMasterFirm_TradeFeePerShare
func (m TradeAccountDataResponsePointerBuilder) SetUseMasterFirm_TradeFeePerShare(value uint8) {
	message.SetUint8VLS(m.Ptr, 221, 220, value)
}

// SetUseMasterFirm_RequireSufficientMarginForNewPositions
func (m TradeAccountDataResponsePointerBuilder) SetUseMasterFirm_RequireSufficientMarginForNewPositions(value uint8) {
	message.SetUint8VLS(m.Ptr, 222, 221, value)
}

// SetUseMasterFirm_UsePercentOfMargin
func (m TradeAccountDataResponsePointerBuilder) SetUseMasterFirm_UsePercentOfMargin(value uint8) {
	message.SetUint8VLS(m.Ptr, 223, 222, value)
}

// SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage
func (m TradeAccountDataResponsePointerBuilder) SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage(value uint8) {
	message.SetUint8VLS(m.Ptr, 224, 223, value)
}

// SetUseMasterFirm_MinimumRequiredAccountValue
func (m TradeAccountDataResponsePointerBuilder) SetUseMasterFirm_MinimumRequiredAccountValue(value uint8) {
	message.SetUint8VLS(m.Ptr, 225, 224, value)
}

// SetUseMasterFirm_MarginTimeSettings
func (m TradeAccountDataResponsePointerBuilder) SetUseMasterFirm_MarginTimeSettings(value uint8) {
	message.SetUint8VLS(m.Ptr, 226, 225, value)
}

// SetUseMasterFirm_TradingIsDisabled
func (m TradeAccountDataResponsePointerBuilder) SetUseMasterFirm_TradingIsDisabled(value uint8) {
	message.SetUint8VLS(m.Ptr, 227, 226, value)
}

// SetIsTradeStatisticsPublicallyShared
func (m TradeAccountDataResponsePointerBuilder) SetIsTradeStatisticsPublicallyShared(value uint8) {
	message.SetUint8VLS(m.Ptr, 228, 227, value)
}

// SetIsReadOnlyFollowingRequestsAllowed
func (m TradeAccountDataResponsePointerBuilder) SetIsReadOnlyFollowingRequestsAllowed(value uint8) {
	message.SetUint8VLS(m.Ptr, 229, 228, value)
}

// SetIsTradeAccountSharingAllowed
func (m TradeAccountDataResponsePointerBuilder) SetIsTradeAccountSharingAllowed(value uint8) {
	message.SetUint8VLS(m.Ptr, 230, 229, value)
}

// SetReadOnlyShareToAllSCUsernames
func (m TradeAccountDataResponsePointerBuilder) SetReadOnlyShareToAllSCUsernames(value uint8) {
	message.SetUint8VLS(m.Ptr, 231, 230, value)
}

// SetUseMasterFirm_SymbolCommissionsArray
func (m TradeAccountDataResponsePointerBuilder) SetUseMasterFirm_SymbolCommissionsArray(value uint8) {
	message.SetUint8VLS(m.Ptr, 232, 231, value)
}

// SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit
func (m TradeAccountDataResponsePointerBuilder) SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit(value uint8) {
	message.SetUint8VLS(m.Ptr, 233, 232, value)
}

// SetUseMasterFirm_UsePercentOfMarginForDayTrading
func (m TradeAccountDataResponsePointerBuilder) SetUseMasterFirm_UsePercentOfMarginForDayTrading(value uint8) {
	message.SetUint8VLS(m.Ptr, 234, 233, value)
}

// SetOpenPositionsProfitLossBasedOnSettlement
func (m TradeAccountDataResponsePointerBuilder) SetOpenPositionsProfitLossBasedOnSettlement(value float64) {
	message.SetFloat64VLS(m.Ptr, 242, 234, value)
}

// SetMarginRequirementFull
func (m TradeAccountDataResponsePointerBuilder) SetMarginRequirementFull(value float64) {
	message.SetFloat64VLS(m.Ptr, 250, 242, value)
}

// SetMarginRequirementFullPositionsOnly
func (m TradeAccountDataResponsePointerBuilder) SetMarginRequirementFullPositionsOnly(value float64) {
	message.SetFloat64VLS(m.Ptr, 258, 250, value)
}

// SetUseMasterFirm_TradeFeesFullOverride
func (m TradeAccountDataResponsePointerBuilder) SetUseMasterFirm_TradeFeesFullOverride(value uint8) {
	message.SetUint8VLS(m.Ptr, 259, 258, value)
}

// SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders
func (m TradeAccountDataResponsePointerBuilder) SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders(value uint8) {
	message.SetUint8VLS(m.Ptr, 260, 259, value)
}

// SetUseMasterFirm_UsePercentOfMarginFullOverride
func (m TradeAccountDataResponsePointerBuilder) SetUseMasterFirm_UsePercentOfMarginFullOverride(value uint8) {
	message.SetUint8VLS(m.Ptr, 261, 260, value)
}

// SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverride
func (m TradeAccountDataResponsePointerBuilder) SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverride(value uint8) {
	message.SetUint8VLS(m.Ptr, 262, 261, value)
}

// SetPeakMarginRequirement
func (m TradeAccountDataResponsePointerBuilder) SetPeakMarginRequirement(value float64) {
	message.SetFloat64VLS(m.Ptr, 270, 262, value)
}

// SetLiquidationOnlyMode
func (m TradeAccountDataResponsePointerBuilder) SetLiquidationOnlyMode(value uint8) {
	message.SetUint8VLS(m.Ptr, 271, 270, value)
}

// SetFlattenPositionsWhenUnderMarginIntradayTriggered
func (m TradeAccountDataResponsePointerBuilder) SetFlattenPositionsWhenUnderMarginIntradayTriggered(value uint8) {
	message.SetUint8VLS(m.Ptr, 272, 271, value)
}

// SetFlattenPositionsWhenUnderMinimumAccountValueTriggered
func (m TradeAccountDataResponsePointerBuilder) SetFlattenPositionsWhenUnderMinimumAccountValueTriggered(value uint8) {
	message.SetUint8VLS(m.Ptr, 273, 272, value)
}

// SetAccountValueAtEndOfDayCaptureTime
func (m TradeAccountDataResponsePointerBuilder) SetAccountValueAtEndOfDayCaptureTime(value float64) {
	message.SetFloat64VLS(m.Ptr, 281, 273, value)
}

// SetEndOfDayCaptureTime
func (m TradeAccountDataResponsePointerBuilder) SetEndOfDayCaptureTime(value int64) {
	message.SetInt64VLS(m.Ptr, 289, 281, value)
}

// SetCustomerOrFirm
func (m TradeAccountDataResponsePointerBuilder) SetCustomerOrFirm(value uint8) {
	message.SetUint8VLS(m.Ptr, 290, 289, value)
}

// SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMet
func (m TradeAccountDataResponsePointerBuilder) SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMet(value uint8) {
	message.SetUint8VLS(m.Ptr, 291, 290, value)
}

// SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValue
func (m TradeAccountDataResponsePointerBuilder) SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValue(value uint8) {
	message.SetUint8VLS(m.Ptr, 292, 291, value)
}

// SetMasterFirm_FlattenCancelWhenUnderMarginIntraday
func (m TradeAccountDataResponsePointerBuilder) SetMasterFirm_FlattenCancelWhenUnderMarginIntraday(value uint8) {
	message.SetUint8VLS(m.Ptr, 293, 292, value)
}

// SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDay
func (m TradeAccountDataResponsePointerBuilder) SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDay(value uint8) {
	message.SetUint8VLS(m.Ptr, 294, 293, value)
}

// SetMasterFirm_MaximumOrderQuantity
func (m TradeAccountDataResponsePointerBuilder) SetMasterFirm_MaximumOrderQuantity(value uint32) {
	message.SetUint32VLS(m.Ptr, 298, 294, value)
}

// SetLastTriggerDateTimeUTCForDailyLossLimit
func (m TradeAccountDataResponsePointerBuilder) SetLastTriggerDateTimeUTCForDailyLossLimit(value int64) {
	message.SetInt64VLS(m.Ptr, 306, 298, value)
}

// SetOpenPositionsProfitLossIsDelayed
func (m TradeAccountDataResponsePointerBuilder) SetOpenPositionsProfitLossIsDelayed(value uint8) {
	message.SetUint8VLS(m.Ptr, 307, 306, value)
}

// SetExchangeTraderId
func (m *TradeAccountDataResponsePointerBuilder) SetExchangeTraderId(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 311, 307, value)
}

// Copy
func (m TradeAccountDataResponsePointer) Copy(to TradeAccountDataResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccountNotExist(m.TradeAccountNotExist())
	to.SetTradeAccount(m.TradeAccount())
	to.SetIsSimulated(m.IsSimulated())
	to.SetCurrencyCode(m.CurrencyCode())
	to.SetCurrentCashBalance(m.CurrentCashBalance())
	to.SetAvailableFundsForNewPositions(m.AvailableFundsForNewPositions())
	to.SetMarginRequirement(m.MarginRequirement())
	to.SetAccountValue(m.AccountValue())
	to.SetOpenPositionsProfitLoss(m.OpenPositionsProfitLoss())
	to.SetDailyProfitLoss(m.DailyProfitLoss())
	to.SetTransactionIdentifierForCashBalanceAdjustment(m.TransactionIdentifierForCashBalanceAdjustment())
	to.SetLastTransactionDateTime(m.LastTransactionDateTime())
	to.SetTrailingAccountValueAtWhichToNotAllowNewPositions(m.TrailingAccountValueAtWhichToNotAllowNewPositions())
	to.SetCalculatedDailyNetLossLimitInAccountCurrency(m.CalculatedDailyNetLossLimitInAccountCurrency())
	to.SetDailyNetLossLimitHasBeenReached(m.DailyNetLossLimitHasBeenReached())
	to.SetLastResetDailyNetLossManagementVariablesDateTimeUTC(m.LastResetDailyNetLossManagementVariablesDateTimeUTC())
	to.SetIsUnderRequiredMargin(m.IsUnderRequiredMargin())
	to.SetDailyNetLossLimitInAccountCurrency(m.DailyNetLossLimitInAccountCurrency())
	to.SetPercentOfCashBalanceForDailyNetLossLimit(m.PercentOfCashBalanceForDailyNetLossLimit())
	to.SetUseTrailingAccountValueToNotAllowIncreaseInPositions(m.UseTrailingAccountValueToNotAllowIncreaseInPositions())
	to.SetDoNotAllowIncreaseInPositionsAtDailyLossLimit(m.DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	to.SetFlattenPositionsAtDailyLossLimit(m.FlattenPositionsAtDailyLossLimit())
	to.SetClosePositionsAtEndOfDay(m.ClosePositionsAtEndOfDay())
	to.SetFlattenPositionsWhenUnderMarginIntraday(m.FlattenPositionsWhenUnderMarginIntraday())
	to.SetFlattenPositionsWhenUnderMarginAtEndOfDay(m.FlattenPositionsWhenUnderMarginAtEndOfDay())
	to.SetSenderSubID(m.SenderSubID())
	to.SetSenderLocationId(m.SenderLocationId())
	to.SetSelfMatchPreventionID(m.SelfMatchPreventionID())
	to.SetCTICode(m.CTICode())
	to.SetTradeAccountIsReadOnly(m.TradeAccountIsReadOnly())
	to.SetMaximumGlobalPositionQuantity(m.MaximumGlobalPositionQuantity())
	to.SetTradeFeePerContract(m.TradeFeePerContract())
	to.SetTradeFeePerShare(m.TradeFeePerShare())
	to.SetRequireSufficientMarginForNewPositions(m.RequireSufficientMarginForNewPositions())
	to.SetUsePercentOfMargin(m.UsePercentOfMargin())
	to.SetUsePercentOfMarginForDayTrading(m.UsePercentOfMarginForDayTrading())
	to.SetMaximumAllowedAccountBalanceForPositionsAsPercentage(m.MaximumAllowedAccountBalanceForPositionsAsPercentage())
	to.SetFirmID(m.FirmID())
	to.SetTradingIsDisabled(m.TradingIsDisabled())
	to.SetDescriptiveName(m.DescriptiveName())
	to.SetIsMasterFirmControlAccount(m.IsMasterFirmControlAccount())
	to.SetMinimumRequiredAccountValue(m.MinimumRequiredAccountValue())
	to.SetBeginTimeForDayMargin(m.BeginTimeForDayMargin())
	to.SetEndTimeForDayMargin(m.EndTimeForDayMargin())
	to.SetDayMarginTimeZone(m.DayMarginTimeZone())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
	to.SetIsDeleted(m.IsDeleted())
	to.SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntraday(m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday())
	to.SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay(m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay())
	to.SetUseMasterFirm_SymbolLimitsArray(m.UseMasterFirm_SymbolLimitsArray())
	to.SetUseMasterFirm_TradeFees(m.UseMasterFirm_TradeFees())
	to.SetUseMasterFirm_TradeFeePerShare(m.UseMasterFirm_TradeFeePerShare())
	to.SetUseMasterFirm_RequireSufficientMarginForNewPositions(m.UseMasterFirm_RequireSufficientMarginForNewPositions())
	to.SetUseMasterFirm_UsePercentOfMargin(m.UseMasterFirm_UsePercentOfMargin())
	to.SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage(m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage())
	to.SetUseMasterFirm_MinimumRequiredAccountValue(m.UseMasterFirm_MinimumRequiredAccountValue())
	to.SetUseMasterFirm_MarginTimeSettings(m.UseMasterFirm_MarginTimeSettings())
	to.SetUseMasterFirm_TradingIsDisabled(m.UseMasterFirm_TradingIsDisabled())
	to.SetIsTradeStatisticsPublicallyShared(m.IsTradeStatisticsPublicallyShared())
	to.SetIsReadOnlyFollowingRequestsAllowed(m.IsReadOnlyFollowingRequestsAllowed())
	to.SetIsTradeAccountSharingAllowed(m.IsTradeAccountSharingAllowed())
	to.SetReadOnlyShareToAllSCUsernames(m.ReadOnlyShareToAllSCUsernames())
	to.SetUseMasterFirm_SymbolCommissionsArray(m.UseMasterFirm_SymbolCommissionsArray())
	to.SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit(m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	to.SetUseMasterFirm_UsePercentOfMarginForDayTrading(m.UseMasterFirm_UsePercentOfMarginForDayTrading())
	to.SetOpenPositionsProfitLossBasedOnSettlement(m.OpenPositionsProfitLossBasedOnSettlement())
	to.SetMarginRequirementFull(m.MarginRequirementFull())
	to.SetMarginRequirementFullPositionsOnly(m.MarginRequirementFullPositionsOnly())
	to.SetUseMasterFirm_TradeFeesFullOverride(m.UseMasterFirm_TradeFeesFullOverride())
	to.SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders(m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders())
	to.SetUseMasterFirm_UsePercentOfMarginFullOverride(m.UseMasterFirm_UsePercentOfMarginFullOverride())
	to.SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverride(m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride())
	to.SetPeakMarginRequirement(m.PeakMarginRequirement())
	to.SetLiquidationOnlyMode(m.LiquidationOnlyMode())
	to.SetFlattenPositionsWhenUnderMarginIntradayTriggered(m.FlattenPositionsWhenUnderMarginIntradayTriggered())
	to.SetFlattenPositionsWhenUnderMinimumAccountValueTriggered(m.FlattenPositionsWhenUnderMinimumAccountValueTriggered())
	to.SetAccountValueAtEndOfDayCaptureTime(m.AccountValueAtEndOfDayCaptureTime())
	to.SetEndOfDayCaptureTime(m.EndOfDayCaptureTime())
	to.SetCustomerOrFirm(m.CustomerOrFirm())
	to.SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMet(m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet())
	to.SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValue(m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValue())
	to.SetMasterFirm_FlattenCancelWhenUnderMarginIntraday(m.MasterFirm_FlattenCancelWhenUnderMarginIntraday())
	to.SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDay(m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDay())
	to.SetMasterFirm_MaximumOrderQuantity(m.MasterFirm_MaximumOrderQuantity())
	to.SetLastTriggerDateTimeUTCForDailyLossLimit(m.LastTriggerDateTimeUTCForDailyLossLimit())
	to.SetOpenPositionsProfitLossIsDelayed(m.OpenPositionsProfitLossIsDelayed())
	to.SetExchangeTraderId(m.ExchangeTraderId())
}

// Copy
func (m TradeAccountDataResponsePointerBuilder) Copy(to TradeAccountDataResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccountNotExist(m.TradeAccountNotExist())
	to.SetTradeAccount(m.TradeAccount())
	to.SetIsSimulated(m.IsSimulated())
	to.SetCurrencyCode(m.CurrencyCode())
	to.SetCurrentCashBalance(m.CurrentCashBalance())
	to.SetAvailableFundsForNewPositions(m.AvailableFundsForNewPositions())
	to.SetMarginRequirement(m.MarginRequirement())
	to.SetAccountValue(m.AccountValue())
	to.SetOpenPositionsProfitLoss(m.OpenPositionsProfitLoss())
	to.SetDailyProfitLoss(m.DailyProfitLoss())
	to.SetTransactionIdentifierForCashBalanceAdjustment(m.TransactionIdentifierForCashBalanceAdjustment())
	to.SetLastTransactionDateTime(m.LastTransactionDateTime())
	to.SetTrailingAccountValueAtWhichToNotAllowNewPositions(m.TrailingAccountValueAtWhichToNotAllowNewPositions())
	to.SetCalculatedDailyNetLossLimitInAccountCurrency(m.CalculatedDailyNetLossLimitInAccountCurrency())
	to.SetDailyNetLossLimitHasBeenReached(m.DailyNetLossLimitHasBeenReached())
	to.SetLastResetDailyNetLossManagementVariablesDateTimeUTC(m.LastResetDailyNetLossManagementVariablesDateTimeUTC())
	to.SetIsUnderRequiredMargin(m.IsUnderRequiredMargin())
	to.SetDailyNetLossLimitInAccountCurrency(m.DailyNetLossLimitInAccountCurrency())
	to.SetPercentOfCashBalanceForDailyNetLossLimit(m.PercentOfCashBalanceForDailyNetLossLimit())
	to.SetUseTrailingAccountValueToNotAllowIncreaseInPositions(m.UseTrailingAccountValueToNotAllowIncreaseInPositions())
	to.SetDoNotAllowIncreaseInPositionsAtDailyLossLimit(m.DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	to.SetFlattenPositionsAtDailyLossLimit(m.FlattenPositionsAtDailyLossLimit())
	to.SetClosePositionsAtEndOfDay(m.ClosePositionsAtEndOfDay())
	to.SetFlattenPositionsWhenUnderMarginIntraday(m.FlattenPositionsWhenUnderMarginIntraday())
	to.SetFlattenPositionsWhenUnderMarginAtEndOfDay(m.FlattenPositionsWhenUnderMarginAtEndOfDay())
	to.SetSenderSubID(m.SenderSubID())
	to.SetSenderLocationId(m.SenderLocationId())
	to.SetSelfMatchPreventionID(m.SelfMatchPreventionID())
	to.SetCTICode(m.CTICode())
	to.SetTradeAccountIsReadOnly(m.TradeAccountIsReadOnly())
	to.SetMaximumGlobalPositionQuantity(m.MaximumGlobalPositionQuantity())
	to.SetTradeFeePerContract(m.TradeFeePerContract())
	to.SetTradeFeePerShare(m.TradeFeePerShare())
	to.SetRequireSufficientMarginForNewPositions(m.RequireSufficientMarginForNewPositions())
	to.SetUsePercentOfMargin(m.UsePercentOfMargin())
	to.SetUsePercentOfMarginForDayTrading(m.UsePercentOfMarginForDayTrading())
	to.SetMaximumAllowedAccountBalanceForPositionsAsPercentage(m.MaximumAllowedAccountBalanceForPositionsAsPercentage())
	to.SetFirmID(m.FirmID())
	to.SetTradingIsDisabled(m.TradingIsDisabled())
	to.SetDescriptiveName(m.DescriptiveName())
	to.SetIsMasterFirmControlAccount(m.IsMasterFirmControlAccount())
	to.SetMinimumRequiredAccountValue(m.MinimumRequiredAccountValue())
	to.SetBeginTimeForDayMargin(m.BeginTimeForDayMargin())
	to.SetEndTimeForDayMargin(m.EndTimeForDayMargin())
	to.SetDayMarginTimeZone(m.DayMarginTimeZone())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
	to.SetIsDeleted(m.IsDeleted())
	to.SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntraday(m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday())
	to.SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay(m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay())
	to.SetUseMasterFirm_SymbolLimitsArray(m.UseMasterFirm_SymbolLimitsArray())
	to.SetUseMasterFirm_TradeFees(m.UseMasterFirm_TradeFees())
	to.SetUseMasterFirm_TradeFeePerShare(m.UseMasterFirm_TradeFeePerShare())
	to.SetUseMasterFirm_RequireSufficientMarginForNewPositions(m.UseMasterFirm_RequireSufficientMarginForNewPositions())
	to.SetUseMasterFirm_UsePercentOfMargin(m.UseMasterFirm_UsePercentOfMargin())
	to.SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage(m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage())
	to.SetUseMasterFirm_MinimumRequiredAccountValue(m.UseMasterFirm_MinimumRequiredAccountValue())
	to.SetUseMasterFirm_MarginTimeSettings(m.UseMasterFirm_MarginTimeSettings())
	to.SetUseMasterFirm_TradingIsDisabled(m.UseMasterFirm_TradingIsDisabled())
	to.SetIsTradeStatisticsPublicallyShared(m.IsTradeStatisticsPublicallyShared())
	to.SetIsReadOnlyFollowingRequestsAllowed(m.IsReadOnlyFollowingRequestsAllowed())
	to.SetIsTradeAccountSharingAllowed(m.IsTradeAccountSharingAllowed())
	to.SetReadOnlyShareToAllSCUsernames(m.ReadOnlyShareToAllSCUsernames())
	to.SetUseMasterFirm_SymbolCommissionsArray(m.UseMasterFirm_SymbolCommissionsArray())
	to.SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit(m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	to.SetUseMasterFirm_UsePercentOfMarginForDayTrading(m.UseMasterFirm_UsePercentOfMarginForDayTrading())
	to.SetOpenPositionsProfitLossBasedOnSettlement(m.OpenPositionsProfitLossBasedOnSettlement())
	to.SetMarginRequirementFull(m.MarginRequirementFull())
	to.SetMarginRequirementFullPositionsOnly(m.MarginRequirementFullPositionsOnly())
	to.SetUseMasterFirm_TradeFeesFullOverride(m.UseMasterFirm_TradeFeesFullOverride())
	to.SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders(m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders())
	to.SetUseMasterFirm_UsePercentOfMarginFullOverride(m.UseMasterFirm_UsePercentOfMarginFullOverride())
	to.SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverride(m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride())
	to.SetPeakMarginRequirement(m.PeakMarginRequirement())
	to.SetLiquidationOnlyMode(m.LiquidationOnlyMode())
	to.SetFlattenPositionsWhenUnderMarginIntradayTriggered(m.FlattenPositionsWhenUnderMarginIntradayTriggered())
	to.SetFlattenPositionsWhenUnderMinimumAccountValueTriggered(m.FlattenPositionsWhenUnderMinimumAccountValueTriggered())
	to.SetAccountValueAtEndOfDayCaptureTime(m.AccountValueAtEndOfDayCaptureTime())
	to.SetEndOfDayCaptureTime(m.EndOfDayCaptureTime())
	to.SetCustomerOrFirm(m.CustomerOrFirm())
	to.SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMet(m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet())
	to.SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValue(m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValue())
	to.SetMasterFirm_FlattenCancelWhenUnderMarginIntraday(m.MasterFirm_FlattenCancelWhenUnderMarginIntraday())
	to.SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDay(m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDay())
	to.SetMasterFirm_MaximumOrderQuantity(m.MasterFirm_MaximumOrderQuantity())
	to.SetLastTriggerDateTimeUTCForDailyLossLimit(m.LastTriggerDateTimeUTCForDailyLossLimit())
	to.SetOpenPositionsProfitLossIsDelayed(m.OpenPositionsProfitLossIsDelayed())
	to.SetExchangeTraderId(m.ExchangeTraderId())
}

// CopyPointer
func (m TradeAccountDataResponsePointer) CopyPointer(to TradeAccountDataResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccountNotExist(m.TradeAccountNotExist())
	to.SetTradeAccount(m.TradeAccount())
	to.SetIsSimulated(m.IsSimulated())
	to.SetCurrencyCode(m.CurrencyCode())
	to.SetCurrentCashBalance(m.CurrentCashBalance())
	to.SetAvailableFundsForNewPositions(m.AvailableFundsForNewPositions())
	to.SetMarginRequirement(m.MarginRequirement())
	to.SetAccountValue(m.AccountValue())
	to.SetOpenPositionsProfitLoss(m.OpenPositionsProfitLoss())
	to.SetDailyProfitLoss(m.DailyProfitLoss())
	to.SetTransactionIdentifierForCashBalanceAdjustment(m.TransactionIdentifierForCashBalanceAdjustment())
	to.SetLastTransactionDateTime(m.LastTransactionDateTime())
	to.SetTrailingAccountValueAtWhichToNotAllowNewPositions(m.TrailingAccountValueAtWhichToNotAllowNewPositions())
	to.SetCalculatedDailyNetLossLimitInAccountCurrency(m.CalculatedDailyNetLossLimitInAccountCurrency())
	to.SetDailyNetLossLimitHasBeenReached(m.DailyNetLossLimitHasBeenReached())
	to.SetLastResetDailyNetLossManagementVariablesDateTimeUTC(m.LastResetDailyNetLossManagementVariablesDateTimeUTC())
	to.SetIsUnderRequiredMargin(m.IsUnderRequiredMargin())
	to.SetDailyNetLossLimitInAccountCurrency(m.DailyNetLossLimitInAccountCurrency())
	to.SetPercentOfCashBalanceForDailyNetLossLimit(m.PercentOfCashBalanceForDailyNetLossLimit())
	to.SetUseTrailingAccountValueToNotAllowIncreaseInPositions(m.UseTrailingAccountValueToNotAllowIncreaseInPositions())
	to.SetDoNotAllowIncreaseInPositionsAtDailyLossLimit(m.DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	to.SetFlattenPositionsAtDailyLossLimit(m.FlattenPositionsAtDailyLossLimit())
	to.SetClosePositionsAtEndOfDay(m.ClosePositionsAtEndOfDay())
	to.SetFlattenPositionsWhenUnderMarginIntraday(m.FlattenPositionsWhenUnderMarginIntraday())
	to.SetFlattenPositionsWhenUnderMarginAtEndOfDay(m.FlattenPositionsWhenUnderMarginAtEndOfDay())
	to.SetSenderSubID(m.SenderSubID())
	to.SetSenderLocationId(m.SenderLocationId())
	to.SetSelfMatchPreventionID(m.SelfMatchPreventionID())
	to.SetCTICode(m.CTICode())
	to.SetTradeAccountIsReadOnly(m.TradeAccountIsReadOnly())
	to.SetMaximumGlobalPositionQuantity(m.MaximumGlobalPositionQuantity())
	to.SetTradeFeePerContract(m.TradeFeePerContract())
	to.SetTradeFeePerShare(m.TradeFeePerShare())
	to.SetRequireSufficientMarginForNewPositions(m.RequireSufficientMarginForNewPositions())
	to.SetUsePercentOfMargin(m.UsePercentOfMargin())
	to.SetUsePercentOfMarginForDayTrading(m.UsePercentOfMarginForDayTrading())
	to.SetMaximumAllowedAccountBalanceForPositionsAsPercentage(m.MaximumAllowedAccountBalanceForPositionsAsPercentage())
	to.SetFirmID(m.FirmID())
	to.SetTradingIsDisabled(m.TradingIsDisabled())
	to.SetDescriptiveName(m.DescriptiveName())
	to.SetIsMasterFirmControlAccount(m.IsMasterFirmControlAccount())
	to.SetMinimumRequiredAccountValue(m.MinimumRequiredAccountValue())
	to.SetBeginTimeForDayMargin(m.BeginTimeForDayMargin())
	to.SetEndTimeForDayMargin(m.EndTimeForDayMargin())
	to.SetDayMarginTimeZone(m.DayMarginTimeZone())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
	to.SetIsDeleted(m.IsDeleted())
	to.SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntraday(m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday())
	to.SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay(m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay())
	to.SetUseMasterFirm_SymbolLimitsArray(m.UseMasterFirm_SymbolLimitsArray())
	to.SetUseMasterFirm_TradeFees(m.UseMasterFirm_TradeFees())
	to.SetUseMasterFirm_TradeFeePerShare(m.UseMasterFirm_TradeFeePerShare())
	to.SetUseMasterFirm_RequireSufficientMarginForNewPositions(m.UseMasterFirm_RequireSufficientMarginForNewPositions())
	to.SetUseMasterFirm_UsePercentOfMargin(m.UseMasterFirm_UsePercentOfMargin())
	to.SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage(m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage())
	to.SetUseMasterFirm_MinimumRequiredAccountValue(m.UseMasterFirm_MinimumRequiredAccountValue())
	to.SetUseMasterFirm_MarginTimeSettings(m.UseMasterFirm_MarginTimeSettings())
	to.SetUseMasterFirm_TradingIsDisabled(m.UseMasterFirm_TradingIsDisabled())
	to.SetIsTradeStatisticsPublicallyShared(m.IsTradeStatisticsPublicallyShared())
	to.SetIsReadOnlyFollowingRequestsAllowed(m.IsReadOnlyFollowingRequestsAllowed())
	to.SetIsTradeAccountSharingAllowed(m.IsTradeAccountSharingAllowed())
	to.SetReadOnlyShareToAllSCUsernames(m.ReadOnlyShareToAllSCUsernames())
	to.SetUseMasterFirm_SymbolCommissionsArray(m.UseMasterFirm_SymbolCommissionsArray())
	to.SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit(m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	to.SetUseMasterFirm_UsePercentOfMarginForDayTrading(m.UseMasterFirm_UsePercentOfMarginForDayTrading())
	to.SetOpenPositionsProfitLossBasedOnSettlement(m.OpenPositionsProfitLossBasedOnSettlement())
	to.SetMarginRequirementFull(m.MarginRequirementFull())
	to.SetMarginRequirementFullPositionsOnly(m.MarginRequirementFullPositionsOnly())
	to.SetUseMasterFirm_TradeFeesFullOverride(m.UseMasterFirm_TradeFeesFullOverride())
	to.SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders(m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders())
	to.SetUseMasterFirm_UsePercentOfMarginFullOverride(m.UseMasterFirm_UsePercentOfMarginFullOverride())
	to.SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverride(m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride())
	to.SetPeakMarginRequirement(m.PeakMarginRequirement())
	to.SetLiquidationOnlyMode(m.LiquidationOnlyMode())
	to.SetFlattenPositionsWhenUnderMarginIntradayTriggered(m.FlattenPositionsWhenUnderMarginIntradayTriggered())
	to.SetFlattenPositionsWhenUnderMinimumAccountValueTriggered(m.FlattenPositionsWhenUnderMinimumAccountValueTriggered())
	to.SetAccountValueAtEndOfDayCaptureTime(m.AccountValueAtEndOfDayCaptureTime())
	to.SetEndOfDayCaptureTime(m.EndOfDayCaptureTime())
	to.SetCustomerOrFirm(m.CustomerOrFirm())
	to.SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMet(m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet())
	to.SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValue(m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValue())
	to.SetMasterFirm_FlattenCancelWhenUnderMarginIntraday(m.MasterFirm_FlattenCancelWhenUnderMarginIntraday())
	to.SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDay(m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDay())
	to.SetMasterFirm_MaximumOrderQuantity(m.MasterFirm_MaximumOrderQuantity())
	to.SetLastTriggerDateTimeUTCForDailyLossLimit(m.LastTriggerDateTimeUTCForDailyLossLimit())
	to.SetOpenPositionsProfitLossIsDelayed(m.OpenPositionsProfitLossIsDelayed())
	to.SetExchangeTraderId(m.ExchangeTraderId())
}

// CopyPointer
func (m TradeAccountDataResponsePointerBuilder) CopyPointer(to TradeAccountDataResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccountNotExist(m.TradeAccountNotExist())
	to.SetTradeAccount(m.TradeAccount())
	to.SetIsSimulated(m.IsSimulated())
	to.SetCurrencyCode(m.CurrencyCode())
	to.SetCurrentCashBalance(m.CurrentCashBalance())
	to.SetAvailableFundsForNewPositions(m.AvailableFundsForNewPositions())
	to.SetMarginRequirement(m.MarginRequirement())
	to.SetAccountValue(m.AccountValue())
	to.SetOpenPositionsProfitLoss(m.OpenPositionsProfitLoss())
	to.SetDailyProfitLoss(m.DailyProfitLoss())
	to.SetTransactionIdentifierForCashBalanceAdjustment(m.TransactionIdentifierForCashBalanceAdjustment())
	to.SetLastTransactionDateTime(m.LastTransactionDateTime())
	to.SetTrailingAccountValueAtWhichToNotAllowNewPositions(m.TrailingAccountValueAtWhichToNotAllowNewPositions())
	to.SetCalculatedDailyNetLossLimitInAccountCurrency(m.CalculatedDailyNetLossLimitInAccountCurrency())
	to.SetDailyNetLossLimitHasBeenReached(m.DailyNetLossLimitHasBeenReached())
	to.SetLastResetDailyNetLossManagementVariablesDateTimeUTC(m.LastResetDailyNetLossManagementVariablesDateTimeUTC())
	to.SetIsUnderRequiredMargin(m.IsUnderRequiredMargin())
	to.SetDailyNetLossLimitInAccountCurrency(m.DailyNetLossLimitInAccountCurrency())
	to.SetPercentOfCashBalanceForDailyNetLossLimit(m.PercentOfCashBalanceForDailyNetLossLimit())
	to.SetUseTrailingAccountValueToNotAllowIncreaseInPositions(m.UseTrailingAccountValueToNotAllowIncreaseInPositions())
	to.SetDoNotAllowIncreaseInPositionsAtDailyLossLimit(m.DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	to.SetFlattenPositionsAtDailyLossLimit(m.FlattenPositionsAtDailyLossLimit())
	to.SetClosePositionsAtEndOfDay(m.ClosePositionsAtEndOfDay())
	to.SetFlattenPositionsWhenUnderMarginIntraday(m.FlattenPositionsWhenUnderMarginIntraday())
	to.SetFlattenPositionsWhenUnderMarginAtEndOfDay(m.FlattenPositionsWhenUnderMarginAtEndOfDay())
	to.SetSenderSubID(m.SenderSubID())
	to.SetSenderLocationId(m.SenderLocationId())
	to.SetSelfMatchPreventionID(m.SelfMatchPreventionID())
	to.SetCTICode(m.CTICode())
	to.SetTradeAccountIsReadOnly(m.TradeAccountIsReadOnly())
	to.SetMaximumGlobalPositionQuantity(m.MaximumGlobalPositionQuantity())
	to.SetTradeFeePerContract(m.TradeFeePerContract())
	to.SetTradeFeePerShare(m.TradeFeePerShare())
	to.SetRequireSufficientMarginForNewPositions(m.RequireSufficientMarginForNewPositions())
	to.SetUsePercentOfMargin(m.UsePercentOfMargin())
	to.SetUsePercentOfMarginForDayTrading(m.UsePercentOfMarginForDayTrading())
	to.SetMaximumAllowedAccountBalanceForPositionsAsPercentage(m.MaximumAllowedAccountBalanceForPositionsAsPercentage())
	to.SetFirmID(m.FirmID())
	to.SetTradingIsDisabled(m.TradingIsDisabled())
	to.SetDescriptiveName(m.DescriptiveName())
	to.SetIsMasterFirmControlAccount(m.IsMasterFirmControlAccount())
	to.SetMinimumRequiredAccountValue(m.MinimumRequiredAccountValue())
	to.SetBeginTimeForDayMargin(m.BeginTimeForDayMargin())
	to.SetEndTimeForDayMargin(m.EndTimeForDayMargin())
	to.SetDayMarginTimeZone(m.DayMarginTimeZone())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
	to.SetIsDeleted(m.IsDeleted())
	to.SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntraday(m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday())
	to.SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay(m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay())
	to.SetUseMasterFirm_SymbolLimitsArray(m.UseMasterFirm_SymbolLimitsArray())
	to.SetUseMasterFirm_TradeFees(m.UseMasterFirm_TradeFees())
	to.SetUseMasterFirm_TradeFeePerShare(m.UseMasterFirm_TradeFeePerShare())
	to.SetUseMasterFirm_RequireSufficientMarginForNewPositions(m.UseMasterFirm_RequireSufficientMarginForNewPositions())
	to.SetUseMasterFirm_UsePercentOfMargin(m.UseMasterFirm_UsePercentOfMargin())
	to.SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage(m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage())
	to.SetUseMasterFirm_MinimumRequiredAccountValue(m.UseMasterFirm_MinimumRequiredAccountValue())
	to.SetUseMasterFirm_MarginTimeSettings(m.UseMasterFirm_MarginTimeSettings())
	to.SetUseMasterFirm_TradingIsDisabled(m.UseMasterFirm_TradingIsDisabled())
	to.SetIsTradeStatisticsPublicallyShared(m.IsTradeStatisticsPublicallyShared())
	to.SetIsReadOnlyFollowingRequestsAllowed(m.IsReadOnlyFollowingRequestsAllowed())
	to.SetIsTradeAccountSharingAllowed(m.IsTradeAccountSharingAllowed())
	to.SetReadOnlyShareToAllSCUsernames(m.ReadOnlyShareToAllSCUsernames())
	to.SetUseMasterFirm_SymbolCommissionsArray(m.UseMasterFirm_SymbolCommissionsArray())
	to.SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit(m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	to.SetUseMasterFirm_UsePercentOfMarginForDayTrading(m.UseMasterFirm_UsePercentOfMarginForDayTrading())
	to.SetOpenPositionsProfitLossBasedOnSettlement(m.OpenPositionsProfitLossBasedOnSettlement())
	to.SetMarginRequirementFull(m.MarginRequirementFull())
	to.SetMarginRequirementFullPositionsOnly(m.MarginRequirementFullPositionsOnly())
	to.SetUseMasterFirm_TradeFeesFullOverride(m.UseMasterFirm_TradeFeesFullOverride())
	to.SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders(m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders())
	to.SetUseMasterFirm_UsePercentOfMarginFullOverride(m.UseMasterFirm_UsePercentOfMarginFullOverride())
	to.SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverride(m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride())
	to.SetPeakMarginRequirement(m.PeakMarginRequirement())
	to.SetLiquidationOnlyMode(m.LiquidationOnlyMode())
	to.SetFlattenPositionsWhenUnderMarginIntradayTriggered(m.FlattenPositionsWhenUnderMarginIntradayTriggered())
	to.SetFlattenPositionsWhenUnderMinimumAccountValueTriggered(m.FlattenPositionsWhenUnderMinimumAccountValueTriggered())
	to.SetAccountValueAtEndOfDayCaptureTime(m.AccountValueAtEndOfDayCaptureTime())
	to.SetEndOfDayCaptureTime(m.EndOfDayCaptureTime())
	to.SetCustomerOrFirm(m.CustomerOrFirm())
	to.SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMet(m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet())
	to.SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValue(m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValue())
	to.SetMasterFirm_FlattenCancelWhenUnderMarginIntraday(m.MasterFirm_FlattenCancelWhenUnderMarginIntraday())
	to.SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDay(m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDay())
	to.SetMasterFirm_MaximumOrderQuantity(m.MasterFirm_MaximumOrderQuantity())
	to.SetLastTriggerDateTimeUTCForDailyLossLimit(m.LastTriggerDateTimeUTCForDailyLossLimit())
	to.SetOpenPositionsProfitLossIsDelayed(m.OpenPositionsProfitLossIsDelayed())
	to.SetExchangeTraderId(m.ExchangeTraderId())
}
