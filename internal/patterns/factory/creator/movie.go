package creator

import (
	"fmt"

	"github.com/sh-yamaga/design-patterns-go/internal/patterns/factory/category"
	"github.com/sh-yamaga/design-patterns-go/internal/patterns/factory/work"
)

type MovieCreator struct {
}

func (mc MovieCreator) New(title, creator string) *work.Movie {
	m := mc.createMovie(title, creator)
	mc.registerMovie(m)

	return m
}

func (mc MovieCreator) createMovie(title, creator string) *work.Movie {
	return &work.Movie{
		Work: work.Work{
			Title:    title,
			Creater:  creator,
			Category: category.Movie,
		}}
}

func (mc MovieCreator) registerMovie(m *work.Movie) *work.Movie {
	fmt.Printf("registered movie: %s\n", m.Title)

	return m
}
