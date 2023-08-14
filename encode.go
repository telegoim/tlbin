package tlbin

import (
	"encoding/binary"
	"math"
	"math/big"
)

type EncodeBuf struct {
	buf []byte
}

func NewEncodeBuf(cap int) *EncodeBuf {
	return &EncodeBuf{make([]byte, 0, cap)}
}

func (e *EncodeBuf) GetBuf() []byte {
	return e.buf
}

func (e *EncodeBuf) Append(bytes ...byte) {
	e.buf = append(e.buf, bytes...)
}

func (e *EncodeBuf) GetOffset() int {
	return len(e.buf)
}

func (e *EncodeBuf) Int16(s int16) {
	e.buf = append(e.buf, 0, 0)
	binary.LittleEndian.PutUint16(e.buf[len(e.buf)-2:], uint16(s))
}

func (e *EncodeBuf) UInt16(s uint16) {
	e.buf = append(e.buf, 0, 0)
	binary.LittleEndian.PutUint16(e.buf[len(e.buf)-2:], s)
}

func (e *EncodeBuf) Int(s int32) {
	e.buf = append(e.buf, 0, 0, 0, 0)
	binary.LittleEndian.PutUint32(e.buf[len(e.buf)-4:], uint32(s))
}

func (e *EncodeBuf) IntOffset(offset int, s int32) {
	binary.LittleEndian.PutUint32(e.buf[offset:], uint32(s))
}

func (e *EncodeBuf) UInt(s uint32) {
	e.buf = append(e.buf, 0, 0, 0, 0)
	binary.LittleEndian.PutUint32(e.buf[len(e.buf)-4:], s)
}

func (e *EncodeBuf) Long(s int64) {
	e.buf = append(e.buf, 0, 0, 0, 0, 0, 0, 0, 0)
	binary.LittleEndian.PutUint64(e.buf[len(e.buf)-8:], uint64(s))
}

func (e *EncodeBuf) Double(s float64) {
	e.buf = append(e.buf, 0, 0, 0, 0, 0, 0, 0, 0)
	binary.LittleEndian.PutUint64(e.buf[len(e.buf)-8:], math.Float64bits(s))
}

func (e *EncodeBuf) String(s string) {
	e.StringBytes([]byte(s))
}

func (e *EncodeBuf) BigInt(s *big.Int) {
	e.StringBytes(s.Bytes())
}

func (e *EncodeBuf) StringBytes(s []byte) {
	var res []byte
	size := len(s)
	if size < 254 {
		nl := 1 + size + (4-(size+1)%4)&3
		res = make([]byte, nl)
		res[0] = byte(size)
		copy(res[1:], s)

	} else {
		nl := 4 + size + (4-size%4)&3
		res = make([]byte, nl)
		binary.LittleEndian.PutUint32(res, uint32(size<<8|254))
		copy(res[4:], s)

	}
	e.buf = append(e.buf, res...)
}

func (e *EncodeBuf) Bytes(s []byte) {
	e.buf = append(e.buf, s...)
}

func (e *EncodeBuf) VectorInt(vList []int32) {
	x := make([]byte, 4+4+len(vList)*4)
	var c = int32(481674261)
	binary.LittleEndian.PutUint32(x, uint32(c))
	binary.LittleEndian.PutUint32(x[4:], uint32(len(vList)))
	i := 8
	for _, v := range vList {
		binary.LittleEndian.PutUint32(x[i:], uint32(v))
		i += 4
	}
	e.buf = append(e.buf, x...)
}

func (e *EncodeBuf) VectorLong(vList []int64) {
	x := make([]byte, 4+4+len(vList)*8)
	var c = int32(481674261)
	binary.LittleEndian.PutUint32(x, uint32(c))
	binary.LittleEndian.PutUint32(x[4:], uint32(len(vList)))
	i := 8
	for _, v := range vList {
		binary.LittleEndian.PutUint64(x[i:], uint64(v))
		i += 8
	}
	e.buf = append(e.buf, x...)
}

func (e *EncodeBuf) VectorString(vList []string) {
	x := make([]byte, 8)
	var c = int32(481674261)
	binary.LittleEndian.PutUint32(x, uint32(c))
	binary.LittleEndian.PutUint32(x[4:], uint32(len(vList)))
	e.buf = append(e.buf, x...)
	for _, v := range vList {
		e.String(v)
	}
}

func (e *EncodeBuf) VectorBytes(vList [][]byte) {
	x := make([]byte, 8)
	var c = int32(481674261)
	binary.LittleEndian.PutUint32(x, uint32(c))
	binary.LittleEndian.PutUint32(x[4:], uint32(len(vList)))
	e.buf = append(e.buf, x...)
	for _, v := range vList {
		e.StringBytes(v)
	}
}
