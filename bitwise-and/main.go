package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

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

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    t := int32(tTemp)

    for tItr := 0; tItr < int(t); tItr++ {
        nk := strings.Split(readLine(reader), " ")

        nTemp, err := strconv.ParseInt(nk[0], 10, 64)
        checkError(err)
        n := int32(nTemp)

        kTemp, err := strconv.ParseInt(nk[1], 10, 64)
        checkError(err)
        k := int32(kTemp)
        
        var result, num, i, j int32
        for i = 1; i < n; i++ {
            for j = i + 1; j <= n; j++ {
                num = i & j
                if num > result && num < k{
                    result = num
                }
            }
        }
        fmt.Println(result)
    }
}
