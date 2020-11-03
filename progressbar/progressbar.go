package progressbar

import (
	"fmt"
	"strconv"
)

// ProgressBar ...
type ProgressBar struct {
	TotalPorts    int
	Current       int
	Percentage    float64
	LastDisplayed string
	
}

func (p *ProgressBar) setPercentage(i int) {
	
	perc := float64(i) / float64((*p).TotalPorts)
	// log.Info("Unrounded percentage: %v", perc)
	(*p).Percentage = perc
}

// PercentageHelper ...
func PercentageHelper(pb *ProgressBar, n int) {
	pb.setPercentage(n)
	
	currentPercentage := strconv.FormatInt(int64((*pb).Percentage*100), 10)
	if currentPercentage != (*pb).LastDisplayed {
		(*pb).LastDisplayed = strconv.FormatInt(int64((*pb).Percentage*100), 10)
		fmt.Print("#")
		// fmt.Print("\r" + (*pb).LastDisplayed + "%")
	}

}