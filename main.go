package main

import (
	"fmt"
	"flag"
	hangman "hangman/GAME"
)



 func main() {
    var h hangman.HANGMAN
	help := flag.Bool("h",false, "Afficher de l'aide\n")
	flag.Parse()
	if *help {
		flag.Usage()
		return 
	}
	h.Init()
	fmt.Println("lancement du jeux")
	h.Start()
	h.Run()
	h.Stop()
	
 }