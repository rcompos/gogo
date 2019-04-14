package main

import (
	"image/color"
	//"math"
	_ "math/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"github.com/gonum/stat/distuv"
)

func main() {
	// Draw some random values from the standard
	// normal distribution.
	//rand.Seed(int64(0))
	//v := make(plotter.Values, 10000)
	//for i := range v {
	//	v[i] = rand.NormFloat64()
	//}
	v := plotter.Values{4, 4, 4, 5, 1, 5, 2, 3, 6, 3, 8, 4, 6, 9, 5, 4, 8}

	// Make a plot and set its title.
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "Histogram"

	// Create a histogram of our values drawn
	// from the standard normal.
	h, err := plotter.NewHist(v, 16)
	if err != nil {
		panic(err)
	}
	// Normalize the area under the histogram to
	// sum to one.
	//h.Normalize(1)
	p.Add(h)

	// The normal distribution function
	norm := plotter.NewFunction(distuv.UnitNormal.Prob)
	norm.Color = color.RGBA{R: 255, A: 255}
	norm.Width = vg.Points(2)
	p.Add(norm)

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "hist.png"); err != nil {
		panic(err)
	}
}
