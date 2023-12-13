package containers

type Maybe[T any] struct {
	just    T
	nothing bool
}

func (m Maybe[T]) Nothing() bool {
	return m.nothing
}

func (m Maybe[T]) Just() T {
	return m.just
}

func Just[T any](v T) Maybe[T] {
	return Maybe[T]{just: v}
}

func Nothing[T any]() Maybe[T] {
	return Maybe[T]{nothing: true}
}
