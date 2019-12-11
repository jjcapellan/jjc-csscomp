package csscomp

import (
	"fmt"
	"testing"
)

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
	str := "prop: 0.456 10.2;"
	wanted := "prop: .456 10.2;"
	result := rmLeadingZeros(str)
	if result != wanted {
		t.Errorf("Incorrect, got: %s, want: %s.", result, wanted)
	}
}

func TestRmSpaces(t *testing.T) {
	str := "  id1, id2, id3, .class1 .class2 { prop : 12px ; prop:  15px 10px; }  "
	wanted := " id1,id2,id3,.class1 .class2{prop:12px;prop:15px 10px;}"
	result := rmSpaces(str)
	if result != wanted {
		t.Errorf("Incorrect, got: %s, want: %s.", result, wanted)
	}
}

func TestGet(t *testing.T) {
	result := Get(fixtures)
	fmt.Println("bytes original--> ", len(fixtures), " bytes compressed--> ", len(result))
	//fmt.Println(result)
}

func BenchmarkGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Get(fixtures)
	}
}

var fixtures string = `.cc_banner-wrapper {
    z-index: 100001;
    position: relative
}

.cc_container .cc_btn {
    cursor: pointer;
    text-align: center;
    font-size: 0.6em;
    transition: font-size 200ms;
    line-height: 1em
}

.cc_container .cc_message {
    font-size: 0.6em;
    transition: font-size 200ms;
    margin: 0;
    padding: 0;
    line-height: 1.5em
}

.cc_container .cc_logo {
    display: none;
    text-indent: -1000px;
    overflow: hidden;
    width: 100px;
    height: 22px;
    background-size: cover;
    background-image: url(https://s3-eu-west-1.amazonaws.com/assets.cookieconsent.silktide.com/cookie-consent-logo.png);
    opacity: 0.9;
    transition: opacity 200ms
}

.cc_container .cc_logo:hover,.cc_container .cc_logo:active {
    opacity: 1
}

@media screen and (min-width: 500px) {
    .cc_container .cc_btn {
        font-size:0.8em
    }

    .cc_container .cc_message {
        font-size: 0.8em
    }
}

@media screen and (min-width: 768px) {
    .cc_container .cc_btn {
        font-size:1em
    }

    .cc_container .cc_message {
        font-size: 1em;
        line-height: 1em
    }
}

@media screen and (min-width: 992px) {
    .cc_container .cc_message {
        font-size:1em
    }
}

@media print {
    .cc_banner-wrapper,.cc_container {
        display: none
    }
}

.cc_container {
    position: fixed;
    left: 0;
    right: 0;
    bottom: 0;
    overflow: hidden;
    padding: 10px
}

.cc_container .cc_btn {
    padding: 8px 10px;
    background-color: #f1d600;
    cursor: pointer;
    transition: font-size 200ms;
    text-align: center;
    font-size: 0.6em;
    display: block;
    width: 33%;
    margin-left: 10px;
    float: right;
    max-width: 120px
}

.cc_container .cc_message {
    transition: font-size 200ms;
    font-size: 0.6em;
    display: block
}

@media screen and (min-width: 500px) {
    .cc_container .cc_btn {
        font-size:0.8em
    }

    .cc_container .cc_message {
        margin-top: 0.5em;
        font-size: 0.8em
    }
}

@media screen and (min-width: 768px) {
    .cc_container {
        padding:15px 30px 15px
    }

    .cc_container .cc_btn {
        font-size: 1em;
        padding: 8px 15px
    }

    .cc_container .cc_message {
        font-size: 1em
    }
}

@media screen and (min-width: 992px) {
    .cc_container .cc_message {
        font-size:1em
    }
}

.cc_container {
    background: #222;
    color: #fff;
    font-size: 17px;
    font-family: "Helvetica Neue Light", "HelveticaNeue-Light", "Helvetica Neue", Calibri, Helvetica, Arial;
    box-sizing: border-box
}

.cc_container ::-moz-selection {
    background: #ff5e99;
    color: #fff;
    text-shadow: none
}

.cc_container .cc_btn,.cc_container .cc_btn:visited {
    color: #000;
    background-color: #f1d600;
    transition: background 200ms ease-in-out,color 200ms ease-in-out,box-shadow 200ms ease-in-out;
    -webkit-transition: background 200ms ease-in-out,color 200ms ease-in-out,box-shadow 200ms ease-in-out;
    border-radius: 5px;
    -webkit-border-radius: 5px
}

.cc_container .cc_btn:hover,.cc_container .cc_btn:active {
    background-color: #fff;
    color: #000
}

.cc_container a,.cc_container a:visited {
    text-decoration: none;
    color: #31a8f0;
    transition: 200ms color
}

.cc_container a:hover,.cc_container a:active {
    color: #b2f7ff
}

@-webkit-keyframes slideUp {
    0% {
        -webkit-transform: translateY(66px);
        transform: translateY(66px)
    }

    100% {
        -webkit-transform: translateY(0);
        transform: translateY(0)
    }
}

@keyframes slideUp {
    0% {
        -webkit-transform: translateY(66px);
        -ms-transform: translateY(66px);
        transform: translateY(66px)
    }

    100% {
        -webkit-transform: translateY(0);
        -ms-transform: translateY(0);
        transform: translateY(0)
    }
}

.cc_container,.cc_message,.cc_btn {
    animation-duration: 0.8s;
    -webkit-animation-duration: 0.8s;
    -moz-animation-duration: 0.8s;
    -o-animation-duration: 0.8s;
    -webkit-animation-name: slideUp;
    animation-name: slideUp
}
`