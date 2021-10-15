package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Gère tour à tour le combat entre le monstre et le personnage
func (p *perso) trainingFight(g *monstre) {
	round := true
	// Génère de façon aléatoirement entre 1 et 10 l'initiative du personnage et du monstre à chaque combat
	p.speed, g.speed = RandomNbr(1, 10), RandomNbr(1, 10)
	var count int
	for int(g.life) > 0 && p.life > 0 { // Permets de donner le premier coup à celui qui a l'initiative la plus importante
		count++
		if p.speed > g.speed {
			round = false
		}
		if !round { // Quand round est FALSE le personnage attaque
			round = p.charTurn(g)
		}
		if round { // Quand round est TRUE le monstre attaque
			if count%3 == 0 {
				p.life -= goblinPattern(g.attack)
				fmt.Print(g.name, " inflige à ", p.name, goblinPattern(g.attack), " de dégâts \n")
			} else {
				p.life -= g.attack
				fmt.Print(g.name, " inflige à ", p.name, " ", g.attack, " de dégâts \n")
			}
			round = false
		}
		p.lifeLevel()
		fmt.Print(g.name, " a ", g.life, " PV \n")
		fmt.Println(Yellow, "round", count, Reset)
	}
	if int(g.life) <= 0 {
		xp, money := RandomNbr(1, 35), RandomNbr(1, 25) // Génère aléatoirement la récompense à chaque combat
		p.xp += xp
		p.money += money
		Clear()
		fmt.Print("Victoire tu as gagnée en ", count, " round \n")
		fmt.Print("Tu as gagner ", xp, " xp et ", money, " golds \n")
	} else {
		Clear()
		fmt.Print("Perdue t'es vraiment un noob, tu tes fait demolire en ", count, " round \n")
		p.dead()
	}
	p.managementXp()
	p.Menu()
}

func goblinPattern(attack int) int {
	attack *= 2
	return attack
}

// Menu d'attaque du personnage
func (p *perso) charTurn(g *monstre) bool {
	fmt.Println("--------------Combat d'entrainement--------------")
	PrintMenu([]string{" Attaque Coup de point", " Attaque coup de feu ", " Inventaire"})
	fmt.Println("------------------------------------------------")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Indique moi ce que vous voulez faire : ")
	scanner.Scan()
	choice, _ := strconv.Atoi(scanner.Text())
	switch choice {
	case 1:
		Clear()
		g.life -= 8
		fmt.Println(p.name, " inflige à", g.name, 8, "de dégâts")
	case 2:
		Clear()
		if p.skillFind("coup de feu") {
			g.life -= 18
			fmt.Println(p.name, " inflige à", g.name, 18, "de dégâts")
		} else {
			fmt.Println(Yellow, "Tu ne possede pas ce skill !", Reset)
		}
	case 3:
		Clear()
		p.menuFigth(g)
	default:
		Clear()
		fmt.Println(Red, "Erreur indiquer un entier compris entre 1-2", Reset)
		p.charTurn(g)
	}
	return true
}

// Menu de l'inventaire pendant le combat
func (p *perso) menuFigth(g *monstre) {
	p.displayinventory()
	fmt.Println("------------------------------------------------")
	PrintMenu([]string{" Utiliser bouteille alcool", " Utiliser bouteille d'eau", Red + "Retour" + Reset})
	fmt.Println("------------------------------------------------")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Indique moi ce que vous voulez faire : ")
	scanner.Scan()
	choice, _ := strconv.Atoi(scanner.Text())
	switch choice {
	case 1:
		Clear()
		p.takepot()
	case 2:
		Clear()
		if cond, i := p.findInventory("bouteille d'eau", 1); cond {
			p.removeInventory(i)
			poisonPot(&g.life)
			fmt.Print(p.name, " inflige à ", g.name, " 30 de dégâts \n")
		} else {
			fmt.Println("Tu as pas d'eau du con !!")
		}
	case 3:
		p.charTurn(g)
	default:
		fmt.Println(Red, "Erreur indiquer un entier compris entre 1-3", Reset)
		p.menuFigth(g)
	}
}
