package infrastructure

import "testing"

func TestGetSeasonAndEpisode(t *testing.T) {

	tests := []struct {
		fileName string
		expected []uint
	}{
		{
			"random title-2011-s1e11.mov",
			[]uint{1, 11},
		},
		{
			"random title-2011-s10e1.wav",
			[]uint{10, 1},
		},
		{
			"random title-2011-s100e100.wav",
			[]uint{100, 100},
		},
		{
			"random title-2011s9e11.mp4",
			[]uint{9, 11},
		},
		{
			"s2e1.wav",
			[]uint{2, 1},
		},
	}

	for i, tt := range tests {
		s, e := getSeasonAndEpisode(tt.fileName)

		if s != tt.expected[0] {
			t.Errorf("case %v) incorrect season=\"%v\", got=\"%v\"", i, tt.expected[0], s)
		}

		if e != tt.expected[1] {
			t.Errorf("case %v) incorrect episode=\"%v\", got=\"%v\"", i, tt.expected[1], e)
		}
	}
}
