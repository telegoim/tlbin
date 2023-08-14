package tlbin

import "encoding/json"

type TLObject interface {
	Encode(x *EncodeBuf, layer int32) error
	Decode(dBuf *DecodeBuf) error
	String() string
	DebugString() string
}

func TLObjectToJson(object TLObject) (b []byte) {
	b, _ = json.Marshal(object)
	return
}
