package dataPacket

import (
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/google/uuid"
	"goproxy/item"
	"goproxy/nbt"

	//"github.com/sandertv/gophertunnel/minecraft/nbt"
	"image/color"
	"io"
	"unsafe"
)

type PacketWriter struct {
	w interface {
		io.Writer
		io.ByteWriter
	}
	shieldID int32
}

// NewWriter creates a new initialised Writer with an underlying io.ByteWriter to write to.
func NewWriter(w interface {
	io.Writer
	io.ByteWriter
}, shieldID int32) *PacketWriter {
	return &PacketWriter{w: w, shieldID: shieldID}
}

func (w *PacketWriter) EntityMetadata(x *map[uint32]interface{}) {
	var count uint32
	w.Varuint32(&count)
	for i := uint32(0); i < count; i++ {
		var key, dataType uint32
		w.Varuint32(&key)
		w.Varuint32(&dataType)
		switch dataType {
		case EntityDataByte:
			var v byte
			w.Uint8(&v)
			(*x)[key] = v
		case EntityDataInt16:
			var v int16
			w.Int16(&v)
			(*x)[key] = v
		case EntityDataInt32:
			var v int32
			w.Varint32(&v)
			(*x)[key] = v
		case EntityDataFloat32:
			var v float32
			w.Float32(&v)
			(*x)[key] = v
		case EntityDataString:
			var v string
			w.String(&v)
			(*x)[key] = v
		case EntityDataCompoundTag:
			//todo nbt
		case EntityDataBlockPos:
			//todo blockpos
		case EntityDataInt64:
			var v int64
			w.Varint64(&v)
			(*x)[key] = v
		case EntityDataVec3:
			//todo vec3 x y z
		default:
			w.UnknownEnumOption(dataType, "entity metadata")

		}
	}

}

// Uint8 writes a uint8 to the underlying buffer.
func (w *PacketWriter) Uint8(x *uint8) {
	_ = w.w.WriteByte(*x)
}

func (w *PacketWriter) ByteRotation(x *float32) {
	g := byte(*x / (float32(360) / float32(256)))
	_ = w.w.WriteByte(g)
}

// Bool writes a bool as either 0 or 1 to the underlying buffer.
func (w *PacketWriter) Bool(x *bool) {
	_ = w.w.WriteByte(*(*byte)(unsafe.Pointer(x)))
}

// StringUTF ...
func (w *PacketWriter) StringUTF(x *string) {
	l := int16(len(*x))
	w.Int16(&l)
	_, _ = w.w.Write([]byte(*x))
}

// String writes a string, prefixed with a varuint32, to the underlying buffer.
func (w *PacketWriter) String(x *string) {
	l := uint32(len(*x))
	w.Varuint32(&l)
	_, _ = w.w.Write([]byte(*x))
}

// ByteSlice writes a []byte, prefixed with a varuint32, to the underlying buffer.
func (w *PacketWriter) ByteSlice(x *[]byte) {
	l := uint32(len(*x))
	w.Varuint32(&l)
	_, _ = w.w.Write(*x)
}

// Bytes appends a []byte to the underlying buffer.
func (w *PacketWriter) Bytes(x *[]byte) {
	_, _ = w.w.Write(*x)
}

func (w *PacketWriter) BytesCopy(x []byte) {
	_, _ = w.w.Write(x)
}

// ByteFloat writes a rotational float32 as a single byte to the underlying buffer.
func (w *PacketWriter) ByteFloat(x *float32) {
	_ = w.w.WriteByte(byte(*x / (360.0 / 256.0)))
}

// Vec3 writes an mgl32.Vec3 as 3 float32s to the underlying buffer.
func (w *PacketWriter) Vec3(x *mgl32.Vec3) {
	w.Float32(&x[0])
	w.Float32(&x[1])
	w.Float32(&x[2])
}

// Vec2 writes an mgl32.Vec2 as 2 float32s to the underlying buffer.
func (w *PacketWriter) Vec2(x *mgl32.Vec2) {
	w.Float32(&x[0])
	w.Float32(&x[1])
}

// VarRGBA writes a color.RGBA x as a varuint32 to the underlying buffer.
func (w *PacketWriter) VarRGBA(x *color.RGBA) {
	val := uint32(x.R) | uint32(x.G)<<8 | uint32(x.B)<<16 | uint32(x.A)<<24
	w.Varuint32(&val)
}

