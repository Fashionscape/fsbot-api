// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type InputOutfit struct {
	ID         string  `json:"id"`
	Link       string  `json:"link"`
	Submitter  string  `json:"submitter"`
	Tag        string  `json:"tag"`
	Meta       *string `json:"meta"`
	Deleted    bool    `json:"deleted"`
	Featured   bool    `json:"featured"`
	DeleteHash string  `json:"delete_hash"`
}

type Outfit struct {
	ID         string  `json:"id"`
	Link       string  `json:"link"`
	Submitter  string  `json:"submitter"`
	Tag        string  `json:"tag"`
	Meta       *string `json:"meta"`
	Created    string  `json:"created"`
	Updated    string  `json:"updated"`
	Deleted    bool    `json:"deleted"`
	Featured   bool    `json:"featured"`
	DeleteHash string  `json:"delete_hash"`
}