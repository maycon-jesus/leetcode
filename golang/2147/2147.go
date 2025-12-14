package main

var MOD = int(1e9 + 7)

func numberOfWays(corridor string) int {
	if len(corridor) < 2 {
		return 0
	}
	ways := 1
	seatCount := 0
	plantCount := 0
	totalSeatCount := 0

	for _, ch := range corridor {
		if ch == 'S' {
			if seatCount == 2 {
				ways = ways * (plantCount + 1) % MOD
				seatCount = 0
				plantCount = 0
			}

			seatCount++
			totalSeatCount++
		} else if seatCount == 2 {
			plantCount++
		}
	}

	if totalSeatCount == 0 || totalSeatCount%2 != 0 {
		return 0
	}

	return ways
}
