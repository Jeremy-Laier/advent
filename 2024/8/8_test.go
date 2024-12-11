package main

import (
	"advent/lib/grid"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAntiNodes(t *testing.T) {
	t.Run("basic case 1", func(t *testing.T) {
		nodes := generateAntinodes(grid.Coord{X: 1, Y: 1}, grid.Coord{X: 2, Y: 2})

		assert.ElementsMatch(t, nodes, []grid.Coord{
			{
				X: 0,
				Y: 0,
			},
			{
				X: 3,
				Y: 3,
			},
		})
		assert.Len(t, nodes, 2)
	})
	t.Run("basic case 1 backwards", func(t *testing.T) {
		nodes := generateAntinodes(grid.Coord{X: 2, Y: 2}, grid.Coord{X: 1, Y: 1})

		assert.ElementsMatch(t, nodes, []grid.Coord{
			{
				X: 0,
				Y: 0,
			},
			{
				X: 3,
				Y: 3,
			},
		})
		assert.Len(t, nodes, 2)
	})

	t.Run("basic case 2", func(t *testing.T) {
		nodes := generateAntinodes(
			grid.Coord{X: 1, Y: 2},
			grid.Coord{X: 2, Y: 1},
		)

		assert.ElementsMatch(t, nodes, []grid.Coord{
			{
				X: 0,
				Y: 3,
			},
			{
				X: 3,
				Y: 0,
			},
		})
		assert.Len(t, nodes, 2)
	})
	t.Run("basic case 2 backwards", func(t *testing.T) {
		nodes := generateAntinodes(
			grid.Coord{X: 2, Y: 1},
			grid.Coord{X: 1, Y: 2},
		)

		assert.ElementsMatch(t, nodes, []grid.Coord{
			{
				X: 0,
				Y: 3,
			},
			{
				X: 3,
				Y: 0,
			},
		})
		assert.Len(t, nodes, 2)
	})

	// . # . . . . . . . . . . . .
	// . . . . . . . . . . . . . .
	// . . . . . . . . . . . . . .
	// . . . . . . . . . . . . . .
	// . . . . . A . . . . . . . .
	// . . . . . . . . . . . . . .
	// . . . . . . . . . . . . . .
	// . . . . . . . . . . . . . .
	// . . . . . . . . . A . . . .
	// . . . . . . . . . . . . . .
	// . . . . . . . . . . . . . .
	// . . . . . . . . . . . . . .
	// . . . . . . . . . . . . . #
	t.Run("long distance case 1", func(t *testing.T) {
		nodes := generateAntinodes(grid.Coord{X: 0, Y: 0}, grid.Coord{X: 4, Y: 4})

		assert.ElementsMatch(t, nodes, []grid.Coord{
			{
				X: -4,
				Y: -4,
			},
			{
				X: 8,
				Y: 8,
			},
		})
		assert.Len(t, nodes, 2)
	})
	// . # . . . . . . . . . . . .
	// . . . . . . . . . . . . . .
	// . . . . . . . . . . . . . .
	// . . . . . . . . . . . . . .
	// . . . . . A . . . . . . . .
	// . . . . . . . . . . . . . .
	// . . . . . . . . . . . . . .
	// . . . . . . . . . . . . . .
	// . . . . . . . . . A . . . .
	// . . . . . . . . . . . . . .
	// . . . . . . . . . . . . . .
	// . . . . . . . . . . . . . .
	// . . . . . . . . . . . . . #
	t.Run("long distance in map case 1", func(t *testing.T) {
		nodes := generateAntinodes(grid.Coord{X: 5, Y: 4}, grid.Coord{X: 9, Y: 8})

		assert.ElementsMatch(t, nodes, []grid.Coord{
			{
				X: 1,
				Y: 0,
			},
			{
				X: 13,
				Y: 12,
			},
		})
		assert.Len(t, nodes, 2)
	})
}
