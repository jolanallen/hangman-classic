package hangman
import (
	"fmt"
	"os"
	//"bufio"
	"flag"
)



func (hangman *HANGMAN) Init() {
	
	if !hangman.flag {
		hangman.drawDisplay()
		hangman.testWord()
		hangman.Start()
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
	fmt.Println("Bienvenue dans le HANGMAN *Ynov 2024 B1 Info Montpellier* )                                                 		  ")
	fmt.Println("--------------------------------------------------------------------------------------------------------------------------------------")
	fmt.Println((`		
					__ __   ____  ____    ____  ___ ___   ____  ____  
   					|  |  | /    ||    \  /    ||   |   | /    ||    \ 
   					|  |  ||  o  ||  _  ||   __|| _   _ ||  o  ||  _  |
   					|  _  ||     ||  |  ||  |  ||  \_/  ||     ||  |  |
   					|  |  ||  _  ||  |  ||  |_ ||   |   ||  _  ||  |  |
   					|  |  ||  |  ||  |  ||     ||   |   ||  |  ||  |  |
   					|__|__||__|__||__|__||___,_||___|___||__|__||__|__|
	   `))
	fmt.Println("                                                   	  ")
	fmt.Println("					Bonne chance !						  ")
	fmt.Println("                                                   	  ")
	fmt.Println("                                                   	  ")
	fmt.Println("-----------------------------------------------------------------------------------------------------------------------------------")
}


func (hangman *HANGMAN) testWord() {
	TesTab, err := os.ReadFile("utile/wordlist/words.txt")
	if err != nil {
		fmt.Println(err)
		hangman.wordIsGood = false
		
	} 
	if err == nil {
		hangman.wordIsGood = true
		
		

	}
	if hangman.wordIsGood {
		hangman.TabByte = TesTab
		
	}
	
}

