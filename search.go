package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()["query"][0]
	var games []*Game
	query = strings.ToLower(query)
	query = fmt.Sprint(query, "%")
	err := db.
		Where("lower_title LIKE ?", query).
		Order("adjusted_rating desc").
		Limit(6).
		Find(&games).Error

	if err != nil {
		http.Error(w, "Something wrong happened", http.StatusInternalServerError)
		return
	}

	var data []interface{}
	for _, game := range games {
		item := map[string]interface{}{
			"value": game.Title,
			"data":  game.ID,
		}

		data = append(data, item)
	}

	var content []byte
	content, err = json.Marshal(data)

	// Return json to client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(content)
}
