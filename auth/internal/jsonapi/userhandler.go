package jsonapi

import (
	"errors"
	"net/http"
)

// IndexUsers endpoint
func (ep *Endpoint) IndexUsers(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}

// NewUser endpoint
func (ep *Endpoint) NewUser(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}

// CreateUser endpoint
func (ep *Endpoint) CreateUser(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}

// ShowUser api endpoint
func (ep *Endpoint) ShowUser(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}

// EditUser api endpoint
func (ep *Endpoint) EditUser(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}

// UpdateUser api endpoint
func (ep *Endpoint) UpdateUser(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}

// DeleteUser api endpoint
func (ep *Endpoint) DeleteUser(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}

// SignUpUser api endpoint
func (ep *Endpoint) SignUpUser(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}

// SignInUser api endpoint
func (ep *Endpoint) SignInUser(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}

// SignOutUser api endpoint
func (ep *Endpoint) SignOutUser(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}

// Misc
func (ep *Endpoint) getToken(r *http.Request) (token string, err error) {
	ctx := r.Context()
	token, ok := ctx.Value(ConfCtxKey).(string)
	if !ok {
		err := errors.New("no token provided")
		return "", err
	}

	return token, nil
}
