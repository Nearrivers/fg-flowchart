package filetree

import (
	"reflect"
	"testing"
)

func TestInsertNode(t *testing.T) {
	t.Run("Node insertion", func(t *testing.T) {
		testNode := Node{
			Name:  "Lab",
			Type:  DIR,
			Files: []*Node{},
		}

		wantedName := "nodetest"
		newNode := InsertNode(false, &testNode, wantedName)
		if newNode.Name != wantedName {
			t.Errorf("got %s want %s", wantedName, newNode.Name)
		}

		if newNode.Type != FILE {
			t.Errorf("")
		}
	})
}

func TestRemoveIndex(t *testing.T) {
	t.Run("Existing index deletion", func(t *testing.T) {
		nodes := []*Node{
			{
				Name: "t1",
				Type: FILE,
			},
			{
				Name: "t2",
				Type: FILE,
			},
			{
				Name: "t3",
				Type: FILE,
			},
		}

		want := []*Node{
			{
				Name: "t1",
				Type: FILE,
			},
			{
				Name: "t3",
				Type: FILE,
			},
		}

		got, err := removeIndex(nodes, 1)
		if err != nil {
			t.Errorf("An error occured %v", err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v want %+v", got, want)
		}
	})

	t.Run("Out of bounds index deletion", func(t *testing.T) {
		nodes := []*Node{
			{
				Name: "t1",
				Type: FILE,
			},
			{
				Name: "t2",
				Type: FILE,
			},
			{
				Name: "t3",
				Type: FILE,
			},
		}

		_, got := removeIndex(nodes, 99)
		if got == nil {
			t.Error("An occured should've occured")
		}

		want := ErrIndexOutOfBounds
		if got != ErrIndexOutOfBounds {
			t.Errorf("want %v got %v", want, got)
		}
	})
}
