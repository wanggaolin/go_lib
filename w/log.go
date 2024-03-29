package w

import (
	"fmt"
	"log"
	"os"
	"time"
)

// date: 2021/12/10
// email: brach@lssin.com

var Log *loging

func init() {
	log.SetFlags(log.Ldate | log.Ltime)
	Log = &loging{}

}

func (l *loging) Log_debug(arg ...any) {
	log.Println("[DEBUG]", arg)
}

func (l *loging) Log_error(arg ...any) {
	log.Println("[ERROR]", arg)
}

func (l *loging) Log_std_err(arg ...any) {
	timespace := time.Now().Format("2006-01-02 15:04:05")
	fmt.Fprintf(os.Stderr, fmt.Sprintln(timespace, "[ERROR]", arg))
}
func (l *loging) Log_panic(arg ...any) {
	log.Panic("[ERROR]", arg)
}
func (l *loging) Log_warr(arg ...any) {
	log.Println("[WARRING]", arg)
}

func (l *loging) Set_log_file(file_paht string) (err error) {
	log_file, err := os.OpenFile(file_paht, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	log.SetOutput(log_file)
	log.SetFlags(log.Ldate | log.Ltime)
	return err
}
