package main

import (
	g "github.com/AllenDang/giu"
)

var (
	datafont_16 *g.FontInfo
	datafont_24 *g.FontInfo
	datafont_32 *g.FontInfo
	datafont_48 *g.FontInfo
	datafont_64 *g.FontInfo
	datafont_72 *g.FontInfo
	datafont_96 *g.FontInfo
)

func loop() {
	g.SingleWindow().Layout(
		g.Label("0verlay Text 96").Font(datafont_96),
		g.Label("0verlay Text 72").Font(datafont_72),
		g.Label("0verlay Text 64").Font(datafont_64),
		g.Label("0verlay Text 48").Font(datafont_48),
		g.Label("0verlay Text 32").Font(datafont_32),
		g.Label("0verlay Text 24").Font(datafont_24),
		g.Label("0verlay Text 16").Font(datafont_16),
	)
}

func main() {
	wnd := g.NewMasterWindow("Font test", 600, 400, 0)
	datafont_16 = g.Context.FontAtlas.AddFont("robotomono-m.ttf", 16)
	datafont_24 = g.Context.FontAtlas.AddFont("robotomono-m.ttf", 24)
	datafont_32 = g.Context.FontAtlas.AddFont("robotomono-m.ttf", 32)
	datafont_48 = g.Context.FontAtlas.AddFont("robotomono-m.ttf", 48)
	datafont_64 = g.Context.FontAtlas.AddFont("robotomono-m.ttf", 64)
	datafont_72 = g.Context.FontAtlas.AddFont("robotomono-m.ttf", 72)
	datafont_96 = g.Context.FontAtlas.AddFont("robotomono-m.ttf", 96)

	wnd.Run(loop)
}