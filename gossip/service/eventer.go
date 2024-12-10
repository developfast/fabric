/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package service

import (
	"github.com/hyperledger/fabric-protos-go-apiv2/peer"
	"github.com/hyperledger/fabric/common/channelconfig"
	"github.com/hyperledger/fabric/msp"
	"google.golang.org/protobuf/proto"
)

// A ConfigUpdate holds the portions of channelconfig Bundle update used by
// gossip.
type ConfigUpdate struct {
	ChannelID        string
	Organizations    map[string]channelconfig.ApplicationOrg
	Sequence         uint64
	OrdererAddresses []string
}

// ConfigProcessor receives config updates
type ConfigProcessor interface {
	// ProcessConfigUpdate should be invoked whenever a channel's configuration is initialized or updated
	ProcessConfigUpdate(configUpdate ConfigUpdate)
}

type configStore struct {
	anchorPeers []*peer.AnchorPeer
	orgMap      map[string]channelconfig.ApplicationOrg
}

type configEventReceiver interface {
	updateAnchors(config ConfigUpdate)
}

type configEventer struct {
	lastConfig *configStore
	receiver   configEventReceiver
}

func newConfigEventer(receiver configEventReceiver) *configEventer {
	return &configEventer{
		receiver: receiver,
	}
}

// ProcessConfigUpdate should be invoked whenever a channel's configuration is initialized or updated
// it invokes the associated method in configEventReceiver when configuration is updated
// but only if the configuration value actually changed
// Note, that a changing sequence number is ignored as changing configuration
func (ce *configEventer) ProcessConfigUpdate(configUpdate ConfigUpdate) {
	logger.Debugf("Processing new config for channel %s", configUpdate.ChannelID)
	orgMap := cloneOrgConfig(configUpdate.Organizations)
	if ce.lastConfig != nil && orgMapEqual(ce.lastConfig.orgMap, orgMap) {
		logger.Debugf("Ignoring new config for channel %s because it contained no anchor peer updates", configUpdate.ChannelID)
	} else {

		var newAnchorPeers []*peer.AnchorPeer
		for _, group := range configUpdate.Organizations {
			newAnchorPeers = append(newAnchorPeers, group.AnchorPeers()...)
		}

		newConfig := &configStore{
			orgMap:      orgMap,
			anchorPeers: newAnchorPeers,
		}
		ce.lastConfig = newConfig

		logger.Debugf("Calling out because config was updated for channel %s", configUpdate.ChannelID)
		ce.receiver.updateAnchors(configUpdate)
	}
}

func cloneOrgConfig(src map[string]channelconfig.ApplicationOrg) map[string]channelconfig.ApplicationOrg {
	clone := make(map[string]channelconfig.ApplicationOrg)
	for k, v := range src {
		clone[k] = &appGrp{
			name:        v.Name(),
			mspID:       v.MSPID(),
			anchorPeers: v.AnchorPeers(),
		}
	}
	return clone
}

func orgMapEqual(orgs1 map[string]channelconfig.ApplicationOrg, orgs2 map[string]channelconfig.ApplicationOrg) bool {
	if len(orgs1) != len(orgs2) {
		return false
	}

	for k, org1 := range orgs1 {
		org2, ok := orgs2[k]
		if !ok {
			return false
		}
		if org1.Name() != org2.Name() {
			return false
		}
		if org1.MSPID() != org2.MSPID() {
			return false
		}
		aps1 := org1.AnchorPeers()
		aps2 := org2.AnchorPeers()
		if len(aps1) != len(aps2) {
			return false
		}
		verified := make(map[int]bool, len(aps1))
	external:
		for _, ap1 := range aps1 {
			for i, ap2 := range aps2 {
				if !verified[i] && proto.Equal(ap1, ap2) {
					verified[i] = true
					continue external
				}
			}
			return false
		}
	}

	return true
}

type appGrp struct {
	name        string
	mspID       string
	anchorPeers []*peer.AnchorPeer
}

func (ag *appGrp) Name() string {
	return ag.name
}

func (ag *appGrp) MSPID() string {
	return ag.mspID
}

func (ag *appGrp) AnchorPeers() []*peer.AnchorPeer {
	return ag.anchorPeers
}

func (ag *appGrp) MSP() msp.MSP {
	return nil
}
