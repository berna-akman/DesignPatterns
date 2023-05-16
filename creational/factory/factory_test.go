package factory

import (
	"reflect"
	"testing"
)

func TestGetPizza(t *testing.T) {
	type args struct {
		pizzaType string
	}
	tests := []struct {
		name    string
		args    args
		want    IPizza
		wantErr bool
	}{
		{
			name: "should_return_neapolitan_pizza_details",
			args: args{
				pizzaType: "Neapolitan",
			},
			want:    NewNeapolitan(),
			wantErr: false,
		},
		{
			name: "should_return_ny_style_pizza_details",
			args: args{
				pizzaType: "New York Style",
			},
			want:    NewNYStyle(),
			wantErr: false,
		},
		{
			name: "should_return_margherita_pizza_details",
			args: args{
				pizzaType: "Margherita",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPizza(tt.args.pizzaType)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPizza() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPizza() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPizza_getCookingTime(t *testing.T) {
	type fields struct {
		Name        string
		CookingTime uint
	}
	tests := []struct {
		name   string
		fields fields
		want   uint
	}{
		{
			name: "should_return_cooking_time",
			fields: fields{
				CookingTime: 4,
			},
			want: 4,
		},
		{
			name: "should_return_zero_when_cooking_time_is_zero",
			fields: fields{
				CookingTime: 0,
			},
			want: 0,
		},
		{
			name:   "should_return_zero_when_cooking_time_is_nil",
			fields: fields{},
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pizza{
				Name:        tt.fields.Name,
				CookingTime: tt.fields.CookingTime,
			}
			if got := p.getCookingTime(); got != tt.want {
				t.Errorf("getCookingTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPizza_getName(t *testing.T) {
	type fields struct {
		Name        string
		CookingTime uint
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "should_return_name",
			fields: fields{
				Name: "sample name",
			},
			want: "sample name",
		},
		{
			name: "should_return_empty_when_name_is_empty",
			fields: fields{
				Name: "",
			},
			want: "",
		},
		{
			name:   "should_return_empty_when_name_is_nil",
			fields: fields{},
			want:   "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pizza{
				Name:        tt.fields.Name,
				CookingTime: tt.fields.CookingTime,
			}
			if got := p.getName(); got != tt.want {
				t.Errorf("getName() = %v, want %v", got, tt.want)
			}
		})
	}
}
