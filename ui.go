package main

import "fyne.io/fyne/v2/container"

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

	finalContent := container.NewVBox(priceContent, toolBar)
	app.MainWindow.SetContent(finalContent)
}
