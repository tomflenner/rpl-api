package database

import (
	_ "embed"

	"github.com/b4cktr4ck5r3/rpl-api/models"
)

var (
	//go:embed sql/selectplayers.sql
	queryPlayers string

	//go:embed sql/selectplayerbysteamid.sql
	queryPlayerBySteamId string

	//go:embed sql/selectplayerstop10bykd.sql
	queryPlayersTop10ByKd string

	//go:embed sql/selectplayerstop10byhs.sql
	queryPlayersTop10ByHs string
)

func SelectPlayers() ([]models.Player, error) {
	players := []models.Player{}

	rows, err := Db.Query(queryPlayers)

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

func SelectPlayerBySteamId(steamId string) (models.Player, error) {

	row := Db.QueryRow(queryPlayerBySteamId, steamId)

	player := models.Player{}
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

	if err != nil {
		return player, err
	}

	return player, err
}

func SelectPlayersTop10ByKd() ([]models.Player, error) {
	players := []models.Player{}

	rows, err := Db.Query(queryPlayersTop10ByKd)

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

func SelectPlayersTop10ByHs() ([]models.Player, error) {
	players := []models.Player{}

	rows, err := Db.Query(queryPlayersTop10ByHs)

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
