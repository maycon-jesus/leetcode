package main

func maximumProfit(prices []int, k int) int64 {
	n := len(prices)
	if n == 0 || k == 0 {
		return 0
	}

	// dp[i][j][0] = lucro máximo após dia i, com j transações completas, sem posição
	// dp[i][j][1] = lucro máximo após dia i, com j transações, segurando posição long (comprou)
	// dp[i][j][2] = lucro máximo após dia i, com j transações, segurando posição short (vendeu a descoberto)

	const INF = int64(1e18)

	// Inicializa com valores muito negativos
	dp := make([][][]int64, n+1)
	for i := range dp {
		dp[i] = make([][]int64, k+1)
		for j := range dp[i] {
			dp[i][j] = make([]int64, 3)
			dp[i][j][0] = -INF
			dp[i][j][1] = -INF
			dp[i][j][2] = -INF
		}
	}

	dp[0][0][0] = 0

	for i := 0; i < n; i++ {
		price := int64(prices[i])

		for j := 0; j <= k; j++ {
			// Estado sem posição
			if dp[i][j][0] > -INF {
				// Não fazer nada
				dp[i+1][j][0] = max(dp[i+1][j][0], dp[i][j][0])

				// Iniciar normal transaction (comprar)
				if j < k {
					dp[i+1][j][1] = max(dp[i+1][j][1], dp[i][j][0]-price)
				}

				// Iniciar short selling (vender a descoberto)
				if j < k {
					dp[i+1][j][2] = max(dp[i+1][j][2], dp[i][j][0]+price)
				}
			}

			// Estado com posição long (segurando ação)
			if dp[i][j][1] > -INF {
				// Não fazer nada
				dp[i+1][j][1] = max(dp[i+1][j][1], dp[i][j][1])

				// Vender (completar normal transaction)
				dp[i+1][j+1][0] = max(dp[i+1][j+1][0], dp[i][j][1]+price)
			}

			// Estado com posição short (vendeu a descoberto)
			if dp[i][j][2] > -INF {
				// Não fazer nada
				dp[i+1][j][2] = max(dp[i+1][j][2], dp[i][j][2])

				// Comprar de volta (completar short selling)
				dp[i+1][j+1][0] = max(dp[i+1][j+1][0], dp[i][j][2]-price)
			}
		}
	}

	result := int64(0)
	for j := 0; j <= k; j++ {
		result = max(result, dp[n][j][0])
	}

	return result
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
