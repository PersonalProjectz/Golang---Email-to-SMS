package Mail

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"sync"

	imap "github.com/emersion/go-imap"
	client "github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/mail"
)

type User struct {
	Ids []uint32
	// ReceiptsArray []Receipts
	ClientSession *client.Client
	// ByteData []byte
}

func NewUser() *User {
	var newUser User
	return &newUser
}

func (user *User) Login(username string, password string) {
	//watch := stopwatch.Start()
	log.Println("Connecting to server...")

	// Connect to server
	var err error
	user.ClientSession, err = client.DialTLS("imap.gmail.com:993", nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected")

	// Don't forget to logout
	//defer user.clientSession.Logout()

	// Login
	if err := user.ClientSession.Login(username, password); err != nil {

		log.Fatal(err)
		os.Exit(1)
	}
	log.Println("Logged in")

	_, err = user.ClientSession.Select("INBOX", false)
	if err != nil {
		log.Fatal(err)
		return
	}

}

func (user *User) EmailSearch(from string, body string) <-chan *imap.Message {
	newSearchCriteria := imap.NewSearchCriteria()
	newSearchCriteria.Header.Add("From", from)
	newSearchCriteria.WithoutFlags = []string{"\\Seen"}

	if body != "" {
		newSearchCriteria.Body = []string{body}
	}

	ids, err := user.ClientSession.Search(newSearchCriteria)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	if len(ids) > 0 {
		seqset := new(imap.SeqSet)
		seqset.AddNum(ids...)

		var section imap.BodySectionName
		messages := make(chan *imap.Message)
		go func() {
			if err := user.ClientSession.Fetch(seqset, []imap.FetchItem{section.FetchItem()}, messages); err != nil {
			}
		}()

		return messages
	}
	return nil
}

func EmailChecker(messages <-chan *imap.Message, wg *sync.WaitGroup) string {
	var mess string
	defer wg.Done()
	if messages == nil {
		return mess
	}

	for msg := range messages {
		if msg == nil {
			log.Fatal("Server didn't returned message")
			return mess
		}
		var section imap.BodySectionName
		r := msg.GetBody(&section)
		if r == nil {
			log.Fatal("Server didn't returned message body")
		}

		// Create a new mail reader
		mr, err := mail.CreateReader(r)
		if err != nil {
			log.Fatal(err)
		}
		for {
			p, err := mr.NextPart()
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}

			switch h := p.Header.(type) {
			case *mail.InlineHeader:
				b, _ := ioutil.ReadAll(p.Body)
				log.Printf("Got text: %v\n", string(b))
				strings := string(b)
				mess = strings
			case *mail.AttachmentHeader:
				filename, _ := h.Filename()
				log.Printf("Got attachment: %v\n", filename)
			}
		}

	}
	return mess
}
