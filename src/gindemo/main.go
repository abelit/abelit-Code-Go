package main

import (
	"fmt"
	"gindemo/router"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// Disable Console Color, you don't need console color when writing the logs to file.
	// gin.DisableConsoleColor()

	// Logging to a file.
	// f, _ := os.Create("gin.log")
	dir, _ := filepath.Abs(filepath.Dir(""))
	logDir := dir + "/log/"
	logFileName := time.Now().Format("20060102") + ".log"
	logFilePath := logDir + logFileName
	var f *os.File

	_, err := os.Stat(logDir)

	if err != nil {
		os.MkdirAll(logDir, 0666)
	}

	_, err = os.Stat(logFilePath)

	if err != nil {
		f, _ = os.Create(logFilePath)
	} else {
		f, _ = os.OpenFile(logFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	}

	// gin.DefaultWriter = io.MultiWriter(f)
	// Use the following code if you need to write the logs to file and console at the same time.
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()

	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	r.Use(gin.Recovery())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "home",
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSONP(200, gin.H{
			"msg": "pong",
		})
	})

	r.GET("/JSONP?callback=x", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}

		//callback is x
		// Will output  :   x({\"foo\":\"bar\"})
		c.JSONP(http.StatusOK, data)
	})

	r.GET("/test", router.TestGet)

	r.Run(":8080")
}
