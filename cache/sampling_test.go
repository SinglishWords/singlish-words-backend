package cache

import (
	"math/rand"
	"testing"

	"singlishwords/model"
)

func questionsWithCounts(counts []int64) []model.Question {
	questions := make([]model.Question, len(counts))
	for i, count := range counts {
		questions[i] = model.Question{Id: int64(i + 1), Word: "cue", Enable: 1, UsableCount: count}
	}
	return questions
}

func newTestRNG(seed int64) *rand.Rand {
	return rand.New(rand.NewSource(seed))
}

func assertUniqueIDs(t *testing.T, got []model.Question) {
	t.Helper()
	seen := make(map[int64]bool, len(got))
	for _, q := range got {
		if seen[q.Id] {
			t.Fatalf("duplicate question id %d in sample", q.Id)
		}
		seen[q.Id] = true
	}
}

func TestSampleQuestionsUses75To25PrimaryAllocation(t *testing.T) {
	// 20 cues => bottom 4 and top 4 pools. A batch of 4 should select
	// ceil(75%)=3 from the under-covered pool and 1 from the well-covered pool.
	counts := make([]int64, 20)
	for i := range counts {
		counts[i] = int64(i)
	}
	got := sampleQuestions(questionsWithCounts(counts), 4, newTestRNG(7))
	if len(got) != 4 {
		t.Fatalf("got %d questions, want 4", len(got))
	}
	underIDs := map[int64]bool{1: true, 2: true, 3: true, 4: true}
	wellIDs := map[int64]bool{17: true, 18: true, 19: true, 20: true}
	under, well := 0, 0
	for _, q := range got {
		if underIDs[q.Id] {
			under++
		}
		if wellIDs[q.Id] {
			well++
		}
	}
	if under != 3 || well != 1 {
		t.Fatalf("expected 3 under-covered and 1 well-covered cue; got under=%d well=%d", under, well)
	}
	assertUniqueIDs(t, got)
}

func TestSampleQuestionsFillsShortPrimaryPoolsFromOtherPrimaryThenMiddle(t *testing.T) {
	// Five cues produce one under, three middle, and one well cue. The 3:1
	// target cannot be met, so the remaining primary cue and then middle cues
	// are used to fill the batch.
	questions := questionsWithCounts([]int64{0, 1, 2, 3, 4})
	got := sampleQuestions(questions, 4, newTestRNG(11))
	if len(got) != 4 {
		t.Fatalf("got %d questions, want 4", len(got))
	}
	assertUniqueIDs(t, got)
	ids := map[int64]bool{}
	for _, q := range got {
		ids[q.Id] = true
	}
	if !ids[1] || !ids[5] {
		t.Fatalf("expected available under- and well-covered cues; got ids=%v", ids)
	}
}

func TestSampleQuestionsDoesNotRepeatQuestions(t *testing.T) {
	questions := questionsWithCounts([]int64{0, 0, 0, 1, 1, 2, 2, 3, 3, 4})
	got := sampleQuestions(questions, 10, newTestRNG(19))
	if len(got) != len(questions) {
		t.Fatalf("got %d questions, want %d", len(got), len(questions))
	}
	assertUniqueIDs(t, got)
}

func TestSampleQuestionsReturnsAllWhenLimitExceedsAvailable(t *testing.T) {
	questions := questionsWithCounts([]int64{0, 1, 2})
	got := sampleQuestions(questions, 20, newTestRNG(23))
	if len(got) != 3 {
		t.Fatalf("got %d questions, want all 3", len(got))
	}
	assertUniqueIDs(t, got)
}

func TestSampleQuestionsHandlesEmptyAndNonPositiveLimit(t *testing.T) {
	if got := sampleQuestions(nil, 5, newTestRNG(1)); len(got) != 0 {
		t.Fatalf("empty input returned %d questions", len(got))
	}
	questions := questionsWithCounts([]int64{0, 1})
	for _, limit := range []int{0, -1} {
		if got := sampleQuestions(questions, limit, newTestRNG(1)); len(got) != 0 {
			t.Fatalf("limit %d returned %d questions", limit, len(got))
		}
	}
}

func TestSampleQuestionsRandomizesTiesAcrossPercentileBoundaries(t *testing.T) {
	// With ten equally covered cues, the bottom 20% contains two randomly
	// selected tied cues. A one-cue batch selects one of those two, so every
	// cue should be selected at roughly the same rate over many requests.
	const iterations, cueCount = 5000, 10
	hits := make([]int, cueCount)
	questions := questionsWithCounts(make([]int64, cueCount))
	for seed := int64(0); seed < iterations; seed++ {
		got := sampleQuestions(questions, 1, newTestRNG(seed+100))
		if len(got) != 1 {
			t.Fatalf("seed %d: got %d questions, want 1", seed, len(got))
		}
		hits[int(got[0].Id)-1]++
	}
	expected, tolerance := iterations/cueCount, iterations/cueCount/5
	for i, count := range hits {
		if count < expected-tolerance || count > expected+tolerance {
			t.Fatalf("cue %d selected %d times; expected about %d", i+1, count, expected)
		}
	}
}

func TestSampleQuestionsExcludesDisabledQuestions(t *testing.T) {
	questions := questionsWithCounts([]int64{0, 1, 2, 3, 4})
	questions[0].Enable = 0
	got := sampleQuestions(questions, 5, newTestRNG(29))
	if len(got) != 4 {
		t.Fatalf("got %d enabled questions, want 4", len(got))
	}
	for _, q := range got {
		if q.Enable != 1 {
			t.Fatalf("selected disabled question id %d", q.Id)
		}
	}
	assertUniqueIDs(t, got)
}

func TestSampleQuestionsDoesNotMutateInput(t *testing.T) {
	questions := questionsWithCounts([]int64{3, 1, 2, 0})
	original := append([]model.Question(nil), questions...)
	_ = sampleQuestions(questions, 3, newTestRNG(31))
	for i := range questions {
		if questions[i] != original[i] {
			t.Fatalf("input question at index %d was mutated", i)
		}
	}
}
