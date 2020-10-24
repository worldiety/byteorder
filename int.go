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

// UintSize is either 32 or 64.
const UintSize = 32 << (^uint(0) >> 32 & 1)

const (
	// MaxInt is either 1<<31 - 1 or 1<<63 - 1.
	MaxInt = 1<<(UintSize-1) - 1

	// MinInt is either -1 << 31 or -1 << 63.
	MinInt = -MaxInt - 1

	// MaxUint is either 1<<32 - 1 or 1<<64 - 1.
	MaxUint = 1<<UintSize - 1

	// MaxInt8 is 127.
	MaxInt8 int8 = 1<<7 - 1
	// MinInt8 is -128.
	MinInt8 int8 = -1 << 7
	// MaxInt16 is 32767.
	MaxInt16 int16 = 1<<15 - 1
	// MinInt16 is -32768.
	MinInt16 int16 = -1 << 15
	// MaxInt24 is 8388607.
	MaxInt24 int32 = 1<<23 - 1
	// MinInt24 is -8388608.
	MinInt24 int32 = -1 << 23
	// MaxInt32 is 2147483647.
	MaxInt32 int32 = 1<<31 - 1
	// MinInt32 is -2147483648.
	MinInt32 int32 = -1 << 31
	// MaxInt40 is 549755813887.
	MaxInt40 int64 = 1<<39 - 1
	// MinInt40 is -549755813888.
	MinInt40 int64 = -1 << 39
	// MaxInt48 is 140737488355327.
	MaxInt48 int64 = 1<<47 - 1
	// MinInt48 is -140737488355328.
	MinInt48 int64 = -1 << 47
	// MaxInt56 is 36028797018963967.
	MaxInt56 int64 = 1<<55 - 1
	// MinInt56 is -36028797018963968.
	MinInt56 int64 = -1 << 55
	// MaxInt64 is 9223372036854775807.
	MaxInt64 int64 = 1<<63 - 1
	// MinInt64 is -9223372036854775808.
	MinInt64 int64 = -1 << 63
	// MaxUint8 is 255.
	MaxUint8 uint8 = 1<<8 - 1
	// MaxUint16 is 65535.
	MaxUint16 uint16 = 1<<16 - 1
	// MaxUint24 is 16777215.
	MaxUint24 uint32 = 1<<24 - 1
	// MaxUint32 is 4294967295.
	MaxUint32 uint32 = 1<<32 - 1
	// MaxUint40 is 1099511627775.
	MaxUint40 uint64 = 1<<40 - 1
	// MaxUint48 is 281474976710655.
	MaxUint48 uint64 = 1<<48 - 1
	// MaxUint56 is 72057594037927935.
	MaxUint56 uint64 = 1<<56 - 1
	// MaxUint64 is 18446744073709551615.
	MaxUint64 uint64 = 1<<64 - 1
)
