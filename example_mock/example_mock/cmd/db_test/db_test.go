package db_test

import (
	"example_mock/internal/db"
	"fmt"
	"log"

	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func GetAndPrintNames(dbService db.DBService) {
	names, err := dbService.GetNames()
	if err != nil {
		log.Fatal(err)
	}
	for _, name := range names {
		fmt.Println(name)
	}
}

func TestGetAndPrintNamesIntegration(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("получена неожиданная ошибка '%s' при открытии поддельного соединения с базой данных", err)
	}
	defer mockDB.Close()

	dbService := db.New(mockDB)

	expectedNames := []string{"John", "Doe", "Alice"}

	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("John").AddRow("Doe").AddRow("Alice"))

	names, err := dbService.GetNames()
	if err != nil {
		t.Fatalf("Failed to get names from the database: %v", err)
	}

	for _, name := range expectedNames {
		found := false
		for _, n := range names {
			if n == name {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected name %s not found in the result", name)
		}
	}
}
