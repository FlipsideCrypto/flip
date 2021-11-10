package utils

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
)

func GetFilesInPath(rootPath string, fileExt string) []string {
	var files []string

	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) != fileExt {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	// for _, file := range files {
	// 	fmt.Println(file)
	// }

	return files
}

func GetUserHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}
