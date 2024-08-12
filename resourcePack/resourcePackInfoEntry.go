package resourcePack

type ResourcePackInfoEntry struct {
	PackID string
	Version string
	PackSize int64
}

func NewResourcePackInfoEntry(packid string, version string, size int64) *ResourcePackInfoEntry{
	return &ResourcePackInfoEntry{packid, version, size}
}

func (rpie *ResourcePackInfoEntry) GetPackID() string {
	return rpie.PackID
}

func (rpie *ResourcePackInfoEntry) GetVersion() string {
	return rpie.Version
}

func (rpie *ResourcePackInfoEntry) GetSize() int64 {
	return rpie.PackSize
}
