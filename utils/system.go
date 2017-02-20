package utils

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

const defaultGid = 0 // root

func GetGid() int {
	cmd := exec.Command("grep", "^staff:", "/etc/group")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return defaultGid
	}

	id, err := strconv.ParseInt(
		strings.Split(bytes.NewBuffer(out).String(), ":")[2],
		10,
		0,
	)
	if err != nil {
		fmt.Println(err)
		return defaultGid
	}

	return int(id)
}
