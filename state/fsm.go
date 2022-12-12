package state

import (
    "context"
    "fmt"
    "github.com/qmuntal/stateless"
)

type Machine interface {
    Fire(state Trigger) error
    GetCurrent() State
    GetCurrentTranslated() string
}

type fsm struct {
    StateMachine stateless.StateMachine
}

type State int
type Trigger int

// States
const (
    StateInvalid State = iota
    StateInitial
    StateDeckReady
    StateHandProposed
    StateHandConfirmed
    StateGameStarted
    StateFirstTurnStarted
    StateVigilStarted
    StateVigilFinished
    StateAttackStarted
    StateAttackFinished
    StateDefenseStarted
    StateDefenseFinished
    StateWarStarted
    StateWarFinished
    StateFinalStarted
    StateFinalFinished
    StateEnd
)

// Transitions triggers
const (
    triggerInvalid Trigger = iota
    TriggerSampleDeck
    TriggerProposeHand
    TriggerRejectHand
    TriggerAcceptHand
    triggerStartFirstTurn
    triggerStartVigil
    triggerFinishVigil
    triggerStartAttack
    triggerFinishAttack
    triggerStartDefense
    triggerFinishDefense
    triggerStartWar
    triggerFinishWar
    triggerStartFinal
    triggerFinishFinal
    triggerStartTurn
    triggerEndGame
)

const (
    stateInvalidSigil = "Invalid state"
    stateInitialSigil = "Initial state"
    stateDeckReadySigil = "DeckReady state"
    stateHandProposedSigil = "HandProposed state"
    stateHandConfirmedSigil = "HandConfirmed state"
    stateGameStartedSigil = "HandConfirmed state"
    stateFirstTurnStartedSigil = "FirstTurnStarted state"
    stateVigilStartedSigil
    stateVigilFinishedSigil
    stateAttackStartedSigil
    stateAttackFinishedSigil
    stateDefenseStartedSigil
    stateDefenseFinishedSigil
    stateWarStartedSigil
    stateWarFinishedSigil
    stateFinalStartedSigil
    stateFinalFinishedSigil
    stateEndSigil
)

func NewMachine() (Machine, error) {
    machine := stateless.NewStateMachine(StateInitial)

    // Initial state
    machine.Configure(StateInitial).
        OnEntry(func(_ context.Context, _ ...interface{}) error {
            print(stateToSigil(StateInitial))
            return nil
        }).
        Permit(TriggerSampleDeck, StateDeckReady)
    
    machine.Configure(StateDeckReady).
        OnEntry(func(_ context.Context, _ ...interface{}) error {
            print(stateToSigil(StateDeckReady))
            return nil
        }).
        Permit(TriggerProposeHand, StateHandProposed)
    
    machine.Configure(StateHandProposed).
        OnEntry(func(_ context.Context, _ ...interface{}) error {
            print(stateToSigil(StateHandProposed))
            return nil
        }).
        Permit(TriggerRejectHand, StateDeckReady).
        Permit(TriggerAcceptHand, StateHandConfirmed)
    
    machine.Configure(StateHandConfirmed).
        OnEntry(func(_ context.Context, _ ...interface{}) error {
            print(stateToSigil(StateHandConfirmed))
            return nil
        })
    
    return &fsm{
        StateMachine: *machine,
    }, nil
}

func (f *fsm) Fire(trigger Trigger) error {
    err := f.StateMachine.Fire(trigger)
    return err
}

func (f *fsm) GetCurrent() State {
    val, _ := f.StateMachine.MustState().(State)
    return val
}

func (f *fsm) GetCurrentTranslated() string {
    val, _ := f.StateMachine.MustState().(State)
    return stateToSigil(val)
}

func print(st string) {
    fmt.Println(fmt.Sprintf("Entering %s...", st))
}
 
func stateToSigil(state State) string {
    switch state {
    case StateInvalid:
        return stateInvalidSigil
    case StateInitial:
        return stateInitialSigil
    case StateDeckReady:
        return stateDeckReadySigil
    case StateHandProposed:
        return stateHandProposedSigil
    case StateHandConfirmed:
        return stateHandConfirmedSigil
    default:
        return stateInvalidSigil
    }
} 