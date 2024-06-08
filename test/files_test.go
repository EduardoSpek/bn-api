package test

import (
	"testing"

	"github.com/eduardospek/notabaiana-backend-golang/internal/utils"
)

func TestFileExists(t *testing.T) {
	filename := "/images/0aa774a9-a897-4623-afdb-6028aa629ba0.jpg"

	exists := utils.FileExsists(filename)

	if !exists {
		t.Error("O arquivo não existe!")
	}

}