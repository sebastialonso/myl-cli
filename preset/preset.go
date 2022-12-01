package preset

import (
	"myl/common"
	"myl/utils"
	"path"
	"os"
	"errors"
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type PresetID int64
type ItemType string
type Frequency string

const (
	CustomPreset PresetID = iota
	ElRetoPreset
)

const (
	Real Frequency = "real"
	Cortesan Frequency = "cortesano"
	Vasal Frequency = "vasallo"
	NoFreq Frequency = "none"
)

const (
	Ally ItemType = "ally"
	Talisman ItemType = "talisman"
	Gold ItemType = "gold"
)

type Item struct {
	Code common.Code `yaml:"code"`
	Name string `yaml:"name"`
	Type ItemType `yaml:"type"`
	Frequency Frequency `yaml:"frequency"`
	Strength *int `yaml:"strength"`
	Cost *int `yaml:"cost"`
	AbilityText string `yaml:"ability"`
}

type Elements []Item

func (el *Elements) String() string {
	elementStr := "["
	for _, element := range *el {
		elementStr += fmt.Sprintf("<Item name:%s type:%s>", element.Name, element.Type)
		elementStr += ", "
	}
	elementStr += "]"
	return elementStr
}

// Preset stores configurations about the card pool from which a Deck is created. Official editions are one type of presets, a bounded set of cards is a differente type of preset
type Preset struct {
	ID PresetID
	Rand utils.Rand
	FixtureSlug string
	Name string
	IsOfficial bool
	Elements Elements
	IsLoaded bool
	Stats Stats
}

func (p *Preset) String() string {
	return fmt.Sprintf("<Preset ID:%d Name:%s Slug:%s Stats:%s>", p.ID, p.Name, p.FixtureSlug, p.Stats.String())
}

type Stats struct {
	NumCards int
	NumGoldCards int
	NumAllyCards int
	NumTalismanCards int
	NumWeaponCards int
	NumTotemCards int
	Elements Elements
	IsLoaded bool
	Allies []Item
	Talismans []Item
	Golds []Item
	Vasals []Item
	Cortesans []Item
	Royals []Item
	NoFreq []Item
}

func (s *Stats) String() string {
	return fmt.Sprintf(
		"<Stats Cards:%d Allies:%d Talismans:%d Golds:%d Weapons:%d Totems:%d>",
		s.NumCards, s.NumAllyCards, s.NumTalismanCards, s.NumGoldCards, s.NumWeaponCards, s.NumTotemCards,
	)
}

func NewPreset(presetID PresetID, seed *int) (*Preset, error) {
	var err error
	
	rand, err := utils.GenerateRand(seed)
        if err != nil {
                return nil, err
        }
	
	preset := &Preset{
		Rand: *rand,
	}
	switch presetID {
	case ElRetoPreset:
		preset.ID = ElRetoPreset
		preset.Name = "El Reto"
		preset.IsOfficial = true
		preset.FixtureSlug = "el_reto"
		
		err = preset.loadFixture()
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("invalid preset")
	}
	err = preset.loadPresetStats()
        if err != nil {
        	return nil, err
        }
	return preset, nil
}

func (p *Preset) loadFixture() error {
	var elements Elements
	cwd, err := os.Getwd()
	filename := path.Join(cwd, fmt.Sprintf("preset/data/%s.yaml", p.FixtureSlug))
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(content, &elements)
	if err != nil {
		return err
	}
	p.Elements = elements
	p.IsLoaded = true
	return nil
}

func (p *Preset) loadPresetStats() error {
	allies := make([]Item, 0)
	talismans := make([]Item, 0)
	golds := make([]Item, 0)
	royals := make([]Item,0)
	cortesans := make([]Item,0)
	vasals := make([]Item, 0)
	noFreq := make([]Item, 0)
	if !p.IsLoaded {
		return errors.New("cannot get stats: preset fixture not loaded")
	}
	for _, item := range p.Elements {
		switch item.Type {
		case Ally:
			allies = append(allies, item)
		case Talisman:
			talismans = append(talismans, item)
		case Gold:
			golds = append(golds, item)
		default:
			return errors.New(fmt.Sprintf("item with unknown type: (%s, %s)", item.Name, item.Type))
		}
		switch item.Frequency {
		case Real:
			royals = append(royals, item)
		case Cortesan:
			cortesans = append(cortesans, item)
		case Vasal:
			vasals = append(vasals, item)
		case NoFreq:
                        noFreq = append(noFreq, item)
		default:
			return errors.New(fmt.Sprintf("item with unkonwn frequency: (%s, %s)", item.Name, item.Frequency))
		}
	}
	stats := Stats{
		Allies: allies,
		Talismans: talismans,
		Golds: golds,
		Royals: royals,
		Cortesans: cortesans,
		Vasals: vasals,
		NoFreq: noFreq,
		NumCards: len(p.Elements),
		NumAllyCards: len(allies),
		NumTalismanCards: len(talismans),
		NumGoldCards: len(golds),
		NumWeaponCards: 0,
		NumTotemCards: 0,
	}
	p.Stats = stats
	return nil
}

func (p *Preset) GetRandomGold() Item {
	return p.getRandomItemFromSlice(p.Stats.Golds)
}

func (p *Preset) GetRandomTalisman() Item {
	return p.getRandomItemFromSlice(p.Stats.Talismans)
}

func (p *Preset) GetRandomAlly() Item {
	return p.getRandomItemFromSlice(p.Stats.Allies)
}

func (p *Preset) getRandomItemFromSlice(slice []Item) Item{
	idx := p.Rand.GetInt(len(slice))
	return slice[idx]
}