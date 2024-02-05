# Etch

Etch is a html generator that aims to be simple and reasonably fast.

If you need to render some html but you don't want to learn yet another
templating language and its idiosyncharies. Etch uses plain old go functions
to define and render html. That means that everything is (reasonably) type-safe
and there's no additional compilations step.

## Use

```shell
go get -u https://github.com/getwhitespace/etch
```

## Usage

```golang
import . "https://github.com/getwhitespace/etch"

func LoginForm() Component {
    return Form(Class("settings-form"),
        Label(
            Text("Email:"),
            Input(Attr("type", "email"))
        ),
        Label(
            Text("Password:"),
            Input(Attr("type", "password"))
        ),
    )
}

str, err := RenderString(LoginForm())

fmt.Println(str)

// <form class="settings-form">
//   <label>Email:<input type="email" /></label>
//   <label>Password:<input type="password" /></label>
// </form>
```

Since everything is just a function, splitting up a component into smaller components is easy. This snippet renders the same html has the previous one.

```golang
func FormInput(label, typ string) Component {
    return Label(Text(label, ":"), Input(Attr("type", typ))),
}


func LoginForm() Component {
    return Form(Class("settings-form"),
        FormInput("Email", "email"),
        FormInput("Password", "password"),
    )
}

```

## Design decisions

The api to build components should make the resulting go code look as close to the html it is representing as possible. This means that the api must be minimal.

For example, this creates a label with a text node and an input node as children.

```golang
Label(
    Text("Email:"),
    Input(Attr("type", "email")),
)

```

This is fairly close to the html it is producing.

```html
<label>
  Email:
  <input type="email" />
</label>
```

This api enables Etch to make a single pass through the component tree to render the html.

To make the single pass render work there's a limitation that all attributes must be put before nodes in the argument list.

```golang

// ✅ OK
Label(Attr("class", "input-label"), Text("Email:"))

// ❌ Not OK
Label(Text("Email:"), Attr("class", "input-label"))

// an error will be returned and the attribute will be ignored
```

## Component

### HTML

See: [nodes.go](nodes.go)

Every element in https://developer.mozilla.org/en-US/docs/Web/HTML/Element

### SVG

See: [svg/nodes.go](svg/nodes.go)

Every element in https://developer.mozilla.org/en-US/docs/Web/SVG/Element

### Helpers

```golang

// render multiple attributes
func Attrs(values ...any) Component

// render `x` if `cond` == true otherwise render `y`
func If[T any](cond bool, x, y T) T

// render `x`` if `cond` == true otherwise nothing
func When[T any](cond bool, x T) *T

// iterate over a slice and render the return value of `fn`
func For[T any](items []T, fn func(item T) Component) Component

// iterate over a slice with indices and render the return value of `fn`
func ForIndex[T any](items []T, fn func(i int) Component) Component

// iterate `count` times
func Times(count int, fn func(i int) Component) Component

// sugar for Attr("class", values)
func Class(values ...any) Component

// sugar for Attr("id", values)
func Id(values ...any) Component

```

## Benchmark

Some numbers to get a rough idea about the performance. Benchmark is taken from https://github.com/slinso/goTemplateBenchmark.

| Simple html            |             |                 |              |                  |
| ---------------------- | ----------- | --------------- | ------------ | ---------------- |
| BenchmarkGolang-10     | 1596957     | 2237 ns/op      | 768 B/op     | 35 allocs/op     |
| BenchmarkGolangText-10 | 5242645     | 689.1 ns/op     | 128 B/op     | 7 allocs/op      |
| BenchmarkAce-10        | 903170      | 3939 ns/op      | 1120 B/op    | 40 allocs/op     |
| BenchmarkAmber-10      | 1546660     | 2307 ns/op      | 848 B/op     | 36 allocs/op     |
| BenchmarkMustache-10   | 2795353     | 1291 ns/op      | 1721 B/op    | 30 allocs/op     |
| BenchmarkPongo2-10     | 2072678     | 1740 ns/op      | 2073 B/op    | 32 allocs/op     |
| BenchmarkHandlebars-10 | 915727      | 3799 ns/op      | 3643 B/op    | 78 allocs/op     |
| BenchmarkSoy-10        | 3295437     | 1084 ns/op      | 1224 B/op    | 19 allocs/op     |
| BenchmarkJetHTML-10    | 8357434     | 432.3 ns/op     | 0 B/op       | 0 allocs/op      |
| **BenchmarkEtch-10**   | **3727230** | **965.4 ns/op** | **896 B/op** | **23 allocs/op** |

| Complex html                  |            |                |               |                   |
| ----------------------------- | ---------- | -------------- | ------------- | ----------------- |
| BenchmarkComplexGolang-10     | 174906     | 20419 ns/op    | 6540 B/op     | 290 allocs/op     |
| BenchmarkComplexGolangText-10 | 405541     | 8637 ns/op     | 2233 B/op     | 107 allocs/op     |
| BenchmarkComplexMustache-10   | 451724     | 7957 ns/op     | 7264 B/op     | 156 allocs/op     |
| BenchmarkComplexJetHTML-10    | 846660     | 4145 ns/op     | 530 B/op      | 5 allocs/op       |
| **BenchmarkComplexEtch-10**   | **636388** | **5686 ns/op** | **6204 B/op** | **124 allocs/op** |

## Roadmap-ish

- Optionally indent output
- Add more relevant helpers
- Add partial rendering. Useful for frontend libraries like htmx.
