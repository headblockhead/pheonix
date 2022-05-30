package main

import (
	"fmt"
	"math/rand"
)

type Scene struct {
	seed    uint16
	pattern Pattern
	frames  []Frame
}

type Pattern struct {
	frametypes []string
}

type Frame struct {
	frameType string
	location  Location
	character Character
}

type Location struct {
	background string
	desk       string
}

type Character struct {
	characterType string
	expression    Expression
	name          string
}

type Expression struct {
	ID      string
	Emotion string
}

func main() {
	pattern1 := Pattern{frametypes: []string{"D", "P", "D", "C"}}
	pattern2 := Pattern{frametypes: []string{"D", "J", "P", "J"}}
	pattern3 := Pattern{frametypes: []string{"D", "W", "D", "W"}}
	pattern4 := Pattern{frametypes: []string{"D", "W", "P", "W"}}
	pattern5 := Pattern{frametypes: []string{"D", "W", "C", "J"}}
	patterns := []Pattern{pattern1, pattern2, pattern3, pattern4, pattern5}

	expressionThinkPheonix := Expression{ID: "think", Emotion: "Confident"}
	expressionConfidentPheonix := Expression{ID: "confident", Emotion: "Confident"}
	expressionReadPheonix := Expression{ID: "read", Emotion: "Confident"}
	expressionPointPheonix := Expression{ID: "point", Emotion: "Accusation"}
	expressionDeskSlamPheonix := Expression{ID: "deskSlam", Emotion: "Accusation"}
	expressionCorneredPheonix := Expression{ID: "cornered", Emotion: "Panic"}
	expressionSillyPheonix := Expression{ID: "silly", Emotion: "Panic"}
	expressionBrekdownPheonix := Expression{ID: "brekdown", Emotion: "Sad"}
	expressionCoffePheonix := Expression{ID: "coffe", Emotion: "Sad"}
	PheonixExpressions := []Expression{expressionThinkPheonix, expressionConfidentPheonix, expressionReadPheonix, expressionPointPheonix, expressionDeskSlamPheonix, expressionCorneredPheonix, expressionSillyPheonix, expressionBrekdownPheonix, expressionCoffePheonix}

	PheonixExpressionsConfident := []Expression{}
	PheonixExpressionsPanic := []Expression{}
	PheonixExpressionsAccusation := []Expression{}
	PheonixExpressionsSad := []Expression{}

	for i := 0; i < len(PheonixExpressions); i++ {
		if PheonixExpressions[i].Emotion == "Confident" {
			PheonixExpressionsConfident = append(PheonixExpressionsConfident, PheonixExpressions[i])
		}
		if PheonixExpressions[i].Emotion == "Panic" {
			PheonixExpressionsPanic = append(PheonixExpressionsPanic, PheonixExpressions[i])
		}
		if PheonixExpressions[i].Emotion == "Accusation" {
			PheonixExpressionsAccusation = append(PheonixExpressionsAccusation, PheonixExpressions[i])
		}
		if PheonixExpressions[i].Emotion == "Sad" {
			PheonixExpressionsSad = append(PheonixExpressionsSad, PheonixExpressions[i])
		}
	}

	expressionArmsEdgeworth := Expression{ID: "arms", Emotion: "Confident"}
	expressionConfidentEdgeworth := Expression{ID: "confident", Emotion: "Confident"}
	expressionReadEdgeworth := Expression{ID: "read", Emotion: "Confident"}
	expressionCorneredEdgeworth := Expression{ID: "cornered", Emotion: "Panic"}
	expressionDeskSlamEdgeworth := Expression{ID: "deskSlam", Emotion: "Accusation"}
	expressionPointEdgeworth := Expression{ID: "point", Emotion: "Accusation"}
	expressionStandEdgeworth := Expression{ID: "stand", Emotion: "Default"}
	EdgeworthExpressions := []Expression{expressionArmsEdgeworth, expressionConfidentEdgeworth, expressionCorneredEdgeworth, expressionDeskSlamEdgeworth, expressionPointEdgeworth, expressionReadEdgeworth, expressionStandEdgeworth}

	EdgeworthExpressionsConfident := []Expression{}
	EdgeworthExpressionsPanic := []Expression{}
	EdgeworthExpressionsAccusation := []Expression{}
	EdgeworthExpressionsDefault := []Expression{}

	for i := 0; i < len(EdgeworthExpressions); i++ {
		if EdgeworthExpressions[i].Emotion == "Confident" {
			EdgeworthExpressionsConfident = append(EdgeworthExpressionsConfident, EdgeworthExpressions[i])
		}
		if EdgeworthExpressions[i].Emotion == "Panic" {
			EdgeworthExpressionsPanic = append(EdgeworthExpressionsPanic, EdgeworthExpressions[i])
		}
		if EdgeworthExpressions[i].Emotion == "Accusation" {
			EdgeworthExpressionsAccusation = append(EdgeworthExpressionsAccusation, EdgeworthExpressions[i])
		}
		if EdgeworthExpressions[i].Emotion == "Default" {
			EdgeworthExpressionsDefault = append(EdgeworthExpressionsDefault, EdgeworthExpressions[i])
		}
	}
	expressionStandJudge := Expression{ID: "stand", Emotion: "Default"}
	expressionAngryJudge := Expression{ID: "angry", Emotion: "Accusation"}
	expressionSurpriseJudge := Expression{ID: "surprise", Emotion: "Default"}
	JudgeExpressions := []Expression{expressionStandJudge, expressionAngryJudge, expressionSurpriseJudge}

	JudgeExpressionsDefault := []Expression{}
	JudgeExpressionsAccusation := []Expression{}

	for i := 0; i < len(JudgeExpressions); i++ {
		if JudgeExpressions[i].Emotion == "Default" {
			JudgeExpressionsDefault = append(JudgeExpressionsDefault, JudgeExpressions[i])
		}
		if JudgeExpressions[i].Emotion == "Accusation" {
			JudgeExpressionsAccusation = append(JudgeExpressionsAccusation, JudgeExpressions[i])
		}
	}

	character1 := Character{characterType: "D", name: "Pheonix Wright"}
	character2 := Character{characterType: "P", name: "Edgeworth"}
	character3 := Character{characterType: "J", name: "Judge"}
	character4 := Character{characterType: "W", name: "Cody Hackins"}
	character5 := Character{characterType: "W", name: "De Vasquez"}
	character6 := Character{characterType: "W", name: "Detective Gumshoe"}
	character7 := Character{characterType: "W", name: "Godot"}
	character8 := Character{characterType: "W", name: "Larry Butz"}
	character9 := Character{characterType: "W", name: "Yanni Yogi"}
	character10 := Character{characterType: "C", name: "Crowd"}
	characters := []Character{character1, character2, character3, character4, character5, character6, character7, character8, character9, character10}

	charactersD := []Character{}
	charactersP := []Character{}
	charactersJ := []Character{}
	charactersW := []Character{}
	charactersC := []Character{}

	for i := 0; i < len(characters); i++ {
		if characters[i].characterType == "D" {
			charactersD = append(charactersD, characters[i])
		}
		if characters[i].characterType == "P" {
			charactersP = append(charactersP, characters[i])
		}
		if characters[i].characterType == "J" {
			charactersJ = append(charactersJ, characters[i])
		}
		if characters[i].characterType == "W" {
			charactersW = append(charactersW, characters[i])
		}
		if characters[i].characterType == "C" {
			charactersC = append(charactersC, characters[i])
		}
	}
	frame1 := Frame{}
	frame2 := Frame{}
	frame3 := Frame{}
	frame4 := Frame{}

	scene := Scene{seed: 121, frames: []Frame{frame1, frame2, frame3, frame4}}

	rand.Seed(int64(scene.seed))
	randomPattern := rand.Intn(len(patterns))
	chosenPattern := patterns[randomPattern]
	scene.pattern = chosenPattern
	frame1.frameType = chosenPattern.frametypes[0]
	frame2.frameType = chosenPattern.frametypes[1]
	frame3.frameType = chosenPattern.frametypes[2]
	frame4.frameType = chosenPattern.frametypes[3]

	frame1 = RandomFrameCharacter(frame1, charactersD, charactersP, charactersJ, charactersW, charactersC)
	frame2 = RandomFrameCharacter(frame2, charactersD, charactersP, charactersJ, charactersW, charactersC)
	frame3 = RandomFrameCharacter(frame3, charactersD, charactersP, charactersJ, charactersW, charactersC)
	frame4 = RandomFrameCharacter(frame4, charactersD, charactersP, charactersJ, charactersW, charactersC)

	scene.frames[0] = frame1
	scene.frames[1] = frame2
	scene.frames[2] = frame3
	scene.frames[3] = frame4
	fmt.Println(scene)
}

func RandomFrameCharacter(frame Frame, charactersD []Character, charactersP []Character, charactersJ []Character, charactersW []Character, charactersC []Character) (outframe Frame) {
	if frame.frameType == "D" {
		frame.character = charactersD[rand.Intn(len(charactersD))]
	}
	if frame.frameType == "P" {
		frame.character = charactersP[rand.Intn(len(charactersP))]
	}
	if frame.frameType == "J" {
		frame.character = charactersJ[rand.Intn(len(charactersJ))]
	}
	if frame.frameType == "W" {
		frame.character = charactersW[rand.Intn(len(charactersW))]
	}
	if frame.frameType == "C" {
		frame.character = charactersC[rand.Intn(len(charactersC))]
	}
	outframe = frame
	return outframe
}
