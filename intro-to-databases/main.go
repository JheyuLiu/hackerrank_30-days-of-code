package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "regexp"
    "sort"
)

var IsLetter = regexp.MustCompile(`^[a-z]+$`).MatchString
var IsEmail = regexp.MustCompile(`.+@gmail.com`).MatchString


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

    NTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    N := int32(NTemp)
    var nameList []string

    for NItr := 0; NItr < int(N); NItr++ {
        firstNameEmailID := strings.Split(readLine(reader), " ")
        firstName := firstNameEmailID[0]
        emailID := firstNameEmailID[1]
        if !IsEmail(emailID) {
            continue
        }

        nameList = append(nameList, strings.ToLower(firstName))
    }
    sort.Strings(nameList)
    for _, name := range(nameList) {
        fmt.Println(name)
    }
}
