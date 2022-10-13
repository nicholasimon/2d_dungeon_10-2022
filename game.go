package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// MARK: VAR VAR VAR VAR VAR VAR VAR VAR VAR VAR VAR VAR VAR VAR VAR VAR VAR VAR VAR VAR VAR
var (

	//INV
	invtabrec, invrec        rl.Rectangle
	inv, invequip            []xitem
	invon, invopen, invclose bool
	invemptynum              int
	invplayerstatstimer      = int32(15)
	selectedinv              = xitem{}
	selectedinvnum int

	//COLORS
	themecol1, themecol2 rl.Color

	//MSG
	msgtimer, msgdelay int32
	msgs               []string

	//MENUS
	infobarrec, infobarrec2 rl.Rectangle

	//PLAYER
	player     = xplayer{}
	pushweight int

	//LEVEL
	rooms          []xroom
	currentroom    int
	tilesize       = u1
	maxuw, maxuh   int
	borderwallrecs []xtile

	//GAME
	fps                       = int32(30)
	frames                    int
	borderrec, borderrecinner rl.Rectangle
	pause                     bool

	//INP
	mousev2, mousev2world rl.Vector2
	fadeblink             = float32(0.2)
	fadeblinkon           bool
	mousin                bool
	selpoint, selpointr   rl.Vector2
	selrec                rl.Rectangle

	//TEXT
	txtbase = int32(10)
	txtss   = txtbase / 2
	txts    = txtbase
	txtm    = txtbase * 2
	txtl    = txtbase * 3
	txtl2   = txtbase * 4
	txtxl   = txtbase * 6
	txtxl2  = txtbase * 8

	//TIMER
	time10, time20, time30 bool

	//DEV
	dev, dev2 bool

	//CAMERAS
	camera rl.Camera2D

	//SCREEN
	scrhf32, scrwf32 float32
	scrh, scrw       int32
	scrhint, scrwint int
	scrcnt           rl.Vector2

	//BLANKS
	blankv2  = rl.NewVector2(77777777777777777, 77777777777777777)
	blankint = 77777777777777777
	blankrec rl.Rectangle

	//MARK: IMAGES IMAGES IMAGES IMAGES IMAGES IMAGES IMAGES IMAGES IMAGES IMAGES
	imgs   rl.Texture2D
	origin = rl.NewVector2(0, 0)

	//TILE IMAGES
	wall1img        = rl.NewRectangle(0, 0, 16, 16)
	floortilesize   = tilesize
	floorimgs       []rl.Rectangle
	floorback       []xtile
	crateimg        = rl.NewRectangle(0, 32, 16, 16)
	blockmoveimg    = rl.NewRectangle(16, 32, 16, 16)
	blockpowerupimg = rl.NewRectangle(32, 32, 16, 16)

	//ROOM EXTRAS IMGS
	flamespoutimg = rl.NewRectangle(0, 48, 16, 16)

	//FX IMGS
	fireimg            = rl.NewRectangle(0, 836, 64, 64)
	yellowexplosionimg = rl.NewRectangle(0, 772, 64, 64)

	//PLAYER IMGS
	pwalkimgd = rl.NewRectangle(1330, 99, 32, 32)
	pwalkimgu = rl.NewRectangle(1330, 35, 32, 32)
	pwalkimgr = rl.NewRectangle(1330, 3, 32, 32)
	pwalkimgl = rl.NewRectangle(1330, 67, 32, 32)
	pidleimg  = rl.NewRectangle(1330, 136, 32, 32)
	patkimgr  = rl.NewRectangle(1330, 176, 32, 32)
	patkimgu  = rl.NewRectangle(1330, 208, 32, 32)
	patkimgl  = rl.NewRectangle(1330, 240, 32, 32)
	patkimgd  = rl.NewRectangle(1330, 272, 32, 32)

	portraitimgs []rl.Rectangle

	//UI ICONS
	invimg   = rl.NewRectangle(0, 64, 16, 16)
	closeimg = rl.NewRectangle(16, 66, 14, 14)

	//INV ICONS
	invweap1img   = rl.NewRectangle(33, 67, 12, 12)
	invweap2img   = rl.NewRectangle(47, 66, 16, 16)
	invweap3img   = rl.NewRectangle(64, 66, 16, 16)
	invweap4img   = rl.NewRectangle(81, 66, 16, 16)
	invglovesimg  = rl.NewRectangle(101, 66, 16, 16)
	invringimg    = rl.NewRectangle(123, 68, 14, 14)
	invjewelryimg = rl.NewRectangle(141, 67, 14, 14)
	invammoimg    = rl.NewRectangle(156, 66, 16, 16)
	invbootsimg   = rl.NewRectangle(174, 67, 16, 16)
	invhelmetimg  = rl.NewRectangle(195, 66, 16, 16)
	invarmorimg   = rl.NewRectangle(214, 67, 16, 16)

	//WEAPONS
	swordimgs []rl.Rectangle

	//UNITS FLOAT32
	ubase = float32(32)

	uq = ubase / 4
	uh = ubase / 2
	ue = ubase / 8

	u1  = ubase
	u1h = u1 + uh
	u1q = u1 + uq

	u2  = ubase * 2
	u2h = u2 + uh
	u2q = u2 + uq

	u3  = ubase * 3
	u3h = u3 + uh
	u3q = u3 + uq

	u4  = ubase * 4
	u4h = u4 + uh
	u4q = u4 + uq

	u5  = ubase * 5
	u5h = u5 + uh
	u5q = u5 + uq

	u6  = ubase * 6
	u6h = u6 + uh
	u6q = u6 + uq

	u7  = ubase * 7
	u7h = u7 + uh
	u7q = u7 + uq

	u8  = ubase * 8
	u8h = u8 + uh
	u8q = u8 + uq

	u9  = ubase * 9
	u9h = u9 + uh
	u9q = u9 + uq

	u10  = ubase * 10
	u10h = u10 + uh
	u10q = u10 + uq

	//UNITS INT32
	nbase = int32(ubase)

	nq = nbase / 4
	nh = nbase / 2
	ne = nbase / 8

	n1  = nbase
	n1h = n1 + nh
	n1q = n1 + nq

	n2  = nbase * 2
	n2h = n2 + nh
	n2q = n2 + nq

	n3  = nbase * 3
	n3h = n3 + nh
	n3q = n3 + nq

	n4  = nbase * 4
	n4h = n4 + nh
	n4q = n4 + nq

	n5  = nbase * 5
	n5h = n5 + nh
	n5q = n5 + nq

	n6  = nbase * 6
	n6h = n6 + nh
	n6q = n6 + nq

	n7  = nbase * 7
	n7h = n7 + nh
	n7q = n7 + nq

	n8  = nbase * 8
	n8h = n8 + nh
	n8q = n8 + nq

	n9  = nbase * 9
	n9h = n9 + nh
	n9q = n9 + nq

	n10  = nbase * 10
	n10h = n10 + nh
	n10q = n10 + nq
)

// MARK: STRUCT STRUCT STRUCT STRUCT STRUCT STRUCT STRUCT STRUCT STRUCT STRUCT STRUCT STRUCT
type xroom struct {
	roomrecs []xroomrec
	objs     []xobj
	extras   []xextra
	fx       []xfx
	items    []xitem

	floortile rl.Rectangle
}
type xtile struct {
	rec, img rl.Rectangle
	color    rl.Color
	fade     float32
	activ    bool
	dir      int
}
type xroomrec struct {
	rec, bordrec rl.Rectangle
	num          int

	wallrecs  []xtile
	floorrecs []xtile
}
type xplayer struct {
	rec, collisrec, boundaryrec, portrait rl.Rectangle

	vel, dirx, diry float32

	cnt rl.Vector2

	direc int

	atk, str int

	atktimer int32

	atkanim bool
}
type xobj struct {
	rec, img, collisrec rl.Rectangle

	hp, weight int

	canmove, solid, destructible, stop bool

	stopv2, hplosstxtv2 rl.Vector2

	name, hplossamouttxt string

	hplosspause int32

	fade  float32
	color rl.Color
}
type xextra struct {
	img, img2, rec, collisrec rl.Rectangle
	timer                     int32
	onoff                     bool

	fade  float32
	color rl.Color
}
type xfx struct {
	numtype  int
	rec, img rl.Rectangle

	fade  float32
	color rl.Color
}
type xitem struct {
	name1, name2, name3, itemname string

	numtype, sockets int

	img, rec rl.Rectangle

	color rl.Color
	fade  float32

	shoot, ammo, equipped bool

	atk, dur, pois, fire, ice, elec, vamp, thorns, regen, invis, diseas, shield float32
}

