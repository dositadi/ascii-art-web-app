package tranformer

type AsciiTransform struct{}

func CreateNewAsciiTransformer() *AsciiTransform {
	return &AsciiTransform{}
}

func (at *AsciiTransform) CalculateStartLine(rn rune) int {
	return (int(rn-' ') * 9) + 1
}
