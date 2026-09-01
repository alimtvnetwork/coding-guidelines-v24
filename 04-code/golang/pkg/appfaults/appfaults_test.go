package appfaults_test

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/appfaults"
	"coding-guidelines/common/pkg/errtype"
)

func TestCollectionBasicOperations(t *testing.T) {
	c := appfaults.New()
	if c.HasError() || !c.IsSuccess() || c.Count() != 0 {
		t.Fatal("expected empty collection to be success")
	}

	c.Add(appfault.New(errtype.Validation, "bad field")).AddError(errors.New("io issue"))
	if !c.HasError() || c.Count() != 2 {
		t.Fatalf("expected 2 items, got %d", c.Count())
	}
}

func TestCollectionFilterAndTransform(t *testing.T) {
	c := appfaults.New()
	c.Add(appfault.New(errtype.Validation, "err1")).Add(appfault.New(errtype.Database, "err2"))

	filtered := c.FilterByType(errtype.Validation)
	if filtered.Count() != 1 || filtered.First().GetType() != errtype.Validation {
		t.Fatalf("expected 1 validation fault, got %d", filtered.Count())
	}

	comp := c.ToAppError(errtype.Execution, "Batch failure")
	if comp == nil || !comp.HasError() {
		t.Fatal("expected composite AppError")
	}
}

func TestMutexCollectionThreadSafety(t *testing.T) {
	mc := appfaults.NewMutexCollection()
	mc.Add(appfault.New(errtype.NotFound, "missing"))
	if !mc.HasError() || mc.Count() != 1 {
		t.Fatalf("expected 1 item in MutexCollection, got %d", mc.Count())
	}
}

func TestContextBinding(t *testing.T) {
	ctx, coll := appfaults.WithFaults(context.Background())
	appfaults.RecordContextError(ctx, appfault.New(errtype.Precondition, "unmet state"))

	if !coll.HasError() || coll.Count() != 1 {
		t.Fatalf("expected 1 recorded error in context, got %d", coll.Count())
	}
}

func TestCollectionJSONMarshaling(t *testing.T) {
	c := appfaults.New().Add(appfault.New(errtype.Validation, "invalid token"))
	data, err := json.Marshal(c)
	if err != nil || len(data) == 0 {
		t.Fatalf("failed to marshal Collection: %v", err)
	}

	restored := appfaults.New()
	if err := json.Unmarshal(data, restored); err != nil || restored.Count() != 1 {
		t.Fatalf("failed to unmarshal Collection: %v", err)
	}
}
