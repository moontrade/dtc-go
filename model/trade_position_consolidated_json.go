//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradePositionConsolidated) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10113,\"F\":["...)
	w.Uint8(m.IsDeleted())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsSimulated())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.String(m.CurrencyCode())
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.AveragePrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.OpenProfitLoss())
	w.Data = append(w.Data, ',')
	w.Float64(m.DailyProfitLoss())
	w.Data = append(w.Data, ',')
	w.Int64(m.LastDailyProfitLossResetDateTimeUTC())
	w.Data = append(w.Data, ',')
	w.Float64(m.ServicePositionQuantity())
	w.Data = append(w.Data, ',')
	w.Uint8(m.PositionHasBeenUpdatedByService())
	w.Data = append(w.Data, ',')
	w.Float64(m.PriceHighDuringPosition())
	w.Data = append(w.Data, ',')
	w.Float64(m.PriceLowDuringPosition())
	w.Data = append(w.Data, ',')
	w.Float64(m.PriceLastDuringPosition())
	w.Data = append(w.Data, ',')
	w.Int64(m.LastProcessedTimeAndSalesSequence())
	w.Data = append(w.Data, ',')
	w.Float64(m.TotalMarginRequirement())
	w.Data = append(w.Data, ',')
	w.Int64(m.InitialEntryDateTimeUTC())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsFromDTCServerReplay())
	w.Data = append(w.Data, ',')
	w.Int64(m.MostRecentPositionIncreaseDateTimeUTC())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsSnapshot())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsFirstMessageInBatch())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsLastMessageInBatch())
	w.Data = append(w.Data, ',')
	w.Float64(m.MarginRequirementFull())
	w.Data = append(w.Data, ',')
	w.Float64(m.MarginRequirementFullPositionsOnly())
	w.Data = append(w.Data, ',')
	w.Float64(m.MaxPotentialPositionQuantity())
	return w.Finish(), nil
}

func (m TradePositionConsolidatedBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10113,\"F\":["...)
	w.Uint8(m.IsDeleted())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsSimulated())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.String(m.CurrencyCode())
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.AveragePrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.OpenProfitLoss())
	w.Data = append(w.Data, ',')
	w.Float64(m.DailyProfitLoss())
	w.Data = append(w.Data, ',')
	w.Int64(m.LastDailyProfitLossResetDateTimeUTC())
	w.Data = append(w.Data, ',')
	w.Float64(m.ServicePositionQuantity())
	w.Data = append(w.Data, ',')
	w.Uint8(m.PositionHasBeenUpdatedByService())
	w.Data = append(w.Data, ',')
	w.Float64(m.PriceHighDuringPosition())
	w.Data = append(w.Data, ',')
	w.Float64(m.PriceLowDuringPosition())
	w.Data = append(w.Data, ',')
	w.Float64(m.PriceLastDuringPosition())
	w.Data = append(w.Data, ',')
	w.Int64(m.LastProcessedTimeAndSalesSequence())
	w.Data = append(w.Data, ',')
	w.Float64(m.TotalMarginRequirement())
	w.Data = append(w.Data, ',')
	w.Int64(m.InitialEntryDateTimeUTC())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsFromDTCServerReplay())
	w.Data = append(w.Data, ',')
	w.Int64(m.MostRecentPositionIncreaseDateTimeUTC())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsSnapshot())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsFirstMessageInBatch())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsLastMessageInBatch())
	w.Data = append(w.Data, ',')
	w.Float64(m.MarginRequirementFull())
	w.Data = append(w.Data, ',')
	w.Float64(m.MarginRequirementFullPositionsOnly())
	w.Data = append(w.Data, ',')
	w.Float64(m.MaxPotentialPositionQuantity())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradePositionConsolidated) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10113)
	w.Uint8Field("m_IsDeleted", m.IsDeleted())
	w.StringField("m_Symbol", m.Symbol())
	w.Uint8Field("m_IsSimulated", m.IsSimulated())
	w.StringField("m_TradeAccount", m.TradeAccount())
	w.StringField("m_CurrencyCode", m.CurrencyCode())
	w.Float64Field("m_Quantity", m.Quantity())
	w.Float64Field("m_AveragePrice", m.AveragePrice())
	w.Float64Field("m_OpenProfitLoss", m.OpenProfitLoss())
	w.Float64Field("m_DailyProfitLoss", m.DailyProfitLoss())
	w.Int64Field("m_LastDailyProfitLossResetDateTimeUTC", m.LastDailyProfitLossResetDateTimeUTC())
	w.Float64Field("m_ServicePositionQuantity", m.ServicePositionQuantity())
	w.Uint8Field("m_PositionHasBeenUpdatedByService", m.PositionHasBeenUpdatedByService())
	w.Float64Field("m_PriceHighDuringPosition", m.PriceHighDuringPosition())
	w.Float64Field("m_PriceLowDuringPosition", m.PriceLowDuringPosition())
	w.Float64Field("m_PriceLastDuringPosition", m.PriceLastDuringPosition())
	w.Int64Field("m_LastProcessedTimeAndSalesSequence", m.LastProcessedTimeAndSalesSequence())
	w.Float64Field("m_TotalMarginRequirement", m.TotalMarginRequirement())
	w.Int64Field("m_InitialEntryDateTimeUTC", m.InitialEntryDateTimeUTC())
	w.Uint8Field("m_IsFromDTCServerReplay", m.IsFromDTCServerReplay())
	w.Int64Field("m_MostRecentPositionIncreaseDateTimeUTC", m.MostRecentPositionIncreaseDateTimeUTC())
	w.Uint8Field("m_IsSnapshot", m.IsSnapshot())
	w.Uint8Field("m_IsFirstMessageInBatch", m.IsFirstMessageInBatch())
	w.Uint8Field("m_IsLastMessageInBatch", m.IsLastMessageInBatch())
	w.Float64Field("m_MarginRequirementFull", m.MarginRequirementFull())
	w.Float64Field("m_MarginRequirementFullPositionsOnly", m.MarginRequirementFullPositionsOnly())
	w.Float64Field("m_MaxPotentialPositionQuantity", m.MaxPotentialPositionQuantity())
	return w.Finish(), nil
}

