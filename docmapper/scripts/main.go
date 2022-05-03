package main

import "github.com/rs/zerolog/log"

func main() {

	err := Gen_AFC_DS01_Pain013_001_07("docmapper/defs/AFC-DS01-pain.013.001.07.yml")
	if err != nil {
		log.Error().Err(err).Msg("Creating Pain013")
	}

	err = Gen_AFC_DS10_Camt055_001_08("docmapper/defs/AFC-DS10-camt.055.001.08.yml")
	if err != nil {
		log.Error().Err(err).Msg("Creating Camt055")
	}

	log.Info().Msg("done")
}
