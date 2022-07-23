package gen

import (
	"embed"
	"errors"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"math/rand"
	"os"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

//go:embed images
var images embed.FS

//go:embed fonts
var fonts embed.FS

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

func Generate(seed int) (frames []image.Image, objection image.Image, objectionLocation int, err error) {
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

	expressionThinkPhoenix := Expression{ID: "think", Emotion: "Confident"}
	expressionConfidentPhoenix := Expression{ID: "confident", Emotion: "Confident"}
	expressionReadPhoenix := Expression{ID: "read", Emotion: "Confident"}
	expressionPointPhoenix := Expression{ID: "point", Emotion: "Accusation"}
	expressionDeskSlamPhoenix := Expression{ID: "deskSlam", Emotion: "Accusation"}
	expressionCorneredPhoenix := Expression{ID: "cornered", Emotion: "Panic"}
	expressionSillyPhoenix := Expression{ID: "silly", Emotion: "Panic"}
	expressionBrekdownPhoenix := Expression{ID: "brekdown", Emotion: "Panic"}
	expressionCoffePhoenix := Expression{ID: "coffe", Emotion: "Sad"}
	PhoenixExpressions := []Expression{expressionThinkPhoenix, expressionConfidentPhoenix, expressionReadPhoenix, expressionPointPhoenix, expressionDeskSlamPhoenix, expressionCorneredPhoenix, expressionSillyPhoenix, expressionBrekdownPhoenix, expressionCoffePhoenix}

	PhoenixExpressionsConfident := []Expression{}
	PhoenixExpressionsPanic := []Expression{}
	PhoenixExpressionsAccusation := []Expression{}
	PhoenixExpressionsSad := []Expression{}

	for i := 0; i < len(PhoenixExpressions); i++ {
		if PhoenixExpressions[i].Emotion == "Confident" {
			PhoenixExpressionsConfident = append(PhoenixExpressionsConfident, PhoenixExpressions[i])
		}
		if PhoenixExpressions[i].Emotion == "Panic" {
			PhoenixExpressionsPanic = append(PhoenixExpressionsPanic, PhoenixExpressions[i])
		}
		if PhoenixExpressions[i].Emotion == "Accusation" {
			PhoenixExpressionsAccusation = append(PhoenixExpressionsAccusation, PhoenixExpressions[i])
		}
		if PhoenixExpressions[i].Emotion == "Sad" {
			PhoenixExpressionsSad = append(PhoenixExpressionsSad, PhoenixExpressions[i])
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
	judgeExpressions := []Expression{expressionStandJudge, expressionAngryJudge, expressionSurpriseJudge}

	judgeExpressionsDefault := []Expression{}
	judgeExpressionsAccusation := []Expression{}
	judgeExpressionsPanic := []Expression{}

	for i := 0; i < len(judgeExpressions); i++ {
		if judgeExpressions[i].Emotion == "Default" {
			judgeExpressionsDefault = append(judgeExpressionsDefault, judgeExpressions[i])
		}
		if judgeExpressions[i].Emotion == "Accusation" {
			judgeExpressionsAccusation = append(judgeExpressionsAccusation, judgeExpressions[i])
		}
		if judgeExpressions[i].Emotion == "Panic" {
			judgeExpressionsPanic = append(judgeExpressionsPanic, judgeExpressions[i])
		}

	}

	character1 := Character{characterType: "D", name: "Phoenix Wright"}
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

	scene := Scene{seed: seed, frames: []Frame{frame1, frame2, frame3, frame4}}
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

	frame1 = setExpressions(frame1, scene, PhoenixExpressionsConfident, EdgeworthExpressionsConfident, EdgeworthExpressionsDefault, judgeExpressionsDefault, judgeExpressionsPanic, judgeExpressionsAccusation, PhoenixExpressionsAccusation, PhoenixExpressionsSad, EdgeworthExpressionsPanic, EdgeworthExpressionsAccusation, judgeExpressionsAccusation)
	frame2 = setExpressions(frame2, scene, PhoenixExpressionsConfident, EdgeworthExpressionsConfident, EdgeworthExpressionsDefault, judgeExpressionsDefault, judgeExpressionsPanic, judgeExpressionsAccusation, PhoenixExpressionsAccusation, PhoenixExpressionsSad, EdgeworthExpressionsPanic, EdgeworthExpressionsAccusation, judgeExpressionsAccusation)
	frame3 = setExpressions(frame3, scene, PhoenixExpressionsConfident, EdgeworthExpressionsConfident, EdgeworthExpressionsDefault, judgeExpressionsDefault, judgeExpressionsPanic, judgeExpressionsAccusation, PhoenixExpressionsAccusation, PhoenixExpressionsSad, EdgeworthExpressionsPanic, EdgeworthExpressionsAccusation, judgeExpressionsAccusation)
	frame4 = setExpressions(frame4, scene, PhoenixExpressionsConfident, EdgeworthExpressionsConfident, EdgeworthExpressionsDefault, judgeExpressionsDefault, judgeExpressionsPanic, judgeExpressionsAccusation, PhoenixExpressionsAccusation, PhoenixExpressionsSad, EdgeworthExpressionsPanic, EdgeworthExpressionsAccusation, judgeExpressionsAccusation)

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
	// fmt.Println(scene)
	if scene.objection.objectionActive {
		objectionPath, err := getObjectionFramePath(scene.objection)
		if err != nil {
			fmt.Println("Error getting objection path: ", err)
			return nil, nil, 0, err
		}
		objection = GetImage(objectionPath)
	}
	for i := 0; i < len(scene.frames); i++ {
		bgPath, fgPath, characterPath := getPathsFromFrame(scene.frames[i])
		bgImage, fgImage, characterImage, err := GetImages(bgPath, fgPath, characterPath)
		if err != nil {
			fmt.Println("Error getting images: ", err)
			fmt.Println("Paths: ", bgPath, fgPath, characterPath)
			return nil, nil, 0, err
		}
		finalImage, err := AssembleImage(bgImage, fgImage, characterImage)
		if err != nil {
			fmt.Println("Error assembling image: ", err)
			return nil, nil, 0, err
		}

		boxImage := GetImage("images/speech_box/box.png")
		textBoxAddedImage, err := combineImages(finalImage, boxImage)
		if err != nil {
			fmt.Println("Error assembling image: ", err)
			return nil, nil, 0, err
		}
		if scene.frames[i].character.characterType == "C" {
			textBoxAddedImage = finalImage
		}
		finalTextAddedImage, err := addLabel(textBoxAddedImage, 0, 0, "hsdkjhdakhdkjsahdkjahdkjahdkjhadkjasd")
		if err != nil {
			fmt.Println("Error adding label: ", err)
			return nil, nil, 0, err
		}
		frames = append(frames, finalTextAddedImage)
	}
	objectionLocation = scene.objection.objectionLocation
	return
}

func getSeedFromBytes(bytes []byte) (seed int) {
	seed = 1
	for i := 0; i < len(bytes); i++ {
		seed *= int(bytes[i])
	}
	return seed
}

func loadFontFaceReader(fontBytes []byte, points float64) (font.Face, error) {
	f, err := truetype.Parse(fontBytes)
	if err != nil {
		return nil, err
	}
	face := truetype.NewFace(f, &truetype.Options{
		Size: points,
		// Hinting: font.HintingFull,
	})
	return face, nil
}

func addLabel(img image.Image, x, y int, label string) (outimage image.Image, err error) {
	var w = img.Bounds().Dx()
	var h = img.Bounds().Dy()
	dc := gg.NewContext(w, h)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	fontBytes, err := fonts.ReadFile("fonts/Roboto-Regular.ttf")
	if err != nil {
		return nil, err
	}
	face, err := loadFontFaceReader(fontBytes, 24)
	if err != nil {
		return nil, err
	}
	dc.SetFontFace(face)

	dc.DrawImage(img, 0, 0)
	dc.DrawStringAnchored(label, float64(x), float64(y), 0.5, 0.5)
	dc.Clip()
	outimage = dc.Image()
	return outimage, nil
}

func getObjectionFramePath(objection Objection) (path string, err error) {
	if objection.objectionType == "Objection" {
		path = "images/speech_bub/objection.png"
	} else if objection.objectionType == "TakeThat" {
		path = "images/speech_bub/takethat.png"
	} else if objection.objectionType == "HoldIt" {
		path = "images/speech_bub/hold_it.png"
	} else {
		return "", errors.New("no objections active")
	}
	return path, nil
}

func getPathsFromFrame(frame Frame) (bg string, fg string, character string) {
	if frame.character.characterType == "D" {
		bg = "images/background/defence.jpg"
		fg = "images/foreground/defence.png"
		if frame.character.expression.ID == "think" {
			character = "images/poses/defence/think.png"
		}
		if frame.character.expression.ID == "angry" {
			character = "images/poses/defence/desk_slam.png"
		}
		if frame.character.expression.ID == "confident" {
			character = "images/poses/defence/confident.png"
		}
		if frame.character.expression.ID == "read" {
			character = "images/poses/defence/read.png"
		}
		if frame.character.expression.ID == "point" {
			character = "images/poses/defence/point.png"
		}
		if frame.character.expression.ID == "deskSlam" {
			character = "images/poses/defence/desk_slam.png"
		}
		if frame.character.expression.ID == "silly" {
			character = "images/poses/defence/silly.png"
		}
		if frame.character.expression.ID == "breakdown" {
			character = "images/poses/defence/breakdown.png"
		}
		if frame.character.expression.ID == "coffe" {
			character = "images/poses/defence/coffe.png"
		}
	}
	if frame.character.characterType == "P" {
		bg = "images/background/prosecution.jpg"
		fg = "images/foreground/prosecution.png"
		if frame.character.expression.ID == "arms" {
			character = "images/poses/prosecution/arms.png"
		}
		if frame.character.expression.ID == "confident" {
			character = "images/poses/prosecution/confident.png"
		}
		if frame.character.expression.ID == "read" {
			character = "images/poses/prosecution/read.png"
		}
		if frame.character.expression.ID == "cornered" {
			character = "images/poses/prosecution/cornered.png"
		}
		if frame.character.expression.ID == "deskSlam" {
			character = "images/poses/prosecution/deskslam.png"
		}
		if frame.character.expression.ID == "point" {
			character = "images/poses/prosecution/point.png"
		}
		if frame.character.expression.ID == "stand" {
			character = "images/poses/prosecution/stand.png"
		}
	}
	if frame.character.characterType == "J" {
		bg = "images/background/judge.jpg"
		fg = "images/foreground/judge.png"
		if frame.character.expression.ID == "stand" {
			character = "images/poses/judge/stand.png"
		}
		if frame.character.expression.ID == "angry" {
			character = "images/poses/judge/angry.png"
		}
		if frame.character.expression.ID == "surprise" {
			character = "images/poses/judge/surprise.png"
		}
	}
	if frame.character.characterType == "W" {
		bg = "images/background/witness.jpg"
		fg = "images/foreground/witness.png"
		if frame.character.name == "Cody Hackins" {
			character = "images/poses/witness/cody.png"
		}
		if frame.character.name == "De Vasquez" {
			character = "images/poses/witness/des.png"
		}
		if frame.character.name == "Detective Gumshoe" {
			character = "images/poses/witness/detective.png"
		}
		if frame.character.name == "Godot" {
			character = "images/poses/witness/godot.png"
		}
		if frame.character.name == "Larry Butz" {
			character = "images/poses/witness/larry.png"
		}
		if frame.character.name == "Yanni Yogi" {
			character = "images/poses/witness/yanni.png"
		}
	}
	if frame.character.characterType == "C" {
		bg = "images/background/gallery.jpg"
		fg = "images/foreground/gallery.png"
		if frame.character.name == "Crowd" {
			character = "images/poses/crowd/blank.png"
		}
	}
	return bg, fg, character
}

func combineImages(img1 image.Image, img2 image.Image) (image.Image, error) {
	rgba := image.NewRGBA(img1.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img1, image.Point{0, 0}, draw.Over)
	draw.Draw(rgba, rgba.Bounds(), img2, image.Point{0, 0}, draw.Over)
	return rgba, nil
}

func AssembleImage(bg image.Image, fg image.Image, character image.Image) (img image.Image, err error) {
	bgchar, err := combineImages(bg, character)
	if err != nil {
		return nil, err
	}
	fgbgchar, err := combineImages(bgchar, fg)
	if err != nil {
		return nil, err
	}
	return fgbgchar, err
}

func GetImage(path string) (img image.Image) {
	imgFile1, err := images.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	img, err = png.Decode(imgFile1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return img
}

func GetImages(bgpath, fgpath, characterpath string) (bg image.Image, fg image.Image, character image.Image, err error) {
	imgFile1, err := images.Open(bgpath)
	if err != nil {
		return nil, nil, nil, err
	}
	imgFile2, err := images.Open(fgpath)
	if err != nil {
		return nil, nil, nil, err
	}
	imgFile3, err := images.Open(characterpath)
	if err != nil {
		return nil, nil, nil, err
	}
	bg, err = jpeg.Decode(imgFile1)
	if err != nil {
		return nil, nil, nil, err
	}
	fg, err = png.Decode(imgFile2)
	if err != nil {
		return nil, nil, nil, err
	}
	character, err = png.Decode(imgFile3)
	if err != nil {
		return nil, nil, nil, err
	}
	return bg, fg, character, nil
}

func randomDialouge(frame Frame) Frame {
	phoenixConfidentDialouge1 := "That is correct."
	phoenixConfidentDialouge2 := "I have evidence to back that up!"
	phoenixConfidentDialouge3 := "There is no problem here."
	phoenixConfidentDialouge4 := "There is no problem with that."
	phoenixConfidentDialouge5 := "That makes no difference."
	phoenixConfidentDialouge6 := "Just as I thought!"
	phoenixConfidentDialouge7 := "I was right!"
	phoenixConfidentDialouge8 := "I knew it!"
	phoenixConfidentDialouges := []string{phoenixConfidentDialouge1, phoenixConfidentDialouge2, phoenixConfidentDialouge3, phoenixConfidentDialouge4, phoenixConfidentDialouge5, phoenixConfidentDialouge6, phoenixConfidentDialouge7, phoenixConfidentDialouge8}
	phoenixAccusationDialouge1 := "I have evidence that proves the defendant was not at the scene of the crime!"
	phoenixAccusationDialouge2 := "You were not inside the boat!"
	phoenixAccusationDialouge3 := "That is not correct!"
	phoenixAccusationDialouge4 := "You are wrong!"
	phoenixAccusationDialouge5 := "The prosection is lying!"
	phoenixAccusationDialouge6 := "There is a vital piece of evidence missing!"
	phoenixAccusationDialouge7 := "You are not correct!"
	phoenixAccusationDialouge8 := "There is not enough evidence!"
	phoenixAccusationDialouges := []string{phoenixAccusationDialouge1, phoenixAccusationDialouge2, phoenixAccusationDialouge3, phoenixAccusationDialouge4, phoenixAccusationDialouge5, phoenixAccusationDialouge6, phoenixAccusationDialouge7, phoenixAccusationDialouge8}
	phoenixSadDialouge1 := "Ouch! That hurt!"
	phoenixSadDialouge2 := "Oww.. that burns!"
	phoenixSadDialouge3 := "Hey!"
	phoenixSadDialouges := []string{phoenixSadDialouge1, phoenixSadDialouge2, phoenixSadDialouge3}
	phoenixPanicDialouge1 := "I.. Uhhh Ummm (How do I get out of this one?)"
	phoenixPanicDialouge2 := "(Urk.)"
	phoenixPanicDialouge3 := "(uh-oh)"
	phoenixPanicDialouge4 := "(oops)"
	phoenixPanicDialouge5 := "(oh no!)"
	phoenixPanicDialouge6 := "(How do I get out of this?)"
	phoenixPanicDialouge7 := "(This isn't going well!)"
	phoenixPanicDialouge8 := "Uhhhh... Ummmm... (oh-no!)"
	phoenixPanicDialouges := []string{phoenixPanicDialouge1, phoenixPanicDialouge2, phoenixPanicDialouge3, phoenixPanicDialouge4, phoenixPanicDialouge5, phoenixPanicDialouge6, phoenixPanicDialouge7, phoenixPanicDialouge8}

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
			frame.character.speech.text = RandomInDialogueList(phoenixConfidentDialouges)
		}
		if frame.character.expression.Emotion == "Sad" {
			frame.character.speech.text = RandomInDialogueList(phoenixSadDialouges)
		}
		if frame.character.expression.Emotion == "Panic" {
			frame.character.speech.text = RandomInDialogueList(phoenixPanicDialouges)
		}
		if frame.character.expression.Emotion == "Accusation" {
			frame.character.speech.text = RandomInDialogueList(phoenixAccusationDialouges)
		}
	}
	if frame.character.characterType == "P" {
		if frame.character.expression.Emotion == "Accusation" {
			frame.character.speech.text = RandomInDialogueList(edgeworthAccusationDialouges)
		}
		if frame.character.expression.Emotion == "Panic" {
			frame.character.speech.text = RandomInDialogueList(edgeworthPanicDialouges)
		}
		if frame.character.expression.Emotion == "Default" {
			frame.character.speech.text = RandomInDialogueList(edgeworthDefaultDialouges)
		}
		if frame.character.expression.Emotion == "Confident" {
			frame.character.speech.text = RandomInDialogueList(edgeworthConfidentDialouges)
		}
	}
	if frame.character.characterType == "J" {
		if frame.character.expression.Emotion == "Accusation" {
			frame.character.speech.text = RandomInDialogueList(JudgeAccusationDialouges)
		}
		if frame.character.expression.Emotion == "Panic" {
			frame.character.speech.text = RandomInDialogueList(JudgePanicDialouges)
		}
		if frame.character.expression.Emotion == "Default" {
			frame.character.speech.text = RandomInDialogueList(JudgeDefaultDialouges)
		}
	}
	if frame.character.characterType == "W" {
		if frame.character.expression.Emotion == "Default" {
			frame.character.speech.text = RandomInDialogueList(WitnessDefaultDialouges)
		}
	}
	return frame
}

func RandomInDialogueList(list []string) string {
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

func setExpressions(frame Frame, scene Scene, PhoenixExpressionsConfident []Expression, EdgeworthExpressionsConfident []Expression, EdgeworthExpressionsDefault []Expression, JudgeExpressionsDefault []Expression, JudgeExpressionsPanic []Expression, PhoenixExpressionsPanic []Expression, PhoenixExpressionsAccusation []Expression, PhoenixExpressionsSad []Expression, EdgeworthExpressionsPanic []Expression, EdgeworthExpressionsAccusation []Expression, JudgeExpressionsAccusation []Expression) Frame {
	if frame.character.characterType == "D" {
		if scene.emotionpattern.emotionTypes[frame.id] == "Confident" {
			frame.character.expression = RandomInExpressionList(PhoenixExpressionsConfident)
		}
		if scene.emotionpattern.emotionTypes[frame.id] == "Panic" {
			frame.character.expression = RandomInExpressionList(PhoenixExpressionsPanic)
		}
		if scene.emotionpattern.emotionTypes[frame.id] == "Accusation" {
			frame.character.expression = RandomInExpressionList(PhoenixExpressionsAccusation)
		}
		if scene.emotionpattern.emotionTypes[frame.id] == "Sad" {
			frame.character.expression = RandomInExpressionList(PhoenixExpressionsSad)
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
