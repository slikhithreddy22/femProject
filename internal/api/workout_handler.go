package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/slikhithreddy22/femProject/internal/store"
)

type WorkoutHandler struct {
	workoutStore store.WorkoutStore
}

func NewWorkoutHandler(workoutStore store.WorkoutStore) *WorkoutHandler {
	return &WorkoutHandler{workoutStore: workoutStore}
}

func (wh *WorkoutHandler) HandleGetWorkoutById(w http.ResponseWriter, r *http.Request) {
	paramsId := chi.URLParam(r, "id")
	if paramsId == "" {
		http.NotFound(w, r)
		return
	}
	workoutID, err := strconv.ParseInt(paramsId, 10, 64)
	if err != nil {
		http.NotFound(w, r)
	}
	workout, err := wh.workoutStore.GetWorkoutById(workoutID)
	if err != nil {
		http.Error(w, "workout not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(workout)
}

func (wh *WorkoutHandler) HandleCreateNewWorkout(w http.ResponseWriter, r *http.Request) {
	var workout store.Workout
	err := json.NewDecoder(r.Body).Decode(&workout)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "failed to create new workout", http.StatusInternalServerError)
		return
	}
	createWorkout, err := wh.workoutStore.CreateWorkout(&workout)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "failed to create new workout", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createWorkout)
}

func (wh *WorkoutHandler) HandleUpdateWorkout(w http.ResponseWriter, r *http.Request) {
	var workout store.Workout
	paramsId := chi.URLParam(r, "id")
	if paramsId == "" {
		http.NotFound(w, r)
		return
	}
	workoutId, err := strconv.ParseInt(paramsId, 10, 64)
	if err != nil {
		return
	}
	err = json.NewDecoder(r.Body).Decode(&workout)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "failed to update workout", http.StatusInternalServerError)
		return
	}
	workout.ID = int(workoutId)
	updateWorkout, err := wh.workoutStore.UpdateWorkout(&workout)
	if err != nil {
		http.Error(w, "workout not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updateWorkout)
}
