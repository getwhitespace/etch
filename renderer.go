package etch

import (
	"errors"
	"fmt"
	"html"
	"io"
	"slices"
)

type Component = func(sb *Renderer)

type position = string

const (
	positionTag     position = "tag"
	positionContent position = "content"
	positionEnd     position = "end"
)

var voidElements = []string{
	"area", "base", "br", "col", "embed", "hr", "img", "input", "link", "meta", "param", "source", "track", "wbr",
}

type Renderer struct {
	bytes    []byte
	position position
	curTag   string
	err      error
}

func (r *Renderer) openStartTag(tag string) {
	if r.position == positionTag {
		r.closeStartTag()
	}
	r.write("<", tag)
	r.curTag = tag
	r.position = positionTag
}

func (r *Renderer) closeStartTag() {
	if slices.Contains(voidElements, r.curTag) {
		r.write(" />")
		r.position = positionEnd
	} else {
		r.write(">")
		r.position = positionContent
	}
}

func (r *Renderer) writeAttr(name string, values ...any) {
	if r.position != positionTag {
		r.err = errors.Join(r.err, fmt.Errorf("tried to write attr '%s', but the tag was already closed. attrs must be put before the nodes. attr '%s' is skipped", name, name))
		return
	}

	r.write(" ", html.EscapeString(name))

	if len(values) > 0 {
		r.write(`=`, `"`)
		for i := range values {
			r.write(Escape(values[i]))
		}
		r.write(`"`)
	}
}

func (r *Renderer) writeEndNode(tag string) {
	if r.position == positionTag {
		r.closeStartTag()
	}

	if r.position != positionEnd {
		r.write("</", tag, ">")
	}

	r.position = positionContent
	r.curTag = ""
}

func (r *Renderer) writeText(text string) {
	if r.position == positionTag {
		r.closeStartTag()
	}

	if r.position == positionEnd {
		r.err = errors.Join(r.err, fmt.Errorf("tried to write text '%s' to '%s', but '%s' is a void element that cannot contain text. text is skipped", text, r.curTag, r.curTag))
		return
	}
	r.write(text)
	r.position = positionContent
}

func (r *Renderer) writeNode(tag string, args []Component) {
	if r.position == positionEnd {
		r.err = errors.Join(r.err, fmt.Errorf("tried to write child node '%s' to '%s', but '%s' is a void element that cannot contain child nodes. child node '%s' is skipped", tag, r.curTag, tag, tag))
		return
	}

	r.openStartTag(tag)
	for i := range args {
		if args[i] != nil {
			args[i](r)
		}
	}
	r.writeEndNode(tag)
}

func (r *Renderer) write(s ...string) {
	for i := range s {
		r.bytes = append(r.bytes, s[i]...)
	}
}

func (r *Renderer) Render(c Component, w io.Writer) error {
	if bytes, err := r.RenderBytes(c); err != nil {
		return err
	} else if _, err := w.Write(bytes); err != nil {
		return err
	}
	return nil
}

func (r *Renderer) RenderString(c Component) (string, error) {
	bytes, err := r.RenderBytes(c)
	return string(bytes), err
}

func (r *Renderer) RenderBytes(c Component) ([]byte, error) {
	r.bytes = r.bytes[:0]
	c(r)
	return r.bytes, r.err
}

func NewRenderer() *Renderer {
	return new(Renderer)
}

func Render(c Component, w io.Writer) error {
	return NewRenderer().Render(c, w)
}

func RenderBytes(c Component) ([]byte, error) {
	return NewRenderer().RenderBytes(c)
}

func RenderString(c Component) (string, error) {
	return NewRenderer().RenderString(c)
}
