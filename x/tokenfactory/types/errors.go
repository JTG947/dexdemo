package types

// DONTCOVER

import (
	fmt "fmt"

	"cosmossdk.io/errors"
)


// x/tokenfactory module sentinel errors
var (
	ErrDenomExists                    = errors.Register(ModuleName, 2, "attempting to create a denom that already exists (has bank metadata)")
	ErrUnauthorized                   = errors.Register(ModuleName, 3, "unauthorized account")
	ErrInvalidDenom                   = errors.Register(ModuleName, 4, "invalid denom")
	ErrInvalidCreator                 = errors.Register(ModuleName, 5, "invalid creator")
	ErrInvalidAuthorityMetadata       = errors.Register(ModuleName, 6, "invalid authority metadata")
	ErrInvalidGenesis                 = errors.Register(ModuleName, 7, "invalid genesis")
	ErrSubdenomTooLong                = errors.Register(ModuleName, 8, fmt.Sprintf("subdenom too long, max length is %d bytes", MaxSubdenomLength))
	ErrCreatorTooLong                 = errors.Register(ModuleName, 9, fmt.Sprintf("creator too long, max length is %d bytes", MaxCreatorLength))
	ErrDenomDoesNotExist              = errors.Register(ModuleName, 10, "denom does not exist")
	ErrEncodeTokenFactoryCreateDenom  = errors.Register(ModuleName, 11, "Error while encoding tokenfactory create denom msg in wasmd")
	ErrEncodeTokenFactoryMint         = errors.Register(ModuleName, 12, "Error while encoding tokenfactory mint denom msg in wasmd")
	ErrEncodeTokenFactoryBurn         = errors.Register(ModuleName, 13, "Error while encoding tokenfactory burn denom msg in wasmd")
	ErrEncodeTokenFactoryChangeAdmin  = errors.Register(ModuleName, 14, "Error while encoding tokenfactory change admin msg in wasmd")
	ErrParsingSeiTokenFactoryQuery    = errors.Register(ModuleName, 15, "Error parsing SeiTokenFactoryQuery")
	ErrAdminAlreadyExists             = errors.Register(ModuleName, 16, "attempting to create a new admin that already exists for the denom")
	ErrEncodeTokenFactorySetMetadata  = errors.Register(ModuleName, 17, "Error while encoding tokenfactory set metadata msg in wasmd")
	ErrEncodingDenomAuthorityMetadata = errors.Register(ModuleName, 18, "Error encoding denom authority metadata as JSON")
	ErrEncodingDenomsFromCreator      = errors.Register(ModuleName, 19, "Error encoding denoms from creator as JSON")
	ErrUnknownSeiTokenFactoryQuery    = errors.Register(ModuleName, 23, "Error unknown sei token factory query")
)