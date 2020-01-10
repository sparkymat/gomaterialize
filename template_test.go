package materialize_test

import (
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/sparkyat/materialize"
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
< />
</body>
</html>
`

	assert.Equal(t, expectedString, tpl.HTML().String())
}
