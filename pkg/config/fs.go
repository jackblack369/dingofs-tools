/*
 *  Copyright (c) 2022 NetEase Inc.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

/*
 * Project: DingoCli
 * Created Date: 2022-08-26
 * Author: chengyi (Cyber-SiKu)
 */
package config

import (
	"strings"
	"time"

	cmderror "github.com/dingodb/dingofs-tools/internal/error"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	// dingofs
	DINGOFS_MDSADDR              = "mdsaddr"
	VIPER_DINGOFS_MDSADDR        = "dingofs.mdsAddr"
	DINGOFS_MDSDUMMYADDR         = "mdsdummyaddr"
	VIPER_DINGOFS_MDSDUMMYADDR   = "dingofs.mdsDummyAddr"
	DINGOFS_ETCDADDR             = "etcdaddr"
	VIPER_DINGOFS_ETCDADDR       = "dingofs.etcdAddr"
	DINGOFS_METASERVERADDR       = "metaserveraddr"
	VIPER_DINGOFS_METASERVERADDR = "dingofs.metaserverAddr"
	DINGOFS_METASERVERID         = "metaserverid"
	VIPER_DINGOFS_METASERVERID   = "dingofs.metaserverId"
	DINGOFS_FSID                 = "fsid"
	VIPER_DINGOFS_FSID           = "dingofs.fsId"
	DINGOFS_FSNAME               = "fsname"
	VIPER_DINGOFS_FSNAME         = "dingofs.fsName"
	DINGOFS_MOUNTPOINT           = "mountpoint"
	VIPER_DINGOFS_MOUNTPOINT     = "dingofs.mountpoint"
	DINGOFS_PARTITIONID          = "partitionid"
	VIPER_DINGOFS_PARTITIONID    = "dingofs.partitionid"
	DINGOFS_NOCONFIRM            = "noconfirm"
	VIPER_DINGOFS_NOCONFIRM      = "dingofs.noconfirm"
	DINGOFS_USER                 = "user"
	VIPER_DINGOFS_USER           = "dingofs.user"
	DINGOFS_CAPACITY             = "capacity"
	VIPER_DINGOFS_CAPACITY       = "dingofs.capacity"
	DINGOFS_DEFAULT_CAPACITY     = "100 GiB"
	DINGOFS_BLOCKSIZE            = "blocksize"
	VIPER_DINGOFS_BLOCKSIZE      = "dingofs.blocksize"
	DINGOFS_DEFAULT_BLOCKSIZE    = "1 MiB"
	DINGOFS_SUMINDIR             = "sumindir"
	VIPER_DINGOFS_SUMINDIR       = "dingofs.sumindir"
	DINGOFS_DEFAULT_SUMINDIR     = false
	DINGOFS_FSTYPE               = "fstype"
	VIPER_DINGOFS_FSTYPE         = "dingofs.fstype"
	DINGOFS_COPYSETID            = "copysetid"
	VIPER_DINGOFS_COPYSETID      = "dingofs.copysetid"
	DINGOFS_POOLID               = "poolid"
	VIPER_DINGOFS_POOLID         = "dingofs.poolid"
	DINGOFS_DETAIL               = "detail"
	VIPER_DINGOFS_DETAIL         = "dingofs.detail"
	DINGOFS_DEFAULT_DETAIL       = false
	DINGOFS_INODEID              = "inodeid"
	VIPER_DINGOFS_INODEID        = "dingofs.inodeid"
	DINGOFS_CLUSTERMAP           = "clustermap"
	VIPER_DINGOFS_CLUSTERMAP     = "dingofs.clustermap"
	DINGOFS_DEFAULT_CLUSTERMAP   = "topo_example.json"
	DINGOFS_MARGIN               = "margin"
	VIPER_DINGOFS_MARGIN         = "dingofs.margin"
	DINGOFS_DEFAULT_MARGIN       = uint64(1000)
	DINGOFS_FILELIST             = "filelist"
	VIPER_DINGOFS_FILELIST       = "dingofs.filelist"
	DINGOFS_SERVERS              = "servers"
	VIPER_DINGOFS_SERVERS        = "dingofs.servers"
	DINGOFS_DEFAULT_SERVERS      = "127.0.0.1:7001,127.0.0.1:7002"
	DINGOFS_INTERVAL             = "interval"
	VIPER_DINGOFS_INTERVAL       = "dingofs.interval"
	DINGOFS_DEFAULT_INTERVAL     = 1 * time.Second
	DINGOFS_DAEMON               = "daemon"
	VIPER_DINGOFS_DAEMON         = "dingofs.daemon"
	DINGOFS_DEFAULT_DAEMON       = false
	DINGOFS_STORAGE              = "storage"
	VIPER_DINGOFS_STORAGE        = "dingofs.storage"
	DINGOFS_DEFAULT_STORAGE      = "disk"
	DINGOFS_STATS_SCHEMA         = "schema"
	VIPER_DINGOFS_STATS_SCHEMA   = "dingofs.schema"
	DINGOFS_STATS_DEFAULT_SCHEMA = "ufmbo"
	DINGOFS_STATS_COUNT          = "count"
	VIPER_DINGOFS_STATS_COUNT    = "dingofs.count"
	DINGOFS_STATS_DEFAULT_COUNT  = uint32(0)
	DINGOFS_QUOTA_PATH           = "path"
	VIPER_DINGOFS_QUOTA_PATH     = "dingofs.quota.path"
	DINGOFS_QUOTA_DEFAULT_PATH   = ""
	DINGOFS_QUOTA_CAPACITY       = "capacity"
	VIPER_DINGOFS_QUOTA_CAPACITY = "dingofs.quota.capacity"
	DINGOFS_QUOTA_DEF_CAPACITY   = uint64(0)
	DINGOFS_QUOTA_INODES         = "inodes"
	VIPER_DINGOFS_QUOTA_INODES   = "dingofs.quota.inodes"
	DINGOFS_QUOTA_DEFAULT_INODES = uint64(0)
	DINGOFS_QUOTA_REPAIR         = "repair"
	VIPER_DINGOFS_QUOTA_REPAIR   = "dingofs.quota.repair"
	DINGOFS_QUOTA_DEFAULT_REPAIR = false

	// S3
	DINGOFS_S3_AK                 = "s3.ak"
	VIPER_DINGOFS_S3_AK           = "dingofs.s3.ak"
	DINGOFS_DEFAULT_S3_AK         = "ak"
	DINGOFS_S3_SK                 = "s3.sk"
	VIPER_DINGOFS_S3_SK           = "dingofs.s3.sk"
	DINGOFS_DEFAULT_S3_SK         = "sk"
	DINGOFS_S3_ENDPOINT           = "s3.endpoint"
	VIPER_DINGOFS_S3_ENDPOINT     = "dingofs.s3.endpoint"
	DINGOFS_DEFAULT_ENDPOINT      = "http://localhost:9000"
	DINGOFS_S3_BUCKETNAME         = "s3.bucketname"
	VIPER_DINGOFS_S3_BUCKETNAME   = "dingofs.s3.bucketname"
	DINGOFS_DEFAULT_S3_BUCKETNAME = "bucketname"
	DINGOFS_S3_BLOCKSIZE          = "s3.blocksize"
	VIPER_DINGOFS_S3_BLOCKSIZE    = "dingofs.s3.blocksize"
	DINGOFS_DEFAULT_S3_BLOCKSIZE  = "4 MiB"
	DINGOFS_S3_CHUNKSIZE          = "s3.chunksize"
	VIPER_DINGOFS_S3CHUNKSIZE     = "dingofs.s3.chunksize"
	DINGOFS_DEFAULT_S3_CHUNKSIZE  = "64 MiB"
	// Volume
	DINGOFS_VOLUME_SIZE                   = "volume.size"
	VIPER_DINGOFS_VOLUME_SIZE             = "dingofs.volume.size"
	DINGOFS_DEFAULT_VOLUME_SIZE           = "1 MiB"
	DINGOFS_VOLUME_BLOCKGROUPSIZE         = "volume.blockgroupsize"
	VIPER_DINGOFS_VOLUME_BLOCKGROUPSIZE   = "dingofs.volume.blockgroupsize"
	DINGOFS_DEFAULT_VOLUME_BLOCKGROUPSIZE = "128 MiB"
	DINGOFS_VOLUME_BLOCKSIZE              = "volume.blocksize"
	VIPER_DINGOFS_VOLUME_BLOCKSIZE        = "dingofs.volume.blocksize"
	DINGOFS_DEFAULT_VOLUME_BLOCKSIZE      = "4 KiB"
	DINGOFS_VOLUME_NAME                   = "volume.name"
	VIPER_DINGOFS_VOLUME_NAME             = "dingofs.volume.name"
	DINGOFS_DEFAULT_VOLUME_NAME           = "volume"
	DINGOFS_VOLUME_USER                   = "volume.user"
	VIPER_DINGOFS_VOLUME_USER             = "dingofs.volume.user"
	DINGOFS_DEFAULT_VOLUME_USER           = "user"
	DINGOFS_VOLUME_PASSWORD               = "volume.password"
	VIPER_DINGOFS_VOLUME_PASSWORD         = "dingofs.volume.password"
	DINGOFS_DEFAULT_VOLUME_PASSWORD       = "password"
	DINGOFS_VOLUME_BITMAPLOCATION         = "volume.bitmaplocation"
	VIPER_DINGOFS_VOLUME_BITMAPLOCATION   = "dingofs.volume.bitmaplocation"
	DINGOFS_DEFAULT_VOLUME_BITMAPLOCATION = "AtStart"
	DINGOFS_VOLUME_SLICESIZE              = "volume.slicesize"
	VIPER_DINGOFS_VOLUME_SLICESIZE        = "dingofs.volume.slicesize"
	DINGOFS_DEFAULT_VOLUME_SLICESIZE      = "1 GiB"
	// gateway
	GATEWAY_LISTEN_ADDRESS          = "listen-address"
	VIPER_GATEWAY_LISTEN_ADDRESS    = "gateway.listen-address"
	GATEWAY_DEFAULT_LISTEN_ADDRESS  = ":19000"
	GATEWAY_CONSOLE_ADDRESS         = "console-address"
	VIPER_GATEWAY_CONSOLE_ADDRESS   = "gateway.console-address"
	GATEWAY_DEFAULT_CONSOLE_ADDRESS = ":19001"
)

