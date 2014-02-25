package utils

import (
	"fmt"
	"html/template"
	"reflect"
	"strings"
)

func GenerateForm(fields interface{}, action string, method string, errs map[string]string) template.HTML {
	var form string
	formname := strings.Split(reflect.TypeOf(fields).String(), ".")[1]

	form += fmt.Sprintf("<form name=\"%s\" action=\"%s\" method=\"%s\">\n", formname, action, strings.ToUpper(method))

	val := reflect.ValueOf(fields).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i).String()
		typeField := val.Type().Field(i)
		tag := typeField.Tag

		// Ignore "-" form fields
		if tag.Get("form") == "-" {
			continue
		}
		
		var extras string
		var label string

		fType := "text"
		fAttrs := strings.Split(tag.Get("attr"), ";")
		for _, element := range fAttrs {
			fAttr := strings.Split(element, ":")
			ele := strings.ToLower(fAttr[0])

			val := ""
			if len(fAttr) == 2 {
				val = fAttr[1]
			}
			if valueField != "" && val == "input" {
				val = valueField
			}
			if valueField == "" && val == "input" {
				val = ""
			}

			switch ele {
				case "alt":
					extras += fmt.Sprintf(" alt=\"%s\"", val)
				case "autofocus":
					extras += " autofocus"
				case "checked":
					extras += " checked"
				case "class":
					extras += fmt.Sprintf(" class=\"%s\"", val)
				case "label":
					label = fmt.Sprintf("\t<label for=\"%s\">%s</label>\n", tag.Get("form"), val)
				case "maxlength":
					extras += fmt.Sprintf(" maxlength=\"%s\"", val)
				case "min":
					extras += fmt.Sprintf(" min=\"%s\"", val)
				case "placeholder":
					extras += fmt.Sprintf(" placeholder=\"%s\"", val)
				case "readonly":
					extras += " readonly"
				case "required":
					extras += " required"
				case "type":
					fType = val
				case "value":
					extras += fmt.Sprintf(" value=\"%s\"", val)
			}
		}

		if label != "" {
			form += label
		}

		form += fmt.Sprintf("\t<input type=\"%s\" name=\"%s\" id=\"%s\"%s>\n", fType, tag.Get("form"), tag.Get("form"), extras)
		if errs[tag.Get("form")] != "" {
			form += fmt.Sprintf("\t<span class=\"flash\">%s</span>\n", errs[tag.Get("form")])
		}
	}

	form += "    </form>\n"

	return template.HTML(form)
}
