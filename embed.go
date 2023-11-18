// GENERATED FILE

package alan

import (
	"embed"
	"github.com/amitybell/piper-asset"
)

var (
	//go:embed dist.tzst dist.json MODEL_CARD.txt
	fs embed.FS

	Asset = asset.Asset{Name: "alan", FS: fs}
)
