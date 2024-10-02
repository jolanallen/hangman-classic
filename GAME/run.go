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
		hangman.testLetter()
		
		

	 }			
}


func (hangman *HANGMAN) letter() {
	var Reader = bufio.NewReader(os.Stdin)
	Rune, _, _:= Reader.ReadRune()
	if Rune <= 'a' || Rune >= 'z' {
		fmt.Println("lettre entrée incorect!!")
	} else {
		hangman.lettre = Rune 	
		
	}
}
func (hangman *HANGMAN) testLetter() {
	

}

func (hangman *HANGMAN) randomWord() {
	rand.Seed(time.Now().UnixMilli())
	hangman.randomNb = rand.Intn(len(hangman.TabMots))
	hangman.Mot = hangman.TabMots[hangman.randomNb]
	fmt.Println(hangman.Mot)
	hangman.MotAdeviner = []rune(hangman.Mot)
	fmt.Println(hangman.MotAdeviner)
	var lettreInconu string = " _ "
	for i := 0; i < len(hangman.MotAdeviner); i++ {
		hangman.motIconnu = append(hangman.motIconnu, lettreInconu)
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
	fmt.Println(hangman.MotAdeviner)
	
	
}