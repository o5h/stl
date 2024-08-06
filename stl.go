package stl

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Vec3 [3]float32

type Solid struct {
	Name   string
	Facets []Facet
}

func New(name string, numFaces int) *Solid {
	return &Solid{name, make([]Facet, 0, numFaces)}
}

func (s *Solid) AddFacet(f Facet) {
	s.Facets = append(s.Facets, f)
}

type Facet struct {
	Normal   Vec3
	Vertices []Vec3
}

func NewFacet(n Vec3) *Facet {
	return &Facet{n, make([]Vec3, 0, 3)}
}

func (f *Facet) AddVec3(v Vec3) {
	f.Vertices = append(f.Vertices, v)
}

func (s *Solid) ReadFrom(ior io.Reader) (n int64, err error) {
	r := bufio.NewReader(ior)
	line, _, _ := r.ReadLine()
	hl := strings.Fields(string(line))

	s.Name = hl[1]
	var facet *Facet
	for {
		line, _, err = r.ReadLine()
		if err == io.EOF {
			break
		}
		fs := strings.Fields(string(line))
		switch fs[0] {
		case "facet":
			facet = &Facet{}
			sliceToVec3(fs[2:5], &facet.Normal)
		case "outer":
		case "vertex":
			v := Vec3{}
			sliceToVec3(fs[1:4], &v)
			facet.AddVec3(v)
		case "endloop":
		case "endfacet":
			s.AddFacet(*facet)
		case "endsolid":
			if s.Name != hl[1] {
				err = fmt.Errorf("expected  %v. was %v", s.Name, hl[1])
			}
			return
		default:
			err = fmt.Errorf("unknown token %s", fs[0])
		}

		if err != nil {
			return
		}
	}
	return
}

func parseFloat32(s string) float32 {
	f, _ := strconv.ParseFloat(s, 32)
	return float32(f)
}

func sliceToVec3(fs []string, v *Vec3) {
	v[0] = parseFloat32(fs[0])
	v[1] = parseFloat32(fs[1])
	v[2] = parseFloat32(fs[2])
}

func (s *Solid) WriteTo(w io.Writer) (n int64, err error) {
	fmt.Fprintf(w, "solid %v\n", s.Name)
	for _, f := range s.Facets {
		fmt.Fprintf(w, "  facet normal %.6f %.6f %.6f\n", f.Normal[0], f.Normal[1], f.Normal[2])
		fmt.Fprintln(w, "    outer loop")
		for _, v := range f.Vertices {
			fmt.Fprintf(w, "      vertex %.6f %.6f %.6f\n", v[0], v[1], v[2])

		}
		fmt.Fprintln(w, "    endloop")
		fmt.Fprintln(w, "  endfacet")
	}
	fmt.Fprintf(w, "endsolid %v\n", s.Name)
	return
}
