package storage

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

func GetUserDirectory(username string) string {
	return filepath.Join(RootPath, username)
}

func GetDocumentPath(username, docID string) string {
	return filepath.Join(GetUserDirectory(username), docID+".json")
}

func SaveDocument(username, docID, content string) (int, error) {
	userDir := GetUserDirectory(username)

	// Crear el directorio del usuario si no existe.
	if err := os.MkdirAll(userDir, os.ModePerm); err != nil {
		return 0, err
	}

	// Guardar el documento.
	docPath := GetDocumentPath(username, docID)
	err := ioutil.WriteFile(docPath, []byte(content), 0644)
	if err != nil {
		return 0, err
	}
	return len(content), nil
}

func LoadDocument(username, docID string) (string, error) {
	docPath := GetDocumentPath(username, docID)
	content, err := ioutil.ReadFile(docPath)
	if err != nil {
		if os.IsNotExist(err) {
			return "", errors.New("document not found")
		}
		return "", err
	}
	return string(content), nil
}

func DeleteDocument(username, docID string) error {
	docPath := GetDocumentPath(username, docID)
	return os.Remove(docPath)
}

func ListAllDocuments(username string) (map[string]string, error) {
	userDir := GetUserDirectory(username)

	files, err := ioutil.ReadDir(userDir)
	if err != nil {
		return nil, err
	}

	documents := make(map[string]string)
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			docID := file.Name()[:len(file.Name())-len(".json")]
			content, _ := LoadDocument(username, docID)
			documents[docID] = content
		}
	}
	return documents, nil
}