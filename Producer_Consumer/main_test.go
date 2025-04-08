package main

import "testing"

func TestProducer(t *testing.T) {
	type args struct {
		ch chan int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "TestProducer",
			args: args{
				ch: make(chan int, 10),
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go func() {
				for _, v := range tt.want {
					tt.args.ch <- v
				}
				close(tt.args.ch)
			}()
			var got []int
			for v := range tt.args.ch {
				got = append(got, v)
			}
			if len(got) != len(tt.want) {
				t.Errorf("Producer() got %v, want %v", got, tt.want)
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("Producer() got %v, want %v", got, tt.want)
					break
				}
			}
		})
	}
}
func TestConsumer(t *testing.T) {
	type args struct {
		ch chan int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "TestConsumer",
			args: args{
				ch: make(chan int, 10),
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go Consumer(tt.args.ch)
			var got []int
			for v := range tt.args.ch {
				got = append(got, v)
			}
			if len(got) != len(tt.want) {
				t.Errorf("Consumer() got %v, want %v", got, tt.want)
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("Consumer() got %v, want %v", got, tt.want)
					break
				}
			}
		})
	}
}