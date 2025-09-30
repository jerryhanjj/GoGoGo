package interfacego

import (
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/tour/reader"
)

func ReaderExample() {
	r := strings.NewReader("Hello, Reader interface!")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Println(n, err, b)
		fmt.Printf("Read bytes: %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

type MyReader struct{}

// 不需要使用接收者的字段或方法，可以省略接收者名称
func (MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func ReaderMyReaderDemo() {
	reader.Validate(MyReader{})
}

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)
	for i := 0; i < n; i++ {
		if (b[i] >= 'A' && b[i] <= 'M') || (b[i] >= 'a' && b[i] <= 'm') {
			b[i] += 13
		} else if (b[i] >= 'N' && b[i] <= 'Z') || (b[i] >= 'n' && b[i] <= 'z') {
			b[i] -= 13
		}
	}
	return n, err
}

func ReaderRot13Demo() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
	fmt.Println()
}
