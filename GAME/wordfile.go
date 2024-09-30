package hangman

import (
	"bufio"
	"fmt"
	//"math/rand"
	"os"
	//"time"
)

func (hangman *HANGMAN) wordlist() {
}
func (hangman *HANGMAN) lireMotsDepuisFichier(nomFichier string) {
	file, err := os.Open(nomFichier)
    if err != nil {
		fmt.Println(err)
	}
    scanner := bufio.NewScanner(file)
	
    for scanner.Scan(){
        hangman.TabMots = append(hangman.TabMots, scanner.Text())
    }
	fmt.Println(hangman.TabMots)
    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }
//	hangman.choisirmotaleatoire(hangman.TabMots)
	defer file.Close()
}

// func (hangman *HANGMAN) choisirmotaleatoire(mot []string)  {
//     rand.Seed(time.Now().UnixNano())
//     index := rand.Intn(len(mots))
    
// }
