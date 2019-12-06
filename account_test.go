package ys7

import "testing"

func TestAccount(t *testing.T) {
	ramAcc, err := ys.CreateAccount("ys7_test", "passwd")
	if err != nil {
		t.Error(err)
	}
	tlog("CreateAccount", ramAcc)

	iAcc, err := ys.RAMAccountGet(ramAcc.AccountID, "")
	if err != nil {
		t.Error(err)
	}
	tlog("RAMAccountGet", iAcc)

	nAcc, err := ys.RAMAccountGet("", "ys7_test")
	if err != nil {
		t.Error(err)
	}
	tlog("RAMAccountGet", nAcc)

	lAcc, page, err := ys.RAMAccountList(0, 10)
	if err != nil {
		t.Error(err)
	}
	tlog("RAMAccountList", lAcc, page)

	err = ys.RAMUpdatePassword(ramAcc.AccountID, "passwd", "new_passwd")
	if err != nil {
		t.Error(err)
	}
	tlog("RAMUpdatePassword success")

	rAc, err := ys.RAMGetAccessToken(ramAcc.AccountID)
	if err != nil {
		t.Error(err)
	}
	tlog(rAc)

	err = ys.DeleteAccount(ramAcc.AccountID)
	if err != nil {
		t.Error(err)
	}
	tlog("DeleteAccount success")

}
