package main

import (
	"flag"
	"os"
	"text/template"
)

type data struct {
	Type string
	Name string
}

func main() {
	var d data
	flag.StringVar(&d.Type, "type", "", "The subtype used for the queue being generated")
	flag.StringVar(&d.Name, "name", "", "The name used for the queue being generated. This should start with a capital letter so that is exported")
	flag.Parse()

	t := template.Must(template.New("queue").Parse(queueTemplate))
	t.Execute(os.Stdout, d)
}

var queueTemplate = `
package queue

import (
	"container/list"
)

func New{{.Name}}() *{{.Name}} {
	return &{{.Name}}{list.New()}
}

type {{.Name}} struct {
	list *list.List
}

func (q *{{.Name}}) Len() int {
	return q.list.len()
}

func (q *{{.Name}}) Enqueue(i {{.Type}}) {
	q.list.PushBack(i)
}

func (q *{{.Name}}) Dequeue() {{.Type}} {
	if q.list.Len() == 0 {
		panic(ErrEmptyQueue)
	}
	raw := q.list.Remove(q.list.Front())
	if typed, ok := raw.({{.Type}}); ok {
		return typed
	}
	panic(ErrInvalidType)
}
`
