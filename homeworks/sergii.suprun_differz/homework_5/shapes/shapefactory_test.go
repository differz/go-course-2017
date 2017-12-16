package shapes

import (
	"reflect"
	"testing"
)

func TestCreate(t *testing.T) {
	type args struct {
		name string
		x    int
	}
	tests := []struct {
		name    string
		args    args
		want    Shaper
		wantErr bool
	}{
		{name: "cone",
			args:    args{name: "cone", x: 10},
			want:    &Cone{m: 261, r: 5, h: 10},
			wantErr: false},
		{name: "cube",
			args:    args{name: "cube", x: 10},
			want:    &Cube{m: 1000, a: 10},
			wantErr: false},
		{name: "cylinder",
			args:    args{name: "cylinder", x: 10},
			want:    &Cylinder{m: 785, r: 5, h: 10},
			wantErr: false},
		{name: "prizm",
			args:    args{name: "prizm", x: 10},
			want:    &Prizm{m: 433, a: 10, h: 10},
			wantErr: false},
		{name: "pyramid",
			args:    args{name: "pyramid", x: 10},
			want:    &Pyramid{m: 333, a: 10, h: 10},
			wantErr: false},
		{name: "sphere",
			args:    args{name: "sphere", x: 10},
			want:    &Sphere{m: 523, r: 5},
			wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Create(tt.args.name, tt.args.x)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
