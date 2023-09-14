package main

import (
	"context"
	"reflect"
	"testing"

	"github.com/ServiceWeaver/weaver/weavertest"
)

func TestPig(t *testing.T) {
	runner := weavertest.Local
	runner.Test(t, func(t *testing.T, searcher Searcher) {
		ctx := context.Background()
		results, err := searcher.Search(ctx, "pig")
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(results, []string{"🐖", "🐗", "🐷", "🐽"}) {
			t.Fail()
		}
	})
}

func TestCaseInsensitivePig(t *testing.T) {
	runner := weavertest.Local
	runner.Test(t, func(t *testing.T, searcher Searcher) {
		ctx := context.Background()
		results, err := searcher.Search(ctx, "PiG")
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(results, []string{"🐖", "🐗", "🐷", "🐽"}) {
			t.Fail()
		}
	})
}

func TestBlackCat(t *testing.T) {
	runner := weavertest.Local
	runner.Test(t, func(t *testing.T, searcher Searcher) {
		ctx := context.Background()
		results, err := searcher.Search(ctx, "black cat")
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(results, []string{"🐈\u200d⬛"}) {
			t.Fail()
		}
	})
}

func TestNoResults(t *testing.T) {
	runner := weavertest.Local
	runner.Test(t, func(t *testing.T, searcher Searcher) {
		ctx := context.Background()
		results, err := searcher.Search(ctx, "foo bar baz")
		if err != nil {
			t.Fatal(err)
		}

		if len(results) > 0 {
			t.Fail()
		}
	})
}

func TestEmptyQuery(t *testing.T) {
	runner := weavertest.Local
	runner.Test(t, func(t *testing.T, searcher Searcher) {
		ctx := context.Background()
		results, err := searcher.Search(ctx, "")
		if err == nil {
			t.Fatal(err)
		}

		if len(results) > 0 {
			t.Fail()
		}
	})
}
