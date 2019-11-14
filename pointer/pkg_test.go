package pointer_test

import (
	"errors"
	"testing"
	"time"

	. "go.octolab.org/pointer"
)

func TestBool(t *testing.T) {
	var x bool
	if *ToBool(x) != x {
		t.Errorf("*ToBool(%v)", x)
	}
	if ToBoolOrNil(x) != nil {
		t.Errorf("ToBoolOrNil(%v)", x)
	}
	if ValueOfBool(nil) != x {
		t.Errorf("ValueOfBool(%v)", nil)
	}

	x = true
	if *ToBool(x) != x {
		t.Errorf("*ToBool(%v)", x)
	}
	if *ToBoolOrNil(x) != x {
		t.Errorf("*ToBoolOrNil(%v)", x)
	}
	if ValueOfBool(&x) != x {
		t.Errorf("ValueOfBool(%v)", &x)
	}
}

func TestByte(t *testing.T) {
	var x byte
	if *ToByte(x) != x {
		t.Errorf("*ToByte(%v)", x)
	}
	if ToByteOrNil(x) != nil {
		t.Errorf("ToByteOrNil(%v)", x)
	}
	if ValueOfByte(nil) != x {
		t.Errorf("ValueOfByte(%v)", nil)
	}

	x = 42
	if *ToByte(x) != x {
		t.Errorf("*ToByte(%v)", x)
	}
	if *ToByteOrNil(x) != x {
		t.Errorf("*ToByteOrNil(%v)", x)
	}
	if ValueOfByte(&x) != x {
		t.Errorf("ValueOfByte(%v)", &x)
	}
}

func TestComplex64(t *testing.T) {
	var x complex64
	if *ToComplex64(x) != x {
		t.Errorf("*ToComplex64(%v)", x)
	}
	if ToComplex64OrNil(x) != nil {
		t.Errorf("ToComplex64OrNil(%v)", x)
	}
	if ValueOfComplex64(nil) != x {
		t.Errorf("ValueOfComplex64(%v)", nil)
	}

	x = 42
	if *ToComplex64(x) != x {
		t.Errorf("*ToComplex64(%v)", x)
	}
	if *ToComplex64OrNil(x) != x {
		t.Errorf("*ToComplex64OrNil(%v)", x)
	}
	if ValueOfComplex64(&x) != x {
		t.Errorf("ValueOfComplex64(%v)", &x)
	}
}

func TestComplex128(t *testing.T) {
	var x complex128
	if *ToComplex128(x) != x {
		t.Errorf("*ToComplex128(%v)", x)
	}
	if ToComplex128OrNil(x) != nil {
		t.Errorf("ToComplex128OrNil(%v)", x)
	}
	if ValueOfComplex128(nil) != x {
		t.Errorf("ValueOfComplex128(%v)", nil)
	}

	x = 42
	if *ToComplex128(x) != x {
		t.Errorf("*ToComplex128(%v)", x)
	}
	if *ToComplex128OrNil(x) != x {
		t.Errorf("*ToComplex128OrNil(%v)", x)
	}
	if ValueOfComplex128(&x) != x {
		t.Errorf("ValueOfComplex128(%v)", &x)
	}
}

func TestError(t *testing.T) {
	var x error
	if *ToError(x) != x {
		t.Errorf("*ToError(%v)", x)
	}
	if ToErrorOrNil(x) != nil {
		t.Errorf("ToErrorOrNil(%v)", x)
	}
	if ValueOfError(nil) != x {
		t.Errorf("ValueOfError(%v)", nil)
	}

	x = errors.New("error")
	if *ToError(x) != x {
		t.Errorf("*ToError(%v)", x)
	}
	if *ToErrorOrNil(x) != x {
		t.Errorf("*ToErrorOrNil(%v)", x)
	}
	if ValueOfError(&x) != x {
		t.Errorf("ValueOfError(%v)", &x)
	}
}

