package main

import (
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
	h.DrawDisplay()
	h.Start()
	h.Run()
	h.Stop()
	
 }