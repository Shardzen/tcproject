package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// 1. Écoute sur le port 8080 (TCP)
	fmt.Println("Lancement du serveur sur le port 9000...")
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println("Erreur lors du lancement du serveur:", err)
		os.Exit(1)
	}
	defer listener.Close()

	for {
		// 2. Attend une connexion entrante
		fmt.Println("En attente de connexion...")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Erreur lors de la connexion:", err)
			continue
		}

		fmt.Println("Un client s'est connecté !")

		// 3. Gère la connexion 
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Scanner pour lire ce qui arrive via la connexion
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()
		fmt.Printf("Message reçu : %s\n", message)
	}

	fmt.Println("Client déconnecté.")
}