package input

import (
	"bufio"
    "fmt"
	"strings"
	"log"
	"os"
)

const (
	newLineStr = "\n"
	newLineRune = '\n'
)

type Input interface {
	WaitForCommand() (Command, error)
	HandleYesOrNo(
		whenYes func(),
		whenNo func(),
	)
}

type input struct {}

func NewInput() (Input, error) {
	return &input{}, nil
}

func GetStdinReader() (*bufio.Reader, error) {
	return bufio.NewReader(os.Stdin), nil
}

func trimInput(str string) string {
	trimmed := strings.Trim(str, newLineStr)
	return trimmed
}

func WaitForInput() {
	reader, err := GetStdinReader()
	if err != nil {
		log.Fatal(err.Error())
	}
	line, err := reader.ReadString(newLineRune)
	if err != nil {
		log.Fatal(err.Error())
	}
	line = trimInput(line)
	cmd, err := NewCommand(line) 
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(cmd)
	fmt.Println(cmd.Translate())
	
}

func (i *input) WaitForCommand() (Command, error) {
	reader, err := GetStdinReader()
	if err != nil {
		return nil, err
	}
	line, err := reader.ReadString(newLineRune)
	if err != nil {
		return nil, err
	}
	line = trimInput(line)
	cmd, err := NewCommand(line) 
	if err != nil {
		return nil, err
	}
	return cmd, nil
}

func (i *input) HandleYesOrNo(
	whenYes func(),
	whenNo func(),
) {
	interactionInProgress := true
	for interactionInProgress {
		command, err := i.WaitForCommand()
		if err != nil {
			continue
		}
		switch {
		case command.IsYes():
			whenYes()
			interactionInProgress = false
		case command.IsNo():
			whenNo()
			interactionInProgress = false
		default:
			continue
		}
	}
}