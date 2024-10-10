class Solution {
    public boolean uniqueOccurrences(int[] arr) {
        int[] numberOfOccurrences = new int[2001];
        boolean[] used = new boolean[1001];

        for (int i : arr) {
            numberOfOccurrences[1000 + i]++;
        }

        for (int i : arr) {
            int count = numberOfOccurrences[1000 + i]; // идем по исходному массиву
            numberOfOccurrences[1000 + i] = 0; // достаем кол-во повторений и зануляем его
            if (count > 0) {
                if (used[count]) return false;
                else used[count] = true;
            }
        }
        return true;
    }
}