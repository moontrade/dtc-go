//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalPriceDataRejectPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetRequestID(r.ReadInt32())
		case 2:
			m.SetRejectText(r.ReadString())
		case 3:
			m.SetRejectReasonCode(HistoricalPriceDataRejectReasonCodeEnum(r.ReadInt16()))
		case 4:
			m.SetRetryTimeInSeconds(r.ReadUint16())
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

func (m HistoricalPriceDataRejectPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 802)
	w.WriteVarint32(1, m.RequestID())
	w.WriteString(2, m.RejectText())
	w.WriteVarint16(3, int16(m.RejectReasonCode()))
	w.WriteUvarint16(4, m.RetryTimeInSeconds())
	return w.Finish(), nil
}
