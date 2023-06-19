class Solution {
    private class Point {
        int x;
        int y;

        public Point(int x, int y) {
            this.x = x;
            this.y = y;
        }
    }

    public boolean exist(char[][] board, String word) {
        boolean[][] visited = new boolean[board.length][board[0].length];
        for (int i = 0; i < board.length; i++) {
            for (int j = 0; j < board[0].length; j++) {
                if (board[i][j] == word.charAt(0)) {
                    if (word.length() == 1) {
                        return true;
                    }
                    if (continueSearch(board, visited, new Point(i, j), word, 1)) {
                        return true;
                    }
                }
            }
        }
        return false;
    }

    private boolean continueSearch(char[][] board, boolean[][] visitedPoints, Point startPoint, String word, int n) {
        int xlim = board.length;
        int ylim = board[0].length;

        visitedPoints[startPoint.x][startPoint.y] = true;
        if (startPoint.x != 0 && board[startPoint.x - 1][startPoint.y] == word.charAt(n) && !visitedPoints[startPoint.x - 1][startPoint.y]) {
            if (n == word.length() - 1) {
                return true;
            }
            if (continueSearch(board, visitedPoints, new Point(startPoint.x - 1, startPoint.y), word, n + 1)) {
                return true;
            }
        }
        if (startPoint.x != xlim - 1 && board[startPoint.x + 1][startPoint.y] == word.charAt(n) && !visitedPoints[startPoint.x + 1][startPoint.y]) {
            if (n == word.length() - 1) {
                return true;
            }
            if (continueSearch(board, visitedPoints, new Point(startPoint.x + 1, startPoint.y), word, n + 1)) {
                return true;
            }
        }
        if (startPoint.y != 0 && board[startPoint.x][startPoint.y - 1] == word.charAt(n) && !visitedPoints[startPoint.x][startPoint.y - 1]) {
            if (n == word.length() - 1) {
                return true;
            }
            if (continueSearch(board, visitedPoints, new Point(startPoint.x, startPoint.y - 1), word, n + 1)) {
                return true;
            }
        }
        if (startPoint.y != ylim - 1 && board[startPoint.x][startPoint.y + 1] == word.charAt(n) && !visitedPoints[startPoint.x][startPoint.y + 1]) {
            if (n == word.length() - 1) {
                return true;
            }
            if (continueSearch(board, visitedPoints, new Point(startPoint.x, startPoint.y + 1), word, n + 1)) {
                return true;
            }
        }
        visitedPoints[startPoint.x][startPoint.y] = false;
        return false;
    }
}