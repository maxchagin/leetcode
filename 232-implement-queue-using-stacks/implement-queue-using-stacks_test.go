package implementqueueusingstacks

import "testing"

// Test single push and pop
func TestSinglePushPop(t *testing.T) {
	q := Constructor()
	q.Push(1)
	if val := q.Pop(); val != 1 {
		t.Errorf("Expected pop 1, got %d", val)
	}
	if !q.Empty() {
		t.Errorf("Expected empty queue")
	}
}

// Test single push and peek
func TestSinglePushPeek(t *testing.T) {
	q := Constructor()
	q.Push(5)
	if val := q.Peek(); val != 5 {
		t.Errorf("Expected peek 5, got %d", val)
	}
	if q.Empty() {
		t.Errorf("Expected non-empty queue")
	}
}

// Test multiple pushes and pops maintain FIFO order
func TestFIFOOrder(t *testing.T) {
	q := Constructor()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	if val := q.Pop(); val != 1 {
		t.Errorf("Expected 1, got %d", val)
	}
	if val := q.Pop(); val != 2 {
		t.Errorf("Expected 2, got %d", val)
	}
	if val := q.Pop(); val != 3 {
		t.Errorf("Expected 3, got %d", val)
	}
	if !q.Empty() {
		t.Errorf("Expected empty after all pops")
	}
}

// Test peek does not remove the element
func TestPeekDoesNotRemove(t *testing.T) {
	q := Constructor()
	q.Push(10)
	q.Push(20)
	if val := q.Peek(); val != 10 {
		t.Errorf("Expected peek 10, got %d", val)
	}
	if val := q.Pop(); val != 10 {
		t.Errorf("Expected pop 10, got %d", val)
	}
}

// Test alternating push and pop
func TestAlternatePushPop(t *testing.T) {
	q := Constructor()
	q.Push(7)
	if val := q.Pop(); val != 7 {
		t.Errorf("Expected pop 7, got %d", val)
	}
	q.Push(8)
	q.Push(9)
	if val := q.Pop(); val != 8 {
		t.Errorf("Expected pop 8, got %d", val)
	}
	if val := q.Peek(); val != 9 {
		t.Errorf("Expected peek 9, got %d", val)
	}
}

// Test empty on new queue
func TestEmptyNewQueue(t *testing.T) {
	q := Constructor()
	if !q.Empty() {
		t.Errorf("Expected new queue to be empty")
	}
}

// Test push then empty check
func TestPushThenEmptyCheck(t *testing.T) {
	q := Constructor()
	q.Push(42)
	if q.Empty() {
		t.Errorf("Expected non-empty after push")
	}
}

// Test repeated push and peek
func TestRepeatedPushPeek(t *testing.T) {
	q := Constructor()
	q.Push(1)
	q.Push(2)
	if val := q.Peek(); val != 1 {
		t.Errorf("Expected peek 1, got %d", val)
	}
	q.Pop()
	if val := q.Peek(); val != 2 {
		t.Errorf("Expected peek 2, got %d", val)
	}
}

// Test mixed operations
func TestMixedOperations(t *testing.T) {
	q := Constructor()
	q.Push(3)
	q.Push(4)
	if val := q.Pop(); val != 3 {
		t.Errorf("Expected pop 3, got %d", val)
	}
	q.Push(5)
	if val := q.Peek(); val != 4 {
		t.Errorf("Expected peek 4, got %d", val)
	}
	if val := q.Pop(); val != 4 {
		t.Errorf("Expected pop 4, got %d", val)
	}
	if val := q.Pop(); val != 5 {
		t.Errorf("Expected pop 5, got %d", val)
	}
}

// Test stress with multiple pushes and pops
func TestStressOperations(t *testing.T) {
	q := Constructor()
	for i := 1; i <= 10; i++ {
		q.Push(i)
	}
	for i := 1; i <= 10; i++ {
		if val := q.Pop(); val != i {
			t.Errorf("Expected %d, got %d", i, val)
		}
	}
	if !q.Empty() {
		t.Errorf("Expected empty after all operations")
	}
}

