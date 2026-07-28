package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type WorkoutHandler struct{}

func NewWorkoutHandler() *WorkoutHandler {
	return &WorkoutHandler{}
}

func (wh *WorkoutHandler) getWorkoutById(w http.ResponseWriter, r *http.Request) {
	paramsId := chi.URLParam(r, "id")
	if paramsId == "" {
		http.NotFound(w, r)
		return
	}
	workoutID, err := strconv.ParseInt(paramsId, 10, 64)
	if err != nil {
		http.NotFound(w, r)
	}
	fmt.Fprintf(w, "the workoutid was %d\n", workoutID)
}
