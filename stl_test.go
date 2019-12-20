package stl

import (
	"fmt"
	"os"
	"testing"

	"github.com/o5h/glm"
)

func TestSolid_ReadFrom(t *testing.T) {
	stlFile, err := os.Open("testdata/cube.stl")
	if err != nil {
		t.Fatal(err)
	}
	defer stlFile.Close()

	s := new(Solid)
	s.ReadFrom(stlFile)

	s.WriteTo(os.Stdout)

}

func Test_sliceToVec3(t *testing.T) {
	v := new(glm.Vec3)
	sliceToVec3([]string{"1", "2", "3"}, v)
	fmt.Println(v)
}
