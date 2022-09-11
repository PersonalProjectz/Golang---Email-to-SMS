package main

import (
	Message "GoMap/Twilio"
	Mail "GoMap/mail"
	"fmt"
	"sync"
	"time"
)

func main() {
	var mail string
	var x int = 1
	var wg sync.WaitGroup
	fmt.Println("Starting...")
	user1 := Mail.NewUser()
	
	
	for x < 2 {
		user1.Login("dom@somersetsdade.org", "xxxxxx") // EMAIL AND APP PASSWORD HERE
		wg.Add(1)
		time.Sleep(10 * time.Second)
		messages := user1.EmailSearch("Dominick", "")            // EMAIL CRITERIA HERE
		mail = Mail.EmailChecker(messages, &wg)
		
		if(len(mail) != 0) {
		Message.Connection("3058332758", mail)                    // YOUR NUMBER GOES HERE
		}

		fmt.Println("itterated")
	}
}
