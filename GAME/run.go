package hangman




func (hangman *HANGMAN) Run() {
	vie := 10 //nombre de vie(a changer comme on veut)
	blanc := []strings{} // pour l'affichage des blancs
	for range mot {
		blanc = append(blanc,"_")
	}
	var input string // pour entrer une lettre 
	fmt.Scanln(&input) // lire la lettre entrer
	for _, inputLetter := range input {
		BonneLettre := false
		for i, motLetter := range mot { // verifier si la lettre entrer est dans le mot
			if inputLetter == motLetter {
				blanc[i] = string(inputLetter) // remplacer le blanc par la lettre
				BonneLettre = true
			}
		}
	if !BonneLettre { // si la lettre n'est pas dans le mot, on perd une vie
		vie -= 1
	}
	}
	if vie <= 0 { // si plus de vie on game over
		hangman.gameOver()
	}

	//for hangman.IsRunning {
		
		
	//}

	
}

func (hangman *HANGMAN) win() {

}

func (hangman *HANGMAN) gameOver() {

}