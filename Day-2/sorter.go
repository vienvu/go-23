package main

import (
	"github.com/vienvu/go-23/day-2/sorting"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type intArrayFlag []int

func (i *intArrayFlag) String() string {
	return fmt.Sprintf("%v", *i)
}

func (i *intArrayFlag) Set(value string) error {
	intVal, err := strconv.Atoi(value)
	if err != nil {
		return fmt.Errorf("Invalid integer value: %s", value)
	}
	*i = append(*i, intVal)
	return nil
}

type floatArrayFlag []float64

func (f *floatArrayFlag) String() string {
	return fmt.Sprintf("%v", *f)
}

func (f *floatArrayFlag) Set(value string) error {
	floatVal, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return fmt.Errorf("Invalid float value: %s", value)
	}
	*f = append(*f, floatVal)
	return nil
}

func main() {
	var ints intArrayFlag
	var floats floatArrayFlag

	flag.Var(&ints, "int", "Specify an array of integers")
	flag.Var(&floats, "float", "Specify an array of floats")
	flag.Parse()

	for _, arg := range flag.Args() {
		intVal, err := strconv.Atoi(arg)
		if err != nil {
			floatVal, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Printf("Invalid value: %s\n", arg)
				return
			}
			floats = append(floats, floatVal)
		} else {
			ints = append(ints, intVal)
		}
	}

	sort.Ints(ints)
	sort.Float64s(floats)

	sortedInts := make([]string, len(ints))
	for i, val := range ints {
		sortedInts[i] = strconv.Itoa(val)
	}

	sortedFloats := make([]string, len(floats))
	for i, val := range floats {
		sortedFloats[i] = strconv.FormatFloat(val, 'f', -1, 64)
	}

	fmt.Println(strings.Join(sortedInts, " "))
	fmt.Println(strings.Join(sortedFloats, " "))
}
