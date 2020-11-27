package db

import (
	"context"

	"go-cqrs/model"
)

// Repository interface
type Repository interface {
	InsertWoof(ctx context.Context, woof model.Woof) error
	ListWoofs(ctx context.Context, offset uint64, limit uint64) ([]model.Woof, error)
	Close()
}

// This is a straightforward way of achieving inversion of control.
// By using Repository interface you allow any concrete implementation
// to be injected at runtime, and all function calls will
// be delegated to the impl object.
var impl Repository

// SetRepository implementation
func SetRepository(repository Repository) {
	impl = repository
}

// InsertWoof implementation
func InsertWoof(ctx context.Context, woof model.Woof) error {
	return impl.InsertWoof(ctx, woof)
}

// ListWoofs implementation
func ListWoofs(ctx context.Context, offset uint64, limit uint64) ([]model.Woof, error) {
	return impl.ListWoofs(ctx, offset, limit)
}

// Close implementation
func Close() {
	impl.Close()
}
