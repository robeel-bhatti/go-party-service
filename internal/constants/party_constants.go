package constants

type PartyIdContextKey string

const (
	ContentType                   = "application/json"
	ServiceName                   = "go-party-service"
	PartyIdKey  PartyIdContextKey = "partyId" // the unique key that stores the party ID in the request context.
)
