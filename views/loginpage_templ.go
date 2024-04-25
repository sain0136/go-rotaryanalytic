// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"
import "strings"

func bodyLogin() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`background-color:white;`)
	templ_7745c5c3_CSSBuilder.WriteString(`display:flex;`)
	templ_7745c5c3_CSSBuilder.WriteString(`justify-content:center;`)
	templ_7745c5c3_CSSBuilder.WriteString(`align-items:center;`)
	templ_7745c5c3_CSSBuilder.WriteString(`height:100vh;`)
	templ_7745c5c3_CSSBuilder.WriteString(`color:black;`)
	templ_7745c5c3_CSSID := templ.CSSID(`bodyLogin`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func container() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`width:100%;`)
	templ_7745c5c3_CSSBuilder.WriteString(`max-width:400px;`)
	templ_7745c5c3_CSSBuilder.WriteString(`padding:20px;`)
	templ_7745c5c3_CSSBuilder.WriteString(`border-radius:5px;`)
	templ_7745c5c3_CSSBuilder.WriteString(`border:1px solid #e0e0e0;`)
	templ_7745c5c3_CSSBuilder.WriteString(`box-shadow:0 0 10px rgba(0, 0, 0, 0.1);`)
	templ_7745c5c3_CSSBuilder.WriteString(`display:flex;`)
	templ_7745c5c3_CSSBuilder.WriteString(`flex-direction:column;`)
	templ_7745c5c3_CSSBuilder.WriteString(`gap:1rem;`)
	templ_7745c5c3_CSSID := templ.CSSID(`container`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func h3() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`margin:0 0 20px;`)
	templ_7745c5c3_CSSBuilder.WriteString(`font-size:24px;`)
	templ_7745c5c3_CSSBuilder.WriteString(`font-weight:500;`)
	templ_7745c5c3_CSSBuilder.WriteString(`color:#333;`)
	templ_7745c5c3_CSSBuilder.WriteString(`text-align:center;`)
	templ_7745c5c3_CSSID := templ.CSSID(`h3`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func onchnage() templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_onchnage_803a`,
		Function: `function __templ_onchnage_803a(){console.log("dds");
}`,
		Call:       templ.SafeScript(`__templ_onchnage_803a`),
		CallInline: templ.SafeScriptInline(`__templ_onchnage_803a`),
	}
}

func buttonClick(mode string) templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_buttonClick_9818`,
		Function: `function __templ_buttonClick_9818(mode){if(	usernameState == "" || passwordState == ""){
		spanValidation.style.display = "block"
		return
	}
	spanValidation.style.display = "none"
	if (mode == "DEV"){
		const url = new URL('http://localhost:3000');  // replace with your URL
  		url.searchParams.append('user', 'true');
  		window.location.href = url.href;
	} else {
		const url = new URL('https://analytics.myrotaryprojects.org');  // replace with your URL
  		url.searchParams.append('user', 'secret');
  		window.location.href = url.href;
	}
}`,
		Call:       templ.SafeScript(`__templ_buttonClick_9818`, mode),
		CallInline: templ.SafeScriptInline(`__templ_buttonClick_9818`, mode),
	}
}

func LoginPage(mode string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<style type=\"text/css\">\r\n\t\tspan {\r\n\t\t\tcolor: red;\r\n\t\t\tmargin-top: 5px;\r\n\t\t\tdisplay: none;\r\n\t\t}\r\n\t</style><script type=\"text/javascript\">\r\n\tlet usernameState = ''\r\n\tlet passwordState = ''\r\n\tlet reqValidation =  false \r\n\tdocument.addEventListener('DOMContentLoaded', (event) => {\r\n\tconst username = document.getElementById('username');\r\n\tusername.addEventListener('sl-input', event => {\r\n\t\tusernameState = event.target.value\r\n\t})\r\n\tconst password = document.getElementById('password');\r\n\tpassword.addEventListener('sl-input', event => {\r\n\t\tpasswordState = event.target.value\r\n\t})\r\n\tconst spanValidation = document.getElementById('spanValidation');\r\n});\r\n</script><html><head><link rel=\"stylesheet\" href=\"https://cdn.jsdelivr.net/npm/@shoelace-style/shoelace@2.15.0/cdn/themes/light.css\"><script type=\"module\" src=\"https://cdn.jsdelivr.net/npm/@shoelace-style/shoelace@2.15.0/cdn/shoelace-autoloader.js\"></script><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>MyRotary Analytics</title></head>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 = []any{bodyLogin()}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var2...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<body class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ.CSSClasses(templ_7745c5c3_Var2).String()))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 = []any{container()}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var3...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ.CSSClasses(templ_7745c5c3_Var3).String()))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 = []any{h3()}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var4...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<h3 class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ.CSSClasses(templ_7745c5c3_Var4).String()))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">Rotary Analytics</h3><sl-input id=\"username\" label=\"Username\" help-text=\"Rotary super user email\" clearable required></sl-input> <sl-input id=\"password\" label=\"Password\" type=\"password\" password-toggle required></sl-input> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderScriptItems(ctx, templ_7745c5c3_Buffer, buttonClick(mode))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<sl-button onClick=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 templ.ComponentScript = buttonClick(mode)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var5.Call)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" variant=\"primary\">Submit</sl-button> <span id=\"spanValidation\">Please fill in all fields</span></div></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
