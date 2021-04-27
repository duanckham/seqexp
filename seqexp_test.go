package seqexp

import "testing"

func TestSeqexp_Match_monotone(t *testing.T) {
	type args struct {
		seq *Seq
	}
	tests := []struct {
		name string
		s    *Seqexp
		args args
		want bool
	}{
		{
			name: "monotone line 1",
			s: &Seqexp{
				expression: "m",
			},
			args: args{
				seq: &Seq{1, 1, 1, 1},
			},
			want: true,
		},
		{
			name: "monotone line 2",
			s: &Seqexp{
				expression: "m",
			},
			args: args{
				seq: &Seq{1, 2, 3, 4},
			},
			want: true,
		},
		{
			name: "monotone line 3",
			s: &Seqexp{
				expression: "m",
			},
			args: args{
				seq: &Seq{4, 3, 2, 1},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.s

			if got := s.Match(tt.args.seq); got != tt.want {
				t.Errorf("Seqexp.Match() = %v, want %v", got, tt.want)
			}
		})
	}
}
