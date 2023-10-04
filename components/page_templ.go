// Code generated by templ@v0.2.364 DO NOT EDIT.

package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Page() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_1 := templ.GetChildren(ctx)
		if var_1 == nil {
			var_1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<!doctype html><html lang=\"en\"><head></head>")
		if err != nil {
			return err
		}
		err = templ.RenderScriptItems(ctx, templBuffer, App())
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("<body onload=\"")
		if err != nil {
			return err
		}
		var var_2 templ.ComponentScript = App()
		_, err = templBuffer.WriteString(var_2.Call)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"><h1>")
		if err != nil {
			return err
		}
		var_3 := `Check the console for a log!`
		_, err = templBuffer.WriteString(var_3)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h1></body></html>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}

func App() templ.ComponentScript {
	return templ.ComponentScript{
		Name:     `__templ_App_1e4f`,
		Function: `function __templ_App_1e4f(){console.log("Loaded!")}`,
		Call:     templ.SafeScript(`__templ_App_1e4f`),
	}
}