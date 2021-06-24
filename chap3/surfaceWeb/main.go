// surface は 3-D 面の関数の SVG レンダリングを計算
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"net/url"
	"strconv"
)

const (
	width, height = 600, 320            // キャンパスの大きさ (画素数)
	cells         = 100                 // 格子のマス目の数
	xyrange       = 30.0                // 軸の範囲 (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // x 単位 及び y 単位当たりの画素数
	zscale        = height * 0.4        // z 単位当たりの画素数
	angle         = math.Pi / 6         // x, y 軸の角度 (30度)
	red           = 0xff0000
)

type params struct {
	height int
	width  int
	color  string
}

func (p *params) set(values url.Values) {
	w, _ := strconv.Atoi(values.Get("width"))
	p.width = w
	h, _ := strconv.Atoi(values.Get("height"))
	p.height = h
	c := values.Get("color")
	p.color = c
}

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		p := new(params)
		p.set(r.URL.Query())
		w.Header().Set("Content-Type", "image/svg+xml")
		plot(w, p)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func plot(out io.Writer, p *params) {
	fmt.Fprintf(out, "<svg xmls='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", p.width, p.height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, _ := corner(i+1, j)
			bx, by, _ := corner(i, j)
			cx, cy, _ := corner(i, j+1)
			dx, dy, _ := corner(i+1, j+1)
			if isValid([]float64{ax, ay, bx, by, cx, cy, dx, dy}) {
				fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='#%s'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy, p.color)
			}
		}
	}
	fmt.Fprintln(out, "</svg>")
}

func isValid(vals []float64) bool {
	for _, v := range vals {
		if math.IsInf(v, 0) {
			return false
		}
	}
	return true
}

func color(az, bz, cz, dz float64) string {
	z := (az + bz + cz + dz) / 4
	b := uint32((1.0 - z) / (1.0 - -0.245) * 0xff)
	c := fmt.Sprintf("%X", red-(b<<16)+b)
	for i := len(c); i < 6; i++ {
		c = "0" + c
	}
	return c
}

func corner(i, j int) (float64, float64, float64) {
	// マス目 (i, j) の角の点 (x, y) を見つける
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// 面の高さ z を計算する
	z := f(x, y)

	// (x, y, z) を 2-D SVG キャンバス (sx, sy) へ等角的に投影
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // (0, 0) からの距離
	return math.Sin(r) / r
	//return math.Min(x, y) / r
}
