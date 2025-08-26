package main

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

func fibo(n int) int {
	if n <= 1 {
		return n
	}

	return fibo(n-1) + fibo(n-2)
}

func work1() {
	cnt := 0
	for {
		cnt++
		fibo(20)
		sha256_busy(10000)
		fmt.Printf("golang................work1 cnt: %d\n", cnt)
	}
}

func work2() {
	cnt := 0
	for {
		cnt++
		fibo(20)
		sha256_busy(10000)
		fmt.Printf("golang................work2 cnt: %d\n", cnt)
	}
}

func sha256_busy(n int) int {
	if n < 1 {
		return 0
	}
	s := fmt.Sprintf("This is a short sha256 test string%d", n)
	h := sha256.New()
	h.Write([]byte(s))
	_ = h.Sum(nil)
	return sha256_busy(n - 1)
}

func work_busy() {
	for {
		sha256_busy(10000)
	}
}

func main() {
	fmt.Println("start to run")
	var wg sync.WaitGroup
	wg.Add(1)
	go work1()
	go work2()
	for i := 0; i < 20; i++ {
		go work_busy()
	}
	wg.Wait()
}
