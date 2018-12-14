package log

import (
	"fmt"

	log "github.com/jeanphorn/log4go"
)

func init() {
	log.LoadConfiguration("./record.json")

}
func WriteInfoLog(fmtstr string, args ...interface{}) {
	log.LOGGER("Info").Info(fmt.Sprintf(fmtstr, args...))
}

func WriteErrorLog(fmtstr string, args ...interface{}) {
	log.LOGGER("Info").Debug(fmt.Sprintf(fmtstr, args...))
}

func Close() {
	log.Close()
}
