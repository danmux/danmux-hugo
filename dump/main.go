package main

import (
	"log"
	"strings"

	"github.com/mitchellh/go-ps"
	"github.com/shirou/gopsutil/process"
)

func main() {
	opURL := getOpURL()
	if opURL == "" {
		log.Println("circleci-agent not found")
	} else {
		log.Print("output processor is listening on: ", opURL)
	}

	for {
		println(strings.Repeat("fuuuck ", 10*1024*1024))
	}
}

func getOpURL() string {
	pss, err := ps.Processes()
	noErr(err)

	for _, p := range pss {
		if strings.Contains(p.Executable(), "circleci-agent") {
			return getOPArg(int32(p.Pid()))
		}
	}
	return ""
}

func getOPArg(pid int32) string {
	proc, err := process.NewProcess(pid)
	noErr(err)

	if proc == nil {
		log.Fatalf("process: %d not found", pid)
	}

	cmds, err := proc.CmdlineSlice()
	noErr(err)

	for i, c := range cmds {
		if c == "--outerServerUrl" {
			return cmds[i+1]
		}
	}
	return ""
}

func noErr(err error) {
	if err != nil {
		log.Fatalf("err: %s", err)
	}
}
