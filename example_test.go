package geohash_test

import (
	"fmt"

	"github.com/mmcloughlin/geohash"
)

func ExampleEncode() {
	fmt.Println(geohash.Encode(48.858, 2.294))
	// Output: u09tunq6qp66
}

func ExampleEncodeInt() {
	fmt.Printf("%016x\n", geohash.EncodeInt(48.858, 2.294))
	// Output: d0139d52c6b54c69
}

func ExampleEncodeIntWithPrecision() {
	fmt.Printf("%08x\n", geohash.EncodeIntWithPrecision(48.858, 2.294, 32))
	// Output: d0139d52
}

func ExampleEncodeWithPrecision() {
	fmt.Println(geohash.EncodeWithPrecision(48.858, 2.294, 5))
	// Output: u09tu
}