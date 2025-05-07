package exception

import (
	"fmt"
)

func PanicLogging(err interface{}) {
	if err != nil {
		fmt.Println("ada error", err)
		panic(err)
	}
}

func MediaError(err interface{}) {
	if err != nil {
		fmt.Println("Error Retrieve file from form-data", err)
		panic(err)
	}
}

func FailGetFile(err interface{}) {
	if err != nil {
		fmt.Println("Failed to get file", err)
		panic(err)
	}
}

func FailReadFile(err interface{}) {
	if err != nil {
		fmt.Println("Failed to read file", err)
		panic(err)
	}
}

func FailSaveFile(err interface{}) {
	if err != nil {
		fmt.Println("Failed to save file", err)
		panic(err)
	}
}

func FailCopyFile(err interface{}) {
	if err != nil {
		fmt.Println("Failed to copy file", err)
		panic(err)
	}
}

func FailDeleteFile(err interface{}) {
	if err != nil {
		fmt.Println("Failed to read delete file", err)
		panic(err)
	}
}

func FailDeleteUploadFile(err interface{}) {
	if err != nil {
		fmt.Println("Failed to read file", err)
		panic(err)
	}
}
