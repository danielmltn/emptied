package utilities

import "testing"

func TestIsEmptyHandlesStrings(t *testing.T) {

	stringDefault := IsEmpty("")
	stringNotDefault := IsEmpty("Test")

	if !stringDefault || stringNotDefault {
		t.Error("This should have handled strings that are empty as true and non-empty as false. ")
	}

}

func TestIsEmptyHandlesUintptr(t *testing.T) {

	uintDefault := IsEmpty(uintptr(0))
	uintNotDefault := IsEmpty(uintptr(12))

	if !uintDefault || uintNotDefault {
		t.Error("This should have handled uints that are empty as true and non-empty as false. ")
	}

}

func TestIsEmptyHandlesUint(t *testing.T) {

	uintDefault := IsEmpty(uint(0))
	uintNotDefault := IsEmpty(uint(12))

	if !uintDefault || uintNotDefault {
		t.Error("This should have handled uints that are empty as true and non-empty as false. ")
	}

}

func TestIsEmptyHandlesUint8(t *testing.T) {

	uintDefault := IsEmpty(uint8(0))
	uintNotDefault := IsEmpty(uint8(12))

	if !uintDefault || uintNotDefault {
		t.Error("This should have handled uints that are empty as true and non-empty as false. ")
	}

}

func TestIsEmptyHandlesUint16(t *testing.T) {

	uintDefault := IsEmpty(uint16(0))
	uintNotDefault := IsEmpty(uint16(12))

	if !uintDefault || uintNotDefault {
		t.Error("This should have handled uints that are empty as true and non-empty as false. ")
	}

}

func TestIsEmptyHandlesUint32(t *testing.T) {

	uintDefault := IsEmpty(uint32(0))
	uintNotDefault := IsEmpty(uint32(12))

	if !uintDefault || uintNotDefault {
		t.Error("This should have handled uints that are empty as true and non-empty as false. ")
	}

}

func TestIsEmptyHandlesUint64(t *testing.T) {

	uintDefault := IsEmpty(uint64(0))
	uintNotDefault := IsEmpty(uint64(12))

	if !uintDefault || uintNotDefault {
		t.Error("This should have handled uints that are empty as true and non-empty as false. ")
	}

}

func TestIsEmptyHandlesInt(t *testing.T) {

	intDefault := IsEmpty(0)
	intNotDefault := IsEmpty(12)

	if !intDefault || intNotDefault {
		t.Error("This should have handled ints that are empty as true and non-empty as false. ")
	}

}

func TestIsEmptyHandlesInt8(t *testing.T) {

	intDefault := IsEmpty(int8(0))
	intNotDefault := IsEmpty(int8(12))

	if !intDefault || intNotDefault {
		t.Error("This should have handled ints that are empty as true and non-empty as false. ")
	}

}

func TestIsEmptyHandlesInt16(t *testing.T) {

	intDefault := IsEmpty(int16(0))
	intNotDefault := IsEmpty(int16(12))

	if !intDefault || intNotDefault {
		t.Error("This should have handled ints that are empty as true and non-empty as false. ")
	}

}

func TestIsEmptyHandlesInt32(t *testing.T) {

	intDefault := IsEmpty(int32(0))
	intNotDefault := IsEmpty(int32(12))

	if !intDefault || intNotDefault {
		t.Error("This should have handled ints that are empty as true and non-empty as false. ")
	}

}

func TestIsEmptyHandlesInt64(t *testing.T) {

	intDefault := IsEmpty(int64(0))
	intNotDefault := IsEmpty(int64(12))

	if !intDefault || intNotDefault {
		t.Error("This should have handled ints that are empty as true and non-empty as false. ")
	}

}

func TestIsEmptyHandlesFloat32(t *testing.T) {

	floatDefault := IsEmpty(float32(0))
	floatNotDefault := IsEmpty(float32(.123))

	if !floatDefault || floatNotDefault {
		t.Error("This should have handled ints that are empty as true and non-empty as false. ")
	}

}

func TestIsEmptyHandlesFloat64(t *testing.T) {

	floatDefault := IsEmpty(float64(0))
	floatNotDefault := IsEmpty(float64(.123))

	if !floatDefault || floatNotDefault {
		t.Error("This should have handled ints that are empty as true and non-empty as false. ")
	}

}

func TestIsEmptyHandlesComplex64(t *testing.T) {

	floatDefault := IsEmpty(complex64(0))
	floatNotDefault := IsEmpty(complex64(123))

	if !floatDefault || floatNotDefault {
		t.Error("This should have handled ints that are empty as true and non-empty as false. ")
	}

}

func TestIsEmptyHandlesComplex128(t *testing.T) {

	floatDefault := IsEmpty(complex128(0))
	floatNotDefault := IsEmpty(complex128(123))

	if !floatDefault || floatNotDefault {
		t.Error("This should have handled ints that are empty as true and non-empty as false. ")
	}

}

func TestIsEmptyHandlesNilAsDefaultCase(t *testing.T) {

	defaultValue := IsEmpty(nil)
	if defaultValue {
		t.Error("This should have handled nil. ")
	}

}

func TestIsEmptyHandlesDefaultCaseAsFalse(t *testing.T) {

	defaultValue := IsEmpty(struct {
	}{})
	if defaultValue {
		t.Error("This should have handled non-string and non-ints as a default value returned. ")
	}

}

func TestIsEmptyHandlesPointerToStruct(t *testing.T) {

	defaultValue := IsEmpty(&struct{}{})
	if defaultValue {
		t.Error("This should have handled non-string and non-ints as a default value returned. ")
	}

}
