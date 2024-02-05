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

## Roadmap-ish

- Optionally indent output
- Add more relevant helpers
- Add partial rendering. Useful for frontend libraries like htmx.
