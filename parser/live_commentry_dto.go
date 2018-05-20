package parser

import (
	"encoding/base64"
	"encoding/xml"
	"log"
	"strings"
	"unicode"

	strip "github.com/grokify/html-strip-tags-go"
)

/*
<LIVECOMMENTRY>
    <COMMENTRIES>
        <COMMENTRY>
            <HEADER>
                <![CDATA[QnVsbGV0IHRyYWluIG1lZXQ6IEZhcm1lcnMgZGV0YWluZWQgYXQgdmVudWUsIHNvbWUgd2FsayBvdXQ=]]>
            </HEADER>
            <DESC>
                <![CDATA[U29tZSBmYXJtZXJzIHdlcmUgZGV0YWluZWQgYW5kIGEgZmV3IG90aGVycyB3YWxrZWQgb3V0IG9mIGEgY29uc3VsdGF0aW9uIG1lZXRpbmcgZm9yIHRoZSBBaG1lZGFiYWQtTXVtYmFpIGhpZ2ggc3BlZWQgcmFpbCBwcm9qZWN0IGluIFN1cmF0IHRvZGF5LCB3aXRoIGEgdG9wIGRpc3RyaWN0IG9mZmljaWFsIHN0YXRpbmcgdGhhdCBpdCB3ZW50IG9mZiBzdWNjZXNzZnVsbHkuPGJyPjxicj5Db2xsZWN0b3IgRGhhdmFsIFBhdGVsIHNhaWQgdGhhdCBzb21lIHBlcnNvbnMgd2VyZSBkZXRhaW5lZCBmcm9tIHRoZSB2ZW51ZSBhcyB0aGUgYWRtaW5pc3RyYXRpb24gZmVhcmVkIHRoYXQgdGhleSBtaWdodCByYWlzZSBzbG9nYW5zIGFuZCBjcmVhdGUgbGF3IGFuZCBvcmRlciBpc3N1ZXMuPGJyPjxicj4iU29tZSBwZXJzb25zIHdlcmUgZGV0YWluZWQgZHVlIHRvIGEgY29uY2VybiBvdmVyIGxhdyBhbmQgb3JkZXIuIEhvd2V2ZXIsIHRoZSBtZWV0aW5nIHdlbnQgb24gc3VjY2Vzc2Z1bGx5IGFzIHdlIHRyaWVkIHRvIGFkZHJlc3MgaXNzdWVzIHJhaXNlZCBieSB0aGUgZmFybWVycywiIFBhdGVsIHNhaWQuPGJyPjxicj5IZSBhZGRlZCB0aGF0IGFyb3VuZCA0MDAgZmFybWVycyBhdHRlbmRlZCB0aGUgcHVibGljIGhlYXJpbmcgd2hpY2ggd2FzIG9yZ2FuaXNlZCBhdCBjaXR5J3MgR2FuZGhpIFNtcnV0aSBCaGF2YW4uIAlEYXJzaGFuIE5haWssIGEgQ29uZ3Jlc3MgbWVtYmVyIG9mIHRoZSBTdXJhdCBkaXN0cmljdCBwYW5jaGF5YXQsIGNsYWltZWQgdGhhdCBoZSB3YXMgZGV0YWluZWQgYnkgdGhlIHBvbGljZSwgYW5kIHNhaWQgdGhhdCBzZXZlcmFsIGZhcm1lcnMgd2Fsa2VkIG91dCBhZnRlciBjb21pbmcgdG8ga25vdyBvZiB0aGUgZGV0ZW50aW9uLjxicj48YnI+IlRoZXJlIHdhcyBhIGxhcmdlIG51bWJlciBvZiBwb2xpY2UgcGVyc29ubmVsIGRlcGxveWVkIGF0IHRoZSBtZWV0aW5nIHdobyB3ZXJlIGZyaXNraW5nIGV2ZXJ5IGZhcm1lciBlbnRlcmluZyB0aGUgdmVudWUgaGFsbC4gSSB3ZW50IGluIHdpdGggYSBzZXQgb2YgZGVtYW5kcyBidXQgd2FzIHN0b3BwZWQgYnkgdGhlIHBvbGljZSBhbmQgZGV0YWluZWQsIiBzYWlkIE5haWsuPGJyPjxicj4iV2hlbiBmYXJtZXJzIHdobyBoYWQgZ29uZSBpbiBmb3IgdGhlIG1lZXRpbmcgY2FtZSB0byBrbm93IGFib3V0IHRoaXMsIHRoZXkgd2Fsa2VkIG91dCBhbmQgdGhlIG1lZXRpbmcgd2FzIG5vdCBoZWxkIGFnYWluLiBXZSBkZW1hbmQgaXRzIHJlc2NoZWR1bGluZywiIGhlIHNhaWQuPGJyPjxicj5BcyBwZXIgdGhlIHByb2Nlc3MsIHN1Y2ggY29uc3VsdGF0aW9uIG1lZXRpbmdzIGhhdmUgdG8gYmUgaGVsZCBiZWZvcmUgdGhlIGFjdHVhbCBwcm9jZXNzIG9mIGFjcXVpcmluZyBsYW5kIGNhbiBzdGFydC48YnI+PGJyPlRoZSBwcm9qZWN0IHdhcyBsYXVuY2hlZCBsYXN0IHllYXIgYnkgUHJpbWUgTWluaXN0ZXIgTmFyZW5kcmEgTW9kaSBhbmQgaGlzIEphcGFuZXNlIGNvdW50ZXJwYXJ0IFNoaW56byBBYmUgYW5kIGl0IGlzIGV4cGVjdGVkIHRvIGJlIHJlYWR5IGZvciBjb21taXNzaW9uaW5nIGJ5IDIwMjIuPGJyPjxicj5UaGUgaGlnaC1zcGVlZCB0cmFpbiB3aWxsIHJ1biBiZXR3ZWVuIEFobWVkYWJhZCBhbmQgTXVtYmFpIGFuZCB3aWxsIGhhbHQgaW4gMTIgc3RhdGlvbnMuJm5ic3A7IC0tIDxpPlBUSTwvaT48YnI+]]>
            </DESC>
            <IMGURL>
                <![CDATA[http://im.rediff.com/money/2018/apr/13bullet.jpg]]>
            </IMGURL>
            <IMGLINK>
                <![CDATA[]]>
            </IMGLINK>
            <CTIME>
                <![CDATA[1526319001]]>
            </CTIME>
            <UTIME>
                <![CDATA[1526319001]]>
            </UTIME>
            <DELETED>
                <![CDATA[0]]>
            </DELETED>
            <UNIQUEID>
                <![CDATA[da2881ebc769a5f036d5e5886bbb2bd4]]>
            </UNIQUEID>
            <TAGS>
                <![CDATA[DQoNCg==]]>
            </TAGS>
            <TOPICS>
                <![CDATA[I25hdGlvbmFsUG9saXRpY3M=]]>
            </TOPICS>
            <CHANNELS>
                <CHANNEL>
                    <![CDATA[news]]>
                </CHANNEL>
            </CHANNELS>
        </COMMENTRY>
*/
//AutoGenerated DTO to fill news data
type AutoGenerated struct {
	XMLName      xml.Name `xml:"LIVECOMMENTRY"`
	TITLE        string   `xml:"TITLE"`
	ENCODING     string   `xml:"ENCODING"`
	SHOWTIME     string   `xml:"SHOWTIME"`
	HEADERIMGURL string   `xml:"HEADER_IMGURL"`
	ZARABOLHASH  string   `xml:"ZARABOL_HASH"`
	COMMENTRIES  struct {
		COMMENTRY []struct {
			HEADER   string `xml:"HEADER"`
			DESC     string `xml:"DESC"`
			IMGURL   string `xml:"IMGURL"`
			IMGLINK  string `xml:"IMGLINK"`
			CTIME    string `xml:"CTIME"`
			UTIME    string `xml:"UTIME"`
			DELETED  string `xml:"DELETED"`
			UNIQUEID string `xml:"UNIQUEID"`
			TAGS     string `xml:"TAGS"`
			TOPICS   string `xml:"TOPICS"`
			CHANNELS struct {
				CHANNEL string `xml:"CHANNEL"`
			} `xml:"CHANNELS"`
		} `xml:"COMMENTRY"`
	} `xml:"COMMENTRIES"`
}

// SpaceMap use to trim string
func SpaceMap(str string) string {
	return strings.TrimFunc(str, func(r rune) bool {
		if unicode.IsSpace(r) {
			return true
		}
		return false
	})
}

// DecodeBase64 string
func DecodeBase64(str string) string {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Fatal("error:", err)
	}
	return strip.StripTags(string(data[:]))

}
