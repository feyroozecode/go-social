// Définir le package principal
package main

// Importer les packages nécessaires
import (
	"log"
	"net/http"
)

// Définir une structure pour le serveur
type api struct {
	addr string // Adresse sur laquelle le serveur écoutera
}

// ServeHTTP implémente l'interface http.Handler.
func (s *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Gérer les requêtes en fonction de la méthode HTTP
	switch r.Method {
	case http.MethodGet:
		// Gérer les requêtes GET en fonction du chemin URL
		switch r.URL.Path {
		case "/":
			// Répondre avec "Index" pour la route racine
			w.Write([]byte("Index "))
		case "/users":
			// Répondre avec "users page" pour la route /users
			w.Write([]byte("users page"))
		}
	default:
		// Répondre avec "404 page" pour toutes les autres requêtes
		w.Write([]byte("404 page"))
	}
}

// Fonction principale pour démarrer le serveur
func main() {
	// Créer une instance de la structure api avec l'adresse ":8080"
	s := &api{addr: ":8080"}
	// Démarrer le serveur HTTP et écouter sur l'adresse spécifiée
	if err := http.ListenAndServe(s.addr, s); err != nil {
		// En cas d'erreur, enregistrer l'erreur et arrêter le programme
		log.Fatal((err))
	}
}
