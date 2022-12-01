package main

import (
    "myl/tui"
    "myl/input"
    "myl/manager"
    "fmt"
    "log"
)

func main() {
    // RunSampleDeck()
    
    // RunNewPreset()
    // runTUI()
    // runGOCUI()
    // WaitForInput()
    // startMachine()
    // start()
    run()
}

func runTUI() {
    tui.RunTUI()
}

func runGOCUI() {
    tui.RunGOCUI()
}

// func RunSampleDeck() {
// 	seed := 3
// 	opts := deck.SampleDeckOpts{
// 		Seed: &seed,
// 		PresetID: preset.ElReto,
// 	}
// 	deck := deck.SampleDeck(opts)
// 	fmt.Println(deck.Preset.String())
// 	fmt.Println(deck)
// }

// func RunNewPreset() {
// 	seed := 37
// 	preset, err := preset.NewPreset(preset.ElReto, &seed)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	fmt.Println(preset)
// 	fmt.Println(preset.Stats.String())
// }

func WaitForInput() {
    input.WaitForInput()
}

func start() {
    manager, err := manager.NewManager()
    if err != nil {
        log.Fatal(err.Error())
    }
    command, err := manager.WaitForUserInput()
    if err != nil {
        log.Fatal(err.Error())
    }
    fmt.Println(command)
}

func run() {
    manager, err := manager.NewManager()
    if err != nil {
        log.Fatal(err.Error())
    }
    err = manager.Run()
    if err != nil {
        log.Fatal(err.Error())
    }
}