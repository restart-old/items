package items

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"sync"
)

// Registered Items.
var items sync.Map

// Register registers the given UsableItem.
func Register(item ...UsableItem) {
	for _, i := range item {
		items.Store(i.Item(), i)
	}
}

// Named ...
type Named interface {
	Name() string
}

// Item ...
type Item interface {
	Item() world.Item
}

// Usable ...
type Usable interface {
	Use(stack item.Stack, p *player.Player)
}

// UsableItem is an interface in which there is a world.Item, and a Use() method.
type UsableItem interface {
	Item
	Usable
}

// NameCompatible returns if the stack passed has the right custom name.
// If the interface passed is not a Named, it returns true.
func NameCompatible(v interface{}, stack item.Stack) bool {
	if named, ok := v.(Named); ok {
		if named.Name() == stack.CustomName() {
			return true
		}
		return false
	}
	return true
}

// Compatible returns the custom item corresponding to the stack if valid.
// Returning false if the item of the stack is not registered.
func Compatible(stack item.Stack) (interface{}, bool) {
	if load, ok := items.Load(stack.Item()); ok {
		return load, NameCompatible(load, stack)
	}
	return nil, false
}
