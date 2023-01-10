package sps

type Store struct {
}

func (s *Store) GetChannel(id string) (*Channel, error) {
	return nil, nil
}

func (s *Store) ListChannels(userId string) ([]*Channel, error) {
	return nil, nil
}

func (s *Store) GetEpisodesByChannel(channelId string) ([]*Episode, error) {
	return nil, nil
}
