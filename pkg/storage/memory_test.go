package storage

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/krlv/goweb/pkg/blog"
	"github.com/krlv/goweb/pkg/note"
)

func TestMemoryStorage_FindNotes(t *testing.T) {
	t.Run("return empty notes result set for empty storage", func(t *testing.T) {
		want := make([]*note.Note, 0)
		s := &MemoryStorage{}

		if got := s.FindNotes(); !reflect.DeepEqual(got, want) {
			t.Errorf("FindNotes() = %v, want %v", got, want)
		}
	})

	t.Run("return list of notes from storage", func(t *testing.T) {
		notes := make(map[int]*note.Note, 3)
		want := make([]*note.Note, 3)

		for i := 0; i < 3; i++ {
			n := &note.Note{ID: i, Title: fmt.Sprintf("test note %d", i)}
			notes[i], want[i] = n, n
		}

		s := &MemoryStorage{notes: notes}

		if got := s.FindNotes(); !reflect.DeepEqual(got, want) {
			t.Errorf("FindNotes() = %v, want %v", got, want)
		}
	})
}

func TestMemoryStorage_FindPages(t *testing.T) {
	t.Run("return empty pages result set for empty storage", func(t *testing.T) {
		want := make([]*blog.Page, 0)
		s := &MemoryStorage{}

		if got := s.FindPages(); !reflect.DeepEqual(got, want) {
			t.Errorf("FindPages() = %v, want %v", got, want)
		}
	})

	t.Run("return list of notes from storage", func(t *testing.T) {
		pages := make(map[string]*blog.Page, 3)
		want := make([]*blog.Page, 3)

		for i := 0; i < 3; i++ {
			slug := fmt.Sprintf("test-page-%d", i)
			n := &blog.Page{ID: i, Slug: slug}
			pages[slug], want[i] = n, n
		}

		s := &MemoryStorage{pages: pages}

		if got := s.FindPages(); !reflect.DeepEqual(got, want) {
			t.Errorf("FindPages() = %v, want %v", got, want)
		}
	})
}

func TestMemoryStorage_GetNoteByID(t *testing.T) {
	notes := make(map[int]*note.Note, 2)

	for i := 0; i < 2; i++ {
		notes[i] = &note.Note{ID: i, Title: fmt.Sprintf("test note %d", i)}
	}

	type fields struct {
		pages map[string]*blog.Page
		notes map[int]*note.Note
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *note.Note
		wantErr error
	}{
		{
			name:   "return first note",
			fields: fields{notes: notes},
			args:   args{id: 0},
			want:   notes[0],
		},
		{
			name:   "return second note",
			fields: fields{notes: notes},
			args:   args{id: 1},
			want:   notes[1],
		},
		{
			name:    "return note not found error",
			fields:  fields{notes: notes},
			args:    args{id: 4},
			wantErr: ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &MemoryStorage{
				pages: tt.fields.pages,
				notes: tt.fields.notes,
			}
			got, err := s.GetNoteByID(tt.args.id)
			if err != nil && err != tt.wantErr {
				t.Errorf("GetNoteByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNoteByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMemoryStorage_GetPageBySlug(t *testing.T) {
	pages := make(map[string]*blog.Page, 2)

	for i := 0; i < 2; i++ {
		slug := fmt.Sprintf("test-page-%d", i+1)
		pages[slug] = &blog.Page{ID: i, Slug: slug}
	}

	type fields struct {
		pages map[string]*blog.Page
		notes map[int]*note.Note
	}
	type args struct {
		slug string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *blog.Page
		wantErr error
	}{
		{
			name:   "return first page",
			fields: fields{pages: pages},
			args:   args{slug: "test-page-1"},
			want:   pages["test-page-1"],
		},
		{
			name:   "return second page",
			fields: fields{pages: pages},
			args:   args{slug: "test-page-2"},
			want:   pages["test-page-2"],
		},
		{
			name:    "return page not found error",
			fields:  fields{pages: pages},
			args:    args{slug: "page-not-found"},
			wantErr: ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &MemoryStorage{
				pages: tt.fields.pages,
				notes: tt.fields.notes,
			}
			got, err := s.GetPageBySlug(tt.args.slug)
			if err != nil && err != tt.wantErr {
				t.Errorf("GetPageBySlug() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPageBySlug() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewStorage(t *testing.T) {
	s := NewStorage()

	if got := len(s.pages); got != 3 {
		t.Errorf("NewStorage() pages count %v, want %v", got, 3)
	}

	if got := len(s.notes); got != 2 {
		t.Errorf("NewStorage() notes count %v, want %v", got, 3)
	}
}
