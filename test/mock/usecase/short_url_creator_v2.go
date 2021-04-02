package mock_usecase

import (
	context "context"

	"github.com/indrasaputra/url-shortener/entity"
)

// MockURLGeneratorV2 is a mock of URLGenerator interface
type MockURLGeneratorV2 struct {
	code string
	url  string
	err  *entity.Error
}

func NewMockURLGeneratorV2() *MockURLGeneratorV2 {
	return &MockURLGeneratorV2{}
}

// SetReturnValues sets return values.
func (m *MockURLGeneratorV2) SetReturnValues(code, url string, err *entity.Error) {
	m.code = code
	m.url = url
	m.err = err
}

// Generate mocks base method
func (m *MockURLGeneratorV2) Generate(_ uint) (string, string, *entity.Error) {
	return m.code, m.url, m.err
}

// MockCreateShortURLRepositoryV2 is a mock of CreateShortURLRepository interface
type MockCreateShortURLRepositoryV2 struct {
	value *entity.Error
}

func NewMockCreateShortURLRepositoryV2() *MockCreateShortURLRepositoryV2 {
	return &MockCreateShortURLRepositoryV2{}
}

// SetReturnValues sets return values.
func (m *MockCreateShortURLRepositoryV2) SetReturnValues(value *entity.Error) {
	m.value = value
}

// Generate mocks base method
func (m *MockCreateShortURLRepositoryV2) Save(_ context.Context, _ *entity.URL) *entity.Error {
	return m.value
}
