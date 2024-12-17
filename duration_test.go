// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the [Golang repository](https://github.com/golang/go/blob/master/LICENSE).

package duration

import (
	"testing"
	"time"
)

var parseDurationTests = []struct {
	in   string
	want time.Duration
}{
	// simple
	{"0", 0},
	{"5s", 5 * time.Second},
	{"30s", 30 * time.Second},
	{"1478s", 1478 * time.Second},
	{"10d", 10 * 24 * time.Hour},
	{"2w", 14 * 24 * time.Hour},
	{"1mo", 30 * 24 * time.Hour},
	{"2y", 2 * 365 * 24 * time.Hour},
	// sign
	{"-5s", -5 * time.Second},
	{"+5s", 5 * time.Second},
	{"-0", 0},
	{"+0", 0},
	// decimal
	{"5.0s", 5 * time.Second},
	{"5.6s", 5*time.Second + 600*time.Millisecond},
	{"5.s", 5 * time.Second},
	{".5s", 500 * time.Millisecond},
	{"1.0s", 1 * time.Second},
	{"1.00s", 1 * time.Second},
	{"1.004s", 1*time.Second + 4*time.Millisecond},
	{"1.0040s", 1*time.Second + 4*time.Millisecond},
	{"100.00100s", 100*time.Second + 1*time.Millisecond},
	// different units
	{"10ns", 10 * time.Nanosecond},
	{"11us", 11 * time.Microsecond},
	{"12µs", 12 * time.Microsecond}, // U+00B5
	{"12μs", 12 * time.Microsecond}, // U+03BC
	{"13ms", 13 * time.Millisecond},
	{"14s", 14 * time.Second},
	{"15m", 15 * time.Minute},
	{"16h", 16 * time.Hour},
	// composite durations
	{"3h30m", 3*time.Hour + 30*time.Minute},
	{"10.5s4m", 4*time.Minute + 10*time.Second + 500*time.Millisecond},
	{"-2m3.4s", -(2*time.Minute + 3*time.Second + 400*time.Millisecond)},
	{"1h2m3s4ms5us6ns", 1*time.Hour + 2*time.Minute + 3*time.Second + 4*time.Millisecond + 5*time.Microsecond + 6*time.Nanosecond},
	{"39h9m14.425s", 39*time.Hour + 9*time.Minute + 14*time.Second + 425*time.Millisecond},
	{"2w10d20h10m15s", (2*Week + 10*Day + 20*time.Hour + 10*time.Minute + 15*time.Second)},
	// large value
	{"52763797000ns", 52763797000 * time.Nanosecond},
	// more than 9 digits after decimal point, see https://golang.org/issue/6617
	{"0.3333333333333333333h", 20 * time.Minute},
	// 9007199254740993 = 1<<53+1 cannot be stored precisely in a float64
	{"9007199254740993ns", (1<<53 + 1) * time.Nanosecond},
	// largest duration that can be represented by int64 in nanoseconds
	{"9223372036854775807ns", (1<<63 - 1) * time.Nanosecond},
	{"9223372036854775.807us", (1<<63 - 1) * time.Nanosecond},
	{"9223372036s854ms775us807ns", (1<<63 - 1) * time.Nanosecond},
	// large negative value
	{"-9223372036854775807ns", -1<<63 + 1*time.Nanosecond},
	// huge string; issue 15011.
	{"0.100000000000000000000h", 6 * time.Minute},
	// This value tests the first overflow check in leadingFraction.
	{"0.830103483285477580700h", 49*time.Minute + 48*time.Second + 372539827*time.Nanosecond},
	{"200y20mo", (200*Year + 20*Month)},
}

func TestParseDuration(t *testing.T) {
	for _, tc := range parseDurationTests {
		d, err := Parse(tc.in)
		if err != nil || d != tc.want {
			t.Errorf("ParseDuration(%q) = %v, %v, want %v, nil", tc.in, d, err, tc.want)
		}
	}
}

func TestParseInvalidDuration(t *testing.T) {
	_, err := Parse("30x")
	if err == nil {
		t.Fatal("expected an error")
	}
	expect := `time: unknown unit "x" in duration "30x". Valid units are "ns", "us", "µs", "μs", "ms", "s", "m", "h", "d", "w", "mo", "y"`
	if err.Error() != expect {
		t.Errorf("ParseDuration(30x): %v, expected %v", err.Error(), expect)
	}
}

func TestKeysLen(t *testing.T) {
	k := len(ValidUnits())
	m := len(unitMap)
	if k != m {
		t.Fatalf("unitKeys: %v, unitMap: %v, both should have the same lenght", k, m)
	}
}
