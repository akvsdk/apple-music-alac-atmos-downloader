package fmp4unfrag

import (
	"testing"

	"github.com/Eyevinn/mp4ff/mp4"
)

func TestCollectSampleGroupRunsIgnoresRollGroup(t *testing.T) {
	traf := &mp4.TrafBox{
		Sbgp: &mp4.SbgpBox{GroupingType: "roll"},
	}
	td := &trackData{}

	if err := collectSampleGroupRuns(td, traf, 0); err != nil {
		t.Fatal(err)
	}
	if len(td.SampleGroups) != 0 {
		t.Fatalf("expected roll group to be ignored, got %d runs", len(td.SampleGroups))
	}
}