var (
	FLAG2VIPER = map[string]string{
		RPCTIMEOUT:             VIPER_GLOBALE_RPCTIMEOUT,
		RPCRETRYTIMES:          VIPER_GLOBALE_RPCRETRYTIMES,
		VERBOSE:                VIPER_GLOBALE_VERBOSE,
		DINGOFS_MDSADDR:        VIPER_DINGOFS_MDSADDR,
		DINGOFS_MDSDUMMYADDR:   VIPER_DINGOFS_MDSDUMMYADDR,
		DINGOFS_ETCDADDR:       VIPER_DINGOFS_ETCDADDR,
		DINGOFS_METASERVERADDR: VIPER_DINGOFS_METASERVERADDR,
		DINGOFS_METASERVERID:   VIPER_DINGOFS_METASERVERID,
		DINGOFS_FSID:           VIPER_DINGOFS_FSID,
		DINGOFS_FSNAME:         VIPER_DINGOFS_FSNAME,
		DINGOFS_MOUNTPOINT:     VIPER_DINGOFS_MOUNTPOINT,
		DINGOFS_PARTITIONID:    VIPER_DINGOFS_PARTITIONID,
		DINGOFS_NOCONFIRM:      VIPER_DINGOFS_NOCONFIRM,
		DINGOFS_USER:           VIPER_DINGOFS_USER,
		DINGOFS_CAPACITY:       VIPER_DINGOFS_CAPACITY,
		DINGOFS_BLOCKSIZE:      VIPER_DINGOFS_BLOCKSIZE,
		DINGOFS_SUMINDIR:       VIPER_DINGOFS_SUMINDIR,
		DINGOFS_FSTYPE:         VIPER_DINGOFS_FSTYPE,
		DINGOFS_COPYSETID:      VIPER_DINGOFS_COPYSETID,
		DINGOFS_POOLID:         VIPER_DINGOFS_POOLID,
		DINGOFS_DETAIL:         VIPER_DINGOFS_DETAIL,
		DINGOFS_INODEID:        VIPER_DINGOFS_INODEID,
		DINGOFS_CLUSTERMAP:     VIPER_DINGOFS_CLUSTERMAP,
		DINGOFS_MARGIN:         VIPER_DINGOFS_MARGIN,
		DINGOFS_SERVERS:        VIPER_DINGOFS_SERVERS,
		DINGOFS_FILELIST:       VIPER_DINGOFS_FILELIST,
		DINGOFS_INTERVAL:       VIPER_DINGOFS_INTERVAL,
		DINGOFS_DAEMON:         VIPER_DINGOFS_DAEMON,
		DINGOFS_STORAGE:        VIPER_DINGOFS_STORAGE,
		DINGOFS_STATS_SCHEMA:   VIPER_DINGOFS_STATS_SCHEMA,
		DINGOFS_STATS_COUNT:    VIPER_DINGOFS_STATS_COUNT,
		DINGOFS_QUOTA_PATH:     VIPER_DINGOFS_QUOTA_PATH,
		DINGOFS_QUOTA_INODES:   VIPER_DINGOFS_QUOTA_INODES,
		DINGOFS_QUOTA_REPAIR:   VIPER_DINGOFS_QUOTA_REPAIR,

		// S3
		DINGOFS_S3_AK:         VIPER_DINGOFS_S3_AK,
		DINGOFS_S3_SK:         VIPER_DINGOFS_S3_SK,
		DINGOFS_S3_ENDPOINT:   VIPER_DINGOFS_S3_ENDPOINT,
		DINGOFS_S3_BUCKETNAME: VIPER_DINGOFS_S3_BUCKETNAME,
		DINGOFS_S3_BLOCKSIZE:  VIPER_DINGOFS_S3_BLOCKSIZE,
		DINGOFS_S3_CHUNKSIZE:  VIPER_DINGOFS_S3CHUNKSIZE,

		// Volume
		DINGOFS_VOLUME_SIZE:           VIPER_DINGOFS_VOLUME_SIZE,
		DINGOFS_VOLUME_BLOCKGROUPSIZE: VIPER_DINGOFS_VOLUME_BLOCKGROUPSIZE,
		DINGOFS_VOLUME_BLOCKSIZE:      VIPER_DINGOFS_VOLUME_BLOCKSIZE,
		DINGOFS_VOLUME_NAME:           VIPER_DINGOFS_VOLUME_NAME,
		DINGOFS_VOLUME_USER:           VIPER_DINGOFS_VOLUME_USER,
		DINGOFS_VOLUME_PASSWORD:       VIPER_DINGOFS_VOLUME_PASSWORD,
		DINGOFS_VOLUME_BITMAPLOCATION: VIPER_DINGOFS_VOLUME_BITMAPLOCATION,
		DINGOFS_VOLUME_SLICESIZE:      VIPER_DINGOFS_VOLUME_SLICESIZE,

		// gateway
		GATEWAY_LISTEN_ADDRESS:  VIPER_GATEWAY_LISTEN_ADDRESS,
		GATEWAY_CONSOLE_ADDRESS: VIPER_GATEWAY_CONSOLE_ADDRESS,
	}
	FLAG2DEFAULT = map[string]interface{}{
		RPCTIMEOUT:           DEFAULT_RPCTIMEOUT,
		RPCRETRYTIMES:        DEFAULT_RPCRETRYTIMES,
		VERBOSE:              DEFAULT_VERBOSE,
		DINGOFS_SUMINDIR:     DINGOFS_DEFAULT_SUMINDIR,
		DINGOFS_DETAIL:       DINGOFS_DEFAULT_DETAIL,
		DINGOFS_CLUSTERMAP:   DINGOFS_DEFAULT_CLUSTERMAP,
		DINGOFS_MARGIN:       DINGOFS_DEFAULT_MARGIN,
		DINGOFS_SERVERS:      DINGOFS_DEFAULT_SERVERS,
		DINGOFS_INTERVAL:     DINGOFS_DEFAULT_INTERVAL,
		DINGOFS_DAEMON:       DINGOFS_DEFAULT_DAEMON,
		DINGOFS_STORAGE:      DINGOFS_DEFAULT_STORAGE,
		DINGOFS_STATS_SCHEMA: DINGOFS_STATS_DEFAULT_SCHEMA,
		DINGOFS_STATS_COUNT:  DINGOFS_STATS_DEFAULT_COUNT,
		DINGOFS_QUOTA_PATH:   DINGOFS_QUOTA_DEFAULT_PATH,
		DINGOFS_QUOTA_INODES: DINGOFS_QUOTA_DEFAULT_INODES,
		DINGOFS_QUOTA_REPAIR: DINGOFS_QUOTA_DEFAULT_REPAIR,

		// S3
		DINGOFS_S3_AK:         DINGOFS_DEFAULT_S3_AK,
		DINGOFS_S3_SK:         DINGOFS_DEFAULT_S3_SK,
		DINGOFS_S3_ENDPOINT:   DINGOFS_DEFAULT_ENDPOINT,
		DINGOFS_S3_BUCKETNAME: DINGOFS_DEFAULT_S3_BUCKETNAME,
		DINGOFS_S3_BLOCKSIZE:  DINGOFS_DEFAULT_S3_BLOCKSIZE,
		DINGOFS_S3_CHUNKSIZE:  DINGOFS_DEFAULT_S3_CHUNKSIZE,

		// Volume
		DINGOFS_VOLUME_SIZE:           DINGOFS_DEFAULT_VOLUME_SIZE,
		DINGOFS_VOLUME_BLOCKGROUPSIZE: DINGOFS_DEFAULT_VOLUME_BLOCKGROUPSIZE,
		DINGOFS_VOLUME_BLOCKSIZE:      DINGOFS_DEFAULT_VOLUME_BLOCKSIZE,
		DINGOFS_VOLUME_NAME:           DINGOFS_DEFAULT_VOLUME_NAME,
		DINGOFS_VOLUME_USER:           DINGOFS_DEFAULT_VOLUME_USER,
		DINGOFS_VOLUME_PASSWORD:       DINGOFS_DEFAULT_VOLUME_PASSWORD,
		DINGOFS_VOLUME_BITMAPLOCATION: DINGOFS_DEFAULT_VOLUME_BITMAPLOCATION,
		DINGOFS_VOLUME_SLICESIZE:      DINGOFS_DEFAULT_VOLUME_SLICESIZE,

		// gateway
		GATEWAY_LISTEN_ADDRESS:  GATEWAY_DEFAULT_LISTEN_ADDRESS,
		GATEWAY_CONSOLE_ADDRESS: GATEWAY_DEFAULT_CONSOLE_ADDRESS,
	}
)

