package exec

import (
	"os/exec"
	"strings"
)

func Exec(cmd, strParam string) error {
	params := strings.Split(strParam, " ")

	//cmdstr := "ip route | grep default"
	//out, err := exec.Command("sh", "-c", cmdstr).Output()

	//err := exec.Command("ls", "-la").Run()
	err := exec.Command(cmd, params...).Run()
	return err
}

func GetExec(cmd, strParam string) (string, error) {
	params := strings.Split(strParam, " ")

	//out, err := exec.Command("ls", "-la").Output()
	out, err := exec.Command(cmd, params...).Output()
	return string(out), err
}

//Fastest exec
func AsyncExec(cmd, strParam string) error {
	params := strings.Split(strParam, " ")

	//err := exec.Command("nc", "-l").Start()
	err := exec.Command(cmd, params...).Start()
	return err
}
