package cli

import (
	"fmt"

	"github.com/zhou-lincong/cloud_station/store"
	"github.com/zhou-lincong/cloud_station/store/aliyun"
	"github.com/zhou-lincong/cloud_station/store/aws"
	"github.com/zhou-lincong/cloud_station/store/txyun"

	"github.com/spf13/cobra"
)

var (
	ossProvier   string
	ossEndpoint  string
	accessKey    string
	accessSecret string
	bucketName   string
	uploadFile   string
)

const (
	default_ak = "xxx"
	default_sk = "xxx"
)

var UploadCmd = &cobra.Command{
	Use:     "upload",
	Long:    "upload 文件上传",
	Short:   "upload 文件上传",
	Example: "upload -f filename",
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			uploader store.Uploader
			err      error
		)
		switch ossProvier {
		case "aliyun":
			//不添加默认值
			// uploader, err = aliyun.NewAliOssStore(&aliyun.Options{
			// 	Endpoint:     ossEndpoint,
			// 	AccessKey:    accessKey,
			// 	AccessSecret: accessSecret,
			// })

			//添加ak 和 sk 默认值
			aliOpts := &aliyun.Options{
				Endpoint:     ossEndpoint,
				AccessKey:    accessKey,
				AccessSecret: accessSecret,
			}
			setAliDefault(aliOpts)
			uploader, err = aliyun.NewAliOssStore(aliOpts)
		case "txyun":
			uploader = txyun.NewTxOssStore()
		case "aws":
			uploader = aws.NewAwsOssStore()
		default:
			return fmt.Errorf("not support oss storage provider")
		}
		if err != nil {
			return err
		}

		// 使用Upload来上传文件
		return uploader.Upload(bucketName, uploadFile, uploadFile)
	},
}

func setAliDefault(opts *aliyun.Options) {
	if opts.AccessKey == "" {
		opts.AccessKey = default_ak
	}

	if opts.AccessSecret == "" {
		opts.AccessSecret = default_sk
	}
}

func init() {
	f := UploadCmd.PersistentFlags()
	f.StringVarP(&ossProvier, "provider", "p", "aliyun", "oss storage provier [aliyun/tx/aws]")
	f.StringVarP(&ossEndpoint, "endpoint", "e", "oss-cn-beijing.aliyuncs.com", "oss storage provier endpoint")
	f.StringVarP(&bucketName, "bucket_name", "b", "devcloud-station", "oss storage provier bucket name")
	f.StringVarP(&accessKey, "access_key", "k", "", "oss storage provier ak")
	f.StringVarP(&accessSecret, "access_secret", "s", "", "oss storage provier sk")
	f.StringVarP(&uploadFile, "upload_file", "f", "", "upload file name")
	//给root添加子命令
	RootCmd.AddCommand(UploadCmd)
}

//不添加ak sk默认值
// go run main.go upload -k ak -s sk -f go.mod

//不用默认aliyun
// go run main.go upload -p aws -f go.mod
