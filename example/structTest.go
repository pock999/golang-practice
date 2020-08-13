package main

import (
    "fmt"
    "math"    
);

type Point struct {
    x float64;
    y float64;
}

func main() {
	p := Point{x: 3.0, y: 4.0};
    q := Point{x: 0.0, y: 0.0};
 
    distance := getDistance(p, q);
    fmt.Println(distance);
}

func getDistance(a Point, b Point) float64 {
    return math.Sqrt(math.Pow(a.x-b.x, 2) + math.Pow(a.y-b.y, 2));
}
