package main

import (
	"goproxy/dataPacket"
	"goproxy/dataPacket/login"
	"goproxy/raklib"
	"goproxy/utils"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"C"
)

var ServerCon *ServerConnection
var ClientCon *ClientConnection
var Listener *raklib.Listener
var LocalAddress string

var CurrentBindedPlayerPort = rand.Intn(65535-40000) + 40000

var Entries = map[string]string{}

var Address string = "play.minenex.ru:19132"
var Nick string = "tetett"
var Input int = 2
var Os int = 1
var Device string = "XIAOMI SS23224242MT"
var CurrentCid = -rand.Int63()

/* START */

func main(){
}

/* END */

func startListener() {
	//log.Println("Listener started")
	ClientCon = NewConnectionToClient()
	/*defer func() {
		_ = ClientCon.Connection.Close()
	}()*/

	//log.Println("Connected to client")

	buffer := make([]byte, 65535)
	ticker := time.NewTicker(3 * time.Millisecond)
	for {
		select {
		case <-ClientCon.IsShutdown:
			//log.Println("client shutdown")
			ClientCon.IsOnline = false
			/*ClientCon.SendPacket(&dataPacket.DisconnectPacket{
				HideDisconnectionScreen: false,
				Message:                 "Выключение",
			})
			_ = ClientCon.Connection.Close()*/
			return

		case <-ticker.C:
			size, _ := ClientCon.Connection.Read(buffer)
			if size == 0 {
				continue
			}

			if pks := ClientCon.DecodeBatch(buffer[1:size]); pks != nil {
				hadpks := 0
				for pos, pk := range pks {

					pkt, _ := dataPacket.ParseDataPacket(pk).TryDecodePacket()
					//C.callHandler(clientPacket, C.CString(string(pk)))
					//log.Println"from client:", reflect.TypeOf(pkt)
					switch p := pkt.(type) {
						
					case *dataPacket.LoginPacket:
						i, c, _, _ := login.Parse(p.ConnectionRequest)
						log.Println("login", c)
						LocalAddress = c.ServerAddress
						c.DeviceOS = Os
						c.ClientRandomID = CurrentCid
						c.DeviceModel = Device
						c.ServerAddress = Address
						c.CurrentInputMode = Input
						c.DefaultInputMode = 2
						c.TenantId = ""
						if Nick == ""{
							i.DisplayName = i.DisplayName
						}else{
							i.DisplayName = Nick
						}
						i.DisplayName = Nick
						c.LanguageCode = "ru_RU"
						c.GuiScale = 0
						c.ADRole = 2
						c.UIProfile = 1
						byted, _ := login.GetLoginEncodedBytes(i, c)
						p.ConnectionRequest = byted

						go startServerConn(p)
						pks = append(pks[:pos], pks[pos+1:]...)
					default:
						//log.Println("from client", p)
						//_, _ = ServerCon.Connection.Write(ServerCon.EncodeBatch([][]byte{pk}))
						// ServerCon.SendPacket(pk)

						hadpks++
					}
				}
				//log.Println("had", hadpks, "packets to send")
				if len(pks) > 0 {
					ServerCon.SendPacketsRaw(pks)
				}
			}
		}
	}
}

func startServerConn(loginpk dataPacket.DataPacket) {
	//hasstartgame := false
	var err error

//	log.Println("Connecting to server")
	ServerCon, err = NewConnectionToServer(Address)
	if err != nil {
		ClientCon.SendPacket(&dataPacket.DisconnectPacket{
			HideDisconnectionScreen: false,
			Message:                 "Сервер не ответил на запрос подключения",
		})
		ServerCon.IsOnline = false
		//log.Println(err)
		time.Sleep(300 * time.Millisecond)
		go func() {
			ClientCon.IsShutdown <- true
			_ = ClientCon.Connection.Close()
		}()
		return
	}
	//log.Println("Connected")

	//log.Println("Sending login")

	ServerCon.SendPacket(loginpk)

	ticker := time.NewTicker(3 * time.Millisecond)

	for {

		select {
		case <-ServerCon.IsShutdown:
			//log.Println("servercon shutdown")
			ServerCon.IsOnline = false
			return
		case <-ticker.C:
			buffer, err := ServerCon.Connection.ReadPacket()

			if err != nil {
				ClientCon.SendPacket(&dataPacket.DisconnectPacket{
					HideDisconnectionScreen: false,
					Message:                 "Ошибка чтения пакетов: " + err.Error(),
				})
				time.Sleep(300 * time.Millisecond)
				go func() {
					ClientCon.IsShutdown <- true
					_ = ClientCon.Connection.Close()
					ServerCon.IsShutdown <- true
					_ = ServerCon.Connection.Close()
				}()
				//log.Println("err read", err)
				continue
			} else if buffer == nil {
				continue
			}
			//log.Println(buffer)
			if pks := ServerCon.DecodeBatch(buffer[1:]); pks != nil {
				hadpks := 0

			pkloop:
				for _, pk := range pks {

					pkt, _ := dataPacket.ParseDataPacket(pk).TryDecodePacket()
					//C.callHandler(serverPacket, C.CString(string(pk)))

					//log.Println("from server:", reflect.TypeOf(pkt))
					switch p := pkt.(type) {
					//case *dataPacket.ResourcePacksInfoPacket:
					//log.Println("skipped packs")
					//pks = append(pks[:pos], pks[pos+1:]...)
					//continue
					
					
					case *dataPacket.TransferPacket:
						/*if !hasstartgame {
							pks = append(pks[:pos], pks[pos+1:]...)
							log.Println("skipped transfer, no startgame", pos)
							continue
						}*/
						Address = p.Address + ":" + strconv.Itoa(int(p.Port))
						//log.Println("transfer to", Address)
						go func(ClientCon *ClientConnection, ServerCon *ServerConnection) {
							//log.Println("transfer servercon shutdown")

							ServerCon.IsShutdown <- true
							//log.Println("transfer servercon shutdown done")
							_ = ServerCon.Connection.Close()
							//log.Println("transfer servercon closed")
							ClientCon.IsOnline = false
							//time.Sleep(1100 * time.Millisecond)
							ClientCon.SendPacket(&dataPacket.TransferPacket{
								Address: strings.Split(LocalAddress, ":")[0],
								Port:    19132,
							})
							ClientCon.IsShutdown <- true
							//log.Println("transfer clientcon shutdown done")
							//log.Println(len(ClientCon.Connection.Closed))
							//_ = ClientCon.Connection.Close()

							//log.Println("transferring...")
						}(ClientCon, ServerCon)
						pks = nil
						break pkloop
					default:
						//_, _ = ClientCon.Connection.Write(ClientCon.EncodeBatch([][]byte{pk}))
						hadpks++
					}
				}
				//log.Println("had", hadpks, "packets to send")
				if len(pks) > 0 {
					ClientCon.SendPacketsRaw(pks)
				}
			} else {
				//log.Println("nil batch from server")
			}
		}
	}
}
