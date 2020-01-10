package materialize_test

import (
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/sparkyat/materialize"
	"github.com/sparkymat/html"
)

func TestTemplate(t *testing.T) {
	tpl := materialize.NewTemplate("Test")

	expectedString := `<!DOCTYPE html>
<html>
<head>
<meta name="viewport" content="width=device-width, initial-scale=1" />
<title>Test</title>
</head>
<body>
</body>
</html>
`

	assert.Equal(t, expectedString, tpl.HTML().String())
}

func TestTemplateWithExtras(t *testing.T) {
	tpl := materialize.NewTemplate("Test").
		ExtraHeadNodes([]html.HeadNode{
			html.MetaInfo("author", "foo"),
		}).
		Css([]string{
			"http://style-link-1",
		}).
		Body(
			html.Div().Class("container").Children(
				html.H1("title"),
			),
		)

	expectedString := `<!DOCTYPE html>
<html>
<head>
<meta name="viewport" content="width=device-width, initial-scale=1" />
<title>Test</title>
<link rel="stylesheet" href="http://style-link-1" />
<meta name="author" content="foo" />
</head>
<body>
<div class="container">
<h1>title</h1>
</div>
</body>
</html>
`

	assert.Equal(t, expectedString, tpl.HTML().String())
}
