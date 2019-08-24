package utils

import "testing"

func TestIsValidFile(t *testing.T) {
	isValid := IsValidFile("file.cod")
	if !isValid {
		t.Error("The file extension is not valid")
	}
	isValid = IsValidFile("file.algo.cod")
	if isValid {
		t.Error("The file extension is not valid")
	}
	isValid = IsValidFile("file_file.cod")
	if !isValid {
		t.Error("The file extension is not valid")
	}
	isValid = IsValidFile("212file.cod")
	if !isValid {
		t.Error("The file extension is not valid")
	}
	isValid = IsValidFile("_file.cod")
	if !isValid {
		t.Error("The file extension is not valid")
	}
	isValid = IsValidFile("file.Cod")
	if !isValid {
		t.Error("The file extension is not valid")
	}
	isValid = IsValidFile("file.c0d")
	if isValid {
		t.Error("The file extension is not valid")
	}
	isValid = IsValidFile("file.COD")
	if !isValid {
		t.Error("The file extension is not valid")
	}
}

func TestOpenFile(t *testing.T) {
	file, err := OpenFile("./test_file.cod")
	if err != nil {
		t.Errorf("Error opening the file. %v", err)
	}
	if len(file) == 0 {
		t.Error("The file is empty")
	}
}
