package Database

import (
	"fmt"
	"testing"
)

func ExampleJsonb_Delete() {
	i := Jsonb{}
	r := i.Delete("home", map[string]interface{}{
		"id": "123345587674",
	})
	fmt.Println(r)
}

func TestJsonb_Delete(t *testing.T) {
	i := Jsonb{}
	r := i.Delete("home", map[string]interface{}{
		"id": "123345587674",
	})
	t.Log(r)
}

func BenchmarkJsonb_Delete(b *testing.B) {
	for i := 0; i > b.N; i++ {
		i := Jsonb{}
		i.Delete("home", map[string]interface{}{
			"id": "123345587674",
		})
	}
}

func TestJsonb_Set(t *testing.T) {
	i := Jsonb{}
	t.Run("without path", func(t *testing.T) {
		r := i.Set("name", map[string]interface{}{
			"id":     12,
			"name":   "ahmed mohamed yassin",
			"male":   true,
			"female": false,
		}, nil)
		t.Log(r)
	})

	t.Run("With path", func(t *testing.T) {
		r := i.Set("name", map[string]interface{}{
			"id":     12,
			"name":   "ahmed mohamed yassin",
			"male":   true,
			"female": false,
		}, "$some.path")
		t.Log(r)
	})
}

func TestJsonb_Get(t *testing.T) {
	i := Jsonb{}
	i.Get()
}

func TestJsonb_Increase(t *testing.T) {
	i := Jsonb{}
	r := i.Increase("notifications_count", map[string]int{
		"new":  1,
		"read": 1,
	})

	t.Log(r)
}

func TestJsonb_Decrease(t *testing.T) {
	i := Jsonb{}
	r := i.Decrease("notifications_count", map[string]int{
		"new":  1,
		"read": 1,
	})

	t.Log(r)
}
func TestJsonb_Find(t *testing.T) {
	i := Jsonb{}
	i.Find()
}
