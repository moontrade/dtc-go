//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *AddCorrectingOrderFillBuilder) UnmarshalProtobuf(b []byte) error {
	var (
		r   = pb.NewBuffer(b)
		f   pb.Number
		err error
	)
	for !r.EOF() {
		f, _, err = r.DecodeTag()
		if err != nil {
			break
		}
		switch f {
		case 1:
			m.SetSymbol(r.ReadString())
		case 2:
			m.SetExchange(r.ReadString())
		case 3:
			m.SetTradeAccount(r.ReadString())
		case 4:
			m.SetClientOrderID(r.ReadString())
		case 5:
			m.SetBuySell(BuySellEnum(r.ReadInt32()))
		case 6:
			m.SetFillPrice(r.ReadFloat64())
		case 7:
			m.SetFillQuantity(r.ReadFloat64())
		case 8:
			m.SetFreeFormText(r.ReadString())
		default:
			r.Consume()
		}
		if err = r.Error(); err != nil {
			return err
		}
	}
	return err
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m AddCorrectingOrderFillBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 309)
	w.WriteString(1, m.Symbol())
	w.WriteString(2, m.Exchange())
	w.WriteString(3, m.TradeAccount())
	w.WriteString(4, m.ClientOrderID())
	w.WriteVarint32(5, int32(m.BuySell()))
	w.WriteFixed64Float64(6, m.FillPrice())
	w.WriteFixed64Float64(7, m.FillQuantity())
	w.WriteString(8, m.FreeFormText())
	return w.Finish(), nil
}
