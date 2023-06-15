package models

const GAME_SCHEMA = `
CREATE TABLE IF NOT EXISTS game (
    game_id INT NOT NULL PRIMARY KEY,
    show_num INT NOT NULL,
    game_date DATE NOT NULL
);`

const CLUE_SCHEMA = `
CREATE TABLE IF NOT EXISTS clue (
    clue_id VARCHAR(64) NOT NULL,
    game_id INT NOT NULL,
    category VARCHAR(256) NOT NULL,
    question TEXT NOT NULL,
    answer VARCHAR(256) NOT NULL,
    PRIMARY KEY (clue_id, game_id),
		CONSTRAINT fk_clue_game FOREIGN KEY (game_id) REFERENCES game(game_id)
);`