func AddStringOptionFlag(cmd *cobra.Command, name string, usage string) {
	defaultValue := FLAG2DEFAULT[name]
	if defaultValue == nil {
		defaultValue = ""
	}
	cmd.Flags().String(name, defaultValue.(string), usage)
	err := viper.BindPFlag(FLAG2VIPER[name], cmd.Flags().Lookup(name))
	if err != nil {
		cobra.CheckErr(err)
	}
}

func AddStringSliceOptionFlag(cmd *cobra.Command, name string, usage string) {
	defaultValue := FLAG2DEFAULT[name]
	if defaultValue == nil {
		defaultValue = []string{}
	}
	cmd.Flags().StringSlice(name, defaultValue.([]string), usage)
	err := viper.BindPFlag(FLAG2VIPER[name], cmd.Flags().Lookup(name))
	if err != nil {
		cobra.CheckErr(err)
	}
}

func AddStringRequiredFlag(cmd *cobra.Command, name string, usage string) {
	cmd.Flags().String(name, "", usage+color.Red.Sprint("[required]"))
	cmd.MarkFlagRequired(name)
	err := viper.BindPFlag(FLAG2VIPER[name], cmd.Flags().Lookup(name))
	if err != nil {
		cobra.CheckErr(err)
	}
}

func AddStringSliceRequiredFlag(cmd *cobra.Command, name string, usage string) {
	cmd.Flags().StringSlice(name, []string{}, usage+color.Red.Sprint("[required]"))
	cmd.MarkFlagRequired(name)
	err := viper.BindPFlag(FLAG2VIPER[name], cmd.Flags().Lookup(name))
	if err != nil {
		cobra.CheckErr(err)
	}
}

