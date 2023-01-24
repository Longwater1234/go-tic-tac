/*
 * Copyright (c) 2023, Davis Tibbz, MIT License.
 */

package player

import "testing"

func TestPlayer_HasWon(t *testing.T) {
	type fields struct {
		Name string
		Vals []int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "Won",
			fields: fields{Name: "X", Vals: []int{2, 3, 5, 8}},
			want:   true,
		},
		{
			name:   "NotWon",
			fields: fields{Name: "X", Vals: []int{0, 3, 8}},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Player{
				Name: tt.fields.Name,
				Vals: tt.fields.Vals,
			}
			if got, _ := p.HasWon(); got != tt.want {
				t.Errorf("HasWon() = %v, want %v", got, tt.want)
			}
		})
	}
}
