package yslice

import "mkuznets.com/go/sps/internal/ytils/yerr"

func Map[E any, R any](slice []E, mapper func(value E) R) (mapped []R) {
	for _, el := range slice {
		mapped = append(mapped, mapper(el))
	}
	return mapped
}

func Unique[R comparable](slice []R) []R {
	unique := make([]R, 0)
	visited := map[R]bool{}

	for _, value := range slice {
		if exists := visited[value]; !exists {
			unique = append(unique, value)
			visited[value] = true
		}
	}
	return unique
}

func UniqueMap[E any, R comparable](slice []E, mapper func(value E) R) []R {
	return Unique(Map(slice, mapper))
}

func MapByKey[T any, R comparable](slice []T, key func(value T) R) (mapped map[R]T) {
	mapped = make(map[R]T)
	for _, el := range slice {
		mapped[key(el)] = el
	}
	return mapped
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
