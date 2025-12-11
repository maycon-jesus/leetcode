package main

func countCoveredBuildings(n int, buildings [][]int) int {
	covered := 0
	m := len(buildings)
	if m < 5 {
		// Se temos menos de 5 prédios, impossível ter um coberto
		return 0
	}

	// Maps para guardar min e max de cada linha/coluna
	xMin := make(map[int]int)
	xMax := make(map[int]int)
	yMin := make(map[int]int)
	yMax := make(map[int]int)

	// Inicializar com valores extremos e calcular min/max
	for _, building := range buildings {
		x, y := building[0], building[1]

		if minY, exists := xMin[x]; !exists || y < minY {
			xMin[x] = y
		}
		if maxY, exists := xMax[x]; !exists || y > maxY {
			xMax[x] = y
		}

		if minX, exists := yMin[y]; !exists || x < minX {
			yMin[y] = x
		}
		if maxX, exists := yMax[y]; !exists || x > maxX {
			yMax[y] = x
		}
	}

	for _, building := range buildings {
		x, y := building[0], building[1]

		hasLeft := y > xMin[x]
		hasRight := y < xMax[x]
		hasAbove := x > yMin[y]
		hasBelow := x < yMax[y]

		if hasLeft && hasRight && hasAbove && hasBelow {
			covered++
		}
	}

	return covered
}
