package net

import (
	"github.com/pkg/errors"
	"github.com/vitelabs/go-vite/common/types"
	"github.com/vitelabs/go-vite/ledger"
	"github.com/vitelabs/go-vite/p2p"
	"github.com/vitelabs/go-vite/vite/net/message"
	"time"
)

var errHandshakeTwice = errors.New("handshake should send only once")
var errMsgTimeout = errors.New("message response timeout")

var subledgerTimeout = 10 * time.Second
var accountBlocksTimeout = 30 * time.Second
var snapshotBlocksTimeout = time.Minute

// @section Cmd
const CmdSetName = "vite"

const CmdSet = 2

type cmd uint64

const (
	HandshakeCode cmd = iota
	StatusCode
	ForkCode // tell peer it has forked, use for respond GetSnapshotBlocksCode
	GetSubLedgerCode
	GetSnapshotBlocksCode // get snapshotblocks without content
	GetSnapshotBlocksContentCode
	GetFullSnapshotBlocksCode   // get snapshotblocks with content
	GetSnapshotBlocksByHashCode // a batch of hash
	GetSnapshotBlocksContentByHashCode
	GetFullSnapshotBlocksByHashCode
	GetAccountBlocksCode       // query single AccountChain
	GetMultiAccountBlocksCode  // query multi AccountChain
	GetAccountBlocksByHashCode // query accountBlocks by hashList
	GetFilesCode
	GetChunkCode
	SubLedgerCode
	FileListCode
	SnapshotBlocksCode
	SnapshotBlocksContentCode
	FullSnapshotBlocksCode
	AccountBlocksCode
	NewSnapshotBlockCode

	ExceptionCode = 127
)

var msgNames = [...]string{
	HandshakeCode:                      "HandShakeMsg",
	StatusCode:                         "StatusMsg",
	ForkCode:                           "ForkMsg",
	GetSubLedgerCode:                   "GetSubLedgerMsg",
	GetSnapshotBlocksCode:              "GetSnapshotBlocksMsg",
	GetSnapshotBlocksContentCode:       "GetSnapshotBlocksContentMsg",
	GetFullSnapshotBlocksCode:          "GetFullSnapshotBlocksMsg",
	GetSnapshotBlocksByHashCode:        "GetSnapshotBlocksByHashMsg",
	GetSnapshotBlocksContentByHashCode: "GetSnapshotBlocksContentByHashMsg",
	GetFullSnapshotBlocksByHashCode:    "GetFullSnapshotBlocksByHashMsg",
	GetAccountBlocksCode:               "GetAccountBlocksMsg",
	GetMultiAccountBlocksCode:          "GetMultiAccountBlocksMsg",
	GetAccountBlocksByHashCode:         "GetAccountBlocksByHashMsg",
	GetFilesCode:                       "GetFileMsg",
	GetChunkCode:                       "GetChunkMsg",
	SubLedgerCode:                      "SubLedgerMsg",
	FileListCode:                       "FileListMsg",
	SnapshotBlocksCode:                 "SnapshotBlocksMsg",
	SnapshotBlocksContentCode:          "SnapshotBlocksContentMsg",
	FullSnapshotBlocksCode:             "FullSnapshotBlocksMsg",
	AccountBlocksCode:                  "AccountBlocksMsg",
	NewSnapshotBlockCode:               "NewSnapshotBlockMsg",
}

func (t cmd) String() string {
	if t == ExceptionCode {
		return "ExceptionMsg"
	}
	return msgNames[t]
}

type MsgHandler interface {
	ID() string
	Cmds() []cmd
	Handle(msg *p2p.Msg, sender *Peer) error
}

// @section statusHandler
type _statusHandler func(msg *p2p.Msg, sender *Peer) error

func statusHandler(msg *p2p.Msg, sender *Peer) error {
	status := new(ledger.HashHeight)
	err := status.Deserialize(msg.Payload)
	if err != nil {
		return err
	}

	sender.SetHead(status.Hash, status.Height)
	return nil
}

func (s _statusHandler) ID() string {
	return "default status handler"
}

func (s _statusHandler) Cmds() []cmd {
	return []cmd{StatusCode}
}

func (s _statusHandler) Handle(msg *p2p.Msg, sender *Peer) error {
	return s(msg, sender)
}

// @section getSubLedgerHandler
type getSubLedgerHandler struct {
	chain Chain
}

func (s *getSubLedgerHandler) ID() string {
	return "default GetSubLedger Handler"
}