// MARK: DRAW DRAW DRAW DRAW DRAW DRAW DRAW DRAW DRAW DRAW DRAW DRAW DRAW DRAW DRAW DRAW DRAW
func cam() { //MARK: cam

	//DRAW FLOOR TILE IMGS
	for a := 0; a < len(floorback); a++ {
		rl.DrawTexturePro(imgs, floorback[a].img, floorback[a].rec, origin, 0, rl.Fade(floorback[a].color, floorback[a].fade))
	}

	//MOUSE CHECK
	for a := 0; a < len(rooms[currentroom].roomrecs); a++ {

		if dev2 {
			rl.DrawRectangleLinesEx(rooms[currentroom].roomrecs[a].rec, 1, rl.White)
			rl.DrawText(fmt.Sprint(rooms[currentroom].roomrecs[a].num), rooms[currentroom].roomrecs[a].rec.ToInt32().X+nh, rooms[currentroom].roomrecs[a].rec.ToInt32().Y+nh, txtm, rl.White)

			for b := 0; b < len(rooms[currentroom].roomrecs[a].wallrecs); b++ {
				rl.DrawRectangleLinesEx(rooms[currentroom].roomrecs[a].wallrecs[b].rec, 1, rl.Magenta)
			}

			//DRAW FLOOR OUTLINES DEV2
			for b := 0; b < len(rooms[currentroom].roomrecs[a].floorrecs); b++ {
				rl.DrawRectangleLinesEx(rooms[currentroom].roomrecs[a].floorrecs[b].rec, 1, rl.Fade(rl.DarkBrown, 0.2))
			}

		} else {

		}

		//MOUSIN CHECK
		mousin = false
		for b := 0; b < len(rooms[currentroom].roomrecs[a].floorrecs); b++ {
			checkrec := rl.NewRectangle(mousev2world.X-tilesize/2, mousev2world.Y-tilesize/2, tilesize, tilesize)
			if rl.CheckCollisionRecs(checkrec, rooms[currentroom].roomrecs[a].floorrecs[b].rec) {
				mousin = true
			}
		}

		if mousin {
			checkmouse()
		}

	}

	//DRAW EXTRAS
	if len(rooms[currentroom].extras) > 0 {

		for a := 0; a < len(rooms[currentroom].extras); a++ {
			rl.DrawTexturePro(imgs, rooms[currentroom].extras[a].img, rooms[currentroom].extras[a].rec, origin, 0, rl.Fade(rooms[currentroom].extras[a].color, rooms[currentroom].extras[a].fade))

			if dev2 {
				rl.DrawRectangleLinesEx(rooms[currentroom].extras[a].rec, 1, brightorange())
				rl.DrawRectangleLinesEx(rooms[currentroom].extras[a].collisrec, 1, brightred())

			}

			uproomextras(a)
		}

	}

	//DRAW OBJS
	if len(rooms[currentroom].objs) > 0 {
		for a := 0; a < len(rooms[currentroom].objs); a++ {

			rl.DrawTexturePro(imgs, rooms[currentroom].objs[a].img, rooms[currentroom].objs[a].rec, origin, 0, rl.Fade(rooms[currentroom].objs[a].color, rooms[currentroom].objs[a].fade))

			if rooms[currentroom].objs[a].hplosspause > 0 {
				rl.DrawText(rooms[currentroom].objs[a].hplossamouttxt, int32(rooms[currentroom].objs[a].hplosstxtv2.X), int32(rooms[currentroom].objs[a].hplosstxtv2.Y), txts, brightred())
				rooms[currentroom].objs[a].hplosstxtv2.Y--

			}

			if dev2 {
				rl.DrawRectangleLinesEx(rooms[currentroom].objs[a].rec, 1, brightorange())
			}

		}

	}

	//DRAW ITEMS
	if len(rooms[currentroom].items) > 0 {

		for a := 0; a < len(rooms[currentroom].items); a++ {

			rl.DrawTexturePro(imgs, rooms[currentroom].items[a].img, rooms[currentroom].items[a].rec, origin, 0, rl.Fade(rooms[currentroom].items[a].color, rooms[currentroom].items[a].fade))

			if rl.CheckCollisionPointRec(mousev2world, rooms[currentroom].items[a].rec) {
				drawiteminfo(a)
			}

			if dev2 {
				rl.DrawRectangleLinesEx(rooms[currentroom].items[a].rec, 1, brightorange())
			}

		}

	}

	//DRAW FX
	if len(rooms[currentroom].fx) > 0 {

		for a := 0; a < len(rooms[currentroom].fx); a++ {

			rl.DrawTexturePro(imgs, rooms[currentroom].fx[a].img, rooms[currentroom].fx[a].rec, origin, 0, rl.Fade(rooms[currentroom].fx[a].color, rooms[currentroom].fx[a].fade))

		}

	}

	//DRAW WALLS
	for a := 0; a < len(borderwallrecs); a++ {

		rl.DrawTexturePro(imgs, borderwallrecs[a].img, borderwallrecs[a].rec, origin, 0, rl.Fade(borderwallrecs[a].color, borderwallrecs[a].fade))
	}

	//SELPOINT
	if selpoint != blankv2 && !rl.CheckCollisionPointRec(player.cnt, selrec) {
		rl.DrawCircleV(selpoint, 2, rl.Fade(themecol1, fadeblink))
	}

	//PLAYER

	destrec := player.rec
	destrec.X -= (player.rec.Width / 2) + ue
	destrec.Y -= (player.rec.Width / 2) + uh + uq
	destrec.Width += player.rec.Width + uq
	destrec.Height += player.rec.Width + uq

	if player.atkanim {
		xdis := absdiff32(player.cnt.X, selpointr.X)
		ydis := absdiff32(player.cnt.Y, selpointr.Y)

		if xdis > ydis {

			if player.cnt.X > selpointr.X || player.direc == 4 {
				rl.DrawTexturePro(imgs, patkimgl, destrec, origin, 0, rl.White)
			} else if player.cnt.X < selpointr.X || player.direc == 1 {
				rl.DrawTexturePro(imgs, patkimgr, destrec, origin, 0, rl.White)
			}

		} else {

			if player.cnt.Y > selpointr.Y || player.direc == 1 {
				rl.DrawTexturePro(imgs, patkimgu, destrec, origin, 0, rl.White)
			} else if player.cnt.Y < selpointr.Y || player.direc == 3 {
				rl.DrawTexturePro(imgs, patkimgd, destrec, origin, 0, rl.White)

			}
		}
	} else {
		switch player.direc {
		case 0:
			rl.DrawTexturePro(imgs, pidleimg, destrec, origin, 0, rl.White)
		case 1:
			rl.DrawTexturePro(imgs, pwalkimgu, destrec, origin, 0, rl.White)
		case 2:
			rl.DrawTexturePro(imgs, pwalkimgr, destrec, origin, 0, rl.White)
		case 3:
			rl.DrawTexturePro(imgs, pwalkimgd, destrec, origin, 0, rl.White)
		case 4:

			rl.DrawTexturePro(imgs, pwalkimgl, destrec, origin, 0, rl.White)

		default:
			rl.DrawRectangleLinesEx(player.rec, 1, brightorange())
		}
	}

	if dev2 {
		rl.DrawCircleV(player.cnt, 3, rl.Magenta)
		rl.DrawRectangleLinesEx(player.rec, 1, brightorange())      //PLAYER REC
		rl.DrawRectangleLinesEx(player.boundaryrec, 1, brightred()) //PLAYER BOUNDSREC
	}

	//DRAW LAYER 2 DRAW LAYER 2 DRAW LAYER 2 DRAW LAYER 2 DRAW LAYER 2 DRAW LAYER 2 DRAW LAYER 2

	//DRAW ITEM INFO COLLECT ITEM
	if len(rooms[currentroom].items) > 0 {

		for a := 0; a < len(rooms[currentroom].items); a++ {

			if rl.CheckCollisionPointRec(mousev2world, rooms[currentroom].items[a].rec) {
				drawiteminfo(a)

				if rl.CheckCollisionRecs(player.boundaryrec, rooms[currentroom].items[a].rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						addtoinv(a)
					}
				}
			}

		}

	}

}
func nocambackg() { //MARK: nocambackg

}
func nocam() { //MARK: nocam

	//INFOBAR REC
	rl.DrawRectangleRec(infobarrec, rl.Fade(rl.Black, 0.8))
	timehere(scrwf32, infobarrec.Y+ue)

	//MSG
	if msgtimer > 0 {
		rl.DrawText(msgs[len(msgs)-1], ne-1, infobarrec.ToInt32().Y+ne+1, txts, rl.Black)
		rl.DrawText(msgs[len(msgs)-1], ne, infobarrec.ToInt32().Y+ne, txts, themecol1)
	}

	//INFOBAR REC2
	rl.DrawRectangleRec(infobarrec2, rl.Fade(rl.Black, 0.7))

	rl.DrawText("playerX "+fmt.Sprintf("%.0f", player.cnt.X)+" playerY "+fmt.Sprintf("%.0f", player.cnt.Y), ne, infobarrec2.ToInt32().Y+ne, txts, rl.White)

	txtlen := rl.MeasureText("playerX  playerY 0000 0000 ", txts)

	rl.DrawText("mouseX "+fmt.Sprintf("%.0f", mousev2world.X)+" mouseY "+fmt.Sprintf("%.0f", mousev2world.Y), ne+txtlen, infobarrec2.ToInt32().Y+ne, txts, rl.White)

	//INV
	if rl.CheckCollisionPointRec(mousev2, invtabrec) || invon {
		rl.DrawRectangleRec(invtabrec, rl.Fade(themecol1, 0.7))
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			pause = true
			invon = true
		}

	} else {
		shadowrec := invtabrec
		shadowrec.X = shadowrec.X - 3
		shadowrec.Y = shadowrec.Y + 3
		shadowrec.Width = shadowrec.Width + ue

		rl.DrawRectangleRec(shadowrec, rl.Fade(rl.DarkGray, 0.9))
		rl.DrawRectangleRec(invtabrec, rl.Black)
		rl.DrawRectangleLinesEx(invtabrec, 1, rl.White)
	}
	iconrec := rl.NewRectangle(invtabrec.X, invtabrec.Y, u1, u1)

	if invon {
		iconrec.Width -= ue
		iconrec.Height -= ue
		iconrec.X += ue / 2
		iconrec.Y += ue / 2
		if rl.CheckCollisionPointRec(mousev2, iconrec) {
			rl.DrawTexturePro(imgs, closeimg, iconrec, origin, 0, brightred())
		} else {
			rl.DrawTexturePro(imgs, closeimg, iconrec, origin, 0, rl.White)
		}
		//CLOSE INV
		if rl.CheckCollisionPointRec(mousev2, iconrec) {
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) && invrec.X < scrwf32-u2 {
				invclose = true
				invopen = false
			}
		}

	} else {
		if rl.CheckCollisionPointRec(mousev2, iconrec) {
			rl.DrawTexturePro(imgs, invimg, iconrec, origin, 0, rl.Green)
		} else {
			rl.DrawTexturePro(imgs, invimg, iconrec, origin, 0, rl.White)
		}

	}

	if invon {
		drawinv()
	}

}

