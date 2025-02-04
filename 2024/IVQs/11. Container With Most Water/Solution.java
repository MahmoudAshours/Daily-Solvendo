class Solution {
    public int maxArea(int[] height) {
        int i = 0;
        int j = height.length - 1;
        int x = 0;
        int[] arr = new int[height.length - 1];
        while (i != j) {
            if (height[i] >= height[j]) {
                int area = height[j] * (j - i);
                arr[x] = area;
                x++;
                j--;
            } else {
                int area = height[i] * (j - i);
                arr[x] = area;
                x++;
                i++;
            }
        }
        int max = 0;
        for (int z : arr) {
            if (max <= z) max = z;
        }
        return max;
    }
}