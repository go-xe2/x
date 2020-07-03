package xhtml

import (
	"github.com/go-xe2/x/xtest"
	"testing"
)

func TestStripTags(t *testing.T) {
	src := `<p>Test paragraph.</p><!-- Comment -->  <a href="#fragment">Other text</a>`
	dst := `Test paragraph.  Other text`
	xtest.Assert(StripTags(src), dst)
}

func TestEntities(t *testing.T) {
	src := `A 'quote' "is" <b>bold</b>`
	dst := `A &#39;quote&#39; &#34;is&#34; &lt;b&gt;bold&lt;/b&gt;`
	xtest.Assert(Entities(src), dst)
	xtest.Assert(EntitiesDecode(dst), src)
}

func TestSpecialChars(t *testing.T) {
	src := `A 'quote' "is" <b>bold</b>`
	dst := `A &#39;quote&#39; &#34;is&#34; &lt;b&gt;bold&lt;/b&gt;`
	xtest.Assert(SpecialChars(src), dst)
	xtest.Assert(SpecialCharsDecode(dst), src)
}
