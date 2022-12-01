package deck

import (
    "github.com/gofrs/uuid"
    "myl/preset"
    types "myl/deck/types"
)

type Deck interface {
    AddCardFromItem(item preset.Item)
    RemoveCard(c types.Card)
    Count() int
    Size() int
    GetCardAtIndex(index int) (*types.Card, error)
    
}

type Hand interface {
    HasByUUID(uuid uuid.UUID) bool
    PutCardInPosition(card types.Card, idx int) error
    List() []string
    Count() int
    Size() int
}

type Builder interface {
	Build() (Deck, error)
	BuildHand(deck Deck, size *int) (Hand, error)
	GetHandConfig() types.HandConfig
}

func newDeck(cfg types.DeckConfig) Deck {
    return types.NewBaseDeck(cfg)
}

func newHand(size int) Hand {
    return types.NewBaseHand(size)
}

var DefaultDeckCfg = types.DefaultDeckCfg
var DefaultHandCfg = types.DefaultHandCfg