func TestFloat32(t *testing.T) {
	var x float32
	if *ToFloat32(x) != x {
		t.Errorf("*ToFloat32(%v)", x)
	}
	if ToFloat32OrNil(x) != nil {
		t.Errorf("ToFloat32OrNil(%v)", x)
	}
	if ValueOfFloat32(nil) != x {
		t.Errorf("ValueOfFloat32(%v)", nil)
	}

	x = 42
	if *ToFloat32(x) != x {
		t.Errorf("*ToFloat32(%v)", x)
	}
	if *ToFloat32OrNil(x) != x {
		t.Errorf("*ToFloat32OrNil(%v)", x)
	}
	if ValueOfFloat32(&x) != x {
		t.Errorf("ValueOfFloat32(%v)", &x)
	}
}

func TestFloat64(t *testing.T) {
	var x float64
	if *ToFloat64(x) != x {
		t.Errorf("*ToFloat64(%v)", x)
	}
	if ToFloat64OrNil(x) != nil {
		t.Errorf("ToFloat64OrNil(%v)", x)
	}
	if ValueOfFloat64(nil) != x {
		t.Errorf("ValueOfFloat64(%v)", nil)
	}

	x = 42
	if *ToFloat64(x) != x {
		t.Errorf("*ToFloat64(%v)", x)
	}
	if *ToFloat64OrNil(x) != x {
		t.Errorf("*ToFloat64OrNil(%v)", x)
	}
	if ValueOfFloat64(&x) != x {
		t.Errorf("ValueOfFloat64(%v)", &x)
	}
}

func TestInt(t *testing.T) {
	var x int
	if *ToInt(x) != x {
		t.Errorf("*ToInt(%v)", x)
	}
	if ToIntOrNil(x) != nil {
		t.Errorf("ToIntOrNil(%v)", x)
	}
	if ValueOfInt(nil) != x {
		t.Errorf("ValueOfInt(%v)", nil)
	}

	x = 42
	if *ToInt(x) != x {
		t.Errorf("*ToInt(%v)", x)
	}
	if *ToIntOrNil(x) != x {
		t.Errorf("*ToIntOrNil(%v)", x)
	}
	if ValueOfInt(&x) != x {
		t.Errorf("ValueOfInt(%v)", &x)
	}
}

func TestInt8(t *testing.T) {
	var x int8
	if *ToInt8(x) != x {
		t.Errorf("*ToInt8(%v)", x)
	}
	if ToInt8OrNil(x) != nil {
		t.Errorf("ToInt8OrNil(%v)", x)
	}
	if ValueOfInt8(nil) != x {
		t.Errorf("ValueOfInt8(%v)", nil)
	}

	x = 42
	if *ToInt8(x) != x {
		t.Errorf("*ToInt8(%v)", x)
	}
	if *ToInt8OrNil(x) != x {
		t.Errorf("*ToInt8OrNil(%v)", x)
	}
	if ValueOfInt8(&x) != x {
		t.Errorf("ValueOfInt8(%v)", &x)
	}
}

func TestInt16(t *testing.T) {
	var x int16
	if *ToInt16(x) != x {
		t.Errorf("*ToInt16(%v)", x)
	}
	if ToInt16OrNil(x) != nil {
		t.Errorf("ToInt16OrNil(%v)", x)
	}
	if ValueOfInt16(nil) != x {
		t.Errorf("ValueOfInt16(%v)", nil)
	}

	x = 42
	if *ToInt16(x) != x {
		t.Errorf("*ToInt16(%v)", x)
	}
	if *ToInt16OrNil(x) != x {
		t.Errorf("*ToInt16OrNil(%v)", x)
	}
	if ValueOfInt16(&x) != x {
		t.Errorf("ValueOfInt16(%v)", &x)
	}
}

func TestInt32(t *testing.T) {
	var x int32
	if *ToInt32(x) != x {
		t.Errorf("*ToInt32(%v)", x)
	}
	if ToInt32OrNil(x) != nil {
		t.Errorf("ToInt32OrNil(%v)", x)
	}
	if ValueOfInt32(nil) != x {
		t.Errorf("ValueOfInt32(%v)", nil)
	}

	x = 42
	if *ToInt32(x) != x {
		t.Errorf("*ToInt32(%v)", x)
	}
	if *ToInt32OrNil(x) != x {
		t.Errorf("*ToInt32OrNil(%v)", x)
	}
	if ValueOfInt32(&x) != x {
		t.Errorf("ValueOfInt32(%v)", &x)
	}
}

