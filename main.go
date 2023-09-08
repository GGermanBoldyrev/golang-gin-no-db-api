package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Grahphics struct {
	Resolution       string `json:"res"`
	Shadow           string `json:"shadow"`
	AntiAliasing     string `json:"anti"`
	Fxaa             string `json:"fxaa"`
	TextureFiltering string `json:"filtering"`
}

type Settings struct {
	Edpi      uint16 `json:"edpi"`
	Crosshair string `json:"crosshair"`
	Viewmodel string `json:"viewmodel"`
	Bob       string `json:"bob"`
}

type Player struct {
	ID        uint16     `json:"id"`
	Name      string     `json:"name"`
	Team      string     `json:"team"`
	Grahphics *Grahphics `json:"grahphics"`
	Settings  *Settings  `json:"settings"`
}

var players = []Player{
	{
		ID:   0,
		Name: "Xantares",
		Team: "Eternal Fire",
		Grahphics: &Grahphics{
			Resolution:       "1024x720",
			Shadow:           "Very Low",
			AntiAliasing:     "4x MSAA",
			Fxaa:             "Disabled",
			TextureFiltering: "Anisotropic 4x",
		},
		Settings: &Settings{
			Edpi:      800,
			Crosshair: "cl_crosshair_drawoutline 0; cl_crosshairalpha 255; cl_crosshaircolor 1; cl_crosshaircolor_b 50; cl_crosshaircolor_g 250; cl_crosshaircolor_r 50; cl_crosshairdot 0; cl_crosshairgap 0; cl_crosshairsize 3; cl_crosshairstyle 4; cl_crosshairthickness 0.5; cl_crosshair_sniper_width 1;",
			Viewmodel: "viewmodel_fov 60; viewmodel_offset_x 1; viewmodel_offset_y 1; viewmodel_offset_z -1; viewmodel_presetpos 1; cl_viewmodel_shift_left_amt 1.5; cl_viewmodel_shift_right_amt 0.75; viewmodel_recoil 0; cl_righthand 1;",
			Bob:       "cl_bob_lower_amt 21; cl_bobamt_lat 0.33; cl_bobamt_vert 0.14; cl_bobcycle 0.98;",
		},
	},
	{
		ID:   1,
		Name: "Niko",
		Team: "G2 Esports",
		Grahphics: &Grahphics{
			Resolution:       "1152x864",
			Shadow:           "High",
			AntiAliasing:     "8x MSAA",
			Fxaa:             "Disabled",
			TextureFiltering: "Anisotropic 8x",
		},
		Settings: &Settings{
			Edpi:      540,
			Crosshair: "cl_crosshair_drawoutline 0; cl_crosshairalpha 255; cl_crosshaircolor 5; cl_crosshaircolor_b 255; cl_crosshaircolor_g 255; cl_crosshaircolor_r 255; cl_crosshairdot 0; cl_crosshairgap -4; cl_crosshairsize 1; cl_crosshairstyle 4; cl_crosshairthickness 1; cl_crosshair_sniper_width 1;",
			Viewmodel: "viewmodel_fov 68; viewmodel_offset_x 2.5; viewmodel_offset_y 0; viewmodel_offset_z -1.5; viewmodel_presetpos 3; cl_viewmodel_shift_left_amt 1.5; cl_viewmodel_shift_right_amt 0.75; viewmodel_recoil 0; cl_righthand 1;",
			Bob:       "cl_bob_lower_amt 15; cl_bobamt_lat 0.33; cl_bobamt_vert 0.14; cl_bobcycle 0.98;",
		},
	},
	{
		ID:   2,
		Name: "Ropz",
		Team: "FaZe Clan",
		Grahphics: &Grahphics{
			Resolution:       "1920x1080",
			Shadow:           "Very Low",
			AntiAliasing:     "4x MSAA",
			Fxaa:             "Disabled",
			TextureFiltering: "Anisotropic 4x",
		},
		Settings: &Settings{
			Edpi:      708,
			Crosshair: "cl_crosshair_drawoutline 0; cl_crosshairalpha 255; cl_crosshaircolor 1; cl_crosshaircolor_b -1000; cl_crosshaircolor_g 0; cl_crosshaircolor_r -1000; cl_crosshairdot 0; cl_crosshairgap -3; cl_crosshairsize 2; cl_crosshairstyle 4; cl_crosshairthickness 0; cl_crosshair_sniper_width 1;",
			Viewmodel: "viewmodel_fov 68; viewmodel_offset_x 2.5; viewmodel_offset_y 0; viewmodel_offset_z -1.5; viewmodel_presetpos 3; cl_viewmodel_shift_left_amt 0.5; cl_viewmodel_shift_right_amt 0.25; viewmodel_recoil 0; cl_righthand 1;",
			Bob:       "cl_bob_lower_amt 5; cl_bobamt_lat 0.1; cl_bobamt_vert 0.1; cl_bobcycle 0.98;",
		},
	},
}

// Returns all players
func getPlayers(cxt *gin.Context) {
	cxt.IndentedJSON(http.StatusOK, players)
}

// Create a new player settings
func postPlayers(cxt *gin.Context) {
	var newPlayer Player

	if err := cxt.BindJSON(&newPlayer); err != nil {
		return
	}

	// Add the new player to the slice
	players = append(players, newPlayer)
	cxt.IndentedJSON(http.StatusCreated, newPlayer)
}

// Returns Player if name matches
func getPlayersByName(cxt *gin.Context) {
	name := cxt.Param("name")

	// Iterate Players to find name from GET param
	for _, player := range players {
		if strings.EqualFold(player.Name, name) {
			cxt.IndentedJSON(http.StatusOK, player)
			return
		}
	}

	// Return not found
	cxt.IndentedJSON(http.StatusNotFound, gin.H{"message": "player not found"})
}

func main() {
	router := gin.Default()

	// GetAll Players
	router.GET("/players", getPlayers)
	// Create the new Player
	router.POST("/players", postPlayers)
	// Search Player by name
	router.GET("/players/:name", getPlayersByName)

	// Enter Point
	router.Run("localhost:8080")
}
