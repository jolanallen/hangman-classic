package hangman

import (
	"fmt"
	"math/rand"
	"time"
)




func (hangman *HANGMAN) Run() {
	hangman.testWord()
	if hangman.wordIsGood {
		hangman.lMot = len(hangman.MotAdeviner)
		fmt.Println(hangman.MotAdeviner)
		for i := 1; i <= hangman.lMot; i++ {
			fmt.Print("_  ")
		}
		fmt.Println(" ")
		hangman.hangman()
		// for hangman.IsRunning {
			



		// }	
	}		
}
func (hangman *HANGMAN) testWord() {
	var TabRune = []rune(hangman.MotAdeviner)
	hangman.randomWord()
	hangman.MotAdeviner = (hangman.TabMots[hangman.randomNb])
	for tab := range TabRune {
		 var Rune = TabRune[tab]
		if Rune <= 'z' && Rune >= 'a' {
			hangman.wordIsGood = true
		} else {
			hangman.wordIsGood = false
			fmt.Println("mot invalid")
			hangman.Stop()
			break
			
		}


	}
	
	
}
func (hangman *HANGMAN) randomWord() {
	rand.Seed(time.Now().UnixMilli())
	hangman.randomNb = rand.Intn(38)


}


func(hangman *HANGMAN) Stop() {
	hangman.IsRunning = false
	
}

func (hangman *HANGMAN) win() {
	fmt.Print("fin du jeu vous avez gagner")
	
}

func (hangman *HANGMAN) gameOver() {
	fmt.Print("fin du jeu le mot était ")
	
	
}