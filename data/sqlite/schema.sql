DROP TABLE IF EXISTS "season";
CREATE TABLE IF NOT EXISTS "season" (
	"season_id" INTEGER NOT NULL  ,
	"start_date" DATE NOT NULL  ,
	"end_date" DATE NOT NULL  ,
	PRIMARY KEY ("season_id")
);

DROP TABLE IF EXISTS "category";
CREATE TABLE IF NOT EXISTS "category" (
	"category_id" INTEGER NOT NULL  ,
	"name" TEXT NOT NULL UNIQUE,
	PRIMARY KEY ("category_id")
);

DROP TABLE IF EXISTS "game";
CREATE TABLE IF NOT EXISTS "game" (
	"game_id" INTEGER NOT NULL  ,
	"season_id" INTEGER NOT NULL  ,
	"show_num" INTEGER NOT NULL  ,
	"game_date" DATE NOT NULL  ,
	"taped_date" DATE NOT NULL  ,
	PRIMARY KEY ("game_id"),
	FOREIGN KEY("season_id") REFERENCES "season" ("season_id") ON UPDATE RESTRICT ON DELETE RESTRICT
);
CREATE INDEX "fk_game_season" ON "game" ("season_id");

DROP TABLE IF EXISTS "clue";
CREATE TABLE IF NOT EXISTS "clue" (
	"clue_id" INTEGER NOT NULL  ,
	"game_id" INTEGER NOT NULL  ,
	"category_id" INTEGER NOT NULL  ,
	"question" TEXT NOT NULL  ,
	"answer" TEXT NOT NULL  ,
	PRIMARY KEY ("clue_id"),
	FOREIGN KEY("category_id") REFERENCES "category" ("category_id") ON UPDATE RESTRICT ON DELETE RESTRICT,
	FOREIGN KEY("game_id") REFERENCES "game" ("game_id") ON UPDATE RESTRICT ON DELETE RESTRICT
);
CREATE INDEX "fk_clue_category" ON "clue" ("category_id");
CREATE INDEX "fk_clue_game" ON "clue" ("game_id");
