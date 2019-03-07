package packet



 const (
	START_CODE_PS  = 0x000001ba
	START_CODE_SYS = 0x000001bb
	START_CODE_MAP = 0x000001bc
	START_CODE_PES_VIDEO = 0x000001e0
	START_CODE_PES_AUDIO = 0x000001c0
)


/*
 * This implement from VLC source code
 * Notice:operate the bit,but not byte
 */

//BitsBuffer bits buffer
type BitsBuffer struct {
	iSize uint32
	iData uint32
	iMask uint8
	pData []byte
}

func bitsInit(isize uint32, buffer []byte) *BitsBuffer {

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

func bitsWrite(bits *BitsBuffer, count uint32, src uint64) *BitsBuffer {

	for count > 0 {
		count--
		if ((src >> count) & 0x01) != 0 {
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
