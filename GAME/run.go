package hangman

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "unicode"
)

// normalize convertit une chaîne en enlevant les accents et en la mettant en minuscules
func normalize(s string) string {
    t := ""
    for _, c := range s {
        switch c {
        case 'à', 'â', 'ä':
            t += "a"
        case 'é', 'è', 'ê', 'ë':
            t += "e"
        case 'î', 'ï':
            t += "i"
        case 'ô', 'ö':
            t += "o"
        case 'ù', 'û', 'ü':
            t += "u"
        case 'ç':
            t += "c"
        default:
            t += string(unicode.ToLower(c))
        }
    }
    return t
}

// Run est la boucle principale du jeu
func (hangman *HANGMAN) Run() {
    hangman.wordIsGood = false
    for hangman.IsRunning {
        hangman.hangman()    // Affiche l'état actuel du pendu
        hangman.readletter() // Lit une lettre ou un mot de l'utilisateur
        hangman.testword()   // Vérifie si le mot est complet
    }
}

// readletter lit l'entrée de l'utilisateur et la traite
func (hangman *HANGMAN) readletter() {
    fmt.Println(hangman.motIconnu)
    var Reader = bufio.NewReader(os.Stdin)
    String, _ := Reader.ReadString('\n')
    String = strings.TrimSpace(String)
    String = normalize(String) // Normalise l'entrée utilisateur (enlève les accents)

    if String >= "a" && String <= "z" {
        if len(String) > 1 {
            // Si l'utilisateur a entré un mot entier
            if String == normalize(hangman.Mot) {
                hangman.win() // Le joueur a gagné
            } else {
                hangman.erreur += 2 // Pénalité pour un mot incorrect
            }
        } else if len(String) == 1 {
            // Si l'utilisateur a entré une seule lettre
            hangman.lettre = String
            hangman.UsedLetter = append(hangman.UsedLetter, hangman.lettre)
            hangman.testLetter()
        }
    } else {
        fmt.Println("Lettre incorrecte !!")
    }
}

// testLetter vérifie si la lettre entrée est dans le mot et met à jour l'état du jeu
func (hangman *HANGMAN) testLetter() {
    hangman.lettreIsGood = false
    normalizedMotAdeviner := normalize(hangman.MotAdeviner)
    for l := range normalizedMotAdeviner {
        if hangman.lettre == string(normalizedMotAdeviner[l]) {
            hangman.lettreIsGood = true
            hangman.motIconnu[l] = hangman.MotAdeviner[l] // Utilise la lettre originale pour l'affichage
        }
    }
    if !hangman.lettreIsGood {
        hangman.erreur++ // Incrémente le compteur d'erreurs si la lettre n'est pas dans le mot
    }
    if hangman.erreur >= 9 {
        hangman.gameOver() // Fin du jeu si le nombre maximum d'erreurs est atteint
    }
}

// testword vérifie si le mot entier a été deviné
func (hangman *HANGMAN) testword() {
    hangman.wordIsGood = true
    for l := range hangman.motIconnu {
        if hangman.motIconnu[l] == '_' {
            hangman.wordIsGood = false
            break
        }
    }
    if hangman.wordIsGood {
        hangman.win() // Le joueur a gagné si toutes les lettres sont devinées
    }
    if hangman.erreur >= 9 {
        hangman.gameOver() // Fin du jeu si le nombre maximum d'erreurs est atteint
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