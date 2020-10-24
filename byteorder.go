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

// A ByteOrder describes a behavior that can serialize integers and floats as bytes.
type ByteOrder interface {

	// ReadUint16 reads the first 2 bytes. Panics when len(b) < 2.
	ReadUint16() uint16

	// WriteUint16 writes an unsigned 16 bit integer. Panics when len(b) < 2.
	WriteUint16(v uint16)

	// ReadUint24 reads the first 3 bytes. Panics when len(b) < 3.
	ReadUint24() uint32

	// WriteUint24 writes the first 3 bytes. Panics when len(b) < 3.
	WriteUint24(v uint32)

	// ReadUint32 reads the first 4 bytes. Panics when len(b) < 4.
	ReadUint32() uint32

	// WriteUint32 writes the first 4 bytes. Panics when len(b) < 4.
	WriteUint32(v uint32)

	// ReadUint40 reads the first 5 bytes. Panics when len(b) < 5.
	ReadUint40() uint64

	// WriteUint40 writes the first 5 bytes. Panics when len(b) < 5.
	WriteUint40(v uint64)

	// ReadUint48 reads the first 6 bytes. Panics when len(b) < 6.
	ReadUint48() uint64

	// WriteUint48 writes the first 6 bytes. Panics when len(b) < 6.
	WriteUint48(v uint64)

	// ReadUint56 reads the first 7 bytes. Panics when len(b) < 7.
	ReadUint56() uint64

	// WriteUint56 writes the first 7 bytes. Panics when len(b) < 7.
	WriteUint56(v uint64)

	// ReadUint64 reads the first 8 bytes. Panics when len(b) < 8.
	ReadUint64() uint64

	// WriteUint64 writes the first 8 bytes. Panics when len(b) < 8.
	WriteUint64(v uint64)

	// ReadFloat64 reads 8 bytes and interprets them as a float64 IEEE 754 4 byte bit sequence.
	// Panics when len(b) < 8.
	ReadFloat64() float64

	// WriteFloat64 writes a float64 IEEE 754 8 byte bit sequence.
	// Panics when len(b) < 8.
	WriteFloat64(v float64)

	// ReadFloat32 reads 4 bytes and interprets them as a float32 IEEE 754 4 byte bit sequence.
	// Panics when len(b) < 4.
	ReadFloat32() float32

	// WriteFloat32 writes a float32 IEEE 754 4 byte bit sequence.
	// Panics when len(b) < 4.
	WriteFloat32(v float32)
}
