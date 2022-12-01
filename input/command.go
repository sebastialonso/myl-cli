package input

import (
    "fmt"
    "strings"
    mylerror "myl/errors"
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
    CemeteryTk
    ExileTk
    ListTk
    QuitTk
    YesTk
    YTk
    NoTk
    NTk
)

const (
    InvalidSigil CmdSigil = "XXX"
    HandSigil CmdSigil = "hand"
    CardSigil CmdSigil = "card"
    PlaySigil CmdSigil = "play"
    ShowSigil CmdSigil = "show"
    NextSigil CmdSigil = "next"
    CountSigil CmdSigil = "count"
    TargetSigil CmdSigil = "TARGET"
    CemeterySigil CmdSigil = "cemetery"
    ExileSigil CmdSigil = "exile"
    ListSigil CmdSigil = "list"
    QuitSigil CmdSigil = "quit"
    YesSigil CmdSigil = "yes"
    YSigil CmdSigil = "y"
    NoSigil CmdSigil = "no"
    NSigil CmdSigil = "n"
)

const (
    newLine = "\n"
    tab = "---"
    whiteSpace = " "
    maxCommandLength = 3
)


type CommandMap map[CmdToken]CommandMap

func (c CommandMap) Translate() string {
    return buildCommandTreeString(0, c)
}

func (c CommandMap) ExpectsTargets() bool {
    _, ok := c[TargetTk]
    return ok
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
    // map for "cemetery" root command
    validCommands[CemeteryTk] = CommandMap{
        ListTk: nil,
    }
    // map for "exile" root command
    validCommands[ExileTk] = CommandMap{
        ListTk: nil,
    }
    // map for "quit"
    validCommands[QuitTk] = nil
    // map for "y", "yes"
    validCommands[YesTk] = nil
    validCommands[YTk] = nil
    // map for "n", "no"
    validCommands[NoTk] = nil
    validCommands[NTk] = nil
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
    case CemeteryTk:
        return CemeterySigil
    case ExileTk:
        return ExileSigil
    case ListTk:
        return ListSigil
    case QuitTk:
        return QuitSigil
    case YTk:
        fallthrough
    case YesTk:
        return YesSigil
    case NTk:
        fallthrough
    case NoTk:
        return NoSigil
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
    case CemeterySigil:
        return CemeteryTk
    case ExileSigil:
        return ExileTk
    case ListSigil:
        return ListTk
    case QuitSigil:
        return QuitTk
    case YesSigil:
        return YesTk
    case YSigil:
        return YTk
    case NoSigil:
        return NoTk
    case NSigil:
        return NTk
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
    Parse(str string, validation CommandMap) error
    Translate() string
    IsQuit() bool
    IsNext() bool
    IsYes() bool
    IsNo() bool
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
    token := fmt.Sprintf("%d", node.Token)
    if node.Token == TargetTk {
      token += fmt.Sprintf(":%s", *node.Value)
    }
        str += fmt.Sprintf("%s ", token)
        node = node.Child
    }
    return fmt.Sprintf("<Command RAW %s>", str)
}

func (c *commandNode) Translate() string {
    node := c
    str := ""
    for node != nil {
    sigil := string(tokenToSigil(node.Token))
    if node.Token == TargetTk {
        sigil += fmt.Sprintf(":%s", *node.Value)
    }
    str += fmt.Sprintf("%s", sigil)
    str += " "
    node = node.Child
    }
    return fmt.Sprintf("<Command TRANSLATED %s>", str)
}

func NewCommand(str string) (Command, error) {
    newCmd := &commandNode{}
    validationTree := getValidCommandTree()
    
    err := newCmd.Parse(str, validationTree)
    
    if err != nil {
        return nil, err
    }
    return newCmd, nil
}

func (c *commandNode) Parse(str string, currentCmdStructure CommandMap) error {
    cmdTokens := strings.Split(str, whiteSpace)
    currentToken := cmdTokens[0]
    restOfTokens := cmdTokens[1:]	
    
    // Check if the passed token is valid one
    recognizedToken := sigilToToken(CmdSigil(currentToken))
    
    // Some commands will expect Targets. Check for this before running validations
    if currentCmdStructure.ExpectsTargets() && recognizedToken == InvalidTk {
        // If indeed we have a target, skip validations
        c.Token = TargetTk
        c.Value = &currentToken
    } else {
        if recognizedToken == InvalidTk {
            return mylerror.NewCommandError(
                mylerror.New(UnrecognizedTokenErr, "unknown token:", nil),
                currentToken,
            )
        }

        // Check if the command validation tree expects more tokens.
        // If we've reached
        nextCmdStructure, _ := currentCmdStructure[recognizedToken]
        if nextCmdStructure == nil && len(restOfTokens) > 0 {
            return mylerror.NewCommandError(
                mylerror.New(InvalidCommandErr, "command entered does not allow that many arguments:", nil),
                currentToken,
            )
        }
    
        c.Token = recognizedToken
    }
    if len(restOfTokens) == 0 {
        return nil
    }

    nextCmdStructure, _ := currentCmdStructure[recognizedToken]
    tailCmdString := strings.Join(restOfTokens, whiteSpace)
    childNode := commandNode{}
    err := childNode.Parse(tailCmdString, nextCmdStructure)
    c.Child = &childNode
    return err
}

func (c commandNode) IsQuit() bool {
    return c.Token == QuitTk
}

func (c commandNode) IsNext() bool {
    return c.Token == NextTk
}

func (c commandNode) IsYes() bool {
    return c.Token == YesTk || c.Token == YTk
}

func (c commandNode) IsNo() bool {
    return c.Token == NoTk || c.Token == NTk
}