func AddUint64RequiredFlag(cmd *cobra.Command, name string, usage string) {
	cmd.Flags().Uint64(name, uint64(0), usage+color.Red.Sprint("[required]"))
	cmd.MarkFlagRequired(name)
	err := viper.BindPFlag(FLAG2VIPER[name], cmd.Flags().Lookup(name))
	if err != nil {
		cobra.CheckErr(err)
	}
}

func AddUint32RequiredFlag(cmd *cobra.Command, name string, usage string) {
	cmd.Flags().Uint32(name, uint32(0), usage+color.Red.Sprint("[required]"))
	cmd.MarkFlagRequired(name)
	err := viper.BindPFlag(FLAG2VIPER[name], cmd.Flags().Lookup(name))
	if err != nil {
		cobra.CheckErr(err)
	}
}

func AddBoolOptionFlag(cmd *cobra.Command, name string, usage string) {
	defaultValue := FLAG2DEFAULT[name]
	if defaultValue == nil {
		defaultValue = false
	}
	cmd.Flags().Bool(name, defaultValue.(bool), usage)
	err := viper.BindPFlag(FLAG2VIPER[name], cmd.Flags().Lookup(name))
	if err != nil {
		cobra.CheckErr(err)
	}
}

func AddBoolOptionPFlag(cmd *cobra.Command, name string, short string, usage string) {
	defaultValue := FLAG2DEFAULT[name]
	if defaultValue == nil {
		defaultValue = false
	}
	cmd.Flags().BoolP(name, short, defaultValue.(bool), usage)
	err := viper.BindPFlag(FLAG2VIPER[name], cmd.Flags().Lookup(name))
	if err != nil {
		cobra.CheckErr(err)
	}
}

func AddDurationOptionFlag(cmd *cobra.Command, name string, usage string) {
	defaultValue := FLAG2DEFAULT[name]
	if defaultValue == nil {
		defaultValue = 0
	}
	cmd.Flags().Duration(name, defaultValue.(time.Duration), usage)
	err := viper.BindPFlag(FLAG2VIPER[name], cmd.Flags().Lookup(name))
	if err != nil {
		cobra.CheckErr(err)
	}
}

func AddInt32OptionFlag(cmd *cobra.Command, name string, usage string) {
	defaultValue := FLAG2DEFAULT[name]
	if defaultValue == nil {
		defaultValue = int32(0)
	}
	cmd.Flags().Int32(name, defaultValue.(int32), usage)
	err := viper.BindPFlag(FLAG2VIPER[name], cmd.Flags().Lookup(name))
	if err != nil {
		cobra.CheckErr(err)
	}
}

func AddUint64OptionFlag(cmd *cobra.Command, name string, usage string) {
	defaultValue := FLAG2DEFAULT[name]
	if defaultValue == nil {
		defaultValue = 0
	}
	cmd.Flags().Uint64(name, defaultValue.(uint64), usage)
	err := viper.BindPFlag(FLAG2VIPER[name], cmd.Flags().Lookup(name))
	if err != nil {
		cobra.CheckErr(err)
	}
}

