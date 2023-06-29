package w

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// date: 2021/12/10
// email: brach@lssin.com

func Make_range(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func Show_version(GoVersion string, Auchar string, AppVersion string) {
	fmt.Fprintf(os.Stderr, `App Name: %v
App Auchar: %v
GoLang Version: %v
`, AppVersion, Auchar, GoVersion)
}

func Shell_run(command string, arg ...string) (string, error) {
	cmd := exec.Command(command, arg...)
	cmd.Env = os.Environ()
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return "", errors.New(fmt.Sprintf("%v(%v)", strings.TrimSpace(stderr.String()), err.Error()))
	}
	return out.String(), nil
}
