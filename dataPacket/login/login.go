package login

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"goproxy/jose"
	"goproxy/jose/jwt"
	"strconv"
	"strings"
	"time"
)

type chain []string

type loginRequest struct {
	// Chain is the client certificate chain. It holds several claims that the server may verify in order to
	// make sure that the client is logged into XBOX Live.
	Chain chain `json:"chain"`
	// RawToken holds the raw token that follows the JWT chain, holding the ClientData.
	RawToken string `json:"-"`
}
type AuthResult struct {
	PublicKey             *ecdsa.PublicKey `json:"PublicKey"`
	XBOXLiveAuthenticated bool             `json:"XBOXLiveAuthenticated,omitempty"`
}
type IdentityData struct {
	// XUID is the XBOX Live user ID of the player, which will remain consistent as long as the player is
	// logged in with the XBOX Live account. It is empty if the user is not logged into its XBL account.
	XUID string `json:"XUID,omitempty"`
	// Identity is the UUID of the player, which will also remain consistent for as long as the user is logged
	// into its XBOX Live account.
	Identity string `json:"identity"`
	// DisplayName is the username of the player, which may be changed by the user. It should for that reason
	// not be used as a key to store information.
	DisplayName string `json:"displayName"`
	// TitleID is a numerical ID present only if the user is logged into XBL. It holds the title ID (XBL
	// related) of the version that the player is on. Some of these IDs may be found below.
	// Win10: 896928775
	// Mobile: 1739947436
	// Nintendo: 2047319603
	// Note that these IDs are protected using XBOX Live, making the spoofing of this data very difficult.
	TitleID string `json:"titleId,omitempty"`
}
type ClientData struct {
	ClientRandomID int64 `json:"ClientRandomId"`
	// CurrentInputMode is the input mode used by the client. It is 1 for mobile and win10, but is different
	ADRole    int    `json:"ADRole"`
	GuiScale  int    `json:"GuiScale"`
	TenantId  string `json:"TenantId"`
	UIProfile int    `json:"UIProfile"`
	// for console input.
	CurrentInputMode int `json:"CurrentInputMode"`
	// DefaultInputMode is t he default input mode used by the device.
	DefaultInputMode int `json:"DefaultInputMode"`
	// DeviceModel is a string indicating the device model used by the player. At the moment, it appears that
	// this name is always '(Standard system devices) System devices'.
	DeviceModel string `json:"DeviceModel"`
	// DeviceOS is a numerical ID indicating the OS of the device.
	DeviceOS int `json:"DeviceOS"`
	// GameVersion is the game version of the player that attempted to join, for example '1.11.0'.
	GameVersion   string `json:"GameVersion"`
	ServerAddress string `json:"ServerAddress"`
	LanguageCode  string `json:"LanguageCode"`
	SkinData      string `json:"SkinData"`
	SkinID        string `json:"SkinId"`
}

func (c identityClaims) Validate(e jwt.Expected) error {
	if err := c.Claims.Validate(e); err != nil {
		return err
	}
	return c.ExtraData.Validate()
}
func (data IdentityData) Validate() error {
	if _, err := strconv.ParseInt(data.XUID, 10, 64); err != nil && len(data.XUID) != 0 {
		return fmt.Errorf("XUID must be parseable as an int64, but got %v", data.XUID)
	}
	if _, err := uuid.Parse(data.Identity); err != nil {
		return fmt.Errorf("UUID must be parseable as a valid UUID, but got %v", data.Identity)
	}
	if len(data.DisplayName) == 0 || len(data.DisplayName) > 15 {
		return fmt.Errorf("DisplayName must not be empty or longer than 15 characters, but got %v characters", len(data.DisplayName))
	}
	if data.DisplayName[0] == ' ' || data.DisplayName[len(data.DisplayName)-1] == ' ' {
		return fmt.Errorf("DisplayName may not have a space as first/last character, but got %v", data.DisplayName)
	}

	// We check here if the name contains at least 2 spaces after each other, which is not allowed. The name
	// is only allowed to have single spaces.
	if strings.Contains(data.DisplayName, "  ") {
		return fmt.Errorf("DisplayName must only have single spaces, but got %v", data.DisplayName)
	}
	return nil
}

func decodeChain(buf *bytes.Buffer) (chain, error) {
	var chainLength int32
	if err := binary.Read(buf, binary.LittleEndian, &chainLength); err != nil {
		return nil, fmt.Errorf("error reading chain length: %v", err)
	}
	chainData := buf.Next(int(chainLength))

	request := &loginRequest{}
	if err := json.Unmarshal(chainData, request); err != nil {
		return nil, fmt.Errorf("error decoding request chain JSON: %v", err)
	}
	// First check if the chain actually has any elements in it.
	if len(request.Chain) == 0 {
		return nil, fmt.Errorf("connection request had no claims in the chain")
	}
	return request.Chain, nil
}

type identityClaims struct {
	jwt.Claims

	// ExtraData holds the extra data of this claim, which is the IdentityData of the player.
	ExtraData IdentityData `json:"extraData"`

	IdentityPublicKey string `json:"identityPublicKey"`
}

type request struct {
	// Chain is the client certificate chain. It holds several claims that the server may verify in order to
	// make sure that the client is logged into XBOX Live.
	Chain chain `json:"chain"`
	// RawToken holds the raw token that follows the JWT chain, holding the ClientData.
	RawToken string `json:"-"`
}

func MarshalPublicKey(key *ecdsa.PublicKey) string {
	data, _ := x509.MarshalPKIXPublicKey(key)
	return base64.StdEncoding.EncodeToString(data)
}

func EncodeOffline(identityData IdentityData, data ClientData, key *ecdsa.PrivateKey) []byte {
	keyData := MarshalPublicKey(&key.PublicKey)
	claims := jwt.Claims{
		Expiry:    jwt.NewNumericDate(time.Now().Add(300 * time.Second)),
		NotBefore: jwt.NewNumericDate(time.Now().Add(300 * -time.Second)),
	}

	signer, _ := jose.NewSigner(jose.SigningKey{Key: key, Algorithm: jose.ES384}, &jose.SignerOptions{
		ExtraHeaders: map[jose.HeaderKey]interface{}{"x5u": keyData},
	})
	firstJWT, _ := jwt.Signed(signer).Claims(identityClaims{
		Claims:            claims,
		ExtraData:         identityData,
		IdentityPublicKey: keyData,
	}).CompactSerialize()

	request := &request{Chain: chain{firstJWT}}
	// We create another token this time, which is signed the same as the claim we just inserted in the chain,
	// just now it contains client data.
	request.RawToken, _ = jwt.Signed(signer).Claims(data).CompactSerialize()
	return loginEncodeRequest(request)
}

// encodeRequest encodes the request passed to a byte slice which is suitable for setting to the Connection
// Request field in a Login packet.
func loginEncodeRequest(req *request) []byte {
	chainBytes, _ := json.Marshal(req)
	//fmt.Println(string(chainBytes))
	buf := bytes.NewBuffer(nil)
	_ = binary.Write(buf, binary.LittleEndian, int32(len(chainBytes)))
	_, _ = buf.WriteString(string(chainBytes))

	_ = binary.Write(buf, binary.LittleEndian, int32(len(req.RawToken)))
	_, _ = buf.WriteString(req.RawToken)
	return buf.Bytes()
}

func GetLoginEncodedBytes(id IdentityData, cd ClientData) ([]byte, *ecdsa.PrivateKey) {
	key, _ := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	return EncodeOffline(id, cd, key), key
}
