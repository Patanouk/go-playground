package main

import "testing"

func Test_getHelloString(t *testing.T) {
	type args struct {
		name     string
		language int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Hello1", args{name: "Chris"}, "Hello Chris"},
		{"Hello2", args{name: ""}, "Hello World"},
		{"Hello2", args{name: "Elodia", language: Spanish}, "Hola Elodia"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHelloString(tt.args.name, tt.args.language); got != tt.want {
				t.Errorf("getHelloString() = %v, want %v", got, tt.want)
			}
		})
	}
}
