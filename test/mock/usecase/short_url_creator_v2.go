package mock_usecase

import (
	context "context"

	"github.com/indrasaputra/aptx/entity"
)

// MockURLGeneratorV2 is a mock of URLGenerator interface
type MockURLGeneratorV2 struct {
	code string
	url  string
	err  error
}

func NewMockURLGeneratorV2() *MockURLGeneratorV2 {
	return &MockURLGeneratorV2{}
}

// SetReturnValues sets return values.
func (m *MockURLGeneratorV2) SetReturnValues(code, url string, err error) {
	m.code = code
	m.url = url
	m.err = err
}

// Generate mocks base method
func (m *MockURLGeneratorV2) Generate(_ uint) (string, string, error) {
	return m.code, m.url, m.err
}

// MockCreateShortURLRepositoryV2 is a mock of CreateShortURLRepository interface
type MockCreateShortURLRepositoryV2 struct {
	value error
}

func NewMockCreateShortURLRepositoryV2() *MockCreateShortURLRepositoryV2 {
	return &MockCreateShortURLRepositoryV2{}
}

// SetReturnValues sets return values.
func (m *MockCreateShortURLRepositoryV2) SetReturnValues(value error) {
	m.value = value
}

// Generate mocks base method
func (m *MockCreateShortURLRepositoryV2) Save(_ context.Context, _ *entity.URL) error {
	return m.value
}
