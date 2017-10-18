package library

import "errors"

type MusicEntry struct {
	Id string
	Name string
	Artist string
	Genre string
	Source string
	Type string
}

func (m *MusicEntry)Equal(m1 *MusicEntry) bool {

	return m.Name == m1.Name &&
	       m.Id == m1.Id &&
	       m.Artist == m1.Artist &&
	       m.Genre == m1.Genre &&
	       m.Source == m1.Source &&
	       m.Type == m1.Type
}

type MusicManager struct {
	musics []MusicEntry
}

func NewMusicManager() *MusicManager{
	return &MusicManager{make([]MusicEntry, 0)}
}

func (m *MusicManager)Len() int {
	return len(m.musics)
}

func (m *MusicManager)Get(index int)(music *MusicEntry, err error) {
	if index<0 || index >= len(m.musics){
		return nil, errors.New("Index out of range")
	}

	return &m.musics[index], nil
}

func (m *MusicManager)Add(music *MusicEntry){
	m.musics = append(m.musics, *music)
}

func (m *MusicManager)Find(name string) *MusicEntry{
	if len(m.musics) == 0{
		return  nil
	}

	for _, m := range m.musics{
		if m.Name == name{
			return &m
		}
	}

	return nil
}

func (m *MusicManager) Remove(index int) * MusicEntry{
	if index<0 || index >= len(m.musics) {
		return nil
	}

	removedMusic := &m.musics[index]
	if index < len(m.musics) - 1 {
		m.musics = append(m.musics[:index-1], m.musics[index +1 :]...)
	}else if index == 0 {
		m.musics = make([]MusicEntry, 0)
	}else {
		m.musics = m.musics[:index-1]
	}
	return  removedMusic
}

func (m *MusicManager) RemoveByName(name string) *MusicEntry {
	var removedMusicEntry *MusicEntry = nil
	var iPos int = -1

	for i:=0; i < m.Len(); i++{
		if m.musics[i].Name == name{
			iPos = i
			break
		}
	}

	if iPos<0{
		return nil
	}

	removedMusicEntry = m.Remove(iPos)
	return removedMusicEntry
}
