package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/RohanSinghCode/rssagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handleCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Id uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON %v", err))
		return
	}

	feedFollow, err := apiCfg.DB.FollowFeed(r.Context(), database.FollowFeedParams{
		ID:        uuid.New(),
		Createdat: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.Id,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Following Feed failed : %v", err))
		return
	}

	respondWithJSON(w, 200, databaseFeedFollowsToFeedFollows(feedFollow))
}

func (apiCfg *apiConfig) handleGetFeedFollowers(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollow, err := apiCfg.DB.GetFollowedFeeds(r.Context(), user.ID)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't feed follows : %v", err))
		return
	}

	respondWithJSON(w, 200, databaseFeedFollowersToFeedFollowers(feedFollow))
}

func (apiCfg *apiConfig) handleDeleteFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIdStr := chi.URLParam(r, "feedFollowId")
	feedFollowId, err := uuid.Parse(feedFollowIdStr)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't parse feed follow : %v", err))
		return
	}

	err = apiCfg.DB.UnfollowFeed(r.Context(), database.UnfollowFeedParams{
		ID:     feedFollowId,
		UserID: user.ID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't delete feed follow : %v", err))
		return
	}

	respondWithJSON(w, 200, struct{}{})
}
