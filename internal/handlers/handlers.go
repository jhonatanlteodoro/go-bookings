package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jhonatanlteodoro/go-bookings/internal/config"
	"github.com/jhonatanlteodoro/go-bookings/internal/forms"
	"github.com/jhonatanlteodoro/go-bookings/internal/models"
	"github.com/jhonatanlteodoro/go-bookings/internal/renders"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHanlders(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	renders.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	StringMap := make(map[string]string)
	StringMap["test"] = "heeey :)"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	StringMap["remote_ip"] = remoteIP

	renders.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{StringMap: StringMap})
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation
	renders.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{
		Form: forms.New(nil, []string{}, r),
		Data: data,
	})
}

func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
	}

	field_names := []string{"first_name", "last_name", "email", "phone"}
	form := forms.New(
		r.PostForm,
		field_names,
		r,
	)

	if !form.Valid() {

		data := make(map[string]interface{})
		data["reservation"] = reservation
		renders.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	// renders.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{})
	m.App.Session.Put(r.Context(), "reservation", reservation)
	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}

func (m *Repository) SearchAvailability(w http.ResponseWriter, r *http.Request) {
	StringMap := make(map[string]string)
	StringMap["test"] = "heeey :)"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	StringMap["remote_ip"] = remoteIP

	renders.RenderTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{StringMap: StringMap})
}

func (m *Repository) PostSearchAvailability(w http.ResponseWriter, r *http.Request) {
	form_data := r.Form
	fmt.Println(form_data.Get("csrf_token"))
	fmt.Println(form_data.Get("start"))
	fmt.Println(form_data.Get("end"))

	renders.RenderTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	MESSAGE string `json:"message"`
}

func (m *Repository) SearchAvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	// form_data := r.Form
	// fmt.Println(form_data.Get("csrf_token"))
	// fmt.Println(form_data.Get("start"))
	// fmt.Println(form_data.Get("end"))
	resp := jsonResponse{
		OK:      true,
		MESSAGE: "Available!",
	}
	out, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		log.Fatal("Deu ruim", err)
	}
	w.Write(out)
}

func (m *Repository) GeneralsQuarters(w http.ResponseWriter, r *http.Request) {
	StringMap := make(map[string]string)
	StringMap["test"] = "heeey :)"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	StringMap["remote_ip"] = remoteIP

	renders.RenderTemplate(w, r, "generals.page.tmpl", &models.TemplateData{StringMap: StringMap})
}

func (m *Repository) MajorsSuite(w http.ResponseWriter, r *http.Request) {
	StringMap := make(map[string]string)
	StringMap["test"] = "heeey :)"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	StringMap["remote_ip"] = remoteIP

	renders.RenderTemplate(w, r, "majors.page.tmpl", &models.TemplateData{StringMap: StringMap})
}

func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		log.Println("cannot get item from session")
		m.App.Session.Put(r.Context(), "error", "Can't load reservation, sorry")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	m.App.Session.Remove(r.Context(), "reservation")

	data := make(map[string]interface{})
	data["reservation"] = reservation
	renders.RenderTemplate(w, r, "reservation-summary.page.tmpl", &models.TemplateData{
		Data: data,
	})
}
