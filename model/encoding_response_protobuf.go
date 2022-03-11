//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *EncodingResponseBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetProtocolVersion(r.ReadInt32())
		case 2:
			m.SetEncoding(EncodingEnum(r.ReadInt32()))
		case 3:
			m.SetProtocolType(r.ReadString())
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

func (m EncodingResponseBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 7)
	w.WriteVarint32(1, m.ProtocolVersion())
	w.WriteVarint32(2, int32(m.Encoding()))
	w.WriteString(3, m.ProtocolType())
	return w.Finish(), nil
}
