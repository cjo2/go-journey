package _switch

func GetEvenNumbers(numbers []int) ([]int, int) {
	oddCounter := 0
	var evens []int
	for _, num := range numbers {
		remainder := num % 2
		switch remainder {
		case 0:
			evens = append(evens, num)
		case 1:
			oddCounter = oddCounter + 1
		}
	}
	return evens, oddCounter
}

type Plane struct {
	XPosition int
	YPosition int
}

func (p *Plane) Navigate(direction string) {
	switch direction {
	case "up":
		p.XPosition = p.XPosition + 1
		break
	case "down":
		p.XPosition = p.XPosition - 1
	case "left":
		p.YPosition = p.YPosition - 1
	case "right":
		p.YPosition = p.YPosition + 1
		fallthrough
	case "runsbecauseoffallthrough":
		p.XPosition = p.XPosition + 100
	}
}
