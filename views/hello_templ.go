// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"
import "strings"

import (
	"github.com/sain0136/go-rotaryanalytic/pkg"
	"strconv"
)

func body() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`background-color:white;`)
	templ_7745c5c3_CSSBuilder.WriteString(`margin:0;`)
	templ_7745c5c3_CSSID := templ.CSSID(`body`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func title() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`color:black;`)
	templ_7745c5c3_CSSBuilder.WriteString(`font-size:1.5rem;`)
	templ_7745c5c3_CSSBuilder.WriteString(`font-weight:bold;`)
	templ_7745c5c3_CSSBuilder.WriteString(`text-align:center;`)
	templ_7745c5c3_CSSID := templ.CSSID(`title`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func th() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`background-color:lightblue;`)
	templ_7745c5c3_CSSBuilder.WriteString(`font-weight:bold;`)
	templ_7745c5c3_CSSBuilder.WriteString(`padding-top:12px;`)
	templ_7745c5c3_CSSBuilder.WriteString(`padding-bottom:12px;`)
	templ_7745c5c3_CSSBuilder.WriteString(`text-align:left;`)
	templ_7745c5c3_CSSID := templ.CSSID(`th`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func main() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`display:flex;`)
	templ_7745c5c3_CSSBuilder.WriteString(`flex-direction:column;`)
	templ_7745c5c3_CSSBuilder.WriteString(`justify-content:center;`)
	templ_7745c5c3_CSSBuilder.WriteString(`align-items:center;`)
	templ_7745c5c3_CSSBuilder.WriteString(`gap:1rem;`)
	templ_7745c5c3_CSSBuilder.WriteString(`padding:1rem;`)
	templ_7745c5c3_CSSID := templ.CSSID(`main`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func bold() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`font-weight:bold;`)
	templ_7745c5c3_CSSID := templ.CSSID(`bold`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func button() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`background-color:lightblue;`)
	templ_7745c5c3_CSSBuilder.WriteString(`border:none;`)
	templ_7745c5c3_CSSBuilder.WriteString(`color:white;`)
	templ_7745c5c3_CSSBuilder.WriteString(`padding:15px 32px;`)
	templ_7745c5c3_CSSBuilder.WriteString(`text-align:center;`)
	templ_7745c5c3_CSSBuilder.WriteString(`text-decoration:none;`)
	templ_7745c5c3_CSSBuilder.WriteString(`display:inline-block;`)
	templ_7745c5c3_CSSBuilder.WriteString(`font-size:1rem;`)
	templ_7745c5c3_CSSBuilder.WriteString(`border-radius:12px;`)
	templ_7745c5c3_CSSBuilder.WriteString(`cursor:pointer;`)
	templ_7745c5c3_CSSID := templ.CSSID(`button`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func subheader() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`display:flex;`)
	templ_7745c5c3_CSSBuilder.WriteString(`flex-direction:row;`)
	templ_7745c5c3_CSSBuilder.WriteString(`justify-content:space-evenly;`)
	templ_7745c5c3_CSSBuilder.WriteString(`padding:1rem;`)
	templ_7745c5c3_CSSBuilder.WriteString(`align-items:center;`)
	templ_7745c5c3_CSSID := templ.CSSID(`subheader`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func sessionsTab() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`width:100%;`)
	templ_7745c5c3_CSSBuilder.WriteString(`display:flex;`)
	templ_7745c5c3_CSSBuilder.WriteString(`flex-direction:column;`)
	templ_7745c5c3_CSSBuilder.WriteString(`justify-content:center;`)
	templ_7745c5c3_CSSBuilder.WriteString(`align-items:center;`)
	templ_7745c5c3_CSSBuilder.WriteString(`gap:1rem;`)
	templ_7745c5c3_CSSBuilder.WriteString(`padding:1rem;`)
	templ_7745c5c3_CSSID := templ.CSSID(`sessionsTab`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func onWindowLoad(mode string) templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_onWindowLoad_8cc8`,
		Function: `function __templ_onWindowLoad_8cc8(mode){const params = new URLSearchParams(window.location.search);
 	 if (!params.has('user')) {
		if (mode == "DEV") {
			window.location.replace("http://localhost:3000/login");
		} else {
			window.location.replace("https://analytics.myrotaryprojects.org/login");
		}
	}
}`,
		Call:       templ.SafeScript(`__templ_onWindowLoad_8cc8`, mode),
		CallInline: templ.SafeScriptInline(`__templ_onWindowLoad_8cc8`, mode),
	}
}

func reloadPage() templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_reloadPage_e431`,
		Function: `function __templ_reloadPage_e431(){window.location.reload();
}`,
		Call:       templ.SafeScript(`__templ_reloadPage_e431`),
		CallInline: templ.SafeScriptInline(`__templ_reloadPage_e431`),
	}
}

