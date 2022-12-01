package types

import (
    "myl/preset"
)

type interval struct {
    Min int
    Max int
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
        Min: 0, Max: 200,
    },
    AllyInterval: interval{
        Min: 200, Max: 800,
    },
    TalismanInterval: interval{
        Min: 800, Max: 1000,
    },
}

var DefaultHandCfg = HandConfig{
    Size: 8,
}