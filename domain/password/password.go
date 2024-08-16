package passworddomain

import (
	"fmt"

	"github.com/dacors-m/fingerp/passwordrepository"
)

type PasswordDomain interface{}
type passwordDomain struct {
	passRepo passwordrepository.PasswordRepository
}

func NewPasswordDomain(
	passwordRepository passwordrepository.PasswordRepository) PasswordDomain {
	return passwordDomain{
		passRepo: passwordRepository,
	}
}

func (pd *passwordDomain) SavePassword(
	userName, passphrase string,
	passwordEntry passwordrepository.PasswordEntry) error {

	filename := fmt.Sprintf("%s.enc", userName)
	currDb, err := pd.passRepo.LoadEncryptedDB(
		filename, []byte(passphrase))
	if err != nil {
		return err
	}

	currDb.Entries = append(currDb.Entries, passwordEntry)

	err = pd.passRepo.SaveEncryptedDB(filename, currDb, []byte(passphrase))
	if err != nil {
		return err
	}
	return nil
}
