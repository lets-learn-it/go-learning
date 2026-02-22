package main

const (
	horiLine = '─'
	vertLine = '│'
	ltCorner = '┌'
	lbCorner = '└'
	rtCorner = '┐'
	rbCorner = '┘'
)

type Style struct {
	bgColor, fgColor int
}

var (
	NormalStyle = Style{fgColor: 34, bgColor: 0}
)

type Pane struct {
	x1, y1 int
	x2, y2 int

	title string
	style Style
}

func NewPane(x1, y1, x2, y2 int, title string, s Style) *Pane {
	return &Pane{
		x1:    x1,
		y1:    y1,
		x2:    x2,
		y2:    y2,
		title: title,
		style: s,
	}
}

func (p *Pane) Draw(s *Screen) {
	// borders
	for i := p.x1; i <= p.x2; i++ {
		s.SetPixel(i, p.y1, horiLine, &p.style)
		s.SetPixel(i, p.y2, horiLine, &p.style)
	}

	for i := p.y1; i <= p.y2; i++ {
		s.SetPixel(p.x1, i, vertLine, &p.style)
		s.SetPixel(p.x2, i, vertLine, &p.style)
	}

	// set corners
	s.SetPixel(p.x1, p.y1, ltCorner, &p.style)
	s.SetPixel(p.x1, p.y2, lbCorner, &p.style)
	s.SetPixel(p.x2, p.y1, rtCorner, &p.style)
	s.SetPixel(p.x2, p.y2, rbCorner, &p.style)

	// title
	titleRunes := []rune(p.title)
	start := p.x1 + 1
	end := p.x1 + len(titleRunes)
	for i := start; i <= end && i < p.x2; i++ {
		s.SetPixel(i, p.y1, titleRunes[i-start], &p.style)
	}
}
