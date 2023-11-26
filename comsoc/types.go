package comsoc

import (
	"errors"
	"fmt"
	"sync"
)

type Alternative int
type Profile [][]Alternative
type Count map[Alternative]int

func rank(alt Alternative, prefs []Alternative) int {
	for rank, alternative := range prefs {
		if alternative == alt {
			return rank
		}
	}
	return -1
}

func IsPref(alt1, alt2 Alternative, prefs []Alternative) bool {
	return rank(alt1, prefs) < rank(alt2, prefs)
}

func maxCount(count Count) (bestAlts []Alternative) {
	bestScore := 0
	for alternative, score := range count {
		if score > bestScore {
			bestScore = score
			bestAlts = []Alternative{alternative}
		} else if score == bestScore {
			bestAlts = append(bestAlts, alternative)
		}
	}
	return bestAlts
}

func checkProfile(prefs []Alternative, alts []Alternative) error {
	altsCounter := make(map[Alternative]int, len(alts))
	for _, alternative := range prefs {
		altsCounter[alternative] += 1
		if altsCounter[alternative] > 1 {
			return errors.New("L'unicité n'est pas respecté dans ce profil.")
		}
	}
	if len(prefs) != len(alts) {
		return errors.New("L'ensemble des alternatives n'est pas représenté dans ce profil.")
	}
	return nil
}

func checkProfileAlternative(prefs Profile, alts []Alternative) error {
	for agentIndex, agentPref := range prefs {
		result := checkProfile(agentPref, alts)
		if result != nil {
			return fmt.Errorf("Le profil de l'agent %d n'est pas correct.\n%v", agentIndex, result)

		}
	}
	return nil
}

func TieBreakFactory(orderedAlts []Alternative) func([]Alternative) (Alternative, error) {
	// Retourne une fonction de tie-break personnalisée basée sur un ordre spécifié d'alternatives.
	return func(alts []Alternative) (Alternative, error) {
		// Parcours de chaque alternative dans l'ordre spécifié.
		for _, orderedAlt := range orderedAlts {
			// Compare avec les alternatives à départager.
			for _, alt := range alts {
				// Si une alternative correspond à l'ordre spécifié, elle est renvoyée.
				if orderedAlt == alt {
					return alt, nil
				}
			}
		}
		// Si aucune correspondance n'est trouvée, retourne une erreur.
		return 0, errors.New("aucune alternative correspondante trouvée")
	}
}

func CondorcetWinner(p Profile) ([]Alternative, error) {
	var wg sync.WaitGroup                           // Crée un groupe de synchronisation pour les goroutines
	winnerChan := make(chan Alternative, len(p[0])) // Crée un canal pour recevoir le gagnant potentiel

	for i := range p[0] {
		wg.Add(1)                  // Ajoute une tâche au groupe de synchronisation
		go func(alt Alternative) { // Démarre une goroutine pour chaque alternative
			defer wg.Done() // Signale que la goroutine est terminée à la fin de son exécution
			winCount := 0
			for j := range p[0] {
				if alt == Alternative(j) {
					continue // Ignore l'auto-comparaison
				}
				for _, prefs := range p {
					if IsPref(alt, Alternative(j), prefs) {
						winCount++ // Incrémente le compteur de victoires si l'alternative est préférée
						break
					}
				}
			}
			if winCount == len(p[0])-1 {
				winnerChan <- alt // Envoie l'alternative dans le canal si elle est gagnante de Condorcet
			}
		}(Alternative(i))
	}

	wg.Wait()         // Attend que toutes les goroutines soient terminées
	close(winnerChan) // Ferme le canal après la fin des goroutines

	winners := make([]Alternative, 0)
	for winner := range winnerChan { // Lit les gagnants potentiels du canal
		winners = append(winners, winner)
	}
	if len(winners) == 1 {
		return winners, nil // Retourne le gagnant de Condorcet s'il est unique
	}
	return []Alternative{}, nil // Retourne un slice vide s'il n'y a pas de gagnant de Condorcet
}
