// Stamp provides functionality for digitally signing files and other artifacts.
package stamp

import (
	"crypto/sha256"
	"encoding/base64"

	"github.com/opencontainers/go-digest"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

// CalculateDigest takes a slice of bytes and returns Digest value.
// The content arg represents JSON doc. Digest is calculated using SHA256 algorithm.
func CalculateDigest(content []byte) (digest.Digest, error) {
	h := sha256.New()
	_, err := h.Write(content)
	if err != nil {
		return "", err
	}
	return digest.NewDigestFromBytes(digest.SHA256, h.Sum(nil)), nil
}

// NewDescriptor takes a slice of bytes representing JSON doc to sign
// and returns Descriptor value.
func NewDescriptor(content []byte) (v1.Descriptor, error) {
	dg, err := CalculateDigest(content)
	if err != nil {
		return v1.Descriptor{}, err
	}
	d := v1.Descriptor{
		MediaType:    "application/json",
		Digest:       dg,
		Size:         int64(len(content)),
		Data:         []byte(base64.StdEncoding.EncodeToString(content)),
		ArtifactType: "application/json",
	}
	return d, nil
}

// Main is the entry point to the `stamp` cli.
func Main() int {
	return 0
}
