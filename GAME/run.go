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
		if hangman.wordIsGood {
			hangman.win()
		} else {
			hangman.gameOver()
		}
		
		
		

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
	hangman.wordIsGood = true
	for l := range hangman.TabWord {
		if hangman.motIconnu[l] == "_" {
			break
		} else if !(hangman.motIconnu[l] ==  hangman.TabMots[l]) {
			hangman.wordIsGood = false
			break
		}
			
		
	}
	
}
func (hangman *HANGMAN) testLetter() {
	hangman.lettreIsGood = false
	for l := range hangman.MotAdeviner {
		if hangman.lettre == hangman.MotAdeviner[l] {
			hangman.lettreIsGood = true
			hangman.motIconnu[l] = string(hangman.lettre)
		}
		
	}
	if hangman.lettreIsGood {
		fmt.Println(hangman.motIconnu)
	} else {
		
		hangman.erreur++
	}
	hangman.hangman()
	
}

	
// pour appeler la fonction test word on vas utiliser la loncguer de mot inconu 
// initialement le tableau mot inconnu suera vide et si la longuer du tableau inconu  = longeur mot a devier alors on appelle test word 
// l'information de la taille du mot sera donnée par un print len(mot a deviner)

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
	
	fmt.Print("fin du jeu vous avez gagner")
	hangman.IsRunning = false
	
	
	
}

func (hangman *HANGMAN) gameOver() {
	
	fmt.Print("fin du jeu le mot était : ")
	fmt.Println(hangman.Mot)
	hangman.IsRunning = false
	
}