package utils

func ToPointer[T interface{}](s T) *T { return &s }
