package items

import (
	"fmt"

	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/player"
)

// PlayerHandler ...
type PlayerHandler struct {
	player.NopHandler
	p *player.Player
}

func (*PlayerHandler) Name() string { return "Item Handler" }

// NewPlayerHandler returns a new *PlayerHandler
func NewPlayerHandler(p *player.Player) *PlayerHandler { return &PlayerHandler{p: p} }

// HandleItemUse will handle when an item has been used in the air.
// It makes sure that the item held is compatible (registered) and that it is usable.
func (h *PlayerHandler) HandleItemUse(*event.Context) {
	p := h.p                     // Handled player
	heldItem, _ := p.HeldItems() // Main hand
	fmt.Println("Held Item: ", heldItem.CustomName())

	if i, ok := Compatible(heldItem); ok { //Making sure it is registered
		fmt.Println("Item Is Compatible")
		if usable, ok := i.(UsableItem); ok { // Hopefully usable
			fmt.Println("Item Is Usable")
			usable.Use(heldItem, p)
			fmt.Println("Used Item")
		}
	}
}
