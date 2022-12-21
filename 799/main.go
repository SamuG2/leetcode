package main

func champagneTower(poured int, query_row int, query_glass int) float64 {
	glasses := make([][]float64, 100)
	for i := range glasses {
		glasses[i] = make([]float64, i+1)
	}
	glasses[0][0] = float64(poured)

	for i, row := range glasses {
		for j, glass := range row {
			t := (glass - 1) / 2
			if t > 0 {
				if i < 100 {
					glasses[i+1][j] += t
					glasses[i+1][j+1] += t
				}
			}
		}
	}
	if glasses[query_row][query_glass] <= 1 {
		return glasses[query_row][query_glass]
	}
	return 1
}
