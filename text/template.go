package text

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
	"text/template"
)

var newT = func(s string) (*template.Template, error) { //nolint:gochecknoglobals
	return template.New("_").Parse(s) //nolint:wrapcheck
}

func SetTemplateFunc(f func(string) (*template.Template, error)) {
	newT = f
}

type Template struct {
	template *template.Template
}

func New(s string) (*Template, error) {
	tpl, err := newT(s)
	if err != nil {
		return nil, fmt.Errorf("parse a template: %w", err)
	}
	return &Template{template: tpl}, nil
}

func NewForTest(t *testing.T, s string) *Template {
	t.Helper()
	a, err := New(s)
	if err != nil {
		t.Fatal(err)
	}
	return a
}

func (tpl *Template) Empty() bool {
	return tpl == nil || tpl.template == nil
}

func (tpl *Template) Execute(param interface{}) (string, error) {
	if tpl.template == nil {
		return "", nil
	}
	buf := &bytes.Buffer{}
	if err := tpl.template.Execute(buf, param); err != nil {
		return "", fmt.Errorf("render a template: %w", err)
	}
	return buf.String(), nil
}

func (tpl *Template) UnmarshalJSON(b []byte) error {
	var a string
	if err := json.Unmarshal(b, &a); err != nil {
		return err //nolint:wrapcheck
	}
	t, err := newT(a)
	if err != nil {
		return err
	}
	tpl.template = t
	return nil
}

func (tpl *Template) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var a string
	if err := unmarshal(&a); err != nil {
		return err
	}
	t, err := newT(a)
	if err != nil {
		return err
	}
	tpl.template = t
	return nil
}
