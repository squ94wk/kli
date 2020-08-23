package signing

import (
	"github.com/squ94wk/kli/internal/reference"
)

func MatchKeys(keys []*reference.Key, certs []*reference.Cert) (map[*reference.Cert]*reference.Key, []*reference.Key, error) {
	pairs := make(map[*reference.Cert]*reference.Key, 0)
	leftover := make([]*reference.Key, 0)
	for _, key := range keys {
		for _, cert := range certs {
			if key.BelongsToCert(cert) {
				pairs[cert] = key
				break
			}
		}
		leftover = append(leftover, key)
	}
	return pairs, leftover, nil
}

