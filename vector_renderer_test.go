package chart

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/blendlabs/go-assert"
	"github.com/wcharczuk/go-chart/drawing"
)

func TestVectorRendererPath(t *testing.T) {
	assert := assert.New(t)

	vr, err := SVG(100, 100)
	assert.Nil(err)

	typed, isTyped := vr.(*vectorRenderer)
	assert.True(isTyped)

	typed.MoveTo(0, 0)
	typed.LineTo(100, 100)
	typed.LineTo(0, 100)
	typed.Close()
	typed.FillStroke()

	buffer := bytes.NewBuffer([]byte{})
	err = typed.Save(buffer)
	assert.Nil(err)

	raw := string(buffer.Bytes())

	assert.True(strings.HasPrefix(raw, "<svg"))
	assert.True(strings.HasSuffix(raw, "</svg>"))
}

func TestVectorRendererMeasureText(t *testing.T) {
	assert := assert.New(t)

	f, err := GetDefaultFont()
	assert.Nil(err)

	vr, err := SVG(100, 100)
	assert.Nil(err)

	vr.SetDPI(DefaultDPI)
	vr.SetFont(f)
	vr.SetFontSize(12.0)

	tb := vr.MeasureText("Ljp")
	assert.Equal(21, tb.Width())
	assert.Equal(15, tb.Height())
}

func TestCanvasStyleSVG(t *testing.T) {
	assert := assert.New(t)

	f, err := GetDefaultFont()
	assert.Nil(err)

	set := Style{
		StrokeColor: drawing.ColorWhite,
		StrokeWidth: 5.0,
		FillColor:   drawing.ColorWhite,
		FontColor:   drawing.ColorWhite,
		Font:        f,
		Padding:     DefaultBackgroundPadding,
	}

	canvas := &canvas{dpi: DefaultDPI}

	svgString := canvas.styleAsSVG(set)
	assert.NotEmpty(svgString)
	assert.True(strings.Contains(svgString, "stroke:rgba(255,255,255,1.0)"))
	assert.True(strings.Contains(svgString, "stroke-width:5"))
	assert.True(strings.Contains(svgString, "fill:rgba(255,255,255,1.0)"))
}

func TestEmptySVG(t *testing.T) {
	//	assert := assert.New(t)

	graph := Chart{
		Series: []Series{
			ContinuousSeries{
				XValues: []float64{1.0, 2.0},
				YValues: []float64{0.0, 0.0},
			},
		},
		XAxis: XAxis{
			Style: StyleShow(),
		},
		YAxis: YAxis{
			Style: StyleShow(),
			Range: &ContinuousRange{
				Min: 0.0,
				Max: 10.0,
			},
		},
	}

	buffer := bytes.NewBuffer([]byte{})
	graph.Render(SVG, buffer)
	fmt.Println(string(buffer.Bytes()))
}
