package main

import (
	Mail "GoMap/mail"
	"fmt"
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	fmt.Println("Starting...")
	user1 := Mail.NewUser()
	user1.Login("dom@somersetsdade.org", "oazamksgdrasmrmw")

	messages := user1.EmailSearch("randy", "")
	wg.Add(1)

	mail := Mail.EmailChecker(messages, &wg)

	fmt.Println()
	log.Println("Logged in...")
	fmt.Println(mail)

}
