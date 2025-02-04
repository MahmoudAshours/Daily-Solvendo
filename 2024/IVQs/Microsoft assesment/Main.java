public class Main {
    public static void main(String[] args) {

    }

    private static double sqrt(double n) {
        /*
         * n^2 = 25.434
         * 
         * 1.000
         * 1.250
         * 1.500
         * 
         * 1.999
         * 
         * 
         * pow(double , 2) => n
         * 
         * 
         * 
         * 0.00
         * 1.09
         * 2.25
         * 
         * 
         * 
         * 0
         * 
         * 5
         * 
         * 26
         */

        // for (double i = 0.00; i < n; i += 0.001) {
        // if (n - Math.pow(i, 2) < 0.0001)
        // return i;
        // }

        double l = 0;
        double r = n < 1 ? 1 : n;

        /** overflow */
        if (n < 0)
            return -1;

        // 0.5 * 2 = 0.25

        /*
         * 
         * 
         * 1/4 * 2 = 2/4
         */
        /*
         * 
         * 
         * 
         */
        /**
         * 
         * n < 1
         * 
         * 
         */
        while (l <= r) {

            double mid = (l + r) / 2;

            System.out.println(l + " " + r + " " + mid);

            if (Math.pow(mid, 2) < n) {
                l = mid;
            } else if (Math.pow(mid, 2) > n) {
                r = mid;
            } else {
                return mid;
            }
        }
        return -1;

    }

    /**
     * 
     * 
     * 999999999999 * 999999 = long
     * 
     * 
     * 25.4
     * 
     * func(25){
     * 
     * n ^ 2 = 25
     * 
     * n * n = 25
     * n^2 = 25
     * 
     * 25
     * i = 0;
     * for(i < n/2){
     * 
     * power(i,2) == n ? i;
     * return 5;
     *
     * }
     * 
     * return n*n;
     * }
     */
}
