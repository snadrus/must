// Must simplifies Golang into a qbasic-style (nearly) error handling.
// Consider the following code:
// import . "github.com/snadrus/must"
//
//		func Copyfile(dest, src string) (err error) {
//		  Recover2Err(&err)
//		  d := One(os.Open(dest))
//	   defer d.Close()
//		  s := One(os.Open(src))
//	   defer s.Close()
//		  _ = One(io.Copy(dest, erc))
//	   return nil
//		}
package must

import (
	"fmt"
	"io"

	"golang.org/x/xerrors"
)

func E2p(err error) {
	if err != nil {
		panic(xerrors.Errorf("%w", err))
	}
}

var Ck = E2p

// One returns the first return of a (T, err) case, panicing the error if not nil:
// Usage:  f := must.One(os.Open("file.txt"))
func One[R any](r R, err error) R {
	if err != nil {
		panic(xerrors.Errorf("%w", err))
	}

	return r
}

// OneWrap wraps any errors with a string
// Usage: f := must.OneWrap(os.Open("file.txt"))("opening file")
func OneWrap[T any](v T, err error) func(s string) T {
	return func(s string) T {
		if err != nil {
			panic(xerrors.Errorf("%s: %w", s, err))
		}
		return v
	}
}

func Two[R1 any, R2 any](r1 R1, r2 R2, err error) (R1, R2) {
	if err != nil {
		panic(xerrors.Errorf("%w", err))
	}
	return r1, r2
}

// TwoWrap wraps any errors with a string
// Usage: host, port := must.TwoWrap(net.SplitHostPort(garbageVar))(fmt.Sprint("parsing h&p: %s", garbageVar))
func TwoWrap[R1 any, R2 any](r1 R1, r2 R2, err error) func(s string) (R1, R2) {
	return func(s string) (R1, R2) {
		if err != nil {
			panic(xerrors.Errorf("%s: %w", s, err))
		}
		return r1, r2
	}
}

func Three[R1 any, R2 any, R3 any](r1 R1, r2 R2, r3 R3, err error) (R1, R2, R3) {
	if err != nil {
		panic(xerrors.Errorf("%w", err))
	}

	return r1, r2, r3
}

type Wrapper func(error) error

// RecoverToErr recovers from a panic and assigns the error to the given pointer.
// Usage:
// func foo() (err error) {
// defer must.RecoverToErr(&err)
// Wrappers like Wrap() can be used to annotate or examine errors.
func RecoverToErr(retErr *error, w ...Wrapper) {
	if r := recover(); r != nil {
		if err, ok := r.(error); ok {
			*retErr = err
		}
		*retErr = xerrors.Errorf("%v", r)
	}
	if *retErr != nil {
		for _, w := range w {
			*retErr = w(*retErr)
		}
	}
}

// Wrap adds a message around any error.
// Usage: defer must.RecoverToErr(&err, must.Wrap("in Launch()"))
func Wrap(fmtStr string, args ...any) Wrapper {
	return func(e error) error {
		var newArgs = make([]any, len(args)+1)
		copy(newArgs, args)
		return fmt.Errorf(fmtStr+": %w", append(newArgs, e)...)
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
