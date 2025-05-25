package scenes

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"rpg-tutorial/animations"
	"rpg-tutorial/camera"
	"rpg-tutorial/components"
	"rpg-tutorial/constants"
	"rpg-tutorial/entities"
	"rpg-tutorial/spritesheet"
	"rpg-tutorial/tilemap"
	"rpg-tutorial/tileset"
)

type GameScene struct {
	scene
	player            *entities.Player
	playerSpriteSheet *spritesheet.SpriteSheet
	enemies           []*entities.Enemy
	potions           []*entities.Potion
	tilemapJSON       *tilemap.TilemapJSON
	tilesets          []tileset.Tileset
	tilemapImg        *ebiten.Image
	cam               *camera.Camera
	colliders         []image.Rectangle
}

var _ Scene = (*GameScene)(nil)

func NewGameScene() *GameScene {
	return &GameScene{
		player:            nil,
		playerSpriteSheet: nil,
		enemies:           make([]*entities.Enemy, 0),
		potions:           make([]*entities.Potion, 0),
		tilemapJSON:       nil,
		tilesets:          nil,
		tilemapImg:        nil,
		cam:               nil,
		colliders:         make([]image.Rectangle, 0),
		scene: scene{
			loaded: false,
			id:     GameSceneId,
			next:   GameSceneId,
		},
	}
}

func (s *GameScene) IsLoaded() bool {
	return s.loaded
}

func (s *GameScene) Draw(screen *ebiten.Image) {
	// fill the screen with a nice sky color
	screen.Fill(color.RGBA{120, 180, 255, 255})

	opts := ebiten.DrawImageOptions{}

	// loop over the layers
	for layerIndex, layer := range s.tilemapJSON.Layers {
		// loop over the tiles in the layer data
		for index, id := range layer.Data {

			if id == 0 {
				continue
			}

			// get the tile position of the tile
			x := index % layer.Width
			y := index / layer.Width

			// convert the tile position to pixel position
			x *= constants.Tilesize
			y *= constants.Tilesize

			img := s.tilesets[layerIndex].Img(id)

			opts.GeoM.Translate(float64(x), float64(y))

			opts.GeoM.Translate(0.0, -(float64(img.Bounds().Dy()) + constants.Tilesize))

			opts.GeoM.Translate(s.cam.X, s.cam.Y)

			screen.DrawImage(img, &opts)

			// reset the opts for the next tile
			opts.GeoM.Reset()
		}
	}

	// set the translation of our drawImageOptions to the player's position
	opts.GeoM.Translate(s.player.X, s.player.Y)
	opts.GeoM.Translate(s.cam.X, s.cam.Y)

	playerFrame := 0
	activeAnim := s.player.ActiveAnimation(int(s.player.Dx), int(s.player.Dy))
	if activeAnim != nil {
		playerFrame = activeAnim.Frame()
	}

	// draw the player
	screen.DrawImage(
		// grab a subimage of the spritesheet
		s.player.Img.SubImage(
			s.playerSpriteSheet.Rect(playerFrame),
		).(*ebiten.Image),
		&opts,
	)

	opts.GeoM.Reset()

	for _, sprite := range s.enemies {
		opts.GeoM.Translate(sprite.X, sprite.Y)
		opts.GeoM.Translate(s.cam.X, s.cam.Y)

		screen.DrawImage(
			sprite.Img.SubImage(
				image.Rect(0, 0, constants.Tilesize, constants.Tilesize),
			).(*ebiten.Image),
			&opts,
		)

		opts.GeoM.Reset()
	}

	opts.GeoM.Reset()

	for _, sprite := range s.potions {
		opts.GeoM.Translate(sprite.X, sprite.Y)
		opts.GeoM.Translate(s.cam.X, s.cam.Y)

		screen.DrawImage(
			sprite.Img.SubImage(
				image.Rect(0, 0, constants.Tilesize, constants.Tilesize),
			).(*ebiten.Image),
			&opts,
		)

		opts.GeoM.Reset()
	}

	for _, collider := range s.colliders {
		vector.StrokeRect(
			screen,
			float32(collider.Min.X)+float32(s.cam.X),
			float32(collider.Min.Y)+float32(s.cam.Y),
			float32(collider.Dx()),
			float32(collider.Dy()),
			1.0,
			color.RGBA{255, 0, 0, 255},
			true,
		)
	}
}

