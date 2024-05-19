package server

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/hans-m-song/depot/pkg/db"
	"github.com/hans-m-song/depot/pkg/view"
	"github.com/rs/zerolog/log"
)

func New(addr string, query *db.Queries) *http.Server {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middlewareLogger)
	r.Use(middleware.Recoverer)

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	r.Handle("/", http.RedirectHandler("/entities", http.StatusFound))

	r.Get("/entities", func(w http.ResponseWriter, r *http.Request) {
		renderRes(w, r, view.EntitiesPage())
	})

	r.Get("/entities/{entity_id}", func(w http.ResponseWriter, r *http.Request) {
		args := parseRequestArgs(r)
		renderRes(w, r, view.EntityPage(args.EntityID))
	})

	r.Get("/relationships", func(w http.ResponseWriter, r *http.Request) {
		renderRes(w, r, view.RelationshipsPage())
	})

	r.Get("/hx/entities", func(w http.ResponseWriter, r *http.Request) {
		args := parseRequestArgs(r)
		params := db.ListEntitiesParams{Limit: args.Limit, Offset: args.Offset}
		entities, err := query.ListEntities(r.Context(), params)
		if err != nil {
			renderErr(w, r, err)
			return
		}

		renderRes(w, r, view.EntitiesTable(entities, args.Limit, args.Offset))
	})

	r.Post("/hx/entities", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			renderErr(w, r, err)
			return
		}

		entityName := r.PostFormValue("entity_name")
		entityTypeIDRaw := r.PostFormValue("entity_type_id")
		entityTypeID, err := strconv.ParseInt(entityTypeIDRaw, 10, 64)
		if err != nil {
			renderErr(w, r, err)
			return
		}

		args := db.CreateEntityParams{
			EntityName:   entityName,
			EntityTypeID: entityTypeID,
		}

		entity, err := query.CreateEntity(r.Context(), args)
		if err != nil {
			renderErr(w, r, err)
			return
		}

		log.Info().Any("entity", entity).Msg("created entity")

		addTriggerHeader(w, view.EntitiesAppEvent)
		renderRes(w, r, view.EntityForm("create"))
	})

	r.Get("/hx/entities/{entity_id}", func(w http.ResponseWriter, r *http.Request) {
		args := parseRequestArgs(r)
		if args.EntityID < 0 {
			renderErr(w, r, fmt.Errorf("invalid entity id: %s", chi.URLParam(r, "entity_id")))
			return
		}

		entity, err := query.GetEntityByID(r.Context(), args.EntityID)
		if err != nil {
			renderErr(w, r, err)
			return
		}

		renderRes(w, r, view.Entity(entity))
	})

	// TODO: update
	// r.Put("/hx/entities/{entity_id}")

	r.Delete("/hx/entities/{entity_id}", func(w http.ResponseWriter, r *http.Request) {
		args := parseRequestArgs(r)
		log.Debug().Any("args", args).Send()
		time.Sleep(time.Second * 3)
		if args.EntityID < 0 {
			renderErr(w, r, fmt.Errorf("invalid entity id: %s", chi.URLParam(r, "entity_id")))
			return
		}

		if err := query.DeleteEntity(r.Context(), args.EntityID); err != nil {
			renderErr(w, r, err)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte{})
	})

	r.Get("/hx/entities/{entity_id}/attributes", func(w http.ResponseWriter, r *http.Request) {
		args := parseRequestArgs(r)
		if args.EntityID < 0 {
			renderErr(w, r, fmt.Errorf("invalid entity id: %s", chi.URLParam(r, "entity_id")))
			return
		}

		attributes, err := query.ListAttributesByEntityID(r.Context(), args.EntityID)
		if err != nil {
			renderErr(w, r, err)
			return
		}

		renderRes(w, r, view.EntityAttributes(args.EntityID, attributes))
	})

	r.Get("/hx/entities/{entity_id}/children", func(w http.ResponseWriter, r *http.Request) {
		args := parseRequestArgs(r)
		if args.EntityID < 0 {
			renderErr(w, r, fmt.Errorf("invalid entity id: %s", chi.URLParam(r, "entity_id")))
			return
		}

		params := db.ListChildrenEntitiesParams{
			EntityID: args.EntityID,
			Limit:    args.Limit,
			Offset:   args.Offset,
		}

		relationships, err := query.ListChildrenEntities(r.Context(), params)
		if err != nil {
			renderErr(w, r, err)
			return
		}

		renderRes(w, r, view.EntityChildrenRelationshipsTable(relationships))
	})

	return &http.Server{Handler: r, Addr: addr}
}
