package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/slikhithreddy22/femProject/internal/store"
	"github.com/slikhithreddy22/femProject/internal/utils"
)

type WorkoutHandler struct {
	workoutStore store.WorkoutStore
	logger       *log.Logger
}

func NewWorkoutHandler(workoutStore store.WorkoutStore, logger *log.Logger) *WorkoutHandler {
	return &WorkoutHandler{workoutStore: workoutStore, logger: logger}
}

func (wh *WorkoutHandler) HandleGetWorkoutById(w http.ResponseWriter, r *http.Request) {
	workoutID, err := utils.ReadIdParam(r)
	if err != nil {
		wh.logger.Printf("Error : ReadIdParam: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envolope{"error": "invalid workout id"})
		return
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
		wh.logger.Printf("Error : Unable to create a new workout: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envolope{"error": "Unable to create a new workout"})
		return
	}
	newWorkout, err := wh.workoutStore.CreateWorkout(&workout)
	if err != nil {
		wh.logger.Printf("Error : Unable to create a new workout: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envolope{"error": "Unable to create a new workout"})
	}
	utils.WriteJSON(w, http.StatusOK, utils.Envolope{"data": newWorkout})
}

func (wh *WorkoutHandler) HandleUpdateWorkout(w http.ResponseWriter, r *http.Request) {
	var workout store.Workout
	workoutId, err := utils.ReadIdParam(r)
	if err != nil {
		wh.logger.Printf("Error : ReadIdParam: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envolope{"error": "invalid workout id"})
	}
	err = json.NewDecoder(r.Body).Decode(&workout)
	if err != nil {
		wh.logger.Printf("Error : failed to update workout : %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envolope{"error": "failed to update workout"})
		return
	}
	workout.ID = int(workoutId)
	updateWorkout, err := wh.workoutStore.UpdateWorkout(&workout)
	if err != nil {
		wh.logger.Printf("Error : workout not found : %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envolope{"error": "workout not found"})
		return
	}
	utils.WriteJSON(w, http.StatusOK, utils.Envolope{"data": updateWorkout})
}

func (wh *WorkoutHandler) HandleDeleteWorkout(w http.ResponseWriter, r *http.Request) {
	workoutId, err := utils.ReadIdParam(r)
	if err != nil {
		wh.logger.Printf("Error : ReadIdParam: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envolope{"error": "invalid workout id"})
	}
	if err != nil {
		wh.logger.Printf("Error : unable to delete workout : %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envolope{"error": "unable to delete workout"})
		return
	}
	err = wh.workoutStore.DeleteWorkout(workoutId)
	if err != nil {
		return
	}
	wh.logger.Printf("Workout of id : %v deleted succesfully", workoutId)
	fmt.Println("deleted succesfully ")
	utils.WriteJSON(w, http.StatusOK, utils.Envolope{"status": "deleted succesfully"})
}
