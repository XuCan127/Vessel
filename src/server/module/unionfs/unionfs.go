package unionfs

import "Vessel/src/common/term"

func CreateUnionFS(containerId string) {
	aufs := term.CheckFS("aufs")

	overlay := term.CheckFS("overlay")
}
