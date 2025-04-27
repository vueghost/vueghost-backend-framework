package System

import (
	"bytes"
	"fmt"
	"github.com/kardianos/osext"
	"html/template"
	"path/filepath"
)

//View
type View struct {
	output       string
	data         interface{}
	viewTemplate *template.Template
}

//View
func (v *View) View(viewName string, data interface{}) {
	v.viewTemplate = template.Must(template.ParseFiles(v.viewDIRPath(viewName)))
	v.data = data
}

//ViewRender
func (v *View) ViewRender(viewName string, data interface{}) string {
	temp := template.Must(template.ParseFiles(v.viewDIRPath(viewName)))
	var output bytes.Buffer

	if err := temp.Execute(&output, data); err != nil {
		v.throwError(err)
		return ""
	}
	return output.String()
}

//viewDIRPath return main Views folder path.
func (v *View) viewDIRPath(viewName string) string {
	_, err := osext.ExecutableFolder()
	defer v.throwError(err)
	//fmt.Sprintf("%s/../Views/%s.gohtml", appPath, viewName)
	filePrefix, _ := filepath.Abs("../")

	return fmt.Sprintf("%s/Views/%s.html", filePrefix, viewName)
}

//throwError development mode throw View error.
func (v View) throwError(err error) {
	_ = fmt.Errorf(
		"View Error %v",
		err,
	)
}
