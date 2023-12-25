package day25

import (
	"adventOfCode/utils"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
)

type wire struct {
	Name           string
	ConnectedNames []string
	Connected      []*wire
	Visited        bool
}

func Part1() {
	count := [2]int{}
	scanner := utils.ReadFile("day25")
	wires := map[string]*wire{}

	compile, err := regexp.Compile(`[ :]+`)
	if err != nil {
		return
	}

	for i := 0; scanner.Scan(); i++ {
		s := compile.Split(scanner.Text(), -1)
		wires[s[0]] = &wire{
			Name:           s[0],
			ConnectedNames: s[1:],
			Connected:      nil,
		}
	}

	for _, w := range wires {
		for _, name := range w.ConnectedNames {
			if v, ok := wires[name]; ok {
				if !slices.Contains(w.Connected, v) {
					w.Connected = append(w.Connected, v)
				}
				if !slices.Contains(v.Connected, w) {
					v.Connected = append(v.Connected, w)
				}
				if !slices.Contains(w.ConnectedNames, v.Name) {
					w.ConnectedNames = append(w.ConnectedNames, v.Name)
				}
				if !slices.Contains(v.ConnectedNames, w.Name) {
					v.ConnectedNames = append(v.ConnectedNames, w.Name)
				}
			} else {
				tmp := &wire{
					Name:           name,
					ConnectedNames: []string{w.Name},
					Connected:      []*wire{w},
				}
				wires[name] = tmp
				if !slices.Contains(w.Connected, tmp) {
					w.Connected = append(w.Connected, tmp)
					w.ConnectedNames = append(w.ConnectedNames, tmp.Name)
				}
				if !slices.Contains(w.ConnectedNames, tmp.Name) {
					w.ConnectedNames = append(w.ConnectedNames, tmp.Name)
				}
			}
		}
	}

	if true {
		err = os.Mkdir("solution/", 0750)
		if err != nil {
			log.Println("An Error Occurred:", err)
		}

		for _, w := range wires {
			created, err := os.Create("solution/" + w.Name + ".md")
			if err != nil {
				log.Fatalln("An Error Occurred1:", err)
			}

			for _, name := range w.ConnectedNames {
				_, err = created.Write([]byte(fmt.Sprintf("[[%s.md]]", name)))
				if err != nil {
					log.Fatalln("An Error Occurred2:", err)
				}
			}
		}
	} else {
		//open in obsidian, find what to delete:
		deleteable := [][2]string{
			{"pmn", "kdc"},
			{"grd", "hvm"},
			{"jmn", "zfk"},
		}

		for _, s := range deleteable {
			w1 := wires[s[0]]
			w2 := wires[s[1]]

			i := slices.Index(w1.ConnectedNames, s[1])
			w1.ConnectedNames = append(w1.ConnectedNames[:i], w1.ConnectedNames[i+1:]...)
			i = slices.Index(w1.Connected, w2)
			w1.Connected = append(w1.Connected[:i], w1.Connected[i+1:]...)

			i = slices.Index(w2.ConnectedNames, s[0])
			w2.ConnectedNames = append(w2.ConnectedNames[:i], w2.ConnectedNames[i+1:]...)
			i = slices.Index(w2.Connected, w1)
			w2.Connected = append(w2.Connected[:i], w2.Connected[i+1:]...)
		}

		rang := [2]utils.Stack[wire]{
			utils.Stack[wire]{wires[deleteable[0][0]]},
			utils.Stack[wire]{wires[deleteable[0][1]]},
		}

		for i := range rang {
			for !rang[i].IsEmpty() {
				u := rang[i].Pop()
				if !u.Visited {
					count[i]++
					u.Visited = true
					for _, w := range u.Connected {
						rang[i].Push(w)
					}
				}
			}
		}
	}

	fmt.Println(count[0] * count[1])
}
