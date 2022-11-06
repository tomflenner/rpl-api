SELECT id, 
    steam, 
    name,
    score, 
    FIND_IN_SET(score, (SELECT GROUP_CONCAT(score ORDER BY score DESC) FROM rankme)) AS rank,
    mvp,
    kills,
    deaths,
    COALESCE(ROUND((kills/deaths),2), 0) as ratio,
    headshots,
    COALESCE(ROUND((headshots/kills) * 100), 0) as headshots_percent,
    assists, 
    assist_flash,
    no_scope,
    thru_smoke,
    blind,
    wallbang 
FROM rankme
WHERE steam = ?;