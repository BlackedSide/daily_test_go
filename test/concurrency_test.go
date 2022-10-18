package test

import (
	"daily_test_go/service"
	"errors"
	"fmt"
	"reflect"
	"sync"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	time.Sleep(20 * time.Millisecond)
	if url == "https://www.qunar.com" {
		return true
	}
	return false
}

var websites = []string{
	"https://www.baidu.com",
	"https://www.qunar.com",
	"https://www.meituan.com",
}

func TestWebsiteChecker(t *testing.T) {

	actualRes := service.ConcurrentCheckWebsites(mockWebsiteChecker, websites)
	expectRes := map[string]bool{
		"https://www.baidu.com":   false,
		"https://www.qunar.com":   true,
		"https://www.meituan.com": false,
	}

	if !reflect.DeepEqual(actualRes, expectRes) {
		t.Errorf("got : '%v', want : '%v'", actualRes, expectRes)
	}
}

func BenchmarkConcurrentWebsites(b *testing.B) {

	for i := 0; i < b.N; i++ {
		service.ConcurrentCheckWebsites(mockWebsiteChecker, websites)
	}
}

func BenchmarkWebsites(b *testing.B) {

	for i := 0; i < b.N; i++ {
		service.LoopCheckWebsites(mockWebsiteChecker, websites)
	}
}

func TestChannel(t *testing.T) {
	c := make(chan int, 5)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
			fmt.Printf("等待消费……%d\n", i)
		}
		close(c)
	}()

	for i := 0; i < 5; i++ {
		got := <-c
		fmt.Printf("消费：%d\n", got)
	}
}

func TestRWChannel(t *testing.T) {
	c := make(chan int) //	chan	//读写

	go counter(c) //生产者
	printer(c)    //消费者

	fmt.Println("Done...")
}

func counter(out chan<- int) {
	defer close(out)
	for i := 0; i < 5; i++ {
		out <- i
	}
}

func printer(in <-chan int) {
	for num := range in {
		fmt.Printf("Print: %d\n", num)
	}
}

func TestRoutineErr(t *testing.T) {
	err := genErr()
	fmt.Println(err)
}

func genErr() error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	errChan := make(chan error, 2)

	go func() {
		defer wg.Done()
		fmt.Println("goroutine 1 output")
		//errChan <- errors.New("go routine 1 error")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("goroutine 2 output")
		errChan <- errors.New("go routine 2 error")
	}()
	wg.Wait()

	for e := range errChan {
		if e != nil {
			return e
		}
	}

	return nil
}
