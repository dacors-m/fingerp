package passwordstorage

import "github.com/dacors-m/fingerp/utils"

type PasswordRepository interface {
	SaveEncryptedDB(filename string, db PasswordDB, passphrase []byte) error
	LoadEncryptedDB(filename string, passphrase []byte) (PasswordDB, error)
}

type passwordRepository struct {
	encryption *utils.Encryption
}

func NewPasswordRepository() PasswordRepository {
	return &passwordRepository{
		encryption: &utils.Encryption{},
	}
}
