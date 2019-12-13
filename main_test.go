package csscomp

import (
	"fmt"
	"testing"
)

var css string = `h1 { font-size: 2em; margin: 0.67em 0; }`
var css2 string = `font-family: "Helvetica Neue Light", "HelveticaNeue-Light", "Helvetica Neue", Calibri, Helvetica, Arial;`

func TestRgbToHex(t *testing.T) {
	str := "rgb(50,255,120)"
	wanted := "#32ff78"
	result := rgbToHex(str)
	if result != wanted {
		t.Errorf("Incorrect, got: %s, want: %s.", result, wanted)
	}
}

func TestMinColor(t *testing.T) {
	str := "#99ddff"
	wanted := "#9df"
	result := minColor(str)
	if result != wanted {
		t.Errorf("Incorrect, got: %s, want: %s.", result, wanted)
	}
}

func TestRmLeadingZeros(t *testing.T) {
	str := "h1{font-size:2em;margin:0.67em 0;}"
	wanted := "h1{font-size:2em;margin: .67em 0;}"
	result := rmLeadingZeros(str)
	if result != wanted {
		t.Errorf("Incorrect, got: %s, want: %s.", result, wanted)
	}
}

func TestRmSpaces(t *testing.T) {
	str := css
	wanted := "h1{font-size:2em;margin:0.67em 0;}"
	result := rmSpaces(str)
	if result != wanted {
		t.Errorf("Incorrect, got: %s, want: %s.", result, wanted)
	}
}

func TestRmQuotes(t *testing.T) {
	str := `id1[prop*="abc def"], id2, id3[prop="gdsh"], .class1`
	wanted := `id1[prop*="abc def"], id2, id3[prop=gdsh], .class1`
	result := rmQuotes(str)
	if result != wanted {
		t.Errorf("Incorrect, got: %s, want: %s.", result, wanted)
	}
	str = `font-family: "Helvetica Neue Light", "HelveticaNeue-Light", "Helvetica Neue", Calibri, Helvetica, Arial;`
	wanted = `font-family:"Helvetica Neue Light",HelveticaNeue-Light,"Helvetica Neue", Calibri, Helvetica, Arial;`
	result = rmQuotes(str)
	if result != wanted {
		t.Errorf("Incorrect, got: %s, want: %s.", result, wanted)
	}
}

func TestRmSemicolon(t *testing.T) {
	str := "h1{font-size:2em; margin:0.67em 0; }"
	wanted := "h1{font-size:2em; margin:0.67em 0}"
	result := rmSemicolon(str)
	if result != wanted {
		t.Errorf("Incorrect, got: %s, want: %s.", result, wanted)
	}
}

func TestMsToSecond(t *testing.T) {
	str := "h1{font-size:2em; margin:0.67em 0; duration: 220ms }"
	wanted := "h1{font-size:2em; margin:0.67em 0; duration: 0.22s }"
	result := msToSecond(str)
	if result != wanted {
		t.Errorf("Incorrect, got: %s, want: %s.", result, wanted)
	}
	str = `{
        -webkit-transform: translateY(66px);
        -ms-transform: translateY(66px);
        transform: translateY(66px)
    }`
	wanted = `{
        -webkit-transform: translateY(66px);
        -ms-transform: translateY(66px);
        transform: translateY(66px)
    }`
	result = msToSecond(str)
	if result != wanted {
		t.Errorf("Incorrect, got: %s, want: %s.", result, wanted)
	}
}

func TestGet(t *testing.T) {
	result := Get(fixtures)
	fmt.Println("bytes original--> ", len(fixtures), " bytes compressed--> ", len(result))
	fmt.Println(result)
}

func BenchmarkGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Get(fixtures)
	}
}

