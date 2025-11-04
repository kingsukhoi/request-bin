package authentication

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/kingsukhoi/request-bin/pkg/db"
	"github.com/kingsukhoi/request-bin/pkg/sqlc"
)

type CurrKey struct {
	KeyId      uuid.UUID
	PrivateKey *ecdsa.PrivateKey
	PublicKey  *ecdsa.PublicKey
}

var currKey *CurrKey

func InitKeys(ctx context.Context) error {
	pool := db.MustGetDatabase()
	queries := sqlc.New(pool)

	latestKey, err := queries.GetLatestKey(ctx)
	if err == nil {

		// Parse the PEM keys
		privateKeyBlock, _ := pem.Decode([]byte(latestKey.PrivateKey))
		if privateKeyBlock == nil {
			return errors.New("failed to decode private key PEM")
		}

		privateKey, errI := x509.ParseECPrivateKey(privateKeyBlock.Bytes)
		if errI != nil {
			return errI
		}

		publicKeyBlock, _ := pem.Decode([]byte(latestKey.PublicKey))
		if publicKeyBlock == nil {
			return errors.New("failed to decode public key PEM")
		}

		publicKeyInterface, errI := x509.ParsePKIXPublicKey(publicKeyBlock.Bytes)
		if errI != nil {
			return errI
		}

		publicKey, ok := publicKeyInterface.(*ecdsa.PublicKey)
		if !ok {
			return errors.New("public key is not an ECDSA key")
		}

		currKey = &CurrKey{
			KeyId:      latestKey.ID,
			PrivateKey: privateKey,
			PublicKey:  publicKey,
		}

		return nil

	} else if !errors.Is(err, pgx.ErrNoRows) {
		return err
	}

	// Generate ECDSA key pair using P-384 curve (for ES384)
	privateKey, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		return err
	}

	// Encode private key to PEM
	privateKeyBytes, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return err
	}
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: privateKeyBytes,
	})

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return err
	}
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	})

	keyId, err := uuid.NewV7()
	if err != nil {
		return err
	}

	err = queries.CreateKey(ctx, sqlc.CreateKeyParams{
		ID:         keyId,
		PrivateKey: string(privateKeyPEM),
		PublicKey:  string(publicKeyPEM),
	})

	if err != nil {
		return err
	}

	currKey = &CurrKey{
		KeyId:      keyId,
		PrivateKey: privateKey,
		PublicKey:  &privateKey.PublicKey,
	}

	return nil
}

func getCurrentKey() *CurrKey {
	return currKey
}

type JwtClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenJwt(username string) (string, error) {
	currentKey := getCurrentKey()
	currentTime := time.Now()

	claims := JwtClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(currentTime.Add(12 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(currentTime),
		},
	}

	currToken := jwt.NewWithClaims(jwt.SigningMethodES384, claims)

	token, err := currToken.SignedString(currentKey.PrivateKey)
	if err != nil {
		return "", err
	}
	return token, nil
}

func VerifyJwt(jwtString string) (bool, error) {
	currentKey := getCurrentKey()

	tok, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return currentKey.PublicKey, nil
	})
	if err != nil {
		return false, err
	}

	return tok.Valid, nil
}
