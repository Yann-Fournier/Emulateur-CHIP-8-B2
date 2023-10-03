package main

import (
	"image/color"
	"log"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"main.go/emulator"
)

type Game struct {
	cpu emulator.Cpu
}

// update du jeu
func (g *Game) Update() error {
	g.cpu.Update()
	time.Sleep(time.Millisecond)
	return nil
}

// dessin des pixels du jeu
func (g *Game) Draw(screen *ebiten.Image) {
	// ebitenutil.DebugPrint(screen,"Hello")
	for x, row := range g.cpu.Gfx {
		for y, pixel := range row {
			if pixel == 1 {
				screen.Set(x, y, color.White)
			} else {
				screen.Set(x, y, color.Black)
			}
		}
	}
}

// fonction pour set l'écran
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 64, 32
}

// fonction pour start la game, ouverture du screen

func main() {
	filename := os.Args[1]
	rombytes := emulator.ReadROM(filename)
	// fmt.Println(rombytes)
	// PrintROM(rombytes)

	var game Game
	emulator.InitCpu(&game.cpu, rombytes)
	// fmt.Println(game.cpu.Memory)

	ebiten.SetWindowSize(640, 320)
	ebiten.SetWindowTitle("Chip8 Emulator")
	ebiten.RunGame(&game)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}

}
