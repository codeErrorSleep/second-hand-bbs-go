package logging

import (
	"fmt"
	"log"
	"os"
	"syscall"
	"time"
)

var (
	LogSavePath = "runtime/logs/"
	LogSaveName = "log"
	LogFileExt  = "log"
	TimeFormat  = "20060102"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s", LogSavePath)
}

func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)

	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

func openLogFile(filePath string) *os.File {
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		mkDir()
	case os.IsPermission(err):
		log.Fatalf("Permission :%v", err)
	}

	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile :%v", err)
	}

	return handle
}

func mkDir() {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/"+getLogFilePath(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}

type PathError struct {
	Op   string
	Path string
	Err  error
}

const (
	// Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
	O_RDONLY int = syscall.O_RDONLY // 以只读模式打开文件
	O_WRONLY int = syscall.O_WRONLY // 以只写模式打开文件
	O_RDWR   int = syscall.O_RDWR   // 以读写模式打开文件
	// The remaining values may be or'ed in to control behavior.
	O_APPEND int = syscall.O_APPEND // 在写入时将数据追加到文件中
	O_CREATE int = syscall.O_CREAT  // 如果不存在，则创建一个新文件
	O_EXCL   int = syscall.O_EXCL   // 使用O_CREATE时，文件必须不存在
	O_SYNC   int = syscall.O_SYNC   // 同步IO
	O_TRUNC  int = syscall.O_TRUNC  // 如果可以，打开时
)
