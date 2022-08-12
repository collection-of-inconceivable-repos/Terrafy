package util

func MapSlice[S any, D any](slice []S, mapFn func(S) D) []D {
	result := make([]D, len(slice))

	for i, item := range slice {
		result[i] = mapFn(item)
	}

	return result
}

func FilterSlice[T any](slice []T, filterFn func(T) bool) []T {
	result := make([]T, 0, len(slice))

	for _, item := range slice {
		if filterFn(item) {
			result = append(result, item)
		}
	}

	return result
}