// UUID writes a UUID to the underlying buffer.
func (w *PacketWriter) UUID(x *uuid.UUID) {
	b := append((*x)[8:], (*x)[:8]...)
	for i, j := 0, 15; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	_, _ = w.w.Write(b)
}

// Varint64 writes an int64 as 1-10 bytes to the underlying buffer.
func (w *PacketWriter) Varint64(x *int64) {
	u := *x
	ux := uint64(u) << 1
	if u < 0 {
		ux = ^ux
	}
	for ux >= 0x80 {
		_ = w.w.WriteByte(byte(ux) | 0x80)
		ux >>= 7
	}
	_ = w.w.WriteByte(byte(ux))
}

// Varuint64 writes a uint64 as 1-10 bytes to the underlying buffer.
func (w *PacketWriter) Varuint64(x *uint64) {
	u := *x
	for u >= 0x80 {
		_ = w.w.WriteByte(byte(u) | 0x80)
		u >>= 7
	}
	_ = w.w.WriteByte(byte(u))
}

// Varint32 writes an int32 as 1-5 bytes to the underlying buffer.
func (w *PacketWriter) Varint32(x *int32) {
	u := *x
	ux := uint32(u) << 1
	if u < 0 {
		ux = ^ux
	}
	for ux >= 0x80 {
		_ = w.w.WriteByte(byte(ux) | 0x80)
		ux >>= 7
	}
	_ = w.w.WriteByte(byte(ux))
}

// Varuint32 writes a uint32 as 1-5 bytes to the underlying buffer.
func (w *PacketWriter) Varuint32(x *uint32) {
	u := *x
	for u >= 0x80 {
		_ = w.w.WriteByte(byte(u) | 0x80)
		u >>= 7
	}
	_ = w.w.WriteByte(byte(u))
}

func (w *PacketWriter) Item(it *item.Item) {
	if it.GetID() == 0 {
		w.Varint32(&it.ID)
		return
	}
	w.Varint32(&it.ID)
	aux := ((it.GetDamage() & 0x7fff) << 8) | it.GetCount()
	w.Varint32(&aux)

	nbtd := it.NBT
	nlen := int16(len(nbtd))
	if nlen != 0 {
		w.Int16(&nlen)
		w.NBT(&nbtd, nbt.LittleEndian)
	} else {
		w.Int16(&nlen)
	}

	todo := int32(0)
	w.Varint32(&todo) //CanPlaceOn
	w.Varint32(&todo) //CanDestroy

}

// NBT writes a map as NBT to the underlying buffer using the encoding passed.
func (w *PacketWriter) NBT(x *map[string]interface{}, encoding nbt.Encoding) {
	if err := nbt.NewEncoderWithEncoding(w.w, encoding).Encode(*x); err != nil {
		panic(err)
	}
}

// NBTList writes a slice as NBT to the underlying buffer using the encoding passed.
func (w *PacketWriter) NBTList(x *[]interface{}, encoding nbt.Encoding) {
	if err := nbt.NewEncoderWithEncoding(w.w, encoding).Encode(*x); err != nil {
		panic(err)
	}
}

func (w *PacketWriter) BlockCoords(x *int32, y *uint32, z *int32) {
	w.Varint32(x)
	w.Varuint32(y)
	w.Varint32(z)
}

// NBT writes a map as NBT to the underlying buffer using the encoding passed.
/*func (w *PacketWriter) NBT(x *map[string]interface{}, encoding nbt.Encoding) {
	if err := nbt.NewEncoderWithEncoding(w.w, encoding).Encode(*x); err != nil {
		panic(err)
	}
}
*/
// NBTList writes a slice as NBT to the underlying buffer using the encoding passed.
/*func (w *PacketWriter) NBTList(x *[]interface{}, encoding nbt.Encoding) {
	if err := nbt.NewEncoderWithEncoding(w.w, encoding).Encode(*x); err != nil {
		panic(err)
	}
}
*/
// UnknownEnumOption panics with an unknown enum option error.
func (w *PacketWriter) UnknownEnumOption(value interface{}, enum string) {
	w.panicf("unknown value '%v' for enum type '%v'", value, enum)
}

// InvalidValue panics with an invalid value error.
func (w *PacketWriter) InvalidValue(value interface{}, forField, reason string) {
	w.panicf("invalid value '%v' for %v: %v", value, forField, reason)
}

// panicf panics with the format and values passed.
func (w *PacketWriter) panicf(format string, a ...interface{}) {
	panic(fmt.Errorf(format, a...))
}
