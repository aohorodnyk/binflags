# Golang bitset library

[![GitHub Workflow Status](https://github.com/aohorodnyk/binflags/actions/workflows/go.yml/badge.svg)](https://github.com/aohorodnyk/binflags/actions/workflows/go.yml)
[![License](https://img.shields.io/github/license/aohorodnyk/binflags)](https://github.com/aohorodnyk/binflags/blob/main/LICENSE)
[![GitHub issues](https://img.shields.io/github/issues/aohorodnyk/binflags)](https://github.com/aohorodnyk/binflags/issues)
[![GitHub issues](https://img.shields.io/github/issues-pr/aohorodnyk/binflags)](https://github.com/aohorodnyk/binflags/pulls)
[![The latest release version](https://img.shields.io/github/v/release/aohorodnyk/binflags)](https://github.com/aohorodnyk/binflags/releases)
[![GoDoc](https://godoc.org/github.com/aohorodnyk/binflags?status.svg)](https://pkg.go.dev/github.com/aohorodnyk/binflags)

- [Golang bitset library](#golang-bitset-library)
  - [Golang 1.18](#golang-118)
  - [Motivation](#motivation)
  - [Examples](#examples)
    - [Dynamic](#dynamic)
    - [Fixed](#fixed)
    - [Sync](#sync)
    - [Examples in unit tests](#examples-in-unit-tests)
  - [Contributing](#contributing)
    - [Branch Name](#branch-name)
    - [Git Hook](#git-hook)

## Golang 1.18

The package was built to use Go 1.18 with generics. If you need older version, see please the version [v0.0.3](https://github.com/aohorodnyk/binflags/releases/tag/v0.0.3).

## Motivation

Creating a robust bitset implementation for various types, including "Big Flags" through the use of a map or array of INTs, is a crucial necessity in many applications. While the implementation process is straightforward, it requires thorough testing and support to ensure reliable functionality. As such, careful attention is necessary during the implementation, testing, and maintenance stages to deliver a high-quality and dependable solution.

After few time of implementation I made a decision to extract the feature into open-source library which should meet number of requirements:

1. :white_check_mark: It should support all INT types
1. :white_check_mark: It should be implemented for "Big Flags" tasks through array and map
1. :white_check_mark: It should be well testes, in ideal world it should be with 100% of coverage

## Examples

### Dynamic

Dynamic type is a `map[uint]T`, where `T` is any integer type. This type is a perfect match to the situation when flags are not sequence and we can assign an any random flag to a user (in range of `uin32 * sizeof(T)`).

```go
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
```

### Fixed

Fixed type uses an `array[T]` to save flags. It's more efficient if more sequent list of flags will be used.

```go
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
```

### Sync

`Sync` type is a decorator on top of `Dynamic` and `Fixed` types. Just cover any of them and use concurrenctly.

This type uses **RWMutex**, because of usually services have much more reads than writes.

```go
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
```

### Examples in unit tests

Full list of examples can be found in `*_test.go` files. This library has coverage ~100%, and I'm gonna keep the level.

And, do not forget to use [GoDoc](https://pkg.go.dev/github.com/aohorodnyk/binflags)!

## Contributing

The library is open source and you can contribute to it.

Before contrbution, make sure that githook is configured for you and all your commits contain the correct issue tag.

### Branch Name

Before you start the contribution, make sure that you are on the correct branch. Branch name should start from the issue number dash and short explanation with spaces replaced by underscores. Example:

- `1-my_feature`
- `2-fix_bug`
- `234-my_important_pr`

### Git Hook

To configure the git hook, you need to simply run the command: `git config core.hooksPath .githooks`

It will configure the git hook to run the `pre-commit` script. Source code of the hook is in `.githooks/prepare-commit-msg`.
