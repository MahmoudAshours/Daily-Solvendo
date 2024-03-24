class Solution {
    public int countOdds(int low, int high) {
    int count = 0;
    if(high % 2 != 0){
    count+=1;
    high-=1;
    }
    if(low % 2 != 0 && low!=high){
         count+=1;
         low+=1;
    }
    count+= high % 2 != 0 && high % 2 != 0 ? ((high - low) / 2) + 1 : ((high - low) / 2);
    return count;
    }
}