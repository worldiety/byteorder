package byteorder_test

import (
	"fmt"
	"os"
	"testing"

	. "github.com/worldiety/byteorder"
)

//nolint:gochecknoglobals
var src = []byte{0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0xFF}

func TestMain(m *testing.M) {
	rc := m.Run()
	requiredCoverage := 1.0

	if rc == 0 && testing.CoverMode() != "" {
		c := testing.Coverage()
		if c < requiredCoverage {
			fmt.Printf("code coverage to low: %f/%f\n", c, requiredCoverage)
			os.Exit(-1)
		}
	}
}

func assertValues(t *testing.T, tmp []byte) {
	t.Helper()

	for i := 0; i < len(tmp); i++ {
		if tmp[i] != src[i] {
			t.Fatalf("expected %x but got %x", src[i], tmp[i])
		}
	}
}

func Test_Uint16(t *testing.T) {
	tmp := make([]byte, 2)

	LE(tmp).WriteUint16(LE(src).ReadUint16())
	assertValues(t, tmp)

	BE(tmp).WriteUint16(BE(src).ReadUint16())
	assertValues(t, tmp)
}

func Test_Uint24(t *testing.T) {
	tmp := make([]byte, 3)

	LE(tmp).WriteUint24(LE(src).ReadUint24())
	assertValues(t, tmp)

	BE(tmp).WriteUint24(BE(src).ReadUint24())
	assertValues(t, tmp)
}

func Test_Uint32(t *testing.T) {
	tmp := make([]byte, 4)

	LE(tmp).WriteUint32(LE(src).ReadUint32())
	assertValues(t, tmp)

	BE(tmp).WriteUint32(BE(src).ReadUint32())
	assertValues(t, tmp)
}

func Test_Uint40(t *testing.T) {
	tmp := make([]byte, 5)

	LE(tmp).WriteUint40(LE(src).ReadUint40())
	assertValues(t, tmp)

	BE(tmp).WriteUint40(BE(src).ReadUint40())
	assertValues(t, tmp)
}

func Test_Uint48(t *testing.T) {
	tmp := make([]byte, 6)

	LE(tmp).WriteUint48(LE(src).ReadUint48())
	assertValues(t, tmp)

	BE(tmp).WriteUint48(BE(src).ReadUint48())
	assertValues(t, tmp)
}

func Test_Uint56(t *testing.T) {
	tmp := make([]byte, 7)

	LE(tmp).WriteUint56(LE(src).ReadUint56())
	assertValues(t, tmp)

	BE(tmp).WriteUint56(BE(src).ReadUint56())
	assertValues(t, tmp)
}

func Test_Uint64(t *testing.T) {
	tmp := make([]byte, 8)

	LE(tmp).WriteUint64(LE(src).ReadUint64())
	assertValues(t, tmp)

	BE(tmp).WriteUint64(BE(src).ReadUint64())
	assertValues(t, tmp)
}

func Test_Float32(t *testing.T) {
	tmp := make([]byte, 4)

	LE(tmp).WriteFloat32(LE(src).ReadFloat32())
	assertValues(t, tmp)

	BE(tmp).WriteFloat32(BE(src).ReadFloat32())
	assertValues(t, tmp)
}

func Test_Float64(t *testing.T) {
	tmp := make([]byte, 8)

	LE(tmp).WriteFloat64(LE(src).ReadFloat64())
	assertValues(t, tmp)

	BE(tmp).WriteFloat64(BE(src).ReadFloat64())
	assertValues(t, tmp)
}
