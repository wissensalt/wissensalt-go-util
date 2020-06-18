package file

import (
	"os"
	"path/filepath"
)

func CheckFileSize(file *os.File) int64 {
	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	return fileInfo.Size()
}

func CheckFileSizeByPath(filePath string) int64 {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		panic(err)
	}
	return fileInfo.Size()
}

func ListDirectoryContent(pathToScan string) []string {
	var result []string
	err := filepath.Walk(pathToScan, func(path string, info os.FileInfo, err error) error {
		result = append(result, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	return result
}

func CreateDirectoryIfNotExist(pathToScan string)  {
	if _, err := os.Stat(pathToScan); os.IsNotExist(err) {
		err = os.MkdirAll(pathToScan, 0755)
		if err != nil {
			panic(err)
		}
	}
}

//fileName include path
func CreateFileIfNotExist(fileName string) (*os.File, error)  {
	return os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
}