package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/* Tache numero 3
func (p *perso) init(name, class string, lifemax, life, level, money int, inventory, skill []string) {
	p.name = name
	p.class = class
	p.level = level
	p.lifemax = lifemax
	p.life = life
	p.money = money
	p.inventory = inventory
	p.skill = skill
}
*/

// Fonctions qui gère l'intnitialisation du monstre et du personnage
func (g *monstre) initGoblin(name string) {
	g.name = name
	g.lifemax = RandomNbr(1, 150)
	g.life = g.lifemax
	g.attack = RandomNbr(1, 15)
}

func (p *perso) charCreation() {
	p.nominate()
	p.choiseOfClass()
	p.level = 1
	p.skill = append(p.skill, "coup de poing")
	p.money = 100000000
	p.limit = 10
	p.xp = 0
	//Pour tester le changement de piece d'armure
	//p.head = "hello"
}

func (p *perso) nominate() {
	Clear()
	fmt.Println("---Début de la pérsonnalisation du personnage---")
	fmt.Println("")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Indique moi le prénoms de ton personnage : ")
	scanner.Scan()
	choice := strconv.Quote(scanner.Text())
	if choice == "" {
		p.nominate()
	} else {
		p.name = Capitalize(choice)
		fmt.Print("Bivenue à toi ", Yellow, p.name, Reset, " !!! \n")
	}
}

func (p *perso) choiseOfClass() {
	fmt.Println("--------Choix de la classe du personnage--------")
	PrintMenu([]string{" Pirate (100 points de vie max)", " Matelot (80 points de vie max)", " Capitaine (120 points de vie max)"})
	fmt.Println("------------------------------------------------")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Indiquez-moi la classe que vous désirez  : ")
	scanner.Scan()
	choice, _ := strconv.Atoi(scanner.Text())

	switch choice {
	case 1:
		p.class = "Pirate"
		p.lifemax = 100
		p.life = p.lifemax / 2
	case 2:
		p.class = "Matelot"
		p.lifemax = 80
		p.life = p.lifemax / 2
	case 3:
		p.class = "Capitaine"
		p.lifemax = 120
		p.life = p.lifemax / 2
	default:
		fmt.Println(Yellow, "Erreur donnée un entier compris entre 1 et 3 ", Reset)
		p.choiseOfClass()
	}
	fmt.Print("Bienvenu à bord du navire ", Yellow, p.class, Reset, " !! \n")
}

// Affichage des infos du personnage
func (p *perso) displayinfo() {
	Clear()
	fmt.Println("---------Information sur le personnage---------")
	fmt.Println("Nom: ", p.name)
	fmt.Println("Classe: ", p.class)
	fmt.Println("Vie max: ", p.lifemax)
	fmt.Println("Vie actuel: ", p.life)
	fmt.Println("Niveau: ", p.level)
	fmt.Println("Xp : ", p.xp, "/", p.managementXp())
	fmt.Println("Money: ", p.money)
	fmt.Println("Inventaire: ", p.inventory)
	fmt.Println("liste des skills: ", p.skill)
	fmt.Println("la limite dinevntaire >>>>", p.limit)
	fmt.Println("------------------------------------------------")
	fmt.Println(Red + "1- Retour" + Reset)
	fmt.Println("------------------------------------------------")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Indiquez-moi ce que vous voulez faire : ")
	scanner.Scan()
	choice, _ := strconv.Atoi(scanner.Text())

	switch choice {
	case 1:
		Clear()
		p.Menu()
	default:
		Clear()
		fmt.Println(Yellow, "Erreur choix impossible", Reset)
		p.displayinfo()
	}
}

// Gestion de l'xp et du level du personnage
func (p *perso) managementXp() int {
	levelSup := 100.0
	levelSup *= (float64(p.level) + 0.5)
	if p.xp > int(levelSup) {
		p.level += 1
		p.xp = p.xp - int(levelSup)
		levelSup *= (float64(p.level) + 0.5)
		fmt.Println("Tu as gagné un niveau !!, tu est niveau ", p.level)
	}
	return int(levelSup)
}

// Gestion de l'affichage en couleur de la vie (selon le niveau de vie, et variation actif selon la vie max)
func (p *perso) lifeLevel() {
	if p.lifemax*2/3 < p.life/2 {
		fmt.Println(Green, "Tu as", p.life, "pv", Reset)
	} else if p.lifemax*2/3 > p.life && p.lifemax*1/3 < p.life {
		fmt.Println(Yellow, "Tu as", p.life, "pv", Reset)
	} else {
		fmt.Println(Red, "Tu as", p.life, "pv", Reset)
	}
}

// Gestion de l'état du personnage et de sa résurrection =8
func (p *perso) dead() {
	if p.life <= 0 {
		fmt.Println(Red, "Vous êtes mort ! ", Reset)
		p.life = (p.lifemax / 2)
		fmt.Println("Vous êtes ressuscité avec ", p.life, " points de vie")
	}
}