func devui() { //MARK: devui

	siderec := rl.NewRectangle(0, 0, 300, scrhf32)

	rl.DrawRectangleRec(siderec, rl.Fade(rl.Green, 0.5))

	x := int32(siderec.X + 10)
	y := int32(10)

	txt := "mousev2.X"
	rl.DrawText(txt, x, y, txts, rl.White)
	x += rl.MeasureText(txt, txts) + txts
	txt = fmt.Sprint(mousev2.X)
	rl.DrawText(txt, x, y, txts, rl.White)
	x = int32(siderec.X + 10)
	y += txts

	txt = "camera zoom"
	rl.DrawText(txt, x, y, txts, rl.White)
	x += rl.MeasureText(txt, txts) + txts
	txt = fmt.Sprint(camera.Zoom)
	rl.DrawText(txt, x, y, txts, rl.White)
	x = int32(siderec.X + 10)
	y += txts

	txt = "maxuw"
	rl.DrawText(txt, x, y, txts, rl.White)
	x += rl.MeasureText(txt, txts) + txts
	txt = fmt.Sprint(maxuw)
	rl.DrawText(txt, x, y, txts, rl.White)
	x = int32(siderec.X + 10)
	y += txts

	txt = "maxuh"
	rl.DrawText(txt, x, y, txts, rl.White)
	x += rl.MeasureText(txt, txts) + txts
	txt = fmt.Sprint(maxuh)
	rl.DrawText(txt, x, y, txts, rl.White)
	x = int32(siderec.X + 10)
	y += txts

}
func drawinv() { //MARK: drawinv

	rl.DrawRectangleRec(invrec, rl.Fade(themecol1, 0.7))

	playerstats := true
	if invopen {

		//EQUIPPED ITEM ICONS
		x := invrec.X + uh + ue
		y := invtabrec.Y + invtabrec.Height + u1

		rec := rl.NewRectangle(x, y, u1+ue, u1+ue)

		if rl.CheckCollisionPointRec(mousev2, rec) {
			rl.DrawRectangleRec(rec, rl.Fade(rl.Black, fadeblink))

		} else {
			rl.DrawRectangleRec(rec, rl.Fade(rl.Black, 0.8))
		}

		rec2 := rec
		rec2.X += 4
		rec2.Y += 4
		rec2.Width -= 8
		rec2.Height -= 8

		if invequip[0].name1 =="" {
		rl.DrawTexturePro(imgs, invweap1img, rec2, origin, 0, rl.Fade(rl.White, 0.5))
		}

		rec.X += u1q

		if rl.CheckCollisionPointRec(mousev2, rec) {
			rl.DrawRectangleRec(rec, rl.Fade(rl.Black, fadeblink))

		} else {
			rl.DrawRectangleRec(rec, rl.Fade(rl.Black, 0.8))
		}

		rec2 = rec
		rec2.X += 4
		rec2.Y += 4
		rec2.Width -= 8
		rec2.Height -= 8

		if invequip[1].name1 =="" {
		rl.DrawTexturePro(imgs, invweap2img, rec2, origin, 0, rl.Fade(rl.White, 0.5))
		}

		rec.X += u1q

		if rl.CheckCollisionPointRec(mousev2, rec) {
			rl.DrawRectangleRec(rec, rl.Fade(rl.Black, fadeblink))

		} else {
			rl.DrawRectangleRec(rec, rl.Fade(rl.Black, 0.8))
		}

		rec2 = rec
		rec2.X += 4
		rec2.Y += 4
		rec2.Width -= 8
		rec2.Height -= 8
		if invequip[2].name1 =="" {
		rl.DrawTexturePro(imgs, invweap3img, rec2, origin, 0, rl.Fade(rl.White, 0.5))
		}
		rec.X += u1q

		if rl.CheckCollisionPointRec(mousev2, rec) {
			rl.DrawRectangleRec(rec, rl.Fade(rl.Black, fadeblink))

		} else {
			rl.DrawRectangleRec(rec, rl.Fade(rl.Black, 0.8))
		}

		rec2 = rec
		rec2.X += 4
		rec2.Y += 4
		rec2.Width -= 8
		rec2.Height -= 8
		if invequip[3].name1 =="" {
		rl.DrawTexturePro(imgs, invweap4img, rec2, origin, 0, rl.Fade(rl.White, 0.5))
		}
		rec.X += u1q

		if rl.CheckCollisionPointRec(mousev2, rec) {
			rl.DrawRectangleRec(rec, rl.Fade(rl.Black, fadeblink))

		} else {
			rl.DrawRectangleRec(rec, rl.Fade(rl.Black, 0.8))
		}

		rec2 = rec
		rec2.X += 4
		rec2.Y += 4
		rec2.Width -= 8
		rec2.Height -= 8
		if invequip[4].name1 =="" {
		rl.DrawTexturePro(imgs, invammoimg, rec2, origin, 0, rl.Fade(rl.White, 0.5))
		}
		rec.X += u1q

		if rl.CheckCollisionPointRec(mousev2, rec) {
			rl.DrawRectangleRec(rec, rl.Fade(rl.Black, fadeblink))

		} else {
			rl.DrawRectangleRec(rec, rl.Fade(rl.Black, 0.8))
		}

		rec2 = rec
		rec2.X += 4
		rec2.Y += 4
		rec2.Width -= 8
		rec2.Height -= 8
		if invequip[5].name1 =="" {
		rl.DrawTexturePro(imgs, invammoimg, rec2, origin, 0, rl.Fade(rl.White, 0.5))
		}
		rec.X += u1q

		if rl.CheckCollisionPointRec(mousev2, rec) {
			rl.DrawRectangleRec(rec, rl.Fade(rl.Black, fadeblink))

		} else {
			rl.DrawRectangleRec(rec, rl.Fade(rl.Black, 0.8))
		}

		rec2 = rec
		rec2.X += 4
		rec2.Y += 4
		rec2.Width -= 8
		rec2.Height -= 8
		if invequip[6].name1 =="" {
		rl.DrawTexturePro(imgs, invhelmetimg, rec2, origin, 0, rl.Fade(rl.White, 0.5))
		}
		rec.X += u1q

		if rl.CheckCollisionPointRec(mousev2, rec) {
			rl.DrawRectangleRec(rec, rl.Fade(rl.Black, fadeblink))

		} else {
			rl.DrawRectangleRec(rec, rl.Fade(rl.Black, 0.8))
		}

		rec2 = rec
		rec2.X += 4
		rec2.Y += 4
		rec2.Width -= 8
		rec2.Height -= 8
		if invequip[7].name1 =="" {
		rl.DrawTexturePro(imgs, invarmorimg, rec2, origin, 0, rl.Fade(rl.White, 0.5))
		}
		rec.X += u1q

		if rl.CheckCollisionPointRec(mousev2, rec) {
			rl.DrawRectangleRec(rec, rl.Fade(rl.Black, fadeblink))

		} else {
			rl.DrawRectangleRec(rec, rl.Fade(rl.Black, 0.8))
		}

		rec2 = rec
		rec2.X += 4
		rec2.Y += 4
		rec2.Width -= 8
		rec2.Height -= 8

		if invequip[8].name1 =="" {
		rl.DrawTexturePro(imgs, invglovesimg, rec2, origin, 0, rl.Fade(rl.White, 0.5))
		}
		rec.X += u1q

		if rl.CheckCollisionPointRec(mousev2, rec) {
			rl.DrawRectangleRec(rec, rl.Fade(rl.Black, fadeblink))

		} else {
			rl.DrawRectangleRec(rec, rl.Fade(rl.Black, 0.8))
		}

		rec2 = rec
		rec2.X += 4
		rec2.Y += 4
		rec2.Width -= 8
		rec2.Height -= 8
		if invequip[9].name1 =="" {
		rl.DrawTexturePro(imgs, invbootsimg, rec2, origin, 0, rl.Fade(rl.White, 0.5))
		}
		rec.X += u1q

		if rl.CheckCollisionPointRec(mousev2, rec) {
			rl.DrawRectangleRec(rec, rl.Fade(rl.Black, fadeblink))

		} else {
			rl.DrawRectangleRec(rec, rl.Fade(rl.Black, 0.8))
		}

		rec2 = rec
		rec2.X += 4
		rec2.Y += 4
		rec2.Width -= 8
		rec2.Height -= 8
		if invequip[10].name1 =="" {
		rl.DrawTexturePro(imgs, invjewelryimg, rec2, origin, 0, rl.Fade(rl.White, 0.5))
		}
		rec.X += u1q

		if rl.CheckCollisionPointRec(mousev2, rec) {
			rl.DrawRectangleRec(rec, rl.Fade(rl.Black, fadeblink))

		} else {
			rl.DrawRectangleRec(rec, rl.Fade(rl.Black, 0.8))
		}

		rec2 = rec
		rec2.X += 4
		rec2.Y += 4
		rec2.Width -= 8
		rec2.Height -= 8
		if invequip[11].name1 =="" {
		rl.DrawTexturePro(imgs, invringimg, rec2, origin, 0, rl.Fade(rl.White, 0.5))
		}
		rec.X += u1q

		if rl.CheckCollisionPointRec(mousev2, rec) {
			rl.DrawRectangleRec(rec, rl.Fade(rl.Black, fadeblink))

		} else {
			rl.DrawRectangleRec(rec, rl.Fade(rl.Black, 0.8))
		}

		rec2 = rec
		rec2.X += 4
		rec2.Y += 4
		rec2.Width -= 8
		rec2.Height -= 8
		if invequip[12].name1 =="" {
		rl.DrawTexturePro(imgs, invringimg, rec2, origin, 0, rl.Fade(rl.White, 0.5))
		}
		rec.X += u1q

		if rl.CheckCollisionPointRec(mousev2, rec) {
			rl.DrawRectangleRec(rec, rl.Fade(rl.Black, fadeblink))

		} else {
			rl.DrawRectangleRec(rec, rl.Fade(rl.Black, 0.8))
		}

		rec2 = rec
		rec2.X += 4
		rec2.Y += 4
		rec2.Width -= 8
		rec2.Height -= 8
		if invequip[13].name1 =="" {
		rl.DrawTexturePro(imgs, invringimg, rec2, origin, 0, rl.Fade(rl.White, 0.5))
		}
		rec.X += u1q

		if rl.CheckCollisionPointRec(mousev2, rec) {
			rl.DrawRectangleRec(rec, rl.Fade(rl.Black, fadeblink))

		} else {
			rl.DrawRectangleRec(rec, rl.Fade(rl.Black, 0.8))
		}

		rec2 = rec
		rec2.X += 4
		rec2.Y += 4
		rec2.Width -= 8
		rec2.Height -= 8
		if invequip[14].name1 =="" {
		rl.DrawTexturePro(imgs, invringimg, rec2, origin, 0, rl.Fade(rl.White, 0.5))
		}
		rec.X += u1q

		//INV ITEM RECS
		y = invtabrec.Y + invtabrec.Height + u1 + u1q
		x = invrec.X + uh + ue

		count := 0

		for a := 0; a < len(inv); a++ {

			//BACK REC
			rec := rl.NewRectangle(x, y, u1+ue, u1+ue)
			if rl.CheckCollisionPointRec(mousev2, rec) {
				rl.DrawRectangleRec(rec, rl.Fade(rl.Black, fadeblink))
				if inv[a].name1 != "" {
					playerstats = false
					invplayerstatstimer = fps / 3
				}
			} else {
				rl.DrawRectangleRec(rec, rl.Fade(rl.Black, 0.8))
			}
			rl.DrawRectangleLinesEx(rec, 2, rl.Black)

			//INV IMG

			if inv[a].name1 != "" {
				rl.DrawTexturePro(imgs, inv[a].img, rec, origin, 0, rl.White)
				if inv[a].equipped {
					rl.DrawRectangleRec(rec,rl.Fade(rl.Blue,0.5))
				}
				if rl.CheckCollisionPointRec(mousev2, rec) {

					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						selectedinv = inv[a]
						selectedinvnum = a
					}

				}
			}

			x += u1q

			count++
			if count == 15 {
				count = 0
				x = invrec.X + uh + ue
				y += u1q
			}
		}

		//EQUIPPED INV ITEMS
		x = invrec.X + uh + ue
		y = invtabrec.Y + invtabrec.Height + u1

		rec = rl.NewRectangle(x, y, u1+ue, u1+ue)
		for a := 0; a < len(invequip); a++ {

			if invequip[a].name1 != "" {
				rl.DrawTexturePro(imgs, invequip[a].img, rec, origin, 0, rl.White)
			}

			if rl.CheckCollisionPointRec(mousev2, rec) {
				if invequip[a].name1 != "" {

				} else {

					if selectedinv.name1 != "" {

						switch selectedinv.numtype {
						case 1:
							if a == 0 || a == 1 || a == 2 || a == 3 {
								rl.DrawRectangleRec(rec, rl.Fade(rl.Green, 0.2))
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									invequip[a] = selectedinv
									inv[selectedinvnum].equipped = true
									selectedinvnum = blankint
									selectedinv = xitem{}
								}
							} else {
								rl.DrawRectangleRec(rec, rl.Fade(brightred(), 0.2))
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									addmsg("cannot be equipped in this slot", false)
								}
							}

						}

					} else {

						txtx := int32(rec.X)
						txty := int32(rec.Y + rec.Height + ue)

						switch a {
						case 0, 1:
							txtlen := rl.MeasureText("or shield", txts)
							wid := float32(txtlen + nq)
							heig := float32(txts * 3)
							backrec := rl.NewRectangle(rec.X-ue, rec.Y+rec.Height+ue, wid, heig)
							rl.DrawRectangleRec(backrec, rl.Fade(rl.Black, 0.8))

							rl.DrawText("weapon", txtx, txty, txts, rl.White)
							rl.DrawText("or shield", txtx, txty+txts, txts, rl.White)
							rl.DrawText("set 1", txtx, txty+(txts*2), txts, rl.White)

						case 2, 3:
							txtlen := rl.MeasureText("or shield", txts)
							wid := float32(txtlen + nq)
							heig := float32(txts * 3)
							backrec := rl.NewRectangle(rec.X-ue, rec.Y+rec.Height+ue, wid, heig)
							rl.DrawRectangleRec(backrec, rl.Fade(rl.Black, 0.8))

							rl.DrawText("weapon", txtx, txty, txts, rl.White)
							rl.DrawText("or shield", txtx, txty+txts, txts, rl.White)
							rl.DrawText("set 2", txtx, txty+(txts*2), txts, rl.White)
						case 4, 5:
							txtlen := rl.MeasureText("ammo", txts)
							wid := float32(txtlen + nq)
							heig := float32(txts)
							backrec := rl.NewRectangle(rec.X-ue, rec.Y+rec.Height+ue, wid, heig)
							rl.DrawRectangleRec(backrec, rl.Fade(rl.Black, 0.8))

							rl.DrawText("ammo", txtx, txty, txts, rl.White)
						case 6:
							txtlen := rl.MeasureText("helmet", txts)
							wid := float32(txtlen + nq)
							heig := float32(txts)
							backrec := rl.NewRectangle(rec.X-ue, rec.Y+rec.Height+ue, wid, heig)
							rl.DrawRectangleRec(backrec, rl.Fade(rl.Black, 0.8))

							rl.DrawText("helmet", txtx, txty, txts, rl.White)
						case 7:
							txtlen := rl.MeasureText("armor", txts)
							wid := float32(txtlen + nq)
							heig := float32(txts)
							backrec := rl.NewRectangle(rec.X-ue, rec.Y+rec.Height+ue, wid, heig)
							rl.DrawRectangleRec(backrec, rl.Fade(rl.Black, 0.8))

							rl.DrawText("armor", txtx, txty, txts, rl.White)
						case 8:
							txtlen := rl.MeasureText("gloves", txts)
							wid := float32(txtlen + nq)
							heig := float32(txts)
							backrec := rl.NewRectangle(rec.X-ue, rec.Y+rec.Height+ue, wid, heig)
							rl.DrawRectangleRec(backrec, rl.Fade(rl.Black, 0.8))

							rl.DrawText("gloves", txtx, txty, txts, rl.White)
						case 9:
							txtlen := rl.MeasureText("boots", txts)
							wid := float32(txtlen + nq)
							heig := float32(txts)
							backrec := rl.NewRectangle(rec.X-ue, rec.Y+rec.Height+ue, wid, heig)
							rl.DrawRectangleRec(backrec, rl.Fade(rl.Black, 0.8))

							rl.DrawText("boots", txtx, txty, txts, rl.White)
						case 10:
							txtlen := rl.MeasureText("jewelry", txts)
							wid := float32(txtlen + nq)
							heig := float32(txts)
							backrec := rl.NewRectangle(rec.X-ue, rec.Y+rec.Height+ue, wid, heig)
							rl.DrawRectangleRec(backrec, rl.Fade(rl.Black, 0.8))

							rl.DrawText("jewelry", txtx, txty, txts, rl.White)
						case 11, 12, 13, 14:
							txtlen := rl.MeasureText("ring", txts)
							wid := float32(txtlen + nq)
							heig := float32(txts)
							backrec := rl.NewRectangle(rec.X-ue, rec.Y+rec.Height+ue, wid, heig)
							rl.DrawRectangleRec(backrec, rl.Fade(rl.Black, 0.8))

							rl.DrawText("ring", txtx, txty, txts, rl.White)
						}
					}
				}

			}

			rec.X += u1q
		}

		//PLAYER STATS
		if playerstats {
			invplayerstatstimer--
			if invplayerstatstimer <= 0 {
				invplayerstatstimer = 0
				rec := rl.NewRectangle(invrec.X+uh+ue, invrec.Y+uh+ue, u2h, u2h)
				backrec := rec
				backrec.X -= 2
				backrec.Y -= 2
				backrec.Width += 4
				backrec.Height += 4
				rl.DrawRectangleRec(backrec, rl.Black)
				rl.DrawTexturePro(imgs, player.portrait, rec, origin, 0, rl.White)

			}
		}
	}

}
func drawiteminfo(itemnum int) { //MARK: drawiteminfo

	//BACKG
	rec := rl.NewRectangle(0, 0, u3, u5)
	rec.X = (rooms[currentroom].items[itemnum].rec.X + (rooms[currentroom].items[itemnum].rec.Width / 2)) - rec.Width/2
	rec.Y = rooms[currentroom].items[itemnum].rec.Y - (rec.Height + uh)
	rl.DrawRectangleRec(rec, rl.Fade(themecol2, 0.7))

	//IMG
	backrec := rl.NewRectangle(rec.X, rec.Y, rec.Width, rec.Width)
	rl.DrawRectangleRec(backrec, rl.Fade(themecol2, 0.7))
	imgrec := rl.NewRectangle(rec.X, rec.Y, rec.Width-u1, rec.Width-u1)
	imgrec.X += uh
	imgrec.Y += uh
	rl.DrawTexturePro(imgs, rooms[currentroom].items[itemnum].img, imgrec, origin, 0, rl.White)

	//TXT
	txtcnt := int32(rec.X + rec.Width/2)
	txty := int32(backrec.Y + backrec.Height)

	txtlen := rl.MeasureText(rooms[currentroom].items[itemnum].name2, txts)
	txtx := txtcnt - txtlen/2
	rl.DrawText(rooms[currentroom].items[itemnum].name2, txtx, txty, txts, rl.White)
	txty += txts

	txtlen = rl.MeasureText(rooms[currentroom].items[itemnum].name1+" of ", txts)
	txtx = txtcnt - txtlen/2
	rl.DrawText(rooms[currentroom].items[itemnum].name1+" of ", txtx, txty, txts, rl.White)
	txty += txts

	txtlen = rl.MeasureText(rooms[currentroom].items[itemnum].name3, txts)
	txtx = txtcnt - txtlen/2
	rl.DrawText(rooms[currentroom].items[itemnum].name3, txtx, txty, txts, rl.White)

	//SOCKETS
	if rooms[currentroom].items[itemnum].sockets > 0 {
		txty += txts * 3

		socketrec := rl.NewRectangle(rec.X, rec.Y+rec.Height, rec.Width, u1)

		rl.DrawRectangleRec(socketrec, rl.Fade(themecol2, 0.7))

		if rooms[currentroom].items[itemnum].sockets == 1 {
			rl.DrawCircleGradient(txtcnt, txty, uq+ue, themecol1, themecol2)
			rl.DrawCircleLines(txtcnt, txty, uq+ue, rl.Black)

		} else if rooms[currentroom].items[itemnum].sockets == 2 {

			rl.DrawCircleGradient(txtcnt-nh, txty, uq+ue, themecol1, themecol2)
			rl.DrawCircleLines(txtcnt-nh, txty, uq+ue, rl.Black)

			rl.DrawCircleGradient(txtcnt+nh, txty, uq+ue, themecol1, themecol2)
			rl.DrawCircleLines(txtcnt+nh, txty, uq+ue, rl.Black)

		} else if rooms[currentroom].items[itemnum].sockets == 3 {

			rl.DrawCircleGradient(txtcnt, txty, uq+ue, themecol1, themecol2)
			rl.DrawCircleLines(txtcnt, txty, uq+ue, rl.Black)

			rl.DrawCircleGradient(txtcnt-(n1-4), txty, uq+ue, themecol1, themecol2)
			rl.DrawCircleLines(txtcnt-(n1-4), txty, uq+ue, rl.Black)

			rl.DrawCircleGradient(txtcnt+(n1-4), txty, uq+ue, themecol1, themecol2)
			rl.DrawCircleLines(txtcnt+(n1-4), txty, uq+ue, rl.Black)

		}
	}

	//ADD MSG
	msgtxt := ""
	if rooms[currentroom].items[itemnum].atk > 0 {
		msgtxt = msgtxt + " " + "atk: " + fmt.Sprintf("%.2f", rooms[currentroom].items[itemnum].atk)

	}

	if rooms[currentroom].items[itemnum].dur > 0 {
		msgtxt = msgtxt + " " + "dur: " + fmt.Sprintf("%.2f", rooms[currentroom].items[itemnum].dur)
	}

	if rooms[currentroom].items[itemnum].elec > 0 {
		msgtxt = msgtxt + " " + "elec: " + fmt.Sprintf("%.2f", rooms[currentroom].items[itemnum].elec)
	}

	if rooms[currentroom].items[itemnum].fire > 0 {
		msgtxt = msgtxt + " " + "fire: " + fmt.Sprintf("%.2f", rooms[currentroom].items[itemnum].fire)
	}

	if rooms[currentroom].items[itemnum].ice > 0 {
		msgtxt = msgtxt + " " + "ice: " + fmt.Sprintf("%.2f", rooms[currentroom].items[itemnum].ice)
	}

	if rooms[currentroom].items[itemnum].pois > 0 {
		msgtxt = msgtxt + " " + "pois: " + fmt.Sprintf("%.2f", rooms[currentroom].items[itemnum].pois)
	}

	if rooms[currentroom].items[itemnum].diseas > 0 {
		msgtxt = msgtxt + " " + "diseas: " + fmt.Sprintf("%.2f", rooms[currentroom].items[itemnum].diseas)
	}

	if rooms[currentroom].items[itemnum].thorns > 0 {
		msgtxt = msgtxt + " " + "thorns: " + fmt.Sprintf("%.2f", rooms[currentroom].items[itemnum].thorns)
	}

	if rooms[currentroom].items[itemnum].vamp > 0 {
		msgtxt = msgtxt + " " + "vamp: " + fmt.Sprintf("%.2f", rooms[currentroom].items[itemnum].vamp)
	}

	if rooms[currentroom].items[itemnum].regen > 0 {
		msgtxt = msgtxt + " " + "regen: " + fmt.Sprintf("%.2f", rooms[currentroom].items[itemnum].regen)
	}

	if rooms[currentroom].items[itemnum].shield > 0 {
		msgtxt = msgtxt + " " + "shield: " + fmt.Sprintf("%.2f", rooms[currentroom].items[itemnum].shield)
	}

	if rooms[currentroom].items[itemnum].invis > 0 {
		msgtxt = msgtxt + " " + "invis: " + fmt.Sprintf("%.2f", rooms[currentroom].items[itemnum].invis)
	}

	addmsg(msgtxt, true)

}