func (s *getSubLedgerHandler) Cmds() []cmd {
	return []cmd{GetSubLedgerCode}
}

func (s *getSubLedgerHandler) Handle(msg *p2p.Msg, sender *Peer) error {
	req := new(message.GetSnapshotBlocks)
	err := req.Deserialize(msg.Payload)
	if err != nil {
		return err
	}

	var files []*ledger.CompressedFileMeta
	var chunks [][2]uint64
	if req.From.Height != 0 {
		files, chunks = s.chain.GetSubLedgerByHeight(req.From.Height, req.Count, req.Forward)
	} else {
		files, chunks, err = s.chain.GetSubLedgerByHash(&req.From.Hash, req.Count, req.Forward)
	}

	if err != nil {
		return sender.Send(ExceptionCode, msg.Id, message.Missing)
	} else {
		return sender.Send(FileListCode, msg.Id, &message.FileList{
			Files:  files,
			Chunks: chunks,
			Nonce:  0,
		})
	}
}

type getSnapshotBlocksHandler struct {
	chain Chain
}

func (s *getSnapshotBlocksHandler) ID() string {
	return "default GetSnapshotBlocks Handler"
}

func (s *getSnapshotBlocksHandler) Cmds() []cmd {
	return []cmd{GetSnapshotBlocksCode}
}

func (s *getSnapshotBlocksHandler) Handle(msg *p2p.Msg, sender *Peer) error {
	req := new(message.GetSnapshotBlocks)
	err := req.Deserialize(msg.Payload)
	if err != nil {
		return err
	}

	var blocks []*ledger.SnapshotBlock
	if req.From.Height != 0 {
		blocks, err = s.chain.GetSnapshotBlocksByHeight(req.From.Height, req.Count, req.Forward, false)
	} else {
		blocks, err = s.chain.GetSnapshotBlocksByHash(&req.From.Hash, req.Count, req.Forward, false)
	}

	if err != nil {
		return sender.Send(ExceptionCode, msg.Id, message.Missing)
	} else {
		return sender.SendSnapshotBlocks(blocks, msg.Id)
	}
}

type getAccountBlocksHandler struct {
	chain Chain
}

func (a *getAccountBlocksHandler) ID() string {
	return "default GetAccountBlocks Handler"
}

func (a *getAccountBlocksHandler) Cmds() []cmd {
	return []cmd{GetAccountBlocksCode}
}

var NULL_ADDRESS = types.Address{}

func (a *getAccountBlocksHandler) Handle(msg *p2p.Msg, sender *Peer) error {
	as := new(message.GetAccountBlocks)
	err := as.Deserialize(msg.Payload)
	if err != nil {
		return err
	}

	// get correct address
	if as.Address == NULL_ADDRESS {
		block, err := a.chain.GetAccountBlockByHash(&as.From.Hash)
		if err != nil {
			return sender.Send(ExceptionCode, msg.Id, message.Missing)
		}
		as.Address = block.AccountAddress
	}

	var blocks []*ledger.AccountBlock
	if as.From.Height != 0 {
		blocks, err = a.chain.GetAccountBlocksByHeight(as.Address, as.From.Height, as.Count, as.Forward)
	} else {
		blocks, err = a.chain.GetAccountBlocksByHash(as.Address, &as.From.Hash, as.Count, as.Forward)
	}

	if err != nil {
		return sender.SendAccountBlocks(&message.AccountBlocks{
			Address: as.Address,
			Blocks:  blocks,
		}, msg.Id)
	} else {
		return sender.Send(ExceptionCode, msg.Id, message.Missing)
	}
}

// @section getChunkHandler
type getChunkHandler struct {
	chain Chain
}

func (c *getChunkHandler) ID() string {
	return "default GetChunk Handler"
}

func (c *getChunkHandler) Cmds() []cmd {
	return []cmd{GetChunkCode}
}

func (c *getChunkHandler) Handle(msg *p2p.Msg, sender *Peer) error {
	req := new(message.GetChunk)
	err := req.Deserialize(msg.Payload)
	if err != nil {
		return err
	}

	sblocks, mblocks, err := c.chain.GetConfirmSubLedger(req.Start, req.End)
	if err == nil {
		return sender.SendSubLedger(&message.SubLedger{
			SBlocks: sblocks,
			ABlocks: mblocks,
		}, msg.Id)
	} else {
		return sender.Send(ExceptionCode, msg.Id, message.Missing)
	}

	return nil
}