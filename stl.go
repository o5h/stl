package stl

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/o5h/glm"
)

type Solid struct {
	Name   string
	Facets []Facet
}

func New(name string, numFaces int) (s *Solid) {
	s = &Solid{name, make([]Facet, 0, numFaces)}
	return
}

func (s *Solid) AddFacet(f Facet) {
	s.Facets = append(s.Facets, f)
}

type Facet struct {
	Normal   glm.Vec3
	Vertices []glm.Vec3
}

func NewFacet(n glm.Vec3) *Facet {
	return &Facet{n, make([]glm.Vec3, 0, 3)}
}

func (f *Facet) AddVec3(v glm.Vec3) {
	f.Vertices = append(f.Vertices, v)
}

func (s *Solid) ReadFrom(ior io.Reader) (n int64, err error) {
	r := bufio.NewReader(ior)
	line, _, _ := r.ReadLine()
	hl := strings.Fields(string(line))

	s.Name = hl[1]
	var facet *Facet
	for true {
		line, _, err = r.ReadLine()
		if err == io.EOF {
			break
		}
		fs := strings.Fields(string(line))
		switch fs[0] {
		case "facet":
			facet = new(Facet)
			err = sliceToVec3(fs[2:5], &facet.Normal)
		case "outer":
		case "vertex":
			v := &glm.Vec3{}
			err = sliceToVec3(fs[1:4], v)
			facet.AddVec3(*v)
		case "endloop":
		case "endfacet":
			s.AddFacet(*facet)
		default:
			err = fmt.Errorf("Unknown token %s", fs[0])
		}

		if err != nil {
			return
		}
	}
	return
}

func sliceToVec3(fs []string, v *glm.Vec3) (err error) {
	x, _ := strconv.ParseFloat(fs[0], 32)
	y, _ := strconv.ParseFloat(fs[1], 32)
	z, _ := strconv.ParseFloat(fs[2], 32)
	v.SetXYZ(float32(x), float32(y), float32(z))
	return
}

func (s *Solid) WriteTo(w io.Writer) (n int64, err error) {
	fmt.Fprintf(w, "solid %v\n", s.Name)
	for _, f := range s.Facets {
		fmt.Fprintf(w, "  facet normal %.1f %.1f %.1f\n", f.Normal.X, f.Normal.Y, f.Normal.Z)
		fmt.Fprintln(w, "    outer loop")
		for _, v := range f.Vertices {
			fmt.Fprintf(w, "      vertex %.1f %.1f %.1f\n", v.X, v.Y, v.Z)

		}
		fmt.Fprintln(w, "    endloop")
		fmt.Fprintln(w, "  endfacet")
	}
	fmt.Fprintf(w, "endsolid %v\n", s.Name)
	return
}