func AddUint32OptionFlag(cmd *cobra.Command, name string, usage string) {
	defaultValue := FLAG2DEFAULT[name]
	if defaultValue == nil {
		defaultValue = 0
	}
	cmd.Flags().Uint32(name, defaultValue.(uint32), usage)
	err := viper.BindPFlag(FLAG2VIPER[name], cmd.Flags().Lookup(name))
	if err != nil {
		cobra.CheckErr(err)
	}
}

// dingofs
// mds addr
func AddFsMdsAddrFlag(cmd *cobra.Command) {
	cmd.Flags().String(DINGOFS_MDSADDR, "", "mds address, should be like 127.0.0.1:6700,127.0.0.1:6701,127.0.0.1:6702")
	err := viper.BindPFlag(VIPER_DINGOFS_MDSADDR, cmd.Flags().Lookup(DINGOFS_MDSADDR))
	if err != nil {
		cobra.CheckErr(err)
	}
}

func GetAddrSlice(cmd *cobra.Command, addrType string) ([]string, *cmderror.CmdError) {
	var addrsStr string
	if cmd.Flag(addrType).Changed {
		addrsStr = cmd.Flag(addrType).Value.String()
	} else {
		addrsStr = viper.GetString(FLAG2VIPER[addrType])
	}
	addrslice := strings.Split(addrsStr, ",")
	for _, addr := range addrslice {
		if !IsValidAddr(addr) {
			err := cmderror.ErrGetAddr()
			err.Format(addrType, "["+addr+"]")
			return addrslice, err
		}
	}
	return addrslice, cmderror.ErrSuccess()
}

func GetFlagString(cmd *cobra.Command, flagName string) string {
	var value string
	if cmd.Flag(flagName).Changed {
		value = cmd.Flag(flagName).Value.String()
	} else {
		value = viper.GetString(FLAG2VIPER[flagName])
	}
	return value
}

func GetFlagBool(cmd *cobra.Command, flagName string) bool {
	var value bool
	flag := cmd.Flag(flagName)
	if flag == nil {
		return false
	}
	if flag.Changed {
		value, _ = cmd.Flags().GetBool(flagName)
	} else {
		value = viper.GetBool(FLAG2VIPER[flagName])
	}
	return value
}

func GetFlagUint64(cmd *cobra.Command, flagName string) uint64 {
	var value uint64
	if cmd.Flag(flagName).Changed {
		value, _ = cmd.Flags().GetUint64(flagName)
	} else {
		value = viper.GetUint64(FLAG2VIPER[flagName])
	}
	return value
}

func GetFlagUint32(cmd *cobra.Command, flagName string) uint32 {
	var value uint32
	if cmd.Flag(flagName).Changed {
		value, _ = cmd.Flags().GetUint32(flagName)
	} else {
		value = viper.GetUint32(FLAG2VIPER[flagName])
	}
	return value
}

func GetFlagStringSlice(cmd *cobra.Command, flagName string) []string {
	var value []string
	if cmd.Flag(flagName).Changed {
		value, _ = cmd.Flags().GetStringSlice(flagName)
	} else {
		value = viper.GetStringSlice(FLAG2VIPER[flagName])
	}
	return value
}

func GetFlagStringSliceDefaultAll(cmd *cobra.Command, flagName string) []string {
	var value []string
	if cmd.Flag(flagName).Changed {
		value, _ = cmd.Flags().GetStringSlice(flagName)
	} else {
		value = []string{"*"}
	}
	return value
}

func GetFlagDuration(cmd *cobra.Command, flagName string) time.Duration {
	var value time.Duration
	if cmd.Flag(flagName).Changed {
		value, _ = cmd.Flags().GetDuration(flagName)
	} else {
		value = viper.GetDuration(FLAG2VIPER[flagName])
	}
	return value
}

func GetFlagInt32(cmd *cobra.Command, flagName string) int32 {
	var value int32
	if cmd.Flag(flagName).Changed {
		value, _ = cmd.Flags().GetInt32(flagName)
	} else {
		value = viper.GetInt32(FLAG2VIPER[flagName])
	}
	return value
}

func GetFsMdsAddrSlice(cmd *cobra.Command) ([]string, *cmderror.CmdError) {
	return GetAddrSlice(cmd, DINGOFS_MDSADDR)
}

func GetRpcTimeout(cmd *cobra.Command) time.Duration {
	return GetFlagDuration(cmd, RPCTIMEOUT)
}

// mds dummy addr
func AddFsMdsDummyAddrFlag(cmd *cobra.Command) {
	cmd.Flags().String(DINGOFS_MDSDUMMYADDR, "", "mds dummy address, should be like 127.0.0.1:7700,127.0.0.1:7701,127.0.0.1:7702")
	err := viper.BindPFlag(VIPER_DINGOFS_MDSDUMMYADDR, cmd.Flags().Lookup(DINGOFS_MDSDUMMYADDR))
	if err != nil {
		cobra.CheckErr(err)
	}
}

func GetFsMdsDummyAddrSlice(cmd *cobra.Command) ([]string, *cmderror.CmdError) {
	return GetAddrSlice(cmd, DINGOFS_MDSDUMMYADDR)
}

// etcd addr
func AddEtcdAddrFlag(cmd *cobra.Command) {
	cmd.Flags().String(DINGOFS_ETCDADDR, "", "etcd address, should be like 127.0.0.1:8700,127.0.0.1:8701,127.0.0.1:8702")
	err := viper.BindPFlag(VIPER_DINGOFS_ETCDADDR, cmd.Flags().Lookup(DINGOFS_ETCDADDR))
	if err != nil {
		cobra.CheckErr(err)
	}
}

func GetFsEtcdAddrSlice(cmd *cobra.Command) ([]string, *cmderror.CmdError) {
	return GetAddrSlice(cmd, DINGOFS_ETCDADDR)
}

