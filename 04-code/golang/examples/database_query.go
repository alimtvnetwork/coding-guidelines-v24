package examples

import (
	"context"
	"database/sql"
	"errors"

	"coding-guidelines/common/pkg/apperror"
)

// PluginSummary represents an active plugin record in the system.
type PluginSummary struct {
	ID       int64  `json:"id"`
	Slug     string `json:"slug"`
	Name     string `json:"name"`
	IsActive bool   `json:"isActive"`
}

// PluginRepository provides data access methods using Result wrappers.
type PluginRepository struct {
	db *sql.DB
}

// NewPluginRepository constructs a repository instance.
func NewPluginRepository(db *sql.DB) *PluginRepository {
	return &PluginRepository{db: db}
}

// validatePluginId checks input and simulates database error states.
func validatePluginId(id int64) *apperror.AppError {
	if id <= 0 {
		return apperror.New(apperror.ErrValidation, "plugin id must be positive")
	}

	if id == 404 {
		return apperror.New(apperror.ErrDatabaseNotFound, "plugin record not found")
	}

	if id == 500 {
		rawErr := errors.New("connection reset by peer")

		return apperror.Wrap(rawErr, apperror.ErrDatabaseQuery, "query execution failed")
	}

	return nil
}

// FindById queries a single plugin record by ID.
func (r *PluginRepository) FindById(ctx context.Context, id int64) apperror.Result[PluginSummary] {
	if err := validatePluginId(id); err != nil {
		return apperror.Fail[PluginSummary](err)
	}

	return apperror.Ok(PluginSummary{
		ID:       id,
		Slug:     "seo-optimizer",
		Name:     "SEO Optimizer Pro",
		IsActive: true,
	})
}

// ListActive retrieves all active plugins as a ResultSlice.
func (r *PluginRepository) ListActive(ctx context.Context) apperror.ResultSlice[PluginSummary] {
	items := []PluginSummary{
		{ID: 1, Slug: "cache-booster", Name: "Cache Booster", IsActive: true},
		{ID: 2, Slug: "security-shield", Name: "Security Shield", IsActive: true},
	}

	return apperror.OkSlice(items)
}
