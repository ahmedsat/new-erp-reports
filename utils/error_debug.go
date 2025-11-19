//go:build !release

package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func HandelErr(err error) {
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		fmt.Printf("ERROR: %v\n", err)
		if ok {
			fmt.Printf("%s:%d\n", relativePath(file), line)
		} else {
			fmt.Println("unknown")
		}
		os.Exit(1)
	}
}

func relativePath(path string) string {
	cwd, err := os.Getwd()
	HandelErr(err)
	rel, err := filepath.Rel(cwd, path)
	HandelErr(err)
	return rel
}

func WhereAmI() string {
	_, file, line, _ := runtime.Caller(1)
	return fmt.Sprintf("%s:%d", relativePath(file), line)
}