// metaserver addr
func AddMetaserverAddrOptionFlag(cmd *cobra.Command) {
	cmd.Flags().StringSlice(DINGOFS_METASERVERADDR, nil, "metaserver address, should be like 127.0.0.1:9700,127.0.0.1:9701,127.0.0.1:9702")
	err := viper.BindPFlag(VIPER_DINGOFS_METASERVERADDR, cmd.Flags().Lookup(DINGOFS_METASERVERADDR))
	if err != nil {
		cobra.CheckErr(err)
	}
}

// metaserver id
func AddMetaserverIdOptionFlag(cmd *cobra.Command) {
	cmd.Flags().StringSlice(DINGOFS_METASERVERID, nil, "metaserver id, should be like 1,2,3")
	err := viper.BindPFlag(VIPER_DINGOFS_METASERVERID, cmd.Flags().Lookup(DINGOFS_METASERVERID))
	if err != nil {
		cobra.CheckErr(err)
	}
}

// fs id [required]
func AddFsIdFlag(cmd *cobra.Command) {
	cmd.Flags().StringSlice(DINGOFS_FSID, nil, "fs Id, should be like 1,2,3 "+color.Red.Sprint("[required]"))
	cmd.MarkFlagRequired(DINGOFS_FSID)
	err := viper.BindPFlag(VIPER_DINGOFS_FSID, cmd.Flags().Lookup(DINGOFS_FSID))
	if err != nil {
		cobra.CheckErr(err)
	}
}

// fs id
func AddFsIdOptionDefaultAllFlag(cmd *cobra.Command) {
	cmd.Flags().StringSlice(DINGOFS_FSID, []string{"*"}, "fs Id, should be like 1,2,3 not set means all fs")
	err := viper.BindPFlag(VIPER_DINGOFS_FSID, cmd.Flags().Lookup(DINGOFS_FSID))
	if err != nil {
		cobra.CheckErr(err)
	}
}

// fs id
func AddFsIdSliceOptionFlag(cmd *cobra.Command) {
	cmd.Flags().StringSlice(DINGOFS_FSID, nil, "fs Id, should be like 1,2,3")
	err := viper.BindPFlag(VIPER_DINGOFS_FSID, cmd.Flags().Lookup(DINGOFS_FSID))
	if err != nil {
		cobra.CheckErr(err)
	}
}

// fs id
func AddFsIdUint32OptionFlag(cmd *cobra.Command) {
	cmd.Flags().Uint32(DINGOFS_FSID, 0, "fileSystem Id")
	err := viper.BindPFlag(VIPER_DINGOFS_FSID, cmd.Flags().Lookup(DINGOFS_FSID))
	if err != nil {
		cobra.CheckErr(err)
	}
}

// partition id [required]
func AddPartitionIdRequiredFlag(cmd *cobra.Command) {
	cmd.Flags().StringSlice(DINGOFS_PARTITIONID, nil, "partition Id, should be like 1,2,3"+color.Red.Sprint("[required]"))
	cmd.MarkFlagRequired(DINGOFS_PARTITIONID)
	err := viper.BindPFlag(FLAG2VIPER[DINGOFS_PARTITIONID], cmd.Flags().Lookup(DINGOFS_PARTITIONID))
	if err != nil {
		cobra.CheckErr(err)
	}
}

// fs name [required]
func AddFsNameRequiredFlag(cmd *cobra.Command) {
	cmd.Flags().String(DINGOFS_FSNAME, "", "fs name"+color.Red.Sprint("[required]"))
	cmd.MarkFlagRequired(DINGOFS_FSNAME)
	err := viper.BindPFlag(VIPER_DINGOFS_FSNAME, cmd.Flags().Lookup(DINGOFS_FSNAME))
	if err != nil {
		cobra.CheckErr(err)
	}
}

// fs name
func AddFsNameSliceOptionFlag(cmd *cobra.Command) {
	cmd.Flags().StringSlice(DINGOFS_FSNAME, nil, "fs name")
	err := viper.BindPFlag(VIPER_DINGOFS_FSNAME, cmd.Flags().Lookup(DINGOFS_FSNAME))
	if err != nil {
		cobra.CheckErr(err)
	}
}

// fs name
func AddFsNameStringOptionFlag(cmd *cobra.Command) {
	cmd.Flags().String(DINGOFS_FSNAME, "", "fsname=[fileSystem name]")
	err := viper.BindPFlag(VIPER_DINGOFS_FSNAME, cmd.Flags().Lookup(DINGOFS_FSNAME))
	if err != nil {
		cobra.CheckErr(err)
	}
}

// mountpoint [required]
func AddMountpointFlag(cmd *cobra.Command) {
	cmd.Flags().String(DINGOFS_MOUNTPOINT, "", "umount fs mountpoint"+color.Red.Sprint("[required]"))
	cmd.MarkFlagRequired(DINGOFS_MOUNTPOINT)
	err := viper.BindPFlag(VIPER_DINGOFS_MOUNTPOINT, cmd.Flags().Lookup("mountpoint"))
	if err != nil {
		cobra.CheckErr(err)
	}
}

// noconfirm
func AddNoConfirmOptionFlag(cmd *cobra.Command) {
	cmd.Flags().Bool(DINGOFS_NOCONFIRM, false, "do not confirm the command")
	err := viper.BindPFlag(VIPER_DINGOFS_NOCONFIRM, cmd.Flags().Lookup(DINGOFS_NOCONFIRM))
	if err != nil {
		cobra.CheckErr(err)
	}
}

/* option */
// User [option]
func AddUserOptionFlag(cmd *cobra.Command) {
	cmd.Flags().String(DINGOFS_USER, "anonymous", "user of request")
	err := viper.BindPFlag(VIPER_DINGOFS_USER, cmd.Flags().Lookup(DINGOFS_USER))
	if err != nil {
		cobra.CheckErr(err)
	}
}