func TestMyQueue(t *testing.T) {
	tests := []struct {
		name       string
		operations []string
		args       [][]int
		want       []interface{}
	}{
		{
			name:       "Single push and pop",
			operations: []string{"MyQueue", "push", "pop", "empty"},
			args:       [][]int{{}, {1}, {}, {}},
			want:       []interface{}{nil, nil, 1, true},
		},
		{
			name:       "Push two elements then peek and pop",
			operations: []string{"MyQueue", "push", "push", "peek", "pop", "empty"},
			args:       [][]int{{}, {1}, {2}, {}, {}, {}},
			want:       []interface{}{nil, nil, nil, 1, 1, false},
		},
		{
			name:       "Multiple pushes and pops in order",
			operations: []string{"MyQueue", "push", "push", "push", "pop", "pop", "pop", "empty"},
			args:       [][]int{{}, {1}, {2}, {3}, {}, {}, {}, {}},
			want:       []interface{}{nil, nil, nil, nil, 1, 2, 3, true},
		},
		{
			name:       "Peek does not remove element",
			operations: []string{"MyQueue", "push", "push", "peek", "peek", "pop"},
			args:       [][]int{{}, {5}, {6}, {}, {}, {}},
			want:       []interface{}{nil, nil, nil, 5, 5, 5},
		},
		{
			name:       "Alternate push and pop",
			operations: []string{"MyQueue", "push", "pop", "push", "push", "pop", "peek"},
			args:       [][]int{{}, {10}, {}, {20}, {30}, {}, {}},
			want:       []interface{}{nil, nil, 10, nil, nil, 20, 30},
		},
		{
			name:       "Pop all after several pushes",
			operations: []string{"MyQueue", "push", "push", "push", "pop", "pop", "pop", "empty"},
			args:       [][]int{{}, {7}, {8}, {9}, {}, {}, {}, {}},
			want:       []interface{}{nil, nil, nil, nil, 7, 8, 9, true},
		},
		{
			name:       "Peek after pops",
			operations: []string{"MyQueue", "push", "push", "push", "pop", "peek"},
			args:       [][]int{{}, {11}, {12}, {13}, {}, {}},
			want:       []interface{}{nil, nil, nil, nil, 11, 12},
		},
		{
			name:       "Empty check when queue never used",
			operations: []string{"MyQueue", "empty"},
			args:       [][]int{{}, {}},
			want:       []interface{}{nil, true},
		},
		{
			name:       "Push multiple then check empty before pop",
			operations: []string{"MyQueue", "push", "push", "empty"},
			args:       [][]int{{}, {42}, {43}, {}},
			want:       []interface{}{nil, nil, nil, false},
		},
		{
			name:       "Stress small sequence",
			operations: []string{"MyQueue", "push", "push", "push", "peek", "pop", "peek", "pop", "pop", "empty"},
			args:       [][]int{{}, {1}, {2}, {3}, {}, {}, {}, {}, {}, {}},
			want:       []interface{}{nil, nil, nil, nil, 1, 1, 2, 2, 3, true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var q MyQueue
			var got []interface{}

			for i, op := range tt.operations {
				switch op {
				case "MyQueue":
					q = Constructor()
					got = append(got, nil)
				case "push":
					q.Push(tt.args[i][0])
					got = append(got, nil)
				case "pop":
					val := q.Pop()
					got = append(got, val)
				case "peek":
					val := q.Peek()
					got = append(got, val)
				case "empty":
					val := q.Empty()
					got = append(got, val)
				default:
					t.Fatalf("unknown operation: %s", op)
				}
			}

			if len(got) != len(tt.want) {
				t.Fatalf("got length %d, want length %d", len(got), len(tt.want))
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Errorf("case %s: operation %s index %d got %v, want %v",
						tt.name, tt.operations[i], i, got[i], tt.want[i])
				}
			}
		})
	}
}
