/*
默认通道是**无缓冲**的，这意味着只有在对应的接收(<- CHAN)
通道准备好接收时，才允许进行发送(CHAN <-)。
可缓存通道允许在没有对应接收方的情况下，缓存限定数量的值。
*/
package main

import (
	"fmt"
)

func channelBuffering() {
	// 使用make创建一个通道，最多允许缓存两个值
	messages := make(chan string, 2)
	// 因为这个通道有缓冲区，即使没有一个对应的并发接收方，仍然可以发送这些值
	messages <- "buffered"
	messages <- "channel"
	// 可以接收者两个值
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
