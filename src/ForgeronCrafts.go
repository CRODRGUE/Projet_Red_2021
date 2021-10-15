package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Menu du Forgeron (qui permet de craft et d'afficher le livre de craft)
func (p *perso) forgeron() {
	fmt.Println("-----------------Forgeron----------------")
	PrintMenu([]string{" Chapeaux de pirate (+ 10 points de vie max)", " Veste de pirate (+ 25 points de vie max)", " Jambe de bois (+ 15 points de vie max)", " Livre de crafts", Red + " Quitter" + Reset})
	fmt.Println("-----------------------------------------")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Indiquez-moi ce que vous voulez faire : ")
	scanner.Scan()
	choice, _ := strconv.Atoi(scanner.Text())

	switch choice {
	case 1:
		Clear()
		p.hataventurer()
		p.forgeron()
	case 2:
		Clear()
		p.vestaventurer()
		p.forgeron()
	case 3:
		Clear()
		p.bootaventurer()
		p.forgeron()
	case 4:
		Clear()
		LivreDeCraft()
		p.forgeron()
	case 5:
		Clear()
		p.Menu()
	default:
		Clear()
		fmt.Println(Yellow, "Erreur donnée un entier compris entre 1 et 4 ", Reset)
		p.forgeron()
	}
}

func LivreDeCraft() {
	fmt.Println("--------------Livre de crafts-------------")
	fmt.Println("Chapeau de pirate")
	fmt.Println("    -->  1 plume de perroquet 1 cuir de sanglier")
	fmt.Println("Veste de pirate")
	fmt.Println("    -->  2 fourrure de loup 1 tissu de squelette")
	fmt.Println("Jambe de bois")
	fmt.Println("    -->  1 fourrure de loup 1 cuir de sanglier")
}

// Fonction qui gère les crafts des différentes pièces d'armure (vérifie, supprime les ressources demandées et ajoute la pièce d'armure) et gère les erreurs
func (p *perso) hataventurer() {
	condItem1, _ := p.findInventory("plume de perroquet", 1)
	condItem2, _ := p.findInventory("cuir de sanglier", 1)
	var counter1, counter2 int
	if condItem1 && condItem2 {
		fmt.Println("+1 Chapeau de pirate")
		p.addInventory("chapeau de pirate")
		for index := len(p.inventory) - 1; index >= 0; index-- {
			if p.inventory[index] == "plume de perroquet" && counter1 == 0 {
				p.removeInventory(index)
				counter1++
			}
			if p.inventory[index] == "cuir de sanglier" && counter2 == 0 {
				p.removeInventory(index)
				counter2++
			}
		}
	} else {
		fmt.Println(Yellow, "resources insufissants", Reset)
	}
}

func (p *perso) vestaventurer() {
	condItem1, _ := p.findInventory("tissu de squellette", 1)
	condItem2, _ := p.findInventory("fourrure de loup", 2)
	var counter1, counter2 int
	if condItem1 && condItem2 {
		p.addInventory("veste de pirate")
		fmt.Println("+1 veste de pirate")
		for index := len(p.inventory) - 1; index >= 0; index-- {
			if p.inventory[index] == "tissu de squellette" && counter1 == 0 {
				p.removeInventory(index)
				counter1++
			}
			if p.inventory[index] == "fourrure de loup" && counter2 < 2 {
				p.removeInventory(index)
				counter2++
			}
		}
	} else {
		fmt.Println(Yellow, "resources insufissants", Reset)
	}
}

func (p *perso) bootaventurer() {
	condItem1, _ := p.findInventory("fourrure de loup", 1)
	condItem2, _ := p.findInventory("cuir de sanglier", 1)
	var counter1, counter2 int
	if condItem1 && condItem2 {
		p.addInventory("jambe de bois")
		fmt.Println("+1 Jambe de bois")
		for index := len(p.inventory) - 1; index >= 0; index-- {
			if p.inventory[index] == "fourrure de loup" && counter1 == 0 {
				p.removeInventory(index)
				counter1++
			}
			if p.inventory[index] == "cuir de sanglier" && counter2 == 0 {
				p.removeInventory(index)
				counter2++
			}
		}
	} else {
		fmt.Println(Yellow, "resources insufissants", Reset)
	}
}

// Fonction qui s'occupe de gèrer l'équipement des pièces d'armures (avec gestion du changement entre celle déjà équiper et la nouvelle)
func (p *perso) activestuff(item string, lifeup, part int) {
	count := 0
	switch part {
	case 1:
		for i := 0; i <= len(p.inventory)-1; i++ {
			if p.inventory[i] == item && count == 0 {
				count++
				if p.head != "" {
					p.addInventory(p.head)
				}
				p.removeInventory(i)
				p.lifemax += lifeup
				p.head = item
			}
		}
	case 2:
		for i := 0; i <= len(p.inventory)-1; i++ {
			if p.inventory[i] == item && count == 0 {
				count++
				if p.torso != "" {
					p.addInventory(p.torso)
				}
				p.removeInventory(i)
				p.lifemax += lifeup
				p.torso = item
			}
		}
	case 3:
		for i := 0; i <= len(p.inventory)-1; i++ {
			if p.inventory[i] == item && count == 0 {
				count++
				if p.torso != "" {
					p.addInventory(p.leg)
				}
				p.removeInventory(i)
				p.lifemax += lifeup
				p.leg = item
			}
		}
	}
}
