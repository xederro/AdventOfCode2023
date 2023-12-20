package day20

import (
	"adventOfCode/utils"
	"fmt"
	"regexp"
	"strings"
)

var (
	pulseQueue2 utils.Queue[pulse2]
	modules2    map[string]module2
)

type module2 interface {
	recive(*pulse2)
	connect()
}

type pulse2 struct {
	Src  string
	Dest string
	Type bool
}

type ff2 struct {
	On   bool
	Dest []string
	Self string
}

func (f *ff2) recive(p *pulse2) {
	if p.Type {
		return
	}

	f.On = !f.On
	for _, dest := range f.Dest {
		pulseQueue2.Enqueue(&pulse2{
			Dest: dest,
			Type: f.On,
			Src:  f.Self,
		})
	}
}

func (f *ff2) connect() {
	for _, dest := range f.Dest {
		if v, ok := modules2[dest].(*conj2); ok {
			v.Modules[f.Self] = false
		}
	}
}

type conj2 struct {
	Modules map[string]bool
	Dest    []string
	Self    string
}

func (f *conj2) recive(p *pulse2) {
	f.Modules[p.Src] = p.Type
	high := false
	for _, b := range f.Modules {
		if !b {
			high = true
			break
		}
	}
	for _, dest := range f.Dest {
		pulseQueue2.Enqueue(&pulse2{
			Dest: dest,
			Src:  f.Self,
			Type: high,
		})
	}
}

func (f *conj2) connect() {
	for _, dest := range f.Dest {
		if v, ok := modules2[dest].(*conj2); ok {
			v.Modules[f.Self] = false
		}
	}
}

type broad2 struct {
	Dest []string
	Self string
}

func (f *broad2) recive(p *pulse2) {
	for _, dest := range f.Dest {
		pulseQueue2.Enqueue(&pulse2{
			Dest: dest,
			Src:  f.Self,
			Type: p.Type,
		})
	}
}

func (f *broad2) connect() {
	for _, s := range f.Dest {
		if v, ok := modules2[s].(*conj2); ok {
			v.Modules[f.Self] = false
		}
	}
}

type out2 struct {
}

func (f *out2) recive(p *pulse2) {
}

func (f *out2) connect() {

}

func Part1() {
	modules2 = map[string]module2{}
	pulseQueue2 = utils.Queue[pulse2]{}
	scanner := utils.ReadFile("day20")
	countH := 0
	countL := 0
	modules2["output"] = &out2{}

	compile, err := regexp.Compile(`[, ]+`)
	if err != nil {
		return
	}
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " -> ")
		dests := compile.Split(s[1], -1)

		switch s[0][0] {
		case 'b':
			modules2[s[0]] = &broad2{
				Dest: dests,
				Self: s[0],
			}
			break
		case '%':
			modules2[s[0][1:]] = &ff2{
				Dest: dests,
				Self: s[0][1:],
			}
			break
		case '&':
			modules2[s[0][1:]] = &conj2{
				Modules: map[string]bool{},
				Dest:    dests,
				Self:    s[0][1:],
			}
			break
		}
	}

	for _, m := range modules2 {
		m.connect()
	}

	for i := 0; i < 1000; i++ {
		pulseQueue2.Enqueue(&pulse2{
			Dest: "broadcaster",
			Src:  "btn",
			Type: false,
		})
		for !pulseQueue2.IsEmpty() {
			p := pulseQueue2.Dequeue()
			if p.Type {
				countH++
			} else {
				countL++
			}

			if v, ok := modules2[p.Dest]; ok {
				v.recive(p)
			}
		}
	}

	fmt.Println(countH, countL, countH*countL)
}
