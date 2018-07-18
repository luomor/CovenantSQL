/*
 * Copyright 2018 The ThunderDB Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package rpc

import (
	"gitlab.com/thunderdb/ThunderDB/conf"
	"gitlab.com/thunderdb/ThunderDB/crypto/asymmetric"
	"gitlab.com/thunderdb/ThunderDB/crypto/kms"
	"gitlab.com/thunderdb/ThunderDB/proto"
	"gitlab.com/thunderdb/ThunderDB/route"
	"gitlab.com/thunderdb/ThunderDB/utils/log"
)

// GetSharedSecretWith gets shared symmetric key with ECDH
func GetSharedSecretWith(nodeID *proto.RawNodeID, isAnonymous bool) (symmetricKey []byte, err error) {
	if isAnonymous {
		symmetricKey = []byte(`!&\\!qEyey*\cbLc,aKl`)
		log.Debug("using anonymous ETLS")
	} else {
		//log.Debugf("ECDH for %v and %v", localPrivateKey, nodePublicKey)
		var remotePublicKey *asymmetric.PublicKey
		if route.IsBPNodeID(nodeID) {
			remotePublicKey = kms.BP.PublicKey
		} else if conf.RoleTag[0] == 'B' {
			remotePublicKey, err = kms.GetPublicKey(proto.NodeID(nodeID.String()))
			if err != nil {
				log.Errorf("get public key locally failed, node id: %s, err: %s", nodeID.ToNodeID(), err)
				return
			}
		} else {
			// if non BP running and key not found, ask BlockProducer
			var nodeInfo *proto.Node
			nodeInfo, err = GetNodeInfo(nodeID)
			if err != nil {
				log.Errorf("get public key failed, node id: %s, err: %s", nodeID.ToNodeID(), err)
				return
			}
			remotePublicKey = nodeInfo.PublicKey
		}

		var localPrivateKey *asymmetric.PrivateKey
		localPrivateKey, err = kms.GetLocalPrivateKey()
		if err != nil {
			log.Errorf("get local private key failed: %s", err)
			return
		}
		symmetricKey = asymmetric.GenECDHSharedSecret(localPrivateKey, remotePublicKey)
	}
	return
}