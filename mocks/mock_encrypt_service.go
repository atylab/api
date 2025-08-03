package mocks

type MockEncryptService struct {
	CalledWithText   string
	CalledWithVector int
	ReturnValue      string
}

func (m *MockEncryptService) EncryptCaesarShift(text string, vector int) string {
	m.CalledWithText = text
	m.CalledWithVector = vector
	return m.ReturnValue
}
