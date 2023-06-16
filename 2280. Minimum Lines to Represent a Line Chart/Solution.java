class Solution {
    public boolean isSameLine(int[] point1, int[] point2, int[] point3) {
        int crossProduct = (point2[0] - point1[0]) * (point3[1] - point1[1]) - (point2[1] - point1[1]) * (point3[0] - point1[0]);
        return crossProduct == 0;
    }

    public int minimumLines(int[][] stockPrices) {
        if (stockPrices.length == 1) {
            return 0;
        }
        if (stockPrices.length == 2) {
            return 1;
        }
        Arrays.sort(stockPrices, Comparator.comparingInt(a -> a[0]));
        int count = 1;
        int[] lastPoint = stockPrices[0];
        for (int i = 1; i < stockPrices.length - 1; i++) {
            if (!isSameLine(lastPoint, stockPrices[i], stockPrices[i + 1])) {
                count++;
                lastPoint = stockPrices[i];
            }
        }
        return count;
    }
}