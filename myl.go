package main

import (
	"myl/deck"
	"myl/preset"
	"myl/tui"
	"fmt"
	"log"
)

func main() {
	// RunSampleDeck()
	
	// RunNewPreset()
	// runTUI()
	runGOCUI()
}

func runTUI() {
	tui.RunTUI()
}

func runGOCUI() {
	tui.RunGOCUI()
}

func RunSampleDeck() {
	seed := 3
	opts := deck.SampleDeckOpts{
		Seed: &seed,
		PresetID: preset.ElReto,
	}
	deck := deck.SampleDeck(opts)
	fmt.Println(deck.Preset.String())
	fmt.Println(deck)
}

func RunNewPreset() {
	seed := 37
	preset, err := preset.NewPreset(preset.ElReto, &seed)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(preset)
	fmt.Println(preset.Stats.String())
}
