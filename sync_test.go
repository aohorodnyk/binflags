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

func ExampleSync_IsSet() {
	flags := binflags.NewSync(&binflags.Dynamic[uint8]{432795: 8, 54529418: 4, 1234320440529: 8})
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

func ExampleSync_Unset() {
	flags := binflags.NewSync(binflags.Dynamic[uint8]{432795: 8, 54529418: 4, 1234320440529: 8})
	fmt.Println(flags.Unset(0))
	fmt.Println(flags.Unset(342))
	fmt.Println(flags.Unset(3462363))
	fmt.Println(flags.Unset(9874563524235))
	fmt.Println(flags.Unset(9874563524234))
	fmt.Println(flags.Unset(9874563524236))

	fmt.Println(flags.IsSet(0))
	fmt.Println(flags.IsSet(342))
	fmt.Println(flags.IsSet(3462363))
	fmt.Println(flags.IsSet(9874563524235))
	fmt.Println(flags.IsSet(9874563524234))
	fmt.Println(flags.IsSet(9874563524236))
	fmt.Println(flags.IsSet(436235346))

	// Output:
	// true
	// true
	// true
	// true
	// true
	// true
	// false
	// false
	// false
	// false
	// false
	// false
	// true
}

func ExampleSync_Set() {
	flags := binflags.NewSync(binflags.Dynamic[uint8]{})
	fmt.Println(flags.Set(234))
	fmt.Println(flags.Set(65234))
	fmt.Println(flags.Set(123))
	fmt.Println(flags.Set(2))
	fmt.Println(flags.Set(1))
	fmt.Println(flags.Set(16))

	fmt.Println(flags.IsSet(233))
	fmt.Println(flags.IsSet(234))
	fmt.Println(flags.IsSet(235))
	fmt.Println(flags.IsSet(122))
	fmt.Println(flags.IsSet(123))
	fmt.Println(flags.IsSet(124))
	fmt.Println(flags.IsSet(15))
	fmt.Println(flags.IsSet(16))
	fmt.Println(flags.IsSet(17))

	// Output:
	// true
	// true
	// true
	// true
	// true
	// true
	// false
	// true
	// false
	// false
	// true
	// false
	// false
	// true
	// false
}
