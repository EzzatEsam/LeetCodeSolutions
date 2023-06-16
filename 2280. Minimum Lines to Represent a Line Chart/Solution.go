func minimumLines(stockPrices [][]int) int {
    if len(stockPrices) == 1 {
		return 0
	}
	if len(stockPrices) == 2 {
		return 1
	}
	sort.Slice(stockPrices , func(i,j int) bool {
		return stockPrices[i][0] < stockPrices[j][0]
	})
	count := 1
	lastPoint := stockPrices[0]

	for i := 1; i < len(stockPrices) -1; i++ {
		if !isSameLine(lastPoint , stockPrices[i] , stockPrices[i+1]) {
			count++
			lastPoint = stockPrices[i]
		}
	}
	return count
}

func isSameLine(point1 []int , point2 []int , point3 []int) bool  {
	crossProduct := (point2[0] - point1[0]) * (point3[1] - point1[1]) - (point2[1] - point1[1]) * (point3[0] - point1[0])
	return crossProduct == 0	
}

