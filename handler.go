package items

import (
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/player"
)

// PlayerHandler ...
type PlayerHandler struct {
	player.NopHandler
	p *player.Player
}

// NewPlayerHandler returns a new *PlayerHandler
func NewPlayerHandler(p *player.Player) *PlayerHandler { return &PlayerHandler{p: p} }

// HandleItemUse will handle when an item has been used in the air.
// It makes sure that the item held is compatible (registered) and that it is usable.
func (h *PlayerHandler) HandleItemUse(*event.Context) {
	p := h.p                     // Handled player
	heldItem, _ := p.HeldItems() // Main hand

	if i, ok := Compatible(heldItem); ok { //Making sure it is registered
		if usable, ok := i.(UsableItem); ok { // Hopefully usable
			usable.Use(heldItem, p)
		}
	}
}
