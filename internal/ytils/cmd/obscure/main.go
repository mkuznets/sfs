package main

import (
	"bytes"
	"fmt"
	"io"
	"mkuznets.com/go/sps/internal/ytils/ycrypto"
	"os"
	"strings"
)

func main() {
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, os.Stdin); err != nil {
		panic(err)
	}
	fmt.Println(ycrypto.MustObscure(strings.TrimSpace(buf.String())))
}
