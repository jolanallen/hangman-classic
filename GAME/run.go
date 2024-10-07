package hangman

import (
	"bufio"
	"fmt"

	"math/rand"
	"os"
	"time"
)




func (hangman *HANGMAN) Run() {
	hangman.wordIsGood = false
	
	hangman.randomWord()
	 for hangman.IsRunning {
		hangman.letter()
		hangman.testword()
		fmt.Println(hangman.MotAdeviner)
		hangman.hangman()
		if hangman.wordIsGood {
			hangman.win()
		}
	 }			
}


func (hangman *HANGMAN) letter() {
	var Reader = bufio.NewReader(os.Stdin)
	l, _, _:= Reader.ReadRune()
	if l >= 'a' && l <= 'z' {
		hangman.lettre = string(l)
		hangman.testLetter()
		fmt.Println(hangman.motIconnu)
		
	} else {
		fmt.Println(hangman.motIconnu)
		fmt.Println("lettre entrée incorect!!")
			
		
	}
}
func (hangman *HANGMAN) testword() {
	hangman.wordIsGood = true
	for l := range hangman.motIconnu {
		if !(hangman.motIconnu[l] == hangman.MotAdeviner[l]) {
			hangman.wordIsGood = false
			break
		}
	}
}
func (hangman *HANGMAN) testLetter() {
	hangman.lettreIsGood = false
	for l := range hangman.MotAdeviner {
		if hangman.lettre == hangman.MotAdeviner[l] {
			hangman. lettreIsGood = true
			fmt.Println(hangman.motIconnu)
			hangman.motIconnu[l] = hangman.lettre
		}
		
	}
	if !hangman.lettreIsGood {
		hangman.erreur++
		fmt.Println("Lettre incorrecte")
	}

	
	
}

func (hangman *HANGMAN) randomWord() {
	rand.Seed(time.Now().UnixMilli())                 /// EVERYONE bien se renseigner sur la library RAND !!!!!!!!§!!!!
	hangman.randomNb = rand.Intn(len(hangman.TabMots))
	hangman.Mot = hangman.TabMots[hangman.randomNb]
	for _, char := range hangman.Mot {
		hangman.MotAdeviner = append(hangman.MotAdeviner, string(char))
	}
	fmt.Println(hangman.MotAdeviner)
	var l string = "_"
	for i :=  0; i < len(hangman.Mot); i++ {
		hangman.motIconnu = append(hangman.motIconnu, l)
	}
	fmt.Println(hangman.motIconnu)

}





func (hangman *HANGMAN) win() {
	for i := 0; i <= 20; i++ {
		fmt.Println("fin du jeu vous avez gagner le mot était bien", hangman.Mot)
	}
	hangman.IsRunning = false
	
	
	
}

func (hangman *HANGMAN) gameOver() {
	
	fmt.Print("fin du jeu le mot était : ")
	fmt.Println(hangman.Mot)
	hangman.IsRunning = false
	
}