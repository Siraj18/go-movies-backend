package main

import (
	"errors"
	"github.com/julienschmidt/httprouter"

	"net/http"
	"strconv"
)

func (app *application) getOneMovie(w http.ResponseWriter, r *http.Request) {
	//TODO посмотреть урок по context на ютубе
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Println(errors.New("invalid id parameter"))
		app.errorJson(w, err)
		return

	}

	movie, err := app.models.DB.Get(id)
	if err != nil {
		app.logger.Println(err)
	}

	err = app.writeJSON(w, http.StatusOK, movie, "movie")
	if err != nil {
		app.errorJson(w, err)
		return
	}
}

func (app *application) getAllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := app.models.DB.All()
	if err != nil {
		app.errorJson(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, movies, "movies")
	if err != nil {
		app.errorJson(w, err)
		return
	}
}

func (app *application) getAllGenres(w http.ResponseWriter, r *http.Request) {
	genres, err := app.models.DB.GenresAll()
	if err != nil {
		app.errorJson(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, genres, "genres")
	if err != nil {
		app.errorJson(w, err)
		return
	}
}

func (app *application) getAllMoviesByGenre(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	genreID, err := strconv.Atoi(params.ByName("genre_id"))

	if err != nil {
		app.errorJson(w, err)
		return
	}

	movies, err := app.models.DB.All(genreID)
	if err != nil {
		app.errorJson(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, movies, "movies")

	if err != nil {
		app.errorJson(w, err)
		return
	}
}

func (app *application) deleteMovie(w http.ResponseWriter, r *http.Request) {

}
func (app *application) insertMovie(w http.ResponseWriter, r *http.Request) {

}
func (app *application) updateMovie(w http.ResponseWriter, r *http.Request) {

}
func (app *application) searchMovie(w http.ResponseWriter, r *http.Request) {

}
