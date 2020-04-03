package types_test

import (
	"bytes"
	"time"

	"github.com/tendermint/tendermint/crypto/tmhash"
	tmtypes "github.com/tendermint/tendermint/types"

	clientexported "github.com/cosmos/cosmos-sdk/x/ibc/02-client/exported"
	ibctmtypes "github.com/cosmos/cosmos-sdk/x/ibc/07-tendermint/types"
)

func (suite *TendermintTestSuite) TestEvidence() {
	signers := []tmtypes.PrivValidator{suite.privVal}

	ev := ibctmtypes.Evidence{
		Header1:  suite.header,
		Header2:  ibctmtypes.CreateTestHeader(chainID, height, suite.now, suite.tmValSet, signers),
		ChainID:  chainID,
		ClientID: "gaiamainnet",
	}

	suite.Require().Equal(ev.ClientType(), clientexported.Tendermint)
	suite.Require().Equal(ev.GetClientID(), "gaiamainnet")
	suite.Require().Equal(ev.Route(), "client")
	suite.Require().Equal(ev.Type(), "client_misbehaviour")
	suite.Require().Equal(ev.GetHeight(), int64(height))
}

func (suite *TendermintTestSuite) TestEvidenceValidateBasic() {
	altPrivVal := tmtypes.NewMockPV()
	altVal := tmtypes.NewValidator(altPrivVal.GetPubKey(), height)

	// Create bothValSet with both suite validator and altVal
	bothValSet := tmtypes.NewValidatorSet(append(suite.tmValSet.Validators, altVal))
	// Create alternative validator set with only altVal
	altValSet := tmtypes.NewValidatorSet([]*tmtypes.Validator{altVal})

	signers := []tmtypes.PrivValidator{suite.privVal}
	// Create signer array and ensure it is in same order as bothValSet
	var bothSigners []tmtypes.PrivValidator
	if bytes.Compare(altPrivVal.GetPubKey().Address(), suite.privVal.GetPubKey().Address()) == -1 {
		bothSigners = []tmtypes.PrivValidator{altPrivVal, suite.privVal}
	} else {
		bothSigners = []tmtypes.PrivValidator{suite.privVal, altPrivVal}
	}

	altSigners := []tmtypes.PrivValidator{altPrivVal}

	testCases := []struct {
		name             string
		evidence         ibctmtypes.Evidence
		malleateEvidence func(ev *ibctmtypes.Evidence) error
		expPass          bool
	}{
		{
			"valid evidence",
			ibctmtypes.Evidence{
				Header1:  suite.header,
				Header2:  ibctmtypes.CreateTestHeader(chainID, height, suite.now.Add(time.Minute), suite.tmValSet, signers),
				ChainID:  chainID,
				ClientID: "gaiamainnet",
			},
			func(ev *ibctmtypes.Evidence) error { return nil },
			true,
		},
		{
			"invalid client ID ",
			ibctmtypes.Evidence{
				Header1:  suite.header,
				Header2:  ibctmtypes.CreateTestHeader(chainID, height, suite.now, suite.tmValSet, signers),
				ChainID:  chainID,
				ClientID: "GAIA",
			},
			func(ev *ibctmtypes.Evidence) error { return nil },
			false,
		},
		{
			"wrong chainID on header1",
			ibctmtypes.Evidence{
				Header1:  suite.header,
				Header2:  ibctmtypes.CreateTestHeader("ethermint", height, suite.now, suite.tmValSet, signers),
				ChainID:  "ethermint",
				ClientID: "gaiamainnet",
			},
			func(ev *ibctmtypes.Evidence) error { return nil },
			false,
		},
		{
			"wrong chainID on header2",
			ibctmtypes.Evidence{
				Header1:  suite.header,
				Header2:  ibctmtypes.CreateTestHeader("ethermint", height, suite.now, suite.tmValSet, signers),
				ChainID:  chainID,
				ClientID: "gaiamainnet",
			},
			func(ev *ibctmtypes.Evidence) error { return nil },
			false,
		},
		{
			"mismatched heights",
			ibctmtypes.Evidence{
				Header1:  suite.header,
				Header2:  ibctmtypes.CreateTestHeader(chainID, 6, suite.now, suite.tmValSet, signers),
				ChainID:  chainID,
				ClientID: "gaiamainnet",
			},
			func(ev *ibctmtypes.Evidence) error { return nil },
			false,
		},
		{
			"same block id",
			ibctmtypes.Evidence{
				Header1:  suite.header,
				Header2:  suite.header,
				ChainID:  chainID,
				ClientID: "gaiamainnet",
			},
			func(ev *ibctmtypes.Evidence) error { return nil },
			false,
		},
		{
			"header 1 doesn't have 2/3 majority",
			ibctmtypes.Evidence{
				Header1:  ibctmtypes.CreateTestHeader(chainID, height, suite.now, bothValSet, bothSigners),
				Header2:  suite.header,
				ChainID:  chainID,
				ClientID: "gaiamainnet",
			},
			func(ev *ibctmtypes.Evidence) error {
				// voteSet contains only altVal which is less than 2/3 of total power (height/1height)
				wrongVoteSet := tmtypes.NewVoteSet(chainID, ev.Header1.SignedHeader.Header.Height, 1, tmtypes.PrecommitType, altValSet)
				commit, err := tmtypes.MakeCommit(ev.Header1.SignedHeader.Commit.BlockID.ToTmTypes(), ev.Header2.SignedHeader.Header.Height, ev.Header1.SignedHeader.Commit.ToTmTypes().Round, wrongVoteSet, altSigners, suite.now)
				if err != nil {
					return err
				}
				ev.Header1.SignedHeader.Commit = ibctmtypes.CommitFromTmTypes(commit)
				return nil
			},
			false,
		},
		{
			"header 2 doesn't have 2/3 majority",
			ibctmtypes.Evidence{
				Header1:  suite.header,
				Header2:  ibctmtypes.CreateTestHeader(chainID, height, suite.now, bothValSet, bothSigners),
				ChainID:  chainID,
				ClientID: "gaiamainnet",
			},
			func(ev *ibctmtypes.Evidence) error {
				// voteSet contains only altVal which is less than 2/3 of total power (height/1height)
				wrongVoteSet := tmtypes.NewVoteSet(chainID, ev.Header2.SignedHeader.Header.Height, 1, tmtypes.PrecommitType, altValSet)
				commit, err := tmtypes.MakeCommit(ev.Header2.SignedHeader.Commit.BlockID.ToTmTypes(), ev.Header2.SignedHeader.Header.Height, ev.Header2.SignedHeader.Commit.ToTmTypes().Round, wrongVoteSet, altSigners, suite.now)
				if err != nil {
					return err
				}
				ev.Header2.SignedHeader.Commit = ibctmtypes.CommitFromTmTypes(commit)
				return nil
			},
			false,
		},
		{
			"validators sign off on wrong commit",
			ibctmtypes.Evidence{
				Header1:  suite.header,
				Header2:  ibctmtypes.CreateTestHeader(chainID, height, suite.now, bothValSet, bothSigners),
				ChainID:  chainID,
				ClientID: "gaiamainnet",
			},
			func(ev *ibctmtypes.Evidence) error {
				blockID := ibctmtypes.MakeBlockID(tmhash.Sum([]byte("other_hash")), 3, tmhash.Sum([]byte("other_partset")))
				ev.Header2.SignedHeader.Commit.BlockID = ibctmtypes.BlockIDFromTmTypes(blockID)
				return nil
			},
			false,
		},
	}

	for i, tc := range testCases {
		tc := tc

		err := tc.malleateEvidence(&tc.evidence)
		suite.Require().NoError(err)

		if tc.expPass {
			suite.Require().NoError(tc.evidence.ValidateBasic(), "valid test case %d failed: %s", i, tc.name)
		} else {
			suite.Require().Error(tc.evidence.ValidateBasic(), "invalid test case %d passed: %s", i, tc.name)
		}
	}
}