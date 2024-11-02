package tag

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		separator string
		want      map[string]string
	}{
		{
			name:      "basic key-value pairs",
			input:     "key1:value1;key2:value2",
			separator: ";",
			want: map[string]string{
				"KEY1": "value1",
				"KEY2": "value2",
			},
		},
		{
			name:      "empty value",
			input:     "key1:;key2:value2",
			separator: ";",
			want: map[string]string{
				"KEY1": "KEY1",
				"KEY2": "value2",
			},
		},
		{
			name:      "escaped separator",
			input:     "key1:value1\\;part2;key2:value2",
			separator: ";",
			want: map[string]string{
				"KEY1": "value1;part2",
				"KEY2": "value2",
			},
		},
		{
			name:      "multiple escaped separators",
			input:     "key1:value1\\;part2\\;part3;key2:value2",
			separator: ";",
			want: map[string]string{
				"KEY1": "value1;part2;part3",
				"KEY2": "value2",
			},
		},
		{
			name:      "empty input",
			input:     "",
			separator: ";",
			want:      map[string]string{},
		},
		{
			name:      "single key without value",
			input:     "key1",
			separator: ";",
			want: map[string]string{
				"KEY1": "KEY1",
			},
		},
		{
			name:      "different separator",
			input:     "key1:value1|key2:value2",
			separator: "|",
			want: map[string]string{
				"KEY1": "value1",
				"KEY2": "value2",
			},
		},
		{
			name:      "whitespace handling",
			input:     " key1 : value1 ; key2: value2",
			separator: ";",
			want: map[string]string{
				"KEY1": "value1",
				"KEY2": "value2",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Parse(tt.input, tt.separator)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitEscaped(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		separator string
		want      []string
	}{
		{
			name:      "basic split",
			input:     "part1;part2;part3",
			separator: ";",
			want:      []string{"part1", "part2", "part3"},
		},
		{
			name:      "escaped separator",
			input:     "part1\\;continued;part2",
			separator: ";",
			want:      []string{"part1;continued", "part2"},
		},
		{
			name:      "multiple escaped separators",
			input:     "part1\\;continued\\;more;part2",
			separator: ";",
			want:      []string{"part1;continued;more", "part2"},
		},
		{
			name:      "empty parts",
			input:     ";;part1;;",
			separator: ";",
			want:      []string{"", "", "part1", "", ""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := splitEscaped(tt.input, tt.separator)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitEscaped() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtractKeyValue(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantKey string
		wantVal string
	}{
		{
			name:    "basic key-value",
			input:   "key:value",
			wantKey: "KEY",
			wantVal: "value",
		},
		{
			name:    "empty value",
			input:   "key:",
			wantKey: "KEY",
			wantVal: "KEY",
		},
		{
			name:    "no colon",
			input:   "key",
			wantKey: "KEY",
			wantVal: "KEY",
		},
		{
			name:    "whitespace handling",
			input:   " key : value ",
			wantKey: "KEY",
			wantVal: "value",
		},
		{
			name:    "lowercase to uppercase key",
			input:   "lowercase:value",
			wantKey: "LOWERCASE",
			wantVal: "value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotKey, gotVal := extractKeyValue(tt.input)
			if gotKey != tt.wantKey || gotVal != tt.wantVal {
				t.Errorf("extractKeyValue() = (%v, %v), want (%v, %v)",
					gotKey, gotVal, tt.wantKey, tt.wantVal)
			}
		})
	}
}
