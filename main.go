package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"math/rand"
	"os"
)

type Scene struct {
	seed           int
	pattern        Pattern
	frames         []Frame
	emotionpattern EmotionPattern
	objection      Objection
}

type Objection struct {
	objectionActive   bool
	objectionType     string
	objectionLocation int
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
	speech        Dialouge
}

type Dialouge struct {
	text string
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
	scene.objection.objectionActive = false
	for i := 0; i < len(scene.frames); i++ {
		if scene.frames[i].character.characterType == "D" {
			if scene.frames[i].character.expression.Emotion == "Accusation" {
				scene.objection.objectionLocation = i - 1
				scene.objection.objectionType = "Objection"
				scene.objection.objectionActive = true
				break
			}
		}
		if i > 0 {
			if scene.frames[i].character.characterType == "C" && scene.frames[i-1].character.characterType == "W" {
				scene.objection.objectionLocation = i - 1
				scene.objection.objectionType = "TakeThat"
				scene.objection.objectionActive = true
				break
			}
		}
	}
	for i := 0; i < len(scene.frames); i++ {
		scene.frames[i] = randomDialouge(scene.frames[i])
	}
	fmt.Println(scene)
	bgpath, fgpath, characterpath := getpathsFromScene(scene)
	bgimage, fgimage, characterimage := GetImages(bgpath, fgpath, characterpath)
	finalimage := AssembleImage(bgimage, fgimage, characterimage)
	saveimage(finalimage)
}

func saveimage(img image.Image) {
	f, err := os.Create("img.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err = jpeg.Encode(f, img, nil); err != nil {
		log.Printf("failed to encode: %v", err)
	}
}

func getpathsFromScene(scene Scene) (bg string, fg string, character string) {
	for i := 0; i < len(scene.frames); i++ {
		if scene.frames[i].character.characterType == "D" {
			bg = "./images/background/defence.jpg"
			fg = "./images/foreground/defence.png"
			if scene.frames[i].character.expression.ID == "think" {
				character = "./images/poses/defence/think.png"
			}
			if scene.frames[i].character.expression.ID == "confident" {
				character = "./images/poses/defence/confident.png"
			}
			if scene.frames[i].character.expression.ID == "read" {
				character = "./images/poses/defence/read.png"
			}
			if scene.frames[i].character.expression.ID == "point" {
				character = "./images/poses/defence/point.png"
			}
			if scene.frames[i].character.expression.ID == "deskSlam" {
				character = "./images/poses/defence/desk_slam.png"
			}
			if scene.frames[i].character.expression.ID == "silly" {
				character = "./images/poses/defence/silly.png"
			}
			if scene.frames[i].character.expression.ID == "breakdown" {
				character = "./images/poses/defence/breakdown.png"
			}
			if scene.frames[i].character.expression.ID == "coffe" {
				character = "./images/poses/defence/coffe.png"
			}
		}
		if scene.frames[i].character.characterType == "P" {
			bg = "./images/background/prosecution.jpg"
			fg = "./images/foreground/prosecution.png"
			if scene.frames[i].character.expression.ID == "arms" {
				character = "./images/poses/prosecution/arms.png"
			}
			if scene.frames[i].character.expression.ID == "confident" {
				character = "./images/poses/prosecution/confident.png"
			}
			if scene.frames[i].character.expression.ID == "read" {
				character = "./images/poses/prosecution/read.png"
			}
			if scene.frames[i].character.expression.ID == "cornered" {
				character = "./images/poses/prosecution/cornered.png"
			}
			if scene.frames[i].character.expression.ID == "deskSlam" {
				character = "./images/poses/prosecution/deskslam.png"
			}
			if scene.frames[i].character.expression.ID == "point" {
				character = "./images/poses/prosecution/point.png"
			}
			if scene.frames[i].character.expression.ID == "stand" {
				character = "./images/poses/prosecution/stand.png"
			}
		}
		if scene.frames[i].character.characterType == "J" {
			bg = "./images/background/judge.jpg"
			fg = "./images/foreground/judge.png"
			if scene.frames[i].character.expression.ID == "stand" {
				character = "./images/poses/judge/stand.png"
			}
			if scene.frames[i].character.expression.ID == "angry" {
				character = "./images/poses/judge/angry.png"
			}
			if scene.frames[i].character.expression.ID == "surprise" {
				character = "./images/poses/judge/surprise.png"
			}
		}
	}
	return bg, fg, character
}

