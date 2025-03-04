func checkPowersOfThree(n int) bool {
    /**
    12 / 3 => 0
    4 / 3 => 1
    1 / 3 => 1
    
    21 / 3 => 0
    7 / 3 => 1
    2 / 3 => 0

    **/
    pow := ""
    for n > 0{
        bit := n % 3
        pow += string('0'+bit)
        n/=3
    }
    if strings.Contains(pow,"2"){
        return false
    }
    return true
}