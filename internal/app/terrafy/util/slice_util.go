package util

import (
	"errors"
	"fmt"

	"golang.org/x/exp/constraints"
)

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

func FindFirstIndex[T constraints.Ordered](slice []T, target T) (int, error) {
	for i, item := range slice {
		if item == target {
			return i, nil
		}
	}
	return -1, errors.New("no matching element found")
}

func MustFindFirstIndex[T constraints.Ordered](slice []T, target T) int {
	index, err := FindFirstIndex(slice, target)
	if err != nil {
		msg := fmt.Sprintf("error finding first index: %s", err)
		panic(msg)
	}

	return index
}
