package examples_test

import (
	"bytes"
	"context"
	"testing"

	"coding-guidelines/common/examples"
	"coding-guidelines/common/pkg/apperror"
	"coding-guidelines/common/pkg/logger"
)

func TestDatabaseQuerySuccess(t *testing.T) {
	repo := examples.NewPluginRepository(nil)
	res := repo.FindById(context.Background(), 1)
	if !res.IsSafe() {
		t.Fatal("expected successful db query")
	}

	if res.Value().Slug != "seo-optimizer" {
		t.Fatalf("expected slug 'seo-optimizer', got %s", res.Value().Slug)
	}
}

func TestDatabaseQueryNotFound(t *testing.T) {
	repo := examples.NewPluginRepository(nil)
	res := repo.FindById(context.Background(), 404)
	if !res.HasError() {
		t.Fatal("expected error for 404 id")
	}

	if res.AppError().Code() != apperror.ErrDatabaseNotFound {
		t.Fatalf("expected ErrDatabaseNotFound, got %s", res.AppError().Code())
	}
}

func newTestWorkflowService(buf *bytes.Buffer) *examples.PluginWorkflowService {
	log := logger.New(logger.DefaultOptions().WithOutput(buf))
	repo := examples.NewPluginRepository(nil)
	client := examples.NewWordPressClient("https://site.example.com")

	return examples.NewPluginWorkflowService(repo, client, log)
}

func TestWorkflowServiceSuccess(t *testing.T) {
	svc := newTestWorkflowService(&bytes.Buffer{})
	res := svc.ActivateWorkflow(context.Background(), 10, 1)
	if !res.IsSafe() {
		t.Fatal("expected workflow to succeed")
	}

	if res.Value().PluginSummary.Slug != "seo-optimizer" {
		t.Fatalf("expected plugin slug 'seo-optimizer', got %s", res.Value().PluginSummary.Slug)
	}
}

func TestWorkflowServicePropagatesErrorWithoutRewrapping(t *testing.T) {
	svc := newTestWorkflowService(&bytes.Buffer{})
	res := svc.ActivateWorkflow(context.Background(), 10, 404)
	if !res.HasError() {
		t.Fatal("expected workflow to fail when plugin is not found in db")
	}

	if res.AppError().Code() != apperror.ErrDatabaseNotFound {
		t.Fatalf("expected original ErrDatabaseNotFound code, got %s", res.AppError().Code())
	}
}
