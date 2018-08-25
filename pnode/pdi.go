
package pnode



import (

   
    //"io"
    //"io/ioutil"
    //"strings"
    //"sync"
    //"time"
    //"sort"
    //"encoding/hex"
    //"encoding/json"
    //"encoding/base64"

    //"github.com/tidwall/redcon"

    "github.com/plan-tools/go-plan/pdi"
    "github.com/plan-tools/go-plan/ski"
    "github.com/plan-tools/go-plan/plan"

    // This inits in with sql, so no named import is needed
    _ "github.com/mattn/go-sqlite3"
    
    //"crypto/md5"
    //"hash"
    //"crypto/rand"

    //"github.com/stretchr/testify/assert"

    //"github.com/ethereum/go-ethereum/rlp"
    //"github.com/ethereum/go-ethereum/common/hexutil"

    //"github.com/plan-tools/go-plan/pservice"

    //"golang.org/x/net/context"

)


// entryWorkspace is a workspace used to pass around
type entryWorkspace struct {
    CR              *CommunityRepo    

    timeStart       plan.Time

    skiVersion      []byte                  
    entry           *pdi.EntryCrypt

    entryHash       []byte
    entryHeader     pdi.EntryHeader
    entryBody       pdi.Body
    
    authorInfo      IdentityInfo


    skiSession      ski.Session
    skiProvider     ski.Provider

    accessChannel   *ChannelStore
    targetChannel   *ChannelStore

}





//            ski.ArgKeyVersion:        ws.skiVersion,


// internal: unpackHeader
//   decrypts and deserializes a pdi header
func (ws *entryWorkspace) unpackHeader(
    inOnCompletion func(*plan.Perror),
    ) {

    ws.timeStart = plan.Now()

    switch ws.entry.GetEntryVersion() {
        case pdi.EntryVersion1: 
            //ws.skiVersion = ski.CryptSKIVersion
        default:
            inOnCompletion(plan.Error(nil, plan.BadPDIEntryFormat, "bad or unsupported PDI entry format"))
            return
    }

    // The entry header is encrypted using one of the community keys.
    ws.skiSession.DispatchOp( 

        &ski.OpArgs {
            OpName: ski.OpDecryptFromCommunity,
            CryptoKeyID: ws.entry.GetCommunityKeyID(),
            Msg: ws.entry.HeaderCrypt,
        }, 

        func(inErr *plan.Perror, inResults *pdi.Body) {
            if inErr != nil {
                inOnCompletion(plan.Error(inErr, plan.FailedToProcessPDIHeader, "failed to decrypt PDI header"))
                return
            }

            err := ws.entryHeader.Unmarshal(inResults.Parts[0].Content)
            if err != nil {
                inOnCompletion(plan.Error(err, plan.FailedToProcessPDIHeader, "failed to unmarshal PDI header"))
                return
            }

            // At this point, ws.entryHeader is ready for use
            inOnCompletion(nil)
        },
    )

}



// internal: validateEntry
//   before we write to the pnode, we need to verify the author is
//   valid and that they had permissions to do the things the entry wants to do. 
//   note that because permissions are immutable at a point in time, it doesn't matter
//   when we check permissions if they're changed later -- they'll
//   always be the same for an entry at a specific point in time.
func (ws *entryWorkspace) validateEntry() *plan.Perror {

    if ws.entry.TimeCreated < ws.CR.Info.TimeCreated.UnixSecs {
        return plan.Error(nil, plan.BadTimestamp, "PDI entry has timestamp earlier than community creation timestamp")
    }
    if ws.timeStart.UnixSecs - ws.entry.TimeCreated + ws.CR.Info.MaxPeerClockDelta < 0 {
        return plan.Error(nil, plan.BadTimestamp, "PDI entry has timestamp too far in the future")
    }

    perr := ws.CR.LookupIdentity(ws.entryHeader.AuthorMemberId, ws.entryHeader.AuthorMemberRev, &ws.authorInfo)
    if perr != nil {
        return perr
    }

    ws.entryHash = ws.entry.ComputeHash()

    perr = ws.skiProvider.VerifySignature( 
        ws.entry.Sig,
        ws.entryHash,
        ws.authorInfo.SigningPubKey,
    )

    if perr != nil {
        return plan.Error(perr, plan.FailedToProcessPDIHeader, "PDI entry signature verification failed")
    }


    err := ws.prepChannelAccess()

    // At this point, the PDI entry's signature has been verified
    return err

}



