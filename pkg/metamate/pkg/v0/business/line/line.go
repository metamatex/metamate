package line

import (
	"fmt"
	"github.com/metamatex/metamatemono/pkg/metamate/pkg/v0/types"
	"strings"
	"sync"
)

const (
	Default = "default"
)

type RequireCtx struct {
	stage        bool
	ctx          bool
	gCliReq      bool
	gCliRsp      bool
	gSvcReq      bool
	gSvcRsp      bool
	gSvcRsps     bool
	gSvcFilter   bool
	endpointKind bool
	gSvcs        bool
	gSvc         bool
	svcErrs      bool
	g            bool
	gSlice       bool
	err          bool
}

type Line struct {
	err        *Line
	errReturn  bool
	isErrLine  bool
	name       string
	pipes      []Pipe
	parallel   *Line
	concurrent []*Line
	if0        *Line
	switch0    map[string]*Line
	next       *Line
	prev       *Line
	Log        bool
}

func New(name string) (*Line) {
	return &Line{name: name}
}

var c = 0

func getUnnamed() (string) {
	c++
	return fmt.Sprintf("unnamed-%v", c)
}

func Do(ts ...types.Transformer) (*Line) {
	l := &Line{name: getUnnamed()}

	return l.Do(ts...)
}

func ErrLine() (*Line) {
	return &Line{name: getUnnamed(), isErrLine:true}
}

func If(f func(ctx types.ReqCtx) (bool), l0 *Line) (*Line) {
	l := &Line{name: getUnnamed()}

	return l.If(f, l0)
}

func Parallel(c int, mapF func(ctx types.ReqCtx) (ctxs []types.ReqCtx), to *Line, reduceF func(ctx types.ReqCtx, ctxs []types.ReqCtx) (types.ReqCtx)) (*Line) {
	l := &Line{name: getUnnamed()}

	return l.Parallel(c, mapF, to, reduceF)
}

func Concurrent(to []*Line, reduceF func(ctx types.ReqCtx, ctxs []types.ReqCtx) (types.ReqCtx)) (*Line) {
	l := &Line{name: getUnnamed()}

	return l.Concurrent(to, reduceF)
}

func Switch(mapF func(ctx types.ReqCtx) (string), to map[string]*Line) (*Line) {
	l := &Line{name: getUnnamed()}

	return l.Switch(mapF, to)
}

func Error(errL *Line, return0 bool) (*Line) {
	l := &Line{name: getUnnamed()}

	return l.Error(errL, return0)
}

func (l *Line) NewCtx() (types.ReqCtx) {
	return types.ReqCtx{}
}

type Pipe struct {
	Name string
	Func func(types.ReqCtx) (types.ReqCtx)
}

func (l *Line) Do(ts ...types.Transformer) (*Line) {
	for _, t := range ts {
		l.pipes = append(l.pipes, Pipe{Name: t.Name(), Func: t.Transform})
	}

	return l
}

func (l *Line) Func(f func(types.ReqCtx) (types.ReqCtx)) (*Line) {
	l.pipes = append(l.pipes, Pipe{Name: "func", Func: f})

	return l
}

func (l *Line) Draw() (s string) {
	return l.draw(0)
}

func (l *Line) draw(indent int) (s string) {
	indentStr := strings.Repeat("\t", indent)

	for _, pipe := range l.pipes {
		if pipe.Name == "next" {
			continue
		}

		s += indentStr + pipe.Name + "\n"
	}

	for k, p0 := range l.switch0 {
		s += indentStr + "\t" + k + "\n"
		s += indentStr + getFirst(p0).draw(indent+2)
	}

	if l.parallel != nil {
		s += l.parallel.draw(indent + 2)
	}

	if l.if0 != nil {
		s += l.if0.draw(indent + 2)
	}

	if l.next != nil {
		s += l.next.draw(indent)
	}

	return
}

func getFirst(l *Line) (f *Line) {
	f = l
	for {
		if f.prev == nil {
			break
		}

		f = f.prev
	}

	return
}

func (l *Line) Name() (string) {
	return l.name
}

