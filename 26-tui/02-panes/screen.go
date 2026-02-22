package main

type Screen struct {
	w, h        int
	front, back []rune
	styles      []*Style

	dirtyCols []bool
}

func NewScreen(w, h int) *Screen {
	return &Screen{
		w:         w,
		h:         h,
		front:     make([]rune, w*h),
		back:      make([]rune, w*h),
		styles:    make([]*Style, w*h),
		dirtyCols: make([]bool, w),
	}
}

func (s *Screen) Draw() {
	for i := 0; i < s.w; i++ {
		if !s.dirtyCols[i] {
			continue
		}

		for j := 0; j < s.h; j++ {
			if s.front[i*s.h+j] != s.back[i*s.h+j] {
				if s.styles[i*s.h+j] == nil {
					SetAt(i, j, string(s.back[i*s.h+j]))
				} else {
					style := s.styles[i*s.h+j]
					SetAtWithStyle(i, j, string(s.back[i*s.h+j]), style.fgColor, style.bgColor)
				}

				s.front[i*s.h+j] = s.back[i*s.h+j]
			}
		}

		s.dirtyCols[i] = false
	}
}

func (s *Screen) SetPixel(x, y int, c rune, style *Style) {
	if x < 0 || x >= s.w || y < 0 || y >= s.h {
		return
	}

	s.dirtyCols[x] = true
	s.back[x*s.h+y] = c
	s.styles[x*s.h+y] = style
}