func combineImages(img1 image.Image, img2 image.Image) (image.Image, error) {
	//starting position of the second image (bottom left)
	sp2 := image.Point{img1.Bounds().Dx(), 0}
	//new rectangle for the second image
	r2 := image.Rectangle{sp2, sp2.Add(img2.Bounds().Size())}
	//rectangle for the big image
	r := image.Rectangle{image.Point{0, 0}, r2.Max}
	rgba := image.NewRGBA(r)
	draw.Draw(rgba, img1.Bounds(), img1, image.Point{0, 0}, draw.Src)
	draw.Draw(rgba, r2, img2, image.Point{0, 0}, draw.Src)
	return rgba, nil
}

func AssembleImage(bg image.Image, fg image.Image, character image.Image) image.Image {
	bgchar, err := combineImages(bg, character)
	if err != nil {
		fmt.Println(err)
	}
	fgbgchar, err := combineImages(bgchar, fg)
	if err != nil {
		fmt.Println(err)
	}
	return fgbgchar
}

func GetImages(bgpath, fgpath, characterpath string) (bg image.Image, fg image.Image, character image.Image) {
	imgFile1, err := os.Open(bgpath)
	if err != nil {
		fmt.Println(err)
	}
	imgFile2, err := os.Open(fgpath)
	if err != nil {
		fmt.Println(err)
	}
	imgFile3, err := os.Open(characterpath)
	if err != nil {
		fmt.Println(err)
	}
	bg, err = jpeg.Decode(imgFile1)
	if err != nil {
		fmt.Println(err)
	}
	fg, err = png.Decode(imgFile2)
	if err != nil {
		fmt.Println(err)
	}
	character, err = png.Decode(imgFile3)
	if err != nil {
		fmt.Println(err)
	}
	return bg, fg, character
}

