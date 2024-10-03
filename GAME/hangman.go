package hangman

import (
	"fmt"

)



func (hangman *HANGMAN) hangman() {
	if hangman.erreur == 0 {
		hangman.Etat0()
	}
	if hangman.erreur == 1 {
		hangman.Etat1()
        
	}
	if hangman.erreur == 2 {
		hangman.Etat2()
	}
	if hangman.erreur == 3 {
		hangman.Etat3()
	}
	if hangman.erreur == 4 {
		hangman.Etat4()
	}
	if hangman.erreur == 5 {
		hangman.Etat5()
	}
	if hangman.erreur == 6 {
		hangman.Etat6()
	}
}


func (hangman *HANGMAN) Etat0() {
    fmt.Print(`
     +---+
     |   |
         |
         |
         |
         |
    =========
    `)
}


func (hangman *HANGMAN) Etat1() {
    fmt.Print(`
     +---+
     |   |
     O   |
         |
         |
         |
    =========
    `)
}


func (hangman *HANGMAN) Etat2() {
    fmt.Print(`
     +---+
     |   |
     O   |
     |   |
         |
         |
    =========
    `)
}


func (hangman *HANGMAN) Etat3() {
    fmt.Println(`
     +---+
     |   |
     O   |
    /|   |
         |
         |
    =========
    `)
}


func (hangman *HANGMAN) Etat4() {
    fmt.Println(`
     +---+
     |   |
     O   |
    /|\  |
         |
         |
    =========
    `)
}


func (hangman *HANGMAN) Etat5() {
    fmt.Println(`
     +---+
     |   |
     O   |
    /|\  |
    /    |
         |
    =========
    `)
}


func (hangman *HANGMAN) Etat6() {
    fmt.Println(`
     +---+
     |   |
     O   |
    /|\	 |
    / \  |
         |
    =========
    `)
	hangman.gameOver()
}