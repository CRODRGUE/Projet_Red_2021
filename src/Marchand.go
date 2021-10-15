package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Menu du marchand qui affiche les items disponibles et permet l'achat d'items (avec gestion des erreurs)
func (p *perso) Marchand() {
	fmt.Println("--------------Marchand-------------------")
	PrintMenu([]string{" Bouteille de rhum ~3 golds", " Bouteille d'eau ~6 golds", " Apptitude : Coup de feu  ~25 golds", " Fourrure de Loup ~4 golds", " Tissu de Squellette ~7 golds", " Cuir de Sanglier ~3 golds", " Plume de Perroquet ~1 golds", " Upgrade cale (+10 emplacements) ~35 golds", Red + " Quitter" + Reset})
	fmt.Println("-----------------------------------------")
	fmt.Println("Solde : ", p.money, "golds")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Indiquez-moi ce que vous voulez faire : ")
	scanner.Scan()
	choice, _ := strconv.Atoi(scanner.Text())

	switch choice {
	case 1:
		// Vérifie la disponiblité des fond et la place dans l'inventaire
		if p.VerifMoney(3) || !p.limitInventory() {
			p.Marchand()

		} else {
			Clear()
			fmt.Println("+1 bouteille de rhum")
			p.addInventory("bouteille de rhum")
			p.money -= 3
			p.Marchand()
		}

	case 2:
		if p.VerifMoney(6) || !p.limitInventory() {
			p.Marchand()

		} else {
			Clear()
			fmt.Println("+1 bouteille d'eau")
			p.addInventory("bouteille d'eau")
			p.money -= 6
			p.Marchand()
		}
	case 3:
		Clear()
		if p.VerifMoney(25) || !p.limitInventory() {
			p.Marchand()

		} else {
			if !p.SpellBook() {
				fmt.Println("+1 apptitude coup de feu")
				p.money -= 25
			} else {
				fmt.Println("achat impossible skill déjà acheter")
			}
			p.Marchand()
		}
	case 4:
		if p.VerifMoney(4) || !p.limitInventory() {
			p.Marchand()

		} else {
			Clear()
			fmt.Println("+1 fourrure de Loup")
			p.addInventory("fourrure de loup")
			p.money -= 4
			p.Marchand()
		}
	case 5:
		Clear()
		if p.VerifMoney(7) || !p.limitInventory() {
			p.Marchand()

		} else {
			fmt.Println("+1 tissu de squellette")
			p.addInventory("tissu de squellette")
			p.money -= 7
			p.Marchand()
		}
	case 6:
		if p.VerifMoney(3) || !p.limitInventory() {
			p.Marchand()

		} else {
			Clear()
			fmt.Println("+1 cuir de sanglier")
			p.addInventory("cuir de sanglier")
			p.money -= 3
			p.Marchand()
		}
	case 7:
		if p.VerifMoney(1) || !p.limitInventory() {
			p.Marchand()

		} else {
			Clear()
			fmt.Println("+1 plume de perroquet")
			p.addInventory("plume de perroquet")
			p.money -= 1
			p.Marchand()
		}
	case 8:
		if p.VerifMoney(35) {
			p.Marchand()

		} else {
			Clear()
			p.upgradeIventory()
			p.money -= 35
			p.Marchand()
		}
	case 9:
		Clear()
		p.Menu()
	default:
		Clear()
		fmt.Println("Erreur donnée un entier compris entre 1 et 8 ")
		p.Marchand()
	}
	fmt.Println(p.money)
}

// Vérifie la disponibilité des fonds
func (p *perso) VerifMoney(prise int) bool {
	cond := false
	if p.money < prise {
		cond = true
		fmt.Println(Yellow, "Golds insuffisant", Reset)
	}
	return cond
}
