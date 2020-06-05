package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

const (
	JPG string = "jpg"
	GIF string = "gif"
	PNG string = "png"
	UNK string = "unknown"
)

const (
	KB = 1 << (10 * (iota + 1))
	MB
	GB
)

type image struct {
	contents []byte
	fileType string
	fileSize int
}

func (i *image) Type() string {
	return i.fileType
}

func NewImage(loc string) *image {
	f, err := os.Open(loc)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	m := new(image)

	// check file size and chunk file if too big
	fileInfo, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}
	if fileInfo.Size() >= 5 * MB {
		// chunk the file
	}
	chunkSize := fileInfo.Size()
	bytes := chunkBytes(f, chunkSize)

	fmt.Printf("%X, %X\n", bytes[0], bytes[1])

	return m
}

//func NewImageReader(fileLoc string) *imageReader {
//	contents := make([]byte)
//	contents, err =
//	return &imageReader{fileLoc: fileLoc}
//}

func chunkBytes(file *os.File, chunkSize int64) []byte {
	bytes := make([]byte, chunkSize)

	_, err := file.Read(bytes)
	if err != nil {
		if err == io.EOF {
			return bytes
		}
		log.Fatal(err)
	}

	return bytes
}

//func (a *imageReader) Read(p []byte) (n int, err error) {
//	if a.
//}

//func ImageType(c *image) {
//
//}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Filename must be passed as an argument.")
		os.Exit(0)
	}
	filename := os.Args[1]
	imageFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer imageFile.Close()

	image := NewImage(filename)
	image.Type()

	//fmt.Println(chunkBytes(imageFile, 10))
}

