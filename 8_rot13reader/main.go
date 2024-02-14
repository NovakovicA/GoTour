package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

const TOTAL_NUMBER_OF_LETTERS = 'z' - 'a' + 1

func getLetterASCIITableType(b byte) int {
	if b >= 'a' && b <= 'z' {
		return 1
	} else if b >= 'A' && b <= 'Z' {
		return 2
	} else {
		return 0
	}
}

func (reader rot13Reader) Read(b []byte) (int, error) {
	var inputBytes []byte = make([]byte, 4, 4)
	totalBytes := 0

	for {
		n, err := reader.r.Read(inputBytes)

		if n < 1 || (err != nil && err != io.EOF) {
			break;
		}

		for i := 0; i < n; i++ {
			currentChar := inputBytes[i]

			switch getLetterASCIITableType(currentChar) {
			case 1:
				b[totalBytes + i] = ((inputBytes[i] - 'a' + 13) % TOTAL_NUMBER_OF_LETTERS) + 'a'
			case 2:
				b[totalBytes + i] = ((inputBytes[i] - 'A' + 13) % TOTAL_NUMBER_OF_LETTERS) + 'A'
			default:
				b[totalBytes + i] = inputBytes[i]
			}
		}

		totalBytes += n

		if err == io.EOF {
			break
		}
	}

	return totalBytes, io.EOF
}


func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)

	fmt.Println("")
	fmt.Println("")

	s = strings.NewReader("Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. Gur dhvpx oebja sbk whzcf bire 13 ynml qbtf. ")
	r = rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