// MARK: MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE
func makeroom() { //MARK: makeroom

	//MAKE ROOM MAP
	borderwallrecs = nil
	floorback = nil
	msgs = nil
	msgtimer = 0

	selpoint = blankv2
	selrec = blankrec

	x := scrcnt.X
	y := scrcnt.Y
	zroom := xroom{}
	zroom.floortile = floorimgs[rInt(0, len(floorimgs))]

	//CENTRE ROOM
	multiw := rInt(4, maxuw/4)
	multih := rInt(4, maxuh/4)
	wid := tilesize * float32(multiw)
	widorig := wid
	heig := tilesize * float32(multih)
	heigorig := heig
	rec := rl.NewRectangle(x-wid/2, y-heig/2, wid, heig)
	recorig := rec

	zroomrec := xroomrec{}
	zroomrec.rec = rec

	zroom.roomrecs = append(zroom.roomrecs, zroomrec)

	num := rInt(10, 51)
	exitcount2 := 0

	for {
		choose := rInt(0, 4)

		chooserec := rInt(0, len(zroom.roomrecs))
		recorig = zroom.roomrecs[chooserec].rec
		widorig = recorig.Width
		heigorig = recorig.Height

		exitswitch := false
		exitcount := 0

		for !exitswitch && exitcount < 20 {

			divwmin := rInt(3, 6)
			divhmin := rInt(3, 6)

			divwmin2 := rInt(3, 6)
			divhmin2 := rInt(3, 6)

			rec = recorig
			multiw = rInt(divwmin, maxuw/divwmin2)
			multih = rInt(divhmin, maxuh/divhmin2)
			wid = tilesize * float32(multiw)
			heig = tilesize * float32(multih)

			rec.Height = heig
			rec.Width = wid

			switch choose {

			case 2: //BOTTOM

				rec.Y += heigorig

				minx := wid - (tilesize * 2)
				maxx := widorig - (tilesize * 2)

				rec.X += rFloat32(-minx, maxx)

			case 3: //LEFT

				rec.X -= wid

				miny := heig - (tilesize * 2)
				maxy := heigorig - (tilesize * 2)

				rec.Y += rFloat32(-miny, maxy)

			case 1: //RIGHT

				rec.X += widorig

				miny := heig - (tilesize * 2)
				maxy := heigorig - (tilesize * 2)

				rec.Y += rFloat32(-miny, maxy)

			case 0: //TOP

				rec.Y -= heig

				minx := wid - (tilesize * 2)
				maxx := widorig - (tilesize * 2)

				rec.X += rFloat32(-minx, maxx)

			} //SWITCH

			if checkaddroomrec(rec, zroom) {
				num--
				zroomrec.rec = rec
				zroomrec.bordrec = makeroomrecborderec(zroomrec.rec)
				zroomrec.num = len(zroom.roomrecs)

				zroom.roomrecs = append(zroom.roomrecs, zroomrec)
				exitswitch = true
			} else {
				exitcount++
			}

		} //FOR SWITCH LOOP

		exitcount2++
		if num == 0 || exitcount2 == 200 {
			break
		}
	}

	for a := 0; a < len(zroom.roomrecs); a++ {
		zroom.roomrecs[a].wallrecs = makeroomwallrecs(zroom.roomrecs[a].rec, zroom)
		zroom.roomrecs[a].floorrecs = makeroomfloortilerecs(zroom.roomrecs[a].rec, zroom)
	}

	rooms = append(rooms, zroom)

	//MAKE FLOOR
	makefloorback(zroom)

	//ADD ROOM EXTRAS
	makeextras()
	makeweapons()
	makecrates()
	makemoveblocks()

}
func makeextras() { //MARK: makeextras

	for a := 0; a < len(rooms); a++ {

		for b := 0; b < len(rooms[a].roomrecs); b++ {
			if rolldice() == 6 {

				zextra := makegenericextra(tilesize)
				zextra.img = flamespoutimg
				v2 := findpointinroom(a, b, zextra.rec)
				zextra.rec.X = v2.X
				zextra.rec.Y = v2.Y
				zextra.collisrec = makecollisrec(zextra.rec)
				rooms[a].extras = append(rooms[a].extras, zextra)

			}
		}

	}

}
func makecollisrec(rec rl.Rectangle) rl.Rectangle {

	rec.X -= rec.Width
	rec.Y -= rec.Height

	rec.Width += rec.Width * 2
	rec.Height += rec.Height * 2

	return rec

}

