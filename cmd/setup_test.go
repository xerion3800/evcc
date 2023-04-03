package cmd

import (
	"strings"
	"testing"

	"github.com/xerion3800/evcc/api"
	"github.com/xerion3800/evcc/core"
	"github.com/xerion3800/evcc/util"
	"github.com/spf13/viper"
)

const sample = `
loadpoints:
- mode: off
`

func TestYamlOff(t *testing.T) {
	var conf config
	viper.SetConfigType("yaml")
	if err := viper.ReadConfig(strings.NewReader(sample)); err != nil {
		t.Error(err)
	}

	if err := viper.UnmarshalExact(&conf); err != nil {
		t.Error(err)
	}

	var lp core.Loadpoint
	if err := util.DecodeOther(conf.Loadpoints[0], &lp); err != nil {
		t.Error(err)
	}

	if lp.Mode != api.ModeOff {
		t.Errorf("expected `off`, got %s", lp.Mode)
	}
}
