package metrics

// HealthMetric defines a metric within a Squad Health Check
type HealthMetric struct {
	title    string
	tendency int
	Indicator
}

// New generates a new HealthMetric
func New(title string) HealthMetric {
	return HealthMetric{title: title}
}

// Title returns the title of the HealthMetric
func (h *HealthMetric) Title() string {
	return h.title
}

// Tendency returns the tendency of the HealthMetric
func (h *HealthMetric) Tendency() int {
	return h.tendency
}

// SetTitle sets the title of the HealthMetric
func (h *HealthMetric) SetTitle(title string) {
	h.title = title
}

// SetTendency sets the title of the HealthMetric
func (h *HealthMetric) SetTendency(tendency int) {
	h.tendency += tendency
}
