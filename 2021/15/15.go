package main

import (
	"math"

	"github.com/liennie/AdventOfCode/common/load"
	"github.com/liennie/AdventOfCode/common/log"
	"github.com/liennie/AdventOfCode/common/util"
)

func parse(filename string) [][]int {
	res := [][]int{}

	for line := range load.File(filename) {
		res = append(res, util.Split(line, ""))
	}

	return res
}

func smallestRisk(risk [][]int) int {
	totalRisk := make([][]int, len(risk))
	for i := range risk {
		totalRisk[i] = make([]int, len(risk[i]))
		for j := range totalRisk[i] {
			totalRisk[i][j] = math.MaxInt
		}
	}
	totalRisk[0][0] = 0

	points := map[util.Point]bool{
		{X: 0, Y: 0}: true,
	}
	for len(points) > 0 {
		var p util.Point
		for p = range points {
			break
		}
		delete(points, p)

		cur := totalRisk[p.Y][p.X]

		for _, dir := range []util.Point{{Y: -1}, {Y: 1}, {X: -1}, {X: 1}} {
			n := p.Add(dir)
			if n.Y >= 0 && n.X >= 0 && n.Y < len(risk) && n.X < len(risk[n.Y]) &&
				cur+risk[n.Y][n.X] < totalRisk[n.Y][n.X] {
				points[n] = true
				totalRisk[n.Y][n.X] = cur + risk[n.Y][n.X]
			}
		}
	}

	end := util.Point{Y: len(risk) - 1}
	end.X = len(risk[end.Y]) - 1

	return totalRisk[end.Y][end.X]
}

func main() {
	defer util.Recover(log.Err)

	const filename = "input.txt"

	risk := parse(filename)

	// Part 1
	log.Part1(smallestRisk(risk))
}
