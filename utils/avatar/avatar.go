package avatar

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func colorHash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	hashed := h.Sum(nil)
	color := "#" + hex.EncodeToString(hashed[:3])
	return color
}

func CreateUser(hash string) []byte {
	hex := colorHash(hash)
	svg := fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 28 28"><path fill="%v" d="M0 0h28v28H0z"/><g opacity=".2"><path d="M14 19.7h-3.2a1.2 1.2 0 0 1-1.3-1.3c.2-5.8 7.9-6.4 8.9-.7.2.9 0 2-1.1 2H14Z"/><circle cx="14" cy="10.5" r="2.5"/></g></svg>`, hex)
	return []byte(svg)
}

func CreateChat(hash string) []byte {
	hex := colorHash(hash)
	svg := fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 28 28"><path fill="%v" d="M0 0h28v28H0z"/><g opacity=".2"><path d="M14 17.8h-2.5a1 1 0 0 1-1-1.1 3.5 3.5 0 0 1 7-.5c.1.7-.1 1.6-.9 1.6Z"/><circle cx="14" cy="10.6" r="2"/></g><g opacity=".2"><path d="M21.6 17.8a3.5 3.5 0 0 0-3.1-3c-.9 0-.2 1.1-.2 1.6.5 2.9-1.9 2-3.7 2.2a.9.9 0 0 0 1 .8h5.1c.8 0 1-.8.9-1.6Z"/><circle cx="18.3" cy="12.1" r="2"/></g><g opacity=".2"><path d="M10.9 18.6a1.3 1.3 0 0 1-1.3-1.4 4.8 4.8 0 0 1 .3-1.8.4.4 0 0 0-.4-.6 3.5 3.5 0 0 0-3.1 3.6.9.9 0 0 0 1 1h5.1a1 1 0 0 0 .9-.8Z"/><circle cx="9.7" cy="12.1" r="2"/></g></svg>`, hex)
	return []byte(svg)
}
