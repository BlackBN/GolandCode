package factory

import (
	"reflect"
	"testing"
)

func TestGetFoodCreateMachine(t *testing.T) {
	tests := []struct {
		rawMaterial string
		want        FoodCreateMachine
	}{
		{
			rawMaterial: "milk",
			want:        &milkCreateMachine{},
		},
		{
			rawMaterial: "flour",
			want:        &breadCreateMachine{},
		},
	}
	for _, test := range tests {
		t.Run(test.rawMaterial, func(t *testing.T) {
			if got := GetFoodCreateMachine(test.rawMaterial); !reflect.DeepEqual(got, test.want) {
				t.Errorf("GetFoodCreateMachine() = %v, want %v", got, test.want)
			} else {
				got.CreateFood(test.rawMaterial)
			}

		})
	}
}

func TestGetFoodCreateMachineFactory(t *testing.T) {
	tests := []struct {
		rawMaterial string
		want        FoodCreateMachineFactory
	}{
		{
			rawMaterial: "milk",
			want:        &milkCreateMachineFactory{},
		},
		{
			rawMaterial: "flour",
			want:        &breadCreateMachineFactory{},
		},
	}
	for _, test := range tests {
		t.Run(test.rawMaterial, func(t *testing.T) {
			if got := GetFoodCreateMachineFactory(test.rawMaterial); !reflect.DeepEqual(got, test.want) {
				t.Errorf("GetFoodCreateMachineFactory() = %v, want = %v", got, test.want)
			} else {
				got.CreateFoodCreateMachine().CreateFood(test.rawMaterial)
			}
		})
	}
}
