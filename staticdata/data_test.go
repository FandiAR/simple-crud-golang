package staticdata

import (
	"reflect"
	"testing"

	"github.com/FandiAR/simple-crud-golang/models"
)

func TestGetPeople(t *testing.T) {
	tests := []struct {
		name string
		want []models.Person
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPeople(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPeople() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddPerson(t *testing.T) {
	type args struct {
		person models.Person
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddPerson(tt.args.person)
		})
	}
}

func TestRemovePerson(t *testing.T) {
	type args struct {
		id int64
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RemovePerson(tt.args.id)
		})
	}
}
