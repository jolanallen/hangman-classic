package hangman




func (hangman *HANGMAN) Run() {
	blanc := []strings{}
	for range mot {
		blanc = append(blanc,"_")
	}
	for {
		var input string
		fmt.Scanln(&input)
		for _, inputLetter := range input {
			BonneLettre := false
			for i, motLetter := range mot {
				if inputLetter == motLetter {
					blanc[i] = string(inputLetter)
					BonneLettre = true
				}
			}
	

	}
	}

	//for hangman.IsRunning {
		
		
	//}

	
}

func (hangman *HANGMAN) win() {

}

func (hangman *HANGMAN) gameOver() {

}