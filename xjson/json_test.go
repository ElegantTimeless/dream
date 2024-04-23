package xjson

import "testing"

func TestIsJSON(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"cs1", args{`{"ngupf":{"nqfsx":-191970134.40492377,"ikuxvbxovjfe":true},"nkkikjdvj":{"pjeybw":"2vji","zcxfghjmtavn":1136696480.232761}}`}, true},
		{"cs2", args{`{"ngupf":{"nqfsx":-191970134.40492377,"ikuxvbxovjfe":true},"nkkikjdvj":{"pjeybw":"2vji","zcxfghjmtavn":1136696480.232761}`}, false},
		{"cs3", args{``}, false},
		{"cs4", args{`{"ngupf":{"ikuxvbxovjfe":true}"nkkikjdvj":{"pjeybw":"2vji","zcxfghjmtavn":1136696480.232761}}`}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsJSON(tt.args.str); got != tt.want {
				t.Errorf("IsJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
