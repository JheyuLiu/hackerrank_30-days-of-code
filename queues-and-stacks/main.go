package main

import (
	"fmt"
	"bufio"
	"os"
)

type stack struct {
	sdata []byte
    len int
}

type queue struct {
	qdata []byte
	start int
	last int
}

func (s *stack) pushCharacter(c byte) {
	s.sdata[s.len] = c
	s.len++
}

func (q *queue) enqueueCharacter(c byte) {
	q.qdata[q.last] = c
	q.last++
}

func (s *stack) popCharacter() byte {
	s.len = s.len - 1

	return s.sdata[s.len]
}

func (q *queue) dequeueCharacter() byte {
    if q.start == q.last {
    	q.start = 0
    	q.last = 0
    }
	c := q.qdata[q.start]
	q.start++

	return c
}

func main() {
	isPalindrome := true
    scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
    str := scanner.Text()
    length := len(str)

    s := stack{}
    s.sdata = make([]byte, length)

    q := queue{}
    q.qdata = make([]byte, length)

    for i := 0; i < length; i++ {
    	s.pushCharacter(str[i])
    	q.enqueueCharacter(str[i])
    }

    for i := 0; i < length / 2; i++ {
    	if s.popCharacter() != q.dequeueCharacter() {
    		isPalindrome = false
    		break;
    	}
    }

    if isPalindrome {
    	fmt.Println("The word, ", str, ", is a palindrome")
    } else {
    	fmt.Println("The word, ", str, ", is not a palindrome")
    }
}