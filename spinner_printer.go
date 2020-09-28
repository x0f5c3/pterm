package pterm

import "time"

// DefaultSpinner is the default spinner
var DefaultSpinner = Spinner{
	Sequence:       []string{"▀ ", " ▀", " ▄", "▄ "},
	Style:          Style{FgLightCyan},
	Delay:          time.Millisecond * 200,
	MessageStyle:   Style{FgLightWhite},
	SuccessPrinter: Success,
	FailPrinter:    Error,
	WarningPrinter: Warning,
}

// Spinner is a loading animation, which can be used if the progress is unknown.
// It's an animation loop, which can have a text and supports throwing errors or warnings.
// A GenericPrinter is used to display all outputs, after the spinner is done.
type Spinner struct {
	Text           string
	Sequence       []string
	Style          Style
	Delay          time.Duration
	MessageStyle   Style
	SuccessPrinter GenericPrinter
	FailPrinter    GenericPrinter
	WarningPrinter GenericPrinter

	IsActive bool
}

// WithText adds a text to the spinner
func (s Spinner) WithText(text string) *Spinner {
	s.Text = text
	return &s
}

// WithSequence adds a sequence to the spinner
func (s Spinner) WithSequence(sequence ...string) *Spinner {
	s.Sequence = sequence
	return &s
}

// WithStyle adds a style to the spinner
func (s Spinner) WithStyle(colors ...Color) *Spinner {
	s.Style = colors
	return &s
}

// WithDelay adds a delay to the spinner
func (s Spinner) WithDelay(delay time.Duration) *Spinner {
	s.Delay = delay
	return &s
}

// WithMessageStyle adds a style to the spinner message
func (s Spinner) WithMessageStyle(colors ...Color) *Spinner {
	s.MessageStyle = colors
	return &s
}

// Start starts the spinner
func (s *Spinner) Start(text ...interface{}) *Spinner {
	s.IsActive = true

	if len(text) != 0 {
		s.Text = Sprint(text...)
	}

	go func(s *Spinner) {
		for s.IsActive {
			for _, seq := range s.Sequence {
				if s.IsActive {
					Printo(s.Style.Sprint(seq) + " " + s.MessageStyle.Sprint(s.Text))
					time.Sleep(s.Delay)
				}
			}
		}
	}(s)
	return s
}

// Success displays the success printer.
// If no message is given, the text of the spinner will be reused as the default message.
func (s *Spinner) Success(message ...interface{}) {
	if len(message) == 0 {
		message = []interface{}{s.Text}
	}
	clearLine()
	Printo(s.SuccessPrinter.Sprint(message...))
	s.Stop()
}

// Fail displays the fail printer.
// If no message is given, the text of the spinner will be reused as the default message.
func (s *Spinner) Fail(message ...interface{}) {
	if len(message) == 0 {
		message = []interface{}{s.Text}
	}
	clearLine()
	Printo(s.FailPrinter.Sprint(message...))
	s.Stop()
}

// Warning displays the warning printer.
// If no message is given, the text of the spinner will be reused as the default message.
func (s *Spinner) Warning(message ...interface{}) {
	if len(message) == 0 {
		message = []interface{}{s.Text}
	}
	clearLine()
	Printo(s.WarningPrinter.Sprint(message...))
	s.Stop()
}

// Stop terminates the Spinner immediately.
// The Spinner will not resolve into anything.
func (s *Spinner) Stop() {
	s.IsActive = false
	Println()
}
