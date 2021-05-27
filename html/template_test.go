package html_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-template-unmarshaler/html"
)

func TestTemplate_Empty(t *testing.T) {
	t.Parallel()
	b := html.Template{}
	if !b.Empty() {
		t.Fatal("Tempalte.Empty() should be true")
	}
}

func TestNew(t *testing.T) {
	t.Parallel()
	b, err := html.New("foo")
	if err != nil {
		t.Fatal(err)
	}
	f, err := b.Execute(nil)
	if err != nil {
		t.Fatal(err)
	}
	if f != "foo" {
		t.Fatal(`Template must be "foo"`)
	}
}
