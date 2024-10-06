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

	hangman.hangman()
	hangman.testword()
	
}

	
// pour appeler la fonction test word on vas utiliser la loncguer de mot inconu 
// initialement le tableau mot inconnu suera vide et si la longuer du tableau inconu  = longeur mot a devier alors on appelle test word 
// l'information de la taille du mot sera donnée par un print len(mot a deviner)

func (hangman *HANGMAN) randomWord() {
	rand.Seed(time.Now().UnixMilli())
	hangman.randomNb = rand.Intn(len(hangman.TabMots))
	hangman.Mot = hangman.TabMots[hangman.randomNb]
	for _, char := range hangman.Mot {
		hangman.MotAdeviner = append(hangman.MotAdeviner, string(char))
	}
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