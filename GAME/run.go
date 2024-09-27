package hangman




func (hangman *HANGMAN) Run() {
	
	for hangman.IsRunning {
		hangman.win()
		hangman.gameOver()
		
	}

	
}

func (hangman *HANGMAN) win() {

}

func (hangman *HANGMAN) gameOver() {

}