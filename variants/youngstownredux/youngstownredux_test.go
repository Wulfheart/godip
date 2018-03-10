package youngstownredux

import (
	"testing"

	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/classical/orders"

	dip "github.com/zond/godip/common"
	cla "github.com/zond/godip/variants/classical/common"
	tst "github.com/zond/godip/variants/testing"
)

func init() {
	dip.Debug = true
}

func startState(t *testing.T) *state.State {
	judge, err := YoungstownReduxStart()
	if err != nil {
		t.Fatalf("%v", err)
	}
	return judge
}

func blankState(t *testing.T) *state.State {
	startPhase := classical.Phase(1901, cla.Spring, cla.Movement)
	judge := YoungstownReduxBlank(startPhase)
	return judge
}

func TestHebei(t *testing.T) {
	judge := startState(t)

	// Test (and document) that there is no connection from Hebei South Coast to Yellow Sea.
	judge.SetUnit("heb/sc", dip.Unit{cla.Fleet, Japan})
	tst.AssertOrderValidity(t, judge, orders.Move("heb/sc", "yel"), "", cla.ErrIllegalMove)

	// Check that this is possible from the North Coast.
	judge.RemoveUnit("heb/sc")
	judge.SetUnit("heb/nc", dip.Unit{cla.Fleet, Japan})
	tst.AssertOrderValidity(t, judge, orders.Move("heb/nc", "yel"), Japan, nil)
}

func TestBoxes(t *testing.T) {
	judge := startState(t)

	// Test some of the connections between boxes.
	judge.SetUnit("bxa", dip.Unit{cla.Fleet, Britain})
	tst.AssertOptionToMove(t, judge, Britain, "bxa", "bxb")
	tst.AssertOptionToMove(t, judge, Britain, "bxa", "bxc")
	tst.AssertOptionToMove(t, judge, Britain, "bxa", "bxd")
	tst.AssertNoOptionToMoveTo(t, judge, Britain, "bxa", "npo")
	tst.AssertNoOptionToMoveTo(t, judge, Britain, "bxa", "bxe")

	judge.SetUnit("bxb", dip.Unit{cla.Fleet, France})
	tst.AssertOptionToMove(t, judge, France, "bxb", "bxa")
	tst.AssertOptionToMove(t, judge, France, "bxb", "bxc")
	tst.AssertOptionToMove(t, judge, France, "bxb", "bxe")
	tst.AssertNoOptionToMoveTo(t, judge, France, "bxb", "bxg")

	judge.SetUnit("bxc", dip.Unit{cla.Fleet, Italy})
	tst.AssertOptionToMove(t, judge, Italy, "bxc", "bxa")
	tst.AssertOptionToMove(t, judge, Italy, "bxc", "bxb")
	tst.AssertOptionToMove(t, judge, Italy, "bxc", "bxf")
	tst.AssertOptionToMove(t, judge, Italy, "bxc", "bxg")
	tst.AssertOptionToMove(t, judge, Italy, "bxc", "bxh")
}