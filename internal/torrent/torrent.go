package torrent

type Torrent struct {
	Announce    string
	Name        string
	Length      int
	PieceLength int
	Pieces      [][]byte
	InfoHash    [20]byte
}