func makefloorback(room xroom) { //MARK: makefloorback

	x := float32(0)
	y := float32(0)

	ztile := xtile{}
	ztile.img = room.floortile

	for {
		rec := rl.NewRectangle(x, y, tilesize+(tilesize/3), tilesize+(tilesize/3))

		for a := 0; a < len(room.roomrecs); a++ {
			if rl.CheckCollisionRecs(rec, room.roomrecs[a].rec) {

				rec.Width = tilesize
				rec.Height = tilesize
				ztile.rec = rec
				ztile.color = rl.DarkGray
				ztile.fade = rFloat32(0.1, 0.2)
				floorback = append(floorback, ztile)
			}

		}

		x += tilesize
		if x >= scrwf32 {
			x = 0
			y += tilesize
		}
		if y >= scrhf32 {
			break
		}
	}

}
func makeroomrecborderec(rec rl.Rectangle) rl.Rectangle { //MARK: makeroomrecborderec

	rec.X -= tilesize
	rec.Y -= tilesize
	rec.Width += tilesize * 2
	rec.Height += tilesize * 2

	return rec

}
func makeroomfloortilerecs(rec rl.Rectangle, room xroom) []xtile { //MARK: makeroomfloortiles

	var recs []xtile

	x := rec.X
	y := rec.Y

	ztile := xtile{}
	ztile.img = room.floortile

	for {

		newrec := rl.NewRectangle(x, y, floortilesize, floortilesize)
		ztile.rec = newrec

		recs = append(recs, ztile)

		if ztile.activ {
			ztile.activ = false
		}

		x += floortilesize
		if x >= rec.X+rec.Width {
			x = rec.X
			y += floortilesize
		}

		if y >= rec.Y+rec.Height {
			break
		}

	}

	return recs
}

func makeroomwallrecs(rec rl.Rectangle, room xroom) []xtile { //MARK: makeroomwallrecs

	var recs []xtile

	x := rec.X - tilesize
	y := rec.Y - tilesize
	ztile := xtile{}
	ztile.color = rl.White
	ztile.fade = 1.0
	ztile.img = wall1img

	for { //TOP TILES

		checkrec := rl.NewRectangle(x, y, tilesize, tilesize)
		canadd := true
		for a := 0; a < len(room.roomrecs); a++ {

			if rl.CheckCollisionRecs(checkrec, room.roomrecs[a].rec) {
				canadd = false
			}
		}

		if canadd {
			ztile.rec = checkrec
			ztile.dir = 1
			recs = append(recs, ztile)
			borderwallrecs = append(borderwallrecs, ztile)
		}

		x += tilesize
		if x > rec.X+rec.Width {
			break
		}

	}

	x = rec.X - tilesize
	y = rec.Y + rec.Height

	for { //BOTTOM TILES

		checkrec := rl.NewRectangle(x, y, tilesize, tilesize)
		canadd := true
		for a := 0; a < len(room.roomrecs); a++ {

			if rl.CheckCollisionRecs(checkrec, room.roomrecs[a].rec) {
				canadd = false
			}
		}

		if canadd {
			ztile.rec = checkrec
			ztile.dir = 3
			recs = append(recs, ztile)
			borderwallrecs = append(borderwallrecs, ztile)
		}

		x += tilesize
		if x > rec.X+rec.Width {
			break
		}

	}

	x = rec.X - tilesize
	y = rec.Y

	for { //LEFT TILES

		checkrec := rl.NewRectangle(x, y, tilesize, tilesize)
		canadd := true
		for a := 0; a < len(room.roomrecs); a++ {

			if rl.CheckCollisionRecs(checkrec, room.roomrecs[a].rec) {
				canadd = false
			}
		}

		if canadd {
			ztile.rec = checkrec
			ztile.dir = 4
			recs = append(recs, ztile)
			borderwallrecs = append(borderwallrecs, ztile)
		}

		y += tilesize
		if y == rec.Y+rec.Height {
			break
		}

	}

	x = rec.X + rec.Width
	y = rec.Y

	for { //RIGHT TILES

		checkrec := rl.NewRectangle(x, y, tilesize, tilesize)
		canadd := true
		for a := 0; a < len(room.roomrecs); a++ {

			if rl.CheckCollisionRecs(checkrec, room.roomrecs[a].rec) {
				canadd = false
			}
		}

		if canadd {
			ztile.rec = checkrec
			ztile.dir = 2
			recs = append(recs, ztile)
			borderwallrecs = append(borderwallrecs, ztile)
		}

		y += tilesize
		if y == rec.Y+rec.Height {
			break
		}

	}

	return recs

}

func makeplayer() { //MARK: makeplayer

	player.portrait = portraitimgs[rInt(0, len(portraitimgs))]

	player.atk = 1
	player.str = 2

	player.vel = 8

	x := rooms[currentroom].roomrecs[0].rec.X
	y := rooms[currentroom].roomrecs[0].rec.Y
	x += rooms[currentroom].roomrecs[0].rec.Width / 2
	y += rooms[currentroom].roomrecs[0].rec.Height / 2

	player.cnt = rl.NewVector2(x, y)
	player.rec = rl.NewRectangle(player.cnt.X-tilesize/2, player.cnt.Y-tilesize/2, tilesize, tilesize)

	player.boundaryrec = player.rec
	player.boundaryrec.X -= tilesize
	player.boundaryrec.Y -= tilesize
	player.boundaryrec.Width += tilesize * 2
	player.boundaryrec.Height += tilesize * 2

}
func makecrates() { //MARK: makecrates

	zobj := xobj{}
	zobj.name = "movable destructible wooden crate"
	zobj.rec = rl.NewRectangle(scrcnt.X, scrcnt.Y, tilesize, tilesize)

	zobj.rec.X += tilesize * 2

	zobj.canmove = true
	zobj.solid = true
	zobj.destructible = true
	zobj.hp = 2
	zobj.weight = 1
	zobj.color = randomgrey()
	zobj.fade = rFloat32(0.7, 0.95)
	zobj.img = crateimg

	rooms[currentroom].objs = append(rooms[currentroom].objs, zobj)

}

func makemoveblocks() { //MARK: makemoveblocks

	zobj := xobj{}
	zobj.name = "movable indestructible titanium block"
	zobj.rec = rl.NewRectangle(scrcnt.X, scrcnt.Y, tilesize, tilesize)

	zobj.rec.X += tilesize * 2

	zobj.canmove = true
	zobj.solid = true
	zobj.weight = 1
	zobj.color = randomgrey()
	zobj.fade = rFloat32(0.7, 0.95)
	zobj.img = blockmoveimg

	zobj.rec.X += tilesize * 2
	rooms[currentroom].objs = append(rooms[currentroom].objs, zobj)

}

