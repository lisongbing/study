package main

import (
	"os/exec"
	"fmt"
	"os"
)

func main()  {
	cmd := exec.Command("E:/mygo/server/src/diabloserver/diabloserver.exe")
	cmd.Dir = "E:/mygo/server/src/diabloserver"

	os.FindProcess()
	err := cmd.Start()
	fmt.Println(err)
}
