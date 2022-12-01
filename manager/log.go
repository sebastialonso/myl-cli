package manager

import (
	"fmt"
)

const (
	_infoPrefix = ">"
)

type Log interface {
	Output(string)
	Raw(string)
}

func NewLog() Log {
	return &log{
		infoPrefix: _infoPrefix,
	}
}

type log struct {
	infoPrefix string
}

func (l *log) Output(st string) {
	fmt.Println(l.infoPrefix, st)
}

func (l *log) Raw(st string) {
	fmt.Println(st)
}