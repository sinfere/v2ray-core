// Package common contains common utilities that are shared among other packages.
// See each sub-package for detail.
package common

import (
	"errors"
)

var (
	ErrObjectReleased   = errors.New("Object already released.")
	ErrBadConfiguration = errors.New("Bad configuration.")
	ErrObjectNotFound   = errors.New("Object not found.")
	ErrDuplicatedName   = errors.New("Duplicated name.")
)

// Releasable interface is for those types that can release its members.
type Releasable interface {
	// Release releases all references to accelerate garbage collection.
	Release()
}

// Release tries to release the given object.
func Release(v interface{}) {
	if releasable, ok := v.(Releasable); ok {
		releasable.Release()
	}
}

// Must panics if err is not nil.
func Must(err error) {
	if err != nil {
		panic(err)
	}
}