// fs name [option]
func AddFsNameOptionFlag(cmd *cobra.Command) {
	cmd.Flags().String(DINGOFS_FSNAME, "", "fs name")
	err := viper.BindPFlag(VIPER_DINGOFS_FSNAME, cmd.Flags().Lookup(DINGOFS_FSNAME))
	if err != nil {
		cobra.CheckErr(err)
	}
}

// Capacity [option]
func AddCapacityOptionFlag(cmd *cobra.Command) {
	cmd.Flags().String(DINGOFS_CAPACITY, DINGOFS_DEFAULT_CAPACITY, "capacity of fs")
	err := viper.BindPFlag(VIPER_DINGOFS_CAPACITY, cmd.Flags().Lookup(DINGOFS_CAPACITY))
	if err != nil {
		cobra.CheckErr(err)
	}
}

// BlockSize [option]
func AddBlockSizeOptionFlag(cmd *cobra.Command) {
	cmd.Flags().String(DINGOFS_BLOCKSIZE, DINGOFS_DEFAULT_BLOCKSIZE, "block size")
	err := viper.BindPFlag(VIPER_DINGOFS_BLOCKSIZE, cmd.Flags().Lookup(DINGOFS_BLOCKSIZE))
	if err != nil {
		cobra.CheckErr(err)
	}
}

// SumInDir [option]
func AddSumInDIrOptionFlag(cmd *cobra.Command) {
	cmd.Flags().Bool(DINGOFS_SUMINDIR, false, "statistic info in xattr")
	err := viper.BindPFlag(VIPER_DINGOFS_SUMINDIR, cmd.Flags().Lookup(DINGOFS_SUMINDIR))
	if err != nil {
		cobra.CheckErr(err)
	}
}

// Fsype [option]
func AddFsTypeOptionFlag(cmd *cobra.Command) {
	cmd.Flags().String(DINGOFS_FSTYPE, "s3", "fs type, should be: s3, volume or hybrid")
	err := viper.BindPFlag(VIPER_DINGOFS_FSTYPE, cmd.Flags().Lookup(DINGOFS_FSTYPE))
	if err != nil {
		cobra.CheckErr(err)
	}
}

// S3.Ak [option]
func AddS3AkOptionFlag(cmd *cobra.Command) {
	AddStringOptionFlag(cmd, DINGOFS_S3_AK, "s3 ak")
}

// S3.Sk [option]
func AddS3SkOptionFlag(cmd *cobra.Command) {
	AddStringOptionFlag(cmd, DINGOFS_S3_SK, "s3 sk")
}

// S3.Endpoint [option]
func AddS3EndpointOptionFlag(cmd *cobra.Command) {
	AddStringOptionFlag(cmd, DINGOFS_S3_ENDPOINT, "s3 endpoint")
}

// S3.Buckname [option]
func AddS3BucknameOptionFlag(cmd *cobra.Command) {
	AddStringOptionFlag(cmd, DINGOFS_S3_BUCKETNAME, "s3 buckname")
}

// S3.Blocksize [option]
func AddS3BlocksizeOptionFlag(cmd *cobra.Command) {
	AddStringOptionFlag(cmd, DINGOFS_S3_BLOCKSIZE, "s3 blocksize")
}

// S3.Chunksize [option]
func AddS3ChunksizeOptionFlag(cmd *cobra.Command) {
	AddStringOptionFlag(cmd, DINGOFS_S3_CHUNKSIZE, "s3 chunksize")
}

// volume.size [option]
func AddVolumeSizeOptionFlag(cmd *cobra.Command) {
	AddStringOptionFlag(cmd, DINGOFS_VOLUME_SIZE, "volume size")
}

// volume.blockgroupsize [option]
func AddVolumeBlockgroupsizeOptionFlag(cmd *cobra.Command) {
	AddStringOptionFlag(cmd, DINGOFS_VOLUME_BLOCKGROUPSIZE, "volume block group size")
}

// volume.blocksize [option]
func AddVolumeBlocksizeOptionFlag(cmd *cobra.Command) {
	AddStringOptionFlag(cmd, DINGOFS_VOLUME_BLOCKSIZE, "volume blocksize")
}

// volume.name [option]
func AddVolumeNameOptionFlag(cmd *cobra.Command) {
	AddStringOptionFlag(cmd, DINGOFS_VOLUME_NAME, "volume name")
}

// volume.user [option]
func AddVolumeUserOptionFlag(cmd *cobra.Command) {
	AddStringOptionFlag(cmd, DINGOFS_VOLUME_USER, "volume user")
}

// volume.password [option]
func AddVolumePasswordOptionFlag(cmd *cobra.Command) {
	AddStringOptionFlag(cmd, DINGOFS_VOLUME_PASSWORD, "volume password")
}

// volume.bitmaplocation [option]
func AddVolumeBitmaplocationOptionFlag(cmd *cobra.Command) {
	AddStringOptionFlag(cmd, DINGOFS_VOLUME_BITMAPLOCATION, "volume space bitmap location, support |AtStart| and |AtEnd|")
}

// volume.slicesize [option]
func AddVolumeSlicesizeOptionFlag(cmd *cobra.Command) {
	AddStringOptionFlag(cmd, DINGOFS_VOLUME_SLICESIZE, "volume extents slice size")
}

// volume.slicesize [option]
func AddDetailOptionFlag(cmd *cobra.Command) {
	AddBoolOptionFlag(cmd, DINGOFS_DETAIL, "show more infomation")
}

// margin [option]
func AddMarginOptionFlag(cmd *cobra.Command) {
	AddUint64OptionFlag(cmd, DINGOFS_MARGIN, "the maximum gap between peers")
}

func GetMarginOptionFlag(cmd *cobra.Command) uint64 {
	return GetFlagUint64(cmd, DINGOFS_MARGIN)
}

// filelist [option]
func AddFileListOptionFlag(cmd *cobra.Command) {
	AddStringOptionFlag(cmd, DINGOFS_FILELIST,
		"filelist path, save the files(dir) to warmup absPath, and should be in dingofs")
}

