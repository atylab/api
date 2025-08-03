package services

type EncryptService interface {
	EncryptCaesarShift(text string, vector int) string
}

type EncryptServiceImpl struct{}

func (s *EncryptServiceImpl) EncryptCaesarShift(input string, shift int) string {
	shiftLetter := func(base, max rune, ch rune) rune {
		return base + (ch-base+rune(shift)+max)%max
	}

	runes := []rune(input)
	for i, ch := range runes {
		switch {
		case 'a' <= ch && ch <= 'z':
			runes[i] = shiftLetter('a', 26, ch)
		case 'A' <= ch && ch <= 'Z':
			runes[i] = shiftLetter('A', 26, ch)
		case '0' <= ch && ch <= '9':
			runes[i] = shiftLetter('0', 10, ch)
		default:
			// 記号・スペース・漢字などはそのまま
		}
	}
	return string(runes)
}
