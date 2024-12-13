package grid

import "errors"

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
	var result T

	if coord.X < 0 || coord.X > len(g.Li) {
		return result
	}

	if coord.Y < 0 || coord.Y > len(g.Li) {
		return result
	}

	return g.Li[coord.Y][coord.X]
}

func (g *Grid[T]) SafeGet(coord Coord) (T, error) {
	var result T

	if coord.X < 0 || coord.X >= len(g.Li) {
		return result, errors.New("invalid X coordinate")
	}

	if coord.Y < 0 || coord.Y >= len(g.Li) {
		return result, errors.New("invalid Y coordinate")
	}

	return g.Li[coord.Y][coord.X], nil
}

func (g *Grid[T]) Set(coord Coord, value T) {
	g.Li[coord.Y][coord.X] = value
}
