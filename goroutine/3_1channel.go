package goroutine

import (
	"fmt"
	"runtime"
)

func printMessage(msg chan string) {
	fmt.Println(<-msg)
}

// simple channel
func main() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)

	var names = []string{"Imam Permana", "Budi", "Anya Geraldine"}

	var sayHelloTo = func(who string) {
		// send data ke channel 'message'
		var data = fmt.Sprintf("hello %s", who)
		messages <- data
	}

	for _, each := range names {
		go sayHelloTo(each)
	}

	for i := 0; i < len(names); i++ {
		// fungsi ini sebagai receive data dari channel 'messages'
		printMessage(messages)
	}
}

func SayHelloTo(name string, msg chan string) {
	// send data ke parameter channel 'msg' yabg memiliki referensi dari channel 'messages'
	msg <- fmt.Sprintf("hello %s", name)
}

// channel as parameter
func TestChannelAsParameter() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)

	var names = []string{"Imam Permana", "Budi", "Anya Geraldine"}

	for _, each := range names {
		// fungsi ini sebagai send data ke channel 'messages'
		go SayHelloTo(each, messages)
	}

	for i := 0; i < len(names); i++ {
		// fungsi ini sebagai receive data dari channel 'messages'
		fmt.Println(<-messages)
	}
}
