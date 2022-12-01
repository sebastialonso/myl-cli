package deck

import (
	"myl/utils"
	"myl/preset"
	"errors"
	"fmt"
)

type Builder interface {
	Build() (Deck, error)
	BuildHand(deck Deck, size *int) (Hand, error)
	GetHandConfig() HandConfig
}

type builder struct {
	DeckCfg DeckConfig
	HandCfg HandConfig
	preset *preset.Preset
	rand utils.Rand
}

func NewBuilder(
	dCfg DeckConfig,
	hCfg HandConfig,
) (Builder, error) {
	newPreset, err := preset.NewPreset(dCfg.PresetID, nil)
	if err != nil {
		return nil, err
	}
	rand, err := utils.GenerateRand(dCfg.Seed)
	if err != nil {
		return nil, err
	}

	return &builder{
		preset: newPreset,
		rand: *rand,
		DeckCfg: dCfg,
		HandCfg: hCfg,
	}, nil
}

func (b *builder) Build() (Deck, error) {
	deck := newBaseDeck(b.DeckCfg)
	for i := 0; i < b.DeckCfg.Size; i++ {
		sample := b.rand.GetInt(1000)
		item, err := b.drawPresetItemBySample(sample)
		if err != nil {
			return nil, err
		}
		deck.AddCardFromItem(item)
	}
	return deck, nil
}

func (b *builder) BuildHand(deck Deck, size *int) (Hand, error) {
	_size := b.HandCfg.Size
	if size != nil {
		_size = *size
	}
	hand := newBaseHand(_size)
	for i := 0; i < hand.Size(); i++ {
		card, err := b.getRandomCardFromDeck(deck)
		if err != nil {
			return nil, err
		}
		
		err = hand.PutCardInPosition(*card, i)
		if err != nil {
			return nil, err
		}
	}
	return hand, nil
}

func (b *builder) GetHandConfig() HandConfig {
	return b.HandCfg
}

// TODO Move preset.Item to better package
func (b *builder) drawPresetItemBySample(sample int) (preset.Item, error) {
	switch {
	case b.DeckCfg.GoldInterval.min <= sample && sample < b.DeckCfg.GoldInterval.max:
		return b.preset.GetRandomGold(), nil
	case b.DeckCfg.AllyInterval.min <= sample && sample < b.DeckCfg.AllyInterval.max:
		return b.preset.GetRandomAlly(), nil
	case b.DeckCfg.TalismanInterval.min <= sample && sample < b.DeckCfg.TalismanInterval.max:
		return b.preset.GetRandomTalisman(), nil
	default:
		// TODO turn into error
		return preset.Item{}, errors.New(fmt.Sprintf("sample out of reach (%d). Review deck config", sample))
	}
}

func (b *builder) getRandomCardFromDeck(deck Deck) (*Card, error) {
	idx := b.rand.GetInt(deck.Size())
	return deck.GetCardAtIndex(idx)
}