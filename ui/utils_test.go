package ui

import (
	"reflect"
	"testing"
	"timseriakov/ricer/pkg/types"
)

func TestFilter(t *testing.T) {
	apps := []types.App{
		{ID: "1", Name: "Firefox"},
		{ID: "2", Name: "Chromium"},
		{ID: "3", Name: "Terminal"},
		{ID: "4", Name: "Files"},
	}

	tests := []struct {
		query string
		want  []string
	}{
		{"", []string{"Firefox", "Chromium", "Terminal", "Files"}},
		{"fire", []string{"Firefox"}},
		{"TER", []string{"Terminal"}},
		{"x", []string{"Firefox"}},
		{"z", []string{}},
	}

	for _, tt := range tests {
		got := Filter(apps, tt.query)
		gotNames := extractNames(got)

		if !reflect.DeepEqual(gotNames, tt.want) {
			t.Errorf("Filter(%q) = %v, want %v", tt.query, gotNames, tt.want)
		}
	}
}

func TestFilter_Fuzzy(t *testing.T) {
	apps := []types.App{
		{ID: "1", Name: "Firefox"},
		{ID: "2", Name: "Chromium"},
		{ID: "3", Name: "Terminal"},
		{ID: "4", Name: "Files"},
	}

	tests := []struct {
		query string
		want  []string
	}{
		{"Fir", []string{"Firefox"}},
		{"Chrom", []string{"Chromium"}},
		{"Term", []string{"Terminal"}},
		{"Files", []string{"Files"}},
		{"Nope", []string{}},
	}

	for _, tt := range tests {
		got := extractNames(Filter(apps, tt.query))
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Filter(%q) = %v, want %v", tt.query, got, tt.want)
		}
	}
}

func TestMergeAndSort(t *testing.T) {
	apps := []types.App{
		{ID: "1", Name: "Firefox"},
		{ID: "2", Name: "Chromium"},
		{ID: "3", Name: "Terminal"},
		{ID: "4", Name: "Files"},
	}
	pinned := map[string]bool{
		"3": true,
		"1": true,
	}
	weights := map[string]int{
		"1": 5,
		"2": 10,
		"3": 2,
		"4": 7,
	}

	got := MergeAndSort(apps, pinned, weights)
	gotNames := extractNames(got)
	wantNames := []string{"Firefox", "Terminal", "Chromium", "Files"} // Chromium > Files по весу

	// Сортировка весов и pinned проверяется через порядок
	if !reflect.DeepEqual(gotNames, wantNames) {
		t.Errorf("MergeAndSort order = %v, want %v", gotNames, wantNames)
	}

	// Проверка: закреплённые идут первыми
	foundUnpinned := false
	for _, app := range got {
		if app.Pinned {
			if foundUnpinned {
				t.Errorf("Pinned app %s found after unpinned apps", app.Name)
			}
		} else {
			foundUnpinned = true
		}
	}
}

func extractNames(apps []types.App) []string {
	names := make([]string, len(apps))
	for i, a := range apps {
		names[i] = a.Name
	}
	return names
}
