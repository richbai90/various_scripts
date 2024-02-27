/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"image/color"
	"ml_preprocessor/helpers"
	"os"
	"path"
	"path/filepath"

	"github.com/sunshineplan/imgconv"

	"github.com/spf13/cobra"
)

var recursive bool
var recursiveDepth int
var output string
var isDir bool
var dir string
var file string
var width int
var height int

var supportedExtensions = []string{".png", ".gif", ".bmp", ".tiff", ".tif", ".jpg", ".jpeg"}

// processCmd represents the process command
var processCmd = &cobra.Command{
	Use:   "process",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MaximumNArgs(0),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		cwd, _ := os.Getwd()
		// Check if file or dir is set
		if cmd.Flags().Changed("file") && cmd.Flags().Changed("dir") {
			return fmt.Errorf("file and dir cannot be set at the same time")
		}
		if cmd.Flags().Changed("recursive") && !cmd.Flags().Changed("dir") {
			return fmt.Errorf("recursive can only be used with dir")
		}
		if cmd.Flags().Changed("recursive-depth") && !cmd.Flags().Changed("dir") {
			return fmt.Errorf("recursive-depth can only be used with dir")
		}
		if cmd.Flags().Changed("width") {
			width, _ = cmd.Flags().GetInt("width")
		}
		if cmd.Flags().Changed("height") {
			height, _ = cmd.Flags().GetInt("height")
		}
		// Check if file or dir is set
		if !cmd.Flags().Changed("file") && !cmd.Flags().Changed("dir") {
			return fmt.Errorf("file or dir must be set")
		}
		// set global variables
		recursiveDepth, _ = cmd.Flags().GetInt("recursive-depth")
		recursive, _ = cmd.Flags().GetBool("recursive")
		output, _ = cmd.Flags().GetString("output")
		dir, _ = cmd.Flags().GetString("dir")
		isDir = cmd.Flags().Changed("dir")
		file, _ = cmd.Flags().GetString("file")
		// check if output is default
		if output == "Current Working Directory" {
			if isDir {
				output = cwd
			} else {
				file, _ := cmd.Flags().GetString("file")
				output = path.Join(cwd, fmt.Sprintf("ouput_%s", file))
			}
		}
		// check if output exists
		var outdir string
		fmt.Printf("Checking if output directory %s exists\n", output)
		if !isDir {
			outdir = path.Dir(output)
		} else {
			outdir = output
		}
		if _, err := os.Stat(outdir); os.IsNotExist(err) {
			fmt.Printf("Output directory %s does not exist, creating it now", outdir)
			os.MkdirAll(outdir, os.ModePerm)
		} else if err != nil {
			return err
		} else {
			fmt.Printf("Output directory %s exists\n", outdir)

		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if isDir {
			return processDir(dir, 0)
		}

		return processFile(file, output)
	},
}

func init() {
	rootCmd.AddCommand(processCmd)
	processCmd.Flags().StringP("file", "f", "", "File to process")
	processCmd.Flags().StringP("output", "o", "Current Working Directory", "Output file or directory")
	processCmd.Flags().StringP("dir", "d", "", "Directory to process cannot be used with file")
	processCmd.Flags().BoolP("recursive", "r", false, "Process directories recursively")
	processCmd.Flags().StringP("recursive-depth", "n", "0", "Process directories recursively to a certain depth (0 = unlimited)")
	processCmd.Flags().IntP("width", "W", 0, "Width of the output image - zero means leave unchanged")
	processCmd.Flags().IntP("height", "H", 0, "Height of the output image - zero means leave unchanged")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// processCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// processCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func processDir(dir string, levels_traversed int) error {
	// walk the directory and look for image files
	// process each file
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && recursive && (levels_traversed < recursiveDepth || recursiveDepth == 0) {
			err := processDir(path, levels_traversed+1)
			if err != nil {
				return err
			}
		}
		// process the file
		if helpers.Contains(supportedExtensions, filepath.Ext(path)) {
			err := processFile(path, output)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return err
}

func processFile(file string, output string) error {
	// process the file
	im, err := imgconv.Open(file)
	if err != nil {
		return err
	}
	if isDir {
		output = path.Join(output, fmt.Sprintf("output_%s.jpg", filepath.Base(file)))
	}
	resizeOption := &imgconv.ResizeOption{}
	if width != 0 {
		resizeOption.Width = width
	}
	if height != 0 {
		resizeOption.Height = height
	}
	imgconv.Resize(im, resizeOption)
	return imgconv.Save(output, im, &imgconv.FormatOption{Format: imgconv.JPEG, EncodeOption: []imgconv.EncodeOption{imgconv.BackgroundColor(color.Black)}})
}
