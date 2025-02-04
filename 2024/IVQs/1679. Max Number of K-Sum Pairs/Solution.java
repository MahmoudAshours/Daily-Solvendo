class Solution {
    public int maxOperations(int[] nums, int k) {
        /**
        
        3 2 3 4 5 , k = 7
        
        sort

        2 3 3 4 5 , k = 7
        
         */
        Arrays.sort(nums);
        int i = 0;
        int j = nums.length - 1;
        int count = 0;
        while(i < j){
            int sum = nums[i] + nums[j];
            if( sum == k){
                count++;
                i++;
                j--;
            }else if (sum > k){
                j--;
            }else{
                i++;
            }
        }
        return count;
    }
}