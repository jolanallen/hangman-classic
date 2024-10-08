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
	randomlettre  string
	erreur       int
	lettreIsGood bool
	wordFile     string
	lettre       string
	UsedLetter   []string
	soft         bool
	medium       bool
	hard         bool
	
	
	
}

