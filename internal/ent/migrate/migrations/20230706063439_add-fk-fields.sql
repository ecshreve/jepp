-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_clues" table
CREATE TABLE `new_clues` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `question` text NOT NULL, `answer` text NOT NULL, `category_id` integer NOT NULL, `game_id` integer NOT NULL, CONSTRAINT `clues_categories_clues` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`) ON DELETE NO ACTION, CONSTRAINT `clues_games_clues` FOREIGN KEY (`game_id`) REFERENCES `games` (`id`) ON DELETE NO ACTION);
-- Copy rows from old table "clues" to new temporary table "new_clues"
INSERT INTO `new_clues` (`id`, `question`, `answer`) SELECT `id`, `question`, `answer` FROM `clues`;
-- Drop "clues" table after copying rows
DROP TABLE `clues`;
-- Rename temporary table "new_clues" to "clues"
ALTER TABLE `new_clues` RENAME TO `clues`;
-- Create "new_games" table
CREATE TABLE `new_games` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `show` integer NOT NULL, `air_date` datetime NOT NULL, `tape_date` datetime NOT NULL, `season_id` integer NOT NULL, CONSTRAINT `games_seasons_games` FOREIGN KEY (`season_id`) REFERENCES `seasons` (`id`) ON DELETE NO ACTION);
-- Copy rows from old table "games" to new temporary table "new_games"
INSERT INTO `new_games` (`id`, `show`, `air_date`, `tape_date`) SELECT `id`, `show`, `air_date`, `tape_date` FROM `games`;
-- Drop "games" table after copying rows
DROP TABLE `games`;
-- Rename temporary table "new_games" to "games"
ALTER TABLE `new_games` RENAME TO `games`;
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
