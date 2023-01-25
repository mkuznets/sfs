package main

import (
	"bytes"
	"fmt"
	"io"
	"mkuznets.com/go/sfs/ytils/ycrypto"
	"mkuznets.com/go/sfs/ytils/yerr"
	"os"
	"strings"
)

func main() {
	var buf bytes.Buffer
	yerr.Must(io.Copy(&buf, os.Stdin))
	fmt.Println(yerr.Must(ycrypto.Obscure(strings.TrimSpace(buf.String()))))
}
