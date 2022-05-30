package main

import (
	"fmt"
	"math/rand"
)

type Scene struct {
	seed           int
	pattern        Pattern
	frames         []Frame
	emotionpattern EmotionPattern
}

type Pattern struct {
	frametypes []string
	id         int
}

type EmotionPattern struct {
	emotionTypes []string
	pattern      Pattern
}

type Frame struct {
	frameType string
	location  Location
	character Character
	id        int
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
	pattern1 := Pattern{frametypes: []string{"D", "P", "D", "C"}, id: 1}
	pattern2 := Pattern{frametypes: []string{"D", "J", "P", "J"}, id: 2}
	pattern3 := Pattern{frametypes: []string{"D", "W", "D", "W"}, id: 3}
	pattern4 := Pattern{frametypes: []string{"D", "W", "P", "W"}, id: 4}
	pattern5 := Pattern{frametypes: []string{"D", "W", "C", "J"}, id: 5}
	patterns := []Pattern{pattern1, pattern2, pattern3, pattern4, pattern5}

	emotion_pattern1_pattern1 := EmotionPattern{emotionTypes: []string{"Accusation", "Panic", "Confident", "Default"}, pattern: pattern1}
	emotion_pattern1_pattern2 := EmotionPattern{emotionTypes: []string{"Confident", "Confident", "Panic", "Default"}, pattern: pattern1}
	emotion_pattern1_pattern3 := EmotionPattern{emotionTypes: []string{"Accusation", "Confident", "Panic", "Default"}, pattern: pattern1}
	emotions_pattern1 := []EmotionPattern{emotion_pattern1_pattern1, emotion_pattern1_pattern2, emotion_pattern1_pattern3}

	emotion_pattern2_pattern1 := EmotionPattern{emotionTypes: []string{"Panic", "Accusation", "Confident", "Panic"}, pattern: pattern2}
	emotion_pattern2_pattern2 := EmotionPattern{emotionTypes: []string{"Accusation", "Panic", "Panic", "Accusation"}, pattern: pattern2}
	emotion_pattern2_pattern3 := EmotionPattern{emotionTypes: []string{"Accusation", "Panic", "Confident", "Accusation"}, pattern: pattern2}
	emotions_pattern2 := []EmotionPattern{emotion_pattern2_pattern1, emotion_pattern2_pattern2, emotion_pattern2_pattern3}

	emotion_pattern3_pattern1 := EmotionPattern{emotionTypes: []string{"Accusation", "Default", "Sad", "Default"}, pattern: pattern3}
	emotion_pattern3_pattern2 := EmotionPattern{emotionTypes: []string{"Accusation", "Default", "Panic", "Default"}, pattern: pattern3}
	emotion_pattern3_pattern3 := EmotionPattern{emotionTypes: []string{"Accusation", "Default", "Confident", "Default"}, pattern: pattern3}
	emotions_pattern3 := []EmotionPattern{emotion_pattern3_pattern1, emotion_pattern3_pattern2, emotion_pattern3_pattern3}

	emotion_pattern4_pattern1 := EmotionPattern{emotionTypes: []string{"Panic", "Default", "Confident", "Default"}, pattern: pattern4}
	emotion_pattern4_pattern2 := EmotionPattern{emotionTypes: []string{"Accusation", "Default", "Panic", "Default"}, pattern: pattern4}
	emotion_pattern4_pattern3 := EmotionPattern{emotionTypes: []string{"Confident", "Default", "Confident", "Default"}, pattern: pattern4}
	emotions_pattern4 := []EmotionPattern{emotion_pattern4_pattern1, emotion_pattern4_pattern2, emotion_pattern4_pattern3}

	emotion_pattern5_pattern1 := EmotionPattern{emotionTypes: []string{"Accusation", "Default", "Default", "Panic"}, pattern: pattern5}
	emotion_pattern5_pattern2 := EmotionPattern{emotionTypes: []string{"Accusation", "Default", "Default", "Panic"}, pattern: pattern5}
	emotion_pattern5_pattern3 := EmotionPattern{emotionTypes: []string{"Accusation", "Default", "Default", "Panic"}, pattern: pattern5}
	emotions_pattern5 := []EmotionPattern{emotion_pattern5_pattern1, emotion_pattern5_pattern2, emotion_pattern5_pattern3}

	expressionThinkPheonix := Expression{ID: "think", Emotion: "Confident"}
	expressionConfidentPheonix := Expression{ID: "confident", Emotion: "Confident"}
	expressionReadPheonix := Expression{ID: "read", Emotion: "Confident"}
	expressionPointPheonix := Expression{ID: "point", Emotion: "Accusation"}
	expressionDeskSlamPheonix := Expression{ID: "deskSlam", Emotion: "Accusation"}
	expressionCorneredPheonix := Expression{ID: "cornered", Emotion: "Panic"}
	expressionSillyPheonix := Expression{ID: "silly", Emotion: "Panic"}
	expressionBrekdownPheonix := Expression{ID: "brekdown", Emotion: "Panic"}
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
	expressionSurpriseJudge := Expression{ID: "surprise", Emotion: "Panic"}
	JudgeExpressions := []Expression{expressionStandJudge, expressionAngryJudge, expressionSurpriseJudge}

	JudgeExpressionsDefault := []Expression{}
	JudgeExpressionsAccusation := []Expression{}
	JudgeExpressionsPanic := []Expression{}

	for i := 0; i < len(JudgeExpressions); i++ {
		if JudgeExpressions[i].Emotion == "Default" {
			JudgeExpressionsDefault = append(JudgeExpressionsDefault, JudgeExpressions[i])
		}
		if JudgeExpressions[i].Emotion == "Accusation" {
			JudgeExpressionsAccusation = append(JudgeExpressionsAccusation, JudgeExpressions[i])
		}
		if JudgeExpressions[i].Emotion == "Panic" {
			JudgeExpressionsPanic = append(JudgeExpressionsPanic, JudgeExpressions[i])
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
	frame1 := Frame{id: 0}
	frame2 := Frame{id: 1}
	frame3 := Frame{id: 2}
	frame4 := Frame{id: 3}

	var seed string
	fmt.Printf("Enter a seed: ")
	fmt.Scanln(&seed)
	seedBytes := []byte(seed)
	var finalSeed int
	for i := 0; i < len(seedBytes); i++ {
		finalSeed += int(seedBytes[i])
	}
	scene := Scene{seed: finalSeed, frames: []Frame{frame1, frame2, frame3, frame4}}

	rand.Seed(int64(scene.seed))
	randomPattern := rand.Intn(len(patterns))
	chosenPattern := patterns[randomPattern]
	scene.pattern = chosenPattern

	if scene.pattern.id == 1 {
		scene.emotionpattern = RandomInEmotionList(emotions_pattern1)
	} else if scene.pattern.id == 2 {
		scene.emotionpattern = RandomInEmotionList(emotions_pattern2)
	} else if scene.pattern.id == 3 {
		scene.emotionpattern = RandomInEmotionList(emotions_pattern3)
	} else if scene.pattern.id == 4 {
		scene.emotionpattern = RandomInEmotionList(emotions_pattern4)
	} else if scene.pattern.id == 5 {
		scene.emotionpattern = RandomInEmotionList(emotions_pattern5)
	}

	frame1.frameType = chosenPattern.frametypes[0]
	frame2.frameType = chosenPattern.frametypes[1]
	frame3.frameType = chosenPattern.frametypes[2]
	frame4.frameType = chosenPattern.frametypes[3]

	frame1 = RandomFrameCharacter(frame1, charactersD, charactersP, charactersJ, charactersW, charactersC)
	frame2 = RandomFrameCharacter(frame2, charactersD, charactersP, charactersJ, charactersW, charactersC)
	frame3 = RandomFrameCharacter(frame3, charactersD, charactersP, charactersJ, charactersW, charactersC)
	frame4 = RandomFrameCharacter(frame4, charactersD, charactersP, charactersJ, charactersW, charactersC)

	frame1 = setexpressions(frame1, scene, PheonixExpressionsConfident, EdgeworthExpressionsConfident, EdgeworthExpressionsDefault, JudgeExpressionsDefault, JudgeExpressionsPanic, JudgeExpressionsAccusation, PheonixExpressionsAccusation, PheonixExpressionsSad, EdgeworthExpressionsPanic, EdgeworthExpressionsAccusation, JudgeExpressionsAccusation)
	frame2 = setexpressions(frame2, scene, PheonixExpressionsConfident, EdgeworthExpressionsConfident, EdgeworthExpressionsDefault, JudgeExpressionsDefault, JudgeExpressionsPanic, JudgeExpressionsAccusation, PheonixExpressionsAccusation, PheonixExpressionsSad, EdgeworthExpressionsPanic, EdgeworthExpressionsAccusation, JudgeExpressionsAccusation)
	frame3 = setexpressions(frame3, scene, PheonixExpressionsConfident, EdgeworthExpressionsConfident, EdgeworthExpressionsDefault, JudgeExpressionsDefault, JudgeExpressionsPanic, JudgeExpressionsAccusation, PheonixExpressionsAccusation, PheonixExpressionsSad, EdgeworthExpressionsPanic, EdgeworthExpressionsAccusation, JudgeExpressionsAccusation)
	frame4 = setexpressions(frame4, scene, PheonixExpressionsConfident, EdgeworthExpressionsConfident, EdgeworthExpressionsDefault, JudgeExpressionsDefault, JudgeExpressionsPanic, JudgeExpressionsAccusation, PheonixExpressionsAccusation, PheonixExpressionsSad, EdgeworthExpressionsPanic, EdgeworthExpressionsAccusation, JudgeExpressionsAccusation)

	scene.frames[0] = frame1
	scene.frames[1] = frame2
	scene.frames[2] = frame3
	scene.frames[3] = frame4
}

func RandomInExpressionList(list []Expression, scene Scene) Expression {
	random := rand.Intn(len(list))
	return list[random]
}

func RandomInEmotionList(list []EmotionPattern) EmotionPattern {
	random := rand.Intn(len(list))
	return list[random]
}

func setexpressions(frame Frame, scene Scene, PheonixExpressionsConfident []Expression, EdgeworthExpressionsConfident []Expression, EdgeworthExpressionsDefault []Expression, JudgeExpressionsDefault []Expression, JudgeExpressionsPanic []Expression, PheonixExpressionsPanic []Expression, PheonixExpressionsAccusation []Expression, PheonixExpressionsSad []Expression, EdgeworthExpressionsPanic []Expression, EdgeworthExpressionsAccusation []Expression, JudgeExpressionsAccusation []Expression) Frame {
	if frame.character.characterType == "D" {
		if scene.emotionpattern.emotionTypes[frame.id] == "Confident" {
			frame.character.expression = RandomInExpressionList(PheonixExpressionsConfident, scene)
		}
		if scene.emotionpattern.emotionTypes[frame.id] == "Panic" {
			frame.character.expression = RandomInExpressionList(PheonixExpressionsPanic, scene)
		}
		if scene.emotionpattern.emotionTypes[frame.id] == "Accusation" {
			frame.character.expression = RandomInExpressionList(PheonixExpressionsAccusation, scene)
		}
		if scene.emotionpattern.emotionTypes[frame.id] == "Sad" {
			frame.character.expression = RandomInExpressionList(PheonixExpressionsSad, scene)
		}
	}
	if frame.character.characterType == "P" {
		if scene.emotionpattern.emotionTypes[frame.id] == "Confident" {
			frame.character.expression = RandomInExpressionList(EdgeworthExpressionsConfident, scene)
		}
		if scene.emotionpattern.emotionTypes[frame.id] == "Panic" {
			frame.character.expression = RandomInExpressionList(EdgeworthExpressionsPanic, scene)
		}
		if scene.emotionpattern.emotionTypes[frame.id] == "Accusation" {
			frame.character.expression = RandomInExpressionList(EdgeworthExpressionsAccusation, scene)
		}
		if scene.emotionpattern.emotionTypes[frame.id] == "Default" {
			frame.character.expression = RandomInExpressionList(EdgeworthExpressionsDefault, scene)
		}
	}
	if frame.character.characterType == "J" {
		if scene.emotionpattern.emotionTypes[frame.id] == "Panic" {
			frame.character.expression = RandomInExpressionList(JudgeExpressionsPanic, scene)
		}
		if scene.emotionpattern.emotionTypes[frame.id] == "Accusation" {
			frame.character.expression = RandomInExpressionList(JudgeExpressionsAccusation, scene)
		}
		if scene.emotionpattern.emotionTypes[frame.id] == "Default" {
			frame.character.expression = RandomInExpressionList(JudgeExpressionsDefault, scene)
		}
	}
	if frame.character.characterType == "W" {
		if scene.emotionpattern.emotionTypes[frame.id] == "Default" {
			frame.character.expression = Expression{ID: "default", Emotion: "Default"}
		}
	}
	if frame.character.characterType == "C" {
		if scene.emotionpattern.emotionTypes[frame.id] == "Default" {
			frame.character.expression = Expression{ID: "default", Emotion: "Default"}
		}
	}
	return frame
}

func RandomFrameCharacter(frame Frame, charactersD []Character, charactersP []Character, charactersJ []Character, charactersW []Character, charactersC []Character) (outframe Frame) {
	if frame.frameType == "D" {
		frame.character = charactersD[rand.Intn(len(charactersD))]
		frame.location.background = "defence"
		frame.location.desk = "defence"
	}
	if frame.frameType == "P" {
		frame.character = charactersP[rand.Intn(len(charactersP))]
		frame.location.desk = "prosecution"
	}
	if frame.frameType == "J" {
		frame.character = charactersJ[rand.Intn(len(charactersJ))]
		frame.location.desk = "judge"
	}
	if frame.frameType == "W" {
		frame.character = charactersW[rand.Intn(len(charactersW))]
		frame.location.desk = "witness"
	}
	if frame.frameType == "C" {
		frame.character = charactersC[rand.Intn(len(charactersC))]
		frame.location.desk = "gallery"
	}
	outframe = frame
	return outframe
}
