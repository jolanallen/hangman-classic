package hangman


import (
	"fmt"
)

func (hangman *HANGMAN) hangman() {
	if hangman.erreur == 0 {
		hangman.Etat0()
	}
	if hangman.erreur == 1 {
		hangman.Etat1()
	}
	if hangman.erreur == 2 {
		hangman.Etat2()
	}
	if hangman.erreur == 3 {
		hangman.Etat3()
	}
	if hangman.erreur == 4 {
		hangman.Etat4()
	}
	if hangman.erreur == 5 {
		hangman.Etat5()
	}
	if hangman.erreur == 6 {
		hangman.Etat6()
	}
	if hangman.erreur == 7 {
		hangman.Etat7()
	}
	if hangman.erreur == 8 {
		hangman.Etat8()
	}
	if hangman.erreur == 9 {
		hangman.Etat9()
        hangman.gameOver()
	}

}

func (hangman *HANGMAN) Etat0() {
	fmt.Println(`
         
      |  
      |  
      |  
      |  
      |  
      |  
      |
  ==========
    `)
}

func (hangman *HANGMAN) Etat1() {
	fmt.Println(`
          
  +---+  
      |  
      |  
      |  
      |  
      |
      | 
  =========
    `)
}

func (hangman *HANGMAN) Etat2() {
	fmt.Println(` 
         
  +---+  
  |   |  
      |  
      |  
      |  
      |  
      |
  =========
    `)
}

func (hangman *HANGMAN) Etat3() {
	fmt.Println(`
         
  +---+  
  |   |  
  O   |  
      |  
      |  
      |  
      |
  =========
    `)
}

func (hangman *HANGMAN) Etat4() {
	fmt.Println(`
    
  +---+  
  |   |  
  O   |  
  |   |  
      |  
      | 
      | 
  =========
    `)
}

func (hangman *HANGMAN) Etat5() {
	fmt.Println(`
          
  +---+  
  |   |  
  O   |  
 /|   |  
      |  
      |  
      |
  =========
    `)
}

func (hangman *HANGMAN) Etat6() {
	fmt.Println(`
          
  +---+  
  |   |  
  O   |  
 /|\  |  
      |  
      | 
      | 
  =========
    `)
}

func (hangman *HANGMAN) Etat7() {
	fmt.Println(`
     
  +---+  
  |   |  
  O   |  
 /|\  |  
 /    | 
      |
  =========
    `)
}

func (hangman *HANGMAN) Etat8() {
	fmt.Println(`
         
  +---+  
  |   |  
  O   |  
 /|\  |  
 / \  |  
      |
  =========
    `)
}


func (hangman *HANGMAN) Etat9() {
	fmt.Println(`
          
  +---+  
  |   |  
  O   |  
 /|\  |  
 / \  |  
      |
  =========
    `)
}
