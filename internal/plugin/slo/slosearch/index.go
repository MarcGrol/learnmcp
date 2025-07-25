package slosearch

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/sahilm/fuzzy"
	"github.com/samber/lo"

	"github.com/MarcGrol/service-catalog-mcp-server/internal/plugin/slo/repo"
)

// Index defines the interface for a search index.
//
//go:generate mockgen -source=index.go -destination=mock_index.go -package=slosearch Index
type Index interface {
	Search(ctx context.Context, keyword string, limit int) Result
}

type searchIndex struct {
	SLOs         []string
	Teams        []string
	Applications []string
}

// NewSearchIndex creates a new search index.
func NewSearchIndex(ctx context.Context, r repo.SLORepo) Index {
	slos, err := r.ListSLOs(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Error listing slos for search index")
	}
	sloNames := lo.Uniq(lo.Map(slos, func(m repo.SLO, index int) string {
		return m.UID
	}))
	teams := lo.Uniq(lo.Map(slos, func(m repo.SLO, index int) string {
		return m.Team
	}))
	applications := lo.Uniq(lo.Map(slos, func(m repo.SLO, index int) string {
		return m.Application
	}))

	return &searchIndex{
		SLOs:         sloNames,
		Teams:        teams,
		Applications: applications,
	}
}

// Result represents the search results.
type Result struct {
	SLOs         []string
	Teams        []string
	Applications []string
}

const flowSearchLimitMultiplier = 2

func (idx *searchIndex) Search(ctx context.Context, keyword string, limit int) Result {
	return Result{
		SLOs:         matchesToSlice(fuzzy.Find(keyword, idx.SLOs), limit),
		Teams:        matchesToSlice(fuzzy.Find(keyword, idx.Teams), limit),
		Applications: matchesToSlice(fuzzy.Find(keyword, idx.Applications), limit),
	}
}

func matchesToSlice(matches fuzzy.Matches, limit int) []string {
	slice := lo.Map(matches, func(item fuzzy.Match, index int) string {
		return item.Str
	})
	// Limit to top 5 per category
	return slice[0:min(len(slice), limit)]
}
