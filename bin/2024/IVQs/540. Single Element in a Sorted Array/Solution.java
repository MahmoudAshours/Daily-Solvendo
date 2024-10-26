class Solution {
    public int singleNonDuplicate(int[] nums) {
        // log n -> bs
        int loIndex = 0;
        int hiIndex = nums.length - 1;
        while (loIndex < hiIndex) {
            int mid = loIndex + (hiIndex - loIndex) / 2;
            boolean isEven = (hiIndex - mid) % 2 == 0;
            if (nums[mid] == nums[mid - 1]) {
                if (isEven) {
                    hiIndex = mid - 2;
                } else {
                    loIndex = mid + 1;
                }
            } else if (nums[mid] == nums[mid + 1]) {
                if (isEven) {
                    loIndex = mid + 2;
                } else {
                    hiIndex = mid - 1;
                }
            } else {
                return nums[mid];
            }
        }
        return nums[loIndex];
    }
}