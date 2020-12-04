package main

import (
	"reflect"
	"testing"
)

func Test_getCheapestProduct(t *testing.T) {
	testProduct1 := Product{
		ID:     1,
		Name:   "TestProduct_bollie",
		Price:  13.37,
		Source: "bollie",
	}
	testProduct2 := Product{
		ID:     1,
		Name:   "TestProduct_coolbere",
		Price:  13.37,
		Source: "coolbere",
	}
	testProduct3 := Product{
		ID:     1,
		Name:   "TestProduct_aliblabla",
		Price:  13.36,
		Source: "aliblabla",
	}

	type args struct {
		products []Product
	}
	tests := []struct {
		name    string
		args    args
		want    Product
		wantErr bool
	}{
		{
			name: "same value products",
			args: args{products: []Product{
				testProduct1,
				testProduct2,
			}},
			want:    testProduct1,
			wantErr: false,
		},
		{
			name: "lower value products",
			args: args{products: []Product{
				testProduct1,
				testProduct3,
			}},
			want:    testProduct3,
			wantErr: false,
		},
		{
			name:    "empty slice",
			args:    args{products: []Product{}},
			want:    Product{},
			wantErr: true,
		},
		{
			name:    "nil slice",
			args:    args{products: nil},
			want:    Product{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getCheapestProduct(tt.args.products)
			if (err != nil) != tt.wantErr {
				t.Errorf("getCheapestProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCheapestProduct() got = %v, want %v", got, tt.want)
			}
		})
	}
}
