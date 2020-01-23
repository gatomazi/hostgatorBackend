package log

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/sadlil/gologger"
)

//LoggerConsole para logar no console
var loggerConsole gologger.GoLogger

// //LoggerFile para logar em arquivo
var loggerFile gologger.GoLogger

//InicializaLog inicializa os dois loggers
func InicializaLog() {

	// gin.DefaultWriter = io.MultiWriter(f)

	// Use the following code if you need to write the logs to file and console at the same time.

	loggerFile = gologger.GetLogger(gologger.FILE, "gin.log")
	loggerConsole = gologger.GetLogger(gologger.CONSOLE, gologger.ColoredLog)
	logfile, err := os.OpenFile("gin.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to create request log file:", err)
	}
	gin.DefaultWriter = io.MultiWriter(logfile, os.Stdout)
}

//Log - Loga a mensagem no arquivo e no console
func Log(mensagem ...interface{}) {
	msg := fmt.Sprintf("%v", mensagem)
	loggerConsole.Log(msg)
	loggerFile.Log(msg)
}

//Warn - Loga a mensagem no arquivo e no console como um warning
func Warn(mensagem ...interface{}) {
	msg := fmt.Sprintf("%v", mensagem)
	loggerConsole.Warn(msg)
	loggerFile.Warn(msg)
}

//Error - Loga a mensagem no arquivo e no console como um erro
func Error(mensagem ...interface{}) {
	_, file, no, ok := runtime.Caller(1)
	var msg string
	if ok {
		msg = fmt.Sprintf("!ERRO! Arquivo: %v | Linha %v |  %v", file, no, mensagem)
	} else {
		msg = fmt.Sprintf("%v", mensagem)
	}
	loggerConsole.Error(msg)
	loggerFile.Error(msg)
}

//Fatal - Loga a mensagem no arquivo e no console como um erro fatal
func Fatal(mensagem ...interface{}) {
	msg := fmt.Sprintf("%v", mensagem)
	loggerConsole.Fatal(msg)
	loggerFile.Fatal(msg)
}

//Info - Loga a mensagem no arquivo e no console como um Info
func Info(mensagem ...interface{}) {
	msg := fmt.Sprintf("%v", mensagem)
	loggerConsole.Info(msg)
	loggerFile.Info(msg)
}

//Debug - Loga a mensagem somente no console para debug
func Debug(mensagem ...interface{}) {
	msg := fmt.Sprintf("%v", mensagem)
	loggerConsole.Debug(msg)

}
