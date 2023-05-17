package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}

		got := SumArray(numbers)
		want := 15

		assertNoError(t, got, want, numbers)
	})

	t.Run("collection of any numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}

		got := SumSlice(numbers)
		want := 21

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	assertNoErrorDeep(t, got, want)

}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2}, []int{0, 9})
	want := []int{2, 9}

	assertNoErrorDeep(t, got, want)
}

func assertNoError(t testing.TB, got, want int, numbers [5]int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}

func assertNoErrorDeep(t testing.TB, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}