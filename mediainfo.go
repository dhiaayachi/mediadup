package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var mediainfoBinary = flag.String("mediainfo-bin", "mediainfo", "the path to the mediainfo binary if it is not in the system $PATH")

type MediaInfo struct {
	Media Media `json:"media"`
}
type Extra struct {
}
type Track struct {
	Type                           string `json:"@type"`
	UniqueID                       string `json:"UniqueID,omitempty"`
	VideoCount                     string `json:"VideoCount,omitempty"`
	AudioCount                     string `json:"AudioCount,omitempty"`
	TextCount                      string `json:"TextCount,omitempty"`
	MenuCount                      string `json:"MenuCount,omitempty"`
	FileExtension                  string `json:"FileExtension,omitempty"`
	Format                         string `json:"Format,omitempty"`
	FormatVersion                  string `json:"Format_Version,omitempty"`
	FileSize                       string `json:"FileSize,omitempty"`
	Duration                       string `json:"Duration,omitempty"`
	OverallBitRateMode             string `json:"OverallBitRate_Mode,omitempty"`
	OverallBitRate                 string `json:"OverallBitRate,omitempty"`
	FrameRate                      string `json:"FrameRate,omitempty"`
	FrameCount                     string `json:"FrameCount,omitempty"`
	StreamSize                     string `json:"StreamSize,omitempty"`
	IsStreamable                   string `json:"IsStreamable,omitempty"`
	Title                          string `json:"Title,omitempty"`
	Movie                          string `json:"Movie,omitempty"`
	EncodedDate                    string `json:"Encoded_Date,omitempty"`
	FileModifiedDate               string `json:"File_Modified_Date,omitempty"`
	FileModifiedDateLocal          string `json:"File_Modified_Date_Local,omitempty"`
	EncodedApplication             string `json:"Encoded_Application,omitempty"`
	EncodedLibrary                 string `json:"Encoded_Library,omitempty"`
	StreamOrder                    string `json:"StreamOrder,omitempty"`
	ID                             string `json:"ID,omitempty"`
	FormatProfile                  string `json:"Format_Profile,omitempty"`
	FormatLevel                    string `json:"Format_Level,omitempty"`
	FormatTier                     string `json:"Format_Tier,omitempty"`
	CodecID                        string `json:"CodecID,omitempty"`
	BitRate                        string `json:"BitRate,omitempty"`
	Width                          string `json:"Width,omitempty"`
	Height                         string `json:"Height,omitempty"`
	SampledWidth                   string `json:"Sampled_Width,omitempty"`
	SampledHeight                  string `json:"Sampled_Height,omitempty"`
	PixelAspectRatio               string `json:"PixelAspectRatio,omitempty"`
	DisplayAspectRatio             string `json:"DisplayAspectRatio,omitempty"`
	FrameRateMode                  string `json:"FrameRate_Mode,omitempty"`
	ColorSpace                     string `json:"ColorSpace,omitempty"`
	ChromaSubsampling              string `json:"ChromaSubsampling,omitempty"`
	BitDepth                       string `json:"BitDepth,omitempty"`
	Delay                          string `json:"Delay,omitempty"`
	Default                        string `json:"Default,omitempty"`
	Forced                         string `json:"Forced,omitempty"`
	ColourDescriptionPresent       string `json:"colour_description_present,omitempty"`
	ColourDescriptionPresentSource string `json:"colour_description_present_Source,omitempty"`
	ColourRange                    string `json:"colour_range,omitempty"`
	ColourRangeSource              string `json:"colour_range_Source,omitempty"`
	ColourPrimaries                string `json:"colour_primaries,omitempty"`
	ColourPrimariesSource          string `json:"colour_primaries_Source,omitempty"`
	TransferCharacteristics        string `json:"transfer_characteristics,omitempty"`
	TransferCharacteristicsSource  string `json:"transfer_characteristics_Source,omitempty"`
	MatrixCoefficients             string `json:"matrix_coefficients,omitempty"`
	MatrixCoefficientsSource       string `json:"matrix_coefficients_Source,omitempty"`
	FormatCommercialIfAny          string `json:"Format_Commercial_IfAny,omitempty"`
	FormatSettingsMode             string `json:"Format_Settings_Mode,omitempty"`
	FormatSettingsEndianness       string `json:"Format_Settings_Endianness,omitempty"`
	FormatAdditionalFeatures       string `json:"Format_AdditionalFeatures,omitempty"`
	BitRateMode                    string `json:"BitRate_Mode,omitempty"`
	Channels                       string `json:"Channels,omitempty"`
	ChannelPositions     string `json:"ChannelPositions,omitempty"`
	ChannelLayout        string `json:"ChannelLayout,omitempty"`
	SamplesPerFrame      string `json:"SamplesPerFrame,omitempty"`
	SamplingRate         string `json:"SamplingRate,omitempty"`
	SamplingCount        string `json:"SamplingCount,omitempty"`
	CompressionMode      string `json:"Compression_Mode,omitempty"`
	DelaySource          string `json:"Delay_Source,omitempty"`
	StreamSizeProportion string `json:"StreamSize_Proportion,omitempty"`
	Language             string `json:"Language,omitempty"`
	ElementCount         string `json:"ElementCount,omitempty"`
	Extra                Extra  `json:"extra,omitempty"`
}
type Media struct {
	Ref   string  `json:"@ref"`
	Track []Track `json:"track"`
}


func IsInstalled() bool {
	cmd := exec.Command(*mediainfoBinary)
	err := cmd.Run()
	if err != nil {
		if strings.HasSuffix(err.Error(), "no such file or directory") ||
			strings.HasSuffix(err.Error(), "executable file not found in %PATH%") ||
			strings.HasSuffix(err.Error(), "executable file not found in $PATH") {
			return false
		} else if strings.HasPrefix(err.Error(), "exit status 255") {
			return true
		}
	}
	return true
}

func (info MediaInfo) IsMedia() bool {
	return len(info.Media.Track) > 0
}

func GetMediaInfo(fname string) ([]MediaInfo, error) {
	if !IsInstalled() {
		return nil, fmt.Errorf("must install mediainfo")
	}
	_, err := os.Stat(fname)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("media file not found")
	}

	var mInfo []MediaInfo

	err = filepath.Walk(fname, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			println(info.Name())
			out, err := exec.Command(*mediainfoBinary, "--Output=JSON", "-f", fname).Output()

			if err != nil {
				return nil
			}
			var i MediaInfo
			if err := json.Unmarshal(out, &i); err != nil {
				return nil
			}
			mInfo = append(mInfo, i)
		}
		return nil
	})
	if len(mInfo)==0 && err != nil {
		return nil, err
	}
	return mInfo, nil
}

func (m *MediaInfo) GetMovieTrackID()  (*Track,error){
	for _,t := range m.Media.Track {
		if t.Type == "General" {
			return &t, nil
		}
	}
	return nil,fmt.Errorf("no video track available")
}
