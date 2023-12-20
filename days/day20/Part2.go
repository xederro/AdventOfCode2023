package day20

import (
	"adventOfCode/utils"
	"fmt"
	"regexp"
	"strings"
)

var (
	pulseQueue utils.Queue[pulse]
	modules    map[string]module
	count      int
	end        []int64
)

type module interface {
	recive(*pulse)
	connect()
}

type pulse struct {
	Src  string
	Dest string
	Type bool
}

type ff struct {
	On   bool
	Dest []string
	Self string
}

func (f *ff) recive(p *pulse) {
	if p.Type {
		return
	}

	f.On = !f.On
	for _, dest := range f.Dest {
		pulseQueue.Enqueue(&pulse{
			Dest: dest,
			Type: f.On,
			Src:  f.Self,
		})
	}
}

func (f *ff) connect() {
	for _, dest := range f.Dest {
		if v, ok := modules[dest].(*conj); ok {
			v.Modules[f.Self] = false
		}
	}
}

type conj struct {
	Modules map[string]bool
	Dest    []string
	Self    string
}

func (f *conj) recive(p *pulse) {
	f.Modules[p.Src] = p.Type
	high := false
	for _, b := range f.Modules {
		if !b {
			high = true
			break
		}
	}
	for _, dest := range f.Dest {
		pulseQueue.Enqueue(&pulse{
			Dest: dest,
			Src:  f.Self,
			Type: high,
		})
	}
}

func (f *conj) connect() {
	for _, dest := range f.Dest {
		if v, ok := modules[dest].(*conj); ok {
			v.Modules[f.Self] = false
		}
	}
}

type broad struct {
	Dest []string
	Self string
}

func (f *broad) recive(p *pulse) {
	for _, dest := range f.Dest {
		pulseQueue.Enqueue(&pulse{
			Dest: dest,
			Src:  f.Self,
			Type: p.Type,
		})
	}
}

func (f *broad) connect() {
	for _, s := range f.Dest {
		if v, ok := modules[s].(*conj); ok {
			v.Modules[f.Self] = false
		}
	}
}

type out struct {
	Self string
}

func (f *out) recive(p *pulse) {
	if !p.Type {
		fmt.Println(f.Self, count)
		end = append(end, int64(count))
	}
}

func (f *out) connect() {

}

func Part2() {
	modules = map[string]module{}
	pulseQueue = utils.Queue[pulse]{}
	scanner := utils.ReadFile("day20")
	modules["output"] = &out{
		Self: "output",
	}
	modules["rx"] = &out{
		Self: "rx",
	}

	compile, err := regexp.Compile(`[, ]+`)
	if err != nil {
		return
	}
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " -> ")
		dests := compile.Split(s[1], -1)

		switch s[0][0] {
		case 'b':
			modules[s[0]] = &broad{
				Dest: dests,
				Self: s[0],
			}
			break
		case '%':
			modules[s[0][1:]] = &ff{
				Dest: dests,
				Self: s[0][1:],
			}
			break
		case '&':
			if dests[0] != "lv" {
				modules[s[0][1:]] = &conj{
					Modules: map[string]bool{},
					Dest:    dests,
					Self:    s[0][1:],
				}
			} else {
				modules[s[0][1:]] = &out{
					Self: s[0][1:],
				}
			}

			break
		}
	}

	for _, m := range modules {
		m.connect()
	}

	for len(end) < 4 {
		count++
		pulseQueue.Enqueue(&pulse{
			Dest: "broadcaster",
			Src:  "btn",
			Type: false,
		})
		for !pulseQueue.IsEmpty() {
			p := pulseQueue.Dequeue()

			if v, ok := modules[p.Dest]; ok {
				v.recive(p)
			}
		}
	}

	fmt.Println(end)

	fmt.Println(utils.LCM(end...))
}
