package dataPacket

type StartGamePacket struct {
	EntityUniqueID          int32   `json:"EntityUniqueID"`
	EntityRuntimeID         int32   `json:"EntityRuntimeID"`
	PlayerGamemode          int32   `json:"PlayerGamemode"`
	X                       float32 `json:"X"`
	Y                       float32 `json:"Y"`
	Z                       float32 `json:"Z"`
	Pitch                   float32 `json:"Pitch"`
	Yaw                     float32 `json:"Yaw"`
	Seed                    int32   `json:"Seed"`
	Dimension               int32   `json:"Dimension"`
	Generator               int32   `json:"Generator"`
	WorldGamemode           int32   `json:"WorldGamemode"`
	Difficulty              int32   `json:"Difficulty"`
	SpawnX                  int32   `json:"SpawnX"`
	SpawnY                  uint32  `json:"SpawnY"`
	SpawnZ                  int32   `json:"SpawnZ"`
	HasAchievementsDisabled bool    `json:"HasAchievementsDisabled"`
	DayCycleStopTime        int32   `json:"DayCycleStopTime"`
	EduMode                 bool    `json:"EduMode"`
	RainLevel               float32 `json:"RainLevel"`
	LightningLevel          float32 `json:"LightningLevel"`
	CommandsEnabled         bool    `json:"CommandsEnabled"`
	IsTexturePacksRequired  bool    `json:"IsTexturePacksRequired"`
	Gamerules               uint32  `json:"Gamerules"`
	LevelID                 string  `json:"LevelID"`
	WorldName               string  `json:"WorldName"`
	PremiumWorldTemplateID  string  `json:"PremiumWorldTemplateID"`
	PacketName              string  `json:"PacketName"`
}

// ID ...
func (*StartGamePacket) ID() byte {
	return IDStartGamePacket
}

// Marshal ...
func (pk *StartGamePacket) Marshal(w *PacketWriter) {
	pk.PacketName = getName(pk)
	w.Varint32(&pk.EntityRuntimeID)
	w.Varint32(&pk.EntityUniqueID)
	w.Varint32(&pk.PlayerGamemode)
	w.Float32(&pk.X)
	w.Float32(&pk.Y)
	w.Float32(&pk.Z)
	w.Float32(&pk.Pitch)
	w.Float32(&pk.Yaw)
	w.Varint32(&pk.Seed)
	w.Varint32(&pk.Dimension)
	w.Varint32(&pk.Generator)
	w.Varint32(&pk.WorldGamemode)
	w.Varint32(&pk.Difficulty)
	w.Varint32(&pk.SpawnX)
	w.Varuint32(&pk.SpawnY)
	w.Varint32(&pk.SpawnZ)
	w.Bool(&pk.HasAchievementsDisabled)
	w.Varint32(&pk.DayCycleStopTime)
	w.Bool(&pk.EduMode)
	w.Float32(&pk.RainLevel)
	w.Float32(&pk.LightningLevel)
	w.Bool(&pk.CommandsEnabled)
	w.Bool(&pk.IsTexturePacksRequired)
	w.Varuint32(&pk.Gamerules)
	w.String(&pk.LevelID)
	w.String(&pk.WorldName)
	w.String(&pk.PremiumWorldTemplateID)
}

// Unmarshal ...
func (pk *StartGamePacket) Unmarshal(r *PacketReader) {
	pk.PacketName = getName(pk)
	r.Varint32(&pk.EntityRuntimeID)
	r.Varint32(&pk.EntityUniqueID)
	r.Varint32(&pk.PlayerGamemode)
	r.Float32(&pk.X)
	r.Float32(&pk.Y)
	r.Float32(&pk.Z)
	r.Float32(&pk.Pitch)
	r.Float32(&pk.Yaw)
	r.Varint32(&pk.Seed)
	r.Varint32(&pk.Dimension)
	r.Varint32(&pk.Generator)
	r.Varint32(&pk.WorldGamemode)
	r.Varint32(&pk.Difficulty)
	r.Varint32(&pk.SpawnX)
	r.Varuint32(&pk.SpawnY)
	r.Varint32(&pk.SpawnZ)
	r.Bool(&pk.HasAchievementsDisabled)
	r.Varint32(&pk.DayCycleStopTime)
	r.Bool(&pk.EduMode)
	r.Float32(&pk.RainLevel)
	r.Float32(&pk.LightningLevel)
	r.Bool(&pk.CommandsEnabled)
	r.Bool(&pk.IsTexturePacksRequired)
	r.Varuint32(&pk.Gamerules)
	r.String(&pk.LevelID)
	r.String(&pk.WorldName)
	r.String(&pk.PremiumWorldTemplateID)
}
