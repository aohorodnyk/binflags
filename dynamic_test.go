package binflags_test

import (
	"fmt"

	"github.com/aohorodnyk/binflags"
)

var (
	_ binflags.Flags = &binflags.Dynamic[int]{}
	_ binflags.Flags = &binflags.Dynamic[uint]{}
	_ binflags.Flags = &binflags.Dynamic[uint16]{}
	_ binflags.Flags = &binflags.Dynamic[int16]{}
	_ binflags.Flags = &binflags.Dynamic[byte]{}
)

func ExampleDynamic() {
	flags := make(binflags.Dynamic[uint8], 3)
	fmt.Println(flags)
	fmt.Println(flags.Set(436235346))
	fmt.Println(flags.Set(0))
	fmt.Println(flags.Set(52))
	fmt.Println(flags.Set(3462363))
	fmt.Println(flags.Set(9874563524235))
	fmt.Println(flags)
	fmt.Println(flags.IsSet(0))
	fmt.Println(flags.IsSet(52))
	fmt.Println(flags.IsSet(3462363))
	fmt.Println(flags.IsSet(9874563524235))
	fmt.Println(flags.IsSet(436235346))
	fmt.Println(flags.IsSet(1))
	fmt.Println(flags.IsSet(436235345))
	fmt.Println(flags.IsSet(436235347))
	fmt.Println(flags.Unset(0))
	fmt.Println(flags.Unset(9874563524235))
	fmt.Println(flags.Unset(2523))
	fmt.Println(flags)

	// Output:
	// map[]
	// true
	// true
	// true
	// true
	// true
	// map[0:1 6:16 432795:8 54529418:4 1234320440529:8]
	// true
	// true
	// true
	// true
	// true
	// false
	// false
	// false
	// true
	// true
	// true
	// map[6:16 432795:8 54529418:4]
}

func ExampleDynamic_IsSet() {
	flags := binflags.Dynamic[uint8]{432795: 8, 54529418: 4, 1234320440529: 8}
	fmt.Println(flags.IsSet(0))
	fmt.Println(flags.IsSet(342))
	fmt.Println(flags.IsSet(3462363))
	fmt.Println(flags.IsSet(9874563524235))
	fmt.Println(flags.IsSet(9874563524234))
	fmt.Println(flags.IsSet(9874563524236))

	// Output:
	// false
	// false
	// true
	// true
	// false
	// false
}

func ExampleDynamic_Unset() {
	flags := binflags.Dynamic[uint8]{432795: 8, 54529418: 4, 1234320440529: 8}
	fmt.Println(flags)
	fmt.Println(flags.Unset(0))
	fmt.Println(flags.Unset(342))
	fmt.Println(flags.Unset(3462363))
	fmt.Println(flags.Unset(9874563524235))
	fmt.Println(flags.Unset(9874563524234))
	fmt.Println(flags.Unset(9874563524236))
	fmt.Println(flags)

	// Output:
	// map[432795:8 54529418:4 1234320440529:8]
	// true
	// true
	// true
	// true
	// true
	// true
	// map[54529418:4]
}

func ExampleDynamic_Set() {
	flags := binflags.Dynamic[uint8]{}
	fmt.Println(flags.Set(234))
	fmt.Println(flags.Set(65234))
	fmt.Println(flags.Set(123))
	fmt.Println(flags.Set(2))
	fmt.Println(flags.Set(1))
	fmt.Println(flags.Set(16))

	fmt.Println(flags)

	// Output:
	// true
	// true
	// true
	// true
	// true
	// true
	// map[0:6 2:1 15:8 29:4 8154:4]
}
