package binflags_test

import (
	"fmt"

	"github.com/aohorodnyk/binflags"
)

var (
	_ binflags.Flags = &binflags.Fixed[int]{}
	_ binflags.Flags = &binflags.Fixed[uint]{}
	_ binflags.Flags = &binflags.Fixed[uint16]{}
	_ binflags.Flags = &binflags.Fixed[int16]{}
	_ binflags.Flags = &binflags.Fixed[int8]{}
	_ binflags.Flags = &binflags.Fixed[byte]{}
)

func ExampleFixed() {
	flags := make(binflags.Fixed[uint8], 3)
	fmt.Println(flags)
	fmt.Println(flags.Set(20))
	fmt.Println(flags)
	fmt.Println(flags.Set(23))
	fmt.Println(flags)
	fmt.Println(flags.Set(13))
	fmt.Println(flags)
	fmt.Println(flags.Set(24))
	fmt.Println(flags)
	fmt.Println(flags.IsSet(5))
	fmt.Println(flags.IsSet(13))
	fmt.Println(flags.Unset(24))
	fmt.Println(flags)
	fmt.Println(flags.Unset(7))
	fmt.Println(flags)
	fmt.Println(flags.Unset(13))
	fmt.Println(flags)

	// Output:
	// [0 0 0]
	// true
	// [0 0 16]
	// true
	// [0 0 144]
	// true
	// [0 32 144]
	// false
	// [0 32 144]
	// false
	// true
	// false
	// [0 32 144]
	// true
	// [0 32 144]
	// true
	// [0 0 144]
}

func ExampleFixed_IsSet() {
	flags := binflags.Fixed[uint8]{0, 255, 123, 53, 255}
	fmt.Println(flags.IsSet(0))
	fmt.Println(flags.IsSet(7))
	fmt.Println(flags.IsSet(8))
	fmt.Println(flags.IsSet(15))
	fmt.Println(flags.IsSet(39))
	fmt.Println(flags.IsSet(40))
	fmt.Println(flags.IsSet(12423))

	// Output:
	// false
	// false
	// true
	// true
	// true
	// false
	// false
}

func ExampleFixed_Set() {
	flags := make(binflags.Fixed[uint8], 3)
	fmt.Println(flags)
	fmt.Println(flags.Set(0))
	fmt.Println(flags)
	fmt.Println(flags.Set(14))
	fmt.Println(flags)
	fmt.Println(flags.Set(21))
	fmt.Println(flags)
	fmt.Println(flags.Set(23))
	fmt.Println(flags)
	fmt.Println(flags.Set(24))
	fmt.Println(flags)

	// Output:
	// [0 0 0]
	// true
	// [1 0 0]
	// true
	// [1 64 0]
	// true
	// [1 64 32]
	// true
	// [1 64 160]
	// false
	// [1 64 160]
}

func ExampleFixed_Unset() {
	flags := binflags.Fixed[uint8]{255, 255, 255}
	fmt.Println(flags)
	fmt.Println(flags.Unset(0))
	fmt.Println(flags)
	fmt.Println(flags.Unset(14))
	fmt.Println(flags)
	fmt.Println(flags.Unset(21))
	fmt.Println(flags)
	fmt.Println(flags.Unset(23))
	fmt.Println(flags)
	fmt.Println(flags.Unset(24))
	fmt.Println(flags)

	// Output:
	// [255 255 255]
	// true
	// [254 255 255]
	// true
	// [254 191 255]
	// true
	// [254 191 223]
	// true
	// [254 191 95]
	// false
	// [254 191 95]
}
