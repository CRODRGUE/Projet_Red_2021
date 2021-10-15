package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Menu principale du jeux
func (p *perso) Menu() {
	var g1 monstre
	g1.initGoblin("Barbare des mers")
	fmt.Println("------------------Menu------------------")
	PrintMenu([]string{" Afficher les informations du personnage", " Accéder aux contenus de l’inventaire", " Marchand", " Forgeron", " Camps d'entrainement", Red + " Quitter" + Reset, Yellow + " Qui sont-ils ?" + Reset})
	fmt.Println("-----------------------------------------")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Indique moi ce que vous voulez faire : ")
	scanner.Scan()
	choice, _ := strconv.Atoi(scanner.Text())

	switch choice {
	case 1:
		p.displayinfo()
	case 2:
		Clear()
		p.MenuInventory()
	case 3:
		Clear()
		p.Marchand()
	case 4:
		Clear()
		p.forgeron()
	case 5:
		Clear()
		p.trainingFight(&g1)
	case 6:
		break
	case 7:
		Clear()
		QuiSontIls()
		p.Menu()
	default:
		Clear()
		fmt.Println(Yellow, "Erreur donnée un entier compris entre 1 et 6 ", Reset)
		p.Menu()
	}
}
