package dataPacket

import (
	"errors"
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/google/uuid"
	"goproxy/item"
	"goproxy/nbt"

	//"github.com/sandertv/gophertunnel/minecraft/nbt"
	"image/color"
	"io"
	"io/ioutil"
	"math"
	"unsafe"
)

const (
	EntityDataByte uint32 = iota
	EntityDataInt16
	EntityDataInt32
	EntityDataFloat32
	EntityDataString
	EntityDataCompoundTag
	EntityDataBlockPos
	EntityDataInt64
	EntityDataVec3
)

type PacketReader struct {
	r interface {
		io.Reader
		io.ByteReader
	}
	shieldID int32
}

// NewReader creates a new Reader using the io.ByteReader passed as underlying source to read bytes from.
func NewReader(r interface {
	io.Reader
	io.ByteReader
}, shieldID int32) *PacketReader {
	return &PacketReader{r: r, shieldID: shieldID}
}

// Uint8 reads a uint8 from the underlying buffer.
func (r *PacketReader) Uint8(x *uint8) {
	var err error
	*x, err = r.r.ReadByte()
	if err != nil {
		r.panic(err)
	}
}

// Bool reads a bool from the underlying buffer.
func (r *PacketReader) Bool(x *bool) {
	u, err := r.r.ReadByte()
	if err != nil {
		r.panic(err)
	}
	*x = *(*bool)(unsafe.Pointer(&u))
}

// errStringTooLong is an error set if a string decoded using the String method has a length that is too long.
var errStringTooLong = errors.New("string length overflows a 32-bit integer")

// StringUTF ...
func (r *PacketReader) StringUTF(x *string) {
	var length int16
	r.Int16(&length)
	l := int(length)
	if l > math.MaxInt16 {
		r.panic(errStringTooLong)
	}
	data := make([]byte, l)
	if _, err := r.r.Read(data); err != nil {
		r.panic(err)
	}
	*x = *(*string)(unsafe.Pointer(&data))
}

// String reads a string from the underlying buffer.
func (r *PacketReader) String(x *string) {
	var length uint32
	r.Varuint32(&length)
	l := int(length)
	if l > math.MaxInt32 {
		r.panic(errStringTooLong)
	}
	data := make([]byte, l)
	if _, err := r.r.Read(data); err != nil {
		r.panic(err)
	}
	*x = *(*string)(unsafe.Pointer(&data))
}

// ByteSlice reads a byte slice from the underlying buffer, similarly to String.
func (r *PacketReader) ByteSlice(x *[]byte) {
	var length uint32
	r.Varuint32(&length)
	l := int(length)
	if l > math.MaxInt32 {
		r.panic(errStringTooLong)
	}
	data := make([]byte, l)
	if _, err := r.r.Read(data); err != nil {
		r.panic(err)
	}
	*x = data
}

// Vec3 reads three float32s into an mgl32.Vec3 from the underlying buffer.
func (r *PacketReader) Vec3(x *mgl32.Vec3) {
	r.Float32(&x[0])
	r.Float32(&x[1])
	r.Float32(&x[2])
}

// Vec2 reads two float32s into an mgl32.Vec2 from the underlying buffer.
func (r *PacketReader) Vec2(x *mgl32.Vec2) {
	r.Float32(&x[0])
	r.Float32(&x[1])
}

// ByteFloat reads a rotational float32 from a single byte.
func (r *PacketReader) ByteFloat(x *float32) {
	var v uint8
	r.Uint8(&v)
	*x = float32(v) * (360.0 / 256.0)
}

// VarRGBA reads a color.RGBA x from a varuint32.
func (r *PacketReader) VarRGBA(x *color.RGBA) {
	var v uint32
	r.Varuint32(&v)
	*x = color.RGBA{
		R: byte(v),
		G: byte(v >> 8),
		B: byte(v >> 16),
		A: byte(v >> 24),
	}
}

// Bytes reads the leftover bytes into a byte slice.
func (r *PacketReader) Bytes(p *[]byte) {
	var err error
	*p, err = ioutil.ReadAll(r.r)
	if err != nil {
		r.panic(err)
	}
}

