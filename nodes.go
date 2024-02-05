package etch

func Text(args ...any) Component {
	return func(sb *Renderer) {
		for i := range args {
			sb.writeText(Escape(args[i]))
		}
	}
}

func Raw(args ...any) Component {
	return func(sb *Renderer) {
		for i := range args {
			sb.writeText(anyToString(args[i]))
		}
	}
}

func Fragment(args ...Component) Component {
	return func(sb *Renderer) {
		for i := range args {
			args[i](sb)
		}
	}
}

// generated html nodes
func A(args ...Component) Component {
	return Node("a", args)
}

func Abbr(args ...Component) Component {
	return Node("abbr", args)
}

func Acronym(args ...Component) Component {
	return Node("acronym", args)
}

func Address(args ...Component) Component {
	return Node("address", args)
}

func Area(args ...Component) Component {
	return Node("area", args)
}

func Article(args ...Component) Component {
	return Node("article", args)
}

func Aside(args ...Component) Component {
	return Node("aside", args)
}

func Audio(args ...Component) Component {
	return Node("audio", args)
}

func B(args ...Component) Component {
	return Node("b", args)
}

func Base(args ...Component) Component {
	return Node("base", args)
}

func Bdi(args ...Component) Component {
	return Node("bdi", args)
}

func Bdo(args ...Component) Component {
	return Node("bdo", args)
}

func Big(args ...Component) Component {
	return Node("big", args)
}

func Blockquote(args ...Component) Component {
	return Node("blockquote", args)
}

func Body(args ...Component) Component {
	return Node("body", args)
}

func Br(args ...Component) Component {
	return Node("br", args)
}

func Button(args ...Component) Component {
	return Node("button", args)
}

func Canvas(args ...Component) Component {
	return Node("canvas", args)
}

func Caption(args ...Component) Component {
	return Node("caption", args)
}

func Center(args ...Component) Component {
	return Node("center", args)
}

func Cite(args ...Component) Component {
	return Node("cite", args)
}

func Code(args ...Component) Component {
	return Node("code", args)
}

func Col(args ...Component) Component {
	return Node("col", args)
}

func Colgroup(args ...Component) Component {
	return Node("colgroup", args)
}

func Data(args ...Component) Component {
	return Node("data", args)
}

func Datalist(args ...Component) Component {
	return Node("datalist", args)
}

func Dd(args ...Component) Component {
	return Node("dd", args)
}

func Del(args ...Component) Component {
	return Node("del", args)
}

func Details(args ...Component) Component {
	return Node("details", args)
}

func Dfn(args ...Component) Component {
	return Node("dfn", args)
}

func Dialog(args ...Component) Component {
	return Node("dialog", args)
}

func Dir(args ...Component) Component {
	return Node("dir", args)
}

func Div(args ...Component) Component {
	return Node("div", args)
}

func Dl(args ...Component) Component {
	return Node("dl", args)
}

func Dt(args ...Component) Component {
	return Node("dt", args)
}

func Em(args ...Component) Component {
	return Node("em", args)
}

func Embed(args ...Component) Component {
	return Node("embed", args)
}

func Fieldset(args ...Component) Component {
	return Node("fieldset", args)
}

func Figcaption(args ...Component) Component {
	return Node("figcaption", args)
}

func Figure(args ...Component) Component {
	return Node("figure", args)
}

func Font(args ...Component) Component {
	return Node("font", args)
}

func Footer(args ...Component) Component {
	return Node("footer", args)
}

func Form(args ...Component) Component {
	return Node("form", args)
}

func Frame(args ...Component) Component {
	return Node("frame", args)
}

func Frameset(args ...Component) Component {
	return Node("frameset", args)
}

func H1(args ...Component) Component {
	return Node("h1", args)
}

func H2(args ...Component) Component {
	return Node("h2", args)
}

func H3(args ...Component) Component {
	return Node("h3", args)
}

func H4(args ...Component) Component {
	return Node("h4", args)
}

func H5(args ...Component) Component {
	return Node("h5", args)
}

func H6(args ...Component) Component {
	return Node("h6", args)
}

func Head(args ...Component) Component {
	return Node("head", args)
}

func Header(args ...Component) Component {
	return Node("header", args)
}

func Hgroup(args ...Component) Component {
	return Node("hgroup", args)
}

func Hr(args ...Component) Component {
	return Node("hr", args)
}

func Html(args ...Component) Component {
	return Node("html", args)
}

func I(args ...Component) Component {
	return Node("i", args)
}

func Iframe(args ...Component) Component {
	return Node("iframe", args)
}

func Image(args ...Component) Component {
	return Node("image", args)
}

func Img(args ...Component) Component {
	return Node("img", args)
}

func Input(args ...Component) Component {
	return Node("input", args)
}

func Ins(args ...Component) Component {
	return Node("ins", args)
}

func Kbd(args ...Component) Component {
	return Node("kbd", args)
}

