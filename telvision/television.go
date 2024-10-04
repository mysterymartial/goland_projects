package telvision

import "errors"

type Tv struct {
	channel        int
	volume         int
	is_on          bool
	previousVolume int
}

func NewTv() *Tv {
	return &Tv{channel: 1, volume: 0, is_on: false, previousVolume: 0}

}
func (tv *Tv) Turn_on() {
	tv.is_on = true

}
func (tv *Tv) Turn_off() {
	tv.is_on = false
}

func (tv *Tv) setChannel(channel_number int) error {
	if !tv.is_on {
		return errors.New("Tv is off, turn on")
	}
	if channel_number < 1 || channel_number > 100 {
		return errors.New("channel number must be between 1 and 100")
	}
	tv.channel = channel_number
	return nil

}
func (tv *Tv) get_channel() (int, error) {
	if !tv.is_on {
		return 0, errors.New("tv is off, tv should be on first")
	}
	return tv.channel, nil
}

func (tv *Tv) setVolume(volume_level int) error {
	if !tv.is_on {
		return errors.New("tv is off, tv should be on first")
	}
	if volume_level < 0 || volume_level > 10 {
		return errors.New("volume level must be between 0 and 10")
	}
	tv.volume = volume_level
	return nil

}
func (tv *Tv) getVolume() (int, error) {
	if !tv.is_on {
		return 0, errors.New("tv is off, tv should be on first")
	}
	return tv.volume, nil
}

func (tv *Tv) mute() error {
	if !tv.is_on {
		return errors.New("tv is off, tv should be on first")
	}
	tv.previousVolume = tv.volume
	tv.volume = 0
	return nil

}

func (tv *Tv) unMute() error {
	if !tv.is_on {
		return errors.New("tv is off, tv should be on first")
	}
	tv.volume = tv.previousVolume
	return nil

}

func (tv *Tv) volume_up() error {
	if !tv.is_on {
		return errors.New("tv is off, tv should be on first")
	}
	if tv.volume < 10 {
		tv.volume++
	}

	return nil

}

func (tv *Tv) volume_down() error {
	if !tv.is_on {
		return errors.New("tv is off, tv should be on first")
	}
	if tv.volume < 10 {
		tv.volume--
	}
	return nil

}

func (tv *Tv) channel_up() error {
	if !tv.is_on {
		return errors.New("tv is off, tv should be on first")
	}
	if tv.channel < 100 {
		tv.channel++

	}

	return nil

}

func (tv *Tv) channel_down() error {
	if !tv.is_on {
		return errors.New("tv is off, tv should be on first")
	}
	if tv.channel < 100 {
		tv.channel--
	}
	return nil
}
