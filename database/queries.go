package database

import (
	_ "embed"

	"github.com/b4cktr4ck5r3/rpl-api/models"
)

var (
	//go:embed sql/selectplayers.sql
	queryPlayers string
)

func SelectPlayers() (models.Players, error) {
	result := models.Players{}

	rows, err := Db.Query(queryPlayers)

	if err != nil {
		return result, err
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
			return result, err
		}

		result.Players = append(result.Players, player)
	}

	return result, err
}
