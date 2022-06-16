package plugin

import (
	"github.com/meson-network/peer-node/basic"
)

//todo: ---
func InitPlugin() {

	/////////////////////////
	err := initReference()
	if err != nil {
		basic.Logger.Fatalln(err)
	}

	/////////////////////////
	err = initSqlite()
	if err != nil {
		basic.Logger.Fatalln(err)
	}

}
