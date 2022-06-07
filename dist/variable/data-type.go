package variable

import "fmt"

func NonDecimal() {
	fmt.Println("#Non Decimal Numeric")
	fmt.Println(`
	Go's non-decimal numeric data type
	uint8   ==> 0 - 255
	uint16  ==> 0 - 65535
	uint32  ==> 0 - 4294967295
	uint64  ==> 0 - 18446744073709551615
	uint    ==> same as uint32 or uint64 (depending on value)
	byte    ==> same as uint8
	int8    ==> -128 - 127
	int16   ==> -32768 - 32767
	int32   ==> -2147483648 - 2147483647
	int64   ==> -9223372036854775808 - 9223372036854775807
	int     ==> same as int32 or int64 (depending on value)
	rune    ==> same as int32
	`)
}

func Decimal() {
	fmt.Println("#Decimal")
	fmt.Println(`
	There are two types of decimal numeric data: float32 and float64, the difference between them is the capacity of the decimal value
	`)
}

func BooleanType() {
	var loading bool = true
	fmt.Println("#Boolean")
	fmt.Printf("loading... %t \n", loading)
}