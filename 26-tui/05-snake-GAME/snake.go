package main

type snake struct {
	body      []position
	direction direction
	w         int
	h         int
}

func newSnake(w, h int) *snake {
	s := &snake{
		w: w,
		h: h,
	}

	s.body = append(s.body, position{x: 5, y: 5})
	s.direction = east

	return s
}

// moves head & tries to eat food, if success return true
func (s *snake) moveAndEat(d direction, food position) bool {
	head := s.body[0]

	switch d {
	case north:
		head.y -= 1
	case east:
		head.x += 1
	case south:
		head.y += 1
	case west:
		head.x -= 1
	}

	if food.x == head.x && food.y == head.y {
		s.body = append([]position{head}, s.body...)
		return true
	} else {
		s.body = append([]position{head}, s.body[:len(s.body)-1]...)
		return false
	}
}

func (s *snake) teleport(w, h int) {
	for i := 0; i < len(s.body); i++ {
		s.body[i] = teleport(s.w, s.h, w, h, s.body[i])
	}
	s.w = w
	s.h = h
}

func (s *snake) hitWall(w, h int) bool {
	return s.body[0].x <= 1 || s.body[0].y <= 1 || s.body[0].x >= w || s.body[0].y >= h
}

func (s *snake) ateItself() bool {
	n := len(s.body)
	head := s.body[0]

	for i := 1; i < n; i++ {
		if head.x == s.body[i].x && head.y == s.body[i].y {
			return true
		}
	}

	return false
}
