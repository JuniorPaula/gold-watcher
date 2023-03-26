package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func (app *Config) makeUI() {
	openPrice, currentPrice, priceChange := app.GetPriceText()

	priceContent := container.NewGridWithColumns(3,
		openPrice,
		currentPrice,
		priceChange,
	)

	app.PriceContainer = priceContent

	toolBar := app.getToolBar()
	app.ToolBar = toolBar

	priceTabContent := app.pricesTab()
	holdingsTabContent := app.holdingsTab()

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Prices", theme.HomeIcon(), priceTabContent),
		container.NewTabItemWithIcon("Holdings", theme.InfoIcon(), holdingsTabContent),
	)
	tabs.SetTabLocation(container.TabLocationTop)

	finalContent := container.NewVBox(priceContent, toolBar, tabs)
	app.MainWindow.SetContent(finalContent)

	go func() {
		for range time.Tick(time.Second * 30) {
			app.refreshPriceContent()
		}
	}()
}

func (app *Config) refreshPriceContent() {
	app.InfoLog.Print("refreshing price")
	open, current, change := app.GetPriceText()
	app.PriceContainer.Objects = []fyne.CanvasObject{open, current, change}
	app.PriceContainer.Refresh()

	chart := app.getChart()
	app.PriceContainer.Objects = []fyne.CanvasObject{chart}
	app.PriceContainer.Refresh()
}

func (app *Config) refreshHoldingsTable() {
	app.Holdings = app.getHoldingsSlice()
	app.HoldingsTable.Refresh()
}
