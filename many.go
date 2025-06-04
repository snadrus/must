package must

import (
	"golang.org/x/sync/errgroup"
)

// Many starts parallel goroutines that handle errors through panic forwarding.
func Many(fs ...func()) {
	g := errgroup.Group{}
	for a := range fs {
		g.Go(func() (err error) {
			defer RecoverToErr(&err)
			fs[a]()
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		panic(err)
	}
}

// ErrOnly is a helper function to return an error if it is not nil.
// This is good for Many() when you don't need the first argument.
// Usage:
//
//	must.Many(
//	 must.ErrOnly2(db.Exec("INSERT INTO foo VALUES (?)", 42))
//	 must.ErrOnly2(db.Exec("INSERT INTO foo VALUES (?)", 43))
//	)
func ErrOnly2[R any](r R, err error) error {
	if err != nil {
		return err
	}

	return nil
}
