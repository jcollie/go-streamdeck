package streamdeck

//KeyDirection are keys L-to-R or R-to-L
type KeyDirection int

const (
	//KeyDirectionLeftToRight keys are L-to-R
	KeyDirectionLeftToRight KeyDirection = iota
	//KeyDirectionRightToLeft keys are R-to-L
	KeyDirectionRightToLeft
)
