package System

import "testing"

//@global
type Todo struct {
	Title string
	Done  bool
}
type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

var viewData = TodoPageData{
	PageTitle: "My TODO list",
	Todos: []Todo{
		{Title: "Task 1", Done: false},
		{Title: "Task 2", Done: true},
		{Title: "Task 3", Done: true},
	},
}

//@View
func ExampleView_View() {
	v := View{}
	v.View("Home", viewData)
}
func TestView_View(t *testing.T) {
	v := View{}
	v.View("Home", viewData)
}
func BenchmarkView_View(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := View{}
		v.View("Home", viewData)
	}
}

//@ViewRender
func ExampleView_ViewRender() {
	v := View{}
	_ = v.ViewRender("Home", viewData)
}
func TestView_ViewRender(t *testing.T) {
	v := View{}
	t.Log(v.ViewRender("Home", viewData))
}
func BenchmarkView_ViewRender(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := View{}
		b.Log(v.ViewRender("Home", viewData))
	}
}

//@viewDIRPath
func ExampleView_viewDIRPath() {
	v := View{}
	v.viewDIRPath("viewFileName")
}
func TestView_viewDIRPath(t *testing.T) {
	v := View{}
	t.Log(v.viewDIRPath("viewFileName"))
}
func BenchmarkView_viewDIRPath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := View{}
		v.viewDIRPath("viewFileName")
	}
}
