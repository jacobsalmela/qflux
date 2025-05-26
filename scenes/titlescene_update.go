package scenes

func (s *TitleScene) Update() error {
	s.next = TitleSceneId // default to this scene
	if err := s.menu.Update(); err != nil {
		return err
	}

	s.elapsed += 0.1
	return nil
}
