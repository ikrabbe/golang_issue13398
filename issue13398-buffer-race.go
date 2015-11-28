package main

import (
	"log"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	out := bytes.Buffer{}
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
