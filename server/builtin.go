package server

import (
	"github.com/onlylayer/onlylayer/consensus"
	consensusDev "github.com/onlylayer/onlylayer/consensus/dev"
	consensusDummy "github.com/onlylayer/onlylayer/consensus/dummy"
	consensusIBFT "github.com/onlylayer/onlylayer/consensus/ibft"
	"github.com/onlylayer/onlylayer/secrets"
	"github.com/onlylayer/onlylayer/secrets/awsssm"
	"github.com/onlylayer/onlylayer/secrets/gcpssm"
	"github.com/onlylayer/onlylayer/secrets/hashicorpvault"
	"github.com/onlylayer/onlylayer/secrets/local"
)

type ConsensusType string

const (
	DevConsensus   ConsensusType = "dev"
	IBFTConsensus  ConsensusType = "ibft"
	DummyConsensus ConsensusType = "dummy"
)

var consensusBackends = map[ConsensusType]consensus.Factory{
	DevConsensus:   consensusDev.Factory,
	IBFTConsensus:  consensusIBFT.Factory,
	DummyConsensus: consensusDummy.Factory,
}

// secretsManagerBackends defines the SecretManager factories for different
// secret management solutions
var secretsManagerBackends = map[secrets.SecretsManagerType]secrets.SecretsManagerFactory{
	secrets.Local:          local.SecretsManagerFactory,
	secrets.HashicorpVault: hashicorpvault.SecretsManagerFactory,
	secrets.AWSSSM:         awsssm.SecretsManagerFactory,
	secrets.GCPSSM:         gcpssm.SecretsManagerFactory,
}

func ConsensusSupported(value string) bool {
	_, ok := consensusBackends[ConsensusType(value)]

	return ok
}
