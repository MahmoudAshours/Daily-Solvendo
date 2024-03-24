package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the freqQuery function below.

func freqQuery(queries [][]int32) []int32 {

    output:= make([]int32, 0)
    outMap:= make(map[int32]int32, 0)

    for _, v := range queries {
        switch v[0] {

        case 1:
            outMap[v[1]]++

        case 2:

            _ , present := outMap[v[1]]
        if len(outMap) != 0{
            if present {
                outMap[v[1]]--
                if outMap[v[1]] == 0{
                    delete(outMap,v[1])
                }
            }
    }
        case 3:
            isPresent := false
            if len(outMap) != 0{
            for _, m := range outMap {
                  if m == v[1]{
                    isPresent = true
                    break
                  }
              }
            }

              if isPresent{
        
               output = append(output,1)
              
              } else {
              
               output = append(output,0)  
              
              }  
      
        }
    }

    return output
}
func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create("test")
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    q := int32(qTemp)

    var queries [][]int32
    for i := 0; i < int(q); i++ {
        queriesRowTemp := strings.Split(strings.TrimRight(readLine(reader)," \t\r\n"), " ")

        var queriesRow []int32
        for _, queriesRowItem := range queriesRowTemp {
            queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
            checkError(err)
            queriesItem := int32(queriesItemTemp)
            queriesRow = append(queriesRow, queriesItem)
        }

        if len(queriesRow) != 2 {
            panic("Bad input")
        }

        queries = append(queries, queriesRow)
    }

    ans := freqQuery(queries)

    for i, ansItem := range ans {
        fmt.Fprintf(writer, "%d", ansItem)

        if i != len(ans) - 1 {
            fmt.Fprintf(writer, "\n")
        }
    }

    fmt.Fprintf(writer, "\n")

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
