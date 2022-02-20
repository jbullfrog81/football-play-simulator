# football-play-simulator

Program for 7 on 7 American Flag Football teams to build and simulate their playbooks.

## Menu

### Menu Items

#### Menu Order
Cycle through the different menus with ESC key

View Plays in Offensive Playbook -> View Offensive Formations -> Build Offensive Playbook -> Run Offensive Plays in Playbook

#### View Plays in Offensive Playbook
View the plays in the default offensive playbook
- Controls
  - left / right: cycles through all of the plays

#### View Offensive Formations
View all of the built in offensive formations
- Controls
  - left / right: cycles through all of the offensive formations

#### Offensive Playbook
There are several different options for the Offensive Playbook Menu:
- Use Default Playbook
- Load Offensive Playbook
- Create Offensive Playbook

##### Use Default Playbook
This will load the default playbook and allow the user to view and run the plays.

##### Load Offensive Playbook
This will allow the user to load a playbook from file then do the following:
- Edit the playbook
  - Add plays to the playbook
  - Edit the plays of the playbook
- View plays in the playbook
- Run plays in the playbook
- Create PDF of Playbook
  - Generates a PDF of all the plays in group of 8
    - Useful for printing off the playbook to fit into 3 play panel compartment wristband

##### Create Offensive Playbook
Build an offensive playbook by first selecting your offensive formation then the route for each player.

Flow for building a playbook
- First - "Formation": select the formation for the play
  - Controls:
    - enter - move to next step (select route for each player)
    - up / down - cycle through all the available formations
- Second - "Routes": select the route per player in the formation
  - Route is displayed in red on the field in front of the player
    Shows selected routes for each player in yellow
  - Controls:
    - Enter - to select the route for the current player
    - up / down - cycle through all the available routes
    - left / right - cycle through all of the players
    - tab - done with building play; move to confirmation step
- Third - "Confirmation": confirm the playbook page is done
  - Controls:
    - enter - confirms newly created playbook is correct and goes to Done
    - tab - goes back to routes menu
- Fourth - "Done": add the page to the playbook
  - Adds page to the playbook
  - Saves the playbook to file "Build.playbook
  - returns back to formation flow


#### Run Offensive Plays in Playbook
View the plays in the offensive playbook with each player running their route.
- Controls
  - left / right: cycles through the plays
  - space - pause/unpuase the running play
  - enter - restart the play

## Offensive Formations

### Bunch receivers - Quarterback in shotgun
- Bunch left with no receivers
- Bunch right with no receivers

### Split running backs - Quarterback in shotgun
- Reciever on the left
- Reciever on the right

### Twin receivers - Quarterback in shotgun
- Twins left with running back on left side
- Twins left with running back on right side
- Twins right with Running Back on left side
- Twins right with Running Back on right side

### Trips receivers - Quarterback in shotgun
- Trips left with no receivers
- Trips right with no receivers

## Offensive Routes

### Side Agnostic Routes
- Block
- Go

### Left Side Routes
- Out to sideline
- Slant - 3 yard
- Out - 5 yard
- Out - 10 yard
- In - 5 yard
- In - 10 yard
- Post - 10 yard
- Whip - 5 yard
- Out and Up - 7 yard

### Right Side Routes
- Out to sideline
- Slant - 3 yard
- Out - 5 yard
- Out - 10 yard
- In - 5 yard
- In - 10 yard
- Post - 10 yard
- Whip - 5 yard
- Out and Up - 7 yard

## Dependencies


### [Pixel](https://github.com/faiface/pixel)
[Pixel](https://github.com/faiface/pixel) is the backbone and has a backend called [PixelGL](https://godoc.org/github.com/faiface/pixel/pixelgl) backend which uses OpenGL to render graphics. So you'll need OpenGL development libraries. More details on Pixel [requirements](https://github.com/faiface/pixel#requirements) before you can compile unless you want to see the error message[1]. I was able to run `go get -u github.com/go-gl/glfw/v3.3/glfw` then `go mod vendor` to pull the dependency without manually building glfw with CMake.

[1]
```
# github.com/go-gl/glfw/v3.3/glfw
vendor/github.com/go-gl/glfw/v3.3/glfw/c_glfw.go:4:10: fatal error: 'glfw/src/context.c' file not found
#include "glfw/src/context.c"
         ^~~~~~~~~~~~~~~~~~~~
1 error generated.
```