package cube

type Cube struct {
	Faces [6][3][3]string
}

func NewCube() *Cube {
	colours := []string{"W", "Y", "G", "B", "R", "O"}
	cube := &Cube{}
	for i, colour := range colours {
		for row := 0; row < 3; row++ {
			for col := 0; col < 3; col++ {
				cube.Faces[i][row][col] = colour
			}
		}
	}
	return cube
}