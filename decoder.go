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
	decoded := make([]byte, 160)
	cDecoded := (*C.short)(unsafe.Pointer(&decoded[0]))
	var cEncoded *C.uint8_t
	if len(data) > 0 {
		cEncoded = (*C.uint8_t)(unsafe.Pointer(&data[0]))
	}else{
		cEncoded = (*C.uint8_t)(C.NULL)
	}

	if len(data) == 0{
		C.bcg729Decoder(thiz.dec, cEncoded, 0, 1, 0, 0, cDecoded)
		return decoded
	}


	if len(data) < 8 {
		//TODO: should consume 2 bytes every time decode in loop? or just decode once
		C.bcg729Decoder(thiz.dec, cEncoded, C.uint8_t(len(data)), 0, 1, 0, cDecoded)
		return decoded
	}

	C.bcg729Decoder(thiz.dec, cEncoded, C.uint8_t(len(data)), 0, 0, 0, cDecoded)
	return decoded
}
