go build -o mandelbrot mandelbrot_parallel.go
go build -o convert main.go

./mandelbrot > mandelbrot.png
./convert -format jpeg < mandelbrot.png > mandelbrot.jpeg
./convert -format gif < mandelbrot.png > mandelbrot.gif