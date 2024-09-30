package hangman




func (hangman *HANGMAN) Run() {
	vie := 10
	blanc := []strings{}
	for range mot {
		blanc = append(blanc,"_")
	}
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
	if !BonneLettre {
		vie -= 1
	}
	}
	if vie <= 0 {
		hangman.gameOver()
	}

	//for hangman.IsRunning {
		
		
	//}

	
}

func (hangman *HANGMAN) win() {

}

func (hangman *HANGMAN) gameOver() {

}