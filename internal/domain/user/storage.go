package user

import "context"

type Storage interface {
	Create(context.Context, *User) (string, error)
	FindOne(context.Context, string) (*User, error)
	Update(context.Context, *User) error
	Delete(context.Context, string) error
}
