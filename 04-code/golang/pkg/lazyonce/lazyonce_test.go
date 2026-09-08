package lazyonce_test

import (
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

	val1, _ := lazy.Value()
	if val1 != 1 {
		t.Fatalf("expected 1, got %d", val1)
	}

	lazy.Reset()
	if lazy.IsEvaluated() {
		t.Fatalf("expected not evaluated after reset")
	}

	val2, _ := lazy.Value()
	if val2 != 2 {
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

	val2, _ := lazy.Value(99)
	if val2 != 50 {
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

	val2, _ := lazy.Value("ignored", 0)
	if val2 != "hello!" {
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
