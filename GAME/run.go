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
	var l string = "_"
	for i :=  0; i < len(hangman.Mot); i++ {
		hangman.motIconnu = append(hangman.motIconnu, l)
	}
	fmt.Println(hangman.motIconnu)

}





func (hangman *HANGMAN) win() {
		fmt.Println(`
 		__     __          __          ___       _ 
 		\ \   / /          \ \        / (_)     | |
	 	 \ \_/ /__  _   _   \ \  /\  / / _ _ __ | |
 	 	  \   / _ \| | | |   \ \/  \/ / | | '_ \| |
  		   | | (_) | |_| |    \  /\  /  | | | | |_|
  	 	   |_|\___/ \__,_|     \/  \/   |_|_| |_(_)
                                           
                                           

		`)
		fmt.Println("fin du jeu vous avez gagner le mot était bien", hangman.Mot)
	hangman.IsRunning = false
	
	
	
}

func (hangman *HANGMAN) gameOver() {
	fmt.Println(` 
		   ____                          ___                 
		  / ___| __ _ _ __ ___   ___    / _ \__   _____ _ __ 
		 | |  _ / _' | '_ ' _ \ / _ \  | | | \ \ / / _ \ '__|
		 | |_| | (_| | | | | | |  __/  | |_| |\ V /  __/ |
	    	  \____|\__,_|_| |_| |_|\___|   \___/  \_/ \___|_|
                                                                                              

	`)

	fmt.Print("fin du jeu le mot était : ")
	fmt.Println(hangman.Mot)
	hangman.IsRunning = false
	
}