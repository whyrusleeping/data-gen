package main

import (
	"flag"
	"fmt"
	human "github.com/dustin/go-humanize"
	randbo "github.com/dustin/randbo"
	"io"
	"os"
)

type byteRepeater struct {
	data []byte
	i    int
}

func (br *byteRepeater) Read(b []byte) (int, error) {
	for i, _ := range b {
		b[i] = br.data[(i+br.i)%len(br.data)]
	}
	return len(b), nil
}

func main() {
	size := flag.String("size", "", "specify number of bytes to output")
	data := flag.String("data", "", "specify string to repeat for contents")
	flag.Parse()

	if *size == "" {
		fmt.Fprintln(os.Stderr, "No Size Specified!")
		return
	}

	var read io.Reader
	if *data == "" {
		read = randbo.New()
	} else {
		read = &byteRepeater{[]byte(*data), 0}
	}

	isize, err := human.ParseBytes(*size)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	read = io.LimitReader(read, int64(isize))

	_, err = io.Copy(os.Stdout, read)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
