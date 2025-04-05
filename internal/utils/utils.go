package utils

func PtrToVal[T any](ptr *T) T {
	var val T
	if ptr == nil {
		return val
	}
	return *ptr
}

func PtrToSliceVal[T any](ptr []*T) []T {
	var val []T
	if ptr == nil {
		return val
	}
	for _, p := range ptr {
		if p != nil {
			val = append(val, *p)
		}
	}
	return val
}
