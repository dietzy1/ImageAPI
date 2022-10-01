package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dietzy1/imageAPI/internal/application/core"
)

// Should have middleware that checks that the user is sending the request from the website itself
func (a Application) RequestMatch(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	images, err := a.dbElo.FindMatch(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode([]any{"Unable to convert file to jpg. Here is the error value:", core.Errconv(err)})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if images == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(images)
}

func (a Application) MatchResult(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	image := []core.Image{}
	err := json.NewDecoder(r.Body).Decode(&image)
	if err != nil || len(image) != 2 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	p := core.CalculateElo(image[0].Elo, image[1].Elo)

	for i := range image {
		if i == 0 {
			image[i].Elo += p
		} else {
			image[i].Elo -= p
		}
		err = a.dbImage.UpdateImage(ctx, &image[i])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
}

func (a Application) GetLeaderboard(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	//Should implement redis caching here

	images, err := a.dbElo.GetLeaderBoardImages(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(images)
}
