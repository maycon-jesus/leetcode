package main

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	rows := make([][]rune, numRows)

	storeRow := 0
	goingDown := false
	for _, k := range s {
		rows[storeRow] = append(rows[storeRow], k)
		if goingDown {
			storeRow--
			if storeRow < 0 {
				storeRow = 1
				goingDown = false
			}
		} else {
			storeRow++
			if storeRow >= numRows {
				storeRow = numRows - 2
				goingDown = true
			}
		}
	}

	stringBuilder := make([]rune, 0, len(s))
	for _, row := range rows {
		stringBuilder = append(stringBuilder, row...)
	}
	return string(stringBuilder)
}
