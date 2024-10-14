package hangman

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func (hangman *HANGMAN) Run() {
	hangman.wordIsGood = false
	for hangman.IsRunning {    /// Boucle principale qui continue tant que le jeu est actif.
		hangman.hangman()      
		hangman.readletter()  
		hangman.testword()     
	}			
}

func (hangman *HANGMAN) readletter() { 
	fmt.Println(hangman.motIconnu)    
	var Reader = bufio.NewReader(os.Stdin) ///// Crée un lecteur pour récupérer l'entrée de l'utilisateur dans le terminal
	String,_ := Reader.ReadString('\n')    /// Récupère la saisie de l'utilisateur sous forme de chaîne.
	String = strings.TrimSpace(String)     // supprime les espaces ou retours à la ligne 

	if String >= "a" && String <= "z" {   			 /// Vérifie que l'entrée est bien une lettre (de a à z)
		if len(String) > 1 {               /// Si l'utilisateur a entré un mot entier
			if String == hangman.Mot {    	 // Si le mot deviné est correct, alors le joueur  gagne
				hangman.win()              
			} else {
				hangman.erreur += 2        /// Sinon le mot est incorrect, + 2 a erreur
			}
		} else if len(String) == 1 {        /// Si l'utilisateur a entré une seule lettre.
			hangman.lettre = String        
			hangman.UsedLetter = append(hangman.UsedLetter, hangman.lettre) /// Ajoute cette lettre au tableau de lettres déjà utilisées.
			hangman.testLetter()           
		} 
	} else {
		fmt.Println("lettre incorrecte !!") 
	}
}

func (hangman *HANGMAN) testLetter() {
	hangman.lettreIsGood = false
	for l := range hangman.MotAdeviner {
		if hangman.lettre == hangman.MotAdeviner[l] { 			/// Si la lettre correspond à une lettre du mot à deviner.
			hangman.lettreIsGood = true              	
			hangman.motIconnu[l] = hangman.lettre    /// Remplace le tiret par la lettre trouvée dans le mot affiché.
			       
		}
	}
	if !hangman.lettreIsGood {                        /// Si la lettre n'est pas dans le mot
		hangman.erreur++                  /// Ajoute une erreur
	}
	if hangman.erreur >= 9 {                          /// Si 9 erreurs sont atteintes, la partie est terminée.
		hangman.gameOver()                           
	}
}

func (hangman *HANGMAN) testword() {
	hangman.wordIsGood = true
	for l := range hangman.motIconnu {
		if !(hangman.motIconnu[l] == hangman.MotAdeviner[l]) { /// Vérifie si toutes les lettres du mot deviné sont correctes.
			hangman.wordIsGood = false             /// Si une lettre ne correspond pas, le mot n'est pas complet.
			break                          // arrêt de la boucle 
		}
	}
	if hangman.wordIsGood {    /// Si le mot est entièrement correct le joueur gagne 
		hangman.win()          
	}
	if hangman.erreur >= 9 {   /// Si il ya 9 erreurs  la partie est perdue.
		hangman.gameOver()     
	}
}


func (hangman *HANGMAN) win() {
		fmt.Println(`
 		__     __          __          ___       _ 
 		\ \   / /          \ \        / (_)     | |
	 	 \ \_/ /__  _   _   \ \  /\  / / _ _ __ | |
 	 	  \   / _ \| | | |   \ \/  \/ / | | '_ \| |
  		   | | (_) | |_| |    \  /\  /  | | | | |_|
  	 	   |_|\___/ \__,_|     \/  \/   |_|_| |_(_)
                                           
                                           

		`)
		fmt.Println("fin du jeu vous avez gagner le mot était bien", hangman.Mot)
	hangman.IsRunning = false
	
}

func (hangman *HANGMAN) gameOver() {
	fmt.Println(` 
		   ____                          ___                 
		  / ___| __ _ _ __ ___   ___    / _ \__   _____ _ __ 
		 | |  _ / _' | '_ ' _ \ / _ \  | | | \ \ / / _ \ '__|
		 | |_| | (_| | | | | | |  __/  | |_| |\ V /  __/ |
	    	  \____|\__,_|_| |_| |_|\___|   \___/  \_/ \___|_|
                                                                                              

	`)

	fmt.Println("fin du jeu le mot était : ", hangman.Mot)
	hangman.IsRunning = false
	
}