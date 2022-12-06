package main

import (
	"testing"
)

func TestFindMarker(t *testing.T) {
	testCases := []struct {
		packet string
		marker int
		offset int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7, 4},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5, 4},
		{"nppdvjthqldpwncqszvftbrmjlhg", 6, 4},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10, 4},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11, 4},
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19, 14},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 23, 14},
		{"nppdvjthqldpwncqszvftbrmjlhg", 23, 14},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29, 14},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26, 14},
	}

	for _, testCase := range testCases {
		marker := find_marker(testCase.packet, testCase.offset)

		if marker != testCase.marker {
			t.Errorf("FAILED: %s offset %d, marker was %d, expected %d", testCase.packet, testCase.offset, marker, testCase.marker)
		} else {
			t.Log("SUCCESS:", testCase.packet, testCase.marker)
		}
	}
}
