package handlers

import (
	"github.com/benacook/PartyPlanner-Assignment/model"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)

//======================================================================================

//InviteHandler type for class definition.
type InviteHandler struct {

}

//======================================================================================

//invite struct is used to generate the html invitation via templating.
type invite struct {
	VenueName string
	GuestName string
	TableNumber int
}

//======================================================================================

//NewInviteHandler returns a InviteHandler.
func NewInviteHandler() *InviteHandler {
	return &InviteHandler{}
}

//======================================================================================

//Get handles the get http requests to the /invite route.
//It returns an invite with the guests name, table number,
//and venue with status 200 if successful.
//Status 500 is returned if the call to get the venue of guest fails.
func (h *InviteHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	v, err := model.GetVenue()
	if err != nil {
		log.Println(err)
		respondWithMessage("no venue set", http.StatusInternalServerError, w)
		return
	}

	g, err := model.GetGuest(name)
	if err != nil {
		log.Println(err)
		respondWithMessage("no guest by that name",
			http.StatusInternalServerError, w)
		return
	}

	inv := invite{v.Name, g.Name, g.TableNumber}

	tmpl := template.Must(template.ParseFiles("invite.html"))
	f, err := os.Create("./" + g.Name + ".html")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to generate invite file"))
	}
	defer f.Close()
	tmpl.Execute(f, inv)

	hdr := make(map[string]string)
	for key, value := range headers {
		hdr[key] = value
	}

	hdr["Content-Type"] = "application/octet-stream"
	hdr["Content-Disposition"] = "attachment; filename="+strconv.Quote("test.html")
	httpAddHeaders(hdr, w)

	http.ServeFile(w, r, f.Name())
	os.Remove("./" + f.Name())
}
