package database

import "github.com/fashionscape/fsbot-api/graph/model"

type Statement string

const (
	/* Outfit statements */
	insertOutfit    Statement = "INSERT INTO outfits (id, link, submitter, tag, meta, created, updated, deleted, featured, delete_hash) VALUES (:id, :link, :submitter, :tag, :meta, :created, :updated, :deleted, :featured, :delete_hash)"
	updateOutfit    Statement = "UPDATE outfits SET (link, submitter, tag, meta, created, updated, deleted, featured, deleted_hash) VALUES (:link, :submitter, :tag, :meta, :created, :updated, :deleted, :featured, :delete_has) WHERE id = :id"
	getOutfit       Statement = "SELECT * FROM outfits WHERE id = :id"
	getRandomOutfit Statement = "SELECT 1 FROM outfits WHERE deleted = 0 ORDER BY RANDOM()"
)

/* Outfit ops */

func InsertOutfit(outfit model.Outfit) (*model.Outfit, error) {

}
