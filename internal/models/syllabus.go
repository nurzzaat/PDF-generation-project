package models

import "context"

type SyllabusRepository interface {
	Create(c context.Context) error
}
