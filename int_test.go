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

package byteorder_test

import (
	"testing"

	. "github.com/worldiety/byteorder"
)

//nolint:funlen
func TestConst(t *testing.T) {
	if MaxInt8 != 127 {
		fail(t)
	}

	if MinInt8 != -128 {
		fail(t)
	}

	if MaxInt16 != 32767 {
		fail(t)
	}

	if MinInt16 != -32768 {
		fail(t)
	}

	if MaxInt24 != 8388607 {
		fail(t)
	}

	if MinInt24 != -8388608 {
		fail(t)
	}

	if MaxInt32 != 2147483647 {
		fail(t)
	}

	if MinInt32 != -2147483648 {
		fail(t)
	}

	if MaxInt40 != 549755813887 {
		fail(t)
	}

	if MinInt40 != -549755813888 {
		fail(t)
	}

	if MaxInt48 != 140737488355327 {
		fail(t)
	}

	if MinInt48 != -140737488355328 {
		fail(t)
	}

	if MaxInt56 != 36028797018963967 {
		fail(t)
	}

	if MinInt56 != -36028797018963968 {
		fail(t)
	}

	if MaxInt64 != 9223372036854775807 {
		fail(t)
	}

	if MinInt64 != -9223372036854775808 {
		fail(t)
	}

	if MaxUint8 != 255 {
		fail(t)
	}

	if MaxUint16 != 65535 {
		fail(t)
	}

	if MaxUint24 != 16777215 {
		fail(t)
	}

	if MaxUint32 != 4294967295 {
		fail(t)
	}

	if MaxUint40 != 1099511627775 {
		fail(t)
	}

	if MaxUint48 != 281474976710655 {
		fail(t)
	}

	if MaxUint56 != 72057594037927935 {
		fail(t)
	}

	if MaxUint64 != 18446744073709551615 {
		fail(t)
	}

	a := MinInt8
	if int32(a) != -128 {
		fail(t)
	}
}

func fail(t *testing.T) {
	t.Helper()
	t.Fatalf("expected other const")
}
