package gift

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type Gift struct {
	Description string
	Cents       int
}

func (i *Gift) String() string {
	return fmt.Sprintf("%s %d", i.Description, i.Cents)
}

type GiftList []*Gift

func (l GiftList) Len() int {
	return len(l)
}
func (l GiftList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}
func (l GiftList) Less(i, j int) bool {
	// sort ascending
	return l[i].Cents < l[j].Cents
}

// Optimize returns the two gifts whose sum is closest to the given limit without going over
// returns (nil, nil) if no suitable combination exists
// modifies (sorts) this instance in place
func (l GiftList) Optimize(limit int) (cheap *Gift, expensive *Gift) {
	if limit < 1 {
		// no free stuff
		return nil, nil
	}

	if l.Len() < 2 {
		// we only optimize for exactly gifts
		return nil, nil
	}

	// sort first, O(n*log(n))
	sort.Sort(l)

	var maxSumSoFar int
	var leftIndex int
	rightIndex := l.Len() - 1

	// scan second, O(n)
	// prices are ordered ascending, move 1 pointer at a time, and don't let them cross
	for leftIndex < rightIndex {
		sum := l[leftIndex].Cents + l[rightIndex].Cents

		if sum > limit {
			// combination is too expensive, try replacing our more expensive Gift with a slightly less expensive one
			rightIndex--
			continue
		}

		if sum > maxSumSoFar {
			// combination is the best one so far...
			maxSumSoFar = sum
			cheap = l[leftIndex]
			expensive = l[rightIndex]
		}

		// ...but try replacing our cheaper Gift with a slightly more expensive one
		leftIndex++
	}

	// if we found any combination
	if maxSumSoFar > 0 {
		return cheap, expensive
	}

	return nil, nil
}

// NewGiftList parses the given unordered lines in the price sheet to create a new GiftList
// each line must be a description followed by a comma then a nonzero price in cents
// whitespace will be stripped
// ex.: {"taffy, 11", "chocolate, 390"}
func NewGiftList(lines []string) (GiftList, error) {
	list := make(GiftList, len(lines))
	for index, line := range lines {
		tokens := strings.Split(line, ",")
		description := strings.TrimSpace(tokens[0])
		if len(description) < 1 {
			return nil, errors.Errorf("a description is required for %s", line)
		}

		cents, err := strconv.Atoi(strings.TrimSpace(tokens[1]))
		if err != nil {
			return nil, errors.Wrapf(err, "invalid line: %s", line)
		}
		if cents < 1 {
			// free items not supported
			return nil, errors.Errorf("invalid price for %s; must be > 0", list[index].Description)
		}

		list[index] = &Gift{description, cents}
	}

	return list, nil
}
