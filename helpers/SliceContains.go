package helpers

func Contains(elems []string, v *string) bool {
	for _, s := range elems {
		if v == &s {
			return true
		}
	}
	return false
}
