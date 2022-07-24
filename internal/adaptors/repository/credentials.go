package repository

func (a *DbAdapter) StoreKey(string) error {
	return nil
}

func (a *DbAdapter) AuthenticateKey(string) bool {
	return true
}

func (a *DbAdapter) DeleteKey(string) error {
	return nil
}
