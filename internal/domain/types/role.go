package types

import (
	"encoding/json"
)

type Role byte

const (
	Admin               Role = 0b00000100
	Modifier            Role = 0b00000010
	Watcher             Role = 0b00000001
	AdminModifier       Role = Admin | Modifier
	AdminModifierWacher Role = Admin | Modifier | Watcher
	ModifierWatcher     Role = Modifier | Watcher
)

func FromStringRole(stringRole StringRole) Role {
	switch stringRole.String() {
	case admin:
		return Admin
	case modifier:
		return Modifier
	default: // watcher it's the lower role that any user could be
		return Watcher
	}
}

func (r Role) Update(new Role) Role {
	return r | new
}

func (r Role) MarshalJSON() ([]byte, error) {
	roles := []string{}

	if r&Admin == Admin {
		roles = append(roles, "ADMIN")
	}

	if r&Modifier == Modifier {
		roles = append(roles, "MODIFIER")
	}

	if r&Watcher == Watcher {
		roles = append(roles, "WATCHER")
	}

	return json.Marshal(roles)
}
