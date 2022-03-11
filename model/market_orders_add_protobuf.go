//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketOrdersAddBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetSymbolID(r.ReadUint32())
		case 2:
			m.SetSide(BuySellEnum(r.ReadInt32()))
		case 3:
			m.SetPrice(r.ReadFloat64())
		case 4:
			m.SetQuantity(r.ReadUint32())
		case 5:
			m.SetOrderID(r.ReadUint64())
		case 6:
			m.SetDateTime(DateTimeWithMicrosecondsInt(r.ReadInt64()))
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

func (m MarketOrdersAddBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 152)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteVarint32(2, int32(m.Side()))
	w.WriteFixed64Float64(3, m.Price())
	w.WriteUvarint32(4, m.Quantity())
	w.WriteUvarint64(5, m.OrderID())
	w.WriteVarint64(6, int64(m.DateTime()))
	return w.Finish(), nil
}
