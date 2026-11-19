package ggit

import (
	"compress/zlib"
	"fmt"
	"os"
	"path/filepath"
)

type HashObjectOptions struct {
	Path  string
	Write bool
}

func HashObject(opts *HashObjectOptions) {
	// read the file contents as slice of bytes
	data, err := os.ReadFile(opts.Path)
	panicOnErr(err)

	// create blob
	blob := &Blob{Data: data}
	objectID := blob.ObjectID()

	// print the sha1 hash of the blob - object id
	fmt.Println(objectID)

	// write the object to .git/objects
	if opts.Write {
		// get the .ggit directory and the .ggit/objects dir
		wd, err := os.Getwd()
		panicOnErr(err)
		objectsDir := filepath.Join(wd, ".ggit/objects", objectID[:2])
		err = os.MkdirAll(objectsDir, 0755)
		panicOnErr(err)

		// create the object file
		objectFilePath := filepath.Join(objectsDir, objectID[2:])
		objectFile, err := os.Create(objectFilePath)
		panicOnErr(err)
		defer objectFile.Close()

		// compress the blob using zlib and write to file
		zlibWriter := zlib.NewWriter(objectFile)
		defer zlibWriter.Close()
		err = blob.Encode(zlibWriter)
		panicOnErr(err)
	}
}
