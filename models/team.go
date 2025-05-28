package models

// Team represents a football team in the league
// Only 4 teams in the league
// Premier League rules apply

type Team struct {
	ID           uint   `json:"ID"`
	Name         string `json:"Name"`
	Strength     int    `json:"Strength"`
	Played       int    `json:"Played"`
	Won          int    `json:"Won"`
	Drawn        int    `json:"Drawn"`
	Lost         int    `json:"Lost"`
	GoalsFor     int    `json:"GoalsFor"`
	GoalsAgainst int    `json:"GoalsAgainst"`
	Points       int    `json:"Points"`
}
