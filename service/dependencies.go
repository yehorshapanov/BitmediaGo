package service

import "github.com/yehorshapanov/BitmediaGo/db"

type Dependencies struct {
	DB db.Storer

	// define other service dependencies
}

