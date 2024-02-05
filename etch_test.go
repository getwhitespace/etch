package etch_test

import (
	"testing"

	. "github.com/getwhitespace/etch"
	svg "github.com/getwhitespace/etch/svg"
)

func TestPlainDiv(t *testing.T) {
	c := Div()
	html, err := RenderString(c)
	if err != nil {
		t.Fatalf("got err %s", err)
	}

	expected := `<div></div>`
	if html != expected {
		t.Fatalf("%s != %s", html, expected)
	}
}

func TestDivMultipleChildren(t *testing.T) {
	c := Div(Div(), Span())
	html, err := RenderString(c)
	if err != nil {
		t.Fatalf("got err %s", err)
	}

	expected := `<div><div></div><span></span></div>`
	if html != expected {
		t.Fatalf("%s != %s", html, expected)
	}
}

func TestFragmentRendersChildren(t *testing.T) {
	c := Div(Div(), Span(), B())
	c2 := Div(Fragment(Div(), Span()), B())
	htmlc, err := RenderString(c)
	if err != nil {
		t.Fatalf("got err %s", err)
	}
	htmlc2, err := RenderString(c2)
	if err != nil {
		t.Fatalf("got err %s", err)
	}
	if htmlc != htmlc2 {
		t.Fatalf("%s != %s", htmlc, htmlc2)
	}
}

func TestDivWithAttributes(t *testing.T) {
	c := Div(Id("id1"), Class("bg-black"), Attr("href", "http://google.com"), Attr("data-url", "http://facebook.com"))
	html, err := RenderString(c)
	if err != nil {
		t.Fatalf("got err %s", err)
	}

	expected := `<div id="id1" class="bg-black" href="http://google.com" data-url="http://facebook.com"></div>`
	if html != expected {
		t.Fatalf("%s != %s", html, expected)
	}
}

func TestDivWithBooleanAttributes(t *testing.T) {
	c := Div(Attr("disabled"))
	html, err := RenderString(c)
	if err != nil {
		t.Fatalf("got err %s", err)
	}

	expected := `<div disabled></div>`
	if html != expected {
		t.Fatalf("%s != %s", html, expected)
	}
}

func TestDivWithSpecialCharAttributes(t *testing.T) {
	c := Div(Id(`"`), Class(`"klass"`), Attr("href", "<div>"), Attr("data-url", "===="))
	html, err := RenderString(c)
	if err != nil {
		t.Fatalf("got err %s", err)
	}

	expected := `<div id="&#34;" class="&#34;klass&#34;" href="&lt;div&gt;" data-url="===="></div>`
	if html != expected {
		t.Fatalf("%s != %s", html, expected)
	}
}

func TestVoidElements(t *testing.T) {
	c := Fragment(Area(), Base(), Br(), Col(), Embed(), Hr(), Img(), Input(), Link(), Meta(), Param(), Source(), Track(), Wbr())
	html, err := RenderString(c)
	if err != nil {
		t.Fatalf("got err %s", err)
	}

	expected := `<area /><base /><br /><col /><embed /><hr /><img /><input /><link /><meta /><param /><source /><track /><wbr />`
	if html != expected {
		t.Fatalf("%s != %s", html, expected)
	}
}

func TestVoidElementWithAttr(t *testing.T) {
	c := Img(Attr("src", "http://google.com"))
	html, err := RenderString(c)
	if err != nil {
		t.Fatalf("got err %s", err)
	}

	expected := `<img src="http://google.com" />`
	if html != expected {
		t.Fatalf("%s != %s", html, expected)
	}
}

func TestVoidElementChildren(t *testing.T) {
	c := Img(Attr("src", "http://google.com"), Text("hello"))
	html, err := RenderString(c)

	if err == nil {
		t.Fatalf("test should have errored but didn't")
	}

	expected := `<img src="http://google.com" />`
	if html != expected {
		t.Fatalf("%s != %s", html, expected)
	}
}

func TestRendersSvg(t *testing.T) {
	c := svg.Svg(
		Attr("viewBox", "0 0 100 100"),
		Attr("xmlns", "http://www.w3.org/2000/svg"),
		svg.G(Attr("fill", "white"), Attr("stroke", "green"), Attr("stroke-width", "5"),
			svg.Circle(Attr("cx", "40"), Attr("cy", "40"), Attr("r", "25")),
			svg.Circle(Attr("cx", "60"), Attr("cy", "60"), Attr("r", "25")),
		),
	)
	html, _ := RenderString(c)

	expected := `<svg viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg"><g fill="white" stroke="green" stroke-width="5"><circle cx="40" cy="40" r="25"></circle><circle cx="60" cy="60" r="25"></circle></g></svg>`
	if html != expected {
		t.Fatalf("%s != %s", html, expected)
	}
}

func BenchmarkEtch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := RenderBytes(
			Html(
				Body(
					H1(Text("Bob")),
					P(
						Text("Here's a list of your favorite colors:"),
					),
					Ul(
						Li(Text("blue")),
						Li(Text("green")),
						Li(Text("mauve")),
					),
				),
			),
		)
		if err != nil {
			b.Fatal(err)
		}
	}
}
