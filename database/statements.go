package database

type Statement string

const (
	// InsertOutfit inserts outfits. Requires a fully populated &Outfit
	InsertOutfit Statement = "INSERT INTO outfits (id, link, submitter, tag, meta, created, updated, deleted, featured, delete_hash) VALUES (:id, :link, :submitter, :tag, :meta, :created, :updated, :deleted, :featured, :delete_hash)"

	// UpdateOutfit Requires a fully populated &Outfit. Will overwrite the outfit with the provided ID
	UpdateOutfit Statement = "UPDATE outfits SET (link, submitter, tag, meta, created, updated, deleted, featured, deleted_hash) VALUES (:link, :submitter, :tag, :meta, :created, :updated, :deleted, :featured, :delete_has) WHERE id = :id"

	// GetOutfit Requires an ID string. Returns one outfit
	GetOutfit Statement = "SELECT * FROM outfits WHERE id = :id"

	// GetOutfitFromTag Requires a Tag. Returns one outfit
	GetOutfitFromTag Statement = "SELECT * FROM outfits WHERE tag = :tag AND deleted = 0" // TODO: Way to add a LIMIT X?

	// GetRandomOutfit Requires nothing. Returns one outfit
	GetRandomOutfit Statement = "SELECT 1 FROM outfits WHERE deleted = 0 ORDER BY RANDOM()"

	// GetRandomOutfitFromSubmitter Requires a user ID. Returns one outfit
	GetRandomOutfitFromSubmitter Statement = "SELECT 1 FROM outfits WHERE submitter = :submitter AND deleted = 0 ORDER BY RANDOM()"

	// GetTopFive requires nothing. Returns 5 user IDs and the number of outfits
	GetTopFive Statement = "SELECT submitter, count(*) AS num FROM outfits WHERE deleted = false GROUP BY submitter ORDER BY num DESC LIMIT 5"

	// TODO: Implement counts, tag counts, and soft/hard deletions.
)

/* Operations? */
