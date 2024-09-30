package hangman





func (hangman *HANGMAN) Start() {
	hangman.IsRunning = false
	hangman.wordlist()
	hangman.hangman()
}

func(hangman *HANGMAN) Stop() {
	hangman.IsRunning = false
	// save 
}