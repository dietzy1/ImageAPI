package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

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
	image := []core.Image{{Uuid: "", Elo: 0}, {Uuid: "", Elo: 0}}
	//Accept formdata
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	image[0].Uuid = r.FormValue("winneruuid")
	image[1].Uuid = r.FormValue("looseruuid")

	//convert to float
	image[0].Elo, err = strconv.ParseFloat(r.FormValue("winnerelo"), 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	image[1].Elo, err = strconv.ParseFloat(r.FormValue("looserelo"), 64)
	if err != nil {
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

	images, err := a.dbElo.GetLeaderBoardImages(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(images)
}
