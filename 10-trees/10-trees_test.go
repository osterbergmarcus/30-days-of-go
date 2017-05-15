package Trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// create a memory pointer to a new instance
var tree = &Tree{
	&Node{
		value:    nil,
		children: make([]Node, 0),
	},
}

/*
 *create a tree structure as illustrated
 *         O
 *        / \
 *       O   O
 *      /\    \
 *     O  O    O
 */

func TestAdd(t *testing.T) {
	tree.Add("SkyBlue", nil)
	tree.Add("MidnightBlue", "SkyBlue")
	tree.Add("PowderBlue", "SkyBlue")
	tree.Add("LawnGreen", "PowderBlue")
	tree.Add("HoneyDew", "PowderBlue")
	tree.Add("ForestGreen", "HoneyDew")

	colorCodes := &Tree{
		&Node{
			"SkyBlue",
			[]Node{
				{
					"MidnightBlue",
					[]Node{},
				},
				{
					"PowderBlue",
					[]Node{
						{
							"LawnGreen",
							[]Node{},
						},
						{
							"HoneyDew",
							[]Node{
								{
									"ForestGreen",
									[]Node{},
								},
							},
						},
					},
				},
			},
		},
	}

	assert.Equal(t, *tree.rootNode, *colorCodes.rootNode)
}
