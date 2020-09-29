package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"github.com/vorteil/vorteil/pkg/provisioners"
	"github.com/vorteil/vorteil/pkg/provisioners/amazon"
	"github.com/vorteil/vorteil/pkg/provisioners/azure"
	"github.com/vorteil/vorteil/pkg/provisioners/google"
	"github.com/vorteil/vorteil/pkg/vdisk"
	"github.com/vorteil/vorteil/pkg/vio"
	"github.com/vorteil/vorteil/pkg/vpkg"
)

var provisionCmd = &cobra.Command{
	Use:  "provision BUILDABLE",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		// Load the provided provisioner file
		if _, err := os.Stat(provisionProvisionerFile); err != nil {
			setError(err, 1)

			return
		}

		b, err := ioutil.ReadFile(provisionProvisionerFile)
		if err != nil {
			setError(err, 2)

			return
		}

		data, err := provisioners.Decrypt(b, provisionPassPhrase)
		if err != nil {
			setError(err, 3)

			return
		}

		m := make(map[string]interface{})
		err = json.Unmarshal(data, &m)
		if err != nil {
			setError(err, 4)

			return
		}

		ptype, ok := m[provisioners.MapKey]
		if !ok {
			setError(err, 5)

			return
		}

		var prov provisioners.Provisioner

		switch ptype {
		case google.ProvisionerType:
			fmt.Println("Provisioning to Google Cloud Platform")
			p := &google.Provisioner{}
			err = p.Initialize(data)
			if err != nil {
				setError(err, 6)

				return
			}

			prov = p

		case amazon.ProvisionerType:
			fmt.Println("Provisioning to Amazon Web Services")
			p := &amazon.Provisioner{}
			err = p.Initialize(data)
			if err != nil {
				setError(err, 7)

				return
			}

			prov = p

		case azure.ProvisionerType:
			fmt.Println("Provisioning to Azure")
			p := &azure.Provisioner{}
			err = p.Initialize(data)
			if err != nil {
				setError(err, 8)

				return
			}

			prov = p
		}

		buildablePath := "."
		if len(args) >= 1 {
			buildablePath = args[0]
		}

		pkgBuilder, err := getPackageBuilder("BUILDABLE", buildablePath)
		if err != nil {
			setError(err, 9)

			return
		}

		err = modifyPackageBuilder(pkgBuilder)
		if err != nil {
			setError(err, 10)

			return
		}

		pkgReader, err := vpkg.ReaderFromBuilder(pkgBuilder)
		if err != nil {
			setError(err, 11)

			return
		}
		defer pkgReader.Close()

		pkgReader, err = vpkg.PeekVCFG(pkgReader)
		if err != nil {
			setError(err, 12)

			return
		}

		err = initKernels()
		if err != nil {
			setError(err, 13)

			return
		}

		f, err := ioutil.TempFile(os.TempDir(), "vorteil.disk")
		if err != nil {
			setError(err, 14)

			return
		}
		defer os.Remove(f.Name())
		defer f.Close()

		err = vdisk.Build(context.Background(), f, &vdisk.BuildArgs{
			PackageReader: pkgReader,
			Format:        prov.DiskFormat(),
			SizeAlign:     int64(prov.SizeAlign()),
			KernelOptions: vdisk.KernelOptions{
				Shell: flagShell,
			},
			Logger: log,
		})
		if err != nil {
			setError(err, 15)

			return
		}

		err = f.Close()
		if err != nil {
			setError(err, 16)

			return
		}

		err = pkgReader.Close()
		if err != nil {
			setError(err, 17)

			return
		}

		image, err := vio.LazyOpen(f.Name())
		if err != nil {
			setError(err, 18)

			return
		}

		if provisionName == "" {
			provisionName = strings.ReplaceAll(uuid.New().String(), "-", "")
		}

		ctx := context.TODO()
		err = prov.Provision(&provisioners.ProvisionArgs{
			Context:         ctx,
			Image:           image,
			Name:            provisionName,
			Description:     provisionDescription,
			Force:           provisionForce,
			Logger:          log,
			ReadyWhenUsable: provisionReadyWhenUsable,
		})
		if err != nil {
			setError(err, 19)

			return
		}

		fmt.Printf("Finished creating image.\n")
	},
}

var (
	provisionName            string
	provisionDescription     string
	provisionForce           bool
	provisionReadyWhenUsable bool
	provisionProvisionerFile string
	provisionPassPhrase      string
)

func init() {
	f := provisionCmd.Flags()
	f.StringVarP(&provisionName, "name", "n", "", "Name of the resulting image on the remote platform.")
	f.StringVarP(&provisionDescription, "description", "D", "", "Description for the resulting image, if supported by the platform.")
	f.BoolVarP(&provisionForce, "force", "f", false, "Force an overwrite if an existing image conflicts with the new.")
	f.BoolVarP(&provisionReadyWhenUsable, "ready-when-usable", "r", false, "Return successfully as soon as the operation is complete, regardless of whether or not the platform is still processing the image.")
	f.StringVarP(&provisionProvisionerFile, "provisioner", "p", "", "Path to file containing provisioner data.")
	f.StringVarP(&provisionPassPhrase, "passphrase", "s", "", "Passphrase used to decrypt encrypted provisioner data.")
}

var provisionersCmd = &cobra.Command{
	Use:     "provisioners",
	Short:   "Helper commands for working with Vorteil provisioners",
	Long:    ``,
	Example: ``,
}

var provisionersNewCmd = &cobra.Command{
	Use:   "new",
	Short: "Add a new provisioner.",
}

