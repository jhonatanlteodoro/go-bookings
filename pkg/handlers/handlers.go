package handlers

import (
	"net/http"

	"github.com/jhonatanlteodoro/go-bookings/pkg/config"
	"github.com/jhonatanlteodoro/go-bookings/pkg/models"
	"github.com/jhonatanlteodoro/go-bookings/pkg/renders"
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
	renders.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	StringMap := make(map[string]string)
	StringMap["test"] = "heeey :)"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	StringMap["remote_ip"] = remoteIP

	renders.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{StringMap: StringMap})
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	StringMap := make(map[string]string)
	StringMap["test"] = "heeey :)"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	StringMap["remote_ip"] = remoteIP

	renders.RenderTemplate(w, "contact.page.tmpl", &models.TemplateData{StringMap: StringMap})
}

func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	StringMap := make(map[string]string)
	StringMap["test"] = "heeey :)"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	StringMap["remote_ip"] = remoteIP

	renders.RenderTemplate(w, "make-reservation.page.tmpl", &models.TemplateData{StringMap: StringMap})
}

func (m *Repository) SearchAvailability(w http.ResponseWriter, r *http.Request) {
	StringMap := make(map[string]string)
	StringMap["test"] = "heeey :)"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	StringMap["remote_ip"] = remoteIP

	renders.RenderTemplate(w, "search-availability.page.tmpl", &models.TemplateData{StringMap: StringMap})
}

func (m *Repository) GeneralsQuarters(w http.ResponseWriter, r *http.Request) {
	StringMap := make(map[string]string)
	StringMap["test"] = "heeey :)"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	StringMap["remote_ip"] = remoteIP

	renders.RenderTemplate(w, "generals.page.tmpl", &models.TemplateData{StringMap: StringMap})
}

func (m *Repository) MajorsSuite(w http.ResponseWriter, r *http.Request) {
	StringMap := make(map[string]string)
	StringMap["test"] = "heeey :)"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	StringMap["remote_ip"] = remoteIP

	renders.RenderTemplate(w, "majors.page.tmpl", &models.TemplateData{StringMap: StringMap})
}
