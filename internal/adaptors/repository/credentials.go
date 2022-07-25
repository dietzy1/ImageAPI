package repository

import "github.com/dietzy1/imageAPI/internal/application/core"

func (a *DbAdapter) StoreKey(string) error {
	return nil
}

func (a *DbAdapter) AuthenticateKey(string) bool {
	return true
}

func (a *DbAdapter) DeleteKey(string) error {
	return nil
}

func (a *DbAdapter) Signup(creds *core.Credentials) error {
	return nil
}

func (a *DbAdapter) Signin(string, string) error {
	return nil
}
