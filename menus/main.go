package menu

import (
	"jbullfrog81/football-play-simulator/menu"

	"github.com/gen2brain/dlgs"
)

type MenuMainOptionsEnum menu.MenuOptionsEnum

//Build Enums for the main menu
const (
	MenuMainOffensivePlaybook MenuMainOptionsEnum = iota + 1
	MenuMainViewOffensiveFormations
	MenuMainViewOffensiveRoutes
	MenuMainViewDefensiveFormations
	MainMenuExit
	MainMenuError
)

//Declare the text representation for each of the main menu options
const (
	MenuMainOffensivePlaybookText       menu.MenuOptionsText = "Offensive Playbook"
	MenuMainViewOffensiveFormationsText menu.MenuOptionsText = "View Offensive Formations"
	MenuMainViewOffensiveRoutesText     menu.MenuOptionsText = "View Offensive Routes"
	MenuMainViewDefensiveFormationsText menu.MenuOptionsText = "View Defensive Formations"
	MenuMainExitText                    menu.MenuOptionsText = "Exit"
	MainMenuErrorText                   menu.MenuOptionsText = "Error"
)

func (m MenuMainOptionsEnum) String() menu.MenuOptionsText {
	switch m {
	case MenuMainOffensivePlaybook:
		return MenuMainOffensivePlaybookText
	case MenuMainViewOffensiveFormations:
		return MenuMainViewOffensiveFormationsText
	case MenuMainViewOffensiveRoutes:
		return MenuMainViewOffensiveRoutesText
	case MenuMainViewDefensiveFormations:
		return MenuMainViewDefensiveFormationsText
	case MainMenuExit:
		return MenuMainExitText
	}
	return MainMenuErrorText
}

func MainMenu() MenuMainOptionsEnum {

	//TODO - for dlgs do we use List and make a slice of string or is there another way to do a menu?
	selectedMenuItem, okSelected, err := dlgs.List("Main Menu", "Program Options:", mainMenuOptions)

	if err != nil {
		panic(err)
	}

	if !okSelected {
		return windowMenuPrevious
	} else {
		return mainMenuLookup[selectedMenuItem]
	}

}

//func mainMenu() (wMenu string) {
//
//}
