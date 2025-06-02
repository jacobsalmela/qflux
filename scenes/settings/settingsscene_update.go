package settings

import "qflux/scenes"

func (s *SettingsScene) Update() error {
	s.Next = scenes.SettingsId
	if err := s.Menu.Update(); err != nil {
		return err
	}

	s.elapsed += 0.1
	return nil
}
