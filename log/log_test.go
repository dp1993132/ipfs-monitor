package log

import (
	"testing"
	"time"
)

func TestLog(t *testing.T) {
	WriteInfoLog("msg: %d", 111)
	WriteInfoLog("msg: %d", 111)
	WriteInfoLog("msg: %d", 111)
	WriteInfoLog("msg: %d", 111)
	time.Sleep(1 * time.Second)
}
