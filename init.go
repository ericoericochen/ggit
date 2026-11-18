package ggit

import (
	"fmt"
	"os"
	"path/filepath"
)

type InitOptions struct {
}

func Init(opts *InitOptions) {
	// get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// initialize the .ggit directory with objects/, refs/heads, refs/tags, and HEAD
	ggitDir := filepath.Join(wd, ".ggit")
	objectsDir := filepath.Join(ggitDir, "objects")
	refsHeadsDir := filepath.Join(ggitDir, "refs/heads")
	refsTagsDir := filepath.Join(ggitDir, "refs/tags")

	panicOnErr(os.MkdirAll(objectsDir, 0755))
	panicOnErr(os.MkdirAll(refsHeadsDir, 0755))
	panicOnErr(os.MkdirAll(refsTagsDir, 0755))

	// write ref: refs/heads/main to HEAD
	headFile := filepath.Join(ggitDir, "HEAD")
	headContent := []byte("ref: refs/heads/main\n")
	panicOnErr(os.WriteFile(headFile, headContent, 0644))

	fmt.Println(fmt.Sprintf("Initialized empty ggit repository in %v", wd))
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
