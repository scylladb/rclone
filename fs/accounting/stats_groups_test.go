package accounting

import (
	"testing"
)

func TestStatsGroupOperations(t *testing.T) {

	t.Run("empty group returns nil", func(t *testing.T) {
		t.Parallel()
		sg := newStatsGroups()
		sg.get("invalid-group")
	})

	t.Run("set assigns stats to group", func(t *testing.T) {
		t.Parallel()
		stats := NewStats()
		sg := newStatsGroups()
		sg.set("test", stats)
		sg.set("test1", stats)
		if len(sg.m) != len(sg.names()) || len(sg.m) != 2 {
			t.Fatalf("Expected two stats got %d, %d", len(sg.m), len(sg.order))
		}
	})

	t.Run("get returns correct group", func(t *testing.T) {
		t.Parallel()
		stats := NewStats()
		sg := newStatsGroups()
		sg.set("test", stats)
		sg.set("test1", stats)
		got := sg.get("test")
		if got != stats {
			t.Fatal("get returns incorrect stats")
		}
	})

	t.Run("sum returns correct values", func(t *testing.T) {
		t.Parallel()
		stats1 := NewStats()
		stats1.bytes = 5
		stats1.errors = 5
		stats2 := NewStats()
		sg := newStatsGroups()
		sg.set("test1", stats1)
		sg.set("test2", stats2)
		sum := sg.sum()
		if sum.bytes != stats1.bytes+stats2.bytes {
			t.Fatalf("sum() => bytes %d, expected %d", sum.bytes, stats1.bytes+stats2.bytes)
		}
		if sum.errors != stats1.errors+stats2.errors {
			t.Fatalf("sum() => errors %d, expected %d", sum.errors, stats1.errors+stats2.errors)
		}
	})

	t.Run("delete removes stats", func(t *testing.T) {
		t.Parallel()
		stats := NewStats()
		sg := newStatsGroups()
		sg.set("test", stats)
		sg.set("test1", stats)
		sg.delete("test1")
		if sg.get("test1") != nil {
			t.Fatal("stats not deleted")
		}
		if len(sg.m) != len(sg.names()) || len(sg.m) != 1 {
			t.Fatalf("Expected two stats got %d, %d", len(sg.m), len(sg.order))
		}
	})
}
