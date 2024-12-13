package higherorder

// Filter
// given a slice and a sort func,
// return a new slice with only elemnts from the original list that meet the funcs true comparison
func Filter[T any](ts []T, f func(T) bool) []T {
	var us []T
	for i := range ts {
		if f(ts[i]) {
			us = append(us, ts[i])
		}
	}
	return us
}
