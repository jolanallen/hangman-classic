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
	TabMots, err := os.Open("utile/wordlist/words.txt")
	if err != nil {
		fmt.Println(err)
	} 
	if err == nil {
		defer TabMots.Close()
		hangman.wordIsGood = true
	}
}
