package main

import (

	"fmt"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"main.go/opcodes"
)

type Game struct {
	cpu opcodes.Cpu
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 320
}

func main() {
	filename := os.Args[1]
	rombytes := readROM(filename)
	// fmt.Println(rombytes)
	PrintROM(rombytes)

	var game Game
	opcodes.InitCpu(&game.cpu, rombytes)
	// fmt.Println(game.cpu.Memory)

	ebiten.SetWindowSize(640, 320)
	ebiten.SetWindowTitle("Chip8 Emulator")
	ebiten.RunGame(&Game{})
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}

}

func readROM(filename string) []byte {
	dat, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return dat
}

func PrintROM(rom []byte) {
	for i, byt := range rom {
		if i%2 == 0 {
			fmt.Printf("0x%03x: ", 0x200+i)
		}
		fmt.Printf("%02x", byt)
		if i%2 == 1 {
			fmt.Print("\n")
		}
	}
}
