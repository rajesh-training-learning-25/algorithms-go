package tree

import (
	"testing"
)

func TestAVLInsert(t *testing.T) {
	t.Run("LLRotation-Test", func(t *testing.T) {
		tree := NewAVLTree[int]()
		tree.Insert(5, 4, 3)

		root := tree.Root
		if root.Key != 4 {
			t.Errorf("Root should have value = 4, not %v", root.Key)
		}
		if root.Height != 2 {
			t.Errorf("Height of Root should be = 2, not %d", root.Height)
		}

		if root.Left.Key != 3 {
			t.Errorf("Left child should have value = 3")
		}
		if root.Left.Height != 1 {
			t.Errorf("Height of Left child should be 1")
		}

		if root.Right.Key != 5 {
			t.Errorf("Right child should have value = 5")
		}
		if root.Right.Height != 1 {
			t.Errorf("Height of Right should be 1")
		}

	})
	t.Run("LRRotation-Test", func(t *testing.T) {
		tree := NewAVLTree[int]()
		tree.Insert(5, 3, 4)

		root := tree.Root
		if root.Key != 4 {
			t.Errorf("Root should have value = 4")
		}
		if root.Height != 2 {
			t.Errorf("Height of Root should be = 2")
		}

		if root.Left.Key != 3 {
			t.Errorf("Left child should have value = 3")
		}
		if root.Left.Height != 1 {
			t.Errorf("Height of Left child should be 1")
		}

		if root.Right.Key != 5 {
			t.Errorf("Right child should have value = 5")
		}
		if root.Right.Height != 1 {
			t.Errorf("Height of Right should be 1")
		}
	})

	t.Run("RRRotation-Test", func(t *testing.T) {
		tree := NewAVLTree[int]()
		tree.Insert(3)
		tree.Insert(4)
		tree.Insert(5)

		root := tree.Root
		if root.Key != 4 {
			t.Errorf("Root should have value = 4")
		}
		if root.Height != 2 {
			t.Errorf("Height of Root should be = 2")
		}

		if root.Left.Key != 3 {
			t.Errorf("Left child should have value = 3")
		}
		if root.Left.Height != 1 {
			t.Errorf("Height of Left child should be 1")
		}

		if root.Right.Key != 5 {
			t.Errorf("Right child should have value = 5")
		}
		if root.Right.Height != 1 {
			t.Errorf("Height of Right should be 1")
		}
	})
	t.Run("RLRotaion-Test", func(t *testing.T) {
		tree := NewAVLTree[int]()
		tree.Insert(3)
		tree.Insert(5)
		tree.Insert(4)

		root := tree.Root
		if root.Key != 4 {
			t.Errorf("Root should have value = 4")
		}
		if root.Height != 2 {
			t.Errorf("Height of Root should be = 2")
		}

		if root.Left.Key != 3 {
			t.Errorf("Left child should have value = 3")
		}
		if root.Left.Height != 1 {
			t.Errorf("Height of Left child should be 1")
		}

		if root.Right.Key != 5 {
			t.Errorf("Right child should have value = 5")
		}
		if root.Right.Height != 1 {
			t.Errorf("Height of Right should be 1")
		}
	})
}

func TestAVLDelete(t *testing.T) {
	t.Run("LLRotaion-Test", func(t *testing.T) {
		tree := NewAVLTree[int]()

		tree.Insert(5)
		tree.Insert(4)
		tree.Insert(3)
		tree.Insert(2)

		tree.Delete(5)

		root := tree.Root
		if root.Key != 3 {
			t.Errorf("Root should have value = 3")
		}
		if root.Height != 2 {
			t.Errorf("Height of Root should be = 2")
		}

		if root.Left.Key != 2 {
			t.Errorf("Left child should have value = 2")
		}
		if root.Left.Height != 1 {
			t.Errorf("Height of Left child should be 1")
		}

		if root.Right.Key != 4 {
			t.Errorf("Right child should have value = 5")
		}
		if root.Right.Height != 1 {
			t.Errorf("Height of Right should be 1")
		}
	})

	t.Run("LRRotaion-Test", func(t *testing.T) {
		tree := NewAVLTree[int]()

		tree.Insert(10)
		tree.Insert(8)
		tree.Insert(6)
		tree.Insert(7)

		tree.Delete(10)

		root := tree.Root
		if root.Key != 7 {
			t.Errorf("Root should have value = 7")
		}
		if root.Height != 2 {
			t.Errorf("Height of Root should be = 2")
		}

		if root.Left.Key != 6 {
			t.Errorf("Left child should have value = 6")
		}
		if root.Left.Height != 1 {
			t.Errorf("Height of Left child should be 1")
		}

		if root.Right.Key != 8 {
			t.Errorf("Right child should have value = 8")
		}
		if root.Right.Height != 1 {
			t.Errorf("Height of Right should be 1")
		}

	})

	t.Run("RRRotaion-Test", func(t *testing.T) {
		tree := NewAVLTree[int]()

		tree.Insert(2)
		tree.Insert(3)
		tree.Insert(4)
		tree.Insert(5)

		tree.Delete(2)

		root := tree.Root
		if root.Key != 4 {
			t.Errorf("Root should have value = 4")
		}
		if root.Height != 2 {
			t.Errorf("Height of Root should be = 2")
		}

		if root.Left.Key != 3 {
			t.Errorf("Left child should have value = 3")
		}
		if root.Left.Height != 1 {
			t.Errorf("Height of Left child should be 1")
		}

		if root.Right.Key != 5 {
			t.Errorf("Right child should have value = 5")
		}
		if root.Right.Height != 1 {
			t.Errorf("Height of Right should be 1")
		}

	})

	t.Run("RLRotaion-Test", func(t *testing.T) {
		tree := NewAVLTree[int]()

		tree.Insert(7)
		tree.Insert(6)
		tree.Insert(9)
		tree.Insert(8)

		tree.Delete(6)

		root := tree.Root
		if root.Key != 8 {
			t.Errorf("Root should have value = 8")
		}
		if root.Height != 2 {
			t.Errorf("Height of Root should be = 2")
		}

		if root.Left.Key != 7 {
			t.Errorf("Left child should have value = 7")
		}
		if root.Left.Height != 1 {
			t.Errorf("Height of Left child should be 1")
		}

		if root.Right.Key != 9 {
			t.Errorf("Right child should have value = 9")
		}
		if root.Right.Height != 1 {
			t.Errorf("Height of Right should be 1")
		}

	})
}
