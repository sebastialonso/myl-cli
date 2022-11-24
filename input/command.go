package input

import (
	"fmt"
	"strings"
	"errors"
)

type CmdToken int
type CmdSigil string

const (
	InvalidTk CmdToken = iota
	NextTk
	HandTk
	PlayTk
	ShowTk
	CountTk
	TargetTk
)

const (
	InvalidSigil CmdSigil = "XXX"
        HandSigil CmdSigil = "hand"
        CardSigil CmdSigil = "card"
        PlaySigil CmdSigil = "play"
        ShowSigil CmdSigil = "show"
        NextSigil CmdSigil = "next"
	CountSigil CmdSigil = "count"
	TargetSigil = "TARGET"
)

const (
	newLine = "\n"
	tab = "---"
	whiteSpace = " "
	maxCommandLength = 3
)


type CommandMap map[CmdToken]CommandMap

func (c CommandMap) String() string {
	return buildCommandTreeString(0, c)
}

func getValidCommandTree() CommandMap {
	// "next" --> move to next game state
	// "hand play X" --> moves card X in hand into field
	// "hand show" --> shows cards in hand
	// "hand count" --> returns number of cards in hand
	
	validCommands := make(CommandMap, 0)
	// map for "next" root command
	validCommands[NextTk] = nil
	// map for "hand" root command
	validCommands[HandTk] = CommandMap{
		ShowTk: nil,
		CountTk: nil,
		PlayTk: CommandMap{
			TargetTk: nil,			
		},
	}
	return validCommands
}

func tokenToSigil(t CmdToken) CmdSigil {
	switch(t) {
	case NextTk:
		return NextSigil
	case HandTk:
		return HandSigil
	case ShowTk:
		return ShowSigil
	case CountTk:
		return CountSigil
	case PlayTk:
		return PlaySigil
	case TargetTk:
		return TargetSigil
	default:
		return InvalidSigil
	}
}

func sigilToToken(s CmdSigil) CmdToken {
	switch(s) {
	case NextSigil:
		return NextTk
	case HandSigil:
		return HandTk
	case ShowSigil:
		return ShowTk
	case CountSigil:
		return CountTk
	case PlaySigil:
		return PlayTk
	case TargetSigil:
		return TargetTk
	default:
		return InvalidTk
	}
}

func buildCommandTreeString(lvl int, cm CommandMap) string {
        if cm == nil { return "" }
        tabs := ""
        for i := 0; i < lvl; i++ {
                tabs += tab
        }
        st := ""
        for key := range cm {
                st += tabs + fmt.Sprintf("%s", tokenToSigil(key)) + newLine
                st += buildCommandTreeString(lvl + 1, cm[key])
        }
        return st
}


type Command interface {
	Parse(str string) error
	Translate() string
}


type commandNode struct {
	Token CmdToken
	Value *string
	Child *commandNode
}

func (c *commandNode) String() string {
	node := c
	str := ""
	for node != nil {
		str += fmt.Sprintf("%d ", node.Token)
		node = node.Child
	}
	return fmt.Sprintf("<Command %s>", str)
}

func NewCommand(str string) (Command, error) {
	newCmd := &commandNode{}
	err := newCmd.Parse(str)
	if err != nil {
		return nil, err
	}
	return newCmd, nil
}

func (c *commandNode) Parse(str string) error {
	parts := strings.Split(str, whiteSpace)
	currentPart := parts[0]
	
	token := sigilToToken(CmdSigil(currentPart))
	if token == InvalidTk {
		return errors.New(fmt.Sprintf("invalid command: %s not valid token", currentPart))
	}
	c.Token = token
        
	if len(parts[1:]) == 0 {
		return nil
	}
	tailParts := strings.Join(parts[1:], whiteSpace)	
	childNode := commandNode{}
	childNode.Parse(tailParts)
	c.Child = &childNode
	return nil
}

func (c *commandNode) Translate() string {
	node := c
        str := ""
        for node != nil {
                str += fmt.Sprintf("%s", tokenToSigil(node.Token))
		str += " "
                node = node.Child
        }
        return fmt.Sprintf("<Command %s>", str)
}
