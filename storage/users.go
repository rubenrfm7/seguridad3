package storage

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

var UsersFile = filepath.Join(RootPath, "users.json")

type userRecord struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func CreateUser(username, password string) error {
	users, err := loadUsers()
	if err != nil {
		return err
	}

	// Verificar si el usuario ya existe.
	for _, user := range users {
		if user.Username == username {
			return errors.New("user already exists")
		}
	}

	// Agregar nuevo usuario.
	users = append(users, userRecord{Username: username, Password: password})
	return saveUsers(users)
}

func ValidateUser(username, password string) bool {
	users, err := loadUsers()
	if err != nil {
		return false
	}

	for _, user := range users {
		if user.Username == username && user.Password == password {
			return true
		}
	}
	return false
}

func loadUsers() ([]userRecord, error) {
	// Si el archivo no existe, devolver una lista vacía.
	if _, err := os.Stat(UsersFile); os.IsNotExist(err) {
		return []userRecord{}, nil
	}

	file, err := os.Open(UsersFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var users []userRecord
	err = json.NewDecoder(file).Decode(&users)
	return users, err
}

func saveUsers(users []userRecord) error {
	file, err := os.Create(UsersFile)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(users)
}
