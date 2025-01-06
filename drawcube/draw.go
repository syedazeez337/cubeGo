package drawcube

import (
	"github.com/gdamore/tcell/v2"
	"github.com/syedazeez337/cubeGo/cube"
)

var faceColors = map[string]tcell.Color{
	"W": tcell.ColorWhite,  // White for Top
	"Y": tcell.ColorYellow, // Yellow for Bottom
	"G": tcell.ColorGreen,  // Green for Left
	"B": tcell.ColorBlue,   // Blue for Right
	"R": tcell.ColorRed,    // Red for Front
	"O": tcell.ColorOrange, // Orange for Back
}

func DrawCube(screen tcell.Screen, cube *cube.Cube) {
	// Define offsets for each face
	offsets := map[int][2]int{
		0: {5, 1},  // Top face
		1: {5, 5},  // Bottom face
		2: {1, 3},  // Left face
		3: {9, 3},  // Right face
		4: {5, 3},  // Front face
		5: {13, 3}, // Back face (Optional: visualize separately)
	}

	// Draw each face
	for face, offset := range offsets {
		startX, startY := offset[0], offset[1]
		for row := 0; row < 3; row++ {
			for col := 0; col < 3; col++ {
				// Sticker background color
				color := faceColors[cube.Faces[face][row][col]]
				style := tcell.StyleDefault.Background(color).Foreground(tcell.ColorBlack)

				// Draw the cell as a block
				screen.SetContent(startX+col*4, startY+row*2, '+', nil, tcell.StyleDefault) // Top-left corner
				screen.SetContent(startX+col*4+3, startY+row*2, '+', nil, tcell.StyleDefault) // Top-right corner
				screen.SetContent(startX+col*4, startY+row*2+1, '|', nil, tcell.StyleDefault) // Vertical border
				screen.SetContent(startX+col*4+3, startY+row*2+1, '|', nil, tcell.StyleDefault) // Vertical border
				for x := 1; x <= 2; x++ {
					screen.SetContent(startX+col*4+x, startY+row*2, '-', nil, tcell.StyleDefault)   // Horizontal top
					screen.SetContent(startX+col*4+x, startY+row*2+2, '-', nil, tcell.StyleDefault) // Horizontal bottom
				}

				// Draw centered text for the color
				screen.SetContent(startX+col*4+1, startY+row*2+1, rune(cube.Faces[face][row][col][0]), nil, style)
			}
		}
	}
}
