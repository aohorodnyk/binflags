package binflags

type ErrorOutOfRange string

func (e ErrorOutOfRange) Error() string {
	return string(e)
}

type ErrorFlagsArrayNil string

func (e ErrorFlagsArrayNil) Error() string {
	return string(e)
}

type ErrorFlagsMapNil string

func (e ErrorFlagsMapNil) Error() string {
	return string(e)
}
