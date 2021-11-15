package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEasyCopy(t *testing.T) {
	tests := []struct {
		msg      string
		offset   int64
		limit    int64
		fromPath string
		toPath   string
	}{
		{
			msg:      "Easy coping from source to dest",
			offset:   0,
			limit:    0,
			fromPath: ".//testdata//input.txt",
			toPath:   "c://tmp//out_offset0_limit0.txt",
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.msg, func(t *testing.T) {
			fromPath := tc.fromPath
			toPath := tc.toPath
			if _, err := os.Stat(toPath); err == nil {
				os.Remove(toPath)
			}
			offset, limit := tc.offset, tc.limit
			if err := Copy(fromPath, toPath, offset, limit); err != nil {
				t.Fatal(err)
			}
			fromInfo, _ := os.Stat(fromPath)
			toInfo, _ := os.Stat(toPath)
			if fromInfo.Size() != toInfo.Size() {
				t.Fatal("Files hase different size ")
			}
		})
	}
}

func TestCopyOffset0LimitNot(t *testing.T) {
	tests := []struct {
		msg      string
		offset   int64
		limit    int64
		fromPath string
		toPath   string
	}{
		{
			msg:      "coping from source to dest - offset 0 , limit 10",
			offset:   0,
			limit:    10,
			fromPath: ".//testdata//input.txt",
			toPath:   "c://tmp//out_offset0_limit10.txt",
		},
		{
			msg:      "coping from source to dest - offset 0 , limit 1000",
			offset:   0,
			limit:    1000,
			fromPath: ".//testdata//input.txt",
			toPath:   "c://tmp//out_offset0_limit1000.txt",
		},
		{
			msg:      "coping from source to dest - offset 0 , limit 10000",
			offset:   0,
			limit:    10000,
			fromPath: ".//testdata//input.txt",
			toPath:   "c://tmp//out_offset0_limit10000.txt",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.msg, func(t *testing.T) {
			fromPath := tc.fromPath
			toPath := tc.toPath
			if _, err := os.Stat(toPath); err == nil {
				os.Remove(toPath)
			}
			offset, limit := tc.offset, tc.limit
			if err := Copy(fromPath, toPath, offset, limit); err != nil {
				t.Fatal(err)
			}
			fromInfo, _ := os.Stat(fromPath)
			toInfo, _ := os.Stat(toPath)
			if fromInfo.Size() != toInfo.Size() && limit == 0 {
				t.Fatal("Files hase different size ")
			}
			if toInfo.Size() >= limit {
				if toInfo.Size() != limit {
					t.Fatalf("Dest files not equal to limit %d != %d \n", toInfo.Size(), limit)
				}
			}
			if limit > toInfo.Size() {
				if toInfo.Size() != fromInfo.Size() {
					msg := "Limit bigger then source file  %d != %d \n and source and dest should be equal by size"
					t.Fatalf(msg, toInfo.Size(), fromInfo.Size())
				}
			}
		})
	}
}

func TestCopyOffsetLimit(t *testing.T) {
	tests := []struct {
		msg      string
		offset   int64
		limit    int64
		fromPath string
		toPath   string
	}{
		{
			msg:      "coping from source to dest - offset 100 , limit 1000",
			offset:   100,
			limit:    1000,
			fromPath: ".//testdata//input.txt",
			toPath:   "c://tmp//out_offset100_limit1000.txt",
		},
		{
			msg:      "coping from source to dest - offset 6000 , limit 1000",
			offset:   6000,
			limit:    1000,
			fromPath: ".//testdata//input.txt",
			toPath:   "c://tmp//out_offset6000_limit1000.txt",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.msg, func(t *testing.T) {
			fromPath := tc.fromPath
			toPath := tc.toPath
			if _, err := os.Stat(toPath); err == nil {
				os.Remove(toPath)
			}
			offset, limit := tc.offset, tc.limit
			if err := Copy(fromPath, toPath, offset, limit); err != nil {
				t.Fatal(err)
			}
			fromInfo, _ := os.Stat(fromPath)
			toInfo, _ := os.Stat(toPath)
			if fromInfo.Size() < (limit + offset) {
				if toInfo.Size() != (fromInfo.Size() - offset) {
					t.Fatalf("Files not copied complitly wait %d got %d \n -- dest file size smaller then limit + offset ",
						fromInfo.Size()-offset, toInfo.Size())
				}
			}
			if fromInfo.Size() > (limit + offset) {
				if toInfo.Size() != limit {
					t.Fatalf("Files not copied complitly wait %d got %d \n -- dest file size bigger then limit + offset ",
						limit, toInfo.Size())
				}
			}
		})
	}
}

func TestCopyErrs(t *testing.T) {
	tests := []struct {
		msg      string
		offset   int64
		limit    int64
		fromPath string
		toPath   string
		errmt    string
	}{
		{
			msg:      "Too big offset ",
			offset:   7000,
			limit:    100,
			fromPath: ".//testdata//input.txt",
			toPath:   "c://tmp//out_offset7000_limit100.txt",
			errmt:    "Too big offset wait %d got %d \n",
		},
	}
	for _, tc := range tests {
		_ = tc
		t.Run(tc.msg, func(t *testing.T) {
			fromPath := tc.fromPath
			toPath := tc.toPath
			if _, err := os.Stat(toPath); err == nil {
				os.Remove(toPath)
			}
			fromInfo, _ := os.Stat(fromPath)
			offset, limit := tc.offset, tc.limit
			err := Copy(fromPath, toPath, offset, limit)
			ms := fmt.Sprintf(tc.errmt, fromInfo.Size(), limit)
			require.Error(t, err, ms)
		})
	}
}
