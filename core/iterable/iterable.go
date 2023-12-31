package iterable

import (
	"io"
)

// Iterator returns items in a collection with every call to Next().
// The error will be set to io.EOF when the iterator is complete.
type Iterator[T any] interface {
	Next() (T, error)
}

type iterator[T any] struct {
	next func() (T, error)
}

func (it *iterator[T]) Next() (T, error) {
	return it.next()
}

func NewIterator[T any](next func() (T, error)) Iterator[T] {
	return &iterator[T]{next}
}

func From[T any](slice []T) Iterator[T] {
	i := 0
	return NewIterator(func() (T, error) {
		if i < len(slice) {
			item := slice[i]
			i++
			return item, nil
		}
		var undef T
		return undef, io.EOF
	})
}

func Collect[T any](it Iterator[T]) ([]T, error) {
	var items []T
	for {
		item, err := it.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func Concat[T any](iterators ...Iterator[T]) Iterator[T] {
	if len(iterators) == 0 {
		return From([]T{})
	}

	i := 0
	iterator := iterators[i]
	return NewIterator(func() (T, error) {
		for {
			item, err := iterator.Next()
			if err != nil {
				if err == io.EOF {
					i++
					if i < len(iterators) {
						iterator = iterators[i]
						continue
					}
				}
				var undef T
				return undef, err
			}
			return item, nil
		}
	})
}
