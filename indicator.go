package teamhealthcheck

// Indicator defines red, yellow and green as health metric
type Indicator struct {
	red    int
	yellow int
	green  int
}

// Indicator is Getter for the Indicator
func (i *Indicator) Indicator() Indicator {
	return *i
}

// SetRed is Setter for Red
func (i *Indicator) SetRed(red int) {
	i.red = red
}

// SetYellow is Setter for Yellow
func (i *Indicator) SetYellow(yellow int) {
	i.yellow = yellow
}

// SetGreen is Setter for Green
func (i *Indicator) SetGreen(green int) {
	i.green = green
}

// Red is Getter for the Indicators indicator red
func (i *Indicator) Red() int {
	return i.red
}

// Yellow is Getter for the Indicators indicator yellow
func (i *Indicator) Yellow() int {
	return i.yellow
}

// Green is Getter for the Indicators indicator green
func (i *Indicator) Green() int {
	return i.green
}
