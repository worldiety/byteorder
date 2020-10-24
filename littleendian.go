/*
 * Copyright 2020 Torben Schinke
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package byteorder

import "math"

var _ = (LittleEndian)(nil)

// LE is an Alias for LittleEndian.
type LE = LittleEndian

// LittleEndian defines little-endian serialization.
type LittleEndian []byte

// ReadUint16 reads the first 2 bytes. An optimized unsafe read on a little endian machine may look like this:
//  b := *(*uint16)(unsafe.Pointer(uintptr(unsafe.Pointer(*(**uint16)(unsafe.Pointer(&f.Bytes))))+uintptr(f.Pos)))
//	f.Pos+=2
// Warning: such unsafe code relies on UB (undefined behavior) for unaligned access but x86 allows that with a
// speed penalty (when crossing cache lines). Panics when len(b) < 2.
func (b LittleEndian) ReadUint16() uint16 {
	_ = b[1] // bounds check hint to compiler; see golang.org/issue/14808

	return uint16(b[0]) | uint16(b[1])<<8
}

// WriteUint16 writes an unsigned 16 bit integer. Panics when len(b) < 2.
func (b LittleEndian) WriteUint16(v uint16) {
	_ = b[1] // bounds check hint to compiler; see golang.org/issue/14808
	b[0] = byte(v)
	b[1] = byte(v >> 8) //nolint:gomnd
}

// ReadUint24 reads the first 3 bytes. Panics when len(b) < 3.
func (b LittleEndian) ReadUint24() uint32 {
	_ = b[2] // bounds check hint to compiler; see golang.org/issue/14808

	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16
}

// WriteUint24 writes the first 3 bytes. Panics when len(b) < 3.
func (b LittleEndian) WriteUint24(v uint32) {
	_ = b[2] // early bounds check to guarantee safety of writes below
	b[0] = byte(v)
	b[1] = byte(v >> 8)  //nolint:gomnd
	b[2] = byte(v >> 16) //nolint:gomnd
}

// ReadUint32 reads the first 4 bytes. An optimized unsafe read may look like this:
//  b := *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&f.Bytes))))+uintptr(f.Pos)))
//	f.Pos+=4
// Warning: such unsafe code relies on UB (undefined behavior) for unaligned access but x86 allows that with a
// speed penalty (when crossing cache lines). Panics when len(b) < 4.
func (b LittleEndian) ReadUint32() uint32 {
	_ = b[3] // bounds check hint to compiler; see golang.org/issue/14808

	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
}

// WriteUint32 writes the first 4 bytes. Panics when len(b) < 4.
func (b LittleEndian) WriteUint32(v uint32) {
	_ = b[3] // bounds check hint to compiler; see golang.org/issue/14808
	b[0] = byte(v)
	b[1] = byte(v >> 8)  //nolint:gomnd
	b[2] = byte(v >> 16) //nolint:gomnd
	b[3] = byte(v >> 24) //nolint:gomnd
}

// ReadUint40 reads the first 5 bytes. Panics when len(b) < 5.
func (b LittleEndian) ReadUint40() uint64 {
	_ = b[4] // bounds check hint to compiler; see golang.org/issue/14808

	return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 | uint64(b[4])<<32
}

// WriteUint40 writes the first 5 bytes. Panics when len(b) < 5.
func (b LittleEndian) WriteUint40(v uint64) {
	_ = b[4] // bounds check hint to compiler; see golang.org/issue/14808
	b[0] = byte(v)
	b[1] = byte(v >> 8)  //nolint:gomnd
	b[2] = byte(v >> 16) //nolint:gomnd
	b[3] = byte(v >> 24) //nolint:gomnd
	b[4] = byte(v >> 32) //nolint:gomnd
}

// ReadUint48 reads the first 6 bytes. Panics when len(b) < 6.
func (b LittleEndian) ReadUint48() uint64 {
	_ = b[5] // bounds check hint to compiler; see golang.org/issue/14808

	return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
		uint64(b[4])<<32 | uint64(b[5])<<40
}

// WriteUint48 writes the first 6 bytes. Panics when len(b) < 6.
func (b LittleEndian) WriteUint48(v uint64) {
	_ = b[5] // bounds check hint to compiler; see golang.org/issue/14808
	b[0] = byte(v)
	b[1] = byte(v >> 8)  //nolint:gomnd
	b[2] = byte(v >> 16) //nolint:gomnd
	b[3] = byte(v >> 24) //nolint:gomnd
	b[4] = byte(v >> 32) //nolint:gomnd
	b[5] = byte(v >> 40) //nolint:gomnd
}

// ReadUint56 reads the first 7 bytes. Panics when len(b) < 7.
func (b LittleEndian) ReadUint56() uint64 {
	_ = b[6] // bounds check hint to compiler; see golang.org/issue/14808

	return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
		uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48
}

// WriteUint56 writes the first 7 bytes. Panics when len(b) < 7.
func (b LittleEndian) WriteUint56(v uint64) {
	_ = b[6] // bounds check hint to compiler; see golang.org/issue/14808
	b[0] = byte(v)
	b[1] = byte(v >> 8)  //nolint:gomnd
	b[2] = byte(v >> 16) //nolint:gomnd
	b[3] = byte(v >> 24) //nolint:gomnd
	b[4] = byte(v >> 32) //nolint:gomnd
	b[5] = byte(v >> 40) //nolint:gomnd
	b[6] = byte(v >> 48) //nolint:gomnd
}

// ReadUint64 reads the first 8 bytes. Panics when len(b) < 8.
func (b LittleEndian) ReadUint64() uint64 {
	_ = b[7] // bounds check hint to compiler; see golang.org/issue/14808

	return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
		uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
}

// WriteUint64 writes the first 8 bytes. Panics when len(b) < 8.
func (b LittleEndian) WriteUint64(v uint64) {
	_ = b[7] // bounds check hint to compiler; see golang.org/issue/14808
	b[0] = byte(v)
	b[1] = byte(v >> 8)  //nolint:gomnd
	b[2] = byte(v >> 16) //nolint:gomnd
	b[3] = byte(v >> 24) //nolint:gomnd
	b[4] = byte(v >> 32) //nolint:gomnd
	b[5] = byte(v >> 40) //nolint:gomnd
	b[6] = byte(v >> 48) //nolint:gomnd
	b[7] = byte(v >> 56) //nolint:gomnd
}

// ReadFloat64 reads 8 bytes and interprets them as a float64 IEEE 754 4 byte bit sequence.
// Panics when len(b) < 8.
func (b LittleEndian) ReadFloat64() float64 {
	bits := b.ReadUint64()

	return math.Float64frombits(bits)
}

// WriteFloat64 writes a float64 IEEE 754 8 byte bit sequence.
// Panics when len(b) < 8.
func (b LittleEndian) WriteFloat64(v float64) {
	bits := math.Float64bits(v)
	b.WriteUint64(bits)
}

// ReadFloat32 reads 4 bytes and interprets them as a float32 IEEE 754 4 byte bit sequence.
// Panics when len(b) < 4.
func (b LittleEndian) ReadFloat32() float32 {
	bits := b.ReadUint32()

	return math.Float32frombits(bits)
}

// WriteFloat32 writes a float32 IEEE 754 4 byte bit sequence.
// Panics when len(b) < 4.
func (b LittleEndian) WriteFloat32(v float32) {
	bits := math.Float32bits(v)
	b.WriteUint32(bits)
}
