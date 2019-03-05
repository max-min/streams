package packet

/*
import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
)
*/

/*
 * this implement from VLC source code
 * operate the bit,but not byte
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
