package deck

import (
    "github.com/gofrs/uuid"
    "fmt"
    "myl/utils"
    "errors"
    "myl/preset"
)

type deck struct {
    UUID uuid.UUID
    Cards Cards
    Config DeckConfig
}

func newBaseDeck(cfg DeckConfig) Deck {
    return &deck{
        UUID: utils.NewUUID4(),
        Config: cfg,
    }
}

func (d *deck) String() string {
    cardStr := "["
    for _, card := range d.Cards {
        cardStr += card.String()
        cardStr += ", "
    }
    cardStr += "]"
    return fmt.Sprintf("<Deck ID:%s> Cards:%s>", d.UUID, cardStr)
}

func (d *deck) AddCardFromItem(item preset.Item) {
    // TODO Turn ItemToCard into item.ToCard, with Item an interface for better testing
    d.Cards = append(d.Cards, ItemToCard(item))
}

func (d *deck) Size() int {
    return d.Config.Size
}

func (d *deck) Count() int {
    return len(d.Cards)
}

func (d *deck) GetCardAtIndex(index int) (*Card, error) {
    if index >= len(d.Cards) {
        return nil, errors.New("invalid index: value out of bound")
    }
    card := d.Cards[index]
    return &card, nil
}

func (d *deck) RemoveCard(c Card) {
    auxCards := d.Cards[:0]
    for _, card := range d.Cards {
        if card.UUID != c.UUID {
            auxCards = append(auxCards, card)
        }
    }
    d.Cards = auxCards
}

type hand struct {
	cards Cards
    size int
}

func newBaseHand(size int) Hand {
    cards := make([]Card, size)
    return &hand{
        cards: cards,
        size: size,
    }
}

func (h *hand) Count() int {
    return len(h.cards)
}

func (h *hand) Size() int {
    return h.size
}

func (h *hand) List() []string {
    codes := make([]string, len(h.cards))
    for i, card := range h.cards {
        codes[i] = fmt.Sprintf("%s (%s)", card.Name, string(card.Code))
    }
    return codes
}

func (h *hand) PutCardInPosition(card Card, idx int) error {
    if idx > len(h.cards) {
        return errors.New("asdas")
    }
    h.cards[idx] = card
    return nil
}

func (h *hand) String() string {
    cardStr := "["
    for _, card := range h.cards {
        cardStr += card.String()
        cardStr += ", "
    }
    cardStr += "]"
    return fmt.Sprintf("<Hand Cards:%s>", cardStr)
}

func (h *hand) HasByUUID(uuid uuid.UUID) bool {
    for _,hand := range h.cards {
        if hand.UUID == uuid {
            return true
        }
    }
    return false
}

func (h *hand) Cards() []Card {
    return h.cards
}