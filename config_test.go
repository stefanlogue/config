package config

import (
	"testing"
	"testing/fstest"
)

func TestFindFileInFileSystem(t *testing.T) {
	t.Run("file exists in current directory", func(t *testing.T) {
		fs := fstest.MapFS{
			"currentDir/wantedfile.txt": {Data: []byte("hello")},
			"unwantedfile.txt":          {Data: []byte("goodbye")},
		}
		got, err := FindFileInFileSystem(fs, ".", "currentDir", "wantedfile.txt")
		want := "currentDir/wantedfile.txt"

		if err != nil {
			t.Errorf("got error %q, want nil", err)
		}
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("file in parent directory", func(t *testing.T) {
		fs := fstest.MapFS{
			"parentDir/wantedfile.txt":              {Data: []byte("hello")},
			"parentDir/currentDir/unwantedfile.txt": {Data: []byte("goodbye")},
		}
		got, err := FindFileInFileSystem(fs, ".", "parentDir/currentDir", "wantedfile.txt")
		want := "parentDir/wantedfile.txt"

		if err != nil {
			t.Errorf("got error %q, want nil", err)
		}
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
