package main

//Cyril RODRIGUES et Sam CONRAUX

// Déclaration des différentes structures (perso, equipement, monstre)
type perso struct {
	name      string
	class     string
	lifemax   int
	life      int
	level     int
	xp        int
	money     int
	inventory []string
	limit     int
	skill     []string
	speed     int
	equipement
}

type equipement struct {
	head  string
	torso string
	leg   string
}

type monstre struct {
	name    string
	lifemax int
	life    int
	attack  int
	speed   int
}

func main() {
	// Déclaration du personnage
	var p1 perso
	p1.charCreation()
	p1.Menu()
}
