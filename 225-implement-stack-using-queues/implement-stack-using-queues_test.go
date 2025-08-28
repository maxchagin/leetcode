package implementstackusingqueues

import "testing"

// Test basic single push and top
func TestSinglePushTop(t *testing.T) {
	st := Constructor()
	st.Push(10)
	if val := st.Top(); val != 10 {
		t.Errorf("Expected top 10, got %d", val)
	}
}

// Test single push and pop
func TestSinglePushPop(t *testing.T) {
	st := Constructor()
	st.Push(5)
	if val := st.Pop(); val != 5 {
		t.Errorf("Expected pop 5, got %d", val)
	}
	if len(st.Val) != 0 {
		t.Errorf("Expected stack to be empty")
	}
}

// Test multiple push, top should return last pushed
func TestMultiplePushTop(t *testing.T) {
	st := Constructor()
	st.Push(1)
	st.Push(2)
	st.Push(3)
	if val := st.Top(); val != 3 {
		t.Errorf("Expected top 3, got %d", val)
	}
}

// Test push and pop order (LIFO check)
func TestPushPopOrder(t *testing.T) {
	st := Constructor()
	st.Push(1)
	st.Push(2)
	st.Push(3)
	if val := st.Pop(); val != 3 {
		t.Errorf("Expected pop 3, got %d", val)
	}
	if val := st.Pop(); val != 2 {
		t.Errorf("Expected pop 2, got %d", val)
	}
	if val := st.Pop(); val != 1 {
		t.Errorf("Expected pop 1, got %d", val)
	}
}

// Test empty on new stack
func TestEmptyNewStack(t *testing.T) {
	st := Constructor()
	if !st.Empty() {
		t.Errorf("Expected new stack to be empty")
	}
}

// Test push, pop, top interleaving
func TestInterleavePushPop(t *testing.T) {
	st := Constructor()
	st.Push(7)
	st.Push(8)
	if val := st.Top(); val != 8 {
		t.Errorf("Expected top 8, got %d", val)
	}
	if val := st.Pop(); val != 8 {
		t.Errorf("Expected pop 8, got %d", val)
	}
	if val := st.Top(); val != 7 {
		t.Errorf("Expected top 7, got %d", val)
	}
}

// Test multiple emptiness checks
func TestEmptyState(t *testing.T) {
	st := Constructor()
	st.Push(1)
	if st.Empty() {
		t.Errorf("Expected non-empty stack")
	}
	st.Pop()
	if !st.Empty() {
		t.Errorf("Expected empty after popping all elements")
	}
}

// Test sequence of pushes and pops
func TestPushPopSequence(t *testing.T) {
	st := Constructor()
	for i := 1; i <= 5; i++ {
		st.Push(i)
	}
	if val := st.Pop(); val != 5 {
		t.Errorf("Expected 5, got %d", val)
	}
	if val := st.Pop(); val != 4 {
		t.Errorf("Expected 4, got %d", val)
	}
	st.Push(99)
	if val := st.Top(); val != 99 {
		t.Errorf("Expected top 99, got %d", val)
	}
}

// Test mix of operations
func TestMixedOperations(t *testing.T) {
	st := Constructor()
	st.Push(3)
	st.Push(4)
	st.Pop()
	st.Push(5)
	if val := st.Top(); val != 5 {
		t.Errorf("Expected top to be 5, got %d", val)
	}
	st.Pop()
	if val := st.Top(); val != 3 {
		t.Errorf("Expected top to be 3, got %d", val)
	}
}

// Test stability with multiple calls
func TestRepeatedOperations(t *testing.T) {
	st := Constructor()
	for i := 0; i < 10; i++ {
		st.Push(i)
	}
	for i := 9; i >= 0; i-- {
		if val := st.Pop(); val != i {
			t.Errorf("Expected %d, got %d", i, val)
		}
	}
	if !st.Empty() {
		t.Errorf("Expected stack empty after all pops")
	}
}
