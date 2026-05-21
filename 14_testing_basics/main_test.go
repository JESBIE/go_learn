package main

import (
	"errors"
	"testing"
)

func devide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func TestDevide(t *testing.T) {
	cases := []struct {
		name    string
		a, b    int
		want    int
		wantErr bool
	}{
		{name: "正常情况", a: 6, b: 2, want: 3, wantErr: false},
		{name: "除数为0", a: 6, b: 0, want: 0, wantErr: true},
		{name: "负数", a: -6, b: 2, want: -3, wantErr: false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := devide(tc.a, tc.b)
			if tc.wantErr && err == nil {
				t.Fatalf("devide(%d, %d) = %d, want error", tc.a, tc.b, got)
			}

			if !tc.wantErr && err != nil {
				t.Fatalf("devide(%d, %d) = %d, want error", tc.a, tc.b, got)
			}

			if !tc.wantErr && got != tc.want {
				t.Fatalf("devide(%d, %d) = %d, want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}

}

func main() {
	t := testing.T{}
	TestDevide(&t)
}
