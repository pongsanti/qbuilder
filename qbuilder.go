package qbuilder

import (
	"fmt"
	"time"

	"github.com/volatiletech/sqlboiler/queries/qm"
)

type Builder struct {
	queries []qm.QueryMod
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) WhereString(sqlTemplate string, value string) *Builder {
	if value != "" {
		b.queries = append(b.queries, qm.Where(fmt.Sprintf(sqlTemplate, value)))
	}
	return b
}

func (b *Builder) WhereDate(sqlTemplate string, value time.Time) *Builder {
	if !value.IsZero() {
		b.queries = append(b.queries, qm.Where(sqlTemplate, value))
	}
	return b
}

func (b *Builder) WhereInt(sqlTemplate string, value int) *Builder {
	if value != 0 {
		b.queries = append(b.queries, qm.Where(sqlTemplate, value))
	}
	return b
}

func (b *Builder) WithQMods(qmods []qm.QueryMod) *Builder {
	if len(qmods) > 0 {
		b.queries = append(b.queries, qmods...)
	}
	return b
}

func (b *Builder) WithQMod(qmod qm.QueryMod) *Builder {
	if qmod != nil {
		b.queries = append(b.queries, qmod)
	}
	return b
}

func (b *Builder) WithLimit(limit int) *Builder {
	b.queries = append(b.queries, qm.Limit(limit))
	return b
}

func (b *Builder) WithOffset(offset int) *Builder {
	b.queries = append(b.queries, qm.Offset(offset))
	return b
}

func (b *Builder) Build() []qm.QueryMod {
	return b.queries
}
