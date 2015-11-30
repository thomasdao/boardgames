package main

import "time"

// User model
type User struct {
	ID       int64
	Username string `sql:"index:ix_users_username"`
	Name     string `sql:"index:ix_users_name"`
}

// Game model
type Game struct {
	ID                  int64
	Title               string `sql:"index:ix_games_title"`
	Description         string
	Designer            string  `sql:"index:ix_games_designer"`
	YearPublished       int64   `sql:"index:ix_games_year_published"`
	Rating              float64 `sql:"index:ix_games_rating"`
	NumRatings          int64   `sql:"index:ix_games_num_ratings"`
	Thumbnail           string
	Minplayers          int64  `sql:"index:ix_games_minplayers"`
	Maxplayers          int64  `sql:"index:ix_games_maxplayers"`
	SuggestedNumplayers string `sql:"index:ix_games_suggested_numplayers"`
	Minplaytime         int64  `sql:"index:ix_games_minplaytime"`
	Maxplaytime         int64  `sql:"index:ix_games_maxplaytime"`
	Minage              int64  `sql:"index:ix_games_minage"`
	Honors              string
	AmazonURL           string `gorm:"column:amazon_url"`
}

// Rating model
type Rating struct {
	ID        int64
	UserID    int64 `sql:"index:ix_ratings_user_id"`
	User      *User
	GameID    int64 `sql:"index:ix_ratings_game_id"`
	Game      *Game
	Comment   string
	Value     float64   `sql:"index:ix_ratings_value"`
	UpdatedAt time.Time `sql:"index:ix_ratings_updated_at"`
}

// Similarity model
type Similarity struct {
	ID            int64
	GameA         *Game
	GameAID       int64 `gorm:"column:game_a_id; foreignkey:game_a_id;associationforeignkey:GameA" sql:"index:ix_similarities_game_a_id"`
	GameB         *Game
	GameBID       int64   `gorm:"column:game_b_id; foreignkey:game_b_id;associationforeignkey:GameB" sql:"index:ix_similarities_game_b_id"`
	PearsonScore  float64 `sql:"index:ix_similarities_pearson_score"`
	ScoreError    float64
	AdjustedScore float64 `sql:"index:ix_similarities_adjusted_score"`
}

// TableName for Similarity table
func (sim *Similarity) TableName() string {
	return "similarities"
}
