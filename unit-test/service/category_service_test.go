package service

import (
	"belajar-golang-unit-test/entity"
	"belajar-golang-unit-test/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

// mock jika gagal maka akan error
func TestCategoryService_GetNotFound(t *testing.T) {

	// program mock
	categoryRepository.Mock.On("FindById", "1").Return(nil) // mock

	category, err := categoryService.Get("1") // program yang di test
	assert.Nil(t, category) // assert
	assert.NotNil(t, err

}

// mock jika berhasil
func TestCategoryService_GetSuccess(t *testing.T) {
	category := entity.Category{
		Id:   "1",
		Name: "Laptop",
	} // mock
	categoryRepository.Mock.On("FindById", "2").Return(category) // mock

	result, err := categoryService.Get("2") // program yang di test
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)
}