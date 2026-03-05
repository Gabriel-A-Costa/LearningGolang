package repository

import (
	"context"
	kivik "github.com/go-kivik/kivik/v4"
	"github.com/gabriel-a-costa/LearningGolang/feb/06/first_service/model"
)

type NotebookCouch struct {
	ID          string
	Name        string
	Description string
	Rev         string
}

type Repository struct {
	couchdb *kivik.DB
}

func New(c *kivik.DB) *Repository {
	return &Repository{
		couchdb: c,
	}
}

fun (repo *Repository) Create(ctx context.Context, nt model.NotebookEntityDomain) {
	notebook_couch := repository.NotebookCouch{
		ID: nt.id,
		Name: nt.Name,
		Description: nt.Description,
		Rev: "",
	}

	resp, err := repo.couchdb.Put(ctx, id, notebook_couch)
	if err != nil {
		nil, err
	}

	fmt.Println("couch resp: ", resp)
}
