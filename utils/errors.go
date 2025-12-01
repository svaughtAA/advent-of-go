package utils

func Assert(cond bool, msg string) {
	if !cond {
		panic(msg)
	}
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func Unpack[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}
