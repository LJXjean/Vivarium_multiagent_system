package sprites

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/exp/rand"
	"image"
	"log"
	"time"
	"vivarium/ebiten/assets/images"
	"vivarium/organisme"
)

const (
	frameOX     = 0
	frameOY     = 0
	frameWidth  = 32
	frameHeight = 32
	frameCount  = 8

	framePerSwitch = 10 // It decides the speed of animation: the greater the slower
)

type SpriteState int

const (
	Idle SpriteState = iota
	Moving
	Attacking
	Dying
	Eating
	Fucking
)

type SpriteType int

const (
	Spider SpriteType = iota
	Snail
)

var SpriteMap = make(map[int]*Sprite)

// 用于存储每个生物agent的状态
type Sprite struct {
	X float64
	Y float64

	Id int

	image *ebiten.Image

	State        SpriteState
	IdleFrames   []*ebiten.Image
	MoveFrames   []*ebiten.Image
	AttackFrames []*ebiten.Image
	DieFrames    []*ebiten.Image

	frameIndex int

	Species string

	IsDead     bool
	DyingCount int
	IsDying    bool

	IsInsect bool

	// 昆虫特有的状态
	IsManger     bool
	IsReproduire bool
	IsSeDeplacer bool
	IsSeBattre   bool
	IsWinner     bool
	IsLooser     bool
	IsNormal     bool

	// 植物特有的状态
	NbParts int
}

// 每次update请求后，都会根据agent更新精灵状态，如果该id不在map中，自动生成精灵
func UpdateOrganisme(org organisme.Organisme) {
	switch o := org.(type) {
	case *organisme.Insecte:
		UpdateInsecte(o) // o 是 *organisme.Insecte 类型
		time.Sleep(time.Millisecond * 100)
	case *organisme.Plante:
		UpdatePlante(o)
		time.Sleep(time.Millisecond * 100)
	}
}

func UpdateInsecte(org *organisme.Insecte) {

	// 假设 Agent 有一个唯一的 ID
	spriteInfo := &Sprite{
		X:  float64(org.GetPosX()),
		Y:  float64(org.GetPosY()),
		Id: org.GetID(),
		// image *ebiten.Image 这里应该是赖子哥来赋值
		// State        SpriteState
		//IdleFrames   []*ebiten.Image
		//MoveFrames   []*ebiten.Image
		//AttackFrames []*ebiten.Image
		//DieFrames    []*ebiten.Image

		//frameIndex int
		Species:    org.GetEspece().String(),
		DyingCount: 0,
		IsDying:    org.GetEtat(),
		IsInsect:   true,

		// 昆虫特有的状态
		IsManger:     org.IsManger,
		IsReproduire: org.IsReproduire,
		IsSeDeplacer: org.IsSeDeplacer,
		IsSeBattre:   org.IsSeBattre,
		IsWinner:     org.IsWinner,
		IsLooser:     org.IsLooser,
		IsNormal:     org.IsNormal,

		// 植物特有的状态
		NbParts: 1,
	}
	SpriteMap[org.GetID()] = spriteInfo
}

func UpdatePlante(org *organisme.Plante) {

	// 假设 Agent 有一个唯一的 ID
	spriteInfo := &Sprite{
		X:  float64(org.GetPosX()),
		Y:  float64(org.GetPosY()),
		Id: org.GetID(),
		// image *ebiten.Image 这里应该是赖子哥来赋值
		// State        SpriteState
		//IdleFrames   []*ebiten.Image
		//MoveFrames   []*ebiten.Image
		//AttackFrames []*ebiten.Image
		//DieFrames    []*ebiten.Image

		//frameIndex int
		Species:    org.GetEspece().String(),
		DyingCount: 0,
		IsDying:    org.GetEtat(),
		IsInsect:   false,

		// 昆虫特有的状态
		IsManger:     false,
		IsReproduire: false,
		IsSeDeplacer: false,
		IsSeBattre:   false,
		IsWinner:     false,
		IsLooser:     false,
		IsNormal:     true,

		// 植物特有的状态
		NbParts: org.NbParts,
	}
	SpriteMap[org.GetID()] = spriteInfo
}

