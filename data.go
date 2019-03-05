package streams

import (
	"github.com/nareix/joy4/av"
)

func init() {
	format.ReigsterAll()
}

func getRawData() {
	fin, err := avutil.Open("out.flv")
	if err != nil {
		log.Errorf("avutil open faile(%v)", err)
		return
	}

	streams, err := fin.Streams()
	if err != nil {
		log.Errorf("streams failed(%v)", err)
		return
	}

	var vstream int
	for index, stream := range streams {
		if stream.Type().IsVideo() {
			vstream = index
		}
	}

	var pts uint64 = 10000

	for {

		select {
		case <-stop:
			break
		default:

		}

		var pkt av.Packet
		var err error
		if pkt, err = fin.ReadPacket(); err != nil {
			break
		}

		time.Sleep(time.Millisecond * 40)
		pts += 40

		psSys := psHeader()

		if pkt.IsKeyFrame {
			psSys = Sysheader()

		}

		psSys = psMapHeader()

		var lens int
		for lens < len(pkt.Data) {
			nallen := int(binary.BigEndian.Uint32(pkt.Data[lens : lens+4]))
			pkt.Data[lens+0], pkt.Data[lens+1], pkt.Data[lens+2], pkt.Data[lens+3] = 0x00, 0x00, 0x00, 0x01
			lens += (nallen + 4)

		}

	}

}
