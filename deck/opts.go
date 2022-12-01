package deck

import (
    "myl/preset"
)

type interval struct {
    min int
    max int
}

type DeckConfig struct {
    Seed *int
    Size int
    PresetID preset.PresetID
    GoldInterval interval
    AllyInterval interval
    TalismanInterval interval
    TotalCards int
}

type HandConfig struct {
    Size int
}

var DefaultDeckCfg = DeckConfig{
    Size: 50,
    PresetID: preset.ElRetoPreset,
    GoldInterval: interval{
        min: 0, max: 200,
    },
    AllyInterval: interval{
        min: 200, max: 800,
    },
    TalismanInterval: interval{
        min: 800, max: 1000,
    },
}

var DefaultHandCfg = HandConfig{
    Size: 8,
}