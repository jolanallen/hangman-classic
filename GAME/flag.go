package hangman


import "flag"


func (hangman *HANGMAN)Flag() {
	help := flag.Bool("h",false, "Afficher de l'aide\n")
	flag.Parse()
	if *help {
		flag.Usage()
		return 
	}
}