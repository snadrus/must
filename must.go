package must

import (
	"io"

	"golang.org/x/xerrors"
)

func One[R any](r R, err error) R {
	if err != nil {
		panic(xerrors.Errorf("%w", err))
	}

	return r
}

func Two[R1 any, R2 any](r1 R1, r2 R2, err error) (R1, R2) {
	if err != nil {
		panic(xerrors.Errorf("%w", err))
	}

	return r1, r2
}

func Three[R1 any, R2 any, R3 any](r1 R1, r2 R2, r3 R3, err error) (R1, R2, R3) {
	if err != nil {
		panic(xerrors.Errorf("%w", err))
	}

	return r1, r2, r3
}

// RecoverToErr recovers from a panic and assigns the error to the given pointer.
// Usage:
// func foo() (err error) {
// defer must.RecoverToErr(&err)
func RecoverToErr(retErr *error) {
	if r := recover(); r != nil {
		*retErr = r.(error)
	}
}

// With is a helper function to close a Closer after a function is done.
// Usage:
//
//	must.With(must.One(os.Open("/foo")), func(f *os.File) {
//	  // do something with f
//	})
func With[C io.Closer](c C, f func(C)) {
	defer c.Close()
	f(c)
}
