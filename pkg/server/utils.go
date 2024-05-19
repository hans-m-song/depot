package server

import (
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/hans-m-song/depot/pkg/view"
)

func addTriggerHeader(w http.ResponseWriter, event view.AppEvent) {
	w.Header().Add("HX-Trigger", event.String())
}

func renderRes(w http.ResponseWriter, r *http.Request, content templ.Component) {
	args := parseRequestArgs(r)
	w.WriteHeader(http.StatusOK)

	if !args.HTMX {
		content = view.Layout(content)
	}

	content.Render(r.Context(), w)
}

func renderErr(w http.ResponseWriter, r *http.Request, err error) {
	args := parseRequestArgs(r)
	w.WriteHeader(http.StatusOK)

	content := view.ErrorMessage(err)
	if !args.HTMX {
		content = view.Layout(content)
	}

	content.Render(r.Context(), w)
}

type RequestArgs struct {
	ReqID    string
	EntityID int64
	Limit    int64
	Offset   int64
	HTMX     bool
}

func parseRequestArgs(r *http.Request) RequestArgs {
	query := r.URL.Query()

	args := RequestArgs{
		HTMX:  r.Header.Get("hx-request") == "true",
		ReqID: middleware.GetReqID(r.Context()),

		EntityID: -1,
		Limit:    10,
		Offset:   0,
	}

	if raw := chi.URLParam(r, "entity_id"); raw != "" {
		if entityID, err := strconv.ParseInt(raw, 10, 64); err == nil {
			args.EntityID = entityID
		}
	}

	if raw := query.Get("limit"); raw != "" {
		if limit, err := strconv.ParseInt(raw, 10, 64); err == nil {
			args.Limit = limit
		}
	}

	if raw := query.Get("offset"); raw != "" {
		if offset, err := strconv.ParseInt(raw, 10, 64); err == nil {
			args.Offset = offset
		}
	}

	return args
}