func makeweapons() { //MARK: makeweapons

	zitem := makerandomweapon(uh + uq)

	zitem.rec.X = scrcnt.X
	zitem.rec.Y = scrcnt.Y
	zitem.rec.Y += u2

	zitem.numtype = 1

	rooms[currentroom].items = append(rooms[currentroom].items, zitem)

}
func makerandomweapon(width float32) xitem { //MARK: makeweapons

	zgen := xitem{}
	zgen.color = rl.White
	zgen.fade = 1.0
	zgen.rec = rl.NewRectangle(0, 0, width, width)

	zgen.img = swordimgs[rInt(0, len(swordimgs))]

	zgen.sockets = rInt(1, 4)

	zgen.atk = rFloat32(0, 11)
	zgen.dur = rFloat32(0, 11)
	zgen.pois = rFloat32(0, 11)
	zgen.fire = rFloat32(0, 11)
	zgen.ice = rFloat32(0, 11)
	zgen.elec = rFloat32(0, 11)
	zgen.vamp = rFloat32(0, 11)
	zgen.thorns = rFloat32(0, 11)
	zgen.regen = rFloat32(0, 11)
	zgen.invis = rFloat32(0, 11)
	zgen.diseas = rFloat32(0, 11)
	zgen.shield = rFloat32(0, 11)

	zgen.name1 = "sword"
	zgen.name2 = "flaming"
	zgen.name3 = "might"
	zgen.itemname = zgen.name2 + " " + zgen.name1 + " of " + zgen.name3

	return zgen

}
func makegenericitem(width float32) xitem { //MARK: makegenericextra

	zgen := xitem{}
	zgen.color = rl.White
	zgen.fade = 1.0
	zgen.rec = rl.NewRectangle(0, 0, width, width)
	return zgen

}
func makegenericextra(width float32) xextra { //MARK: makegenericextra

	zgen := xextra{}
	zgen.color = rl.White
	zgen.fade = 1.0
	zgen.rec = rl.NewRectangle(0, 0, width, width)
	return zgen

}
func makegenericobj(width float32) xobj { //MARK: makegenericextra

	zgen := xobj{}
	zgen.color = rl.White
	zgen.fade = 1.0
	zgen.rec = rl.NewRectangle(0, 0, width, width)

	return zgen

}
func makeimgs() { //MARK: makeimgs

	//FLOOR IMGS
	x := float32(0)
	y := float32(16)

	for {

		floorimgs = append(floorimgs, rl.NewRectangle(x, y, 16, 16))
		x += 16
		if x > 736 {
			break
		}
	}

	//SWORD IMGS
	x = 1130
	y = 0

	for {

		swordimgs = append(swordimgs, rl.NewRectangle(x, y, 32, 32))
		x += 32
		if x > 1290 {
			y += 32
			x = 1130
		}
		if y > 128 {
			break
		}
	}

	//PORTRAIT IMGS
	x = 1344
	y = 644

	for {

		portraitimgs = append(portraitimgs, rl.NewRectangle(x, y, 64, 64))
		x += 64
		if x > 1536 {
			y += 64
			x = 1344
		}
		if y > 836 {
			break
		}
	}

	x = 1344
	y = 512

	for {

		portraitimgs = append(portraitimgs, rl.NewRectangle(x, y, 32, 32))
		x += 32
		if x > 1568 {
			y += 32
			x = 1344
		}
		if y > 608 {
			break
		}
	}

	x = 1330
	y = 459

	for {

		portraitimgs = append(portraitimgs, rl.NewRectangle(x, y, 45, 45))
		x += 45
		if x > 1555 {
			break
		}

	}

}

// MARK: FIND FIND FIND FIND FIND FIND FIND FIND FIND FIND FIND FIND FIND FIND FIND FIND FIND
func findpointinroom(roomnum, roomrecnum int, rec rl.Rectangle) rl.Vector2 { //MARK: findpointinroom

	x := rooms[roomnum].roomrecs[roomrecnum].rec.X + tilesize
	x += rFloat32(0, rooms[roomnum].roomrecs[roomrecnum].rec.Width-(tilesize*2))

	y := rooms[roomnum].roomrecs[roomrecnum].rec.Y + tilesize
	y += rFloat32(0, rooms[roomnum].roomrecs[roomrecnum].rec.Height-(tilesize*2))

	v2 := rl.NewVector2(x, y)

	return v2

}
func findinvemptynum() bool { //MARK: findinvemptynum

	isfull := true
	for a := 0; a < len(inv); a++ {

		if inv[a].name1 == "" {
			invemptynum = a
			isfull = false
			break
		}

	}

	return isfull

}

// MARK: CHECKS CHECKS CHECKS CHECKS CHECKS CHECKS CHECKS CHECKS CHECKS CHECKS CHECKS CHECKS
func checkmouse() { //MARK: checkmouse

	if rl.IsMouseButtonPressed(rl.MouseLeftButton) && !invon {
		selpoint = mousev2world
		selrec = rl.NewRectangle(selpoint.X-tilesize/4, selpoint.Y-tilesize/4, tilesize/2, tilesize/2)
	}

	if rl.IsMouseButtonPressed(rl.MouseRightButton) {
		selpointr = mousev2world
		if !player.atkanim {
			player.atktimer = fps
			player.atkanim = true
		}
	}

}

func checkaddroomrec(rec rl.Rectangle, room xroom) bool { //MARK: checkaddroomrec

	canadd := true

	v1 := rl.NewVector2(rec.X, rec.Y)
	v2 := rl.NewVector2(rec.X+rec.Width, rec.Y)
	v3 := rl.NewVector2(rec.X+rec.Width, rec.Y+rec.Height)
	v4 := rl.NewVector2(rec.X, rec.Y+rec.Height)

	if !rl.CheckCollisionPointRec(v1, borderrecinner) || !rl.CheckCollisionPointRec(v2, borderrecinner) || !rl.CheckCollisionPointRec(v3, borderrecinner) || !rl.CheckCollisionPointRec(v4, borderrecinner) {
		canadd = false
	}

	for a := 0; a < len(room.roomrecs); a++ {
		if rl.CheckCollisionRecs(rec, room.roomrecs[a].rec) {
			canadd = false
		}
	}

	return canadd
}
func checkplayermove() bool { //MARK: checkplayermove
	canmove := true

	for a := 0; a < len(borderwallrecs); a++ {
		if rl.CheckCollisionRecs(player.collisrec, borderwallrecs[a].rec) {
			canmove = false
		}
	}

	for a := 0; a < len(rooms[currentroom].objs); a++ {

		if rl.CheckCollisionRecs(player.collisrec, rooms[currentroom].objs[a].rec) {
			if rooms[currentroom].objs[a].canmove {
				if moveobjwithplayer(a) {
					canmove = false
				}
			}

		}

	}

	return canmove

}

func checkplayercollis() { //MARK: checkobjmovecollis

	for a := 0; a < len(rooms[currentroom].extras); a++ {

		if rl.CheckCollisionRecs(player.collisrec, rooms[currentroom].extras[a].collisrec) {
			upplayerextracollis(a)
		}

	}

}

// MARK: MOVE MOVE MOVE MOVE MOVE MOVE MOVE MOVE MOVE MOVE MOVE MOVE MOVE MOVE MOVE MOVE MOVE
func moveobjwithplayer(objnum int) bool { //MARK: moveobjwithplayer

	stopmove := false

	if rooms[currentroom].objs[objnum].weight <= player.str {

		for a := 0; a < len(borderwallrecs); a++ {
			checkrec := rooms[currentroom].objs[objnum].rec
			checkrec.X += player.dirx
			checkrec.Y += player.diry

			if rl.CheckCollisionRecs(checkrec, borderwallrecs[a].rec) {
				stopmove = true
			}
		}

		if !stopmove {

			for a := 0; a < len(rooms[currentroom].objs); a++ {

				checkrec := rooms[currentroom].objs[objnum].rec
				checkrec.X += player.dirx
				checkrec.Y += player.diry

				if rl.CheckCollisionRecs(checkrec, rooms[currentroom].objs[a].rec) && a != objnum {

					stopmove = true

					rooms[currentroom].objs[objnum].stop = true
					rooms[currentroom].objs[objnum].stopv2 = rl.NewVector2(rooms[currentroom].objs[objnum].rec.X, rooms[currentroom].objs[objnum].rec.Y)
				}
			}

		}

		if !stopmove {
			rooms[currentroom].objs[objnum].rec.X += player.dirx
			rooms[currentroom].objs[objnum].rec.Y += player.diry
		}
	}

	return stopmove

}

// MARK: UPDATE UPDATE UPDATE UPDATE UPDATE UPDATE UPDATE UPDATE UPDATE UPDATE UPDATE UPDATE
func update() { //MARK: update

	inp()
	timers()

	if !pause {
		upplayer()
		uproomobjs()
		upfx()
	} else {
		if invon {
			upinv()
		}

	}

}
func upplayerextracollis(extranum int) { //MARK: upplayerextracollis

	switch rooms[currentroom].extras[extranum].img {

	case flamespoutimg:
		if !rooms[currentroom].extras[extranum].onoff {
			rooms[currentroom].extras[extranum].img2 = fireimg
			rooms[currentroom].extras[extranum].onoff = true
			rooms[currentroom].extras[extranum].timer = fps * 4
		}

	}

}
func uproomextras(extranum int) { //MARK: uproomextras

	if rooms[currentroom].extras[extranum].timer > 0 {
		rooms[currentroom].extras[extranum].timer--
		if rooms[currentroom].extras[extranum].timer == 1 {
			rooms[currentroom].extras[extranum].onoff = false
		}

		switch rooms[currentroom].extras[extranum].img {
		case flamespoutimg:
			if rooms[currentroom].extras[extranum].timer > 0 {

				destrec := rooms[currentroom].extras[extranum].rec
				destrec.X -= tilesize / 2
				destrec.Y -= tilesize / 2
				destrec.Width += tilesize
				destrec.Height += tilesize

				rl.DrawTexturePro(imgs, rooms[currentroom].extras[extranum].img2, destrec, origin, 0, rl.White)
				if frames%3 == 0 {
					rooms[currentroom].extras[extranum].img2.X += 64
					if rooms[currentroom].extras[extranum].img2.X > 512 {
						rooms[currentroom].extras[extranum].img2.X = 0
					}
				}
			}

		}

	}

}
func uproomobjs() { //MARK: uproomobjs

	for a := 0; a < len(rooms[currentroom].objs); a++ {

		//CLEAR OBJ STOP
		checkv2 := rl.NewVector2(rooms[currentroom].objs[a].rec.X, rooms[currentroom].objs[a].rec.Y)
		if rooms[currentroom].objs[a].stopv2 != checkv2 {
			rooms[currentroom].objs[a].stop = false
		}

		//CHECK PLAYER BOUNDARY REC OBJ COLLISION ACTIONS
		if rl.CheckCollisionRecs(player.boundaryrec, rooms[currentroom].objs[a].rec) {

			if rooms[currentroom].objs[a].destructible {
				if rl.IsMouseButtonPressed(rl.MouseRightButton) {

					if player.atktimer == 0 {
						upobjhp(1, a)
					}

				}
			}

		}

		//TIMERS
		if rooms[currentroom].objs[a].hplosspause > 0 {
			rooms[currentroom].objs[a].hplosspause--
		}

		//MOUSE OBJ COLLISIONS
		if rl.CheckCollisionPointRec(mousev2world, rooms[currentroom].objs[a].rec) {
			addmsg(rooms[currentroom].objs[a].name, true)

		}

	}

}
func upobjhp(numtype, objnum int) { //MARK: upobjhp

	switch numtype {
	case 1: //PLAYER ATK
		if rooms[currentroom].objs[objnum].hplosspause == 0 {
			rooms[currentroom].objs[objnum].hplosspause = fps
			rooms[currentroom].objs[objnum].hp -= player.atk
			if rooms[currentroom].objs[objnum].hp <= 0 {
				destroyobj(1, objnum)
			} else {
				rooms[currentroom].objs[objnum].hplossamouttxt = "-" + fmt.Sprint(player.atk)

				rooms[currentroom].objs[objnum].hplosstxtv2 = rl.NewVector2(rooms[currentroom].objs[objnum].rec.X+rooms[currentroom].objs[objnum].rec.Width/2-float32(txts/2), rooms[currentroom].objs[objnum].rec.Y)
			}
		}

	}

}
func upplayer() { //MARK: upplayer

	if selrec != blankrec && !rl.CheckCollisionPointRec(player.cnt, selrec) {

		xdiff := absdiff32(player.cnt.X, selpoint.X)
		ydiff := absdiff32(player.cnt.Y, selpoint.Y)

		if ydiff >= xdiff {
			player.diry = player.vel
			player.direc = 3
			if selpoint.Y < player.cnt.Y {
				player.diry = -player.diry
				player.direc = 1
			}
			xchange := xdiff / (ydiff / player.vel)
			player.dirx = xchange
			if selpoint.X < player.cnt.X {
				player.dirx = -player.dirx
			}

		} else {
			player.dirx = player.vel
			player.direc = 2
			if selpoint.X < player.cnt.X {
				player.dirx = -player.dirx
				player.direc = 4
			}
			ychange := ydiff / (xdiff / player.vel)
			player.diry = ychange
			if selpoint.Y < player.cnt.Y {
				player.diry = -player.diry
			}
		}

		player.collisrec = player.rec
		player.collisrec.X += player.dirx
		player.collisrec.Y += player.diry

		if checkplayermove() { //CHECK COLLISIONS

			player.cnt.X += player.dirx
			player.cnt.Y += player.diry
			player.rec = rl.NewRectangle(player.cnt.X-tilesize/2, player.cnt.Y-tilesize/2, tilesize, tilesize)
		} else {
			player.direc = 0
		}

		player.boundaryrec = player.rec
		player.boundaryrec.X -= tilesize
		player.boundaryrec.Y -= tilesize
		player.boundaryrec.Width += tilesize * 2
		player.boundaryrec.Height += tilesize * 2

		checkplayercollis()

	} else {
		player.direc = 0
		player.collisrec = player.rec
	}

	//PLAYER ANIMATIONS
	if frames%3 == 0 {

		if player.atkanim {

			patkimgl.X += 32
			if patkimgl.X > 1490 {
				patkimgl.X = 1330
				player.atkanim = false
			}

			patkimgd.X += 32
			if patkimgd.X > 1490 {
				patkimgd.X = 1330
				player.atkanim = false
			}

			patkimgu.X += 32
			if patkimgu.X > 1490 {
				patkimgu.X = 1330
				player.atkanim = false
			}

			patkimgr.X += 32
			if patkimgr.X > 1490 {
				patkimgr.X = 1330
				player.atkanim = false

			}

		} else {
			switch player.direc {
			case 2:
				pwalkimgr.X += 32
				if pwalkimgr.X > 1554 {
					pwalkimgr.X = 1330
				}
			case 4:
				pwalkimgl.X += 32
				if pwalkimgl.X > 1554 {
					pwalkimgl.X = 1330
				}
			case 1:
				pwalkimgu.X += 32
				if pwalkimgu.X > 1554 {
					pwalkimgu.X = 1330
				}
			case 3:
				pwalkimgd.X += 32
				if pwalkimgd.X > 1554 {
					pwalkimgd.X = 1330
				}
			case 0:
				pidleimg.X += 32
				if pidleimg.X > 1426 {
					pidleimg.X = 1330
				}
			}
		}
	}

	//PLAYER TIMERS
	if player.atktimer > 0 {
		player.atktimer--
	}

}
func upcams() { //MARK: upcams

	camera.Target = scrcnt
	camera.Offset.X = scrwf32 / 2
	camera.Offset.Y = scrhf32 / 2

}

