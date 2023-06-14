package usecase

import (
	"EchoEdyP/RESTfulAPI-CleanArch/category/repository"
	"EchoEdyP/RESTfulAPI-CleanArch/helper"
	"EchoEdyP/RESTfulAPI-CleanArch/models/domain"
	"EchoEdyP/RESTfulAPI-CleanArch/models/request_response"
	"context"
	"database/sql"
)

type CategoryUseCaseImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
}

func (useCase *CategoryUseCaseImpl) Create(ctx context.Context, request request_response.CategoryCreateRequest) (response request_response.CategoryResponse, err error) {
	tx, err := useCase.DB.Begin()
	if err != nil {
		return request_response.CategoryResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category, err = useCase.CategoryRepository.Save(ctx, tx, category)
	if err != nil {
		return request_response.CategoryResponse{}, err
	}

	return helper.ToCategoryRespones(category), nil
}

func (useCase *CategoryUseCaseImpl) Update(ctx context.Context, request request_response.CategoryUpdateRequest) (response request_response.CategoryResponse, err error) {
	tx, err := useCase.DB.Begin()
	if err != nil {
		return request_response.CategoryResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	category, err := useCase.CategoryRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		return request_response.CategoryResponse{}, err
	}

	category.Name = request.Name

	category, err = useCase.CategoryRepository.Update(ctx, tx, category)
	if err != nil {
		return request_response.CategoryResponse{}, err
	}

	return helper.ToCategoryRespones(category), nil
}

func (useCase *CategoryUseCaseImpl) Delete(ctx context.Context, categoryId int) (err error) {
	tx, err := useCase.DB.Begin()
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	category, err := useCase.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		return err
	}

	err = useCase.CategoryRepository.Delete(ctx, tx, category)
	if err != nil {
		return err
	}

	return nil
}

func (useCase *CategoryUseCaseImpl) FindById(ctx context.Context, categoryId int) (response request_response.CategoryResponse, err error) {
	tx, err := useCase.DB.Begin()
	if err != nil {
		return request_response.CategoryResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	category, err := useCase.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		return request_response.CategoryResponse{}, err
	}

	return helper.ToCategoryRespones(category), nil
}

func (useCase *CategoryUseCaseImpl) FindAll(ctx context.Context) (response []request_response.CategoryResponse, err error) {
	tx, err := useCase.DB.Begin()
	if err != nil {
		return []request_response.CategoryResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	categories, err := useCase.CategoryRepository.FindAll(ctx, tx)
	if err != nil {
		return []request_response.CategoryResponse{}, err
	}

	return helper.ToCategoryResponses(categories), nil
}