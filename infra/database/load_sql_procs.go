package database

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func RunSQLScriptsInFolder(folderPath string) error {
	files, err := os.ReadDir(folderPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".sql") {
			filePath := filepath.Join(folderPath, file.Name())
			fmt.Println("Running SQL script:", filePath)

			sqlScript, err := readSQLScript(filePath)
			if err != nil {
				return err
			}

			err = executeSQL(sqlScript)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func readSQLScript(filePath string) (string, error) {
	sqlBytes, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(sqlBytes), nil
}

func executeSQL(sqlScript string) error {
	tx := instance.Exec(sqlScript)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