func upfx() { //MARK: upfx

	if len(rooms[currentroom].fx) > 0 {

		for a := 0; a < len(rooms[currentroom].fx); a++ {

			switch rooms[currentroom].fx[a].numtype {
			case 1:
				if frames%2 == 0 {
					rooms[currentroom].fx[a].img.X += 64
					if rooms[currentroom].fx[a].img.X > 512 {
						rooms[currentroom].fx = remfx(rooms[currentroom].fx, a)
					}
				}
			}
		}

	}

}

func upinv() { //MARK: upinv

	//OPEN CLOSE
	if !invopen && !invclose {
		if invrec.X > scrwf32-(scrwf32/3) {
			invrec.X -= 50
			invtabrec = rl.NewRectangle(invrec.X-(u1h), invrec.Y+infobarrec.Height+u1, u1h, u1)

		} else if invrec.X <= scrwf32-(scrwf32/3) {
			if invrec.X < scrwf32-(scrwf32/3) {
				invrec.X = scrwf32 - (scrwf32 / 3)
			}
			invtabrec = rl.NewRectangle(invrec.X-(u1h), invrec.Y+infobarrec.Height+u1, u1h, u1)
			invopen = true
		}
	} else if !invopen && invclose {

		if invrec.X < scrwf32 {
			invrec.X += 50
			invtabrec = rl.NewRectangle(invrec.X-(u1h-1), invrec.Y+infobarrec.Height+u1, u1h, u1)
		} else if invrec.X >= scrwf32 {

			if invrec.X > scrwf32 {
				invrec.X = scrwf32
			}

			invtabrec = rl.NewRectangle(invrec.X-(u1h-1), invrec.Y+infobarrec.Height+u1, u1h, u1)
			invclose = false
			invon = false
			pause = false

		}

	}

}
func upscreenressizes() { //MARK: upscreenrestiles

	switch scrwf32 {

	case 2560:
		ubase = float32(42)
		txtbase = int32(20)
	}

	txtss = txtbase / 2
	txts = txtbase
	txtm = txtbase * 2
	txtl = txtbase * 3
	txtl2 = txtbase * 4
	txtxl = txtbase * 6
	txtxl2 = txtbase * 8

	tilesize = ubase
	floortilesize = tilesize

	uq = ubase / 4
	uh = ubase / 2
	ue = ubase / 8

	u1 = ubase
	u1h = u1 + uh
	u1q = u1 + uq

	u2 = ubase * 2
	u2h = u2 + uh
	u2q = u2 + uq

	u3 = ubase * 3
	u3h = u3 + uh
	u3q = u3 + uq

	u4 = ubase * 4
	u4h = u4 + uh
	u4q = u4 + uq

	u5 = ubase * 5
	u5h = u5 + uh
	u5q = u5 + uq

	u6 = ubase * 6
	u6h = u6 + uh
	u6q = u6 + uq

	u7 = ubase * 7
	u7h = u7 + uh
	u7q = u7 + uq

	u8 = ubase * 8
	u8h = u8 + uh
	u8q = u8 + uq

	u9 = ubase * 9
	u9h = u9 + uh
	u9q = u9 + uq

	u10 = ubase * 10
	u10h = u10 + uh
	u10q = u10 + uq

	//UNITS INT32
	nbase = int32(ubase)

	nq = nbase / 4
	nh = nbase / 2
	ne = nbase / 8

	n1 = nbase
	n1h = n1 + nh
	n1q = n1 + nq

	n2 = nbase * 2
	n2h = n2 + nh
	n2q = n2 + nq

	n3 = nbase * 3
	n3h = n3 + nh
	n3q = n3 + nq

	n4 = nbase * 4
	n4h = n4 + nh
	n4q = n4 + nq

	n5 = nbase * 5
	n5h = n5 + nh
	n5q = n5 + nq

	n6 = nbase * 6
	n6h = n6 + nh
	n6q = n6 + nq

	n7 = nbase * 7
	n7h = n7 + nh
	n7q = n7 + nq

	n8 = nbase * 8
	n8h = n8 + nh
	n8q = n8 + nq

	n9 = nbase * 9
	n9h = n9 + nh
	n9q = n9 + nq

	n10 = nbase * 10
	n10h = n10 + nh
	n10q = n10 + nq

}

// MARK: ETC ETC ETC ETC ETC ETC ETC ETC ETC ETC ETC ETC ETC ETC ETC ETC ETC ETC ETC ETC ETC
func destroyobj(numtype, objnum int) { //MARK: destroyobj

	switch numtype {
	case 1: //PLAYER ATK

		switch rooms[currentroom].objs[objnum].img {
		case crateimg:

			zfx := xfx{}
			zfx.numtype = 1
			zfx.img = yellowexplosionimg
			zfx.color = rl.White
			zfx.fade = 0.2
			zfx.rec = rooms[currentroom].objs[objnum].rec
			zfx.rec.X -= uh
			zfx.rec.Y -= uh
			zfx.rec.Width += u1
			zfx.rec.Height += u1
			rooms[currentroom].fx = append(rooms[currentroom].fx, zfx)

			rooms[currentroom].objs = remobj(rooms[currentroom].objs, objnum)

		}

	}

}
func addmsg(txt string, addloc bool) { //MARK: addmsg

	if addloc {
		txt = txt + "  @ " + fmt.Sprint(time.Now().Format("15:04")) + " room num " + fmt.Sprint(currentroom) + " location " + "x:" + fmt.Sprintf("%.0f", player.cnt.X) + "  y:" + fmt.Sprintf("%.0f", player.cnt.Y)
	}
	msgs = append(msgs, txt)
	msgtimer = fps * 2

}

func addtoinv(itemnum int) { //MARK: addtoinv

	if findinvemptynum() {
		addmsg("inventory is full - sell, drop or destroy items", true)
	} else {
		inv[invemptynum] = rooms[currentroom].items[itemnum]
		rooms[currentroom].items = remitem(rooms[currentroom].items, itemnum)
	}

}

// MARK: CORE CORE CORE CORE CORE CORE CORE CORE CORE CORE CORE CORE CORE CORE CORE CORE CORE
func initial() { //MARK: initial

	invon = true
	pause = true

	//dev2 = true

	//MAKE INV
	for a := 0; a < 345; a++ {
		inv = append(inv, xitem{})
	}
	for a := 0; a < 15; a++ {
		invequip = append(invequip, xitem{})
	}

	themecol1 = brightorange()
	themecol2 = darkred()

	upscreenressizes()

	infobarrec = rl.NewRectangle(0, 0, scrwf32, uh+uq)
	infobarrec2 = infobarrec
	infobarrec2.Y = scrhf32 - infobarrec2.Height

	invtabrec = rl.NewRectangle(scrwf32-u1q, infobarrec.Y+infobarrec.Height+u1, u1h, u1)
	invrec = rl.NewRectangle(scrwf32, 0, scrwf32/3, scrhf32)
	invtabrec = rl.NewRectangle(invrec.X-(u1h-1), invrec.Y+infobarrec.Height+u1, u1h, u1)

	camera.Zoom = 1.0
	selpoint = blankv2
	selrec = blankrec

	upcams()

	borderrec = rl.NewRectangle(0, 0, scrwf32, scrhf32)
	borderrecinner = borderrec
	borderrecinner.X += tilesize
	borderrecinner.Y += tilesize
	borderrecinner.Width -= tilesize * 2
	borderrecinner.Height -= tilesize * 2

	maxuw = scrwint / int(tilesize)
	maxuh = scrhint / int(tilesize)

	makeroom()
	makeplayer()
}

