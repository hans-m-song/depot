package view

type AppEvent string

func (e AppEvent) String() string { return string(e) }

const (
	EntitiesAppEvent      AppEvent = "app-entities"
	RelationshipsAppEvent AppEvent = "app-relationships"
)
