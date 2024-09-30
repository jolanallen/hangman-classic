package hangman

import (
	"bufio"
	"fmt"
	"os"
)

//hangman "hangman/GAME"

func (hangman *HANGMAN) wordlist() {
	wordFile, err := os.Open("utile/wordlist/words.txt")
	if err != nil {
		fmt.Println(err)
		hangman.wordIsGood = false
		defer wordFile.Close()
	}
	if err == nil {
		hangman.wordIsGood = true
	}
	if hangman.wordIsGood {
		scanner := bufio.NewScanner(wordFile)
		for scanner.Scan() { 
			hangman.Mot = scanner.Text()
			hangman.TabMots = append(hangman.TabMots,  hangman.Mot)
			fmt.Println(hangman.Mot)
		}
		
	}
}



