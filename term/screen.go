package term

import (
	"github.com/nsf/termbox-go"
)

type Screen struct {
	Width, Height                       int
	DefaultForground, DefaultBackground termbox.Attribute
	Lines                               []Line
	Header                              Line
	Footer                              Line
}

func Init() error {
	return termbox.Init()
}

func Event() termbox.Event {
	return termbox.PollEvent()
}

func NewScreen(fg, bg termbox.Attribute) (*Screen, error) {
	s := &Screen{Lines: []Line{}}
	s.Width, s.Height = termbox.Size()
	s.DefaultForground, s.DefaultBackground = fg, bg

	return s, nil
}

func (s *Screen) Close() error {
	termbox.Close()
	return nil
}

func (s *Screen) Clear() error {
	return termbox.Clear(s.DefaultForground, s.DefaultBackground)
}

func (s *Screen) Display() error {
	if err := s.Clear(); err != nil {
		return nil
	}
	x := 0
	if s.Header != nil {
		if err := s.Header.Display(x, 0, s); err != nil {
			return err
		}
		x++
	}
	for _, l := range s.Lines {
		if err := l.Display(x, 0, s); err != nil {
			return err
		}
		x++
	}
	if s.Footer != nil {
		if err := s.Footer.Display(s.Height-1, 0, s); err != nil {
			return err
		}
	}
	return termbox.Flush()
}

// Resize and display the screen
func (s *Screen) Resize() error {
	s.Width, s.Height = termbox.Size()
	return s.Display()
}