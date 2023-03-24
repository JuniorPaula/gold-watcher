package main

import "testing"

func TestApp_ToolBar(t *testing.T) {
	tb := testApp.getToolBar()

	if len(tb.Items) != 4 {
		t.Error("wrong number of items in toobar")
	}
}
