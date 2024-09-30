package hangman

import (
    "bufio"
    "fmt"
    "log"
    "math/rand"
    "os"
    "time"
)

func (hangman *HANGMAN) wordlist() {
    mots, err := lireMotsDepuisFichier("words.txt")
    if err != nil {
        log.Fatalf("Erreur lors de: %v", err)
    }
    if len(mots) == 0 {
        log.Fatal("fichier vide.")
    }
    motAleatoire := choisirmotaleatoire(mots)
    fmt.Println("Le mot est:", motAleatoire)
}

func lireMotsDepuisFichier(nomFichier string) ([]string, error) {
    file, err := os.Open(nomFichier)
    if err != nil {
return nil, err
    }
    defer file.Close()
    var mots []string
    scanner := bufio.NewScanner(file)

    for scanner.Scan(){
        mots = append(mots, scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        return nil, err
    }
    return mots, nil
}

func choisirmotaleatoire(mots []string) string {
    rand.Seed(time.Now().UnixNano())
    index := rand.Intn(len(mots))
    return mots[index]

}
