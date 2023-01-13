package configuracion

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	. "os"
)

func InitLogs() {
	var file, err = OpenFile("logs.log", O_RDWR|O_CREATE|O_APPEND, 0666)
	if err != nil {
		fmt.Println("Could Not Open Log File : " + err.Error())
	}
	log.SetOutput(file)
	log.SetFormatter(&log.JSONFormatter{})
}