// <div id="subheader" class={ subheader() }>
//
//	<p>
//		User: { name }
//	</p>
//	<div>Mode: { mode }</div>
//
// </div>
// @media (max-width: 600px) {
//
//	#subheader {
//	  flex-direction: column;
//	  gap: 1rem;
//	}
//
// }
func Hello(name string, fileStatus string, LogEntries []pkg.RotaryLog, mode string, lastPage int) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<style type=\"text/css\">\r\n\timg {\r\n\t\theight: 50px;\r\n\t\twidth: 50px;\r\n\t}\r\n\t.tabs {\r\n\t\twidth: 100%;\r\n\t}\r\n\t.header{\r\n\t\tdisplay: flex;\r\n\t\tflex-direction: row;\r\n\t\tjustify-content: space-between;\r\n\t\talign-items: center;\r\n\t\tpadding: 1rem;\r\n\t\tborder-bottom: 1px solid lightblue;\r\n\t}\r\n\t.sessions {\r\n\t\twidth: 100%;\r\n\t\tdisplay: flex;\r\n\t\tflex-direction: column;\r\n\t\tjustify-content: center;\r\n\t\talign-items: center;\r\n\t\tgap: 1rem;\r\n\t\tmargin: 2rem 0;\r\n\t\tfont-size: 3rem;\r\n\t\tfont-weight: bold;\r\n\t}\r\n\t.no-wrap {\r\n\t\twhite-space: nowrap;\r\n\t}\r\n\t.settings {\r\n\t\tcursor: pointer;\r\n\t}\r\n\r\n\t.settings:hover {\r\n\t\tcolor: lightblue;\r\n\t}\r\n\t.icons {\r\n\t\tdisplay: flex;\r\n\t\tflex-direction: row;\r\n\t\tgap: 1rem;\r\n\t}\r\n\ta {\r\n    color: black;      \r\n    text-decoration: none; \r\n}\r\n\t</style><script type=\"text/javascript\">\r\n\t\r\n\r\n  \t</script><html><head><link rel=\"stylesheet\" href=\"https://cdn.jsdelivr.net/npm/@shoelace-style/shoelace@2.15.0/cdn/themes/light.css\"><script type=\"module\" src=\"https://cdn.jsdelivr.net/npm/@shoelace-style/shoelace@2.15.0/cdn/shoelace-autoloader.js\"></script><script src=\"https://unpkg.com/htmx.org@1.9.12\" integrity=\"sha384-ujb1lZYygJmzgSwoxRggbCHcjc0rB2XoQrxeTUQyRjrOnlCoYta87iKBWq3EsdM2\" crossorigin=\"anonymous\"></script><script src=\"https://kit.fontawesome.com/fa78fc50e4.js\" crossorigin=\"anonymous\"></script><link rel=\"stylesheet\" href=\"https://assets.ubuntu.com/v1/vanilla-framework-version-4.10.0.min.css\"><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>MyRotary Analytics</title></head>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = Header().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 = []any{body()}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var2...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderScriptItems(ctx, templ_7745c5c3_Buffer, onWindowLoad(mode))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<body onload=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 templ.ComponentScript = onWindowLoad(mode)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var3.Call)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"")
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
		var templ_7745c5c3_Var4 = []any{main()}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var4...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ.CSSClasses(templ_7745c5c3_Var4).String()))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><sl-tab-group class=\"tabs\"><sl-tab slot=\"nav\" panel=\"logs\">Rotary Logs</sl-tab> <sl-tab slot=\"nav\" panel=\"sessions\">Sessions</sl-tab> <sl-tab-panel name=\"logs\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = Table(LogEntries, lastPage).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</sl-tab-panel> <sl-tab-panel name=\"sessions\"><div class=\"sessions\">Under Construction\r</div></sl-tab-panel></sl-tab-group></div></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func Header() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var5 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var5 == nil {
			templ_7745c5c3_Var5 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<header class=\"header\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 = []any{title()}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var6...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ.CSSClasses(templ_7745c5c3_Var6).String()))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">Myrotaryprojects.org Analytics </div><div class=\"icons\"><a href=\"https://myrotaryprojects.org\" title=\"MyRotaryProjects Homepage\"><i class=\"settings fa-solid fa-house fa-lg\"></i></a> <a title=\"Under Construction\"><i class=\"settings fa-solid fa-gear fa-lg\"></i></a></div></header>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func Table(LogEntries []pkg.RotaryLog, lastPage int) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var7 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var7 == nil {
			templ_7745c5c3_Var7 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<style type=\"text/css\">\r\n\t@media (max-width: 990px) {\r\n\t.column-to-hide {\r\n      display: none;\r\n    }\r\n\t}\r\n\t@media (max-width: 600px) {\r\n\t.column-to-hide-mobile {\r\n\t  display: none;\r\n\t}\r\n\t}\r\n\t.trow:hover {\r\n\t\tbackground-color: #ddd;\r\n\t\tcursor: pointer;\r\n\t}\r\n\t#table td, #table th {\r\n \t border: 1px solid #ddd;\r\n \t padding: 8px;\r\n\t}\r\n\t.no-wrap {\r\n\t\twhite-space: nowrap;\r\n\t}\r\n\t.buttons {\r\n\t\tdisplay: flex;\r\n\t\tflex-direction: row;\r\n\t\tjustify-content: flex-end;\r\n\t\tpadding: 1.5rem;\r\n\t\tgap: 2rem;\r\n\t\talign-items: center;\r\n\t\tflex-wrap: wrap;\r\n\t}\r\n\t.small-column{\r\n\t\twidth: 8rem;\r\n\t}\r\n\t</style><script type=\"text/javascript\">\r\n\twindow.currentPage = 1\r\n\twindow.getPageNumber = function(action) {\r\n\t\tif (action == \"next\") {\r\n\t\twindow.currentPage = window.currentPage + 1; // Increment page number and return it\r\n\t\t}\r\n\t\tif (action == \"prev\" && window.currentPage > 1) {\r\n\t\twindow.currentPage = window.currentPage - 1; // Decrement page number and return it\r\n\t\t} \r\n\t\twindow.prevButton.style.display = window.currentPage > 1 ? 'inline-flex' : 'none';\r\n\t\tif(window.currentPage > 1 && window.refreshIfNotFirstPage && window.reloadButton) {\r\n\t\t\twindow.refreshIfNotFirstPage.style.display = 'inline-flex';\r\n\t\t\twindow.reloadButton.style.display = 'none';\r\n\t\t} else {\r\n\t\t\twindow.refreshIfNotFirstPage.style.display = 'none';\r\n\t\t\twindow.reloadButton.style.display = 'inline-flex';\r\n\t\t}\r\n\t\tif (parseInt(window.currentPage) === (parseInt(window.lastPage) + 1)) {\t\r\n\t\t\twindow.nextButton.style.display = 'none';\r\n\t\t} else {\r\n\t\t\twindow.nextButton.style.display = 'inline-flex';\r\n\t\t}\r\n\t\treturn window.currentPage\r\n\t}\r\n\t\r\n\twindow.onload = function() {\r\n\t\twindow.nextButton = document.getElementById(\"nextPageButton\");\r\n\t\twindow.prevButton = document.getElementById(\"previousPageButton\");\r\n\t\twindow.reloadButton = document.getElementById(\"reload\");\r\n\t\twindow.refreshIfNotFirstPage = document.getElementById(\"refreshIfNotFirstPage\");\r\n\t\tif(window.refreshIfNotFirstPage){\r\n\t\t\twindow.refreshIfNotFirstPage.style.display = 'none';\r\n\t\t}\r\n\t\tif (window.currentPage == 1 && window.prevButton) {\r\n\t\t\twindow.prevButton.style.display = \"none\" ;\r\n\t\t}\t\r\n\t\tlet body = document.getElementById(\"tableContainer\")\r\n\t\tif (body) {\r\n\t\t\twindow.lastPage = body.getAttribute(\"last-page\")\r\n\t\t}\r\n\t\tlet showFileStatus = false\r\n\t\tconst buttonText = document.getElementById(\"showFileStatus\");\r\n\t\tif (buttonText) {\r\n\t\t\tbuttonText.innerHTML = \"Show File Status\";\r\n\t\t\tbuttonText.addEventListener(\"click\", function() {\r\n\t\t\t\tshowFileStatus = !showFileStatus\r\n\t\t\t\tconst fileStatus = document.getElementById(\"fileStatus\");\r\n\t\t\t\tif (showFileStatus) {\r\n\t\t\t\t\tbuttonText.innerHTML = \"Hide File Status\";\r\n\t\t\t\t\tfileStatus.style.display = \"block\";\r\n\t\t\t\t} else {\r\n\t\t\t\t\tif (fileStatus) {\r\n\t\t\t\t\t\tfileStatus.style.display = \"none\";\r\n\t\t\t\t\t}\r\n\t\t\t\t\tbuttonText.innerHTML = \"Show File Status\";\r\n\t\t\t\t}\r\n\t\t\t});\r\n\t\t}\r\n\t}\r\n\t\t\r\n\t</script><div class=\"buttons\"><div id=\"fileStatus\"></div><div><sl-button id=\"showFileStatus\" variant=\"primary\" hx-get=\"/getLogsPath\" hx-target=\"#fileStatus\"></sl-button></div><sl-button id=\"reload\" variant=\"primary\" hx-headers=\"{&#34;Cache-Control&#34;:&#34;no-cache&#34;}\" hx-get=\"/logs\" hx-target=\"#tableContainer\">Reload\r</sl-button> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderScriptItems(ctx, templ_7745c5c3_Buffer, reloadPage())
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<sl-button onclick=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var8 templ.ComponentScript = reloadPage()
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var8.Call)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" id=\"refreshIfNotFirstPage\" variant=\"primary\" oncli>Refresh\r</sl-button></div><div id=\"tableContainer\" last-page=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(strconv.Itoa(lastPage)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var9 = []any{bold()}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var9...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ.CSSClasses(templ_7745c5c3_Var9).String()))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">Raw Logs:</span> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if LogEntries == nil {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span>Error: LogEntries could not be processed</span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<table id=\"table\"><thead>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var10 = []any{th()}
			templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var10...)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<tr class=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ.CSSClasses(templ_7745c5c3_Var10).String()))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><th class=\"column-to-hide\" colspan=\"1\">Id</th><th class=\"column-to-hide-mobile\" colspan=\"1\">Timestamp</th><th class=\"small-column\" colspan=\"1\">Type</th><th class=\"small-column\" colspan=\"1\">Event</th><th class=\"small-column\" colspan=\"1\">Status</th><th class=\"small-column\" colspan=\"1\">Source</th><th class=\"column-to-hide\" colspan=\"1\">Message</th></tr></thead> <tbody id=\"tableBody\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			for _, entry := range LogEntries {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<tr class=\"trow\"><td class=\"column-to-hide\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var11 string
				templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(entry.UniqueId)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views\hello.templ`, Line: 320, Col: 50}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><td class=\"column-to-hide-mobile no-wrap\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var12 string
				templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs(entry.TimeStamp)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views\hello.templ`, Line: 321, Col: 66}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><td>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var13 string
				templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs(entry.Type)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views\hello.templ`, Line: 322, Col: 23}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><td>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var14 string
				templ_7745c5c3_Var14, templ_7745c5c3_Err = templ.JoinStringErrs(entry.Event)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views\hello.templ`, Line: 323, Col: 24}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var14))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><td>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var15 string
				templ_7745c5c3_Var15, templ_7745c5c3_Err = templ.JoinStringErrs(entry.Status)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views\hello.templ`, Line: 324, Col: 25}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var15))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><td>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var16 string
				templ_7745c5c3_Var16, templ_7745c5c3_Err = templ.JoinStringErrs(entry.Source)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views\hello.templ`, Line: 325, Col: 25}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var16))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><td class=\"column-to-hide\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var17 string
				templ_7745c5c3_Var17, templ_7745c5c3_Err = templ.JoinStringErrs(entry.Message)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views\hello.templ`, Line: 326, Col: 49}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var17))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td></tr>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</tbody></table>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div><sl-button id=\"previousPageButton\" variant=\"default\" hx-target=\"#tableContainer\" hx-get=\"/logs\" hx-vals=\"js:{page: getPageNumber(&#39;prev&#39;)}\">Previous Page</sl-button> <sl-button id=\"nextPageButton\" variant=\"default\" hx-target=\"#tableContainer\" hx-get=\"/logs\" hx-vals=\"js:{page: getPageNumber(&#39;next&#39;)}\">Next Page</sl-button></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
