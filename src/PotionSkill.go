package main

import (
	"fmt"
	"time"
)

//Gestion des differents item dynamique (potion, skill...)

// Potion de poison
func poisonPot(life *int) {
	for i := 0; i < 3; i++ {
		*life -= 10
		time.Sleep(1 * time.Second)
	}
}

// Potion de vie
func (p *perso) takepot() {
	cond := false
	for i := 0; i <= len(p.inventory)-1 && !cond; i++ {
		if p.inventory[i] == "bouteille de rhum" {
			cond = true
			if p.life == p.lifemax {
				fmt.Println("vos point de vie son de :", p.life, "/", p.lifemax)
				fmt.Println("vos point de vie sont déjà aux max")
			} else {
				p.life += 50
				p.removeInventory(i)
				if p.life > p.lifemax {
					p.life = p.lifemax
				}
				fmt.Println("vos points de vie son de ", p.life)
				fmt.Println("-1 bouteille de rhum")
			}
		}
	}
	if !cond {
		fmt.Println(Yellow, "Vous n'avez pas de rhum dans votre inventaire", Reset)
	}
}

// Gestion de l'achat du skill et de l'ajout dans l'iventaire des skills
func (p *perso) SpellBook() bool {
	cond := false
	for i := 0; i < len(p.skill); i++ {
		if p.skill[i] == "coup de feu" {
			cond = true
		}
	}
	if !cond {
		p.skill = append(p.skill, "coup de feu")
	}
	return cond
}
