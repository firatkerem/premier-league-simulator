package simulation

import (
	"insider/models"
	"math/rand"
)

type LeagueSimulator interface {
	SimulateMatch(match *models.Match, home *models.Team, away *models.Team)
	CalculateChampionshipProbabilities(teams []models.Team, remainingMatches []models.Match) map[string]float64
}

type PremierLeagueSimulator struct{}

func (pls PremierLeagueSimulator) SimulateMatch(match *models.Match, home *models.Team, away *models.Team) {
	if match.IsPlayed {
		return
	}
	homeAdv := int(float64(home.Strength) * 1.1)
	awayAdv := away.Strength
	total := homeAdv + awayAdv
	homeWinProb := float64(homeAdv) / float64(total)
	awayWinProb := float64(awayAdv) / float64(total)
	r := rand.Float64()
	var homeScore, awayScore int
	if r < homeWinProb {
		homeScore = 1 + rand.Intn(3)
		awayScore = rand.Intn(2)
	} else if r < homeWinProb+awayWinProb {
		homeScore = rand.Intn(2)
		awayScore = 1 + rand.Intn(3)
	} else {
		homeScore = 1 + rand.Intn(2)
		awayScore = homeScore
	}
	match.HomeTeamScore = homeScore
	match.AwayTeamScore = awayScore
	match.IsPlayed = true
	home.Played++
	away.Played++
	home.GoalsFor += homeScore
	home.GoalsAgainst += awayScore
	away.GoalsFor += awayScore
	away.GoalsAgainst += homeScore
	if homeScore > awayScore {
		home.Won++
		home.Points += 3
		away.Lost++
	} else if homeScore < awayScore {
		away.Won++
		away.Points += 3
		home.Lost++
	} else {
		home.Drawn++
		away.Drawn++
		home.Points++
		away.Points++
	}
}

func (pls PremierLeagueSimulator) CalculateChampionshipProbabilities(teams []models.Team, remainingMatches []models.Match) map[string]float64 {
	const simulations = 1000
	championships := make(map[string]int)

	for i := 0; i < simulations; i++ {
		simulatedTeams := make([]models.Team, len(teams))
		copy(simulatedTeams, teams)

		for _, match := range remainingMatches {
			if match.IsPlayed {
				continue
			}

			var homeTeam, awayTeam *models.Team
			for j := range simulatedTeams {
				if simulatedTeams[j].ID == match.HomeTeamID {
					homeTeam = &simulatedTeams[j]
				}
				if simulatedTeams[j].ID == match.AwayTeamID {
					awayTeam = &simulatedTeams[j]
				}
			}

			if homeTeam != nil && awayTeam != nil {
				pls.SimulateMatch(&match, homeTeam, awayTeam)
			}
		}

		winner := simulatedTeams[0]
		for _, team := range simulatedTeams {
			if team.Points > winner.Points {
				winner = team
			}
		}
		championships[winner.Name]++
	}

	probabilities := make(map[string]float64)
	for team, wins := range championships {
		probabilities[team] = float64(wins) / float64(simulations) * 100
	}

	return probabilities
}