var fixtures string = `@font-face {
    font-family: 'Open Sans';
    font-weight: normal;
    font-style: normal;
    src: url("../fonts/OpenSans-Regular-webfont.eot");
    src: local("Open Sans"), local("OpenSans"), url("../fonts/OpenSans-Regular-webfont.eot?#iefix") format("embedded-opentype"), url("../fonts/OpenSans-Regular-webfont.woff") format("woff"), url("../fonts/OpenSans-Regular-webfont.svg#open_sansregular") format("svg"); }
  
  @font-face {
    font-family: 'Open Sans Light';
    font-weight: normal;
    font-style: normal;
    src: url("../fonts/OpenSans-Light-webfont.eot");
    src: local("Open Sans Light"), local("OpenSans Light"), url("../fonts/OpenSans-Light-webfont.eot?#iefix") format("embedded-opentype"), url("../fonts/OpenSans-Light-webfont.woff") format("woff"), url("../fonts/OpenSans-Light-webfont.svg#open_sanslight") format("svg"); }
  
  html {
    overflow: auto;
    background-color: #ffffff;
    font-size: 1em;
    scroll-behavior: smooth; }
  
  body {
    font-family: 'Open Sans', sans-serif;
    line-height: 1.5;
    color: #414042;
    margin: 0;
    padding: 64px 0 0 0; }
  
  a,
  a:visited,
  a:active {
    color: #fe5000;
    text-decoration: none; }
  
  a:hover {
    text-decoration: underline;
    font-weight: bold; }
  
  header {
    display: block;
    padding: 0px 4px; }
    header h2,
    header .signature {
      color: #fe5000;
      overflow-wrap: break-word; }
  
  tt,
  code,
  kbd,
  samp {
    font-family: Consolas, Monaco, 'Andale Mono', monospace; }
  
  p > code {
    background-color: #ececec; }
  
  article dl {
    margin-bottom: 40px; }
  
  article img {
    max-width: 100%; }
  
  section {
    background-color: #fff;
    display: block;
    width: 70%;
    max-width: 960px;
    margin: 5em auto 2.5em;
    padding: 12px 24px;
    border-bottom: 1px solid #ccc; }
  
  nav {
    background-color: #414042;
    position: fixed; }
    nav h3 {
      margin-top: 12px;
      color: #fe5000; }
      nav h3 .subsection-title {
        display: inline-block;
        margin-top: 1em;
        margin-bottom: 0; }
  
  footer {
    background-color: #414042;
    color: #ececec;
    display: block;
    padding: 2em;
    margin-top: 3em;
    text-align: center;
    font-style: italic;
    font-size: 90%; }
  
  h1,
  h2,
  h3,
  h4 {
    font-weight: 200;
    margin: 0; }
  
  h1 {
    font-family: 'Open Sans Light', sans-serif;
    letter-spacing: -2px; }
  
  h2,
  h3.subsection-title {
    font-size: 1.9em;
    font-weight: 700;
    letter-spacing: -1px;
    margin-bottom: 0.75em;
    color: #fe5000; }
  
  h3.subsection-title {
    margin-top: 5em; }
  
  h3 {
    font-size: 1.5em;
    letter-spacing: -0.5px;
    margin-bottom: 0.75em; }
  
  h5,
  .container-overview .subsection-title {
    font-size: 120%;
    font-weight: bold;
    letter-spacing: -0.01em;
    margin: 8px 0 3px 0;
    color: #fe5000; }
  
  h5 {
    color: #414042; }
  
  h6 {
    font-size: 100%;
    letter-spacing: -0.01em;
    margin: 6px 0 3px 0;
    font-style: italic; }
  
  pre {
    border: 1px solid #337ab7;
    padding: 5px;
    overflow: auto;
    margin-bottom: 2em; }
  
  .icon-bars {
    display: none;
    position: absolute;
    right: 1em;
    font-size: 2em; }
  
  .icon-bars-show {
    display: inline-block; }
  
  .topnav {
    background-color: #414042;
    display: flex;
    justify-content: space-between;
    top: 0;
    left: 0;
    flex-direction: column;
    width: 100%;
    height: auto;
    box-sizing: border-box;
    box-shadow: 0 1px 3px #414042;
    /*transition: 0.3s;*/
    z-index: 5; }
  
  .breadcrumb {
    display: flex;
    flex-wrap: wrap;
    background-color: #fe5000;
    color: #414042;
    font-size: 0.8em;
    font-weight: bold;
    padding-left: 2em;
    transition: 0.1s;
    z-index: 5; }
  
  .showBreadcrumb {
    display: flex;
    font-size: 1.5em;
    justify-content: space-between;
    background-color: #414042;
    color: #ececec; }
  
  /*Style of list of links in the top bar*/
  .topnav-menu {
    color: #ececec;
    display: flex;
    align-items: center;
    padding: 0;
    -webkit-padding-start: 0;
            padding-inline-start: 0;
    -webkit-margin-before: 0;
            margin-block-start: 0;
    -webkit-margin-end: 0;
            margin-inline-end: 0;
    -webkit-margin-after: 0;
            margin-block-end: 0;
    list-style-type: none;
    text-transform: uppercase; }
    .topnav-menu a {
      color: unset;
      padding-left: 0.5em;
      text-decoration: none; }
  
  /*Style of .topnav-menu elements*/
  .topnav-menu-option {
    display: flex;
    position: relative;
    align-items: center;
    justify-content: center;
    flex: 1 1 auto;
    text-align: left;
    height: 100%;
    margin: 0;
    padding: 1em;
    cursor: pointer; }
    .topnav-menu-option a {
      margin: 0;
      padding: 0; }
    .topnav-menu-option:hover {
      font-weight: bold;
      color: #ffffff; }
  
  .selectedOption {
    background-color: #fe5000;
    color: #414042;
    font-weight: bold; }
  
  .topnav-submenu {
    background-color: #ececec;
    display: flex;
    position: fixed;
    left: 45%;
    top: -100%;
    flex-direction: column;
    flex-wrap: wrap;
    height: 60%;
    width: 50%;
    padding-top: 16px;
    box-shadow: -1px 1px 3px #414042;
    overflow: auto;
    font-weight: bold;
    transition: top 0.3s ease-out;
    z-index: 4; }
    .topnav-submenu a {
      color: #414042;
      display: block;
      margin: 1px;
      padding: 0.5em;
      border-style: solid;
      border-width: 1px;
      border-color: #cdcdcd; }
      .topnav-submenu a:hover {
        background-color: #414042;
        color: #ffffff;
        font-weight: bold;
        border-color: #fe5000;
        text-decoration: none;
        box-shadow: -1px 1px 3px #414042; }
  
  .breadcrumb-menu {
    display: none;
    right: 1em;
    bottom: 0;
    font-size: 1em; }
  
  .breadcrumb-menu-option {
    display: inline-block;
    padding: 0 1em 0.5em 1em;
    font-size: 14px;
    cursor: pointer; }
  
  #Members.topnav-submenu {
    z-index: 3; }
  
  #Methods.topnav-submenu {
    z-index: 3; }
  
  .showSubMenu {
    top: 62px; }
  
  h4 {
    color: #337ab7;
    position: relative;
    margin-top: 5em;
    margin-bottom: 0.95em;
    padding: 5px;
    font-size: 1.2em;
    font-weight: bold;
    letter-spacing: -0.33px; }
  
  .name,
  .signature {
    font-family: Consolas, Monaco, 'Andale Mono', monospace; }
  
  .signature-attributes {
    color: #8e8c8f;
    font-size: 60%;
    font-style: italic;
    font-weight: lighter; }
  
  .type-signature {
    color: #9fc4e4; }
  
  .underline {
    display: block;
    background-color: #fe5000;
    position: absolute;
    width: 100%;
    padding: 1px;
    margin-top: 0.2em;
    margin-left: -4px; }
  
  table {
    border-spacing: 0;
    border-collapse: collapse;
    width: 100%; }
  
  td,
  th {
    margin: 0px;
    text-align: left;
    vertical-align: top;
    padding: 4px 6px;
    border-bottom: 1px solid #337ab7;
    display: table-cell; }
  
  thead tr {
    font-weight: bold;
    border-bottom: 2px solid #337ab7; }
  
  tr:nth-child(even) {
    background-color: #eff5fa; }
  
  .params .name,
  .props .name,
  .name code {
    color: #414042;
    font-family: Consolas, Monaco, 'Andale Mono', monospace;
    font-size: 100%; }
  
  .params td.description > p:first-child,
  .props td.description > p:first-child {
    margin-top: 0;
    padding-top: 0; }
  
  .params td.description > p:last-child,
  .props td.description > p:last-child {
    margin-bottom: 0;
    padding-bottom: 0; }
  
  .details {
    margin-top: 14px; }
    .details dt {
      width: 120px;
      float: left;
      padding-left: 10px;
      padding-top: 6px; }
    .details dd {
      margin-left: 70px; }
    .details ul {
      margin: 0;
      list-style-type: none; }
    .details li {
      margin-left: 30px;
      padding-top: 6px; }
    .details pre.prettyprint {
      margin: 0; }
    .details .object-value {
      padding-top: 0; }
  
  .source {
    border: 1px solid #ddd;
    width: 80%;
    overflow: auto; }
  
  .prettyprint.source {
    width: inherit; }
  
  .source code {
    font-size: 100%;
    line-height: 18px;
    display: block;
    padding: 4px 12px;
    margin: 0;
    background-color: #fff;
    color: #414042; }
  
  .prettyprint code span.line {
    display: inline-block; }
  
  .prettyprint.linenums {
    padding-left: 70px;
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    user-select: none; }
  
  .prettyprint.linenums ol {
    padding-left: 0; }
  
  .prettyprint.linenums li {
    border-left: 3px #ddd solid; }
  
  .prettyprint.linenums li.selected,
  .prettyprint.linenums li.selected * {
    background-color: lightyellow; }
  
  .prettyprint.linenums li * {
    -webkit-user-select: text;
    -moz-user-select: text;
    -ms-user-select: text;
    user-select: text; }
  
  #main {
    width: 100%;
    margin-top: 38px;
    transition: width .2s; }
  
  .ancestors,
  .attribs {
    color: #999; }
  
  .ancestors a,
  .attribs a {
    color: #999 !important;
    text-decoration: none; }
  
  .clear {
    clear: both; }
  
  .description {
    margin-bottom: 1em;
    margin-top: 1em; }
  
  .code-caption {
    font-style: italic;
    font-size: 107%;
    margin: 0; }
  
  .important {
    font-weight: bold;
    color: #950B02; }
  
  .yes-def {
    text-indent: -1000px; }
  
  .class-description {
    font-size: 130%;
    line-height: 140%;
    margin-bottom: 1em;
    margin-top: 1em; }
    .class-description:empty {
      margin: 0; }
  
  .variation {
    display: none; }
  
  .disabled {
    opacity: 0.2;
    pointer-events: none; }`
