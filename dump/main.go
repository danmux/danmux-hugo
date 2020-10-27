package main

import (
	"github.com/davecgh/go-spew/spew"
	"log"
	"strings"

	"github.com/mitchellh/go-ps"
	"github.com/shirou/gopsutil/process"
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



	cmds, err := proc.CmdlineSlice()
    noErr(err)

    spew.Dump(cmds)

	// 2020/10/27 09:00:25 /bin/circleci-agent --config /.circleci-runner-config.json --task-data /.circleci-task-data --outerServerUrl https://circleci-internal-outer-build-agent:5500 _internal runner<nil>

}


func noErr(err error) {
	if err != nil {
		log.Fatalf("err: %s", err)
	}
}