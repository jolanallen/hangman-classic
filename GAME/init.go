package hangman
import (
	"fmt"
	"os"
	"bufio"
)



func (hangman *HANGMAN) Init() {
	hangman.wordIsGood = false
	hangman.testWord()
	hangman.wordlist()
	hangman.Start()
	hangman.hangman()
}

func (hangman *HANGMAN) testWord() {
	TesTab, err := os.ReadFile("utile/wordlist/words.txt")
	if err != nil {
		fmt.Println(err)
		hangman.wordIsGood = false
		
	} 
	if err == nil {
		hangman.wordIsGood = true
		
		

	}
	if hangman.wordIsGood {
		hangman.TabByte = TesTab
	}
}
