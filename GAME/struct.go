package hangman

import "strings"



type HANGMAN struct {
	IsRunning     bool
	wordIsGood    bool
	TabMots       []string
	TabRune       []rune
	Mot           string
	MotAdeviner   []string
	TabHang       []string
	etat          strings.Builder
	motIconnu     []string
	randomNb      int
	flag           bool
	erreur       int
	lettreIsGood bool
	lettre       string
	tabletter    []string
	UsedLetter   []string
	
	
	
}

