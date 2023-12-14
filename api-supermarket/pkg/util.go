package pkg

import (
	"api-supermarket/internal/domain"
	"encoding/json"
	"io"
	"log"
	"os"
)

// FillDb fills the database with the data from the json file
func FillDb(path string) []domain.Product {
	data, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	datarRead, err := io.ReadAll(data)
	if err != nil {
		log.Fatal(err)
	}

	slice := []domain.Product{}
	json.Unmarshal(datarRead, &slice)

	return slice
}
