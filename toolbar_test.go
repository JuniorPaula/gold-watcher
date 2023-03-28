package main

import (
	"testing"

	"fyne.io/fyne/v2/test"
)

func TestApp_ToolBar(t *testing.T) {
	tb := testApp.getToolBar()

	if len(tb.Items) != 4 {
		t.Error("wrong number of items in toobar")
	}
}

func TestApp_addHoldingsDialogs(t *testing.T) {
	testApp.addHoldingsDialog()

	test.Type(testApp.AddHoldingsPurchseAmountEntry, "1")
	test.Type(testApp.AddHoldingsPurchsePriceEntry, "1000")
	test.Type(testApp.AddHoldingsPurchseDateEntry, "2023-03-28")

	if testApp.AddHoldingsPurchseDateEntry.Text != "2023-03-28" {
		t.Error("date not correct")
	}

	if testApp.AddHoldingsPurchseAmountEntry.Text != "1" {
		t.Error("amount not correct")
	}

	if testApp.AddHoldingsPurchsePriceEntry.Text != "1000" {
		t.Error("price not correct")
	}

}
