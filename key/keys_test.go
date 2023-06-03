package key

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPrivateKey(t *testing.T) {
	testCases := []struct {
		name string
		fn   func(t *testing.T)
	}{
		{
			name: "NewPrivateKey",
			fn: func(t *testing.T) {
				r := NewPrivateKey()
				assert.NotNilf(t, r, "NewPrivateKey() = %v, want %v", r, nil)
				assert.NotNilf(t, r.Key, "NewPrivateKey() = %v, want %v", r.Key, nil)
				assert.NotNilf(t, r.Key.D, "NewPrivateKey() = %v, want %v", r.Key.D, nil)
				assert.NotNilf(t, r.Key.PublicKey, "NewPrivateKey() = %v, want %v", r.Key.PublicKey, nil)
				assert.NotNilf(t, r.Key.PublicKey.X, "NewPrivateKey() = %v, want %v", r.Key.PublicKey.X, nil)
				assert.NotNilf(t, r.Key.PublicKey.Y, "NewPrivateKey() = %v, want %v", r.Key.PublicKey.Y, nil)
				assert.NotNilf(t, r.Key.Curve, "NewPrivateKey() = %v, want %v", r.Key.Curve, nil)

				assert.Equal(t, PRIVATEKEYLEN, len(r.Bytes()), "NewPrivateKey() = %v, want %v", len(r.Key.D.Bytes()), PRIVATEKEYLEN)

				assert.Equalf(t, r.Key.PublicKey.Curve, r.Key.Curve, "NewPrivateKey() = %v, want %v", r.Key.PublicKey.Curve, r.Key.Curve)

				assert.NotNilf(t, r.Key.PublicKey.Curve.Params(), "NewPrivateKey() = %v, want %v", r.Key.PublicKey.Curve.Params(), nil)

				assert.NotNilf(t, r.Key.PublicKey.Curve.Params().Name, "NewPrivateKey() = %v, want %v", r.Key.PublicKey.Curve.Params().Name, nil)

				assert.NotNilf(t, r.Key.PublicKey.Curve.Params().BitSize, "NewPrivateKey() = %v, want %v", r.Key.PublicKey.Curve.Params().BitSize, nil)
			},
		},
		{
			name: "Sign and Verify",
			fn: func(t *testing.T) {
				r := NewPrivateKey()
				assert.NotNilf(t, r, "NewPrivateKey() = %v, want %v", r, nil)

				msg := []byte("Hello World")
				sig, err := r.Sign(msg)
				assert.Nilf(t, err, "Sign() = %v, want %v", err, nil)
				assert.NotNilf(t, sig, "Sign() = %v, want %v", sig, nil)

				pub := r.PublicKey()
				assert.NotNilf(t, pub, "PublicKey() = %v, want %v", pub, nil)

				assert.Equal(t, PUBLICKEYLEN, len(pub.Bytes()), "PublicKey() = %v, want %v", len(pub.Key.X.Bytes()), PUBLICKEYLEN)

				assert.Truef(t, sig.Verify(pub, msg), "Verify() = %v, want %v", sig.Verify(pub, msg), true)

				assert.Falsef(t, sig.Verify(pub, []byte("Hello World!")), "Verify() = %v, want %v", sig.Verify(pub, []byte("Hello World!")), false)

				falseKey := NewPrivateKey()
				assert.NotNilf(t, falseKey, "NewPrivateKey() = %v, want %v", falseKey, nil)

				falsePub := falseKey.PublicKey()
				assert.NotNilf(t, falsePub, "PublicKey() = %v, want %v", falsePub, nil)

				assert.Falsef(t, sig.Verify(falsePub, msg), "Verify() = %v, want %v", sig.Verify(falsePub, msg), false)

			},
		},
		{
			name: "Public Address",
			fn: func(t *testing.T) {
				p := NewPrivateKey()
				assert.NotNilf(t, p, "NewPrivateKey() = %v, want %v", p, nil)

				pub := p.PublicKey()
				assert.NotNilf(t, pub, "PublicKey() = %v, want %v", pub, nil)

				addr := pub.Address()
				assert.Equal(t, ADDRESSLEN, len(addr.Hash), "Address() = %v, want %v", len(addr.Hash), 20)
				t.Logf("Address: %v", addr.String())
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, tc.fn)
	}
}
