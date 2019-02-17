package main

import "fmt"
import "time"
import "github.com/amimof/huego"

var huebridge *huego.Bridge

func GetBridge() (*huego.Bridge, error){
	var err error
	var config *huego.Config
	var UUID string = "BgtBLx2Hfzb0FYYZYyaqtRDZzVydPle5TXNcJpfA"
	
	bridge, _ := huego.Discover()
	
	for {
		bridge = bridge.Login(UUID)
		config, err = bridge.GetConfig(); 
		
		if err != nil {
			fmt.Println(err)
		}

		if config.PortalState.SignedOn {
			break
		}

		if config.LinkButton {
			UUID, err = bridge.CreateUser("hueman")
			if err != nil {
				break
			}
		}

		time.Sleep(200*time.Millisecond)
	}
	
	return bridge, err
}

func FixState(){
	lights, _ := huebridge.GetLights()
	for _, l := range lights {
		if (*l.State).On {
			// fmt.Printf("%s, %+v\n", l.Name, *l.State)
			state := *l.State
			defhsc := state.Hue == 8418 && state.Sat == 140
			defhsm := state.Hue == 0 && state.Sat == 0
			if  (defhsc || defhsm) && state.Ct == 366 {
				l.Bri(254)
				l.Hue(0)
				l.Sat(0)
				l.Ct(1)
			}
		}
	}	
}

func main() {
	huebridge, _ = GetBridge()

	// users, _ := huebridge.GetUsers()
	// for _, u := range users {
	// 	fmt.Printf("%+v\n", u)
	// 	if u.Name == "hueman" && u.Username != "BgtBLx2Hfzb0FYYZYyaqtRDZzVydPle5TXNcJpfA" {
	// 		huebridge.DeleteUser(u.Username)
	// 	}
	// }

	for {
		FixState()
		time.Sleep(5000*time.Millisecond)
	}
}
