package main

import (
	"log"
	"strings"

	"github.com/mitchellh/go-ps"
	"github.com/shirou/gopsutil/process"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	pss, err := ps.Processes()
	noErr(err)

	for _, p := range pss {
		if strings.Contains(p.Executable(), "circleci-agent") {
			dumpProcess(int32(p.Pid()))
			return
		}
	}
	log.Fatal("circleci-agent not found")
}

func dumpProcess(pid int32) {
	proc, err := process.NewProcess(pid)
	noErr(err)

	if proc == nil {
		log.Fatalf("process: %d not found", pid)
	}

	spew.Dump(proc)

	log.Print(proc.Cmdline())
}


func noErr(err error) {
	if err != nil {
		log.Fatalf("err: %s", err)
	}
}