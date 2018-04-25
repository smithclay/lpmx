package log

import (
	"fmt"
	. "github.com/jasonyangshadow/lpmx/error"
	. "github.com/jasonyangshadow/lpmx/utils"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

var (
	LogError   *log.Logger
	LogWarning *log.Logger
	LogInfo    *log.Logger
	LogDebug   *log.Logger
	LogFatal   *log.Logger
)

func LogInit(dir string) *Error {
	multiouts := false
	if strings.TrimSpace(dir) != "" {
		multiouts = true
	}

	if multiouts {
		current_date := time.Now().Local()
		file := fmt.Sprintf("%s/log-%s", dir, current_date.Format("2006-01-02"))
		if !FolderExist(dir) {
			_, err := MakeDir(dir)
			if err != nil {
				return err
			}
		}
		fp, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			cerr := ErrNew(err, fmt.Sprintf("can't open log file: %s", file))
			return &cerr
		}
		LogError = log.New(io.MultiWriter(fp, os.Stdout, os.Stderr), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
		LogWarning = log.New(io.MultiWriter(fp, os.Stdout, os.Stderr), "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
		LogInfo = log.New(io.MultiWriter(fp, os.Stdout), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		LogDebug = log.New(io.MultiWriter(fp, os.Stdout), "DEBUGG: ", log.Ldate|log.Ltime|log.Lshortfile)
		LogFatal = log.New(io.MultiWriter(fp, os.Stdout, os.Stderr), "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)
	} else {
		LogError = log.New(io.MultiWriter(os.Stdout, os.Stderr), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
		LogWarning = log.New(io.MultiWriter(os.Stdout, os.Stderr), "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
		LogInfo = log.New(io.MultiWriter(os.Stdout), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		LogDebug = log.New(io.MultiWriter(os.Stdout), "DEBUGG: ", log.Ldate|log.Ltime|log.Lshortfile)
		LogFatal = log.New(io.MultiWriter(os.Stdout, os.Stderr), "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)
	}
	return nil
}