func inp() { //MARK: inp

	//DEV KEYS

	if rl.IsKeyPressed(rl.KeyLeftControl) {
		rooms = nil
		makeroom()
	}

	if rl.IsKeyPressed(rl.KeyF1) {
		if dev {
			dev = false
		} else {
			dev = true
		}
	}
	if rl.IsKeyPressed(rl.KeyF2) {
		if dev2 {
			dev2 = false
		} else {
			dev2 = true
		}
	}

	if rl.IsKeyPressed(rl.KeyKpSubtract) {

		if camera.Zoom == 2.0 {
			camera.Zoom = 1.5
		} else if camera.Zoom == 1.5 {
			camera.Zoom = 1.0
		} else if camera.Zoom == 1.0 {
			camera.Zoom = 0.5
		}

	}
	if rl.IsKeyPressed(rl.KeyKpAdd) {

		if camera.Zoom == 0.5 {
			camera.Zoom = 1.0
		} else if camera.Zoom == 1.0 {
			camera.Zoom = 1.5
		} else if camera.Zoom == 1.5 {
			camera.Zoom = 2.0
		}
	}
}
func timers() { //MARK: timers

	if msgtimer > 0 {
		msgtimer--
	}
	if msgdelay > 0 {
		msgdelay--
	}

	if frames%30 == 0 {
		if time30 {
			time30 = false
		} else {
			time30 = true
		}
	}
	if frames%20 == 0 {
		if time20 {
			time20 = false
		} else {
			time20 = true
		}
	}
	if frames%10 == 0 {
		if time10 {
			time10 = false
		} else {
			time10 = true
		}
	}

	if fadeblinkon {
		if fadeblink > 0.4 {
			fadeblink -= 0.05
		} else {
			fadeblinkon = false
		}
	} else {
		if fadeblink < 0.9 {
			fadeblink += 0.05
		} else {
			fadeblinkon = true
		}
	}

}

func raylib() { //MARK: raylib
	//rl.SetConfigFlags(rl.FlagMsaa4xHint) // enable 4X anti-aliasing

	rl.InitWindow(scrw, scrh, "GAME TITLE")

	//rl.InitAudioDevice()

	rl.SetExitKey(rl.KeyEnd) // key to end the game and close window

	//rl.SetWindowSize(scrwint, scrhint)

	//rl.ToggleFullscreen()

	rl.HideCursor()
	imgs = rl.LoadTexture("data/imgs.png") // load images
	makeimgs()
	initial()
	rl.SetTargetFPS(fps)

	for !rl.WindowShouldClose() {
		frames++

		mousev2 = rl.GetMousePosition()
		mousev2world = rl.GetScreenToWorld2D(mousev2, camera)

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		nocambackg()

		rl.BeginMode2D(camera)

		cam()

		rl.EndMode2D()

		nocam()

		if dev {
			devui()
		}

		//CURSOR

		cursorv1 := rl.NewVector2(mousev2.X+20, mousev2.Y+5)
		cursorv2 := rl.NewVector2(mousev2.X+5, mousev2.Y+20)
		shadow1 := cursorv1
		shadow1.Y += 5
		shadow1.X -= 2
		shadow2 := cursorv2
		shadow2.Y += 5
		shadow2.X -= 2
		shadow3 := mousev2
		shadow3.Y += 5
		shadow3.X -= 2

		rl.DrawTriangle(shadow1, shadow3, shadow2, rl.Fade(rl.Black, 0.5))
		rl.DrawTriangle(cursorv1, mousev2, cursorv2, rl.Fade(rl.Magenta, fadeblink))
		rl.DrawTriangleLines(cursorv1, mousev2, cursorv2, rl.Black)

		if selectedinv.name1 != "" {

			rec := selectedinv.rec
			rec.X = mousev2.X
			rec.Y = mousev2.Y + uh

			rl.DrawTexturePro(imgs, selectedinv.img, rec, origin, 0, rl.Fade(rl.White, 0.7))
		}

		rl.EndDrawing()
		update()
	}

	//	rl.StopSoundMulti()
	//	rl.CloseAudioDevice()

	rl.CloseWindow()

}
func screen() { //MARK: screen

	rl.InitWindow(0, 0, "")
	scrhint = rl.GetScreenHeight()
	scrwint = rl.GetScreenWidth()
	rl.CloseWindow()

	scrhf32 = float32(scrhint)
	scrwf32 = float32(scrwint)
	scrh = int32(scrhint)
	scrw = int32(scrwint)

	scrcnt = rl.NewVector2(scrwf32/2, scrhf32/2)

	camera.Zoom = 1.0

}
func main() { //MARK: main
	rand.Seed(time.Now().UnixNano()) // random numbers
	rl.SetTraceLog(rl.LogError)      // hides info window
	screen()
	raylib()
}

// MARK: COLORS
// https://www.rapidtables.com/web/color/RGB_Color.html
func darkred() rl.Color {
	color := rl.NewColor(55, 0, 0, 255)
	return color
}
func semidarkred() rl.Color {
	color := rl.NewColor(70, 0, 0, 255)
	return color
}
func brightorange() rl.Color {
	color := rl.NewColor(253, 95, 0, 255)
	return color
}
func brightred() rl.Color {
	color := rl.NewColor(230, 0, 0, 255)
	return color
}
func randomgrey() rl.Color {
	color := rl.NewColor(uint8(rInt(160, 193)), uint8(rInt(160, 193)), uint8(rInt(160, 193)), 255)
	return color
}
func randombrown() rl.Color {
	color := rl.NewColor(uint8(rInt(100, 200)), uint8(rInt(50, 100)), uint8(0), 255)
	return color
}
func randombluelight() rl.Color {
	color := rl.NewColor(uint8(rInt(0, 180)), uint8(rInt(120, 256)), uint8(rInt(120, 256)), 255)
	return color
}
func randombluedark() rl.Color {
	color := rl.NewColor(0, 0, uint8(rInt(120, 250)), 255)
	return color
}
func randomyellow() rl.Color {
	color := rl.NewColor(255, uint8(rInt(150, 256)), 0, 255)
	return color
}
func randomorange() rl.Color {
	color := rl.NewColor(uint8(rInt(250, 256)), uint8(rInt(60, 210)), 0, 255)
	return color
}
func randomred() rl.Color {
	color := rl.NewColor(uint8(rInt(128, 256)), uint8(rInt(0, 129)), uint8(rInt(0, 129)), 255)
	return color
}
func randomgreen() rl.Color {
	color := rl.NewColor(uint8(rInt(0, 170)), uint8(rInt(100, 256)), uint8(rInt(0, 50)), 255)
	return color
}
func randomcolor() rl.Color {
	color := rl.NewColor(uint8(rInt(0, 256)), uint8(rInt(0, 256)), uint8(rInt(0, 256)), 255)
	return color
}
func brightyellow() rl.Color {
	color := rl.NewColor(uint8(255), uint8(255), uint8(0), 255)
	return color
}
func brightbrown() rl.Color {
	color := rl.NewColor(uint8(218), uint8(165), uint8(32), 255)
	return color
}
func brightgrey() rl.Color {
	color := rl.NewColor(uint8(212), uint8(212), uint8(213), 255)
	return color
}

// MARK: random numbers
func rInt(min, max int) int {
	return rand.Intn(max-min) + min
}
func rInt32(min, max int) int32 {
	i := rand.Intn(max-min) + min
	return int32(i)
}
func rFloat32(min, max float32) float32 {
	return (rand.Float32() * (max - min)) + min
}
func flipcoin() bool {
	var b bool
	a := rInt(0, 10001)
	if a < 5000 {
		b = true
	}
	return b
}
func rolldice() int {
	a := rInt(1, 7)
	return a
}

// MARK: other functions
func orbitpoint(cnt, point rl.Vector2, angle float32) rl.Vector2 {

	//THIS DOES NOT WORK AS A FUNCTION COPY & PASTE TO RELEVANT PART

	angle = angle * (math.Pi / 180)

	newx := float32(math.Cos(float64(angle)))*(point.X-cnt.X) - float32(math.Sin(float64(angle)))*(point.Y-cnt.Y) + cnt.X

	newy := float32(math.Sin(float64(angle)))*(point.X-cnt.X) + float32(math.Cos(float64(angle)))*(point.Y-cnt.Y) + cnt.Y

	point2 := rl.NewVector2(newx, newy)

	return point2

}
func distancebetweentwopoints(v1, v2 rl.Vector2) float32 {

	num := float32(0)

	x2 := getabs(v2.X)
	x1 := getabs(v1.X)

	y2 := getabs(v2.Y)
	y1 := getabs(v1.Y)

	num = squareroot(squarenum(x2-x1) + squarenum(y2-y1))

	return num
}
func squareroot(num float32) float32 {
	num = float32(math.Sqrt(float64(num)))
	return num
}
func squarenum(num float32) float32 { // num*num
	num = num * num
	return num
}
func lastdigits(num int) int {
	number := num % 1e2 //change 1e2 to 1e3 to 1e4 etc for more digit places
	return number
}
func firstdigits(num int) int {
	number := num / 1e3 //change 1e2 to 1e3 to 1e4 etc for more digit places
	return number
}
func timehere(x, y float32) {
	currentTime := time.Now()
	txtlen := rl.MeasureText(currentTime.Format("15:04"), txts)
	x -= float32(txtlen + txts)
	rl.DrawText(currentTime.Format("15:04"), int32(x-1), int32(y+1), txts, rl.Black)
	rl.DrawText(currentTime.Format("15:04"), int32(x), int32(y), txts, themecol1)
}
func getabs(num float32) float32 {
	return float32(math.Abs(float64(num)))
}
func absdiff32(num, num2 float32) float32 {

	diff := float32(0)
	if num == num2 {
		diff = 0
	} else if num == 0 || num2 == 0 {
		if num == 0 {
			diff = float32(math.Abs(float64(num2)))
		} else {
			diff = float32(math.Abs(float64(num)))
		}
	} else if num > 0 && num2 > 0 {
		if num > num2 {
			diff = num - num2
		} else {
			diff = num2 - num
		}
	} else if num > 0 && num2 < 0 || num < 0 && num2 > 0 {

		if num > 0 {
			diff = num + float32(math.Abs(float64(num2)))
		} else {
			diff = num2 + float32(math.Abs(float64(num)))
		}

	} else if num < 0 && num2 < 0 {
		num = float32(math.Abs(float64(num)))
		num2 = float32(math.Abs(float64(num2)))
		if num > num2 {
			diff = num - num2
		} else {
			diff = num2 - num
		}
	}

	return diff

}
func angle2points(start, destination rl.Vector2) float32 { //make sure destination vector is vec2
	angle := float32(math.Atan2(float64(destination.Y-start.Y), float64(destination.X-start.X)))*(180/math.Pi) + 90
	//change +30 (addition value at end) to angle to compensate for polygon rotation difference
	return angle

}
func diagsquare(sidelength float32) float32 {
	return sidelength * float32(math.Sqrt(2))
}
func circlearea(radius float32) float32 {
	return math.Pi * radius * radius
}
func gcd(num1, num2 float32) float32 { //greatest common divisor

	num164 := float64(num1)
	num264 := float64(num2)

	//Calculate GCD
	num3 := math.Mod(num164, num264)

	for num3 > 0 {

		num164 = num264
		num264 = num3
		num3 = math.Mod(num164, num264)

	}

	return float32(num264)
}

func remobj(s []xobj, index int) []xobj { //remove struct from a slice
	return append(s[:index], s[index+1:]...)
}
func remfx(s []xfx, index int) []xfx { //remove struct from a slice
	return append(s[:index], s[index+1:]...)
}
func remitem(s []xitem, index int) []xitem { //remove struct from a slice
	return append(s[:index], s[index+1:]...)
}

/*
func remdedmons(s []xdedmonster, index int) []xdedmonster { //remove struct from a slice
	return append(s[:index], s[index+1:]...)
}
func remcirc(s []xcircle, index int) []xcircle { //remove struct from a slice
	return append(s[:index], s[index+1:]...)
}

func remstring(s []string, index int) []string { //remove string from a slice
	return append(s[:index], s[index+1:]...)
}
*/
