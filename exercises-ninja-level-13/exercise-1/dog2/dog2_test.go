package dog2

import (
	"fmt"
	"testing"
)

func TestDogyears(t *testing.T) {
	k := Dogyears(5)
	if k != 35 {
		t.Error("expected", 35, "got", k)
	}
}

func TestYearsTwo(t *testing.T) {
	k := YearsTwo(5)
	if k != 35 {
		t.Error("expected", 35, "got", k)
	}

}

func ExampleDogyears() {
	fmt.Println(Dogyears(10))
	//Output:
	//70
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(10))
	//Output:
	//70
}

func BenchmarkDogyears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Dogyears(10)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}