func (r *PacketReader) BytesLength(p []byte) {
	var err error
	_, err = r.r.Read(p)
	if err != nil {
		r.panic(err)
	}
}

func (r *PacketReader) BlockCoords(x *int32, y *uint32, z *int32) {
	r.Varint32(x)
	r.Varuint32(y)
	r.Varint32(z)
}

func (r *PacketReader) Item() *item.Item {

	var iid int32
	r.Varint32(&iid)
	if iid <= 0 {
		return item.NewItem(0, 0, 0, map[string]interface{}{})
	}
	var auxv int32
	r.Varint32(&auxv)
	data := auxv >> 8
	if data == 0x7fff {
		data = -1
	}
	cnt := auxv & 0xff
	//fmt.Println("icount", cnt)
	//nbt := ""
	var nbtlen int16
	var nbtdata map[string]interface{}
	r.Int16(&nbtlen)
	if nbtlen > 0 {
		r.NBT(&nbtdata, nbt.LittleEndian)
	}
	var canp int32
	r.Varint32(&canp)
	zz := ""
	if canp > 0 {
		for i := int32(0); i < canp; i++ {
			r.String(&zz)
		}
	}
	var cand int32
	r.Varint32(&cand)
	if cand > 0 {
		for i := int32(0); i < canp; i++ {
			r.String(&zz)
		}
	}
	return item.NewItem(iid, data, cnt, nbtdata)
}

func (r *PacketReader) EntityMetadata(x *map[uint32]interface{}) {
	*x = map[uint32]interface{}{}

	var count uint32
	r.Varuint32(&count)
	r.LimitUint32(count, 256)
	for i := uint32(0); i < count; i++ {
		var key, dataType uint32
		r.Varuint32(&key)
		r.Varuint32(&dataType)
		switch dataType {
		case EntityDataByte:
			var v byte
			r.Uint8(&v)
			(*x)[key] = v
		case EntityDataInt16:
			var v int16
			r.Int16(&v)
			(*x)[key] = v
		case EntityDataInt32:
			var v int32
			r.Varint32(&v)
			(*x)[key] = v
		case EntityDataFloat32:
			var v float32
			r.Float32(&v)
			(*x)[key] = v
		case EntityDataString:
			var v string
			r.String(&v)
			(*x)[key] = v
		case EntityDataCompoundTag:
			//todo nbt
		case EntityDataBlockPos:
			//todo blockpos
		case EntityDataInt64:
			var v int64
			r.Varint64(&v)
			(*x)[key] = v
		case EntityDataVec3:
			//todo vec3 x y z
		default:
			r.UnknownEnumOption(dataType, "entity metadata")

		}
		(*x)[key] = []interface{}{dataType, (*x)[key]}
	}
}

// NBT reads a compound tag into a map from the underlying buffer.

func (r *PacketReader) NBT(m *map[string]interface{}, encoding nbt.Encoding) {
	if err := nbt.NewDecoderWithEncoding(r.r, encoding).Decode(m); err != nil {
		r.panic(err)
	}
}

func (r *PacketReader) ByteRotation(x *float32) {
	g, _ := r.r.ReadByte()
	a := float32(g * (360 / 256))
	x = &a
}

// NBTList reads a list of NBT tags from the underlying buffer.
func (r *PacketReader) NBTList(m *[]interface{}, encoding nbt.Encoding) {
	if err := nbt.NewDecoderWithEncoding(r.r, encoding).Decode(m); err != nil {
		r.panic(err)
	}
}

// UUID reads a uuid.UUID from the underlying buffer.
func (r *PacketReader) UUID(x *uuid.UUID) {
	b := make([]byte, 16)
	if _, err := r.r.Read(b); err != nil {
		r.panic(err)
	}

	// The UUIDs we read are Little Endian, but the uuid library is based on Big Endian UUIDs, so we need to
	// reverse the two int64s the UUID is composed of, then reverse their bytes too.
	b = append(b[8:], b[:8]...)
	var arr [16]byte
	for i, j := 0, 15; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = b[j], b[i]
	}
	*x = arr
}

