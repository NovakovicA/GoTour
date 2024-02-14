package main

import (
	"io"
	"golang.org/x/tour/reader"
	"fmt"
)

type MyReader struct{
	charFill byte
}

func (reader MyReader) Read(b []byte) (int, error) {
	if len(b) == 0 {
		return 0, io.EOF
	}

	for i := range(b) {
		b[i] = reader.charFill
	}

	return len(b), nil
}

func main() {
	reader.Validate(MyReader{'A'})

	fmt.Println("")

	testBytes := []byte {0, 0, 0, 0, 0, 0}
	n, err := MyReader{127}.Read(testBytes)
	fmt.Println("testBytes = ", testBytes)
	fmt.Println("Read", n, "bytes in total.")
	fmt.Println("err =", err)

	fmt.Println("")

	testBytes = testBytes[:0]
	n, err = MyReader{127}.Read(testBytes)
	fmt.Println("testBytes = ", testBytes)
	fmt.Println("Read", n, "bytes in total.")
	fmt.Println("err =", err)
}
