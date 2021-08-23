package main

import "os"
import "./CellularAutomata"
import "./Renderers"

func main() {
    CellularAutomata := CellularAutomata{2, 1, 110, false, 100, 100}
    Renderers.generateImage(CellularAutomata)
}