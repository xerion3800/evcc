package meter

import (
	"github.com/xerion3800/evcc/api"
	"github.com/xerion3800/evcc/meter/tplink"
	"github.com/xerion3800/evcc/util"
)

func init() {
	registry.Add("tplink", NewTPLinkFromConfig)
}

// NewTPLinkFromConfig creates a tapo meter from generic config
func NewTPLinkFromConfig(other map[string]interface{}) (api.Meter, error) {
	var cc struct {
		URI string
	}

	if err := util.DecodeOther(other, &cc); err != nil {
		return nil, err
	}

	return tplink.NewConnection(cc.URI)
}
