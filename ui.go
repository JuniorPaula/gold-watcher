package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
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

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Prices", theme.HomeIcon(), priceTabContent),
		container.NewTabItemWithIcon("Holdings", theme.InfoIcon(), canvas.NewText("Holdings content goes here", nil)),
	)
	tabs.SetTabLocation(container.TabLocationTop)

	finalContent := container.NewVBox(priceContent, toolBar, tabs)
	app.MainWindow.SetContent(finalContent)
}

func (app *Config) refreshPriceContent() {
	open, current, change := app.GetPriceText()
	app.PriceContainer.Objects = []fyne.CanvasObject{open, current, change}
	app.PriceContainer.Refresh()

	chart := app.getChart()
	app.PriceContainer.Objects = []fyne.CanvasObject{chart}
	app.PriceContainer.Refresh()
}
