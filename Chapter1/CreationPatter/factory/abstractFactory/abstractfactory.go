package abstractfactory

import (
	"errors"
	"fmt"
)

/*
"An Abstract factory defines a function that helps us to create multiple objects of the same family, that creation is done calling the specific factory function related to the object that we want to create. each one of the family objects is defined with a struct and that struct is composed by a base struct that defines the family properties"
*/

const (
	admintUser  string = "admin"
	regularUser string = "regular"
)

type ImUser interface {
	WhoAmI() string
}
type user struct {
	name     string
	kind     string
	autonomy float32
}

func (r *user) WhoAmI() string {
	return fmt.Sprintf("Name:%s\nKind: %s,\nAutonomy: %.2f", r.name, r.kind, r.autonomy)
}

type RegularUser struct {
	user
}

func NewRegularUser() ImUser {
	return &RegularUser{
		user: user{
			name:     "eminem",
			kind:     "rapper",
			autonomy: 100.00,
		},
	}
}

type AdminUser struct {
	user
}

func NewAdminUser() ImUser {
	return &AdminUser{
		user: user{
			name:     "mohammed ali",
			kind:     "fighter",
			autonomy: 70.00,
		},
	}
}

//abstract factory
func CreateUser(userType string) (ImUser, error) {
	if userType == regularUser {
		return NewRegularUser(), nil
	}
	if userType == admintUser {
		return NewAdminUser(), nil
	}
	return nil, errors.New("invalid type")
}
