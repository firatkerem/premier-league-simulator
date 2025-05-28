package main

import (
	"fmt"
	"insider/models"
	"insider/simulation"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)

var (
	teams = []models.Team{
		{ID: 1, Name: "Manchester City", Strength: 90},
		{ID: 2, Name: "Liverpool", Strength: 85},
		{ID: 3, Name: "Arsenal", Strength: 80},
		{ID: 4, Name: "Chelsea", Strength: 75},
	}
	matches   = []models.Match{}
	simulator = simulation.PremierLeagueSimulator{}
)

func main() {
	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")

	// Initialize matches
	// Each team plays against every other team exactly twice (once home, once away)
	week := 1
	totalTeams := len(teams)

	// Create first half of the season
	for w := 0; w < totalTeams-1; w++ {
		// Create matches for this week
		for i := 0; i < totalTeams/2; i++ {
			homeTeamIndex := i
			awayTeamIndex := totalTeams - 1 - i

			// Skip if it's the same team
			if homeTeamIndex == awayTeamIndex {
				continue
			}

			// Create home match
			match := models.Match{
				Week:          week,
				HomeTeamID:    teams[homeTeamIndex].ID,
				AwayTeamID:    teams[awayTeamIndex].ID,
				IsPlayed:      false,
				HomeTeamScore: 0,
				AwayTeamScore: 0,
			}
			matches = append(matches, match)
		}

		// Rotate teams for next week
		lastTeam := teams[totalTeams-1]
		for i := totalTeams - 1; i > 1; i-- {
			teams[i] = teams[i-1]
		}
		teams[1] = lastTeam

		week++
	}

	// Create second half of the season (reverse home/away)
	for w := 0; w < totalTeams-1; w++ {
		// Create matches for this week
		for i := 0; i < totalTeams/2; i++ {
			homeTeamIndex := i
			awayTeamIndex := totalTeams - 1 - i

			// Skip if it's the same team
			if homeTeamIndex == awayTeamIndex {
				continue
			}

			// Create away match (reversed from first half)
			match := models.Match{
				Week:          week,
				HomeTeamID:    teams[awayTeamIndex].ID,
				AwayTeamID:    teams[homeTeamIndex].ID,
				IsPlayed:      false,
				HomeTeamScore: 0,
				AwayTeamScore: 0,
			}
			matches = append(matches, match)
		}

		// Rotate teams for next week
		lastTeam := teams[totalTeams-1]
		for i := totalTeams - 1; i > 1; i-- {
			teams[i] = teams[i-1]
		}
		teams[1] = lastTeam

		week++
	}

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/api/league/table", func(c *gin.Context) {
		sortedTeams := make([]models.Team, len(teams))
		copy(sortedTeams, teams)

		sort.Slice(sortedTeams, func(i, j int) bool {
			if sortedTeams[i].Points != sortedTeams[j].Points {
				return sortedTeams[i].Points > sortedTeams[j].Points
			}
			// Averaj: GoalsFor - GoalsAgainst
			avgI := sortedTeams[i].GoalsFor - sortedTeams[i].GoalsAgainst
			avgJ := sortedTeams[j].GoalsFor - sortedTeams[j].GoalsAgainst
			return avgI > avgJ
		})

		// DEBUG: Sıralamayı terminalde gör
		for _, t := range sortedTeams {
			fmt.Printf("%s: %d puan, averaj: %d\n", t.Name, t.Points, t.GoalsFor-t.GoalsAgainst)
		}

		// İlk takımın (1. takım) görünüp görünmediğini kontrol et
		if len(sortedTeams) > 0 {
			fmt.Printf("1. Takım: %s, Puan: %d\n", sortedTeams[0].Name, sortedTeams[0].Points)
		}

		c.JSON(http.StatusOK, sortedTeams)
	})

	r.GET("/api/league/matches", func(c *gin.Context) {
		c.JSON(http.StatusOK, matches)
	})

	r.GET("/api/league/championship-probabilities", func(c *gin.Context) {
		probabilities := simulator.CalculateChampionshipProbabilities(teams, matches)
		c.JSON(http.StatusOK, probabilities)
	})

	r.POST("/api/league/simulate-week", func(c *gin.Context) {
		currentWeek := 1
		for _, match := range matches {
			if match.IsPlayed {
				currentWeek = match.Week + 1
			}
		}
		for i := range matches {
			if matches[i].Week == currentWeek {
				var homeTeam, awayTeam *models.Team
				for j := range teams {
					if teams[j].ID == matches[i].HomeTeamID {
						homeTeam = &teams[j]
					}
					if teams[j].ID == matches[i].AwayTeamID {
						awayTeam = &teams[j]
					}
				}
				if homeTeam != nil && awayTeam != nil {
					simulator.SimulateMatch(&matches[i], homeTeam, awayTeam)
				}
			}
		}
		c.JSON(http.StatusOK, gin.H{"message": "Week simulated successfully"})
	})

	r.POST("/api/league/simulate-all", func(c *gin.Context) {
		for i := range matches {
			if !matches[i].IsPlayed {
				var homeTeam, awayTeam *models.Team
				for j := range teams {
					if teams[j].ID == matches[i].HomeTeamID {
						homeTeam = &teams[j]
					}
					if teams[j].ID == matches[i].AwayTeamID {
						awayTeam = &teams[j]
					}
				}
				if homeTeam != nil && awayTeam != nil {
					simulator.SimulateMatch(&matches[i], homeTeam, awayTeam)
				}
			}
		}
		c.JSON(http.StatusOK, gin.H{"message": "All matches simulated successfully"})
	})

	r.POST("/api/league/reset", func(c *gin.Context) {
		for i := range teams {
			teams[i].Played = 0
			teams[i].Won = 0
			teams[i].Drawn = 0
			teams[i].Lost = 0
			teams[i].GoalsFor = 0
			teams[i].GoalsAgainst = 0
			teams[i].Points = 0
		}
		matches = []models.Match{}
		week := 1
		totalTeams := len(teams)

		// Create first half of the season
		for w := 0; w < totalTeams-1; w++ {
			// Create matches for this week
			for i := 0; i < totalTeams/2; i++ {
				homeTeamIndex := i
				awayTeamIndex := totalTeams - 1 - i

				// Skip if it's the same team
				if homeTeamIndex == awayTeamIndex {
					continue
				}

				// Create home match
				match := models.Match{
					Week:          week,
					HomeTeamID:    teams[homeTeamIndex].ID,
					AwayTeamID:    teams[awayTeamIndex].ID,
					IsPlayed:      false,
					HomeTeamScore: 0,
					AwayTeamScore: 0,
				}
				matches = append(matches, match)
			}

			// Rotate teams for next week
			lastTeam := teams[totalTeams-1]
			for i := totalTeams - 1; i > 1; i-- {
				teams[i] = teams[i-1]
			}
			teams[1] = lastTeam

			week++
		}

		// Create second half of the season (reverse home/away)
		for w := 0; w < totalTeams-1; w++ {
			// Create matches for this week
			for i := 0; i < totalTeams/2; i++ {
				homeTeamIndex := i
				awayTeamIndex := totalTeams - 1 - i

				// Skip if it's the same team
				if homeTeamIndex == awayTeamIndex {
					continue
				}

				// Create away match (reversed from first half)
				match := models.Match{
					Week:          week,
					HomeTeamID:    teams[awayTeamIndex].ID,
					AwayTeamID:    teams[homeTeamIndex].ID,
					IsPlayed:      false,
					HomeTeamScore: 0,
					AwayTeamScore: 0,
				}
				matches = append(matches, match)
			}

			// Rotate teams for next week
			lastTeam := teams[totalTeams-1]
			for i := totalTeams - 1; i > 1; i-- {
				teams[i] = teams[i-1]
			}
			teams[1] = lastTeam

			week++
		}
		c.JSON(http.StatusOK, gin.H{"message": "League reset successfully"})
	})

	r.Run(":8080")
}
