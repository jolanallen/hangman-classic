package hangman

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"math/rand"
	"time"
)



func (hangman *HANGMAN) Init() {
		hangman.drawDisplay()
		hangman.intiwordlist()
		hangman.initHangman()
		hangman.randomWord()
}
func (hangman *HANGMAN)Flag() {
	help := flag.Bool("h", false, "Afficher de l'aide")
	hard := flag.Bool("hard", false, "mode de jeu hard mots dificile")
	medium := flag.Bool("medium", false, "mode de jeu medium mot moyennement dificile")
	
	flag.Parse()

	if *help {
		flag.Usage()
	}
	if *hard {
		hangman.hard = true 
		fmt.Println("mode de jeu hard")
	} else if *medium {
		hangman.medium = true
		fmt.Println("mode de jeu medium")
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
	
	if hangman.medium {
		hangman.wordFile = "utile/wordlist/words2.txt"
		fmt.Println("mode de jeu : intermédaire")
	} else if hangman.hard {
		hangman.wordFile = "utile/wordlist/words3.txt"
		fmt.Println("mode de jeu : hard")
	} else {
		hangman.wordFile = "utile/wordlist/words.txt"
		fmt.Println("mode de jeu : soft (defaut)")
		fmt.Println("tapez -medium pour un mode de jeux intermediare")
		fmt.Println("tapez -hard pour un mode de jeux dificile")
	}
		
	wordFile, err := os.Open(hangman.wordFile)
	if err != nil {
		fmt.Println("Erreur ouverture fichier:", err)
		defer wordFile.Close()
	}
	if err == nil {
		scanner := bufio.NewScanner(wordFile)
		for scanner.Scan() { 
			hangman.TabMots = append(hangman.TabMots,  scanner.Text())
		}
		defer wordFile.Close()
	}
}
func (hangman *HANGMAN) randomWord() {
	rand.Seed(time.Now().UnixMilli())                 /// unixmilli c'est la date est l'heure de référence des systéme linux et time.now c'est l'heure actuelle seed c'est le point de départ a partir du quel le nombre aléatoire sera générer 
	hangman.randomNb = rand.Intn(len(hangman.TabMots))
	hangman.Mot = hangman.TabMots[hangman.randomNb]
	for _, char := range hangman.Mot {
		hangman.MotAdeviner = append(hangman.MotAdeviner, string(char))
	}
	//remplace les lettre par des underScore
	for i :=  0; i < len(hangman.Mot); i++ {
		hangman.motIconnu = append(hangman.motIconnu, "_")
	}
	//Révéler une lettre au hasard
	hangman.randomNb = rand.Intn(len(hangman.Mot))
	hangman.motIconnu[hangman.randomNb] = hangman.MotAdeviner[hangman.randomNb] // permet de faire apparaitre  une lettre aléatoirement dans le mot a deviner

	hangman.IsRunning = true
	

}
func (hangman *HANGMAN) initHangman() {
	hangman.erreur = 0 

	hangFile, err := os.Open("utile/hangman/hangman.txt")
	if err != nil {
		fmt.Println(err)

		defer hangFile.Close()
	}
  	if err == nil {
		scanner := bufio.NewScanner(hangFile)
    	
		for scanner.Scan() { 
     		if scanner.Text() == "" {
       			hangman.TabHang = append(hangman.TabHang, hangman.etat.String() )
       			hangman.etat.Reset()
     		} else {
        		hangman.etat.WriteString(scanner.Text() + "\n") 
     		}
		}
		defer hangFile.Close()
		
	}
}