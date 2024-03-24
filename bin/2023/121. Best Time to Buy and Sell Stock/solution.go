package main

func main(){
print(maxProfit([]int{7,1,5,3,6,4}))
}
func maxProfit(prices []int) int {
    var profit = 0
    var minPrice = prices[0]
    
    for i := 1; i < len(prices); i++ {
        // If we find any price which is lower than the current minPrice
        // update the minPrice
        if prices[i] < minPrice {
            minPrice = prices[i]
        } else if (prices[i] - minPrice) > profit {
            // If diff of current stock with minPrice is greater
            // update the profit
            profit = prices[i] - minPrice
        }
    }
    
    return profit
}
