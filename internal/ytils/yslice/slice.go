package yslice

import "mkuznets.com/go/sps/internal/ytils/yerr"

func Map[T any, R any](slice []T, mapper func(value T) R) (mapped []R) {
	for _, el := range slice {
		mapped = append(mapped, mapper(el))
	}
	return mapped
}

func Unique[T comparable](slice []T) []T {
	unique := make([]T, 0)
	visited := map[T]bool{}

	for _, value := range slice {
		if exists := visited[value]; !exists {
			unique = append(unique, value)
			visited[value] = true
		}
	}
	return unique
}

func EnsureOneE[T any](s []T, err error) (T, error) {
	var zero T

	if err != nil {
		return zero, err
	}

	if len(s) < 1 {
		return zero, yerr.NotFound("not found")
	}
	return s[0], nil
}
