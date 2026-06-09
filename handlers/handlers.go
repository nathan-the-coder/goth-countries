package handlers

import (
	"goth-countries/components"
	"goth-countries/models"
	"net/http"
	"strings"
)

func CountryList(w http.ResponseWriter, r *http.Request) {
	component := components.Countries(models.Countries)
	component.Render(r.Context(), w)
}

func CountryDetail(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	for _, country := range models.Countries {
		if country.Name == name {
			component := components.CountryDetail(country)
			component.Render(r.Context(), w)
			return
		}
	}
	http.NotFound(w, r)
}

func SearchCountries(w http.ResponseWriter, r *http.Request) {
	query := strings.ToLower(r.URL.Query().Get("search"))
	var results []models.Country

	for _, country := range models.Countries {
		if strings.Contains(strings.ToLower(country.Name), query) {
			results = append(results, country)
		}
	}

	components.CountryList(results).Render(r.Context(), w)
}
