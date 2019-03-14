package packet

import (
	"github.com/32bitkid/bitreader"
)

type DecPSPackage struct {
	systemClockReferenceBase      uint64
	systemClockReferenceExtension uint64
	programMuxRate                uint32
	streamType                    int
}

func (dec *DecPSPackage) decPSHeader(br bitreader.BitReader) ([]byte, error) {

	startcode, err := br.Read32(32)
	if err != nil {
		return nil, err
	}
	if startcode != StartCodePS {
		return nil, ErrNotFoundStartCode
	}

	if marker, err := br.Read32(2); err != nil {
		return nil, err
	} else if marker != 0x01 {
		return nil, ErrMarkerBit
	}

	if s, err := br.Read32(3); err != nil {
		return nil, err
	} else {
		dec.systemClockReferenceBase |= uint64(s << 30)
	}
	if marker, err := br.Read32(1); err != nil {
		return nil, err
	} else if marker != 0x01 {
		return nil, ErrMarkerBit
	}

	if s, err := br.Read32(15); err != nil {
		return nil, err
	} else {
		dec.systemClockReferenceBase |= uint64(s << 15)
	}
	if marker, err := br.Read32(1); err != nil {
		return nil, err
	} else if marker != 0x01 {
		return nil, ErrMarkerBit
	}
	if s, err := br.Read32(15); err != nil {
		return nil, err
	} else {
		dec.systemClockReferenceBase |= uint64(s)
	}
	if marker, err := br.Read32(1); err != nil {
		return nil, err
	} else if marker != 0x01 {
		return nil, ErrMarkerBit
	}
	if s, err := br.Read32(9); err != nil {
		return nil, err
	} else {
		dec.systemClockReferenceExtension |= uint64(s)
	}
	if marker, err := br.Read32(1); err != nil {
		return nil, err
	} else if marker != 0x01 {
		return nil, ErrMarkerBit
	}

	if pmr, err := br.Read32(22); err != nil {
		return nil, err
	} else {
		dec.programMuxRate |= pmr
	}
	if marker, err := br.Read32(1); err != nil {
		return nil, err
	} else if marker != 0x01 {
		return nil, ErrMarkerBit
	}
	if marker, err := br.Read32(1); err != nil {
		return nil, err
	} else if marker != 0x01 {
		return nil, ErrMarkerBit
	}

	if err := br.Skip(5); err != nil {
		return nil, err
	}
	if psl, err := br.Read32(3); err != nil {
		return nil, err
	} else {
		br.Skip(uint(psl * 8))
	}

	// 判断是否位关键帧， I帧会有system头 systemap头
	for {
		nextStartCode, err := br.Read32(32)
		if err != nil {
			return nil, err
		}

		// how to do that ??
		switch nextStartCode {
		case StartCodeSYS:
			//dec.decSystemHeader(br)
		case StartCodeMAP:
			//dec.decSystemMapHeader(br)
		case StartCodeVideo:
			fallthrough
		case StartCodeAudio:
			dec.decPESHeader(br)
		}
	}

	return nil, ErrNotFoundStartCode
}

func (dec *DecPSPackage) decSystemHeader(br bitreader.BitReader) ([]byte, error) {
	return nil, nil
}

func (dec *DecPSPackage) decSystemMapHeader(br bitreader.BitReader) ([]byte, error) {
	return nil, nil
}

func (dec *DecPSPackage) decPESHeader(br bitreader.BitReader) ([]byte, error) {
	return nil, nil
}
