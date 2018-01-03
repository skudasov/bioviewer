package hbase

type BioRegisterProfile struct {
	Name           string   `json:"name"`
	PhotoOriginals [][]byte `json:"photos"`
	SoundOriginals [][]byte `json:"sounds"`
}

func NewBioRegisterProfile(name string, photos [][]byte, sounds [][]byte) *BioRegisterProfile {
	return &BioRegisterProfile{
		name,
		photos,
		sounds,
	}
}

type BioVerifyProfile struct {
	Name          string
	PhotoOriginal []byte
	SoundOriginal []byte
	//TODO: Определить форму профиля для верификации
}