func TestInt64(t *testing.T) {
	var x int64
	if *ToInt64(x) != x {
		t.Errorf("*ToInt64(%v)", x)
	}
	if ToInt64OrNil(x) != nil {
		t.Errorf("ToInt64OrNil(%v)", x)
	}
	if ValueOfInt64(nil) != x {
		t.Errorf("ValueOfInt64(%v)", nil)
	}

	x = 42
	if *ToInt64(x) != x {
		t.Errorf("*ToInt64(%v)", x)
	}
	if *ToInt64OrNil(x) != x {
		t.Errorf("*ToInt64OrNil(%v)", x)
	}
	if ValueOfInt64(&x) != x {
		t.Errorf("ValueOfInt64(%v)", &x)
	}
}

func TestRune(t *testing.T) {
	var x rune
	if *ToRune(x) != x {
		t.Errorf("*ToRune(%v)", x)
	}
	if ToRuneOrNil(x) != nil {
		t.Errorf("ToRuneOrNil(%v)", x)
	}
	if ValueOfRune(nil) != x {
		t.Errorf("ValueOfRune(%v)", nil)
	}

	x = 'x'
	if *ToRune(x) != x {
		t.Errorf("*ToRune(%v)", x)
	}
	if *ToRuneOrNil(x) != x {
		t.Errorf("*ToRuneOrNil(%v)", x)
	}
	if ValueOfRune(&x) != x {
		t.Errorf("ValueOfRune(%v)", &x)
	}
}

func TestString(t *testing.T) {
	var x string
	if *ToString(x) != x {
		t.Errorf("*ToString(%v)", x)
	}
	if ToStringOrNil(x) != nil {
		t.Errorf("ToStringOrNil(%v)", x)
	}
	if ValueOfString(nil) != x {
		t.Errorf("ValueOfString(%v)", nil)
	}

	x = "x"
	if *ToString(x) != x {
		t.Errorf("*ToString(%v)", x)
	}
	if *ToStringOrNil(x) != x {
		t.Errorf("*ToStringOrNil(%v)", x)
	}
	if ValueOfString(&x) != x {
		t.Errorf("ValueOfString(%v)", &x)
	}
}

func TestUint(t *testing.T) {
	var x uint
	if *ToUint(x) != x {
		t.Errorf("*ToUint(%v)", x)
	}
	if ToUintOrNil(x) != nil {
		t.Errorf("ToUintOrNil(%v)", x)
	}
	if ValueOfUint(nil) != x {
		t.Errorf("ValueOfUint(%v)", nil)
	}

	x = 42
	if *ToUint(x) != x {
		t.Errorf("*ToUint(%v)", x)
	}
	if *ToUintOrNil(x) != x {
		t.Errorf("*ToUintOrNil(%v)", x)
	}
	if ValueOfUint(&x) != x {
		t.Errorf("ValueOfUint(%v)", &x)
	}
}

func TestUint8(t *testing.T) {
	var x uint8
	if *ToUint8(x) != x {
		t.Errorf("*ToUint8(%v)", x)
	}
	if ToUint8OrNil(x) != nil {
		t.Errorf("ToUint8OrNil(%v)", x)
	}
	if ValueOfUint8(nil) != x {
		t.Errorf("ValueOfUint8(%v)", nil)
	}

	x = 42
	if *ToUint8(x) != x {
		t.Errorf("*ToUint8(%v)", x)
	}
	if *ToUint8OrNil(x) != x {
		t.Errorf("*ToUint8OrNil(%v)", x)
	}
	if ValueOfUint8(&x) != x {
		t.Errorf("ValueOfUint8(%v)", &x)
	}
}

func TestUint16(t *testing.T) {
	var x uint16
	if *ToUint16(x) != x {
		t.Errorf("*ToUint16(%v)", x)
	}
	if ToUint16OrNil(x) != nil {
		t.Errorf("ToUint16OrNil(%v)", x)
	}
	if ValueOfUint16(nil) != x {
		t.Errorf("ValueOfUint16(%v)", nil)
	}

	x = 42
	if *ToUint16(x) != x {
		t.Errorf("*ToUint16(%v)", x)
	}
	if *ToUint16OrNil(x) != x {
		t.Errorf("*ToUint16OrNil(%v)", x)
	}
	if ValueOfUint16(&x) != x {
		t.Errorf("ValueOfUint16(%v)", &x)
	}
}

