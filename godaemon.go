package godaemon

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func usage() {
    fmt.Fprintf(os.Stderr, `Options:
start : start the program as a daemon
`)
}

func init() {
	var Daemon string
	flag.Usage = usage
	flag.Parse()
	flag.Usage()
	if len(flag.Args()) > 0 {
		Daemon = flag.Args()[0]
	}

	if Daemon == "start" {
		cmd := exec.Command(os.Args[0], flag.Args()...)
		if err := cmd.Start(); err != nil {
			fmt.Printf("start %s failed, error: %v\n", os.Args[0], err)
			os.Exit(1)
		}
		fmt.Printf("%s [PID] %d running...\n", os.Args[0], cmd.Process.Pid)
		os.Exit(0)
	} else {
		fmt.Println("Invalid parameter")
		os.Exit(1)
	}
}
