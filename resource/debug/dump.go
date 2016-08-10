/*
Copyright (C) 2016 Andreas T Jonsson

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package debug

import (
	"image/color"
	"image/png"
	"os"
	"path"

	"github.com/andreas-jonsson/openwar/resource"
)

func DumpImg(images resource.Images, pal color.Palette, p string) {
	outputPath := "img"
	if p != "" {
		outputPath = p
	}

	os.MkdirAll(outputPath, 0755)

	for file, image := range images {
		outfile, err := os.Create(path.Join(outputPath, file) + ".png")
		if err != nil {
			panic(err)
		}

		orgPal := image.Data.Palette
		image.Data.Palette = pal

		if err := png.Encode(outfile, image.Data); err != nil {
			panic(err)
		}

		outfile.Close()
		image.Data.Palette = orgPal
	}
}

func DumpArchive(arch *resource.Archive, p string) {
	outputPath := "archive"
	if p != "" {
		outputPath = p
	}

	os.MkdirAll(outputPath, 0755)

	for fileName, data := range arch.Files {
		fp, err := os.Create(path.Join(outputPath, fileName))
		if err != nil {
			panic(err)
		}

		if num, err := fp.Write(data); num != len(data) || err != nil {
			panic(err)
		}
		fp.Close()
	}
}