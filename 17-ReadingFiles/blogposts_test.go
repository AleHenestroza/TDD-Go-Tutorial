package blogposts_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "go-tdd-tutorial/17-ReadingFiles"
)

type StubFailingFs struct{}

func (s StubFailingFs) Open(name string) (fs.File, error) {
	return nil, errors.New("Oh now, I will always fail")
}

const (
	firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
world`
	secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
B
L
T`
)

func TestNewBlogPosts(t *testing.T) {
	t.Run("read from fileSystem", func(t *testing.T) {
		fs := fstest.MapFS{
			"testfile.md":  {Data: []byte(firstBody)},
			"testfile2.md": {Data: []byte(secondBody)},
		}

		posts, err := blogposts.NewPostsFromFS(fs)

		if err != nil {
			t.Fatal(err)
		}

		assertPost(t, posts[0], blogposts.Post{
			Title:       "Post 1",
			Description: "Description 1",
			Tags:        []string{"tdd", "go"},
			Body: `Hello
world`,
		})
	})

	t.Run("read from fileSystem throws error", func(t *testing.T) {
		failingFS := StubFailingFs{}

		_, err := blogposts.NewPostsFromFS(failingFS)

		if err == nil {
			t.Errorf("expected error, but got nil")
		}
	})
}

func assertPost(t testing.TB, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