func (ws *entryWorkspace) prepChannelAccess() *plan.Perror {

    var err *plan.Perror

      // Fetch the data structure container for the cited access channel
    ws.targetChannel, err = ws.CR.FetchChannelStore(
        ws.entryHeader.ChannelId, 
        ws.entryHeader.ChannelRev,
        PostingToChannel | LoadIfNeeded)
    
    if ws.targetChannel == nil {
        return plan.Errorf(err, plan.AccessChannelNotFound, "channel 0x%x not found", ws.entryHeader.ChannelId )
    }

    if ws.targetChannel.ACStore == nil {
        return plan.Errorf(nil, plan.NotAnAccessChannel, "invalid channel 0x%x", ws.entryHeader.ChannelId )
    }

    // TODO: do all of ACStore checking!

    return nil
}




func (ws *entryWorkspace) storeEntry(
    inOnCompletion func(*plan.Perror),
    ) {


}






func (ws *entryWorkspace) processAndStoreEntry( 
    inOnCompletion func(*plan.Perror),
    ) {

    ws.unpackHeader( func(inErr *plan.Perror) {
        if inErr != nil {
            inOnCompletion(inErr)
        }

        perr := ws.validateEntry()
        if perr != nil {
            inOnCompletion(perr)
        }

        perr = ws.prepChannelAccess()
        if perr != nil {
            inOnCompletion(perr)
        }

        ws.storeEntry(inOnCompletion)
    })
   

}   



/*
    entry := new( plan.PDIEntry )
    entry.PDIEntryCrypt = ioEntry

    var err error

    entry.HeaderBuf, err = CR.decryptCommunityData( ioEntry.CommunityKeyID, ioEntry.HeaderCrypt )
    if err != nil {
        return err
    }

    // De-serialize inEntry.HeaderBuf into inEntry.Header
    entry.Header = new( plan.PDIEntryHeader )
    err = rlp.DecodeBytes( entry.HeaderBuf, entry.Header )
    if err != nil {
        return err
    }

    // Used in various places
    ioEntry.Hash = new( plan.PDIEntryHash )
    ioEntry.ComputeHash( ioEntry.Hash )

    // Now that we've decrypted and de-serialized the header, we can verify the entry's signature
    err = CR.VerifySig( ioEntry )
    if err != nil {
        return err
    }

    // Fetch (or load and fetch) the ChannelStore associated with the given channel
    CS, err := CR.GetChannelStore( &entry.Header.ChannelID, PostingToChannel | LoadIfNeeded )

    verb := entry.Header.Verb
    switch ( verb ) {

        case plan.PDIEntryVerbPostEntry:

            // First, we must validate the access channel cited by the header used by the author to back permissions for posting this entry.
            // This checks that the author didn't use an invalid or expired access channel to post this entry.  Once we validate this, 
            //    we can trust and use that access channel to check permissions further.
            err = CS.ValidateCitedAccessChannel( entry.Header );

            err = CR.VerifyWriteAccess( CS, entry.Header )
            if err != nil {
                return err
            }

            err := CS.WriteEntryToStorage( entry )
            if err != nil {
                return err
            }

        case plan.PDIEntryVerbChannelAdmin:

            // In general, if the channel already exists, it's an error.  Howeever we need to check if this entry 
            if CS != nil {
                //err = plan.Error( )
            }

        default:
            plan.Assert( false, "Unhandled verb" )

    } 



    return err

}





// VerifyAccess checks that the given PDI Entry has the proper permissions to do what it says it wants to do and that
//    the AccessChannelID cited is in fact a valid access channel to cite (given the timestamp of the entry, etc)
func (CR *CommunityRepo) VerifyWriteAccess( CS *ChannelStore, inHeader *plan.PDIEntryHeader ) error {

    // Get/Load/Create the data structure container for the cited access channel
    AC, _ := CR.GetChannelStore( &inHeader.AccessChannelID, IsAccessChannel | ReadingFromChannel | LoadIfNeeded )
    if AC == nil {
        return plan.Errorf( AccessChannelNotFound, "cited access channel 0x%x not found", inHeader.AccessChannelID )
    }
    if AC.ACStore == nil {
        return plan.Errorf( NotAnAccessChannel, "cited access channel 0x%x not actually an access channel", inHeader.AccessChannelID )
    }

    // Entries posted to a channel cite (and use) the latest/current AccessChannelID associated with the channel.
    // ...but pnodes must check this!


    {
        access := AC.ACStore.AccessByAuthor[inHeader.Author]
        if ( access & AuthorHasWriteAccess ) == 0 {
            return plan.Error( AuthorLacksWritePermission, "Author does not have write access to channel" )
        }
    }

    return nil

}





*/
