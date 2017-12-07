package tower

import (
	"math"
)

type Program struct {
	Name        string
	Weight      float64
	Children    []string
	TotalWeight float64
}

type Tower struct {
	Programs map[string]*Program
	Bases    []string
}

func (t *Tower) FindBase() string {
Loop:
	for _, v := range t.Bases {
		for _, x := range t.Programs {
			// Ignore self or leaf nodes
			if x.Name == v || x.Children == nil {
				continue
			}
			for _, each := range x.Children {
				if v == each {
					continue Loop
				}
			}
		}
		return v
	}
	return ""
}

func (t *Tower) totalWeight(root string) float64 {
	if t.Programs[root].Children != nil {
		for _, v := range t.Programs[root].Children {
			t.Programs[root].TotalWeight += +t.totalWeight(v)
		}
		t.Programs[root].TotalWeight += t.Programs[root].Weight
	} else {
		t.Programs[root].TotalWeight += t.Programs[root].Weight
	}
	return t.Programs[root].TotalWeight
}

func (t *Tower) FindUnbalanced(root string) float64 {
	_ = t.totalWeight(root)
	var unbalanced string
Loop:
	for _, program := range t.Programs {
		programs := make(map[float64]int)
		// Leaf node
		if program.Children == nil || program.Name == root {
			continue
		}
		// Not leaf, for each child entry, hash weights
		// If has has more then one entry, its unbalanced
		// The culprit will have balanced child entries,
		// so we need to find that next and return.
		for _, child := range program.Children {
			programs[t.Programs[child].TotalWeight]++
			if len(programs) > 1 {
				balanced := make(map[float64]int)
				for _, v := range t.Programs[child].Children {
					balanced[t.Programs[v].TotalWeight]++
				}
				if len(balanced) == 1 {
					unbalanced = child
					break Loop
				}
				continue Loop
			}
		}
	}

	parent := t.getParent(unbalanced)
	return t.weightAdjustment(unbalanced, parent)
}

func (t *Tower) getParent(node string) string {
	for _, v := range t.Programs {
		if v.Children == nil {
			continue
		}
		for _, program := range v.Children {
			if program == node {
				return v.Name
			}
		}
	}
	return ""
}

func (t *Tower) weightAdjustment(unbalanced, parent string) float64 {
	var unbalancedWeight float64
	var parentWeight float64
	for _, v := range t.Programs[parent].Children {
		if v == unbalanced {
			continue
		}
		unbalancedWeight = math.Abs(t.Programs[v].TotalWeight - t.Programs[unbalanced].TotalWeight)
		parentWeight = t.Programs[unbalanced].TotalWeight
		break
	}
	return parentWeight - unbalancedWeight
}
