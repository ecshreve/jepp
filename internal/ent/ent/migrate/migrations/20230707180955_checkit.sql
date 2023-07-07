-- Create "categories" table
CREATE TABLE `categories` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `name` text NOT NULL);
-- Create "clues" table
CREATE TABLE `clues` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `question` text NOT NULL, `answer` text NOT NULL, `category_id` integer NOT NULL, `game_id` integer NOT NULL, CONSTRAINT `clues_categories_clues` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`) ON DELETE NO ACTION, CONSTRAINT `clues_games_clues` FOREIGN KEY (`game_id`) REFERENCES `games` (`id`) ON DELETE NO ACTION);
-- Create "games" table
CREATE TABLE `games` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `show` integer NOT NULL, `air_date` datetime NOT NULL, `tape_date` datetime NOT NULL, `season_id` integer NOT NULL, CONSTRAINT `games_seasons_games` FOREIGN KEY (`season_id`) REFERENCES `seasons` (`id`) ON DELETE NO ACTION);
-- Create "seasons" table
CREATE TABLE `seasons` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `number` integer NOT NULL, `start_date` datetime NOT NULL, `end_date` datetime NOT NULL);