func TestUint32(t *testing.T) {
	var x uint32
	if *ToUint32(x) != x {
		t.Errorf("*ToUint32(%v)", x)
	}
	if ToUint32OrNil(x) != nil {
		t.Errorf("ToUint32OrNil(%v)", x)
	}
	if ValueOfUint32(nil) != x {
		t.Errorf("ValueOfUint32(%v)", nil)
	}

	x = 42
	if *ToUint32(x) != x {
		t.Errorf("*ToUint32(%v)", x)
	}
	if *ToUint32OrNil(x) != x {
		t.Errorf("*ToUint32OrNil(%v)", x)
	}
	if ValueOfUint32(&x) != x {
		t.Errorf("ValueOfUint32(%v)", &x)
	}
}

func TestUint64(t *testing.T) {
	var x uint64
	if *ToUint64(x) != x {
		t.Errorf("*ToUint64(%v)", x)
	}
	if ToUint64OrNil(x) != nil {
		t.Errorf("ToUint64OrNil(%v)", x)
	}
	if ValueOfUint64(nil) != x {
		t.Errorf("ValueOfUint64(%v)", nil)
	}

	x = 42
	if *ToUint64(x) != x {
		t.Errorf("*ToUint64(%v)", x)
	}
	if *ToUint64OrNil(x) != x {
		t.Errorf("*ToUint64OrNil(%v)", x)
	}
	if ValueOfUint64(&x) != x {
		t.Errorf("ValueOfUint64(%v)", &x)
	}
}

func TestUintptr(t *testing.T) {
	var x uintptr
	if *ToUintptr(x) != x {
		t.Errorf("*ToUintptr(%v)", x)
	}
	if ToUintptrOrNil(x) != nil {
		t.Errorf("ToUintptrOrNil(%v)", x)
	}
	if ValueOfUintptr(nil) != x {
		t.Errorf("ValueOfUintptr(%v)", nil)
	}

	x = 42
	if *ToUintptr(x) != x {
		t.Errorf("*ToUintptr(%v)", x)
	}
	if *ToUintptrOrNil(x) != x {
		t.Errorf("*ToUintptrOrNil(%v)", x)
	}
	if ValueOfUintptr(&x) != x {
		t.Errorf("ValueOfUintptr(%v)", &x)
	}
}

func TestDuration(t *testing.T) {
	var x time.Duration
	if *ToDuration(x) != x {
		t.Errorf("*ToDuration(%v)", x)
	}
	if ToDurationOrNil(x) != nil {
		t.Errorf("ToDurationOrNil(%v)", x)
	}
	if ValueOfDuration(nil) != x {
		t.Errorf("ValueOfDuration(%v)", nil)
	}

	x = time.Second
	if *ToDuration(x) != x {
		t.Errorf("*ToDuration(%v)", x)
	}
	if *ToDurationOrNil(x) != x {
		t.Errorf("*ToDurationOrNil(%v)", x)
	}
	if ValueOfDuration(&x) != x {
		t.Errorf("ValueOfDuration(%v)", &x)
	}
}

func TestTime(t *testing.T) {
	var x time.Time
	if *ToTime(x) != x {
		t.Errorf("*ToTime(%v)", x)
	}
	if ToTimeOrNil(x) != nil {
		t.Errorf("ToTimeOrNil(%v)", x)
	}
	if ValueOfTime(nil) != x {
		t.Errorf("ValueOfTime(%v)", nil)
	}

	x = time.Date(2014, 6, 25, 12, 24, 40, 0, time.UTC)
	if *ToTime(x) != x {
		t.Errorf("*ToTime(%v)", x)
	}
	if *ToTimeOrNil(x) != x {
		t.Errorf("*ToTimeOrNil(%v)", x)
	}
	if ValueOfTime(&x) != x {
		t.Errorf("ValueOfTime(%v)", &x)
	}
}
