let teams = [];

function getTeamLogo(name) {
    // Basit örnek: Takım adına göre emoji döndür
    switch (name) {
        case 'Manchester City': return '🔵';
        case 'Liverpool': return '🔴';
        case 'Arsenal': return '🔺';
        case 'Chelsea': return '🔷';
        default: return '⚽️';
    }
}

// Function to fetch and display league table
function updateLeagueTable() {
    fetch('/api/league/table')
        .then(response => response.json())
        .then(data => {
            teams = data;
            const tableBody = document.getElementById('leagueTableBody');
            tableBody.innerHTML = '';
            teams.forEach((team, idx) => {
                const row = document.createElement('tr');
                if (idx === 0) row.classList.add('leader');
                row.innerHTML = `
                    <td>${idx + 1}</td>
                    <td>${team.Name}</td>
                    <td>${team.Played}</td>
                    <td>${team.Won}</td>
                    <td>${team.Drawn}</td>
                    <td>${team.Lost}</td>
                    <td>${team.GoalsFor}</td>
                    <td>${team.GoalsAgainst}</td>
                    <td>${team.GoalsFor - team.GoalsAgainst}</td>
                    <td>${team.Points}</td>
                `;
                tableBody.appendChild(row);
            });
        })
        .catch(error => console.error('Error:', error));
}

// Function to fetch and display matches (shows 2 matches for every completed week)
function updateMatches() {
    fetch('/api/league/matches')
        .then(response => response.json())
        .then(matches => {
            const matchesContainer = document.getElementById('matchesContainer');
            matchesContainer.innerHTML = '';

            const matchesByWeek = {};
            matches.forEach(match => {
                if (!matchesByWeek[match.Week]) {
                    matchesByWeek[match.Week] = [];
                }
                matchesByWeek[match.Week].push(match);
            });

            Object.keys(matchesByWeek).sort((a, b) => b - a).forEach(week => {
                const weekMatches = matchesByWeek[week];
                const weekDiv = document.createElement('div');
                weekDiv.className = 'week-matches';
                weekDiv.innerHTML = `<div class="week-header">Week ${week}</div>`;

                weekMatches.forEach(match => {
                    const homeTeam = teams.find(t => t.ID === match.HomeTeamID);
                    const awayTeam = teams.find(t => t.ID === match.AwayTeamID);
                    if (homeTeam && awayTeam) {
                        const matchDiv = document.createElement('div');
                        matchDiv.className = 'match';
                        matchDiv.innerHTML = `
                            <div class="team home ${match.IsPlayed ? 'played' : ''}">${homeTeam.Name}</div>
                            <div class="score-center">
                                <span class="score">${match.IsPlayed ? match.HomeTeamScore : '-'}</span>
                                <span class="score-separator">-</span>
                                <span class="score">${match.IsPlayed ? match.AwayTeamScore : '-'}</span>
                            </div>
                            <div class="team away ${match.IsPlayed ? 'played' : ''}">${awayTeam.Name}</div>
                        `;
                        weekDiv.appendChild(matchDiv);
                    }
                });

                matchesContainer.appendChild(weekDiv);
            });
        })
        .catch(error => console.error('Error:', error));
}

function updateChampionshipProbabilities() {
    fetch('/api/league/championship-probabilities')
        .then(response => response.json())
        .then(probabilities => {
            const probabilitiesContainer = document.getElementById('probabilitiesContainer');
            probabilitiesContainer.innerHTML = '';
            const sortedTeams = Object.entries(probabilities)
                .sort(([,a], [,b]) => b - a);
            const table = document.createElement('table');
            table.className = 'probabilities-table';
            table.innerHTML = `
                <thead>
                    <tr>
                        <th>Team</th>
                        <th>Probability</th>
                    </tr>
                </thead>
                <tbody>
                    ${sortedTeams.map(([team, prob]) => `
                        <tr>
                            <td>${team}</td>
                            <td style="width:60%">
                                <div class="progress-bar-bg">
                                    <div class="progress-bar" style="width:${prob}%;"></div>
                                </div>
                                <span>${prob.toFixed(1)}%</span>
                            </td>
                        </tr>
                    `).join('')}
                </tbody>
            `;
            probabilitiesContainer.appendChild(table);
        })
        .catch(error => console.error('Error:', error));
}

// Function to simulate next week
function simulateWeek() {
    fetch('/api/league/simulate-week', {
        method: 'POST'
    })
    .then(response => response.json())
    .then(data => {
        updateLeagueTable();
        updateMatches();
        updateChampionshipProbabilities();
    })
    .catch(error => console.error('Error:', error));
}

// Function to simulate all remaining matches
function simulateAll() {
    fetch('/api/league/simulate-all', {
        method: 'POST'
    })
    .then(response => response.json())
    .then(data => {
        updateLeagueTable();
        updateMatches();
        updateChampionshipProbabilities();
    })
    .catch(error => console.error('Error:', error));
}

// Function to reset league
function resetLeague() {
    fetch('/api/league/reset', {
        method: 'POST'
    })
    .then(response => response.json())
    .then(data => {
        updateLeagueTable();
        updateMatches();
        updateChampionshipProbabilities();
    })
    .catch(error => console.error('Error:', error));
}

// Initial load
document.addEventListener('DOMContentLoaded', () => {
    updateLeagueTable();
    updateMatches();
    updateChampionshipProbabilities();
}); 