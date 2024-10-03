package hangman

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)



func (hangman *HANGMAN) Init() {
	
	if !hangman.flag {
		hangman.drawDisplay()
		hangman.intiwordlist()
		hangman.initHangman()
	} else {
		hangman.drawDisplay()
		hangman.Flag()
		hangman.intiwordlist()
		hangman.initHangman()

	}
}
func (hangman *HANGMAN)Flag() {
	help := flag.Bool("h",false, "Afficher de l'aide\n")
	flag.Parse()
	if *help {
		flag.Usage()
		hangman.flag = true
	}
}
func (hangman *HANGMAN) drawDisplay() {
	fmt.Println("Bienvenue dans le HANGMAN Ynov 2024 B1 Info Montpellier )")
	fmt.Println("---------------------------------------------------------------------------------------------------------------------------------------------")
	fmt.Println((`																													 
					 __ __   ____  ____    ____  ___ ___   ____  ____  				\ / |\ | / \ \  /
   					|  |  | /    ||    \  /    ||   |   | /    ||    \ 				 |  | \| \_/  \/
   					|  |  ||  o  ||  _  ||   __|| _   _ ||  o  ||  _  |
   					|  _  ||     ||  |  ||  |  ||  \_/  ||     ||  |  |
   					|  |  ||  _  ||  |  ||  |_ ||   |   ||  _  ||  |  |
   					|  |  ||  |  ||  |  ||     ||   |   ||  |  ||  |  |
   					|__|__||__|__||__|__||___,_||___|___||__|__||__|__|
	   `))
	fmt.Println("                                                   	  ")
	fmt.Println("					                     Bienvenue !						  ")
	fmt.Println("                                                   	  ")
	fmt.Println("---------------------------------------------------------------------------------------------------------------------------------------------")
}

func (hangman *HANGMAN) intiwordlist() {
	wordFile, err := os.Open("utile/wordlist/words.txt")
	if err != nil {
		fmt.Println(err)
		defer wordFile.Close()
	}
	if err == nil {
		scanner := bufio.NewScanner(wordFile)
		var mot string
		for scanner.Scan() { 
			mot = scanner.Text()
			hangman.TabMots = append(hangman.TabMots,  mot)
		}
		defer wordFile.Close()
		hangman.IsRunning = true
	}
}
func (hangman *HANGMAN) initHangman() {
	
	hangman.erreur = 0
	hangman.hangman()
  
}