func (s *GameScene) Init() error {
	// load the image from file
	playerImg, _, err := ebitenutil.NewImageFromFile("assets/images/ninja.png")
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	// load the image from file
	skeletonImg, _, err := ebitenutil.NewImageFromFile("assets/images/skeleton.png")
	if err != nil {
		// handle error
		log.Fatal(err)
	}

	potionImg, _, err := ebitenutil.NewImageFromFile("assets/images/potion.png")
	if err != nil {
		// handle error
		log.Fatal(err)
	}

	tilemapImg, _, err := ebitenutil.NewImageFromFile("assets/images/TilesetFloor.png")
	if err != nil {
		// handle error
		log.Fatal(err)
	}

	tilemapJSON, err := tilemap.NewTilemapJSON("assets/maps/spawn.json")
	if err != nil {
		log.Fatal(err)
	}

	tilesets, err := tilemapJSON.GenTilesets()
	if err != nil {
		log.Fatal(err)
	}

	playerSpriteSheet := spritesheet.NewSpriteSheet(4, 7, constants.Tilesize)

	s.player = &entities.Player{
		Sprite: &entities.Sprite{
			Img: playerImg,
			X:   50.0,
			Y:   50.0,
		},
		Health: 3,
		Animations: map[entities.PlayerState]*animations.Animation{
			entities.Up:    animations.NewAnimation(5, 13, 4, 20.0),
			entities.Down:  animations.NewAnimation(4, 12, 4, 20.0),
			entities.Left:  animations.NewAnimation(6, 14, 4, 20.0),
			entities.Right: animations.NewAnimation(7, 15, 4, 20.0),
		},
		CombatComp: components.NewBasicCombat(3, 1),
	}

	s.playerSpriteSheet = playerSpriteSheet

	s.enemies = []*entities.Enemy{
		{
			Sprite: &entities.Sprite{
				Img: skeletonImg,
				X:   100.0,
				Y:   100.0,
			},
			FollowsPlayer: true,
			CombatComp:    components.NewEnemyCombat(3, 1, 30),
		},
		{
			Sprite: &entities.Sprite{
				Img: skeletonImg,
				X:   150.0,
				Y:   50.0,
			},
			FollowsPlayer: false,
			CombatComp:    components.NewEnemyCombat(3, 1, 30),
		},
	}
	s.tilemapJSON = tilemapJSON
	s.tilesets = tilesets
	s.tilemapImg = tilemapImg
	s.cam = camera.NewCamera(0.0, 0.0)
	s.colliders = []image.Rectangle{
		image.Rect(100, 100, 116, 16),
	}
	s.potions = []*entities.Potion{
		{
			Sprite: &entities.Sprite{
				Img: potionImg,
				X:   210.0,
				Y:   100.0,
			},
			AmtHeal: 1.0,
		},
	}
	s.loaded = true
	return nil
}

func (s *GameScene) OnEnter() error {
	return nil
}

func (s *GameScene) OnExit() error {
	return nil
}

