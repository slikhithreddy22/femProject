package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/slikhithreddy22/femProject/internal/app"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", app.HealthCheck)
	r.Post("/workouts", app.WorkoutHandler.HandleCreateNewWorkout)
	r.Get("/workout/{id}", app.WorkoutHandler.HandleGetWorkoutById)
	// r.Put("/workout/id")
	return r
}
