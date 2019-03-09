package packet

type EncPSPacket struct {
}


func (enc* EncPSPacket) Enc(pts, dts uint64, raw []byte ) []byte{

	// ps 

	// sys 
	// map 
	
	//pes 
	return nil
}

func (enc* EncPSPacket) addPSHeader(pts uint64) []byte{
	pack := make([]byte, PS_HEADER_LENGTH)

	bits := bitsInit(PS_HEADER_LENGTH, pack)

	bitsWrite(bits, 32, START_CODE_PS)               
	bitsWrite(bits, 2, 0x01)                   
	bitsWrite(bits, 3, (pts>>30)&0x07)     
	bitsWrite(bits, 1, 1)                     
	bitsWrite(bits, 15, (pts>>15)&0x7fff)  
	bitsWrite(bits, 1, 1)                      
	bitsWrite(bits, 15, (pts & 0x7fff))   
	bitsWrite(bits, 1, 1)                    

	bitsWrite(bits, 9, 0)                      
	bitsWrite(bits, 1, 1)                      

	bitsWrite(bits, 22, 255&0x3fffff)          
	bitsWrite(bits, 1, 1)                      

	bitsWrite(bits, 1, 1)                     
	bitsWrite(bits, 5, 0x1f)                   
	bitsWrite(bits, 3, 0)                     
	return bits.pData
}

func (enc* EncPSPacket) addSystemHeader(data []byte, vrates , arates int ) []byte{
	pack := make([]byte, SYSTEM_HEADER_LENGTH)

	bits := bitsInit(SYSTEM_HEADER_LENGTH, pack)

	bitsWrite(bits, 32, START_CODE_SYS)                 
	bitsWrite(bits, 16, uint64(SYSTEM_HEADER_LENGTH-6)) 
	bitsWrite(bits, 1, 1)                       
	bitsWrite(bits, 22, 40960)                  
	bitsWrite(bits, 1, 1)                      

	bitsWrite(bits, 6, 1)                       
	bitsWrite(bits, 1, 0)                       
	bitsWrite(bits, 1, 0)                       
	bitsWrite(bits, 1, 0)                       
	bitsWrite(bits, 1, 0)                       

	bitsWrite(bits, 1, 1)                       

	bitsWrite(bits, 5, 1)                      
	bitsWrite(bits, 1, 1)                       
	bitsWrite(bits, 7, 0x7f)                    

	// video stream bound
	bitsWrite(bits, 8, 0xe0)                   
	bitsWrite(bits, 2, 3)                       
	bitsWrite(bits, 1, 1)                      
	bitsWrite(bits, 13, uint64(vrates))                 

	// audio stream bound
	bitsWrite(bits, 8, 0xc0)  // 0xc0 音频
	bitsWrite(bits, 2, 3)
	bitsWrite(bits, 1, 0)
	bitsWrite(bits, 13, uint64(arates))

	return append(data, bits.pData...)
}

func (enc *EncPSPacket) addMapHeader(data []byte, crc32 uint64) []byte {

	pack := make([]byte, MAP_HEADER_LENGTH)

	bits := bitsInit(MAP_HEADER_LENGTH, pack)
	bitsWrite(bits, 32, START_CODE_MAP)                  
	bitsWrite(bits, 16, uint64(MAP_HEADER_LENGTH-6))    
	bitsWrite(bits, 1, 1)                           
	bitsWrite(bits, 2, 0xf)                         
	bitsWrite(bits, 5, 0)                          
	bitsWrite(bits, 7, 0xff)                        
	bitsWrite(bits, 1, 1)                          

	bitsWrite(bits, 16, 0)                          
	bitsWrite(bits, 16, 8)                          

	//video
	bitsWrite(bits, 8, STREAM_TYPE_H264)                        
	bitsWrite(bits, 8, STREAM_ID_VIDEO)                        
	bitsWrite(bits, 16, 0)                          
	// audio
	bitsWrite(bits, 8, STREAM_TYPE_AAC)                        // 0x90 G711
	bitsWrite(bits, 8, STREAM_ID_AUDIO)                        // 0x0c0 音频取值（0xc0-0xdf），通常为0xc0
	bitsWrite(bits, 16, 0)

	bitsWrite(bits, 8, crc32 >> 24)                        // CRC_32 : (32) CRC 32字段
	bitsWrite(bits, 8, (crc32 >>16)&0xFF)
	bitsWrite(bits, 8, (crc32 >>8) & 0xFF)
	bitsWrite(bits, 8, crc32 &0xFF)
	return append(data, bits.pData...)
}

func (enc *EncPSPacket) packPESHeader(data []byte, streamtype int, payloadlen uint32, pts, dts uint64) []byte {

	pack := make([]byte, PES_HEADER_LENGTH)

	bits := bitsInit(PES_HEADER_LENGTH, pack)

	bitsWrite(bits, 24, 0x01)  
	bitsWrite(bits, 8, uint64(streamtype))  

	bitsWrite(bits, 16, uint64(payloadlen+13)) 
	bitsWrite(bits, 2, 2) 
	bitsWrite(bits, 2, 0)  
	bitsWrite(bits, 1, 0)  
	bitsWrite(bits, 1, 0)  
	bitsWrite(bits, 1, 0)   
	bitsWrite(bits, 1, 0)  

	bitsWrite(bits, 2, 0x03)  
	bitsWrite(bits, 1, 0)   
	bitsWrite(bits, 1, 0)   

	bitsWrite(bits, 1, 0)    
	bitsWrite(bits, 1, 0)   
	bitsWrite(bits, 1, 0)   
	bitsWrite(bits, 1, 0)  

	bitsWrite(bits, 8, 10)  
	// pts dts //
	bitsWrite(bits, 4, 3)               
	bitsWrite(bits, 3, (pts>>30)&0x07)   
	bitsWrite(bits, 1, 1)                
	bitsWrite(bits, 15, (pts>>15)&0x7fff)
	bitsWrite(bits, 1, 1)
	bitsWrite(bits, 15, pts&0x7fff)
	bitsWrite(bits, 1, 1)

	bitsWrite(bits, 4, 1)                 
	bitsWrite(bits, 3, (dts>>30)&0x07)    
	bitsWrite(bits, 15, (dts>>15)&0x7fff)
	bitsWrite(bits, 1, 1)
	bitsWrite(bits, 15, dts&0x7fff)
	bitsWrite(bits, 1, 1)
	return append(bits.pData, data...)
}

