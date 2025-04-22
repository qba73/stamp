package stamp_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/qba73/stamp"
)

func TestCalculateDigestFromBytes(t *testing.T) {
	t.Parallel()

	b := []byte(`{"key":"value"}`)
	got, err := stamp.CalculateDigest(b)
	if err != nil {
		t.Fatal(err)
	}

	want := "sha256:e43abcf3375244839c012f9633f95862d232a95b00d5bc7348b3098b9fed7f32"
	if !cmp.Equal(want, got.String()) {
		t.Error(cmp.Diff(want, got.String()))
	}
}

func TestCreateNewDescriptorFromBytes(t *testing.T) {
	t.Parallel()

	b := []byte(`{"key":"value"}`)
	got, err := stamp.NewDescriptor(b)
	if err != nil {
		t.Fatal(err)
	}

	want := v1.Descriptor{
		MediaType:    "application/json",
		Digest:       "sha256:e43abcf3375244839c012f9633f95862d232a95b00d5bc7348b3098b9fed7f32",
		Size:         15,
		Data:         []uint8("eyJrZXkiOiJ2YWx1ZSJ9"),
		ArtifactType: "application/json",
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
