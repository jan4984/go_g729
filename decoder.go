package go_g729

//#cgo CFLAGS: -I bcg729
//#cgo LDFLAGS: -L bcg729 -lbcg729 -lm
//#include "bcg729/decoder.h"
import "C"
import "unsafe"

type Decoder struct{
	dec *C.bcg729DecoderChannelContextStruct
}

func NewDecoder() *Decoder{
	return &Decoder{C.initBcg729DecoderChannel()}
}

func (thiz *Decoder)Destroy(){
	C.closeBcg729DecoderChannel(thiz.dec)
}

func (thiz *Decoder)Decode(data []byte) []byte{
	if len(data) > 10 {
		panic("g729b one frame must less than 10 bytes")
	}
	decoded := make([]byte, 160)
	cDecoded := (*C.short)(unsafe.Pointer(&decoded[0]))
	cEncoded := (*C.uint8_t)(unsafe.Pointer(&data[0]))

	if len(data) == 0{
		C.bcg729Decoder(thiz.dec, C.NULL, 0, 1, 0, 0, cDecoded)
		return decoded
	}


	if len(data) < 8 {
		//TODO: should consume 2 bytes every time decode in loop? or just decode once
		C.bcg729Decoder(thiz.dec, cEncoded, len(data), 0, 1, 0, cDecoded)
		return decoded
	}

	C.bcg729Decoder(thiz.dec, cEncoded, len(data), 0, 0, 0, cDecoded)
	return decoded
}
