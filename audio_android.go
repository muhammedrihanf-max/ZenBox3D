// +build android

package main

// Dummy player for when audio is disabled
type dummyPlayer struct{}

func (p *dummyPlayer) Play()             {}
func (p *dummyPlayer) Stop()             {}
func (p *dummyPlayer) SetLooping(b bool) {}
func (p *dummyPlayer) SetGain(v float32) {}
func (p *dummyPlayer) SetRolloffFactor(v float32) {}

// Audio contains all sounds and music used by the game (dummy version for Android)
type Audio struct {
	musicGame *dummyPlayer
	musicMenu *dummyPlayer
	levelDone    *dummyPlayer
	levelRestart *dummyPlayer
	levelFail    *dummyPlayer
	gameComplete    *dummyPlayer
	click *dummyPlayer
	hover *dummyPlayer
	gopherWalk      *dummyPlayer
	gopherBump      *dummyPlayer
	gopherHurt      *dummyPlayer
	gopherFallStart *dummyPlayer
	gopherFallEnd   *dummyPlayer
	boxPush      *dummyPlayer
	boxOnPad     *dummyPlayer
	boxOffPad    *dummyPlayer
	boxFallStart *dummyPlayer
	boxFallEnd   *dummyPlayer
	elevatorUp   *dummyPlayer
	elevatorDown *dummyPlayer
}

func NewAudio() *Audio {
	a := new(Audio)
	d := &dummyPlayer{}
	a.musicGame = d
	a.musicMenu = d
	a.levelDone = d
	a.levelRestart = d
	a.levelFail = d
	a.gameComplete = d
	a.click = d
	a.hover = d
	a.gopherWalk = d
	a.gopherBump = d
	a.gopherHurt = d
	a.gopherFallStart = d
	a.gopherFallEnd = d
	a.boxPush = d
	a.boxOnPad = d
	a.boxOffPad = d
	a.boxFallStart = d
	a.boxFallEnd = d
	a.elevatorUp = d
	a.elevatorDown = d
	return a
}

func NewListener(cam interface{}) {}

func (a *Audio) SetMusicVolume(vol float32) {}
func (a *Audio) SetSfxVolume(vol float32)   {}
