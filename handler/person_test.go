package handler

import (
	"net/http"
	"testing"
)

func TestPerson_GetPersonEndpoint(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name string
		p    *Person
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.GetPersonEndpoint(tt.args.w, tt.args.req)
		})
	}
}

func TestPerson_GetPeopleEndpoint(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name string
		p    *Person
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.GetPeopleEndpoint(tt.args.w, tt.args.req)
		})
	}
}

func TestPerson_CreatePersonEndpoint(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name string
		p    *Person
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.CreatePersonEndpoint(tt.args.w, tt.args.req)
		})
	}
}

func TestPerson_DeletePersonEndpoint(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name string
		p    *Person
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.DeletePersonEndpoint(tt.args.w, tt.args.req)
		})
	}
}
