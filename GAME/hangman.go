package hangman

import (
	"fmt"

)



func (hangman *HANGMAN) hangman() {
	
}


func (hangman *HANGMAN) Etat0() {
    fmt.Println(`
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
    fmt.Println(`
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
    fmt.Println(`
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
    /|\\  |
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
    /|\\  |
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
    /|\\  |
    / \\  |
         |
    =========
    `)
}