func randomDialouge(frame Frame) Frame {
	pheonixConfidentDialouge1 := "That is correct."
	pheonixConfidentDialouge2 := "I have evidence to back that up!"
	pheonixConfidentDialouge3 := "There is no problem here."
	pheonixConfidentDialouge4 := "There is no problem with that."
	pheonixConfidentDialouge5 := "That makes no difference."
	pheonixConfidentDialouge6 := "Just as I thought!"
	pheonixConfidentDialouge7 := "I was right!"
	pheonixConfidentDialouge8 := "I knew it!"
	pheonixConfidentDialouges := []string{pheonixConfidentDialouge1, pheonixConfidentDialouge2, pheonixConfidentDialouge3, pheonixConfidentDialouge4, pheonixConfidentDialouge5, pheonixConfidentDialouge6, pheonixConfidentDialouge7, pheonixConfidentDialouge8}
	pheonixAccusationDialouge1 := "I have evidence that proves the defendant was not at the scene of the crime!"
	pheonixAccusationDialouge2 := "You were not inside the boat!"
	pheonixAccusationDialouge3 := "That is not correct!"
	pheonixAccusationDialouge4 := "You are wrong!"
	pheonixAccusationDialouge5 := "The prosection is lying!"
	pheonixAccusationDialouge6 := "There is a vital piece of evidence missing!"
	pheonixAccusationDialouge7 := "You are not correct!"
	pheonixAccusationDialouge8 := "There is not enough evidence!"
	pheonixAccusationDialouges := []string{pheonixAccusationDialouge1, pheonixAccusationDialouge2, pheonixAccusationDialouge3, pheonixAccusationDialouge4, pheonixAccusationDialouge5, pheonixAccusationDialouge6, pheonixAccusationDialouge7, pheonixAccusationDialouge8}
	pheonixSadDialouge1 := "Ouch! That hurt!"
	pheonixSadDialouge2 := "Oww.. that burns!"
	pheonixSadDialouge3 := "Hey!"
	pheonixSadDialouges := []string{pheonixSadDialouge1, pheonixSadDialouge2, pheonixSadDialouge3}
	pheonixPanicDialouge1 := "I.. Uhhh Ummm (How do I get out of this one?)"
	pheonixPanicDialouge2 := "(Urk.)"
	pheonixPanicDialouge3 := "(uh-oh)"
	pheonixPanicDialouge4 := "(oops)"
	pheonixPanicDialouge5 := "(oh no!)"
	pheonixPanicDialouge6 := "(How do I get out of this?)"
	pheonixPanicDialouge7 := "(This isn't going well!)"
	pheonixPanicDialouge8 := "Uhhhh... Ummmm... (oh-no!)"
	pheonixPanicDialouges := []string{pheonixPanicDialouge1, pheonixPanicDialouge2, pheonixPanicDialouge3, pheonixPanicDialouge4, pheonixPanicDialouge5, pheonixPanicDialouge6, pheonixPanicDialouge7, pheonixPanicDialouge8}

	edgeworthConfidentDialouge1 := "That is incorrect."
	edgeworthConfidentDialouge2 := "There is a problem with your statement."
	edgeworthConfidentDialouge3 := "Your theory is obviously flawed."
	edgeworthConfidentDialouge4 := "There is a severe lack of evidence to back up your claims."
	edgeworthConfidentDialouge5 := "You are completely wrong."
	edgeworthConfidentDialouges := []string{edgeworthConfidentDialouge1, edgeworthConfidentDialouge2, edgeworthConfidentDialouge3, edgeworthConfidentDialouge4, edgeworthConfidentDialouge5}
	edgeworthAccusationDialouge1 := "The defence is not correct in this matter!"
	edgeworthAccusationDialouge2 := "There is something seriously wrong with this case!"
	edgeworthAccusationDialouge3 := "Objection, your honour! That statement was... Objectionable!"
	edgeworthAccusationDialouge4 := "The defence's arguments make no sense!"
	edgeworthAccusationDialouge5 := "That autopsy report is outdated!"
	edgeworthAccusationDialouges := []string{edgeworthAccusationDialouge1, edgeworthAccusationDialouge2, edgeworthAccusationDialouge3, edgeworthAccusationDialouge4, edgeworthAccusationDialouge5}
	edgeworthPanicDialouge1 := "Your honour! I am not going to deal with this child-like behaviour!"
	edgeworthPanicDialouge2 := "Your honour! The defence is not correct in this matter!"
	edgeworthPanicDialouge3 := "I object! Your arguments are not correct!"
	edgeworthPanicDialouge4 := "I object! There is a problem in your reasoning!"
	edgeworthPanicDialouge5 := "You are completely incorrect!"
	edgeworthPanicDialouges := []string{edgeworthPanicDialouge1, edgeworthPanicDialouge2, edgeworthPanicDialouge3, edgeworthPanicDialouge4, edgeworthPanicDialouge5}
	edgeworthDefaultDialouge1 := "Go on."
	edgeworthDefaultDialouge2 := "Go ahead."
	edgeworthDefaultDialouge3 := "State your name."
	edgeworthDefaultDialouge4 := "Please go on."
	edgeworthDefaultDialouge5 := "Ok."
	edgeworthDefaultDialouges := []string{edgeworthDefaultDialouge1, edgeworthDefaultDialouge2, edgeworthDefaultDialouge3, edgeworthDefaultDialouge4, edgeworthDefaultDialouge5}

	JudgeAccusationDialouge1 := "What is going on here?"
	JudgeAccusationDialouge2 := "Order! Order in the court!"
	JudgeAccusationDialouge3 := "I said order!"
	JudgeAccusationDialouge4 := "Will the defence please refrain from taunting the witness!"
	JudgeAccusationDialouge5 := "Will you please stop talking for once!"
	JudgeAccusationDialouge6 := "Order! Will I please have some order!"
	JudgeAccusationDialouges := []string{JudgeAccusationDialouge1, JudgeAccusationDialouge2, JudgeAccusationDialouge3, JudgeAccusationDialouge4, JudgeAccusationDialouge5, JudgeAccusationDialouge6}
	JudgeDefaultDialouge1 := "Witness, please state your name."
	JudgeDefaultDialouge2 := "Please, no flash photography in the courtroom!"
	JudgeDefaultDialouge3 := "Where am I again? Sorry, I dozed off there."
	JudgeDefaultDialouge4 := "ZZZZ...ZZZZ..."
	JudgeDefaultDialouge5 := "There is nothing more to say."
	JudgeDefaultDialouges := []string{JudgeDefaultDialouge1, JudgeDefaultDialouge2, JudgeDefaultDialouge3, JudgeDefaultDialouge4, JudgeDefaultDialouge5}
	JudgePanicDialouge1 := "That is highly unusual."
	JudgePanicDialouge2 := "Is that true?"
	JudgePanicDialouge3 := "Really?"
	JudgePanicDialouge4 := "I didn't expect that!"
	JudgePanicDialouge5 := "My goodness!"
	JudgePanicDialouge6 := "Wow!"
	JudgePanicDialouges := []string{JudgePanicDialouge1, JudgePanicDialouge2, JudgePanicDialouge3, JudgePanicDialouge4, JudgePanicDialouge5, JudgePanicDialouge6}

	WitnessDefaultDialouge1 := "I saw him at the boat house."
	WitnessDefaultDialouge2 := "They were not there!"
	WitnessDefaultDialouge3 := "She had disappeared!"
	WitnessDefaultDialouge4 := "The boat was not there!"
	WitnessDefaultDialouge5 := "Someone has been murdered!"
	WitnessDefaultDialouge6 := "The shop was closed!"
	WitnessDefaultDialouge7 := "There were no footprints."
	WitnessDefaultDialouge8 := "He flew over my head!"
	WitnessDefaultDialouge9 := "A brick fell from the sky!"
	WitnessDefaultDialouge10 := "Something had to be done."
	WitnessDefaultDialouge11 := "There really was no one there."
	WitnessDefaultDialouge12 := "I'm being totally honest."
	WitnessDefaultDialouges := []string{WitnessDefaultDialouge1, WitnessDefaultDialouge2, WitnessDefaultDialouge3, WitnessDefaultDialouge4, WitnessDefaultDialouge5, WitnessDefaultDialouge6, WitnessDefaultDialouge7, WitnessDefaultDialouge8, WitnessDefaultDialouge9, WitnessDefaultDialouge10, WitnessDefaultDialouge11, WitnessDefaultDialouge12}

	if frame.character.characterType == "D" {
		if frame.character.expression.Emotion == "Confident" {
			frame.character.speech.text = RandomInDialougeList(pheonixConfidentDialouges)
		}
		if frame.character.expression.Emotion == "Sad" {
			frame.character.speech.text = RandomInDialougeList(pheonixSadDialouges)
		}
		if frame.character.expression.Emotion == "Panic" {
			frame.character.speech.text = RandomInDialougeList(pheonixPanicDialouges)
		}
		if frame.character.expression.Emotion == "Accusation" {
			frame.character.speech.text = RandomInDialougeList(pheonixAccusationDialouges)
		}
	}
	if frame.character.characterType == "P" {
		if frame.character.expression.Emotion == "Accusation" {
			frame.character.speech.text = RandomInDialougeList(edgeworthAccusationDialouges)
		}
		if frame.character.expression.Emotion == "Panic" {
			frame.character.speech.text = RandomInDialougeList(edgeworthPanicDialouges)
		}
		if frame.character.expression.Emotion == "Default" {
			frame.character.speech.text = RandomInDialougeList(edgeworthDefaultDialouges)
		}
		if frame.character.expression.Emotion == "Confident" {
			frame.character.speech.text = RandomInDialougeList(edgeworthConfidentDialouges)
		}
	}
	if frame.character.characterType == "J" {
		if frame.character.expression.Emotion == "Accusation" {
			frame.character.speech.text = RandomInDialougeList(JudgeAccusationDialouges)
		}
		if frame.character.expression.Emotion == "Panic" {
			frame.character.speech.text = RandomInDialougeList(JudgePanicDialouges)
		}
		if frame.character.expression.Emotion == "Default" {
			frame.character.speech.text = RandomInDialougeList(JudgeDefaultDialouges)
		}
	}
	if frame.character.characterType == "W" {
		if frame.character.expression.Emotion == "Default" {
			frame.character.speech.text = RandomInDialougeList(WitnessDefaultDialouges)
		}
	}
	return frame
}

