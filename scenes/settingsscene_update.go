package scenes

func (s *SettingsScene) Update() error {
	s.next = SettingsId
	if err := s.menu.Update(); err != nil {
		return err
	}

	s.elapsed += 0.1
	return nil
}
