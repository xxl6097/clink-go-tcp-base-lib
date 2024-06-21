package frputil

import (
	"fmt"
	"github.com/xxl6097/clink-go-tcp-base-lib/util/osutil"
	"log"
	"os/exec"
	"runtime"
)

func Run(dir string) {
	exe_path := dir
	exe_path += runtime.GOOS
	if osutil.IsWindows() {
		exe_path += "/frpc.exe"
	} else {
		exe_path += "/frpc"
	}
	cmd := exec.Command(exe_path, "-c", dir+"frpc.toml")
	err := cmd.Start()
	//err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	pid := cmd.Process.Pid
	fmt.Printf("Started program with PID %d\n", pid)
	//err = cmd.Process.Kill()
	//if err != nil {
	//	log.Fatalf("Failed to kill process: %v", err)
	//}
	fmt.Println("============")
}
