package lnwire

import (
	"fmt"
	"io"
)

//Multiple Clearing Requests are possible by putting this inside an array of
//clearing requests
type HTLCSettleAccept struct {
	//We can use a different data type for this if necessary...
	ChannelID uint64

	//ID of this request
	StagingID uint64
}

func (c *HTLCSettleAccept) Decode(r io.Reader, pver uint32) error {
	//ChannelID(8)
	//StagingID(8)
	//Expiry(4)
	//Amount(4)
	//NextHop(20)
	//ContractType(1)
	//RedemptionHashes (numOfHashes * 20 + numOfHashes)
	//Blob(2+blobsize)
	err := readElements(r,
		&c.ChannelID,
		&c.StagingID,
	)
	if err != nil {
		return err
	}

	return nil
}

//Creates a new HTLCSettleAccept
func NewHTLCSettleAccept() *HTLCSettleAccept {
	return &HTLCSettleAccept{}
}

//Serializes the item from the HTLCSettleAccept struct
//Writes the data to w
func (c *HTLCSettleAccept) Encode(w io.Writer, pver uint32) error {
	err := writeElements(w,
		c.ChannelID,
		c.StagingID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (c *HTLCSettleAccept) Command() uint32 {
	return CmdHTLCSettleAccept
}

func (c *HTLCSettleAccept) MaxPayloadLength(uint32) uint32 {
	//16
	return 16
}

//Makes sure the struct data is valid (e.g. no negatives or invalid pkscripts)
func (c *HTLCSettleAccept) Validate() error {
	//We're good!
	return nil
}

func (c *HTLCSettleAccept) String() string {
	return fmt.Sprintf("\n--- Begin HTLCSettleAccept ---\n") +
		fmt.Sprintf("ChannelID:\t%d\n", c.ChannelID) +
		fmt.Sprintf("StagingID:\t%d\n", c.StagingID) +
		fmt.Sprintf("--- End HTLCSettleAccept ---\n")
}
