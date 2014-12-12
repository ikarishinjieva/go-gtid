package gtid

import (
	"testing"
)

func TestGtidAdd(t *testing.T) {
	gtid, err := GtidAdd(
		"ca8035ea-c5d5-11e3-8ce9-e66ccf50db66:5-10,ff92c4da-c5d7-11e3-8cf7-5e10e6a05cfb:7",
		"ca8035ea-c5d5-11e3-8ce9-e66ccf50db66:1-4,ff92c4da-c5d7-11e3-8cf7-5e10e6a05cfb:1-5")
	if nil != err {
		t.Fatalf("unexpected error %v", err)
	}
	if gtid != "CA8035EAC5D511E38CE9E66CCF50DB66:1-10,FF92C4DAC5D711E38CF75E10E6A05CFB:1-5:7" {
		t.Fatalf("wrong gtid %v", gtid)
	}
}

func TestGtidContain_pass(t *testing.T) {
	contain, err := GtidContain(
		"ff92c4da-c5d7-11e3-8cf7-5e10e6a05cfb:1-7,ca8035ea-c5d5-11e3-8ce9-e66ccf50db66:1-10",
		"ca8035ea-c5d5-11e3-8ce9-e66ccf50db66:1-4,ff92c4da-c5d7-11e3-8cf7-5e10e6a05cfb:1-5")
	if nil != err {
		t.Fatalf("unexpected error %v", err)
	}
	if !contain {
		t.Fatalf("contain should == true")
	}
}

func TestGtidContain_not_pass(t *testing.T) {
	contain, err := GtidContain(
		"ff92c4da-c5d7-11e3-8cf7-5e10e6a05cfb:1-7,ca8035ea-c5d5-11e3-8ce9-e66ccf50db66:2-10",
		"ca8035ea-c5d5-11e3-8ce9-e66ccf50db66:1-4,ff92c4da-c5d7-11e3-8cf7-5e10e6a05cfb:1-5")
	if nil != err {
		t.Fatalf("unexpected error %v", err)
	}
	if contain {
		t.Fatalf("contain should == false")
	}
}

func TestGtidContain_null(t *testing.T) {
	contain, err := GtidContain(
		"",
		"")
	if nil != err {
		t.Fatalf("unexpected error %v", err)
	}
	if !contain {
		t.Fatalf("contain should == true")
	}
}

func TestGtidSub_1(t *testing.T) {
	gtid, err := GtidSub(
		"ca8035ea-c5d5-11e3-8ce9-e66ccf50db66:1-10:20-30",
		"ca8035ea-c5d5-11e3-8ce9-e66ccf50db66:11-12")
	if nil != err {
		t.Fatalf("unexpected error %v", err)
	}
	if gtid != "CA8035EAC5D511E38CE9E66CCF50DB66:1-10:20-30" {
		t.Fatalf("wrong gtid %v", gtid)
	}
}

func TestGtidSub_2(t *testing.T) {
	gtid, err := GtidSub(
		"ca8035ea-c5d5-11e3-8ce9-e66ccf50db66:1-10:20-30",
		"ca8035ea-c5d5-11e3-8ce9-e66ccf50db66:15-40")
	if nil != err {
		t.Fatalf("unexpected error %v", err)
	}
	if gtid != "CA8035EAC5D511E38CE9E66CCF50DB66:1-10" {
		t.Fatalf("wrong gtid %v", gtid)
	}
}

func TestGtidSub_3(t *testing.T) {
	gtid, err := GtidSub(
		"ca8035ea-c5d5-11e3-8ce9-e66ccf50db66:1-10:20-30",
		"ca8035ea-c5d5-11e3-8ce9-e66ccf50db66:15-25")
	if nil != err {
		t.Fatalf("unexpected error %v", err)
	}
	if gtid != "CA8035EAC5D511E38CE9E66CCF50DB66:1-10:26-30" {
		t.Fatalf("wrong gtid %v", gtid)
	}
}

func TestGtidSub_4(t *testing.T) {
	gtid, err := GtidSub(
		"ca8035ea-c5d5-11e3-8ce9-e66ccf50db66:1-10:20-30",
		"ca8035ea-c5d5-11e3-8ce9-e66ccf50db66:15-30")
	if nil != err {
		t.Fatalf("unexpected error %v", err)
	}
	if gtid != "CA8035EAC5D511E38CE9E66CCF50DB66:1-10" {
		t.Fatalf("wrong gtid %v", gtid)
	}
}

func TestGtidSub_5(t *testing.T) {
	gtid, err := GtidSub(
		"ca8035ea-c5d5-11e3-8ce9-e66ccf50db66:1-10:20-30",
		"ca8035ea-c5d5-11e3-8ce9-e66ccf50db66:2-29")
	if nil != err {
		t.Fatalf("unexpected error %v", err)
	}
	if gtid != "CA8035EAC5D511E38CE9E66CCF50DB66:1:30" {
		t.Fatalf("wrong gtid %v", gtid)
	}
}

func TestGtidSub_6(t *testing.T) {
	gtid, err := GtidSub(
		"ca8035ea-c5d5-11e3-8ce9-e66ccf50db66:1-10:20-30",
		"ca8035ea-c5d5-11e3-8ce9-e66ccf50db66:1-30")
	if nil != err {
		t.Fatalf("unexpected error %v", err)
	}
	if gtid != "" {
		t.Fatalf("wrong gtid %v", gtid)
	}
}
