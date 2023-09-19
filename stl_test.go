package stl

import (
	"os"
	"testing"

	"github.com/o5h/testing/assert"
)

func TestSolid_ReadFrom(t *testing.T) {
	stlFile, err := os.Open("testdata/bottle.stl")
	if err != nil {
		t.Fatal(err)
	}
	defer stlFile.Close()

	s := new(Solid)
	_, err = s.ReadFrom(stlFile)
	assert.Nil(t, err)

	stlFile2, _ := os.Create("testdata/bottle2.stl")
	s.WriteTo(stlFile2)
	stlFile2.Close()

}

func Test_sliceToVec3(t *testing.T) {
	v := Vec3{}
	sliceToVec3([]string{"1", "2", "3"}, &v)
	assert.Eq(t, v[0], float32(1))
	assert.Eq(t, v[1], float32(2))
	assert.Eq(t, v[2], float32(3))
}
