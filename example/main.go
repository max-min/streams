package main

import (
	"streams/packet"

	"github.com/nareix/joy4/av"
	"github.com/nareix/joy4/av/avutil"
	log "github.com/sirupsen/logrus"
)

func main() {

	rtp := packet.NewRtpService("", packet.UDPTransfer)

	rtp.Service("127.0.0.1", "172.20.25.2", 10086, 10087)

	f, err := avutil.Open("test.flv")
	if err != nil {
		log.Errorf("read file error(%v)", err)
		return
	}

	var pts uint64 = 10000
	streams, _ := f.Streams()
	var vindex int8
	for i, stream := range streams {
		if stream.Type() == av.H264 {
			vindex = int8(i)
			break
		}
	}

	for {
		var pkt av.Packet
		var err error
		if pkt, err = f.ReadPacket(); err != nil {
			log.Errorf("read packet error(%v)", err)
			return
		}
		if pkt.Idx != vindex {
			continue
		}

		rtp.Send2data(pkt.Data, pkt.IsKeyFrame, pts)
		pts += 40
	}
}
