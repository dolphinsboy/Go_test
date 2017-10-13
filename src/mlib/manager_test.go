package library

import "testing"

func TestOps(t *testing.T) {
	mm := NewMusicManager()

	if mm == nil{
		t.Error("NewMusicManager failed.")
	}

	if mm.Len() != 0 {
		t.Error("NewMusicManager failed, not empty")
	}
	m0 := &MusicEntry{"1", "@@","CCC","POP","dddd","End"}
	mm.Add(m0)

	if mm.Len() != 1{
		t.Error("MusicManager Add failed")
	}

	m := mm.Find(m0.Name)

	if m == nil {
		t.Error("MusicManager.Find failed")
	}

	if m0.Equal(m) != true {
		t.Error("MusicManager.Find() failed. Found item mismatch")
	}

	m, err := mm.Get(0)

	if m == nil{
		t.Error("MusicManager.Get() failed.", err)
	}

	m = mm.Remove(0)

	if m == nil || mm.Len() != 0 {
		t.Error("MusicManger.Remove() failed.", err)
	}
}

