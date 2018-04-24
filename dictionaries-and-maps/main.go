package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var num int
	fmt.Scan(&num)

    infoMap := make(map[string]int)
    var name string
    var phone int
	for i := 0; i < num; i++ {
		fmt.Scan(&name)
		fmt.Scan(&phone)
		infoMap[name] = phone
	}

    scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		name = scanner.Text()
		_, exist := infoMap[name]
		if exist {
			fmt.Printf("%s=%d\n", name, infoMap[name])
		} else {
			fmt.Println("Not found")
		}
	}
}