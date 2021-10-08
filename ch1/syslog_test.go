package ch1

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
	"path/filepath"
	"testing"
)

//OS X使用syslog包会把日志输出到/var/log/mail.log
func TestSysLog(t *testing.T) {
	programName := filepath.Base(os.Args[0])
	//New的第一个参数是log等级和log facility的综合表示,第二个参数是日志信息的前缀,一般使用程序名
	sysLog, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, programName)
	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sysLog)
	}
	log.Println("LOG_INFO + LOG_LOCAL7: Logging in Go!")
	sysLog, err = syslog.New(syslog.LOG_MAIL, "Some program!")
	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sysLog)
	}
	log.Println("LOG_MAIL: Logging in Go!")
	fmt.Println("Will you see this?")
}

func TestFatal(t *testing.T) {
	sysLog, err := syslog.New(syslog.LOG_ALERT|syslog.LOG_MAIL,
		"Some program!")
	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sysLog)
	}
	//只有在真的发生很严重的错误的时候才使用Fatal去记录日志
	//调用这个方法只会程序马上回退出
	log.Fatal(sysLog)
	fmt.Println("Will you see this?")
}

func TestPanic(t *testing.T) {
	sysLog, err := syslog.New(syslog.LOG_ALERT|syslog.LOG_MAIL,
		"Some program!")
	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sysLog)
	}
	//调用Panic和Fatal有一样的效果,会让程序马上结束
	log.Panic(sysLog)
	fmt.Println("Will you see this?")
}
