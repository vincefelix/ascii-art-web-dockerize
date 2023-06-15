package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// printres gère l'affichage d'un tableau à 2 dimensions contenant des caractères
func printres(result [][]string) string {
	var temp string          // stocker les caractères à afficher par ligne
	for i := 0; i < 8; i++ { // parcourir la colonne
		for j := 0; j < len(result); j++ { // parcourir la ligne
			temp += result[j][i]
		}
		if i != 7 { // ne pas ajouter de newline au dernier caractère
			if temp == "" {
				break
			} else {
				temp += "\n"
			}
		}
	}
	return temp // retourne la ligne
}

// IsPrintable permet de vérifier si le string comporte un caractère affichable ou pas
// - il renvoie true s'il rencontre un seul caractère affichable
// - et renvoie false s'il n'en voit aucun.
func IsPrintable(s string) bool {
	a := []rune(s) // convertir le tableau string en tableau de runes
	b := len(s)
	rep := true
	for i := 0; i <= b-1; i++ {
		if a[i] < 0 || a[i] > 127 {
			rep = false
			break
		} else if a[i] < 32 || a[i] > 126 { // intervalle des caractères affichables
			rep = false
		} else {
			rep = true
			break
		}
	}
	return rep
}

// Newline enlève le dernier élément d'un tableau exclusivement constitué de newline
func Newline(tab []string) []string {
	var count int // compteur
	for _, v := range tab {
		if v == "" {
			count++
		}
	}
	var res []string

	// il n'y a que des newlines
	if count == len(tab) {
		res = tab[:len(tab)-1]
	} else {
		// on renvoie le tableau d'origine
		res = tab
	}
	return res
}

// Naboufs génère l'art ASCII en fonction du fichier et de la phrase donnés
func Naboufs(filename, phrase string) string {
	// ------------------- 1ère étape : Lire le fichier avec les graphiques ------------------------
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Erreur: nous ne parvenons pas à lire le fichier source : ", err)
		return ""
	}
	defer file.Close()

	// Stocker les caractères des graphiques dans une variable
	longtext := bufio.NewScanner(file)

	// Mettre les données stockées dans un tableau de string
	var tab []string
	for longtext.Scan() {
		tab = append(tab, longtext.Text())
	}

	// ------------------- 2ème étape : Stocker chaque ensemble de caractères pour chaque ASCII -----
	var vinc [][]string
	for i := 1; i < len(tab); i += 9 {
		vinc = append(vinc, tab[i:i+8])
	}

	// ------------------- 3ème étape : Gérer l'affichage --------------------------p----------------
	var asciiArt string

	// Vérifier si l'argument contient un caractère affichable
	test := phrase
	if !IsPrintable(test) {
		return ""
	}

	// L'argument contient un caractère affichable
	splitext := strings.Split(test, "\r\n") // séparer le string en cas de présence d'un "newline"
	splitext = Newline(splitext)
	var num int // variable pour déterminer l'index dans le tableau des caractères

	// Afficher le string de l'argument sous le format ASCII art
	for _, v := range splitext {
		// Récolter les caractères ASCII art à afficher
		var result [][]string
		for _, y := range v {
			num = int(y - 32) // la position correspondant au caractère selon le tableau de caractères dans vinc
			if num > 95 {
				continue
			} else {
				result = append(result, vinc[num])
			}
		}
		asciiArt += printres(result) + "\n" // ajouter la version graphique des caractères récoltées à asciiArt
	}

	return asciiArt

}
