type point struct {
	x int
	y int
}
func exist(board [][]byte, word string) bool {
	visited := make([][]bool,len(board))
	for i := 0; i < len(visited); i++ {  
		visited[i] = make([]bool,len(board[0]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0]  {
				if len(word) == 1 {
					return true
				}
				if continueSearch(board, visited,point{i,j},word,1) {
					return true
				}
			}
		}
	}
    return false
}

func continueSearch(board [][]byte,visitedPoints [][]bool ,startPoint point ,word string ,n int) bool {  
	xlim , ylim := len(board),len(board[0])
	visitedPoints[startPoint.x][startPoint.y] = true
	if startPoint.x != 0 && board[startPoint.x-1][startPoint.y] == word[n] && !visitedPoints[startPoint.x-1][startPoint.y]  {
		if n == len(word)-1  {
			return true
		}
		if continueSearch(board,visitedPoints,point{startPoint.x-1,startPoint.y},word,n+1) {
			return true
		}	
	}
	if startPoint.x != xlim-1 && board[startPoint.x+1][startPoint.y] == word[n] && !visitedPoints[startPoint.x+1][startPoint.y]  {
		if n == len(word)-1  {
			return true
		}
		if continueSearch(board,visitedPoints,point{startPoint.x+1,startPoint.y},word,n+1) {
			return true
		}
	}
	if startPoint.y != 0 && board[startPoint.x][startPoint.y-1] == word[n] && !visitedPoints[startPoint.x][startPoint.y-1]  {
		if n == len(word)-1  {
			return true
		}
		if continueSearch(board,visitedPoints,point{startPoint.x,startPoint.y-1},word,n+1)  {
			return true	
		}
	}
	if startPoint.y != ylim-1 && board[startPoint.x][startPoint.y+1] == word[n] && !visitedPoints[startPoint.x][startPoint.y+1]  {
		if n == len(word)-1  {
			return true
		}
		if continueSearch(board,visitedPoints,point{startPoint.x,startPoint.y+1},word,n+1)   {
			return true
		}
	}
	visitedPoints[startPoint.x][startPoint.y] = false
	return false
}