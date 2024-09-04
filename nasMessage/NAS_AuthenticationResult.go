// Code generated by generate.sh, DO NOT EDIT.

package nasMessage

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/ramprabudgl/nas/nasType"
)

type AuthenticationResult struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.AuthenticationResultMessageIdentity
	nasType.SpareHalfOctetAndNgksi
	nasType.EAPMessage
	*nasType.ABBA
}

func NewAuthenticationResult(iei uint8) (authenticationResult *AuthenticationResult) {
	authenticationResult = &AuthenticationResult{}
	return authenticationResult
}

const (
	AuthenticationResultABBAType uint8 = 0x38
)

func (a *AuthenticationResult) EncodeAuthenticationResult(buffer *bytes.Buffer) error {
	if err := binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS encode error (AuthenticationResult/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return fmt.Errorf("NAS encode error (AuthenticationResult/SpareHalfOctetAndSecurityHeaderType): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.AuthenticationResultMessageIdentity.Octet); err != nil {
		return fmt.Errorf("NAS encode error (AuthenticationResult/AuthenticationResultMessageIdentity): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.SpareHalfOctetAndNgksi.Octet); err != nil {
		return fmt.Errorf("NAS encode error (AuthenticationResult/SpareHalfOctetAndNgksi): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetLen()); err != nil {
		return fmt.Errorf("NAS encode error (AuthenticationResult/EAPMessage): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.EAPMessage.Buffer); err != nil {
		return fmt.Errorf("NAS encode error (AuthenticationResult/EAPMessage): %w", err)
	}
	if a.ABBA != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.ABBA.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (AuthenticationResult/ABBA): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ABBA.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (AuthenticationResult/ABBA): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ABBA.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (AuthenticationResult/ABBA): %w", err)
		}
	}
	return nil
}

func (a *AuthenticationResult) DecodeAuthenticationResult(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS decode error (AuthenticationResult/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return fmt.Errorf("NAS decode error (AuthenticationResult/SpareHalfOctetAndSecurityHeaderType): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.AuthenticationResultMessageIdentity.Octet); err != nil {
		return fmt.Errorf("NAS decode error (AuthenticationResult/AuthenticationResultMessageIdentity): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndNgksi.Octet); err != nil {
		return fmt.Errorf("NAS decode error (AuthenticationResult/SpareHalfOctetAndNgksi): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Len); err != nil {
		return fmt.Errorf("NAS decode error (AuthenticationResult/EAPMessage): %w", err)
	}
	if a.EAPMessage.Len < 4 || a.EAPMessage.Len > 1500 {
		return fmt.Errorf("invalid ie length (AuthenticationResult/EAPMessage): %d", a.EAPMessage.Len)
	}
	a.EAPMessage.SetLen(a.EAPMessage.GetLen())
	if err := binary.Read(buffer, binary.BigEndian, a.EAPMessage.Buffer); err != nil {
		return fmt.Errorf("NAS decode error (AuthenticationResult/EAPMessage): %w", err)
	}
	for buffer.Len() > 0 {
		var ieiN uint8
		var tmpIeiN uint8
		if err := binary.Read(buffer, binary.BigEndian, &ieiN); err != nil {
			return fmt.Errorf("NAS decode error (AuthenticationResult/iei): %w", err)
		}
		// fmt.Println(ieiN)
		if ieiN >= 0x80 {
			tmpIeiN = (ieiN & 0xf0) >> 4
		} else {
			tmpIeiN = ieiN
		}
		// fmt.Println("type", tmpIeiN)
		switch tmpIeiN {
		case AuthenticationResultABBAType:
			a.ABBA = nasType.NewABBA(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ABBA.Len); err != nil {
				return fmt.Errorf("NAS decode error (AuthenticationResult/ABBA): %w", err)
			}
			if a.ABBA.Len < 2 {
				return fmt.Errorf("invalid ie length (AuthenticationResult/ABBA): %d", a.ABBA.Len)
			}
			a.ABBA.SetLen(a.ABBA.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ABBA.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (AuthenticationResult/ABBA): %w", err)
			}
		default:
		}
	}
	return nil
}
