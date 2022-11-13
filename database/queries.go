package database

import (
	"database/sql"
	_ "embed"

	"github.com/b4cktr4ck5r3/rpl-api/models"
)

var (
	//go:embed sql/selectplayers.sql
	queryPlayers string

	//go:embed sql/selectplayerbysteamid.sql
	queryPlayerBySteamId string

	//go:embed sql/selectplayerstopbykd.sql
	queryPlayersTopByKd string

	//go:embed sql/selectplayerstopbyhs.sql
	queryPlayersTopByHs string
)

func sqlRowToPlayer(row *sql.Row, player *models.Player) error {
	err := row.Scan(
		&player.Id,
		&player.SteamID,
		&player.Name,
		&player.Score,
		&player.Rank,
		&player.Mvp,
		&player.Kills,
		&player.Deaths,
		&player.Ratio,
		&player.Headshots,
		&player.HeadshotsPercent,
		&player.Assists,
		&player.FlashAssists,
		&player.NoScope,
		&player.ThruSmoke,
		&player.Blind,
		&player.Wallbang,
	)
	return err
}

func processPlayersQueryWithLimit(query string, limit int) ([]models.Player, error) {
	players := []models.Player{}

	rows, err := Db.Query(query, limit)

	if err != nil {
		return players, err
	}

	for rows.Next() {
		player := models.Player{}
		err := rows.Scan(
			&player.Id,
			&player.SteamID,
			&player.Name,
			&player.Score,
			&player.Rank,
			&player.Mvp,
			&player.Kills,
			&player.Deaths,
			&player.Ratio,
			&player.Headshots,
			&player.HeadshotsPercent,
			&player.Assists,
			&player.FlashAssists,
			&player.NoScope,
			&player.ThruSmoke,
			&player.Blind,
			&player.Wallbang,
		)

		if err != nil {
			break
		}

		players = append(players, player)
	}
	
	return players, err
}

func processPlayersQuery(query string) ([]models.Player, error) {
	players := []models.Player{}

	rows, err := Db.Query(query)

	if err != nil {
		return players, err
	}

	players = processPlayersDatas(rows)
	
	return players, err
}

func processPlayersDatas(rows *sql.Rows) []models.Player {
	players := []models.Player{}

	for rows.Next() {
		player := models.Player{}
		err := rows.Scan(
			&player.Id,
			&player.SteamID,
			&player.Name,
			&player.Score,
			&player.Rank,
			&player.Mvp,
			&player.Kills,
			&player.Deaths,
			&player.Ratio,
			&player.Headshots,
			&player.HeadshotsPercent,
			&player.Assists,
			&player.FlashAssists,
			&player.NoScope,
			&player.ThruSmoke,
			&player.Blind,
			&player.Wallbang,
		)

		if err != nil {
			break
		}

		players = append(players, player)
	}
	return players
}

func SelectPlayers() ([]models.Player, error) {
	return processPlayersQuery(queryPlayers)
}

func SelectPlayerBySteamId(steamId string) (models.Player, error) {
	row := Db.QueryRow(queryPlayerBySteamId, steamId)

	player := models.Player{}
	err := sqlRowToPlayer(row, &player)

	if err != nil {
		return player, err
	}

	return player, err
}

func SelectPlayersTopByKd(limit int) ([]models.Player, error) {
	return processPlayersQueryWithLimit(queryPlayersTopByKd, limit)
}

func SelectPlayersTopByHs(limit int) ([]models.Player, error) {
	return processPlayersQueryWithLimit(queryPlayersTopByHs, limit)
}
