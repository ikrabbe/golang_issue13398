package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"time"
)

type MyBuffer struct {
	intbuf bytes.Buffer
}

func (mb *MyBuffer) WriteTo(w io.Writer) (n int64, err error) {
	return mb.intbuf.WriteTo(w)
}

func (mb *MyBuffer) Len() int {
	return mb.intbuf.Len()
}

func (mb *MyBuffer) Write(p []byte) (n int, err error) {
	return mb.intbuf.Write(p)
}

func main() {
	out := MyBuffer{}
	cmd := exec.Command("cmd/test_command")
	cmd.Stdout = &out
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	step, _ := time.ParseDuration("3s")
	if err != nil {
		log.Fatal(err)
	}
	for {
		time.Sleep(step)
		os.Stderr.Write([]byte(fmt.Sprintf("Buffer Length %d\n", out.Len())))
		out.WriteTo(os.Stdout)
	}
}
