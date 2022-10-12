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

func TestNewBlogPosts(t *testing.T) {
	t.Run("read from fileSystem", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello-github.md": {Data: []byte("Title: hello github")},
			"hello-world.md":  {Data: []byte("Title: hello world")},
		}

		posts, err := blogposts.NewPostsFromFS(fs)

		if err != nil {
			t.Fatal(err)
		}

		got := posts[0]
		want := blogposts.Post{Title: "hello github"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %+v, want %+v", got, want)
		}
	})

	t.Run("read from fileSystem throws error", func(t *testing.T) {
		failingFS := StubFailingFs{}

		_, err := blogposts.NewPostsFromFS(failingFS)

		if err == nil {
			t.Errorf("expected error, but got nil")
		}
	})
}