func (s *Sprite) Update() {

	// 如果精灵已死，不再更新
	if s.IsDead {
		return
	}

	// 处理正在死亡的逻辑
	if s.IsDying {
		// 此处执行死亡的渲染动画
		s.DyingCount++
		if s.DyingCount >= 20 {
			s.IsDead = true
			return
		}
	}

	// 更新精灵帧索引
	s.frameIndex++

	if s.IsNormal == false {
		// 如果是昆虫
		if s.IsInsect {
			if s.IsManger {
				// 执行与进食相关的逻辑 戴个恰饭图标
			}
			if s.IsReproduire {
				// 执行与繁殖相关的逻辑 戴个💗💗💗
			}
			if s.IsSeDeplacer {
				// 执行与移动相关的逻辑 戴个移动图标
			}
			if s.IsSeBattre {
				if s.IsWinner {
					// 执行胜利者的逻辑 戴个1s小王冠
				} else if s.IsLooser {
					// 执行失败者的逻辑 显示1sLoser
				} else {
					// 执行正常战斗的逻辑 戴个打架图标
				}
			}
		} else {
			// 如果是植物
			if s.NbParts > 0 {
				// 根据NbParts显示百分比图标
			}
		}
	} else {
		// 执行正常状态的逻辑 无图标状态
	}
}

func (s *Sprite) MoveTo(x, y float64) {
}

func (s *Sprite) Draw(screen *ebiten.Image, FrameIndex int) {
	var currentFrame *ebiten.Image

	if s.IsDead {
		// 如果精灵已死，不进行渲染
		return
	}

	if s.State == Moving {
		currentFrame = s.MoveFrames[(FrameIndex/framePerSwitch)%len(s.MoveFrames)]
	} else if s.State == Attacking {
		currentFrame = s.AttackFrames[(FrameIndex/framePerSwitch)%len(s.AttackFrames)]
	} else if s.State == Dying {
		currentFrame = s.DieFrames[(FrameIndex/framePerSwitch)%len(s.DieFrames)]
	} else if s.State == Idle {
		currentFrame = s.IdleFrames[(FrameIndex/framePerSwitch)%len(s.IdleFrames)]
	}

	// 应该还有Eating和Fucking的渲染？

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(s.X, s.Y)
	screen.DrawImage(currentFrame, op)
}

func loadFrames(img *ebiten.Image, frameCount, stateIdx int) []*ebiten.Image {
	frames := make([]*ebiten.Image, frameCount)
	for i := range frames {
		sX, sY := frameOX+i*frameWidth, frameOY+frameHeight*stateIdx
		frame := img.SubImage(image.Rect(sX, sY, sX+frameWidth, sY+frameHeight)).(*ebiten.Image)
		frames[i] = frame
	}
	return frames
}

func NewSpiderSprite(X, Y float64, state SpriteState) *Sprite {
	img, _, err := image.Decode(bytes.NewReader(images.Spider_png))
	if err != nil {
		log.Fatal(err)
	}
	spiderImage := ebiten.NewImageFromImage(img)

	s := &Sprite{
		image: spiderImage,
		X:     X,
		Y:     Y,
		State: state,
		Id:    rand.Intn(100000),

		frameIndex: 0,
	}

	s.IdleFrames = loadFrames(spiderImage, 5, 0)
	s.MoveFrames = loadFrames(spiderImage, 6, 1)
	s.AttackFrames = loadFrames(spiderImage, 9, 2)
	s.DieFrames = loadFrames(spiderImage, 9, 6)
	return s
}

func NewSnailSprite(X, Y float64) *Sprite {
	img, _, err := image.Decode(bytes.NewReader(images.Snail_png))
	if err != nil {
		log.Fatal(err)
	}
	snailImage := ebiten.NewImageFromImage(img)

	return &Sprite{
		image: snailImage,
		X:     X,
		Y:     Y,
	}
}
