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

type Player struct {
	Hand deck.Hand
	Field Field
}

type World struct {
	Alpha Player
	Beta Player
}