func (l *Line) Parallel(c int, mapF func(ctx types.ReqCtx) (ctxs []types.ReqCtx), to *Line, reduceF func(ctx types.ReqCtx, ctxs []types.ReqCtx) (types.ReqCtx)) (*Line) {
	l.parallel = getFirst(to)

	if c < 0 {
		f := func(ctx types.ReqCtx) (types.ReqCtx) {
			ctxs := mapF(ctx)

			w := sync.WaitGroup{}
			for i, _ := range ctxs {
				w.Add(1)

				go func(i int) {
					ctxs[i] = l.parallel.Transform(ctxs[i])
					w.Add(-1)
				}(i)
			}

			w.Wait()

			return reduceF(ctx, ctxs)
		}

		l.pipes = append(l.pipes, Pipe{Name: "parallel", Func: f})
	}

	l.next = l.getNext()

	l.pipes = append(l.pipes, Pipe{Name: "next", Func: l.next.Transform})

	return l.next
}

func (l *Line) Concurrent(to []*Line, reduceF func(ctx types.ReqCtx, ctxs []types.ReqCtx) (types.ReqCtx)) (*Line) {
	for i, _ := range to {
		to[i] = getFirst(to[i])
	}

	l.concurrent = to

	f := func(ctx types.ReqCtx) (types.ReqCtx) {
		ctxs := []types.ReqCtx{}

		w := sync.WaitGroup{}
		for i, _ := range to {
			w.Add(1)

			ctx0, err := ctx.Copy()
			if err != nil {
			    panic(err)
			}

			ctxs = append(ctxs, ctx0)

			go func(i int) {
				ctxs[i] = to[i].Transform(ctxs[i])
				w.Add(-1)
			}(i)
		}

		w.Wait()

		return reduceF(ctx, ctxs)
	}

	l.pipes = append(l.pipes, Pipe{Name: "concurrent", Func: f})

	l.next = l.getNext()

	l.pipes = append(l.pipes, Pipe{Name: "next", Func: l.next.Transform})

	return l.next
}

func (l *Line) Add(f func(*Line)(*Line)) (*Line) {
	return f(l)
}

func (l *Line) Switch(mapF func(ctx types.ReqCtx) (string), to map[string]*Line) (*Line) {
	for k, l0 := range to {
		if l0 == nil {
			continue
		}

		to[k] = getFirst(l0)
	}

	l.switch0 = to

	l.pipes = append(l.pipes, Pipe{Name: "switch0", Func: func(ctx types.ReqCtx) (types.ReqCtx) {
		k := mapF(ctx)

		l0, ok := l.switch0[k]
		if !ok {
			l0, _ = l.switch0[Default]
			if l0 == nil {
				return ctx
			}
		}

		return l0.Transform(ctx)
	}})

	l.next = l.getNext()

	l.pipes = append(l.pipes, Pipe{Name: "next", Func: l.next.Transform})

	return l.next
}

func (l *Line) SetLog(b bool) {
	l.Log = b

	for _, l0 := range l.switch0 {
		if l0 != nil {
			l0.SetLog(b)
		}
	}

	if l.parallel != nil {
		l.parallel.SetLog(b)
	}

	if l.if0 != nil {
		l.if0.SetLog(b)
	}

	if l.next != nil {
		l.next.SetLog(b)
	}
}

func (l *Line) getNext() (*Line) {
	return &Line{name: getUnnamed(), prev: l, Log: l.Log, errReturn: l.errReturn, err: l.err, isErrLine: l.isErrLine}
}

func (l *Line) Error(l0 *Line, return0 bool) (*Line) {
	if !l0.isErrLine {
		panic("must be err line, line.ErrLine()")
	}

	l.err = getFirst(l0)
	l.errReturn = return0

	return l
}

func (l *Line) If(f func(ctx types.ReqCtx) (bool), l0 *Line) (*Line) {
	l.if0 = getFirst(l0)

	l.pipes = append(l.pipes, Pipe{Name: "if", Func: func(ctx types.ReqCtx) (types.ReqCtx) {
		if f(ctx) {
			return l.if0.Transform(ctx)
		}

		return ctx
	}})

	l.next = l.getNext()

	l.pipes = append(l.pipes, Pipe{Name: "next", Func: l.next.Transform})

	return l.next
}

func (l *Line) Transform(ctx types.ReqCtx) (types.ReqCtx) {
	if l.Log {
		println("--- " + getFirst(l).Name())
	}

	if l.isErrLine {
		for _, n := range l.pipes {
			ctx = n.Func(ctx)
		}
	} else {}
	for _, n := range l.pipes {
		ctx = n.Func(ctx)

		if len(ctx.Errs) != 0 {
			if l.err == nil {
				return ctx
			} else {
				ctx := l.err.Transform(ctx)
				if l.errReturn {
					return ctx
				}
			}
		}
	}

	return ctx
}