func (m TradePositionConsolidatedBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10113)
	w.Uint8Field("m_IsDeleted", m.IsDeleted())
	w.StringField("m_Symbol", m.Symbol())
	w.Uint8Field("m_IsSimulated", m.IsSimulated())
	w.StringField("m_TradeAccount", m.TradeAccount())
	w.StringField("m_CurrencyCode", m.CurrencyCode())
	w.Float64Field("m_Quantity", m.Quantity())
	w.Float64Field("m_AveragePrice", m.AveragePrice())
	w.Float64Field("m_OpenProfitLoss", m.OpenProfitLoss())
	w.Float64Field("m_DailyProfitLoss", m.DailyProfitLoss())
	w.Int64Field("m_LastDailyProfitLossResetDateTimeUTC", m.LastDailyProfitLossResetDateTimeUTC())
	w.Float64Field("m_ServicePositionQuantity", m.ServicePositionQuantity())
	w.Uint8Field("m_PositionHasBeenUpdatedByService", m.PositionHasBeenUpdatedByService())
	w.Float64Field("m_PriceHighDuringPosition", m.PriceHighDuringPosition())
	w.Float64Field("m_PriceLowDuringPosition", m.PriceLowDuringPosition())
	w.Float64Field("m_PriceLastDuringPosition", m.PriceLastDuringPosition())
	w.Int64Field("m_LastProcessedTimeAndSalesSequence", m.LastProcessedTimeAndSalesSequence())
	w.Float64Field("m_TotalMarginRequirement", m.TotalMarginRequirement())
	w.Int64Field("m_InitialEntryDateTimeUTC", m.InitialEntryDateTimeUTC())
	w.Uint8Field("m_IsFromDTCServerReplay", m.IsFromDTCServerReplay())
	w.Int64Field("m_MostRecentPositionIncreaseDateTimeUTC", m.MostRecentPositionIncreaseDateTimeUTC())
	w.Uint8Field("m_IsSnapshot", m.IsSnapshot())
	w.Uint8Field("m_IsFirstMessageInBatch", m.IsFirstMessageInBatch())
	w.Uint8Field("m_IsLastMessageInBatch", m.IsLastMessageInBatch())
	w.Float64Field("m_MarginRequirementFull", m.MarginRequirementFull())
	w.Float64Field("m_MarginRequirementFullPositionsOnly", m.MarginRequirementFullPositionsOnly())
	w.Float64Field("m_MaxPotentialPositionQuantity", m.MaxPotentialPositionQuantity())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradePositionConsolidatedBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10113 {
		return message.ErrWrongType
	}
	m.SetIsDeleted(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetSymbol(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsSimulated(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetCurrencyCode(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetAveragePrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetOpenProfitLoss(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetDailyProfitLoss(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetLastDailyProfitLossResetDateTimeUTC(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	m.SetServicePositionQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetPositionHasBeenUpdatedByService(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetPriceHighDuringPosition(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetPriceLowDuringPosition(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetPriceLastDuringPosition(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetLastProcessedTimeAndSalesSequence(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	m.SetTotalMarginRequirement(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetInitialEntryDateTimeUTC(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsFromDTCServerReplay(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetMostRecentPositionIncreaseDateTimeUTC(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsSnapshot(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsFirstMessageInBatch(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsLastMessageInBatch(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetMarginRequirementFull(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetMarginRequirementFullPositionsOnly(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetMaxPotentialPositionQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradePositionConsolidatedBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 10113 {
		return message.ErrWrongType
	}
	in := &r.Lexer
LOOP:
	for !in.IsDelim('}') {
		key, err := r.FieldName()
		if err != nil {
			return err
		}
		switch key {
		case "m_IsDeleted":
			m.SetIsDeleted(r.Uint8())
		case "m_Symbol":
			m.SetSymbol(r.String())
		case "m_IsSimulated":
			m.SetIsSimulated(r.Uint8())
		case "m_TradeAccount":
			m.SetTradeAccount(r.String())
		case "m_CurrencyCode":
			m.SetCurrencyCode(r.String())
		case "m_Quantity":
			m.SetQuantity(r.Float64())
		case "m_AveragePrice":
			m.SetAveragePrice(r.Float64())
		case "m_OpenProfitLoss":
			m.SetOpenProfitLoss(r.Float64())
		case "m_DailyProfitLoss":
			m.SetDailyProfitLoss(r.Float64())
		case "m_LastDailyProfitLossResetDateTimeUTC":
			m.SetLastDailyProfitLossResetDateTimeUTC(r.Int64())
		case "m_ServicePositionQuantity":
			m.SetServicePositionQuantity(r.Float64())
		case "m_PositionHasBeenUpdatedByService":
			m.SetPositionHasBeenUpdatedByService(r.Uint8())
		case "m_PriceHighDuringPosition":
			m.SetPriceHighDuringPosition(r.Float64())
		case "m_PriceLowDuringPosition":
			m.SetPriceLowDuringPosition(r.Float64())
		case "m_PriceLastDuringPosition":
			m.SetPriceLastDuringPosition(r.Float64())
		case "m_LastProcessedTimeAndSalesSequence":
			m.SetLastProcessedTimeAndSalesSequence(r.Int64())
		case "m_TotalMarginRequirement":
			m.SetTotalMarginRequirement(r.Float64())
		case "m_InitialEntryDateTimeUTC":
			m.SetInitialEntryDateTimeUTC(r.Int64())
		case "m_IsFromDTCServerReplay":
			m.SetIsFromDTCServerReplay(r.Uint8())
		case "m_MostRecentPositionIncreaseDateTimeUTC":
			m.SetMostRecentPositionIncreaseDateTimeUTC(r.Int64())
		case "m_IsSnapshot":
			m.SetIsSnapshot(r.Uint8())
		case "m_IsFirstMessageInBatch":
			m.SetIsFirstMessageInBatch(r.Uint8())
		case "m_IsLastMessageInBatch":
			m.SetIsLastMessageInBatch(r.Uint8())
		case "m_MarginRequirementFull":
			m.SetMarginRequirementFull(r.Float64())
		case "m_MarginRequirementFullPositionsOnly":
			m.SetMarginRequirementFullPositionsOnly(r.Float64())
		case "m_MaxPotentialPositionQuantity":
			m.SetMaxPotentialPositionQuantity(r.Float64())
		case "f", "F":
			return message.ErrJSONCompactDetected
		case "":
			break LOOP
		default:
			in.SkipRecursive()
		}
		if r.IsError() {
			return r.Error()
		}
		in.WantComma()
	}
	return nil
}
