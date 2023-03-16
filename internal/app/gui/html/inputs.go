package html

type Input struct {
	Type         string `json:"type"`
	DefaultValue any    `json:"defaultValue"`
}

type selection struct {
	Input
	Options map[string]string `json:"options"`
}

func NewSelect(options map[string]string, defaultValue string) selection {
	return selection{
		Input: Input{
			Type:         "select",
			DefaultValue: defaultValue,
		},
		Options: options,
	}
}

type number struct {
	Input
	NumberOptions
}

type NumberOptions struct {
	Min  float64 `json:"min,omitempty"`
	Max  float64 `json:"max,omitempty"`
	Step float64 `json:"step,omitempty"`
}

func NewNumber(options NumberOptions, defaultValue float64) number {
	return number{
		Input: Input{
			Type:         "number",
			DefaultValue: defaultValue,
		},
		NumberOptions: NumberOptions{
			Min:  options.Min,
			Max:  options.Max,
			Step: options.Step,
		},
	}
}
