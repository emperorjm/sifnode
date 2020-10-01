package networks

import (
	"encoding/json"
	"log"
	"time"

	"github.com/Sifchain/sifnode/tools/sifgen/networks/types"

	"github.com/sethvargo/go-password/password"
	"github.com/yelinaung/go-haikunator"
)

type Validator struct {
	name        string
	address     string
	peerAddress string
	keyPassword string
	genesisURL  string
	utils       Utils
}

func NewValidator(defaultNodeHome string) *Validator {
	keyPassword, _ := password.Generate(32, 5, 0, false, false)

	return &Validator{
		name:        haikunator.New(time.Now().UTC().UnixNano()).Haikunate(),
		keyPassword: keyPassword,
		utils:       NewUtils(defaultNodeHome),
	}
}

func (v *Validator) Name() string {
	return v.name
}

func (v *Validator) Address(address *string) *string {
	if address == nil {
		return &v.address
	}

	v.address = *address
	return nil
}

func (v *Validator) PeerAddress() string {
	return v.peerAddress
}

func (v *Validator) KeyPassword() string {
	return v.keyPassword
}

func (v *Validator) GenesisURL() string {
	return v.genesisURL
}

func (v *Validator) CollectPeerAddress() {
	var genesisAppState types.GenesisAppState
	if err := json.Unmarshal([]byte(v.utils.ExportGenesis()), &genesisAppState); err != nil {

		log.Fatal(err)
	}

	v.peerAddress = genesisAppState.AppState.Genutil.Gentxs[0].Value.Memo
}