func GetFileListOptionFlag(cmd *cobra.Command) string {
	return GetFlagString(cmd, DINGOFS_FILELIST)
}

// interval [option]
func AddIntervalOptionFlag(cmd *cobra.Command) {
	AddDurationOptionFlag(cmd, DINGOFS_INTERVAL, "interval time")
}

func GetIntervalFlag(cmd *cobra.Command) time.Duration {
	return GetFlagDuration(cmd, DINGOFS_INTERVAL)
}

// daemon [option]
func AddDaemonOptionFlag(cmd *cobra.Command) {
	AddBoolOptionFlag(cmd, DINGOFS_DAEMON, "run in daemon mode")
}

func AddDaemonOptionPFlag(cmd *cobra.Command) {
	AddBoolOptionPFlag(cmd, DINGOFS_DAEMON, "d", "run in daemon mode")
}

func GetDaemonFlag(cmd *cobra.Command) bool {
	return GetFlagBool(cmd, DINGOFS_DAEMON)
}

// storage [option]
func AddStorageOptionFlag(cmd *cobra.Command) {
	AddStringOptionFlag(cmd, DINGOFS_STORAGE, "warmup storage type, can be: disk/mem")
}

func GetStorageFlag(cmd *cobra.Command) string {
	return GetFlagString(cmd, DINGOFS_STORAGE)
}

// stats schema [option]
func AddFsSchemaOptionalFlag(cmd *cobra.Command) {
	AddStringOptionFlag(cmd, DINGOFS_STATS_SCHEMA, `schema string that controls the output sections (u: usage, f: fuse, m: metaserver,s: mds, b: blockcache, o: object)`)
}

func GetStatsSchemaFlagOptionFlag(cmd *cobra.Command) string {
	return GetFlagString(cmd, DINGOFS_STATS_SCHEMA)
}

// stats count [optional]
func AddFsCountOptionalFlag(cmd *cobra.Command) {
	AddUint32OptionFlag(cmd, DINGOFS_STATS_COUNT, `the max number of outout (Default is 0, no restriction)`)
}

func GetStatsCountFlagOptionFlag(cmd *cobra.Command) uint32 {
	return GetFlagUint32(cmd, DINGOFS_STATS_COUNT)
}

// capacity [option]
func AddFsCapacityOptionalFlag(cmd *cobra.Command) {
	cmd.Flags().Uint64(DINGOFS_QUOTA_CAPACITY, DINGOFS_QUOTA_DEF_CAPACITY, `hard quota for usage space in GiB`)
	err := viper.BindPFlag(VIPER_DINGOFS_QUOTA_CAPACITY, cmd.Flags().Lookup(DINGOFS_QUOTA_CAPACITY))
	if err != nil {
		cobra.CheckErr(err)
	}
}

// inodes [option]
func AddFsInodesOptionalFlag(cmd *cobra.Command) {
	AddUint64OptionFlag(cmd, DINGOFS_QUOTA_INODES, `hard quota for inodes (default: 0)`)
}

/* required */

// copysetid [required]
func AddCopysetidSliceRequiredFlag(cmd *cobra.Command) {
	AddStringSliceRequiredFlag(cmd, DINGOFS_COPYSETID, "copysetid")
}

// poolid [required]
func AddPoolidSliceRequiredFlag(cmd *cobra.Command) {
	AddStringSliceRequiredFlag(cmd, DINGOFS_POOLID, "poolid")
}

// inodeid [required]
func AddInodeIdRequiredFlag(cmd *cobra.Command) {
	AddUint64RequiredFlag(cmd, DINGOFS_INODEID, "inodeid")
}

// fsid [required]
func AddFsIdRequiredFlag(cmd *cobra.Command) {
	AddUint32RequiredFlag(cmd, DINGOFS_FSID, "fsid")
}

// fsid [optional]
func AddFsIdOptionalFlag(cmd *cobra.Command) {
	cmd.Flags().String(DINGOFS_FSID, "", "file system id")
	err := viper.BindPFlag(VIPER_DINGOFS_FSID, cmd.Flags().Lookup("fsid"))
	if err != nil {
		cobra.CheckErr(err)
	}

}

// gateway listen address [required]
func AddListenAddressRequiredFlag(cmd *cobra.Command) {
	AddStringRequiredFlag(cmd, GATEWAY_LISTEN_ADDRESS, "gateway listen address")
}

// gateway console address [optional]
func AddConsoleAddressOptionalFlag(cmd *cobra.Command) {
	AddStringOptionFlag(cmd, GATEWAY_CONSOLE_ADDRESS, "gateway console address")
}

// gateway mountpoint path [optional]
func AddGatewayMountpointOptionalFlag(cmd *cobra.Command) {
	cmd.Flags().String(DINGOFS_MOUNTPOINT, "", "fs gateway mountpoint")
	err := viper.BindPFlag(VIPER_DINGOFS_MOUNTPOINT, cmd.Flags().Lookup("mountpoint"))
	if err != nil {
		cobra.CheckErr(err)
	}
}

// cluserMap [required]
func AddClusterMapRequiredFlag(cmd *cobra.Command) {
	AddStringRequiredFlag(cmd, DINGOFS_CLUSTERMAP, "clusterMap")
}

// mountpoint [required]
func AddMountpointRequiredFlag(cmd *cobra.Command) {
	AddStringRequiredFlag(cmd, DINGOFS_MOUNTPOINT, "dingofs mountpoint path")
}

// servers [required]
func AddFsServersRequiredFlag(cmd *cobra.Command) {
	AddStringRequiredFlag(cmd, DINGOFS_SERVERS, "dingofs memcache servers")
}

// path [required]
func AddFsPathRequiredFlag(cmd *cobra.Command) {
	AddStringRequiredFlag(cmd, DINGOFS_QUOTA_PATH, "full path of the directory within the volume")
}