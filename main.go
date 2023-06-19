package main

import (
	"log"
	"net/http"
	"text/template"
)

//----------------------------------------------------------Fonctions utilisées-------------------------------------------------------//
//                                                          ---------------------                                                     //

// checker si les caracteres sont printable sinon renvoyer une chaine vide
func Printable(s string) bool {
	txt := []rune(s)
	for i := 0; i <= len(txt)-1; i++ {
		if txt[i] != 13 && txt[i] != 10 && (txt[i] < 32 || txt[i] > 126) {
			return false
		}
	}
	return true
}
func landingHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/landing_page.html")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		template, _ := template.ParseFiles("./templates/error500.html")
		template.Execute(w, r)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Error while parsing file", http.StatusInternalServerError)
		w.WriteHeader(http.StatusNotFound)
		template, _ := template.ParseFiles("./templates/error500.html")
		template.Execute(w, r)
		return
	}
}
func ThanksHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/thank_you_page.html")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		template, _ := template.ParseFiles("./templates/error500.html")
		template.Execute(w, r)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Error while parsing file", http.StatusInternalServerError)
		w.WriteHeader(http.StatusNotFound)
		template, _ := template.ParseFiles("./templates/error500.html")
		template.Execute(w, r)
		return
	}
}

// Gestion de la page d'acceuil
func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Vérifier que le lien saisi est valide
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		template, _ := template.ParseFiles("./templates/error405.html")
		template.Execute(w, r)
		return
	}

	if r.URL.Path != "/" && r.URL.Path != "/ascii-art" {
		w.WriteHeader(http.StatusNotFound)
		template, _ := template.ParseFiles("./templates/error404.html")
		template.Execute(w, r)
		return
	}
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		template, _ := template.ParseFiles("./templates/error500.html")
		template.Execute(w, r)
		return
	}
	//Traitement erreur 500
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Error while parsing file", http.StatusInternalServerError)
		w.WriteHeader(http.StatusNotFound)
		template, _ := template.ParseFiles("./templates/error500.html")
		template.Execute(w, r)
		return
	}
}

// Handler pour le traitement du formulaire ASCII Art
func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" { // Vérification de la route de la requête HTTP
		w.WriteHeader(http.StatusBadRequest)
		template, _ := template.ParseFiles("./templates/error400.html")
		template.Execute(w, r)
		return
	}
	/* if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		template, _ := template.ParseFiles("./templates/error404.html")
		template.Execute(w, r)
		return
	} */
	text := r.FormValue("text")
	banner := r.FormValue("banner")
	_, err := template.ParseFiles(banner + ".txt")
	if err != nil {
		if os.IsNotExist(err) {
			w.WriteHeader(http.StatusInternalServerError)
			template, _ := template.ParseFiles("./templates/error500.html")
			template.Execute(w, r)
			return
		}

	}
	asciiArt := Naboufs(banner+".txt", text)
	if !Printable(text) {
		w.WriteHeader(http.StatusNotFound)
		template, _ := template.ParseFiles("./templates/error400.html")
		template.Execute(w, r)
		return

	}
	//Traitement erreur 404
	if Printable(text) && text == "" || banner == "" || (banner != "standard" && banner != "thinkertoy" && banner != "shadow") && r.Method == "GET" {
		w.WriteHeader(http.StatusNotFound)
		template, _ := template.ParseFiles("./templates/error404.html")
		template.Execute(w, r)
		return

	}
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, struct {
		Text     string
		Banner   string
		AsciiArt string
	}{
		Text:     text,
		Banner:   banner,
		AsciiArt: asciiArt,
	})
}

func main() {
	// Configurer le serveur de fichiers statiques pour les fichiers CSS et autres ressources
	fs := http.FileServer(http.Dir("templates"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Définir les gestionnaires de routes pour la page d'accueil et le traitement du formulaire ASCII Art
	http.HandleFunc("/landing_page", landingHandler)
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)
	http.HandleFunc("/thank_you_page", ThanksHandler)

	// Démarrer le serveur HTTP
	log.Println("Server started on http://localhost:8080/landing_page")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
