package etch

import (
	"fmt"
	"math"
)

func Node(tag string, args []Component) Component {
	return func(sb *Renderer) {
		sb.writeNode(tag, args)
	}
}

func Attr(name string, values ...any) Component {
	return func(r *Renderer) {
		r.writeAttr(name, values...)
	}
}

func Attrs(values ...any) Component {
	return func(r *Renderer) {
		var iters int = int(math.Floor(float64(len(values)) / 2.0))
		for i := 0; i < iters; i++ {
			start := i * 2
			r.writeAttr(anyToString(values[start]), values[start+1])
		}
	}
}

func AttrNS(namespace string, name string, values ...any) Component {
	return func(r *Renderer) {
		r.writeAttr(fmt.Sprintf("%s:%s", namespace, name), values...)
	}
}

func If[T any](cond bool, x, y T) T {
	if cond {
		return x
	} else {
		return y
	}
}

func When[T any](cond bool, x T) *T {
	if cond {
		return &x
	} else {
		return nil
	}
}

func For[T any](items []T, fn func(item T) Component) Component {
	return func(r *Renderer) {
		for i := range items {
			fn(items[i])(r)
		}
	}
}

func ForIndex[T any](items []T, fn func(i int) Component) Component {
	return func(r *Renderer) {
		for i := range items {
			fn(i)(r)
		}
	}
}

func Times(count int, fn func(i int) Component) Component {
	return func(r *Renderer) {
		for i := 0; i < count; i++ {
			fn(i)(r)
		}
	}
}

func Class(values ...any) Component {
	v := make([]any, 0, len(values)*2-1)
	for i := range values {
		if i > 0 {
			v = append(v, " ")
		}
		v = append(v, values[i])
	}
	return Attr("class", v...)
}

func Id(values ...any) Component {
	return Attr("id", values...)
}