func (s *GameScene) Update() error {
	// "T" will temporarily trigger a game over
	if inpututil.IsKeyJustPressed(ebiten.KeyT) {
		s.next = GameOverSceneId
		return nil
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		s.next = ExitSceneId
		return nil
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		s.next = PauseSceneId
		return nil
	}
	// move the player based on keyboar input (left, right, up down)

	s.player.Dx = 0.0
	s.player.Dy = 0.0

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		s.player.Dx = -2
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		s.player.Dx = 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		s.player.Dy = -2
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		s.player.Dy = 2
	}

	s.player.X += s.player.Dx

	CheckCollisionHorizontal(s.player.Sprite, s.colliders)

	s.player.Y += s.player.Dy

	CheckCollisionVertical(s.player.Sprite, s.colliders)

	activeAnim := s.player.ActiveAnimation(int(s.player.Dx), int(s.player.Dy))
	if activeAnim != nil {
		activeAnim.Update()
	}

	// add behavior to the enemies
	for _, sprite := range s.enemies {

		sprite.Dx = 0.0
		sprite.Dy = 0.0

		if sprite.FollowsPlayer {
			if sprite.X < s.player.X {
				sprite.Dx += 1
			} else if sprite.X > s.player.X {
				sprite.Dx -= 1
			}
			if sprite.Y < s.player.Y {
				sprite.Dy += 1
			} else if sprite.Y > s.player.Y {
				sprite.Dy -= 1
			}
		}

		sprite.X += sprite.Dx

		CheckCollisionHorizontal(sprite.Sprite, s.colliders)

		sprite.Y += sprite.Dy

		CheckCollisionVertical(sprite.Sprite, s.colliders)
	}

	clicked := inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0)
	cX, cY := ebiten.CursorPosition()
	cX -= int(s.cam.X)
	cY -= int(s.cam.Y)
	s.player.CombatComp.Update()
	pRect := image.Rect(
		int(s.player.X),
		int(s.player.Y),
		int(s.player.X)+constants.Tilesize,
		int(s.player.Y)+constants.Tilesize,
	)

	deadEnemies := make(map[int]struct{})
	for index, enemy := range s.enemies {
		enemy.CombatComp.Update()
		rect := image.Rect(
			int(enemy.X),
			int(enemy.Y),
			int(enemy.X)+constants.Tilesize,
			int(enemy.Y)+constants.Tilesize,
		)

		// if enemy overlaps player
		if rect.Overlaps(pRect) {
			if enemy.CombatComp.Attack() {
				s.player.CombatComp.Damage(enemy.CombatComp.AttackPower())
				fmt.Println(
					fmt.Sprintf("player damaged. health: %d\n", s.player.CombatComp.Health()),
				)
				if s.player.CombatComp.Health() <= 0 {
					fmt.Println("player has died!")
				}
			}
		}

		// is cursor in rect?
		if cX > rect.Min.X && cX < rect.Max.X && cY > rect.Min.Y && cY < rect.Max.Y {
			if clicked &&
				math.Sqrt(
					math.Pow(
						float64(cX)-s.player.X+(constants.Tilesize/2),
						2,
					)+math.Pow(
						float64(cY)-s.player.Y+(constants.Tilesize/2),
						2,
					),
				) < constants.Tilesize*5 {
				fmt.Println("damaging enemy")
				enemy.CombatComp.Damage(s.player.CombatComp.AttackPower())

				if enemy.CombatComp.Health() <= 0 {
					deadEnemies[index] = struct{}{}
					fmt.Println("enemy has been eliminated")
				}
			}
		}
	}
	if len(deadEnemies) > 0 {
		newEnemies := make([]*entities.Enemy, 0)
		for index, enemy := range s.enemies {
			if _, exists := deadEnemies[index]; !exists {
				newEnemies = append(newEnemies, enemy)
			}
		}
		s.enemies = newEnemies
	}

	// handle simple potion functionality
	// for _, potion := range s.potions {
	// 	if s.player.X > potion.X {
	// 		s.player.Health += potion.AmtHeal
	// 		fmt.Printf("Picked up potion! Health: %d\n", s.player.Health)
	// 	}
	// }

	s.cam.FollowTarget(s.player.X+8, s.player.Y+8, 320, 240)
	s.cam.Constrain(
		float64(s.tilemapJSON.Layers[0].Width)*constants.Tilesize,
		float64(s.tilemapJSON.Layers[0].Height)*constants.Tilesize,
		320,
		240,
	)

	s.next = GameSceneId
	return nil
}

func (s *GameScene) ID() SceneId {
	return s.id
}

func (s *GameScene) Next() SceneId {
	return s.next
}

func (s *GameScene) SetFreezeFrame(screen *ebiten.Image) {
	s.freezeFrame = screen
}

func CheckCollisionHorizontal(sprite *entities.Sprite, colliders []image.Rectangle) {
	for _, collider := range colliders {
		if collider.Overlaps(
			image.Rect(
				int(sprite.X),
				int(sprite.Y),
				int(sprite.X)+constants.Tilesize,
				int(sprite.Y)+constants.Tilesize,
			),
		) {
			if sprite.Dx > 0.0 {
				sprite.X = float64(collider.Min.X) - constants.Tilesize
			} else if sprite.Dx < 0.0 {
				sprite.X = float64(collider.Max.X)
			}
		}
	}
}

func CheckCollisionVertical(sprite *entities.Sprite, colliders []image.Rectangle) {
	for _, collider := range colliders {
		if collider.Overlaps(
			image.Rect(
				int(sprite.X),
				int(sprite.Y),
				int(sprite.X)+constants.Tilesize,
				int(sprite.Y)+constants.Tilesize,
			),
		) {
			if sprite.Dy > 0.0 {
				sprite.Y = float64(collider.Min.Y) - constants.Tilesize
			} else if sprite.Dy < 0.0 {
				sprite.Y = float64(collider.Max.X)
			}
		}
	}
}
