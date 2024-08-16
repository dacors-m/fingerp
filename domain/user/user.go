package userdomain

import (
	"fmt"

	"github.com/dacors-m/fingerp/passwordrepository"
)

type UserDomain interface {
	SetupUser(userName, password string) error
}

type userDomain struct {
	passRepo passwordrepository.PasswordRepository
}

func NewUserDomain(
	passwordRepository passwordrepository.PasswordRepository) UserDomain {

	return &userDomain{
		passRepo: passwordRepository,
	}
}

func (u *userDomain) SetupUser(userName, password string) error {

	filename := fmt.Sprintf("%s.enc", userName)
	passphrase := []byte(password)

	newDb := passwordrepository.PasswordDB{}

	err := u.passRepo.SaveEncryptedDB(filename, newDb, passphrase)
	if err != nil {
		return err
	}

	return nil
}
