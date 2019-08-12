package main

import (
	"fmt"
	"time"
	"os/exec"
	"os"
	"syscall"
)

func main() {
/*	fiveDaysFromNow := time.Now().Add(time.Hour * 24 * 5)
	err := SetSystemDate(fiveDaysFromNow)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}*/
	//sec := time.Now().Unix() - 86400
	//t := time.Unix(sec, 0)

	//hms := t.Format("15:04:05")
	// hour-minute-second
/*	err := exec.Command("time").Run()
	if err != nil {
		log.Fatalf("set hour failed,err:%v\n", err)
		return
	}*/
	cmd := exec.Command("time", "-s", time.Unix(1538326861, 1538326861).Format("01/02/2006 15:04:05"))
	cmd.Run()
}

func SetSystemDate(newTime time.Time) error {
	binary, lookErr := exec.LookPath("time")
	if lookErr != nil {
		fmt.Printf("Date binary not found, cannot set system date: %s\n", lookErr.Error())
		return lookErr
	} else {
		//dateString := newTime.Format("2006-01-2 15:4:5")
		dateString := newTime.Format("2 Jan 2006 15:04:05")
		fmt.Printf("Setting system date to: %s\n", dateString)
		args := []string{"--set", dateString}
		env := os.Environ()
		return syscall.Exec(binary, args, env)
	}
}
