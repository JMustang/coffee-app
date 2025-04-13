package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/JMustang/coffee-app/helpers"
	"github.com/JMustang/coffee-app/services"
	"github.com/go-chi/chi/v5"
)

var coffee services.Coffee

func GetAllCoffees(w http.ResponseWriter, r *http.Request) {

	all, err := coffee.GetAllCoffees()
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"coffees": all})
}

func GetCoffeeById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	coffee, err := coffee.GetCoffeeById(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, coffee)
}

func CreateCoffee(w http.ResponseWriter, r *http.Request) {
	var coffeeData services.Coffee
	err := json.NewDecoder(r.Body).Decode(&coffeeData)

	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

	coffeeCreated, err := coffee.CreateCoffee(coffeeData)
	// CHECK
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, coffeeCreated)
}

func UpdateCoffee(w http.ResponseWriter, r *http.Request) {
	var coffeeData services.Coffee
	id := chi.URLParam(r, "id")
	err := json.NewDecoder(r.Body).Decode(&coffeeData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	coffeUpdated, err := coffee.UpdateCoffee(id, coffeeData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}
	helpers.WriteJSON(w, http.StatusOK, coffeUpdated)
}

func DeleteCoffee(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := coffee.DeleteCoffee(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"message": "successfull deletion"})
}
