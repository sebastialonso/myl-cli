package world

import (
	"myl/deck"
)

type Field struct {
	Deck deck.Deck
	// Cemetery
	// Exile
	// ReservePile
	// PayedPile
	// SupportLine
	// DefenseLine
	// AttackLine
}

type Hand struct {
	Limit int
	Cards deck.Cards
}

type Player struct {
	Hand Hand
	Field Field
}

type World struct {
	Alpha Player
	Beta Player
}
