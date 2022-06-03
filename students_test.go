package coverage

import (
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0o644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW

func TestPeople(t *testing.T) {
	p := make(People, 3)
	for i := 0; i < 3; i++ {
		p[i].firstName = "Model" + strconv.Itoa(i)
		p[i].lastName = "Factory" + strconv.Itoa(i)
		p[i].birthDay = time.Now()
	}

	t.Run("Len", func(t *testing.T) {
		require.Equal(t, 3, len(p))
	})

	t.Run("Less", func(t *testing.T) {
		ok := p.Less(0, 2)
		require.True(t, ok)

		ok = p.Less(2, 1)
		require.False(t, ok)
	})

	t.Run("Swap", func(t *testing.T) {
		p1 := p[1]
		p2 := p[2]
		p.Swap(1, 2)
		require.Equal(t, p1, p[2])
		require.Equal(t, p2, p[1])
	})
}

func TestMatrix(t *testing.T) {
	strMatrix := `12 3 4 9
5 45 3 76
4 2 1 53`
	m, err := New(strMatrix)
	if err != nil {
		t.Error(err)
	}

	t.Run("matrix rows/cols/data", func(t *testing.T) {
		require.Equal(t, m.rows, 3)
		require.Equal(t, m.cols, 4)
		require.Equal(t, m.data, []int{12, 3, 4, 9, 5, 45, 3, 76, 4, 2, 1, 53})
	})

	t.Run("matrix Rows get", func(t *testing.T) {
		require.Equal(t, m.Rows()[1], []int{5, 45, 3, 76})
	})

	t.Run("matrix Cols get", func(t *testing.T) {
		require.Equal(t, m.Cols()[3], []int{9, 76, 53})
	})

	t.Run("matrix item Set", func(t *testing.T) {
		ok := m.Set(22, 4, 111)
		require.False(t, ok)

		ok = m.Set(2, 3, 111)
		require.True(t, ok)
		require.Equal(t, m.Rows()[2][3], 111)
	})
}
