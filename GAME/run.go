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
		hangman.testword()
		hangman.win()
		hangman.gameOver()
		
		
		
		

	 }			
}


func (hangman *HANGMAN) letter() {
	var Reader = bufio.NewReader(os.Stdin)
	Rune, _, _:= Reader.ReadRune()
	if Rune >= 'a' && Rune <= 'z' {
		hangman.lettre = Rune 
		
		hangman.testLetter()
		
	} else {
		fmt.Println(hangman.motIconnu)
	
		fmt.Println("lettre entrée incorect!!")
			
		
	}
}
func (hangman *HANGMAN) testword() {
	//for i := range hangman.TabRune {
		var i int = 0
		if hangman.motIconnu[i] == hangman.TabWord[i] {
			i ++
			if i == len(hangman.MotAdeviner) {
				hangman.wordIsGood = true
			}
		} 
}
func (hangman *HANGMAN) testLetter() {
	
	for l := range hangman.MotAdeviner {
		if hangman.lettre == hangman.MotAdeviner[l] {
			hangman.lettreIsGood = true
			hangman.motIconnu[l] = string(hangman.lettre)
			if l == len(hangman.MotAdeviner) {
				break
			}
			
		} else if !(hangman.lettre == hangman.MotAdeviner[l]) {
			hangman.lettreIsGood = false
			
		} 
		
	}
	if hangman.lettreIsGood == true {
		hangman.hangman()
		fmt.Println(hangman.motIconnu)
		
		
	} 
	if hangman.lettreIsGood == false {
		hangman.erreur += 1
		hangman.hangman()
		fmt.Println(hangman.motIconnu)

		
	}
	

}

func (hangman *HANGMAN) randomWord() {
	rand.Seed(time.Now().UnixMilli())
	hangman.randomNb = rand.Intn(len(hangman.TabMots))
	hangman.Mot = hangman.TabMots[hangman.randomNb]
	fmt.Println(hangman.Mot)
	for _, l := range hangman.Mot {
		hangman.TabWord = append(hangman.TabWord, string(l))
	}
	fmt.Println(hangman.TabWord)
	hangman.MotAdeviner = []rune(hangman.Mot)
	fmt.Println(hangman.MotAdeviner)
	var inconnu string = "_"
	for i := 0; i < len(hangman.MotAdeviner); i++ {
		hangman.motIconnu = append(hangman.motIconnu, inconnu)
	}
	fmt.Println(hangman.motIconnu)

	

}





func (hangman *HANGMAN) win() {
	if hangman.wordIsGood && hangman.erreur < 6 {
		fmt.Print("fin du jeu vous avez gagner")
		hangman.IsRunning = false

	}
	
	
}

func (hangman *HANGMAN) gameOver() {
	if !(hangman.wordIsGood) && hangman.erreur == 6 {
		fmt.Print("fin du jeu le mot était : ")
		fmt.Println(hangman.Mot)
		hangman.IsRunning = false
	}
}