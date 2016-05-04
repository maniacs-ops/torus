package block

import (
	"errors"

	"github.com/coreos/agro"
	"github.com/coreos/agro/models"
	"github.com/coreos/pkg/capnslog"
)

var clog = capnslog.NewPackageLogger("github.com/coreos/agro", "block")

type blockMetadata interface {
	agro.MetadataService

	Lock(lease int64) error
	GetINode() (agro.INodeRef, error)
	SyncINode(agro.INodeRef) error
	Unlock() error
	CreateBlockVolume(vol *models.Volume) error
}

func createBlockMetadata(mds agro.MetadataService, vid agro.VolumeID) (blockMetadata, error) {
	switch mds.Kind() {
	case agro.EtcdMetadata:
		return createBlockEtcdMetadata(mds, vid)
	case agro.TempMetadata:
		return createBlockTempMetadata(mds, vid)
	default:
		return nil, errors.New("unimplemented for this kind of metadata")
	}
}