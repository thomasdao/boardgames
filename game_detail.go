package main

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
)

func relatedGames(gameID int64) ([]*Game, error) {
	var sims []*Similarity

	// Query related games from similarities table
	err := db.
		Limit(8).
		Preload("GameA").
		Preload("GameB").
		Where("game_a_id = ?", gameID).
		Or("game_b_id = ?", gameID).
		Order("adjusted_score desc").
		Find(&sims).Error

	if err != nil {
		return nil, err
	}

	var games []*Game
	for _, sim := range sims {
		if sim.GameAID == gameID {
			games = append(games, sim.GameB)
		} else {
			games = append(games, sim.GameA)
		}
	}

	return games, nil
}

func getGame(gameID int64) (*Game, error) {
	var game Game
	err := db.First(&game, gameID).Error

	if err != nil {
		return nil, err
	}

	return &game, err
}

func getComments(gameID int64) ([]*Rating, error) {
	var ratings []*Rating
	// Query related games from similarities table
	err := db.
		Limit(8).
		Preload("User").
		Where("game_id = ?", gameID).
		Order("updated_at desc").
		Find(&ratings).Error

	if err != nil {
		return nil, err
	}

	return ratings, nil
}

func gameDetailHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/game.html")
	vars := mux.Vars(r)
	id := vars["id"]
	gameID, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		http.Error(w, "Invalid Game ID", http.StatusBadRequest)
		return
	}

	var game *Game
	game, err = getGame(gameID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	var related []*Game
	related, err = relatedGames(gameID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var ratings []*Rating
	ratings, err = getComments(gameID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.Execute(w, map[string]interface{}{
		"game":    game,
		"related": related,
		"ratings": ratings,
	})
}
