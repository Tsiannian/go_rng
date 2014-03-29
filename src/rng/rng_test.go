package rng_test

import (
	"rng"
	"fmt"
	"sort"
	"time"
	"testing"
	"strings"
)

func TestGeometryGenerator(t *testing.T) {
	fmt.Println("=====Testing for GeometryGenerator begin=====")
	
	grng := rng.NewGeometricGenerator(time.Now().UnixNano())
	fmt.Println("Geometry(0.2): ")
	hist := map[int64]int{}
	for i := 0; i < 10000; i++ {
		hist[grng.Geometric(0.2)]++
	}
	
	keys := []int64{}
	for k := range hist {
		keys = append(keys, k)
	}
	SortInt64Slice(keys)
	
	for _, key := range keys {
		fmt.Printf("%d:\t%s\n", key, strings.Repeat("*", hist[key] / 100))
	}
	
	fmt.Println("=====Testing for GeometryGenerator end=====")
	fmt.Println()
}

func TestBinomialGenerator(t *testing.T) {
	fmt.Println("=====Testing for BinomialGenerator begin=====")
	bing := rng.NewBinomialGenerator(time.Now().UnixNano())
	
	fmt.Println("Binomial(10^6, 0.02) = ", bing.Binomial(1000000, 0.02))
	var n int64 = 6
	var p float64 = 0.5
	fmt.Printf("X ~ Binomial(%d, %.2f): \n", n, p)
	hist := map[int64]int{}
	for i := 0; i < 10000; i++ {
		hist[bing.Binomial(n, p)]++
	}
	
	keys := []int64{}
	for k := range hist {
		keys = append(keys, k)
	}
	SortInt64Slice(keys)
	
	for _, key := range keys {
		fmt.Printf("%d:\t%s\n", key, strings.Repeat("*", hist[key] / 100))
	}
	
	fmt.Println("=====Testing for BinomialGenerator end=====")
	fmt.Println()
}

func TestBernoulliGenerator(t *testing.T) {
	fmt.Println("=====Testing for BernoulliGenerator begin=====")
	beng := rng.NewBernoulliGenerator(time.Now().UnixNano())
	hist := map[bool]int{}
	for i := 0; i < 10000; i++ {
		hist[beng.Bernoulli_P(0.25)]++
	}
	for k, v := range hist {
		fmt.Printf("%t:\t%s\n", k, strings.Repeat("*", v / 500))
	}
	fmt.Printf("ratio: %.2f\n", float32(hist[true]) / float32(hist[true] + hist[false]))
	fmt.Println("=====Testing for BernoulliGenerator end=====")
	fmt.Println()
}

func TestUniformGenerator(t *testing.T) {
	urng := rng.NewUniformGenerator(time.Now().UnixNano())
	
	fmt.Println("=====Testing for UniformGenerator begin=====")
	fmt.Println("Generating 100 random int64s: ")
	for i := 0; i < 100; i++ {
		fmt.Printf("%d\t", urng.Int64())
		if (i + 1) % 3 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
	fmt.Println()
	
	fmt.Println("Generating 100 random int32s: ")
	for i := 0; i < 100; i++ {
		fmt.Printf("%d\t", urng.Int32())
		if (i + 1) % 5 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
	
	fmt.Println("Generating 100 random float64s: ")
	for i := 0; i < 100; i++ {
		fmt.Printf("%.2f\t", urng.Float64())
		if (i + 1) % 10 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
	
	fmt.Println("Generating 100 random float32s: ")
	for i := 0; i < 100; i++ {
		fmt.Printf("%.2f\t", urng.Float32())
		if (i + 1) % 10 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
	
	fmt.Printf("Random number in [0, 1024): %d and %d\n", 
		urng.Int32n(1024), urng.Int64n(1024))
	fmt.Printf("Random number in [33421, 987584): %d and %d\n", 
		urng.Int32Range(33421, 987584), urng.Int64Range(33421, 987584))
		
	fmt.Printf("Random number in [0.0, 47.339): %.2f and %.2f\n", 
		urng.Float32n(47.339), urng.Float64n(47.339))
	fmt.Printf("Random number in [45.485, 999.458): %.2f and %.2f\n", 
		urng.Float32Range(45.485, 999.458), urng.Float64Range(45.485, 999.458))
		
	fmt.Println("=====Testing for UniformGenerator end=====")
}


func SortInt64Slice(slice []int64) {
	sort.Sort(int64slice(slice))
}

type int64slice []int64

func (slice int64slice) Len() int {
	return len(slice)
}

func (slice int64slice) Less(i, j int) bool {
	return slice[i] < slice[j]
}

func (slice int64slice) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}