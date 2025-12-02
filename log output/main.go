package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 16)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	storedString := string(b)

	fmt.Printf("🚀 App Started. ID stored in memory: %s\n", storedString)

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for t := range ticker.C {
		log.Printf("[%s] %s\n", t.Format(time.DateTime), storedString)
	}
}
