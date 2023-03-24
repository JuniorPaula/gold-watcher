package main

import "testing"

func TestApp_getPriceText(t *testing.T) {

	open, _, _ := testApp.GetPriceText()
	if open.Text != "Open: $10242.9064 USD" {
		t.Error("wrong price returned", open.Text)
	}
}
