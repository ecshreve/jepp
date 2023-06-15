package models

const GAME_SCHEMA = `
CREATE TABLE IF NOT EXISTS game (
    game_id INT NOT NULL PRIMARY KEY,
    show_num INT NOT NULL,
    game_date DATE NOT NULL
);`

const CLUE_SCHEMA = `
CREATE TABLE IF NOT EXISTS clue (
    clue_id BIGINT NOT NULL PRIMARY KEY,
    game_id BIGINT NOT NULL,
    category_id VARCHAR(8) NOT NULL,
    question TEXT NOT NULL,
    answer VARCHAR(256) NOT NULL,
	CONSTRAINT fk_clue_game FOREIGN KEY (game_id) REFERENCES game(game_id)
);`

const CATEGORY_SCHEMA = `
CREATE TABLE IF NOT EXISTS category (
    category_id VARCHAR(8) NOT NULL,
    game_id BIGINT NOT NULL,
    name VARCHAR(256) NOT NULL,
    PRIMARY KEY (category_id, game_id),
    CONSTRAINT fk_category_game FOREIGN KEY (game_id) REFERENCES game(game_id)
);`

const COUNT_VIEW = `
CREATE OR REPLACE VIEW category_counts AS
SELECT cc.id,
       cc.name,
       game_count,
       clue_count
FROM
  (SELECT category.category_id AS id,
          category.name,
          category.game_id AS gid,
          count(*) AS clue_count
   FROM clue
   JOIN category ON clue.category_id = category.category_id
   AND clue.game_id = category.game_id
   GROUP BY category.name) AS cc
JOIN
  (SELECT category_id AS id,
          name,
          count(*) AS game_count
   FROM category
   GROUP BY category.name) AS gc ON cc.id = gc.id
AND cc.name = gc.name;`
