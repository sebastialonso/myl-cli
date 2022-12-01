package deck

import (
	"myl/utils"
	"myl/preset"
	"errors"
	"fmt"
	"myl/deck/types"
)

type builder struct {
	DeckCfg types.DeckConfig
	HandCfg types.HandConfig
	preset *preset.Preset
	rand utils.Rand
}

func NewBuilder(
	dCfg types.DeckConfig,
	hCfg types.HandConfig,
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
	deck := newDeck(b.DeckCfg)
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
	hand := newHand(_size)
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

func (b *builder) GetHandConfig() types.HandConfig {
	return b.HandCfg
}

// TODO Move preset.Item to better package
func (b *builder) drawPresetItemBySample(sample int) (preset.Item, error) {
	switch {
	case b.DeckCfg.GoldInterval.Min <= sample && sample < b.DeckCfg.GoldInterval.Max:
		return b.preset.GetRandomGold(), nil
	case b.DeckCfg.AllyInterval.Min <= sample && sample < b.DeckCfg.AllyInterval.Max:
		return b.preset.GetRandomAlly(), nil
	case b.DeckCfg.TalismanInterval.Min <= sample && sample < b.DeckCfg.TalismanInterval.Max:
		return b.preset.GetRandomTalisman(), nil
	default:
		// TODO turn into error
		return preset.Item{}, errors.New(fmt.Sprintf("sample out of reach (%d). Review deck config", sample))
	}
}

func (b *builder) getRandomCardFromDeck(deck Deck) (*types.Card, error) {
	idx := b.rand.GetInt(deck.Size())
	return deck.GetCardAtIndex(idx)
}