package notifier_test

import (
	"testing"

	"github.com/juliandanciu/go-feature-flag/internal/flag"
	flagv1 "github.com/juliandanciu/go-feature-flag/internal/flagv1"
	"github.com/juliandanciu/go-feature-flag/notifier"
	"github.com/juliandanciu/go-feature-flag/testutils/testconvert"
	"github.com/stretchr/testify/assert"
)

func TestDiffCache_HasDiff(t *testing.T) {
	type fields struct {
		Deleted map[string]flag.Flag
		Added   map[string]flag.Flag
		Updated map[string]notifier.DiffUpdated
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "null fields",
			fields: fields{},
			want:   false,
		},
		{
			name: "empty fields",
			fields: fields{
				Deleted: map[string]flag.Flag{},
				Added:   map[string]flag.Flag{},
				Updated: map[string]notifier.DiffUpdated{},
			},
			want: false,
		},
		{
			name: "only Deleted",
			fields: fields{
				Deleted: map[string]flag.Flag{
					"flag": &flagv1.FlagData{
						Percentage: testconvert.Float64(100),
						True:       testconvert.Interface(true),
						False:      testconvert.Interface(true),
						Default:    testconvert.Interface(true),
					},
				},
				Added:   map[string]flag.Flag{},
				Updated: map[string]notifier.DiffUpdated{},
			},
			want: true,
		},
		{
			name: "only Added",
			fields: fields{
				Added: map[string]flag.Flag{
					"flag": &flagv1.FlagData{
						Percentage: testconvert.Float64(100),
						True:       testconvert.Interface(true),
						False:      testconvert.Interface(true),
						Default:    testconvert.Interface(true),
					},
				},
				Deleted: map[string]flag.Flag{},
				Updated: map[string]notifier.DiffUpdated{},
			},
			want: true,
		},
		{
			name: "only Updated",
			fields: fields{
				Added:   map[string]flag.Flag{},
				Deleted: map[string]flag.Flag{},
				Updated: map[string]notifier.DiffUpdated{
					"flag": {
						Before: &flagv1.FlagData{
							Percentage: testconvert.Float64(100),
							True:       testconvert.Interface(true),
							False:      testconvert.Interface(true),
							Default:    testconvert.Interface(true),
						},
						After: &flagv1.FlagData{
							Percentage: testconvert.Float64(100),
							True:       testconvert.Interface(true),
							False:      testconvert.Interface(true),
							Default:    testconvert.Interface(false),
						},
					},
				},
			},
			want: true,
		},
		{
			name: "all fields",
			fields: fields{
				Added: map[string]flag.Flag{
					"flag": &flagv1.FlagData{
						Percentage: testconvert.Float64(100),
						True:       testconvert.Interface(true),
						False:      testconvert.Interface(true),
						Default:    testconvert.Interface(true),
					},
				},
				Deleted: map[string]flag.Flag{
					"flag": &flagv1.FlagData{
						Percentage: testconvert.Float64(100),
						True:       testconvert.Interface(true),
						False:      testconvert.Interface(true),
						Default:    testconvert.Interface(true),
					},
				},
				Updated: map[string]notifier.DiffUpdated{
					"flag": {
						Before: &flagv1.FlagData{
							Percentage: testconvert.Float64(100),
							True:       testconvert.Interface(true),
							False:      testconvert.Interface(true),
							Default:    testconvert.Interface(true),
						},
						After: &flagv1.FlagData{
							Percentage: testconvert.Float64(100),
							True:       testconvert.Interface(true),
							False:      testconvert.Interface(true),
							Default:    testconvert.Interface(false),
						},
					},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := notifier.DiffCache{
				Deleted: tt.fields.Deleted,
				Added:   tt.fields.Added,
				Updated: tt.fields.Updated,
			}
			assert.Equal(t, tt.want, d.HasDiff())
		})
	}
}
