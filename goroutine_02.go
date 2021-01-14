package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	//  create channel for DB
	DbChannel := make(chan string)
	RedisChannel := make(chan string)

	go connectDB(DbChannel)
	go connectRedis(RedisChannel)

	DbClinet := <-DbChannel // store the return values from channels to variable
	RedisClinet := <-RedisChannel

	fmt.Println(DbClinet, RedisClinet)

	fmt.Println("Total time required ", time.Since(now))
}

func connectDB(ch chan string) {
	time.Sleep(time.Second * 2)
	// fmt.Println("DB Connected")

	ch <- "DB connected"
}

func connectRedis(ch chan string) {
	time.Sleep(time.Second * 2)
	// fmt.Println("Redis")
	ch <- "Redis conneted"

}

// =======================================
// writing both the functions as a go routine leads to terminate the main function, so no one will be execute.
//  Basically main function should have to wait to complete routines defined
//  to avoid this intrdoduce channels
