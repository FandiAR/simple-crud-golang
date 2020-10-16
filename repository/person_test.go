package repository

import (
	"reflect"
	"testing"

	"github.com/FandiAR/simple-crud-golang/models"
)

func TestPerson_GetAll(t *testing.T) {
	tests := []struct {
		name       string
		p          *Person
		wantPeople []models.Person
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPeople, err := tt.p.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("Person.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPeople, tt.wantPeople) {
				t.Errorf("Person.GetAll() = %v, want %v", gotPeople, tt.wantPeople)
			}
		})
	}
}

func TestPerson_GetById(t *testing.T) {
	type args struct {
		id int64
	}
	tests := []struct {
		name       string
		p          *Person
		args       args
		wantPerson models.Person
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPerson, err := tt.p.GetById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Person.GetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPerson, tt.wantPerson) {
				t.Errorf("Person.GetById() = %v, want %v", gotPerson, tt.wantPerson)
			}
		})
	}
}

func TestPerson_Create(t *testing.T) {
	type args struct {
		person models.Person
	}
	tests := []struct {
		name    string
		p       *Person
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.Create(tt.args.person); (err != nil) != tt.wantErr {
				t.Errorf("Person.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPerson_Delete(t *testing.T) {
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		p       *Person
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Person.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
