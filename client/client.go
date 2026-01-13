package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// 1. Se connecte au serveur local
conn, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println("Impossible de se connecter (le serveur est-il lancé ?) :", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connecté ! Tape ton message et appuie sur Entrée.")

	// 2. Lit ce que tu tapes au clavier et l'envoie
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		
		// Envoie le texte + un retour à la ligne pour que le serveur sache que le message est fini
		fmt.Fprintf(conn, "%s\n", text)
	}
}