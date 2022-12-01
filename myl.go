package main

import (
    "myl/tui"
    "myl/manager"
    "log"
)

func main() {
    run()
}

func runTUI() {
    tui.RunTUI()
}

func runGOCUI() {
    tui.RunGOCUI()
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