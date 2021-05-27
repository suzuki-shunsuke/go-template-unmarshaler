package text_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-template-unmarshaler/text"
)

func TestTemplate_Empty(t *testing.T) {
	t.Parallel()
	b := text.Template{}
	if !b.Empty() {
		t.Fatal("Tempalte.Empty() should be true")
	}
}

func TestNew(t *testing.T) {
	t.Parallel()
	b, err := text.New("foo")
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
