package hangman

import (
	"bufio"
	"fmt"
	//"math/rand"
	"os"
	//"time"
)

func (hangman *HANGMAN) wordlist() {
	//hangman.lireMotsDepuisFichier("hangman/utile/wordlist/words.txt")
}
// func (hangman *HANGMAN) lireMotsDepuisFichier(nomFichier string) {
// 	file, err := os.Open(nomFichier)
//     if err != nil {
// 		fmt.Println(err)
// 	}
//     scanner := bufio.NewScanner(file)
	
//     for scanner.Scan(){
//         hangman.TabMots = append(hangman.TabMots, scanner.Text())
//     }
// 	fmt.Println(hangman.TabMots)
//     if err := scanner.Err(); err != nil {
//         fmt.Println(err)
//     }
	defer file.Close()
	// motAleatoire := hangman.choisirmotaleatoire()
    // fmt.Println("Le mot est:", motAleatoire)
	// hangman.wordIsGood = true
	// si les mot sont invalide
		//revoit invalide wordlist
}


func (hangman *HANGMAN) choisirmotaleatoire(mot []string)  {
   //	rand.Seed(time.Now().UnixNano())
  //  index := rand.Intn(len(mots))
    
 }
