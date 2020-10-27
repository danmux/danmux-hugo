package main

import (
	"log"
	"os"
	"strings"

	"github.com/mitchellh/go-ps"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	pss, err := ps.Processes()
	noErr(err)

	for _, p := range pss {
		if strings.Contains(p.Executable(), "circleci-agent") {
			dumpProcess(p.Pid())
			return
		}
	}
	log.Fatal("circleci-agent not found")
}

func dumpProcess(pid int) {
	proc, err := os.FindProcess(pid)
	noErr(err)

	if proc == nil {
		log.Fatalf("process: %d not found", pid)
	}

	spew.Dump(proc)
}


func noErr(err error) {
	if err != nil {
		log.Fatalf("err: %s", err)
	}
}