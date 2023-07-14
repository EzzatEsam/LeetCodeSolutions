package main

func numIslands(grid [][]byte) int {
    num := 0
    visited := make([][]bool, len(grid))
    for i := 0; i < len(grid); i++ {
        visited[i] = make([]bool, len(grid[0]))
    }
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == '1' && !visited[i][j] {
                DFS(grid, visited, i, j)
                num++
            }
        }
    }
    return num
}

func DFS(grid [][]byte, visited [][]bool, i, j int) {
    if i < 0 || i > len(grid)-1 || j < 0 || j > len(grid[0])-1 {
        return
    }
    if grid[i][j] == '0' || visited[i][j] {
        return
    }
    visited[i][j] = true
    DFS(grid, visited, i+1, j)
    DFS(grid, visited, i-1, j)
    DFS(grid, visited, i, j+1)
    DFS(grid, visited, i, j-1)
}

