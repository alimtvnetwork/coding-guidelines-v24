package lazyonce_test

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/errtype"
	"coding-guidelines/common/pkg/lazyonce"
	"coding-guidelines/common/pkg/result"
)

func TestLazyOnce_ExecutesExactlyOnce(t *testing.T) {
	var count int32
	lazy := lazyonce.New(func() (string, *appfault.AppError) {
		atomic.AddInt32(&count, 1)

		return "computed-val", nil
	})

	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			val, fault := lazy.Value()
			if fault != nil || val != "computed-val" {
				t.Errorf("unexpected result: val=%s, fault=%v", val, fault)
			}
		}()
	}

	wg.Wait()

	if atomic.LoadInt32(&count) != 1 {
		t.Fatalf("expected count 1, got %d", count)
	}

	if !lazy.IsEvaluated() {
		t.Fatalf("expected IsEvaluated to be true")
	}
}

func TestLazyOnce_ErrorPropagationAndResult(t *testing.T) {
	lazy := lazyonce.New(func() (int, *appfault.AppError) {
		return 0, appfault.New(errtype.Validation, "invalid state")
	})

	res := lazy.Result()
	if !res.IsFailed() {
		t.Fatalf("expected failure result")
	}

	if res.Fault().Message() != "invalid state" {
		t.Fatalf("expected fault message 'invalid state', got %s", res.Fault().Message())
	}
}

func TestLazyOnce_NewResultConstructor(t *testing.T) {
	lazy := lazyonce.NewResult(func() result.Result[string] {
		return result.Success("result-val")
	})

	val, fault := lazy.Value()
	if fault != nil || val != "result-val" {
		t.Fatalf("expected result-val, got %s, %v", val, fault)
	}
}

func TestLazyOnce_Reset(t *testing.T) {
	var count int32
	lazy := lazyonce.New(func() (int, *appfault.AppError) {
		atomic.AddInt32(&count, 1)

		return int(count), nil
	})

	val1, fault1 := lazy.Value()
	if fault1 != nil || val1 != 1 {
		t.Fatalf("expected 1, got %d", val1)
	}

	lazy.Reset()
	if lazy.IsEvaluated() {
		t.Fatalf("expected not evaluated after reset")
	}

	val2, fault2 := lazy.Value()
	if fault2 != nil || val2 != 2 {
		t.Fatalf("expected 2, got %d", val2)
	}
}

func TestLazyOnce1_SingleParameter(t *testing.T) {
	var callCount int32
	lazy := lazyonce.New1(func(multiplier int) (int, *appfault.AppError) {
		atomic.AddInt32(&callCount, 1)

		return multiplier * 10, nil
	})

	val1, fault1 := lazy.Value(5)
	if fault1 != nil || val1 != 50 {
		t.Fatalf("expected 50, got %d", val1)
	}

	val2, fault2 := lazy.Value(99)
	if fault2 != nil || val2 != 50 {
		t.Fatalf("expected cached 50, got %d", val2)
	}

	res := lazy.Result(100)
	if res.Data() != 50 {
		t.Fatalf("expected 50 from Result(), got %d", res.Data())
	}

	if atomic.LoadInt32(&callCount) != 1 {
		t.Fatalf("expected 1 execution, got %d", callCount)
	}
}

func TestLazyOnce2_TwoParameters(t *testing.T) {
	var callCount int32
	lazy := lazyonce.New2(func(a string, b int) (string, *appfault.AppError) {
		atomic.AddInt32(&callCount, 1)

		return a + "!", nil
	})

	val1, fault1 := lazy.Value("hello", 42)
	if fault1 != nil || val1 != "hello!" {
		t.Fatalf("expected hello!, got %s", val1)
	}

	val2, fault2 := lazy.Value("ignored", 0)
	if fault2 != nil || val2 != "hello!" {
		t.Fatalf("expected cached hello!, got %s", val2)
	}

	res := lazy.Result("x", 1)
	if res.Data() != "hello!" {
		t.Fatalf("expected hello! from Result(), got %s", res.Data())
	}

	if atomic.LoadInt32(&callCount) != 1 {
		t.Fatalf("expected 1 execution, got %d", callCount)
	}
}

func TestLazyOnce_ContextSuccessAndCancellation(t *testing.T) {
	lazy := lazyonce.New(func() (string, *appfault.AppError) {
		return "ctx-val", nil
	})

	val, fault := lazy.ValueContext(context.Background())
	if fault != nil || val != "ctx-val" {
		t.Fatalf("expected ctx-val, got %s, %v", val, fault)
	}

	cancelledCtx, cancel := context.WithCancel(context.Background())
	cancel()

	lazyNew := lazyonce.New(func() (string, *appfault.AppError) {
		return "should not run", nil
	})

	_, cFault := lazyNew.ValueContext(cancelledCtx)
	if cFault == nil {
		t.Fatal("expected timeout fault on canceled context")
	}
}

func TestLazyOnce_ResultContext(t *testing.T) {
	lazy := lazyonce.New(func() (int, *appfault.AppError) {
		return 999, nil
	})

	res := lazy.ResultContext(context.Background())
	if !res.IsSuccess() || res.Data() != 999 {
		t.Fatalf("expected 999, got %v", res)
	}

	resNil := lazy.ResultContext(nil)
	if !resNil.IsSuccess() || resNil.Data() != 999 {
		t.Fatalf("expected 999 with nil context")
	}
}

func TestLazyOnce1_ContextAndReset(t *testing.T) {
	var count int32
	lazy := lazyonce.New1(func(multiplier int) (int, *appfault.AppError) {
		atomic.AddInt32(&count, 1)

		return multiplier * 2, nil
	})

	val, fault := lazy.ValueContext(context.Background(), 10)
	if fault != nil || val != 20 {
		t.Fatalf("expected 20, got %d", val)
	}

	res := lazy.ResultContext(context.Background(), 10)
	if res.Data() != 20 {
		t.Fatalf("expected 20")
	}

	testLazyOnce1Reset(t, lazy)
}

func testLazyOnce1Reset(t *testing.T, lazy *lazyonce.LazyOnce1[int, int]) {
	lazy.Reset()
	if lazy.IsEvaluated() {
		t.Fatal("expected not evaluated after reset")
	}

	val, fault := lazy.Value(5)
	if fault != nil || val != 10 {
		t.Fatalf("expected 10 after reset, got %d", val)
	}
}

func TestLazyOnce2_ContextAndReset(t *testing.T) {
	lazy := lazyonce.New2(func(s string, n int) (string, *appfault.AppError) {
		return fmt.Sprintf("%s:%d", s, n), nil
	})

	val, fault := lazy.ValueContext(context.Background(), "port", 8080)
	if fault != nil || val != "port:8080" {
		t.Fatalf("expected port:8080, got %s", val)
	}

	res := lazy.ResultContext(context.Background(), "port", 8080)
	if res.Data() != "port:8080" {
		t.Fatalf("expected port:8080 from ResultContext")
	}

	lazy.Reset()
	if lazy.IsEvaluated() {
		t.Fatal("expected not evaluated after reset")
	}
}
