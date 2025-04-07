package tasks

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/drizzleent/pluto/x/tasks/keeper"
)

var ModuleName = "tasks"

type AppModule struct {
	keeper *keeper.Keeper
}

func NewAppModule(k *keeper.Keeper) *AppModule {
	return &AppModule{
		keeper: k,
	}
}

func (am *AppModule) Name() string {
	return ModuleName
}
func (am *AppModule) QuerierRoute() string {
	return ModuleName
}

func (am *AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

func (am *AppModule) Route() {

}
