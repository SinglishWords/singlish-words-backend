package cache

import (
	"math/rand"
	"sort"

	"singlishwords/model"
)
import "fmt"

// sampleQuestions selects up to limit unique questions.
//
// Questions are divided by UsableCount into bottom-20%, middle-60%, and
// top-20% pools. Equal-count questions are shuffled before the stable sort,
// giving tied questions equal chances of crossing percentile boundaries.
// The requested allocation is ceil(75%) from the under-covered pool and the
// remainder from the well-covered pool. Any shortfall is filled from unused
// primary-pool questions first, then from the middle pool.
//
// rng must be non-nil. Passing it explicitly makes sampling deterministic in
// unit tests and avoids hidden global randomness.
func sampleQuestions(questions []model.Question, limit int, rng *rand.Rand) []model.Question {
	if limit <= 0 || len(questions) == 0 {
		return []model.Question{}
	}
	if rng == nil {
		panic("sampleQuestions requires a non-nil random source")
	}

	// Keep the sampler defensive: disabled cues are never eligible, even if a
	// future caller accidentally passes them in.
	ranked := make([]model.Question, 0, len(questions))
	for _, question := range questions {
		if question.Enable == 1 {
			ranked = append(ranked, question)
		}
	}
	if len(ranked) == 0 {
		return []model.Question{}
	}
	rng.Shuffle(len(ranked), func(i, j int) {
		ranked[i], ranked[j] = ranked[j], ranked[i]
	})
	sort.SliceStable(ranked, func(i, j int) bool {
		return ranked[i].UsableCount < ranked[j].UsableCount
	})

	total := len(ranked)
	poolSize := (total + 4) / 5 // ceil(20%); cap to prevent overlap.
	if poolSize > total/2 {
		poolSize = total / 2
	}

	underCovered := ranked[:poolSize]
	middleEnd := total - poolSize
	middle := ranked[poolSize:middleEnd]
	wellCovered := ranked[middleEnd:]
	fmt.Println("underCovered: ", underCovered)
	fmt.Println("wellCovered: ", wellCovered)

	if limit > total {
		limit = total
	}
	underTarget := (3*limit + 3) / 4 // ceil(75%); remainder is well-covered.
	wellTarget := limit - underTarget

	selected := make([]model.Question, 0, limit)
	selectedIDs := make(map[int64]struct{}, limit)
	selected = appendRandomQuestions(selected, selectedIDs, underCovered, underTarget, rng)
	selected = appendRandomQuestions(selected, selectedIDs, wellCovered, wellTarget, rng)

	// If either primary pool is short, draw from the unused primary pool
	// before using the middle 60%.
	if len(selected) < limit {
		remainingPrimary := make([]model.Question, 0, len(underCovered)+len(wellCovered))
		remainingPrimary = append(remainingPrimary, underCovered...)
		remainingPrimary = append(remainingPrimary, wellCovered...)
		selected = appendRandomQuestions(selected, selectedIDs, remainingPrimary, limit-len(selected), rng)
	}
	if len(selected) < limit {
		selected = appendRandomQuestions(selected, selectedIDs, middle, limit-len(selected), rng)
	}

	return selected
}

func appendRandomQuestions(
	selected []model.Question,
	selectedIDs map[int64]struct{},
	pool []model.Question,
	requested int,
	rng *rand.Rand,
) []model.Question {
	if requested <= 0 || len(pool) == 0 {
		return selected
	}

	candidates := make([]model.Question, 0, len(pool))
	for _, question := range pool {
		if _, exists := selectedIDs[question.Id]; !exists {
			candidates = append(candidates, question)
		}
	}
	rng.Shuffle(len(candidates), func(i, j int) {
		candidates[i], candidates[j] = candidates[j], candidates[i]
	})

	if requested > len(candidates) {
		requested = len(candidates)
	}
	for _, question := range candidates[:requested] {
		selected = append(selected, question)
		selectedIDs[question.Id] = struct{}{}
	}
	return selected
}