func Label(args ...Component) Component {
	return Node("label", args)
}

func Legend(args ...Component) Component {
	return Node("legend", args)
}

func Li(args ...Component) Component {
	return Node("li", args)
}

func Link(args ...Component) Component {
	return Node("link", args)
}

func Main(args ...Component) Component {
	return Node("main", args)
}

func Map(args ...Component) Component {
	return Node("map", args)
}

func Mark(args ...Component) Component {
	return Node("mark", args)
}

func Marquee(args ...Component) Component {
	return Node("marquee", args)
}

func Menu(args ...Component) Component {
	return Node("menu", args)
}

func Menuitem(args ...Component) Component {
	return Node("menuitem", args)
}

func Meta(args ...Component) Component {
	return Node("meta", args)
}

func Meter(args ...Component) Component {
	return Node("meter", args)
}

func Nav(args ...Component) Component {
	return Node("nav", args)
}

func Nobr(args ...Component) Component {
	return Node("nobr", args)
}

func Noembed(args ...Component) Component {
	return Node("noembed", args)
}

func Noframes(args ...Component) Component {
	return Node("noframes", args)
}

func Noscript(args ...Component) Component {
	return Node("noscript", args)
}

func Object(args ...Component) Component {
	return Node("object", args)
}

func Ol(args ...Component) Component {
	return Node("ol", args)
}

func Optgroup(args ...Component) Component {
	return Node("optgroup", args)
}

func Option(args ...Component) Component {
	return Node("option", args)
}

func Output(args ...Component) Component {
	return Node("output", args)
}

func P(args ...Component) Component {
	return Node("p", args)
}

func Param(args ...Component) Component {
	return Node("param", args)
}

func Picture(args ...Component) Component {
	return Node("picture", args)
}

func Plaintext(args ...Component) Component {
	return Node("plaintext", args)
}

func Portal(args ...Component) Component {
	return Node("portal", args)
}

func Pre(args ...Component) Component {
	return Node("pre", args)
}

func Progress(args ...Component) Component {
	return Node("progress", args)
}

func Q(args ...Component) Component {
	return Node("q", args)
}

func Rb(args ...Component) Component {
	return Node("rb", args)
}

func Rp(args ...Component) Component {
	return Node("rp", args)
}

func Rt(args ...Component) Component {
	return Node("rt", args)
}

func Rtc(args ...Component) Component {
	return Node("rtc", args)
}

func Ruby(args ...Component) Component {
	return Node("ruby", args)
}

func S(args ...Component) Component {
	return Node("s", args)
}

func Samp(args ...Component) Component {
	return Node("samp", args)
}

func Script(args ...Component) Component {
	return Node("script", args)
}

func Search(args ...Component) Component {
	return Node("search", args)
}

func Section(args ...Component) Component {
	return Node("section", args)
}

func Select(args ...Component) Component {
	return Node("select", args)
}

func Slot(args ...Component) Component {
	return Node("slot", args)
}

func Small(args ...Component) Component {
	return Node("small", args)
}

func Source(args ...Component) Component {
	return Node("source", args)
}

func Span(args ...Component) Component {
	return Node("span", args)
}

func Strike(args ...Component) Component {
	return Node("strike", args)
}

func Strong(args ...Component) Component {
	return Node("strong", args)
}

func Style(args ...Component) Component {
	return Node("style", args)
}

func Sub(args ...Component) Component {
	return Node("sub", args)
}

func Summary(args ...Component) Component {
	return Node("summary", args)
}

func Sup(args ...Component) Component {
	return Node("sup", args)
}

func Table(args ...Component) Component {
	return Node("table", args)
}

func Tbody(args ...Component) Component {
	return Node("tbody", args)
}

func Td(args ...Component) Component {
	return Node("td", args)
}

func Template(args ...Component) Component {
	return Node("template", args)
}

func Textarea(args ...Component) Component {
	return Node("textarea", args)
}

func Tfoot(args ...Component) Component {
	return Node("tfoot", args)
}

func Th(args ...Component) Component {
	return Node("th", args)
}

func Thead(args ...Component) Component {
	return Node("thead", args)
}

func Time(args ...Component) Component {
	return Node("time", args)
}

func Title(args ...Component) Component {
	return Node("title", args)
}

func Tr(args ...Component) Component {
	return Node("tr", args)
}

func Track(args ...Component) Component {
	return Node("track", args)
}

func Tt(args ...Component) Component {
	return Node("tt", args)
}

func U(args ...Component) Component {
	return Node("u", args)
}

func Ul(args ...Component) Component {
	return Node("ul", args)
}

func Var(args ...Component) Component {
	return Node("var", args)
}

func Video(args ...Component) Component {
	return Node("video", args)
}

func Wbr(args ...Component) Component {
	return Node("wbr", args)
}

func Xmp(args ...Component) Component {
	return Node("xmp", args)
}
