package hangman
import (
	"fmt"
	"os"
	// "bufio"
)



func (hangman *HANGMAN) Init() {
	hangman.testWord()
	hangman.wordlist()
	hangman.Start()
	hangman.hangman()
}

func (hangman *HANGMAN) testWord() {
	TesTabs, err := os.Open("utile/wordlist/words.txt")
	if err != nil {
		fmt.Println(err)
		hangman.wordIsGood = false
		defer TesTabs.Close()
	} 
	if err == nil {
		defer TesTabs.Close()
		hangman.wordIsGood = true
	}
}
