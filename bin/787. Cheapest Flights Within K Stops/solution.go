func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
    var bestTimeToDst = make([]int, n, n)
    for i:=0; i < n; i++ {
        bestTimeToDst[i] = math.MaxInt32
    }
    bestTimeToDst[src] = 0
    for i:=0; i < min(k + 1, n - 1); i++ {
        var changes = 0
        var tmp = make([]int, n, n)
        copy(tmp, bestTimeToDst)
        for _, flightData := range(flights) {
            var flightSource, flightDest, flightPrice = flightData[0], flightData[1], flightData[2]
            // if there is known path
            if(bestTimeToDst[flightSource] != math.MaxInt32) {
                if(bestTimeToDst[flightSource] + flightPrice < tmp[flightDest]) {
                    tmp[flightDest] = bestTimeToDst[flightSource] + flightPrice
                    changes++
                }
            }
        }
        copy(bestTimeToDst, tmp)
        if(changes == 0) {
            break
        }
    }
    if(bestTimeToDst[dst] == math.MaxInt32) {
        return -1
    }
    return bestTimeToDst[dst]
}

func min(a int, b int) int {
    if(a < b) {
        return a
    } else {
        return b
    }
}