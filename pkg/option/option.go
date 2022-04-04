package option

import "fmt"

// Option represents an optional value: every Option is either Some and contains a value, or None, and does not.
type Option[T any] struct {
	v *T
}

func (o Option[T]) String() string {
	if o.v != nil {
		if s, ok := interface{}(o.v).(fmt.Stringer); ok {
			return s.String()
		}

		return fmt.Sprintf("%v", *o.v)
	}

	return "<none>"
}

// Some value of type T.
func Some[T any](v T) Option[T] { return Option[T]{&v} }

// None means no value.
func None[T any]() Option[T] { return Option[T]{} }

// HasSome returns true if the option is a `Some` value.
func (o Option[T]) HasSome() bool { return o.v != nil }

// IsNone returns true if the option is a `None` value.
func (o Option[T]) IsNone() bool { return o.v == nil }

// Expect returns the contained `Some` value,
// or panic if the `o` value equals `None` with a custom panic message provided by `msg`.
func (o Option[T]) Expect(msg string) T {
	if o.v == nil {
		panic(msg)
	}

	return *o.v
}

// Unwrap returns the contained `Some` value, or panic if the `o` value equals `None`.
func (o Option[T]) Unwrap() T {
	return o.Expect("Unwrap on a `None` value")
}

// UnwrapOr returns the contained `Some` value or a provided `defaultValue`.
func (o Option[T]) UnwrapOr(defaultValue T) T {
	if o.v == nil {
		return defaultValue
	}

	return *o.v
}

// UnwrapOrElse returns the contained `Some` value or computes it from a closure `f`.
func (o Option[T]) UnwrapOrElse(f func() T) T {
	if o.v == nil {
		return f()
	}

	return *o.v
}

// UnwrapOrDefault returns the contained `Some` value or a default.
func (o Option[T]) UnwrapOrDefault() (v T) {
	if o.v != nil {
		v = *o.v
	}

	return
}

// OkOr transforms the Option[T] into a (T, error), mapping `Some(v)` to `v` and `None` to `err`.
func (o Option[T]) OkOr(e error) (v T, err error) {
	if o.v != nil {
		v = *o.v
	} else {
		err = e
	}

	return
}

// OkOrElse transforms the Option[T] into a (T, error), mapping `Some(v)` to `v` and `None` to `f()`.
func (o Option[T]) OkOrElse(f func() error) (v T, err error) {
	if o.v != nil {
		v = *o.v
	} else {
		err = f()
	}

	return
}

// Map maps an Option[T] to Option[U] by applying a function `f` to a contained value.
func Map[T, U any](o Option[T], f func(T) U) Option[U] {
	if o.v == nil {
		return None[U]()
	}

	return Some(f(*o.v))
}

// MapOr returns the provided default result (if none), or applies a function to the contained value (if any).
func MapOr[T, U any](o Option[T], defaultValue U, f func(T) U) Option[U] {
	if o.v == nil {
		return Some(defaultValue)
	}

	return Some(f(*o.v))
}

// MapOrElse computes a default function result (if none),
// or applies a different function to the contained value (if any).
func MapOrElse[T, U any](o Option[T], defaultValue func() U, f func(T) U) Option[U] {
	if o.v == nil {
		return Some(defaultValue())
	}

	return Some(f(*o.v))
}

// And returns `None` if the option `o` is `None``, otherwise returns `b`.
func And[T, U any](o Option[T], b Option[U]) Option[U] {
	if o.v == nil {
		return None[U]()
	}

	return b
}

// AndThen returns `None` if the option `o` is `None`,
// otherwise calls `f` with the wrapped value and returns the result.
func AndThen[T, U any](o Option[T], f func(T) Option[U]) Option[U] {
	if o.v == nil {
		return None[U]()
	}

	return f(*o.v)
}

// Or returns the option `o` if it contains a value, otherwise returns `b`.
func Or[T any](o, b Option[T]) Option[T] {
	if o.v != nil {
		return o
	}

	return b
}

// OrElse returns the option `o` if it contains a value, otherwise calls `f`` and returns the result.
func OrElse[T any](o Option[T], f func() Option[T]) Option[T] {
	if o.v != nil {
		return o
	}

	return f()
}

// Contains returns `true` if the option `o` is a `Some` value containing the given value.
func Contains[T comparable](o Option[T], v T) bool {
	return o.v != nil && *o.v == v
}
