package go_lib

import (
	"fmt"
	"log"
	"os"
	"time"
)

var Logs *log.Logger

// date: 2021/12/10
// email: brach@lssin.com

func init() {
	log.SetFlags(log.Ldate | log.Ltime)
}

func Log_debug(arg ...string) {
	log.Println("[DEBUG]", arg)
}

func Log_error(arg ...string) {
	log.Println("[ERROR]", arg)
}

func Log_std_err(arg ...string) {
	timespace := time.Now().Format("2006-01-02 15:04:05")
	fmt.Fprintf(os.Stderr, fmt.Sprintln(timespace, "[ERROR]", arg))
}
func Log_panic(arg ...string) {
	log.Panic("[ERROR]", arg)
}
func Log_warr(arg ...string) {
	log.Println("[WARRING]", arg)
}

func Log_file(file_paht string) {
	log_file, log_err := os.OpenFile(file_paht, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if log_err != nil {
		Log_std_err("Faild to open error logger file", log_err.Error())
	}
	log.SetOutput(log_file)
	log.SetFlags(log.Ldate | log.Ltime)
}

func Version() {
	fmt.Println("v0.1")
}
