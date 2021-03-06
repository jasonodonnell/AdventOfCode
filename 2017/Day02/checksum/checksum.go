package checksum

import (
	"errors"
	"sort"
)

type Checksum struct {
	Nums []int
}

func (c *Checksum) Checksum() (int, error) {
	i, j, err := c.divides()
	if err != nil {
		return 0, err
	}
	return c.Nums[i] / c.Nums[j], nil
}

func (c *Checksum) Difference() int {
	return c.largest() - c.smallest()
}

func (c *Checksum) divides() (int, int, error) {
	sort.Sort(sort.Reverse(sort.IntSlice(c.Nums)))
	for i := 0; i < (len(c.Nums) - 2); i++ {
		for j := i + 1; j < len(c.Nums); j++ {
			if (c.Nums[i] % c.Nums[j]) == 0 {
				return i, j, nil
			}
		}
	}
	return 0, 0, errors.New("no dividing pairs found")
}

func (c *Checksum) largest() int {
	largest := -1
	for _, num := range c.Nums {
		if num > largest {
			largest = num
		}
	}
	return largest
}

func (c *Checksum) smallest() int {
	smallest := 1000000000
	for _, num := range c.Nums {
		if num < smallest {
			smallest = num
		}
	}
	return smallest
}
