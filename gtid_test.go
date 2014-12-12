package gtid

import (
	"testing"
)

func TestParseGtid_simple(t *testing.T) {
	gtid, err := parseGtid("ca8035ea-c5d5-11e3-8ce9-e66ccf50db66:1-11:13, ff92c4da-c5d7-11e3-8cf7-5e10e6a05cfb:1-5")
	if nil != err {
		t.Fatalf("unexpected error %v", err)
	}
	if gtid.String() != "CA8035EAC5D511E38CE9E66CCF50DB66:1-11:13,FF92C4DAC5D711E38CF75E10E6A05CFB:1-5" {
		t.Fatalf("wrong gtid %v", gtid.String())
	}
}

func TestParseGtid_interval_merge_left(t *testing.T) {
	gtid, err := parseGtid("ca8035ea-c5d5-11e3-8ce9-e66ccf50db66:1-10:13-14:15-20")
	if nil != err {
		t.Fatalf("unexpected error %v", err)
	}
	if gtid.String() != "CA8035EAC5D511E38CE9E66CCF50DB66:1-10:13-20" {
		t.Fatalf("wrong gtid %v", gtid.String())
	}
}

func TestParseGtid_interval_merge_right(t *testing.T) {
	gtid, err := parseGtid("ca8035ea-c5d5-11e3-8ce9-e66ccf50db66:1-10:11-12:15-20")
	if nil != err {
		t.Fatalf("unexpected error %v", err)
	}
	if gtid.String() != "CA8035EAC5D511E38CE9E66CCF50DB66:1-12:15-20" {
		t.Fatalf("wrong gtid %v", gtid.String())
	}
}

func TestParseGtid_interval_merge_cross(t *testing.T) {
	gtid, err := parseGtid("ca8035ea-c5d5-11e3-8ce9-e66ccf50db66:5-10:1-12:15-20")
	if nil != err {
		t.Fatalf("unexpected error %v", err)
	}
	if gtid.String() != "CA8035EAC5D511E38CE9E66CCF50DB66:1-12:15-20" {
		t.Fatalf("wrong gtid %v", gtid.String())
	}
}

func TestParseGtid_uuid_merge_left(t *testing.T) {
	gtid, err := parseGtid("ca8035ea-c5d5-11e3-8ce9-e66ccf50db66:5-10, ff92c4da-c5d7-11e3-8cf7-5e10e6a05cfb:1-5, ca8035ea-c5d5-11e3-8ce9-e66ccf50db66:1-4")
	if nil != err {
		t.Fatalf("unexpected error %v", err)
	}
	if gtid.String() != "CA8035EAC5D511E38CE9E66CCF50DB66:1-10,FF92C4DAC5D711E38CF75E10E6A05CFB:1-5" {
		t.Fatalf("wrong gtid %v", gtid.String())
	}
}

func TestParseGtid_uuid_merge_right(t *testing.T) {
	gtid, err := parseGtid("ca8035ea-c5d5-11e3-8ce9-e66ccf50db66:5-10, ff92c4da-c5d7-11e3-8cf7-5e10e6a05cfb:1-5, ca8035ea-c5d5-11e3-8ce9-e66ccf50db66:1-4")
	if nil != err {
		t.Fatalf("unexpected error %v", err)
	}
	if gtid.String() != "CA8035EAC5D511E38CE9E66CCF50DB66:1-10,FF92C4DAC5D711E38CF75E10E6A05CFB:1-5" {
		t.Fatalf("wrong gtid %v", gtid.String())
	}
}
