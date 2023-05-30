package adapter

import "fmt"

type MediaPlayer interface {
	Play(audioType, filename string)
}

type AdvancedMediaPlayer interface {
	PlayVlc(filename string)
	PlayMp4(filename string)
}

type VlcPlayer struct{}

func (v *VlcPlayer) PlayVlc(filename string) {
	fmt.Printf("Playing VLC file: %s\n", filename)
}

func (v *VlcPlayer) PlayMp4(filename string) {}

type Mp4Player struct{}

func (m *Mp4Player) PlayVlc(filename string) {}

func (m *Mp4Player) PlayMp4(filename string) {
	fmt.Printf("Playing MP4 file: %s\n", filename)
}

type MediaAdapter struct {
	advancedPlayer AdvancedMediaPlayer
}

func (a *MediaAdapter) Play(audioType, filename string) {
	if audioType == "vlc" {
		a.advancedPlayer.PlayVlc(filename)
	} else if audioType == "mp4" {
		a.advancedPlayer.PlayMp4(filename)
	}
}

type AudioPlayer struct {
	adapter MediaPlayer
}

func (a *AudioPlayer) Play(audioType, filename string) {
	if audioType == "mp3" {
		fmt.Printf("Playing MP3 file: %s\n", filename)
	} else if audioType == "vlc" || audioType == "mp4" {
		a.adapter.Play(audioType, filename)
	} else {
		fmt.Printf("Invalid media type: %s\n", audioType)
	}
}

func GetPlayer() {
	audioPlayer := &AudioPlayer{adapter: &MediaAdapter{advancedPlayer: &VlcPlayer{}}}
	audioPlayer.Play("mp3", "song.mp3")
	audioPlayer.Play("vlc", "movie.vlc")
	audioPlayer.Play("mp4", "video.mp4")
	audioPlayer.Play("avi", "video.avi")
}