func RandomInDialougeList(list []string) string {
	random := rand.Intn(len(list))
	return list[random]
}

func RandomInExpressionList(list []Expression) Expression {
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
			frame.character.expression = RandomInExpressionList(PheonixExpressionsConfident)
		}
		if scene.emotionpattern.emotionTypes[frame.id] == "Panic" {
			frame.character.expression = RandomInExpressionList(PheonixExpressionsPanic)
		}
		if scene.emotionpattern.emotionTypes[frame.id] == "Accusation" {
			frame.character.expression = RandomInExpressionList(PheonixExpressionsAccusation)
		}
		if scene.emotionpattern.emotionTypes[frame.id] == "Sad" {
			frame.character.expression = RandomInExpressionList(PheonixExpressionsSad)
		}
	}
	if frame.character.characterType == "P" {
		if scene.emotionpattern.emotionTypes[frame.id] == "Confident" {
			frame.character.expression = RandomInExpressionList(EdgeworthExpressionsConfident)
		}
		if scene.emotionpattern.emotionTypes[frame.id] == "Panic" {
			frame.character.expression = RandomInExpressionList(EdgeworthExpressionsPanic)
		}
		if scene.emotionpattern.emotionTypes[frame.id] == "Accusation" {
			frame.character.expression = RandomInExpressionList(EdgeworthExpressionsAccusation)
		}
		if scene.emotionpattern.emotionTypes[frame.id] == "Default" {
			frame.character.expression = RandomInExpressionList(EdgeworthExpressionsDefault)
		}
	}
	if frame.character.characterType == "J" {
		if scene.emotionpattern.emotionTypes[frame.id] == "Panic" {
			frame.character.expression = RandomInExpressionList(JudgeExpressionsPanic)
		}
		if scene.emotionpattern.emotionTypes[frame.id] == "Accusation" {
			frame.character.expression = RandomInExpressionList(JudgeExpressionsAccusation)
		}
		if scene.emotionpattern.emotionTypes[frame.id] == "Default" {
			frame.character.expression = RandomInExpressionList(JudgeExpressionsDefault)
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
