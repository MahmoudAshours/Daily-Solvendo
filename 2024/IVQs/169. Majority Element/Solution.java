import java.util.HashMap;

public class Solution {
    public int majorityElement(int[] nums) {
        int max = Integer.MIN_VALUE;
        int maxIndex = 0;
        HashMap frequencyArray  = new HashMap();
        for (int i = 0; i < nums.length; i++) {
            if(frequencyArray.get(nums[i])== null){
                frequencyArray.put(nums[i],1);
            }else{
                frequencyArray.put(nums[i],Integer.parseInt(frequencyArray.get(nums[i]).toString())+1);
            }
        }

        for (Object element : frequencyArray.keySet()) {
            System.out.println((int) element);
            if (max <= (int)frequencyArray.get(element)){
                max = (int) frequencyArray.get(element);
                maxIndex = (int) element;
            }
        }
        return maxIndex;
    }
}
