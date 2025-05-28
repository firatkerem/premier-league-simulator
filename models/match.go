package models

// Match represents a football match between two teams in a given week

type Match struct {
	ID            uint `json:"id"`
	Week          int  `json:"Week"`
	HomeTeamID    uint `json:"HomeTeamID"`
	AwayTeamID    uint `json:"AwayTeamID"`
	HomeTeamScore int  `json:"HomeTeamScore"`
	AwayTeamScore int  `json:"AwayTeamScore"`
	IsPlayed      bool `json:"IsPlayed"`
	HomeTeam      Team `gorm:"foreignKey:HomeTeamID" json:"home_team"`
	AwayTeam      Team `gorm:"foreignKey:AwayTeamID" json:"away_team"`
}
