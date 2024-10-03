package hangman

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	
)




func (hangman *HANGMAN) Run() {
	hangman.randomWord()
	 for hangman.IsRunning {
		hangman.letter()
		
		
		
		

	 }			
}


func (hangman *HANGMAN) letter() {
	var Reader = bufio.NewReader(os.Stdin)
	Rune, _, _:= Reader.ReadRune()
	if Rune >= 'a' && Rune <= 'z' {
		hangman.lettre = Rune 
		hangman.testLetter()
		
	} else {
		
		fmt.Println("lettre entrée incorect!!")
			
		
	}
}
func (hangman *HANGMAN) testLetter() {
	for l := range hangman.MotAdeviner {
		if hangman.lettre == hangman.MotAdeviner[l] {
			hangman.lettreIsGood = true
			hangman.motIconnu[l] = string(hangman.lettre)
			fmt.Println(hangman.motIconnu)

		} 
	}
	if !hangman.lettreIsGood {
		hangman.erreur =+ 1
		
	}
}

func (hangman *HANGMAN) randomWord() {
	rand.Seed(time.Now().UnixMilli())
	hangman.randomNb = rand.Intn(len(hangman.TabMots))
	hangman.Mot = hangman.TabMots[hangman.randomNb]
	fmt.Println(hangman.Mot)
	hangman.MotAdeviner = []rune(hangman.Mot)
	fmt.Println(hangman.MotAdeviner)
	var inconnu string = "_"
	for i := 0; i < len(hangman.MotAdeviner); i++ {
		hangman.motIconnu = append(hangman.motIconnu, inconnu)
	}
	fmt.Println(hangman.motIconnu)

	

}



func(hangman *HANGMAN) Stop() {
	hangman.IsRunning = false
	
}

// func (hangman *HANGMAN) win() {
// 	fmt.Print("fin du jeu vous avez gagner")
	
// }

func (hangman *HANGMAN) gameOver() {
	fmt.Print("fin du jeu le mot était : ")
	fmt.Println(hangman.Mot)
	
	
}