package mock_usecase

import (
	context "context"

	"github.com/indrasaputra/url-shortener/entity"
)

// MockURLGeneratorV2 is a mock of URLGenerator interface
type MockURLGeneratorV2 struct {
	value string
	err   *entity.Error
}

func NewMockURLGeneratorV2() *MockURLGeneratorV2 {
	return &MockURLGeneratorV2{}
}

// SetReturnValues sets return values.
func (m *MockURLGeneratorV2) SetReturnValues(value string, err *entity.Error) {
	m.value = value
	m.err = err
}

// Generate mocks base method
func (m *MockURLGeneratorV2) Generate(_ uint) (string, *entity.Error) {
	return m.value, m.err
}

// MockURLRepositoryV2 is a mock of URLRepository interface
type MockURLRepositoryV2 struct {
	value *entity.Error
}

func NewMockURLRepositoryV2() *MockURLRepositoryV2 {
	return &MockURLRepositoryV2{}
}

// SetReturnValues sets return values.
func (m *MockURLRepositoryV2) SetReturnValues(value *entity.Error) {
	m.value = value
}

// Generate mocks base method
func (m *MockURLRepositoryV2) Save(_ context.Context, _ *entity.URL) *entity.Error {
	return m.value
}
