package packet

import (
	"bytes"

	"github.com/32bitkid/bitreader"
)

type RtpParsePacket struct {
	psDenc *DecPSPackage
}

func NewRtpParsePacket() *RtpParsePacket {
	return &RtpParsePacket{
		psDenc: &DecPSPackage{
			rawData: make([]byte, MAXFrameLen),
			rawLen:  0,
		},
	}
}

// data包含 接受到完整一帧数据后，所有的payload, 解析出去后是一阵完整的raw数据
func (rtp *RtpParsePacket) Read(data []byte) ([]byte, error) {

	br := bitreader.NewReader(bytes.NewReader(data))

	if rtp.psDenc != nil {
		return rtp.psDenc.decPackHeader(br)
	}

	return nil, nil
}