// LimitUint32 checks if the value passed is lower than the limit passed. If not, the Reader panics.
func (r *PacketReader) LimitUint32(value uint32, max uint32) {
	if max == math.MaxUint32 {
		// Account for 0-1 overflowing into max.
		max = 0
	}
	if value > max {
		r.panicf("uint32 %v exceeds maximum of %v", value, max)
	}
}

// LimitInt32 checks if the value passed is lower than the limit passed and higher than the minimum. If not,
// the Reader panics.
func (r *PacketReader) LimitInt32(value int32, min, max int32) {
	if value < min {
		r.panicf("int32 %v exceeds minimum of %v", value, min)
	} else if value > max {
		r.panicf("int32 %v exceeds maximum of %v", value, max)
	}
}

// UnknownEnumOption panics with an unknown enum option error.
func (r *PacketReader) UnknownEnumOption(value interface{}, enum string) {
	r.panicf("unknown value '%v' for enum type '%v'", value, enum)
}

// InvalidValue panics with an error indicating that the value passed is not valid for a specific field.
func (r *PacketReader) InvalidValue(value interface{}, forField, reason string) {
	r.panicf("invalid value '%v' for %v: %v", value, forField, reason)
}

// errVarIntOverflow is an error set if one of the Varint methods encounters a varint that does not terminate
// after 5 or 10 bytes, depending on the data type read into.
var errVarIntOverflow = errors.New("varint overflows integer")

// Varint64 reads up to 10 bytes from the underlying buffer into an int64.
func (r *PacketReader) Varint64(x *int64) {
	var ux uint64
	for i := 0; i < 70; i += 7 {
		b, err := r.r.ReadByte()
		if err != nil {
			r.panic(err)
		}

		ux |= uint64(b&0x7f) << uint32(i)
		if b&0x80 == 0 {
			*x = int64(ux >> 1)
			if ux&1 != 0 {
				*x = ^*x
			}
			return
		}
	}
	//	r.panic(errVarIntOverflow)
}

// Varuint64 reads up to 10 bytes from the underlying buffer into a uint64.
func (r *PacketReader) Varuint64(x *uint64) {
	var v uint64
	for i := 0; i < 70; i += 7 {
		b, err := r.r.ReadByte()
		if err != nil {
			r.panic(err)
		}

		v |= uint64(b&0x7f) << uint32(i)
		if b&0x80 == 0 {
			*x = v
			return
		}
	}
	//r.panic(errVarIntOverflow)
}

// Varint32 reads up to 5 bytes from the underlying buffer into an int32.
func (r *PacketReader) Varint32(x *int32) {
	var ux uint32
	for i := 0; i < 35; i += 7 {
		b, err := r.r.ReadByte()
		if err != nil {
			r.panic(err)
		}

		ux |= uint32(b&0x7f) << uint32(i)
		if b&0x80 == 0 {
			*x = int32(ux >> 1)
			if ux&1 != 0 {
				*x = ^*x
			}
			return
		}
	}
	//r.panic(errVarIntOverflow)
}

// Varuint32 reads up to 5 bytes from the underlying buffer into a uint32.
func (r *PacketReader) Varuint32(x *uint32) {
	var v uint32
	for i := 0; i < 35; i += 7 {
		b, err := r.r.ReadByte()
		if err != nil {
			r.panic(err)
		}

		v |= uint32(b&0x7f) << uint32(i)
		if b&0x80 == 0 {
			*x = v
			return
		}
	}
	//	r.panic(errVarIntOverflow)
}

// panicf panics with the format and values passed and assigns the error created to the Reader.
func (r *PacketReader) panicf(format string, a ...interface{}) {
	panic(fmt.Errorf(format, a...))
}

// panic panics with the error passed, similarly to panicf.
func (r *PacketReader) panic(err error) {
	panic(err)
}
