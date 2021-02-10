package commands

import (
	"context"
	"encoding/binary"
	"fmt"
	"github.com/ledgerwatch/turbo-geth/core/types"
	"github.com/ledgerwatch/turbo-geth/rlp"
	"math/big"
	"time"

	"github.com/spf13/cobra"

	"github.com/ledgerwatch/turbo-geth/common"
	"github.com/ledgerwatch/turbo-geth/common/dbutils"
	"github.com/ledgerwatch/turbo-geth/core/rawdb"
	"github.com/ledgerwatch/turbo-geth/ethdb"
	"github.com/ledgerwatch/turbo-geth/log"
	"github.com/ledgerwatch/turbo-geth/turbo/snapshotsync"
)

func init() {
	withChaindata(generateBodiesSnapshotCmd)
	withSnapshotFile(generateBodiesSnapshotCmd)
	withSnapshotData(generateBodiesSnapshotCmd)
	withBlock(generateBodiesSnapshotCmd)
	rootCmd.AddCommand(generateBodiesSnapshotCmd)

}

var generateBodiesSnapshotCmd = &cobra.Command{
	Use:     "bodies",
	Short:   "Generate bodies snapshot",
	Example: "go run cmd/snapshots/generator/main.go bodies --block 11000000 --chaindata /media/b00ris/nvme/snapshotsync/tg/chaindata/ --snapshotDir /media/b00ris/nvme/snapshotsync/tg/snapshots/ --snapshotMode \"hb\" --snapshot /media/b00ris/nvme/snapshots/bodies_test",
	RunE: func(cmd *cobra.Command, args []string) error {
		return BodySnapshot(cmd.Context(), chaindata, snapshotFile, block, snapshotDir, snapshotMode)
	},
}

func BodySnapshot(ctx context.Context, dbPath, snapshotPath string, toBlock uint64, snapshotDir string, snapshotMode string) error {
	kv := ethdb.NewLMDB().Path(dbPath).MustOpen()
	db := ethdb.NewObjectDatabase(kv)
	var (
		hash common.Hash
		err error
	)
	t := time.Now()

	if snapshotDir != "" {
		var mode snapshotsync.SnapshotMode
		mode, err = snapshotsync.SnapshotModeFromString(snapshotMode)
		if err != nil {
			return err
		}

		kv, err = snapshotsync.WrapBySnapshotsFromDir(kv, snapshotDir, mode)
		if err != nil {
			return err
		}
	}

	snKV := ethdb.NewMDBX().WithBucketsConfig(func(defaultBuckets dbutils.BucketsCfg) dbutils.BucketsCfg {
		return dbutils.BucketsCfg{
			dbutils.BlockBodyPrefix:          dbutils.BucketConfigItem{},
			dbutils.EthTx:          dbutils.BucketsConfigs[dbutils.EthTx],
			dbutils.BodiesSnapshotInfoBucket: dbutils.BucketConfigItem{},
			//dbutils.Sequence: dbutils.BucketConfigItem{},
		}
	}).Path(snapshotPath).MustOpen()

	snDB := ethdb.NewObjectDatabase(snKV)

	chunkFile := 100000
	tuples := make(ethdb.MultiPutTuples, 0, chunkFile*3+100)


	for i := uint64(1); i <= toBlock; i++ {
		if common.IsCanceled(ctx) {
			return common.ErrStopped
		}

		hash, err = rawdb.ReadCanonicalHash(db, i)
		if err != nil {
			return fmt.Errorf("getting canonical hash for block %d: %v", i, err)
		}
		bodyRlp:=rawdb.ReadStorageBodyRLP(db, hash, i)
		tuples = append(tuples, []byte(dbutils.BlockBodyPrefix), dbutils.BlockBodyKey(i, hash), bodyRlp)
		if len(tuples) >= chunkFile {
			log.Info("Committed", "block", i)
			if _, err = snDB.MultiPut(tuples...); err != nil {
				log.Crit("Multiput error", "err", err)
				return err
			}
			tuples = tuples[:0]
		}
	}

	if len(tuples) > 0 {
		if _, err = snDB.MultiPut(tuples...); err != nil {
			log.Crit("Multiput error", "err", err)
			return err
		}
	}
	tuples = tuples[:0]


	log.Info("Bodies copied", "t", time.Since(t))

	t2:=time.Now()
	hash, err = rawdb.ReadCanonicalHash(db, toBlock+1)
	if err != nil {
		return fmt.Errorf("getting canonical hash for block %s: %v", hash, err)
	}
	bodyRlp:=rawdb.ReadStorageBodyRLP(db, hash, toBlock+1)
	bodyForStorage :=new(types.BodyForStorage)
	err = rlp.DecodeBytes(bodyRlp, bodyForStorage)
	if err != nil {
		log.Error("Invalid block body RLP", "hash", hash, "err", err)
		return err
	}

	//r,err:=snDB.Sequence(dbutils.EthTx, bodyForStorage.BaseTxId)
	//if err!=nil {
	//	return fmt.Errorf("seq %w",err)
	//}
	//fmt.Println(r, err)
	err = db.Walk(dbutils.EthTx, []byte{},0, func(k, v []byte) (bool, error) {
		if common.IsCanceled(ctx) {
			return false, common.ErrStopped
		}

		if binary.BigEndian.Uint64(k) >= bodyForStorage.BaseTxId {
			return false, nil
		}

		tuples = append(tuples, []byte(dbutils.EthTx), common.CopyBytes(k), common.CopyBytes(v))
		if len(tuples) >= chunkFile {
			log.Info("Committed", "tx", binary.BigEndian.Uint64(k))
			if _, err = snDB.MultiPut(tuples...); err != nil {
				log.Crit("Multiput error", "err", err)
				return false, err
			}
			tuples = tuples[:0]
		}
		return true, nil
	})

	if len(tuples) > 0 {
		if _, err = snDB.MultiPut(tuples...); err != nil {
			log.Crit("Multiput error", "err", err)
			return err
		}
	}
	log.Info("Transactions copied", "t", time.Since(t2))


	err = snDB.Put(dbutils.BodiesSnapshotInfoBucket, []byte(dbutils.SnapshotBodyHeadNumber), big.NewInt(0).SetUint64(toBlock).Bytes())
	if err != nil {
		log.Crit("SnapshotBodyHeadNumber error", "err", err)
		return err
	}
	err = snDB.Put(dbutils.BodiesSnapshotInfoBucket, []byte(dbutils.SnapshotBodyHeadHash), hash.Bytes())
	if err != nil {
		log.Crit("SnapshotBodyHeadHash error", "err", err)
		return err
	}
	snDB.Close()
	//err = os.Remove(snapshotPath + "/lock.mdb")
	//if err != nil {
	//	log.Warn("Remove lock", "err", err)
	//	return err
	//}

	log.Info("Finished", "duration", time.Since(t))
	return nil
}
