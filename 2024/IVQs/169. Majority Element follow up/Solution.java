import java.util.HashMap;

public class Solution {
    public int majorityElement(int[] nums) {
        int majorityNumber = nums[0];
        int majorityNumberCount = 1;
        System.out.println("Number :" + majorityNumber + "Count :" + majorityNumberCount);
        for (int i = 0; i < nums.length - 1; i++) {
            if (majorityNumber == nums[i + 1]) {
                majorityNumberCount++;
            } else {
                if (majorityNumberCount == 0) {
                    majorityNumber = nums[i];
                    majorityNumberCount = 1;
                } else majorityNumberCount--;
            }
            System.out.println("Number :" + majorityNumber + "Count :" + majorityNumberCount);
        }
        return majorityNumber;
    }
}
