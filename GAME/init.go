package hangman
import (
	"fmt"
	"log"
)



func (hangman *HANGMAN) Init() {
	hangman.testWord()
	hangman.start()
}

func (hangman *HANGMAN) testWord() {
	mots, err := lireMotsDepuisFichier("hangman/utile/wordlist/words.txt")
    if err != nil {
        log.Fatalf("Erreur lors de: %v", err)
    }
    if len(mots) == 0 {
        log.Fatal("fichier vide.")
    }
    motAleatoire := hangman.choisirmotaleatoire()
    fmt.Println("Le mot est:", motAleatoire)
	hangman.wordIsGood = true
	// si les mot sont invalide
		//revoit invalide wordlist
}