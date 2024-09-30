package hangman

import "fmt"





func (hangman *HANGMAN) Start() {
	if hangman.wordIsGood {
		hangman.wordlist()
	} 
}

func(hangman *HANGMAN) Stop() {
	hangman.IsRunning = false
	fmt.Print("fin du jeu le mot était ")
	fmt.Println(hangman.MotAdeviner)
	// si la touche ctrl et q est pressé
	//on sauvegarde
	//on passe hangman.IsRunning = false ce qui arrête run()
	// save 
}

