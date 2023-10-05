package components

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

func inline(script templ.ComponentScript) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		// Can this be just script.Raw? (i.e. add a field called `Raw` that has exactly the code in the `script` block)
		// ...as this would make calculating the CSP easier (no need for a repeat construction of `+script.Function+";"+script.Call`)
		if _, err = io.WriteString(w, `<script type="text/javascript">`+script.Function+";"+script.Call+"</script>"); err != nil {
			return err
		}
		return nil
	})
}