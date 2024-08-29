package belajar_golang_goroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Ferdian Alvanza"
		fmt.Println("Selesai Mengirim data ke Channel")
	}()
	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}
func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Ferdian Alvanza"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

// HANYA UNTUK MENGIRIM DATA KE CHANNEL
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Ferdian Alvanza"
}

// HANYA UNTUK MENERIMA DATA DARI CHANNEL
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Moch"
		channel <- "Ferdian"
		channel <- "Alvanza"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")

	fmt.Println(cap(channel)) //MELIHAT JUMLAH DARI BUFFER
	fmt.Println(len(channel)) // MELIHAT PANJANG BUFFER

}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data", data)
	}
	fmt.Println("Selesai")

}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	channel3 := make(chan string)
	defer close(channel1)
	defer close(channel2)
	defer close(channel3)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)
	go GiveMeResponse(channel3)

	counter := 0
	for {

		select {
		case data := <-channel1:
			fmt.Println("Data dari Channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari Channel 2", data)
			counter++
		case data := <-channel3:
			fmt.Println("Data dari Channel 3", data)
			counter++
		}
		if counter == 3 {
			break
		}
	}
	// select {
	// case data := <-channel1:
	// 	fmt.Println("Data dari Channel 1", data)
	// case data := <-channel2:
	// 	fmt.Println("Data dari Channel 2", data)
	// }

	// select {
	// case data := <-channel1:
	// 	fmt.Println("Data dari Channel 1", data)
	// case data := <-channel2:
	// 	fmt.Println("Data dari Channel 2", data)
	// }
}

func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {

		select {
		case data := <-channel1:
			fmt.Println("Data dari Channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari Channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu Data")
		}
		if counter == 2 {
			break
		}
	}
}
