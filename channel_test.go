package goroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// Basic Channel
func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Dimas Eka Adinandra"
		fmt.Println("Finish Send Data to Channel")
	}()

	data := <-channel
	fmt.Println(data)
}

// Channel as Params
func GiveMeResponse(request chan string) {
	time.Sleep(2 * time.Second)
	request <- "Dimas Eka Adinandra"
}

func TestChannelAsParams(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)
}

// Channel In and Channel Out
func OnlyIn(request chan<- string) {
	time.Sleep(2 * time.Second)
	request <- "Dimas Eka Adinandra"
}

func OnlyOut(request <-chan string) {
	data := <-request
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

// Buffered Channel
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	channel <- "Dimas"
	channel <- "Eka"
	channel <- "Adinandra"

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()
	time.Sleep(5 * time.Second)
}

// Range Channel

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for i := range channel {
		fmt.Println(i)
	}
}

// Select Channel
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	counter := 0

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	for {
		select {
		case data := <-channel1:
			fmt.Println("Data Dari Channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data Dari Channel 1", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}
}

// Default Select
func TestDefaultSelect(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	counter := 0

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	for {
		select {
		case data := <-channel1:
			fmt.Println("Data Dari Channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data Dari Channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu Data")
		}

		if counter == 2 {
			break
		}
	}
}
