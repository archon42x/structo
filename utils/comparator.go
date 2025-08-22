package utils

type Comparator[T comparable] func(x, y T) int
