package plan


import (
    "os"
	"os/user"
    "path"
    "strings"
    "encoding/hex"
    //"math/big"

    //"github.com/ethereum/go-ethereum/common/math"
)

/*****************************************************
** Utility & Conversion Helpers
**/

// GetCommunityID returns the CommunityID for the given buffer
func GetCommunityID(in []byte) CommunityID {

	var out CommunityID

	overhang := CommunityIDSz - len(in)
	if overhang < 0 {
		in = in[-overhang:]
		overhang = 0
	}

	copy(out[overhang:], in)
	return out
}

// GetKeyID returns the KeyID for the given buffer
func GetKeyID(in []byte) KeyID {

	var out KeyID

	overhang := KeyIDSz - len(in)
	if overhang < 0 {
		in = in[-overhang:]
		overhang = 0
	}

	copy(out[overhang:], in)
	return out
}

// GetChannelID returns the KeyID for the given buffer
func GetChannelID(in []byte) ChannelID {

	var out ChannelID

	overhang := ChannelIDSz - len(in)
	if overhang < 0 {
		in = in[-overhang:]
		overhang = 0
	}

	copy(out[overhang:], in)
	return out
}

// UseLocalDir ensures the dir pathname associated with PLAN exists and returns the final absolute pathname
// inSubDir can be any relative pathname
func UseLocalDir(inSubDir string) (string, error) {
	usr, err := user.Current()
	if err != nil {
        return "", Error(err, FileSysError, "failed to get current user dir")
	}

    pathname := usr.HomeDir
    pathname = path.Clean(path.Join(pathname, "_.plan"))

    if len(inSubDir) > 0 {
        pathname = path.Join(pathname, inSubDir)
    }

    err = os.MkdirAll(pathname, DefaultFileMode)
	if err != nil {
		return "", Error(err, FileSysError, "os.MkdirAll() failed")
	}

	return pathname, nil

}


var remapCharset = map[rune]rune{
    ' ':  '-',
    '.':  '-',
    '?':  '-',
    '\\': '+',
    '/':  '+',
    '&':  '+',
}
    

// MakeFSFriendly makes a given string safe to use for a file system.
// If inSuffix is given, the hex encoding of those bytes are appended after "-"
func MakeFSFriendly(inName string, inSuffix []byte) string {

    var b strings.Builder
    for _, r := range inName {
        if replace, ok := remapCharset[r]; ok {
            if replace != 0 {
                b.WriteRune(replace)
            }
        } else {
            b.WriteRune(r)
        }
    }
     
    name := b.String()
    if len(inSuffix) > 0 {
        name = name + "-" + hex.EncodeToString(inSuffix)
    }

    return name
}



