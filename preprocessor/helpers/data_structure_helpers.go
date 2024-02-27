package helpers

func Contains[V comparable](s []V, e V) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}