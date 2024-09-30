package hangman

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

//hangman "hangman/GAME"

func (hangman *HANGMAN) wordlist() {
	wordFile, err := os.Open("utile/wordlist/words.txt")
	if err != nil {
		fmt.Println(err)
		hangman.wordIsGood = false
		hangman.IsRunning = false
		defer wordFile.Close()
	}
	if err == nil {
		hangman.wordIsGood = true
		scanner := bufio.NewScanner(wordFile)
		for scanner.Scan() { 
			hangman.Mot = scanner.Text()
			hangman.TabMots = append(hangman.TabMots,  hangman.Mot)
		}	
		hangman.randomWord()
	}
}

func (hangman *HANGMAN) randomWord() {
	rand.Seed(time.Now().UnixMilli())
	hangman.randomNb = rand.Intn(38)
	fmt.Println(hangman.randomNb)
	

}

