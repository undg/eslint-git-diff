package run

import (
	"reflect"
	"testing"
)

func TestFilterDotFiles(t *testing.T) {
	tests := []struct {
		name     string
		files    []string
		path     string
		expected []string
	}{
		{
			name:     "no dot files",
			files:    []string{"file1.txt", "file2.txt"},
			path: ".",
			expected: []string{"file1.txt", "file2.txt"},
		},
		{
			name:     "all dot files",
			files:    []string{".file1", ".file2"},
			path: ".",
			expected: []string{},
		},
		{
			name:     "mixed files",
			files:    []string{"xfile1", ".xfile2", "xfile3"},
			path: ".",
			expected: []string{"xfile1", "xfile3"},
		},
		{
			name:     "files in deep dir",
			files:    []string{"xfile1", "src/.xfile2", "xfile3"},
			path: ".",
			expected: []string{"xfile1", "xfile3"},
		},
		{
			name:     "husky dot file in deep dir",
			files:    []string{".husky/post-checkout", "src/xfile1", "src/xfile3"},
			path: "src",
			expected: []string{"src/xfile1", "src/xfile3"},
		},
		{
			name:     "husky dot file in deep dir",
			files:    []string{"pnpm-lock.yaml", "package.json", "src/xfile1", "src/xfile3"},
			path: "src",
			expected: []string{"src/xfile1", "src/xfile3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := filterFiles(tt.files, tt.path)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("filterDotFiles(%v) = %v, want %v", tt.files, result, tt.expected)
			}
		})
	}
}
