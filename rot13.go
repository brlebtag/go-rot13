package main

import (
    "os"
    "bufio"
	"log"
    "io"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(p []byte) (n int, err error) {
	n,err = rot.r.Read(p)
	for i := 0; i < len(p); i++ {
		if (p[i] >= 'A' && p[i] < 'N') || (p[i] >='a' && p[i] < 'n') {
			p[i] += 13
		} else if (p[i] > 'M' && p[i] <= 'Z') || (p[i] > 'm' && p[i] <= 'z'){
			p[i] -= 13
		}
	}
	return
}

func main() {
    info, err := os.Stdin.Stat()

    if err != nil {
        log.Fatal(err)
    }

    if info.Mode() & os.ModeCharDevice != os.ModeCharDevice {
        stdIn := bufio.NewReader(os.Stdin)
        reader := rot13Reader{stdIn}
        io.Copy(os.Stdout, &reader)
    }
}