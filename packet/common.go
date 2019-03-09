package packet

const (
	STREAM_TYPE_H264 = 0x1b
	STREAM_TYPE_AAC  = 0x90
)

const (
	STREAM_ID_VIDEO = 0xe0
	STREAM_ID_AUDIO = 0xc0
)
 const (
	START_CODE_PS  		 = 0x000001ba
	START_CODE_SYS 	   	 = 0x000001bb
	START_CODE_MAP		 = 0x000001bc
	START_CODE_PES_VIDEO = 0x000001e0
	START_CODE_PES_AUDIO = 0x000001c0
)

//len limit
const (
	RTP_HEADER_LENGTH    int = 12
	PS_HEADER_LENGTH     int = 14
	SYSTEM_HEADER_LENGTH int = 18
	MAP_HEADER_LENGTH    int = 24
	PES_HEADER_LENGTH    int = 19
	RTP_LOAD_LENGTH      int = 1460
	PES_LOAD_LENGTH      int = 0xFFFF
)

/*
 * This implement from VLC source code
 * Notice:operate the bit,but not byte
 */

//BitsBuffer bits buffer
type BitsBuffer struct {
	iSize int
	iData int
	iMask uint8
	pData []byte
}

func bitsInit(isize int, buffer []byte) *BitsBuffer {

	bits := &BitsBuffer{
		iSize: isize,
		iData: 0,
		iMask: 0x80,
		pData: buffer,
	}

	if bits.pData == nil {
		bits.pData = make([]byte, isize)
	}
	return bits
}

func bitsWrite(bits *BitsBuffer, count int, src uint64) *BitsBuffer {

	for count > 0 {
		count--
		if ((src >> uint(count)) & 0x01) != 0 {
			bits.pData[bits.iData] |= bits.iMask
		} else {
			bits.pData[bits.iData] &= ^bits.iMask
		}

		bits.iMask >>= 1
		if bits.iMask == 0 {
			bits.iData++
			bits.iMask = 0x80
		}
	}

	return bits
}
