package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Affiche les différents inventaires et l'équipement du personnage
func (p *perso) displayinventory() {
	if len(p.skill) == 0 {
		fmt.Println("Votre inventaire de skill est vide")
	} else {
		fmt.Println("--------Votre inventaire de skill--------")
		for i := 0; i <= len(p.skill)-1; i++ {
			fmt.Println("- ", p.skill[i])
		}
		fmt.Println("")
	}
	if len(p.inventory) == 0 {
		fmt.Println("Votre inventaire est vide")
	} else {
		fmt.Println("------------Votre inventaire------------")
		for i := 0; i <= len(p.inventory)-1; i++ {
			fmt.Println("- ", p.inventory[i])
		}
	}
	fmt.Println("---------------Equipement---------------")
	fmt.Println("- ", p.head)
	fmt.Println("- ", p.torso)
	fmt.Println("- ", p.leg)
}

// Menu de l'inventaire (permet l'utilisation d'item et l'équipement de pièce d'armure grâce à l'appel de fonctions)
func (p *perso) MenuInventory() {
	p.displayinventory()
	fmt.Println("-----------------------------------------")
	PrintMenu([]string{" Utiliser bouteille de rhum", " Utiliser bouteille d'eau", " Equiper chapeau de pirate (+ 10 points de vie max)", " Equiper veste de pirate (+ 25 points de vie max)", " Equiper jambe de bois (+ 15 points de vie max)", Red + " Quitter" + Reset})
	fmt.Println("-----------------------------------------")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Indiquez-moi ce que vous voulez faire : ")
	scanner.Scan()
	choice, _ := strconv.Atoi(scanner.Text())
	switch choice {
	case 1:
		Clear()
		p.takepot()
		p.MenuInventory()
	case 2:
		Clear()
		if cond, _ := p.findInventory("bouteille d'eau", 1); !cond {
			fmt.Println(Yellow, "item manquant", Reset)
		} else {
			poisonPot(&p.life)
			fmt.Println(Red, "-30 Pv empoisonner à l'eau de source", Reset)
			p.dead()
		}
		p.MenuInventory()
	case 3:
		Clear()
		if cond, _ := p.findInventory("chapeau de pirate", 1); !cond {
			fmt.Println(Yellow, "item manquant", Reset)
		} else {
			p.activestuff("chapeau de pirate", 10, 1)
		}
		p.MenuInventory()
	case 4:
		Clear()
		if cond, _ := p.findInventory("veste de pirate", 1); !cond {
			fmt.Println(Yellow, "item manquant", Reset)
		} else {
			p.activestuff("veste de pirate", 25, 2)
		}
		p.MenuInventory()
	case 5:
		Clear()
		if cond, _ := p.findInventory("jambe de bois", 1); !cond {
			fmt.Println(Yellow, "item manquant", Reset)
		} else {
			p.activestuff("jambe de bois", 15, 3)
		}
		p.MenuInventory()
	case 6:
		Clear()
		p.Menu()
	default:
		Clear()
		fmt.Println(Yellow, "Erreur donnée un entier compris entre 1 et 6 ", Reset)
		p.MenuInventory()
	}
}

// Fonctions utiles pour la gestion des différents inventaires
// Vérifie que le skill est bien disponible
func (p *perso) skillFind(item string) bool {
	cond := false
	for i := 0; i <= len(p.skill)-1 && !cond; i++ {
		if p.skill[i] == item {
			cond = true
		}
	}
	return cond
}

// Ajoute à l'inventaire l'item entré en paramètre
func (p *perso) addInventory(items string) {
	p.inventory = append(p.inventory, items)
}

// Supprime de l'inventaire l'élément qui se situe à la position de l'index
func (p *perso) removeInventory(index int) {
	p.inventory = append(p.inventory[:index], p.inventory[(index+1):]...)
}

// Vérifie qu'un item est présent (item) et en bonne quantité (need) dans l'inventaire
func (p *perso) findInventory(item string, need int) (bool, int) {
	var countItems, index int
	for i := 0; i <= len(p.inventory)-1; i++ {
		if p.inventory[i] == item {
			countItems++
			index = i
		}
	}
	if countItems >= need {
		return true, index
	} else {
		return false, 0
	}
}

// Fonction qui s'occupe de gérer l'amélioration limite de l'inventaire
func (p *perso) limitInventory() bool {
	var test bool
	if len(p.inventory) < p.limit {
		test = true
	} else {
		test = false
		Clear()
		fmt.Println(Yellow, "Tu as atteints la limite de charge possible", Reset)
	}
	return test
}

// Fonctions qui s'occupe de gère l'amelioration limite de l'inventaire
func (p *perso) upgradeIventory() {
	if p.limit < 40 {
		p.limit += 10
		fmt.Println("+10 emplacements dans la cale")
	} else {
		fmt.Println(Yellow, "Cale level max", Reset)
	}
}
