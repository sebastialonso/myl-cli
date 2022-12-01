package manager

import (
	"fmt"
	"myl/world"
	"myl/input"
	"myl/deck"
	"myl/state"
)

type Manager interface {
	WaitForUserInput() (input.Command, error)
	Run() error
}

type manager struct {
	// To keep control of the entities
	World world.World
	Input input.Input
	StateMachine state.Machine
	AlphaBuilder deck.Builder

	// To keep control of the flow
	playerQuitted bool
	alphaAcceptedHand bool

	// Utilities
	log Log
}

func NewManager() (Manager, error) {
	newInput, err := input.NewInput()
	if err != nil {
		return nil, err
	}
	newStateMachine, err := state.NewMachine()
	if err != nil {
		return nil, err
	}
	newDeckBuilder, err := deck.NewBuilder(
		deck.DefaultDeckCfg, deck.DefaultHandCfg,
	)
	if err != nil {
		return nil, err
	}

	llog := NewLog()
	return &manager{
		World: world.World{},
		Input: newInput,
		StateMachine: newStateMachine,
		AlphaBuilder: newDeckBuilder,
		log: llog,
	}, nil
}

func (m *manager) WaitForUserInput() (input.Command, error) {
	return m.Input.WaitForCommand()
}

func (m *manager) Run() error {
	// build decks
	err := m.CreateDeckForAlpha()
	if err != nil {
		return err
	}
	// propose and select hand
	err = m.CreateAndProposeHandForAlpha()
	if err != nil {
		return err
	}
	
	// start game
	for !m.playerQuitted {
		fmt.Println("waiting for input...")
		command, err := m.WaitForUserInput()
		if err != nil {
			continue
		}
		if command.IsQuit() { m.PlayerQuitted() }
		if command.IsNext() {
		
			// m.MovetoNext() // StateMachine.Fire(appropriate_next)
		}
	}
	return nil
	// keep waiting for commands and updating state until player quits
}

func (m *manager) PlayerQuitted() {
	m.playerQuitted = true
}

func (m *manager) CreateDeckForAlpha() error {
	// Creates a deck, asssign it to World Alpha's and fires the TriggerSampleDeck state transition
	deck, err := m.AlphaBuilder.Build()
	if err != nil {
		return err
	}
	m.World.Alpha.Field.Deck = deck
	m.StateMachine.Fire(state.TriggerSampleDeck)
	return nil
}

func (m *manager) CreateAndProposeHandForAlpha() error {
	var attemptNumber = 0
	for !m.alphaAcceptedHand {
		alphaDeck := m.World.Alpha.Field.Deck
		handSize := m.AlphaBuilder.GetHandConfig().Size - attemptNumber
		hand, err := m.AlphaBuilder.BuildHand(alphaDeck, &handSize)
		if err != nil {
			return err
		}

		m.log.ProposeHand(hand.List())
		// m.StateMachine.Fire(state.TriggerHandProposed)
		m.Input.HandleYesOrNo(
			func() {
				m.alphaAcceptedHand = true 
				m.World.Alpha.Hand = hand
				// m.Builder.RemoveHandFromDeck(deck, hand)
				// m.StateMachine.Fire(state.triggerHandAccepted)
			},
			func() {
				attemptNumber++
				// m.StateMachine.Fire(state.triggerHandRejected)
				m.log.Output("Dealing new hand...")
		})
		
	}
	return nil
	// while a hand isn't accepted:
	// Sample a hand from Deck
	// List() the Hand
	// Fire TriggerProposeHand and move state
	// Wait for user input
	// If accepted:
	// * assign hand to world object
	// * from world, take deck and remove cards in hand from deck
	// * fire TriggerAcceptHand and move state
	// If rejected:
	// 
}