package state

import (
        "github.com/qmuntal/stateless"
)

type FSM struct {
        Machine stateless.StateMachine
}

type State int
type Trigger int

// States
const (
        stateInvalid State = iota
        stateInitial
        stateDeckReady
        stateHandProposed
        stateHandConfirmed
        stateGameStarted
        stateFirstTurnStarted
        stateVigilStarted
        stateVigilFinished
)

// Transitions triggers
const (
        triggerInvalid Trigger = iota
        triggerSampleDeck
        triggerProposeHand
        triggerRejectHand
        triggerStartGame
)

func NewMachine() (*stateless.StateMachine, error) {
        machine := stateless.NewStateMachine(stateInitial)
}
