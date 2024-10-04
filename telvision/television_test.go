package telvision

import (
	"testing"
)

func TestNewTv(t *testing.T) {
	tv := NewTv()
	tv.Turn_on()
	if !tv.is_on {
		t.Errorf("tv is off is expected to be on")
	}
	tv.Turn_off()
	if tv.is_on {
		t.Errorf("tv is on is expected to be off")
	}

}
func TestSetChannel(t *testing.T) {
	tv := NewTv()
	tv.Turn_on()
	err := tv.setChannel(5)
	if err != nil {
		t.Errorf("unexpected error, %v", err)
	}
	if tv.channel != 5 {
		t.Errorf("expected 5 , got %d", tv.channel)
	}
	err = tv.setChannel(101)
	if err == nil {
		t.Errorf("expected error, got none")
	}

}
func TestTv_Volume(t *testing.T) {
	tv := NewTv()
	tv.Turn_on()
	err := tv.setVolume(5)
	if err != nil {
		t.Errorf("unexpected error, %v", err)
	}
	if tv.volume != 5 {
		t.Errorf("expected 5 , got %d", tv.volume)
	}
	err = tv.setVolume(12)
	if err == nil {
		t.Errorf("expected error, got none")
	}

}
func Test_To_Mute_And_Unmute(t *testing.T) {
	tv := NewTv()
	tv.Turn_on()
	tv.setVolume(7)
	if tv.volume != 7 {
		t.Errorf("expected 7 , got %d", tv.volume)
	}
	tv.mute()
	if tv.volume != 0 {
		t.Errorf("expected 0 , got %d", tv.volume)
	}
	tv.unMute()
	if tv.volume != 7 {
		t.Errorf("expected 7 , got %d", tv.volume)
	}
}
func Test_Volume_Is_Up(t *testing.T) {
	tv := NewTv()
	tv.Turn_on()
	tv.setVolume(7)
	tv.volume_up()
	if tv.volume != 8 {
		t.Errorf("expected 8 , got %d", tv.volume)
	}
}
func Test_Volume_Down(t *testing.T) {
	tv := NewTv()
	tv.Turn_on()
	tv.setVolume(7)
	tv.volume_down()
	if tv.volume != 6 {
		t.Errorf("expected 6 , got %d", tv.volume)
	}
}
func Test_Channel_Up(t *testing.T) {
	tv := NewTv()
	tv.Turn_on()
	tv.channel_up()
	if tv.channel != 2 {
		t.Errorf("expected 2 got %d", tv.channel)
	}
}
func Test_Channel_Down(t *testing.T) {
	tv := NewTv()
	tv.Turn_on()
	tv.setChannel(4)
	tv.channel_down()
	if tv.channel != 3 {
		t.Errorf("expected 3 got %d", tv.channel)
	}
}
