package components

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

func onLoad(script templ.ComponentScript) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		if _, err = io.WriteString(w, `<script type="text/javascript">`+"\r\n"+script.Call+"\r\n</script>"); err != nil {
			return err
		}
		return nil
	})
}