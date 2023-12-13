package containers

type Result[T any] struct {
	ok      T
	isEmpty bool
	err     error
}

func (m Result[T]) HasErr() bool {
	return m.err != nil
}

func (m Result[T]) IsEmpty() bool {
	return m.isEmpty
}

func (m Result[T]) Err() error {
	return m.err
}

func (m Result[T]) OK() T {
	return m.ok
}

func OK[T any](v T) Result[T] {
	return Result[T]{ok: v}
}

func Err[T any](err error) Result[T] {
	return Result[T]{err: err}
}

func EmptyResult[T any]() Result[T] {
	return Result[T]{err: nil, isEmpty: true}
}
