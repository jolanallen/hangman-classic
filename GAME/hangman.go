package hangman

import (
	"fmt"
	"os"
)



func (hangman *HANGMAN) hangman() {
	hangman.etat1()
}

func (hangman *HANGMAN) etat1() {
	Etat1, err := os.Open("utile/hangman/etat1.txt")
	if err != nil {
		fmt.Println(err)
	}
	if err == nil {
		fmt.Println(Etat1)
	}
}