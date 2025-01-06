package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/syedazeez337/cubeGo/cube"
	"github.com/syedazeez337/cubeGo/drawcube"
)

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("Error creating screen: %v", err)
	}
	defer screen.Fini()

	if err := screen.Init(); err != nil {
		log.Fatalf("Error initialising screen: %v", err)
	}

	cube := cube.NewCube()

	for {
		screen.Clear()
		drawcube.DrawCube(screen, cube)
		screen.Show()

		// Poll for events
		ev := screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape {
				return // Exit on ESC key
			}
			// Handle more keys here for rotations
		case *tcell.EventResize:
			screen.Sync() // Handle terminal resizing
		}
	}

}
