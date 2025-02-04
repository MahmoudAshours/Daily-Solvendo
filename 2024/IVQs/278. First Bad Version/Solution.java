/* The isBadVersion API is defined in the parent class VersionControl.
      boolean isBadVersion(int version); */

public class Solution extends VersionControl {
    public int firstBadVersion(int n) {
        int hi = n;
        int ans = Integer.MAX_VALUE;
        int res = ceil(n, 1, hi, ans);
        return res;
    }

    public int ceil(int n, int lo, int hi, int ans) {
        if (lo > hi) return ans;
        int mid = lo + (hi - lo) / 2;
        if (isBadVersion(mid)) return ceil(n, lo, mid - 1, mid);
        else return ceil(n, mid + 1, hi, ans);
    }
}