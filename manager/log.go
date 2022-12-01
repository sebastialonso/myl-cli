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
	ProposeHand([]string)
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

func (l *log) ProposeHand(list []string) {
	l.Output("This is the proposed hand:")
	proposedHand := ""
	for _, element := range list {
		proposedHand += fmt.Sprintf("%s\n", element)
	}
	l.Raw(proposedHand)
	l.Output("You can redraw your proposed hand, but it will contain one less card.")
	l.Output("Do you want to keep this hand? [y/n]")
}