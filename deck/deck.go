package deck

import (
    "github.com/gofrs/uuid"
    "myl/preset"
)

type Deck interface {
    AddCardFromItem(item preset.Item)
    RemoveCard(c Card)
    Count() int
    Size() int
    GetCardAtIndex(index int) (*Card, error)    
}

type Hand interface {
    HasByUUID(uuid uuid.UUID) bool
    PutCardInPosition(card Card, idx int) error
    List() []string
    Count() int
    Size() int
    Cards() []Card
}