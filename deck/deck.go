package deck

import (
	"github.com/gofrs/uuid"
	"myl/utils"
	"fmt"
	"log"
	"myl/preset"
)

type Deck struct {
	UUID uuid.UUID
	Preset preset.Preset
	Cards Cards
	Rand utils.Rand
 	Config DeckConfig
}

// DeckConfig stores configuration options for a given Deck (like type distribution)
type DeckConfig struct {
	GoldPct int
	AllyPct int
	TalismanPct int
	TotalCards int
}

var baseDeckCfg = DeckConfig{GoldPct: 200, AllyPct: 600, TalismanPct: 200, TotalCards: 20}

func (d *Deck) String() string {
	cardStr := "["
	for _, card := range d.Cards {
		cardStr += card.String()
		cardStr += ", "
	}
	cardStr += "]"
	return fmt.Sprintf("<Deck ID:%s> Cards:%s>", d.UUID, cardStr)
}

func (d *Deck) addCardFromItem(item preset.Item) {
	d.Cards = append(d.Cards, ItemToCard(item))
}

func (d *Deck) AddGoldFromPreset() {
	item := d.Preset.GetRandomGold()
	d.addCardFromItem(item)
}

func (d *Deck) AddAllyFromPreset() {
	item := d.Preset.GetRandomAlly()
	d.addCardFromItem(item)
}

func (d *Deck) AddTalismanFromPreset() {
	item := d.Preset.GetRandomTalisman()
	d.addCardFromItem(item)
}

type SampleDeckOpts struct {
	Seed *int
	PresetID preset.PresetID
}

func SampleDeck(opts SampleDeckOpts) *Deck {
	prst, err := preset.NewPreset(opts.PresetID, nil)
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}
	// SampleDeck should receive options, like edition
	deck := Deck{
		UUID: utils.NewUUID4(),
		Config: baseDeckCfg,
		Preset: *prst,
	}
	rand, err := utils.GenerateRand(opts.Seed)
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}
	deck.Rand = *rand

	for i := 0; i < deck.Config.TotalCards; i++ {
		sample := deck.Rand.GetInt(1000)
		switch {
		case SampleGold(deck.Config, sample):
			deck.AddGoldFromPreset()
		case SampleAlly(deck.Config, sample):
			deck.AddAllyFromPreset()
		case SampleTalisman(deck.Config, sample):
			deck.AddTalismanFromPreset()
		default:
			log.Fatal(fmt.Sprintf("sample error: %d", sample))
			return nil
		}
	}
	return &deck
}

func SampleGold(cfg DeckConfig, sample int) bool {
	if sample <= cfg.GoldPct {
		return true
	}
	return false
}

func SampleAlly(cfg DeckConfig, sample int) bool {
	if sample <= cfg.AllyPct && sample > cfg.GoldPct {
		return true
	}
	return false
}

func SampleTalisman(cfg DeckConfig, sample int) bool {
	if sample >= cfg.TalismanPct {
		return true
	}
	return false
}
