package binflags_test

import (
	"fmt"

	"github.com/aohorodnyk/binflags"
)

var (
	_ binflags.Flags = binflags.NewSync(&binflags.Dynamic[int]{})
	_ binflags.Flags = binflags.NewSync(&binflags.Dynamic[uint]{})
	_ binflags.Flags = binflags.NewSync(&binflags.Dynamic[uint16]{})
	_ binflags.Flags = binflags.NewSync(&binflags.Dynamic[int16]{})
	_ binflags.Flags = binflags.NewSync(&binflags.Dynamic[byte]{})
)

func ExampleSync() {
	flags := binflags.NewSync(&binflags.Dynamic[uint8]{})
	fmt.Println(flags.Set(436235346))
	fmt.Println(flags.Set(0))
	fmt.Println(flags.Set(52))
	fmt.Println(flags.Set(3462363))
	fmt.Println(flags.Set(9874563524235))
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

	// Output:
	// true
	// true
	// true
	// true
	// true
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
}
