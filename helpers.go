package must

// Ternary is a helper function that returns r1 if cond is true, otherwise r2.
func Ternary[R any](cond bool, r1 R, r2 R) R {
	if cond {
		return r1
	}

	return r2
}
