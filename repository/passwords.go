package passwordstorage

import (
	"encoding/json"
	"os"
)

type PasswordEntry struct {
	Reference string `json:"reference"`
	Password  string `json:"password"`
}

type PasswordDB struct {
	Entries []PasswordEntry `json:"entries"`
}

func (r *passwordRepository) SaveEncryptedDB(filename string, db PasswordDB, passphrase []byte) error {
	data, err := json.Marshal(db)
	if err != nil {
		return err
	}

	encryptedData, err := r.encryption.Encrypt(data, passphrase)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, encryptedData, 0644)
}

func (r *passwordRepository) LoadEncryptedDB(filename string, passphrase []byte) (PasswordDB, error) {
	var db PasswordDB
	encryptedData, err := os.ReadFile(filename)
	if err != nil {
		return db, err
	}

	data, err := r.encryption.Decrypt(encryptedData, passphrase)
	if err != nil {
		return db, err
	}

	err = json.Unmarshal(data, &db)
	return db, err
}
