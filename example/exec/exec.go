// Package exec is just sample
package exec

import (
	"os/exec"
	"strings"
)

// Exec is to exec command
func Exec(cmd, strParam string) error {
	params := strings.Split(strParam, " ")

	//cmdstr := "ip route | grep default"
	//out, err := exec.Command("sh", "-c", cmdstr).Output()

	//err := exec.Command("ls", "-la").Run()
	err := exec.Command(cmd, params...).Run()
	return err
}

// GetExec is to exec command and get return
func GetExec(cmd, strParam string) (string, error) {
	params := strings.Split(strParam, " ")

	//out, err := exec.Command("ls", "-la").Output()
	out, err := exec.Command(cmd, params...).Output()
	return string(out), err
}

// AsyncExec is to exec command asynchronously (fastest exec)
func AsyncExec(cmd, strParam string) error {
	params := strings.Split(strParam, " ")

	//err := exec.Command("nc", "-l").Start()
	err := exec.Command(cmd, params...).Start()
	return err
}