var (
	provisionersNewPassphrase string

	// Google Cloud Platform
	provisionersNewGoogleBucket  string
	provisionersNewGoogleKeyFile string

	// Amazon Web Services
	provisionersNewAmazonKey    string
	provisionersNewAmazonRegion string
	provisionersNewAmazonSecret string

	// Azure
	provisionersNewAzureContainer          string
	provisionersNewAzureKeyFile            string
	provisionersNewAzureLocation           string
	provisionersNewAzureResourceGroup      string
	provisionersNewAzureStorageAccountKey  string
	provisionersNewAzureStorageAccountName string
)

var provisionersNewAmazonEC2Cmd = &cobra.Command{
	Use:   "amazon-ec2 <OUTPUT_FILE>",
	Short: "Add a new AWS (Amazon Web Services) Provisioner.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {


		f, err := os.OpenFile(args[0], os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			setError(err, 1)

			return
		}
		defer f.Close()

		p, err := amazon.Create(&amazon.Config{
			Key:    provisionersNewAmazonKey,
			Secret: provisionersNewAmazonSecret,
			Region: provisionersNewAmazonRegion,
		})
		if err != nil {
			setError(err, 2)

			return
		}

		data, err := p.Marshal()
		if err != nil {
			setError(err, 3)

			return
		}

		out := provisioners.Encrypt(data, provisionersNewPassphrase)
		_, err = io.Copy(f, bytes.NewReader(out))
		if err != nil {
			setError(err, 4)

			return
		}

	},
}

func init() {
	f := provisionersNewAmazonEC2Cmd.Flags()
	f.StringVarP(&provisionersNewAmazonKey, "key", "k", "", "Access key ID")
	f.StringVarP(&provisionersNewAmazonSecret, "secret", "s", "", "Secret access key")
	f.StringVarP(&provisionersNewAmazonRegion, "region", "r", "ap-southeast-2", "AWS region")
	f.StringVarP(&provisionersNewPassphrase, "passphrase", "p", "", "Passphrase for encrypting exported provisioner data.")
}

var provisionersNewAzureCmd = &cobra.Command{
	Use:   "azure <OUTPUT_FILE>",
	Short: "Add a new Microsoft Azure Provisioner.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	
		f, err := os.OpenFile(args[0], os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			setError(err, 1)

			return
		}
		defer f.Close()

		path := provisionersNewAzureKeyFile
		_, err = os.Stat(path)
		if err != nil {
			setError(err, 2)

			return
		}

		b, err := ioutil.ReadFile(path)
		if err != nil {
			setError(err, 3)

			return
		}

		p, err := azure.Create(&azure.Config{
			Key:                base64.StdEncoding.EncodeToString(b),
			Container:          provisionersNewAzureContainer,
			Location:           provisionersNewAzureLocation,
			ResourceGroup:      provisionersNewAzureResourceGroup,
			StorageAccountKey:  provisionersNewAzureStorageAccountKey,
			StorageAccountName: provisionersNewAzureStorageAccountName,
		})
		if err != nil {
			setError(err, 4)

			return
		}

		data, err := p.Marshal()
		if err != nil {
			setError(err,5)

			return
		}

		out := provisioners.Encrypt(data, provisionersNewPassphrase)
		_, err = io.Copy(f, bytes.NewReader(out))
		if err != nil {
			setError(err, 6)

			return
		}

	},
}

func init() {
	f := provisionersNewAzureCmd.Flags()
	f.StringVarP(&provisionersNewAzureKeyFile, "key-file", "k", "", "Azure 'Service Principal' credentials file")
	f.StringVarP(&provisionersNewAzureContainer, "container", "c", "", "Azure container name")
	f.StringVarP(&provisionersNewAzureResourceGroup, "resource-group", "r", "", "Azure resource group name")
	f.StringVarP(&provisionersNewAzureLocation, "location", "l", "", "Azure location")
	f.StringVarP(&provisionersNewAzureStorageAccountKey, "storage-account-key", "s", "", "Azure storage account key")
	f.StringVarP(&provisionersNewAzureStorageAccountName, "storage-account-name", "n", "", "Azure storage account name")
}

var provisionersNewGoogleCmd = &cobra.Command{
	Use:   "google <OUTPUT_FILE>",
	Short: "Add a new Google Cloud (Compute Engine) Provisioner.",
	Args:  cobra.ExactArgs(1), // Single arg, points to output file
	Run: func(cmd *cobra.Command, args []string) {
	
		f, err := os.OpenFile(args[0], os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			setError(err, 1)

			return
		}
		defer f.Close()

		path := provisionersNewGoogleKeyFile
		_, err = os.Stat(path)
		if err != nil {
			setError(err, 2)

			return
		}

		b, err := ioutil.ReadFile(path)
		if err != nil {
			setError(err, 3)

			return
		}

		p, err := google.Create(&google.Config{
			Bucket: provisionersNewGoogleBucket,
			Key:    base64.StdEncoding.EncodeToString(b),
		})
		if err != nil {
			setError(err, 4)

			return
		}

		data, err := p.Marshal()
		if err != nil {
			setError(err, 5)

			return
		}

		out := provisioners.Encrypt(data, provisionersNewPassphrase)
		_, err = io.Copy(f, bytes.NewReader(out))
		if err != nil {
			setError(err, 6)

			return
		}
	},
}

func init() {
	f := provisionersNewGoogleCmd.Flags()
	f.StringVarP(&provisionersNewPassphrase, "passphrase", "p", "", "Passphrase for encrypting exported provisioner data.")
	f.StringVarP(&provisionersNewGoogleBucket, "bucket", "b", "", "Name of an existing Google Cloud Storage bucket, for which the provided service account credentials have adequate permissions for object creation/deletion.")
	f.StringVarP(&provisionersNewGoogleKeyFile, "credentials", "f", "", "Path of an existing JSON-formatted Google Cloud Platform service account credentials file.")
}
