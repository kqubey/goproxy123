package dataPacket

const (
	IDTextPacket                       = 0x9
	IDLoginPacket                      = 0x1 //only client
	IDPlayStatusPacket                 = 0x2
	IDResourcePackDataInfoPacket       = 0x53
	IDResourcePacksInfoPacket          = 0x6
	IDResourcePackStackPacket          = 0x7
	IDStartGamePacket                  = 0x0b
	IDDisconnectPacket                 = 0x5
	IDResourcePackClientResponsePacket = 0x8 //only client
	IDRequestChunkRadiusPacket         = 0x45
	IDPlayerListPacket                 = 0x3f
	IDRespawnPacket                    = 0x2d
	IDAvailableCommandsPacket          = 0x4e
	IDFullChunkPacket                  = 0x3a
	IDTransferPacket                   = 0x56
	IDInteractPacket                   = 0x21
	IDContainerSetContentPacket        = 0x34
	IDAnimatePacket                    = 0x2c
	IDContainerSetSlotPacket           = 0x32
	IDSetEntityDataPacket              = 0x27
	IDPlayerMovePacket                 = 0x13
	IDEntityMovePacket                 = 0x12
	IDPlayerActionPacket               = 0x24
	IDMobEquipmentPacket               = 0x1f
	IDSetEntityMotionPacket            = 0x28
	IDUpdateBlockPacket                = 0x16
	IDCommandStepPacket                = 0x4f
	IDRiderJumpPacket                  = 0x14
	IDResourcePackChunkDataPacket      = 0x54
	IDResourcePackChunkRequestPacket   = 0x55
	IDClientBoundMapItemDataPacket     = 0x43
	IDServerToClientHandshakePacket    = 0x03
	IDClientToServerHandshakePacket    = 0x04
	IDSetTitlePacket                   = 0x59
	IDMobArmorEquipmentPacket          = 0x20
	IDSetPlayerGameTypePacket          = 0x3e
	IDUpdateAttributesPacket           = 0x1e
	IDCraftingEventPacket              = 0x36
	IDAdventureSettingsPacket          = 0x37
	IDBlockEntityDataPacket            = 0x38
)
