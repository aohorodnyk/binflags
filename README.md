# Golang bitset library

* [Motivation](#motivation)
* [Quick Architecture Review](#quick-architecture-review)
* [Known limitations](#known-limitations)
* [Examples](#examples)
* [Contributing](#contributing)

[![Codecov](https://codecov.io/gh/aohorodnyk/binflags/branch/main/graph/badge.svg?token=61SCDABJJ0)](https://codecov.io/gh/aohorodnyk/binflags) ![Test](https://github.com/aohorodnyk/binflags/workflows/Test/badge.svg) ![Lint](https://github.com/aohorodnyk/binflags/workflows/Lint/badge.svg)

## Motivation
There are number of needs to implement bitset for various types, and some "Big Flags" implementation through map or array of INTs.
It's pretty simple to implement, but every time it should be well tested, as a result it takes some time to implement, test and support it.
After few time of implementation I made a decision to extract the feature into open-source library which should meet number of requirements:
1. :white_check_mark: It should support all INT types
1. :white_check_mark: It should be implemented for "Big Flags" tasks through array and map
1. :white_check_mark: It should be well testes, in ideal world it should be 100% of coverage

## Quick Architecture Review
This library doesn't have any state. There're number of functions which have some standards in an implementation:
1. Every `has` function returns only `bool` even if a parameter(s) is not valid
1. Every `has` function doesn't change parameters
1. Every `set` function returns `error`, if a parameter(s) is not valid
1. Every `set` function for base types (like `int8`, `uint8`, `int64`, `uint64`) returns copy of flags from a function
1. Every `set` function for an array types (`[]uint64`)/map (`map[uint64]uint64`) changes a flag directly in a argument of the function

## Known limitations
There are known limitations:
1. Go uses int64 as a key for array types, it means that max size of array is `math.MaxInt64` or `9223372036854775807`, then the formula for max flag value in array flags is `9223372036854775807 * sizeof(type)`. `type` is a type for flags, like `int8`, `int16` ... `int64`.

## Examples

### Examples in unit tests
Full list of examples can be found in `*_test.go` files. This library has coverage ~100%, and I'm gonna keep the level.

### Check has flag in base types 
```go
package main

import (
	"fmt"
	"github.com/aohorodnyk/binflags"
)

func main() {
	fmt.Println("=-=-=-=-=-=-=-Int8-=-=-=-=-=-=-=")
	i8 := int8(0)
	fmt.Printf("%d\n", i8) // 0 => 00000000
	i8, _ = binflags.SetFlagInt8(i8, 0, true)
	fmt.Printf("%d\n", i8) // 1 => 00000001
	i8, _ = binflags.SetFlagInt8(i8, 0, true)
	fmt.Printf("%d\n", i8) // 1 => 00000001
	i8, _ = binflags.SetFlagInt8(i8, 5, true)
	fmt.Printf("%d\n", i8) // 33 => 00100001
	i8, _ = binflags.SetFlagInt8(i8, 7, true)
	fmt.Printf("%d\n", i8) // 161 => 10100001
	i8, _ = binflags.SetFlagInt8(i8, 7, false)
	fmt.Printf("%d\n", i8) // 33 => 00100001
	i8, _ = binflags.SetFlagInt8(i8, 7, true)
	fmt.Printf("%d\n", i8) // 161 => 10100001
	// Check first bit in flags, 1 << 0
	fmt.Printf("%d => %t\n", i8, binflags.HasFlagInt8(i8, 0)) // -95 => true
	// Check first bit in flags, 1 << 1
	fmt.Printf("%d => %t\n", i8, binflags.HasFlagInt8(i8, 1)) // -95 => false
	// Check first bit in flags, 1 << 2
	fmt.Printf("%d => %t\n", i8, binflags.HasFlagInt8(i8, 2)) // -95 => false
	// Check first bit in flags, 1 << 3
	fmt.Printf("%d => %t\n", i8, binflags.HasFlagInt8(i8, 3)) // -95 => false
	// Check first bit in flags, 1 << 4
	fmt.Printf("%d => %t\n", i8, binflags.HasFlagInt8(i8, 4)) // -95 => false
	// Check first bit in flags, 1 << 5
	fmt.Printf("%d => %t\n", i8, binflags.HasFlagInt8(i8, 5)) // -95 => true
	// Check first bit in flags, 1 << 6
	fmt.Printf("%d => %t\n", i8, binflags.HasFlagInt8(i8, 6)) // -95 => false
	// Check first bit in flags, -1 << 7
	fmt.Printf("%d => %t\n", i8, binflags.HasFlagInt8(i8, 7)) // -95 => true

	fmt.Println("=-=-=-=-=-=-=-Uint8-=-=-=-=-=-=-=")
	ui8 := uint8(0)
	fmt.Printf("%d\n", ui8) // 0 => 00000000
	ui8, _ = binflags.SetFlagUint8(ui8, 0, true)
	fmt.Printf("%d\n", ui8) // 1 => 00000001
	ui8, _ = binflags.SetFlagUint8(ui8, 0, true)
	fmt.Printf("%d\n", ui8) // 1 => 00000001
	ui8, _ = binflags.SetFlagUint8(ui8, 5, true)
	fmt.Printf("%d\n", ui8) // 33 => 00100001
	ui8, _ = binflags.SetFlagUint8(ui8, 7, true)
	fmt.Printf("%d\n", ui8) // -95 => 10100001
	ui8, _ = binflags.SetFlagUint8(ui8, 7, false)
	fmt.Printf("%d\n", ui8) // 33 => 00100001
	ui8, _ = binflags.SetFlagUint8(ui8, 7, true)
	fmt.Printf("%d\n", ui8) // -95 => 10100001
	// Check first bit in flags, 1 << 0
	fmt.Printf("%d => %t\n", ui8, binflags.HasFlagUint8(ui8, 0)) // 161 => true
	// Check first bit in flags, 1 << 1
	fmt.Printf("%d => %t\n", ui8, binflags.HasFlagUint8(ui8, 1)) // 161 => false
	// Check first bit in flags, 1 << 2
	fmt.Printf("%d => %t\n", ui8, binflags.HasFlagUint8(ui8, 2)) // 161 => false
	// Check first bit in flags, 1 << 3
	fmt.Printf("%d => %t\n", ui8, binflags.HasFlagUint8(ui8, 3)) // 161 => false
	// Check first bit in flags, 1 << 4
	fmt.Printf("%d => %t\n", ui8, binflags.HasFlagUint8(ui8, 4)) // 161 => false
	// Check first bit in flags, 1 << 5
	fmt.Printf("%d => %t\n", ui8, binflags.HasFlagUint8(ui8, 5)) // 161 => true
	// Check first bit in flags, 1 << 6
	fmt.Printf("%d => %t\n", ui8, binflags.HasFlagUint8(ui8, 6)) // 161 => false
	// Check first bit in flags, -1 << 7
	fmt.Printf("%d => %t\n", ui8, binflags.HasFlagUint8(ui8, 7)) // 161 => true
}
```
#### Output
```
=-=-=-=-=-=-=-Int8-=-=-=-=-=-=-=
0
1
1
33
-95
33
-95
-95 => true
-95 => false
-95 => false
-95 => false
-95 => false
-95 => true
-95 => false
-95 => true
=-=-=-=-=-=-=-Uint8-=-=-=-=-=-=-=
0
1
1
33
161
33
161
161 => true
161 => false
161 => false
161 => false
161 => false
161 => true
161 => false
161 => true
```

### Check has flag in an array type
```go
package main

import (
	"fmt"
	"github.com/aohorodnyk/binflags"
)

func main() {
	var err error
	fmt.Println("=-=-=-=-=-=-=-Int8-=-=-=-=-=-=-=")
	i8a := []int8{-128, -2, 0, 25}
	fmt.Println(i8a) // [-128 -2 0 25]
	binflags.SetFlagArrayInt8(i8a, 19, true)
	fmt.Println(i8a) // [-128 -2 8 25]
	binflags.SetFlagArrayInt8(i8a, 27, false)
	fmt.Println(i8a) // [-128 -2 8 17]
	err = binflags.SetFlagArrayInt8(i8a, 32, true)
	fmt.Println(err) // flag is out of range
	fmt.Println(i8a) // [-128 -2 8 17]
	err = binflags.SetFlagArrayInt8(nil, 0, true)
	fmt.Println(err) // flags array is nil
	// Check first bit in flags, 1 << 0
	fmt.Printf("%t\n", binflags.HasFlagArrayInt8(i8a, 0)) // false
	// Check first bit in flags, 1 << 1
	fmt.Printf("%t\n", binflags.HasFlagArrayInt8(i8a, 1)) // false
	// Check first bit in flags, 1 << 7
	fmt.Printf("%t\n", binflags.HasFlagArrayInt8(i8a, 7)) // true
	// Check first bit in flags, 1 << 8
	fmt.Printf("%t\n", binflags.HasFlagArrayInt8(i8a, 8)) // false
	// Check first bit in flags, 1 << 15
	fmt.Printf("%t\n", binflags.HasFlagArrayInt8(i8a, 15)) // true
	// Check first bit in flags, 1 << 32
	fmt.Printf("%t\n", binflags.HasFlagArrayInt8(i8a, 28)) // true
	// Check first bit in flags, 1 << 32
	fmt.Printf("%t\n", binflags.HasFlagArrayInt8(i8a, 32)) // false
	// Check first bit in flags, 1 << 1453456
	fmt.Printf("%t\n", binflags.HasFlagArrayInt8(i8a, 1453456)) // false
}
```

#### Output
```
=-=-=-=-=-=-=-Int8-=-=-=-=-=-=-=
[-128 -2 0 25]
[-128 -2 8 25]
[-128 -2 8 17]
flag is out of range
[-128 -2 8 17]
flags array is nil
false
false
true
false
true
true
false
false
```

### Check has flag in a map type
```go
package main

import (
	"fmt"
	"github.com/aohorodnyk/binflags"
)

func main() {
	var err error
	fmt.Println("=-=-=-=-=-=-=-Int8-=-=-=-=-=-=-=")
	i8m := map[uint64]int8{}
	fmt.Println(i8m) // map[]
	binflags.SetFlagMapInt8(i8m, 0, true)
	fmt.Println(i8m) // map[0:1]
	binflags.SetFlagMapInt8(i8m, 7, true)
	fmt.Println(i8m) // map[0:-127]
	binflags.SetFlagMapInt8(i8m, 76, true)
	fmt.Println(i8m) // map[0:-127 9:16]
	binflags.SetFlagMapInt8(i8m, 35743653548435, true)
	fmt.Println(i8m) // map[0:-127 9:16 4467956693554:8]
	binflags.SetFlagMapInt8(i8m, 35743653548436, true)
	fmt.Println(i8m) // map[0:-127 9:16 4467956693554:24]
	binflags.SetFlagMapInt8(i8m, 76, false)
	fmt.Println(i8m) // map[0:-127 4467956693554:24]
	err = binflags.SetFlagMapInt8(nil, 76, false)
	fmt.Println(err) // flags map is nil
	fmt.Printf("%t\n", binflags.HasFlagMapInt8(i8m, 0))              // true
	fmt.Printf("%t\n", binflags.HasFlagMapInt8(i8m, 5))              // false
	fmt.Printf("%t\n", binflags.HasFlagMapInt8(i8m, 7))              // true
	fmt.Printf("%t\n", binflags.HasFlagMapInt8(i8m, 76))             // false
	fmt.Printf("%t\n", binflags.HasFlagMapInt8(i8m, 2523565235))     // false
	fmt.Printf("%t\n", binflags.HasFlagMapInt8(i8m, 35743653548436)) // true
	fmt.Printf("%t\n", binflags.HasFlagMapInt8(i8m, 35743653548437)) // false
}
```

#### Output
```
=-=-=-=-=-=-=-Int8-=-=-=-=-=-=-=
map[]
map[0:1]
map[0:-127]
map[0:-127 9:16]
map[0:-127 9:16 4467956693554:8]
map[0:-127 9:16 4467956693554:24]
map[0:-127 4467956693554:24]
flags map is nil
true
false
true
false
false
true
false
```

## Contributing
All contributions have to follow the [CONTRIBUTING.md document](https://github.com/aohorodnyk/uid/blob/main/CONTRIBUTING.md)
If you have any questions/issues/feature requests do not hesitate to create a ticket.
