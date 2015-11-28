package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"sync"
	"time"
)

func main() {
	var oblock sync.Mutex
	ob := bytes.Buffer{}
	cmd := exec.Command("cmd/test_command")
	out, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	step, _ := time.ParseDuration("3s")
	if err != nil {
		log.Fatal(err)
	}
	end := make(chan error)
	go func() {
		var buf [1024]byte
		var err error
		var n int
		for err == nil {
			n, err = out.Read(buf[:])
			if n > 0 {
				oblock.Lock()
				ob.Write(buf[:n])
				oblock.Unlock()
			}
		}
		end <- err
	}()
	for err == nil {
		time.Sleep(step)
		select {
		case err = <-end:
		default:
			oblock.Lock()
			ob.WriteTo(os.Stdout)
			oblock.Unlock()
		}
	}
	ob.WriteTo(os.Stdout)
	log.Print("Goodbye!")
}
