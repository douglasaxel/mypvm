package utils

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// descompactar chama a função de descompressão correta.
func Decompress(src, dest string) error {
	switch filepath.Ext(src) {
	case ".zip":
		return decompressZip(src, dest)
	case ".gz":
		return decompressTarGz(src, dest)
	default:
		return fmt.Errorf("tipo de arquivo desconhecido: %s", filepath.Ext(src))
	}
}

// descompactarZip extrai os arquivos de um .zip para um destino.
func decompressZip(src, dest string) error {
	reader, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer reader.Close()

	err = os.MkdirAll(dest, 0755)
	if err != nil {
		return err
	}

	for _, f := range reader.File {
		caminho := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(caminho, f.Mode())
			continue
		}

		outFile, err := os.OpenFile(caminho, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}
		defer outFile.Close()

		inFile, err := f.Open()
		if err != nil {
			return err
		}
		defer inFile.Close()

		_, err = io.Copy(outFile, inFile)
		if err != nil {
			return err
		}
	}
	return nil
}

// descompactarTarGz extrai os arquivos de um .tar.gz para um destino.
func decompressTarGz(src, dest string) error {
	file, err := os.Open(src)
	if err != nil {
		return err
	}
	defer file.Close()

	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	defer gzipReader.Close()

	tarReader := tar.NewReader(gzipReader)
	err = os.MkdirAll(dest, 0755)
	if err != nil {
		return err
	}

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		caminho := filepath.Join(dest, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(caminho, os.FileMode(header.Mode)); err != nil {
				return err
			}
		case tar.TypeReg:
			outFile, err := os.OpenFile(caminho, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.FileMode(header.Mode))
			if err != nil {
				return err
			}
			defer outFile.Close()

			if _, err := io.Copy(outFile, tarReader); err != nil {
				return err
			}
		}
	}
	return nil
}
