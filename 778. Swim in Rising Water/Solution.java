
import java.util.ArrayList;
import java.util.PriorityQueue;
class Solution {
    
    
    class Node implements Comparable<Node> {
        public static boolean[][] visited;
        int i , j;
        int cost;
        public Node(int i, int j, int cost){
            this.i = i;
            this.j = j;
            this.cost = cost;
        }
        @Override
        public int compareTo(Node o) {
            return this.cost - o.cost;
        }
        public void MarkVisited() {
            visited[i][j] = true;
        }


        public ArrayList<int[]> GetUnvisitedNeighbours(){
            var unvisitedNeighbours = new ArrayList<int[]>();
            
            if (i > 0 && !visited[i - 1][j]) unvisitedNeighbours.add(new int[]{i - 1, j});
            if (i < visited.length - 1 && !visited[i + 1][j]) unvisitedNeighbours.add(new int[]{i + 1, j});
            if (j > 0 && !visited[i][j - 1]) unvisitedNeighbours.add(new int[]{i, j - 1});
            if (j < visited[0].length - 1 && !visited[i][j + 1]) unvisitedNeighbours.add(new int[]{i, j + 1});
                
            
            return unvisitedNeighbours;
        }
    }
    
    PriorityQueue<Node> pq = new PriorityQueue<>();
    public int swimInWater(int[][] grid) {
        var cost = grid[0][0];
        Node.visited = new boolean[grid.length][grid.length];

        pq.add(new Node(0, 0, cost));

        while (!pq.isEmpty()){
            var node = pq.poll();
            // System.out.print(node.i );
            // System.out.println(node.j );
            cost = node.cost;
            if (node.i == grid.length - 1 && node.j == grid[0].length - 1){
                return node.cost;
            }
            var unvisitedNeighbours = node.GetUnvisitedNeighbours();
            for (var ij : unvisitedNeighbours) {
                var i = ij[0];
                var j = ij[1];
                pq.add(new Node(i, j, Math.max(cost , grid[i][j])));
            }
            node.MarkVisited();
        }
        return cost;
    }
    

    // public static void main(String[] args) {
    //     var sol = new Solution();
    //     var grid = new int[][]{
    //         {0, 1, 2, 3, 4},
    //         {24, 23, 22, 21, 5},
    //         {12, 13, 14, 15, 16},
    //         {11, 17, 18, 19, 20},
    //         {10, 9, 8, 7, 6}
    //     } ;
    //     var ans = sol.swimInWater(grid);
    //     System.out.println(ans);   
    // }
}