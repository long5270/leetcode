package main

import (
	"fmt"
	"os/exec"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	t1 := time.Now()
	for i := 0; i < 200; i++ {
		go Exec(&wg, i)
	}
	wg.Wait()
	fmt.Println(time.Since(t1))
}

func Exec(wg *sync.WaitGroup, i int) {
	wg.Add(1)
	cmd := exec.Command("./client")
	ret, _ := cmd.CombinedOutput()
	fmt.Println(i)
	fmt.Println(string(ret))
	wg.Done()
}

func isValid(s string) bool {
	length := len(s)
	if length%2 > 0 {
		return false
	}
	notMatch := []int{}
	matchMap := make(map[int]bool, 10)
	sBytes := []byte(s)
	keyMap := map[byte]byte{40: 41, 41: 40, 91: 93, 93: 91, 123: 125, 125: 123}

	for i := 0; i < length; i++ {
		tmp := sBytes[i]
		if !matchMap[i] {
			if i+1 != length {
				if sBytes[i+1] != keyMap[tmp] {
					if len(notMatch) > 0 {
						lenNotMatch := len(notMatch) - 1
						if keyMap[tmp] == sBytes[notMatch[lenNotMatch]] {
							matchMap[i] = true
							matchMap[notMatch[lenNotMatch]] = true
							notMatch = notMatch[:lenNotMatch]
						} else {
							notMatch = append(notMatch, i)
						}
					} else {
						notMatch = append(notMatch, i)
					}
				} else {
					matchMap[i] = true
					matchMap[i+1] = true
				}
			} else {
				lenNotMatch := len(notMatch) - 1
				if keyMap[tmp] == sBytes[notMatch[lenNotMatch]] {
					matchMap[i] = true
					matchMap[notMatch[lenNotMatch]] = true
					notMatch = notMatch[:lenNotMatch]
				} else {
					notMatch = append(notMatch, i)
				}
			}
		}
	}
	if len(notMatch) > 0 {
		return false
	}
	return true
}
