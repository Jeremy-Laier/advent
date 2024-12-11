package grid

type Grid[T any] struct {
	Li [][]T
}

func BuildGrid(len int) Grid[string] {
	grid := Grid[string]{}
	for range len {
		row := make([]string, len)

		grid.Append(row)
	}

	return grid
}

func (g *Grid[T]) Append(row []T) {
	g.Li = append(g.Li, row)
}

func (g *Grid[T]) Get(coord Coord) T {
	return g.Li[coord.Y][coord.X]
}

func (g *Grid[T]) Set(coord Coord, value T) {
	g.Li[coord.Y][coord.X] = value
}
