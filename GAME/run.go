package hangman

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	
)




func (hangman *HANGMAN) Run() {
	hangman.testWord()
	if hangman.wordIsGood {
		hangman.lMot = len(hangman.MotAdeviner)
		for i := 1; i <= hangman.lMot; i++ {
			fmt.Print("_  ")
		}
		fmt.Println(" ")
		hangman.erreur = 0
		
		for hangman.IsRunning {
			hangman.letter()
			




		}	
	}		
}
func (hangman *HANGMAN) letter() {
	var Reader = bufio.NewReader(os.Stdin)
	Rune, _, err := Reader.ReadRune()
	if err != nil {
		fmt.Println(err)
	}
	if Rune <= 'a' || Rune >= 'z' {
		hangman.lettreIsGood = false
		fmt.Println("lettre entrée incorect. Veuliiez entrée un charactére de l'alphabet sans accent.")
	} else {
		Rune = hangman.lettre
		fmt.Println(hangman.lettre)
	}

	
	
}
func (hangman *HANGMAN) testLetter() {
	

}
func (hangman *HANGMAN) testWord() {
	hangman.TabRune = []rune(hangman.MotAdeviner)
	hangman.randomWord()
	hangman.MotAdeviner = (hangman.TabMots[hangman.randomNb])
	for tab := range hangman.TabRune {
		 var Rune = hangman.TabRune[tab]
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
	hangman.randomNb = rand.Intn(len(hangman.TabMots))
	

}



func(hangman *HANGMAN) Stop() {
	hangman.IsRunning = false
	
}

func (hangman *HANGMAN) win() {
	fmt.Print("fin du jeu vous avez gagner")
	
}

func (hangman *HANGMAN) gameOver() {
	fmt.Print("fin du jeu le mot était : ")
	fmt.Println(hangman.MotAdeviner)
	
	
}